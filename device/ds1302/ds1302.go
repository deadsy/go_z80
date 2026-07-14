//-----------------------------------------------------------------------------
/*

DS1302 Real Time Clock Emulation

*/
//-----------------------------------------------------------------------------

package ds1302

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

//-----------------------------------------------------------------------------

func intToBcd(x int) byte {
	return byte(((x / 10) << 4) | (x % 10))
}

func bcdToInt(x byte) int {
	return int((x>>4)*10) + int(x&15)
}

func boolToByte(x bool, val byte) byte {
	if x {
		return val
	}
	return 0
}

//-----------------------------------------------------------------------------
// ds1302 clock registers

const (
	clockSecond        = 0
	clockMinute        = 1
	clockHour          = 2
	clockDayOfMonth    = 3
	clockMonthOfYear   = 4
	clockDayOfWeek     = 5
	clockYear          = 6
	clockWriteProtect  = 7
	clockTrickleCharge = 8
)

const writeProtectEnabled = byte(1 << 7) // in clockWriteProtect
const clockHalted = byte(1 << 7)         // in clockSecond
const mode12Hour = byte(1 << 7)          // in clockHour

// return true if the clock address is valid
func clockAddressValid(adr int) bool {
	return adr >= clockSecond && adr <= clockTrickleCharge
}

// mask the valid bits of the clock registers
var clockMask = [9]byte{0xff, 0x7f, 0xbf, 0x3f, 0x1f, 0x07, 0xff, 0x80, 0xff}

//-----------------------------------------------------------------------------
// ds1302 ram registers

const numRamRegisters = 31

// return true if the ram address is valid
func ramAddressValid(adr int) bool {
	return adr >= 0 && adr < numRamRegisters
}

//-----------------------------------------------------------------------------
// ds1302 command byte

const rwBit = byte(1 << 0)  // read/write
const rcBit = byte(1 << 6)  // ram/clock
const topBit = byte(1 << 7) // should always be 1

const burstAddress = 0x1f

// extract an address value from the command
func cmdAddress(cmd byte) int {
	return (int(cmd) >> 1) & 0x1f
}

//-----------------------------------------------------------------------------

type serialState int

const (
	commandState   serialState = iota // getting the command
	byteReadState                     // reading byte(s)
	byteWriteState                    // writing byte(s)
)

type Config struct {
	Enable        bool          `toml:"enable"`         // is the rtc enabled?
	BaseYear      int           `toml:"base_year"`      // base year
	Mode12        bool          `toml:"mode_12"`        // 12 hour (am/pm) OR 24 hour clock
	WeekDayOffset int           `toml:"weekday_offset"` // weekday offset 0..6
	TimeOffset    time.Duration `toml:"time_offset"`    // time offset from UTC
	RAM           []byte        `toml:"ram"`            // ram contents
}

type RTC struct {
	// configuration
	enable        bool                  // is the rtc enabled?
	baseYear      int                   // base year
	mode12        bool                  // 12 hour (am/pm) OR 24 hour clock
	weekDayOffset int                   // weekday offset 0..6
	timeOffset    time.Duration         // configured time offset
	ram           [numRamRegisters]byte // ram registers

	// serial bus state
	state   serialState // serial transfer state
	clk     bool        // clock state
	command byte        // command register
	data    byte        // data register
	bits    int         // count of shift bits
	address int         // register address
	burst   bool        // burst mode
	out     bool        // output bit

	// register state
	clockHalted   bool    // is the clock halted (must hold lock!)
	writeProtect  bool    // write protect the registers
	offsetDirty   bool    // has the set time been modified?
	trickleCharge byte    // trickle charge register
	clockTime     rtcTime // rtc time (must hold lock!)

	// background ticker
	cancel func()        // cancel the background ticker
	lock   sync.Mutex    // lock for ticker access
	resync chan struct{} // resync the background ticker

}

func New(cfg *Config) (*RTC, error) {
	if cfg == nil {
		return nil, errors.New("no configuration")
	}
	rtc := &RTC{
		enable:        cfg.Enable,
		baseYear:      cfg.BaseYear,
		mode12:        cfg.Mode12,
		weekDayOffset: cfg.WeekDayOffset,
		timeOffset:    cfg.TimeOffset,
		clockHalted:   false,
		writeProtect:  true,
		trickleCharge: 0x5c, // power-on state
	}

	// copy the ram
	copy(rtc.ram[:], cfg.RAM)

	// serial bus reset
	rtc.reset()

	if rtc.enable {
		// start the background second ticker
		rtc.resync = make(chan struct{}, 1) // buffered so unhalt never blocks
		ctx, cancel := context.WithCancel(context.Background())
		rtc.cancel = cancel
		t := time.Now().UTC().Add(rtc.timeOffset)
		rtc.clockTime = newRtcTime(t, rtc.baseYear, rtc.weekDayOffset)
		go backgroundTick(ctx, rtc)
	}

	return rtc, nil
}

// return a config based on the current state
func (rtc *RTC) GetConfig() Config {

	rtc.lock.Lock()
	defer rtc.lock.Unlock()

	if rtc.enable && rtc.offsetDirty {
		rtc.timeOffset = rtc.clockTime.getTime(rtc.baseYear).Sub(time.Now().UTC())
	}

	// copy the ram
	ram := make([]byte, len(rtc.ram))
	copy(ram, rtc.ram[:])

	return Config{
		Enable:        rtc.enable,
		BaseYear:      rtc.baseYear,
		Mode12:        rtc.mode12,
		WeekDayOffset: rtc.weekDayOffset,
		TimeOffset:    rtc.timeOffset,
		RAM:           ram,
	}
}

func (rtc *RTC) Close() {
	// stop the background ticker
	if rtc.enable {
		rtc.cancel()
	}
}

//-----------------------------------------------------------------------------

type rtcTime struct {
	second int // 0..59
	minute int // 0..59
	hour   int // 0..23
	day    int // 0..x (x = 27,28,29,30)
	month  int // 0..11
	year   int // 0..99
	dow    int // 0..6 day of week
}

// convert a go time to an rtc time
func newRtcTime(t time.Time, baseYear, dowOffset int) rtcTime {
	return rtcTime{
		second: t.Second(),
		minute: t.Minute(),
		hour:   t.Hour(),
		day:    t.Day() - 1,
		month:  int(t.Month()) - 1,
		year:   t.Year() - baseYear,
		dow:    (int(t.Weekday()) + dowOffset) % daysPerWeek,
	}
}

// convert the rtc time into a go time
func (t *rtcTime) getTime(baseYear int) time.Time {
	year := t.year + baseYear
	month := time.Month(t.month + 1)
	day := t.day + 1
	hour := t.hour
	minute := t.minute
	second := t.second
	return time.Date(year, month, day, hour, minute, second, 0, time.UTC)
}

// Return the number of days (28,29,30,31) in a month (0..11)
// Leap year checking is simple, supposedly matching the ds1302.
func (t *rtcTime) daysPerMonth() int {
	switch t.month {
	case 0, 2, 4, 6, 7, 9, 11: // jan mar may jul aug oct dec
		return 31
	case 3, 5, 8, 10: // apr jun sep nov
		return 30
	case 1: // feb
		if t.year&3 == 0 {
			return 29
		}
		return 28
	}
	panic("bad month")
}

const secondsPerMinute = 60
const minutesPerHour = 60
const hoursPerDay = 24
const monthsPerYear = 12
const yearsPerCentury = 100
const daysPerWeek = 7

// increment the rtc time by 1 second
func (t *rtcTime) increment() {
	t.second += 1
	if t.second < secondsPerMinute {
		return
	}
	t.second = 0
	t.minute += 1
	if t.minute < minutesPerHour {
		return
	}
	t.minute = 0
	t.hour += 1
	if t.hour < hoursPerDay {
		return
	}
	t.hour = 0
	t.day += 1
	t.dow = (t.dow + 1) % daysPerWeek
	if t.day < t.daysPerMonth() {
		return
	}
	t.day = 0
	t.month += 1
	if t.month < monthsPerYear {
		return
	}
	t.month = 0
	t.year += 1
	if t.year < yearsPerCentury {
		return
	}
	t.year = 0
}

//-----------------------------------------------------------------------------
// background once-a-second ticker

// update the rtc time every second
func backgroundTick(ctx context.Context, rtc *RTC) {
	now := time.Now()
	next := now.Truncate(time.Second).Add(time.Second)
	timer := time.NewTimer(next.Sub(now))
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-rtc.resync:
			// resync: next tick is exactly 1 second from the un-halt event
			if !timer.Stop() {
				// drain a pending fire if there is one
				select {
				case <-timer.C:
				default:
				}
			}
			timer.Reset(time.Second)

		case <-timer.C:
			rtc.lock.Lock()
			if !rtc.clockHalted {
				rtc.clockTime.increment()
			}
			rtc.lock.Unlock()

			// schedule the next aligned tick
			now := time.Now()
			next := now.Truncate(time.Second).Add(time.Second)
			timer.Reset(next.Sub(now))
		}
	}
}

//-----------------------------------------------------------------------------

// check that n is in [min, max]
func checkRange(n, min, max int, msg string) int {
	if n < min || n > max {
		log.Printf("ds1302: bad %s value: %d", msg, n)
		return min
	}
	return n
}

// encode 0..23 to the clockHour byte
func encodeHour(n int, mode12 bool) byte {
	val := boolToByte(mode12, mode12Hour)
	if val != 0 {
		// 12 hour mode
		if n >= 12 /* 12 == 12pm */ {
			val |= (1 << 5) // PM
		}
		if n >= 13 /* 13 == 1pm */ {
			n -= 12
		}
	}
	return val | intToBcd(n)
}

// decode the clockHour byte to 0..23
func decodeHour(val byte) int {
	var pm bool
	if val&mode12Hour != 0 {
		// 12 hour clock
		pm = val&(1<<5) != 0
		val &= 0x1f
	}
	hour := bcdToInt(val)
	if pm && hour != 12 /* 12pm == 12 */ {
		hour += 12
	}
	return hour
}

// read a clock register
func (rtc *RTC) readClock(adr int) byte {
	//log.Printf("clock read [%d] (%s)", adr, rtc.mode())
	rtc.lock.Lock()
	defer rtc.lock.Unlock()

	switch adr {
	case clockSecond:
		return boolToByte(rtc.clockHalted, clockHalted) | intToBcd(rtc.clockTime.second)
	case clockMinute:
		return intToBcd(rtc.clockTime.minute)
	case clockHour:
		return encodeHour(rtc.clockTime.hour, rtc.mode12)
	case clockDayOfMonth:
		return intToBcd(rtc.clockTime.day + 1)
	case clockMonthOfYear:
		return intToBcd(rtc.clockTime.month + 1)
	case clockDayOfWeek:
		return byte(rtc.clockTime.dow + 1)
	case clockYear:
		return intToBcd(rtc.clockTime.year)
	case clockWriteProtect:
		return boolToByte(rtc.writeProtect, writeProtectEnabled)
	case clockTrickleCharge:
		return rtc.trickleCharge
	}

	return 0
}

// write a clock register
func (rtc *RTC) writeClock(adr int, data byte) {
	//log.Printf("clock write [%d]=0x%02x (%s)", adr, data, rtc.mode())

	// zero out any illegal bits
	data &= clockMask[adr]

	rtc.lock.Lock()
	defer rtc.lock.Unlock()

	wasHalted := rtc.clockHalted

	switch adr {
	case clockSecond:
		rtc.offsetDirty = true
		n := bcdToInt(data &^ clockHalted)
		rtc.clockTime.second = checkRange(n, 0, 59, "second")
		rtc.clockHalted = (data & clockHalted) != 0
	case clockMinute:
		rtc.offsetDirty = true
		n := bcdToInt(data)
		rtc.clockTime.minute = checkRange(n, 0, 59, "minute")
	case clockHour:
		rtc.offsetDirty = true
		n := decodeHour(data)
		rtc.clockTime.hour = checkRange(n, 0, 23, "hour")
		rtc.mode12 = (data & mode12Hour) != 0
	case clockDayOfMonth:
		rtc.offsetDirty = true
		n := bcdToInt(data) - 1
		rtc.clockTime.day = checkRange(n, 0, 30, "day")
	case clockMonthOfYear:
		rtc.offsetDirty = true
		n := bcdToInt(data) - 1
		rtc.clockTime.month = checkRange(n, 0, 11, "month")
	case clockDayOfWeek:
		n := int((data & 7) - 1)
		rtc.clockTime.dow = checkRange(n, 0, 6, "day of week")
	case clockYear:
		rtc.offsetDirty = true
		n := bcdToInt(data)
		rtc.clockTime.year = checkRange(n, 0, 99, "year")
	case clockWriteProtect:
		rtc.writeProtect = (data & writeProtectEnabled) != 0
	case clockTrickleCharge:
		rtc.trickleCharge = data
	}

	// are we un-halting the clock?
	if wasHalted && !rtc.clockHalted {
		// non-blocking send: if a resync is pending it's ok
		select {
		case rtc.resync <- struct{}{}:
		default:
		}
	}
}

// read a clock or ram register
func (rtc *RTC) read(adr int) byte {
	if rtc.command&rcBit == 0 {
		if clockAddressValid(adr) {
			return rtc.readClock(adr)
		} else {
			log.Printf("ds1302 read: bad clock address %d", adr)
		}
	} else {
		//log.Printf("ram read [%d] (%s)", adr, rtc.mode())
		if ramAddressValid(adr) {
			return rtc.ram[adr]
		} else {
			log.Printf("ds1302 read: bad ram address %d", adr)
		}
	}
	return 0
}

// write a clock or ram register
func (rtc *RTC) write(adr int, data byte) {
	if rtc.command&rcBit == 0 {
		// check write protect
		if rtc.writeProtect && adr != clockWriteProtect {
			log.Printf("ds1302: write protect enabled")
			return
		}
		if clockAddressValid(adr) {
			rtc.writeClock(adr, data)
		} else {
			log.Printf("ds1302 write: bad clock address %d", adr)
		}
	} else {
		// check write protect
		if rtc.writeProtect {
			log.Printf("ds1302: write protect enabled")
			return
		}
		//log.Printf("ram write [%d]=0x%02x (%s)", adr, data, rtc.mode())
		if ramAddressValid(adr) {
			rtc.ram[adr] = data
		} else {
			log.Printf("ds1302 write: bad ram address %d", adr)
		}
	}
}

//-----------------------------------------------------------------------------

func (rtc *RTC) mode() string {
	if rtc.burst {
		return "burst"
	}
	return "single"
}

// reset the rtc command/byte state variables
func (rtc *RTC) reset() {
	rtc.clk = false
	rtc.state = commandState
	rtc.command = 0
	rtc.data = 0
	rtc.bits = 0
	rtc.address = 0
	rtc.burst = false
	rtc.out = false
}

// increment the address used in burst mode
func (rtc *RTC) burstIncrement() int {
	adr := rtc.address + 1
	if rtc.command&rcBit == 0 {
		// clock 0..7 (8 is skipped)
		if adr == 8 {
			adr = 0
		}
	} else {
		// ram 0..30
		if adr == 31 {
			adr = 0
		}
	}
	return adr
}

// read from the RTC output bit
func (rtc *RTC) Read() bool {
	if !rtc.enable {
		// no device present
		return true
	}
	return rtc.out
}

// write a value to the RTC board latch
func (rtc *RTC) Write(chipEnable, serialClock, inputBit bool) {
	if !rtc.enable {
		// no device present
		return
	}

	if !chipEnable {
		// chip enable is low
		rtc.reset()
		return
	}

	// chip is enabled
	risingEdge := !rtc.clk && serialClock
	fallingEdge := rtc.clk && !serialClock
	rtc.clk = serialClock

	switch rtc.state {
	case commandState:
		if risingEdge {
			rtc.command >>= 1
			if inputBit {
				rtc.command |= 0x80
			}
			rtc.bits += 1
			if rtc.bits == 8 {
				if rtc.command&topBit == 0 {
					log.Printf("ds1302: bad command byte %02x", rtc.command)
					rtc.reset()
					return
				}
				rtc.bits = 0
				rtc.address = cmdAddress(rtc.command)
				if rtc.address == burstAddress {
					// bursting write
					rtc.burst = true
					rtc.address = 0
				}
				if rtc.command&rwBit == 0 {
					// writing
					rtc.state = byteWriteState
				} else {
					// reading
					rtc.state = byteReadState
					rtc.data = rtc.read(rtc.address)
				}
			}
		}
	case byteWriteState:
		if risingEdge {
			rtc.data >>= 1
			if inputBit {
				rtc.data |= 0x80
			}
			rtc.bits += 1
			if rtc.bits == 8 {
				rtc.write(rtc.address, rtc.data)
				if rtc.burst {
					rtc.bits = 0
					rtc.address = rtc.burstIncrement()
				} else {
					rtc.reset()
				}
			}
		}
	case byteReadState:
		if fallingEdge {
			rtc.out = rtc.data&1 != 0
			rtc.data >>= 1
			rtc.bits += 1
			if rtc.bits == 8 {
				if rtc.burst {
					rtc.bits = 0
					rtc.address = rtc.burstIncrement()
					rtc.data = rtc.read(rtc.address)
				} else {
					rtc.reset()
				}
			}
		}
	}
}

//-----------------------------------------------------------------------------
