//-----------------------------------------------------------------------------
/*

Emulate the TEC-1G RTC Board (DS1302)

*/
//-----------------------------------------------------------------------------

package rtc

import (
	"log"
	"time"
)

//-----------------------------------------------------------------------------

func intToBcd(x int) byte {
	return byte(((x / 10) << 4) | (x % 10))
}

func bcdToInt(x byte) int {
	return int((x>>4)*10) + int(x&15)
}

//-----------------------------------------------------------------------------

// latched write bits
const ceBit = byte(1 << 4) // active high
const clkBit = byte(1 << 6)
const inBit = byte(1 << 7) // maps to D7

// buffered read bits
const outBit = byte(1 << 0) // maps to D0

// per the tec-1g monitor mapping
var weekdayMap = [7]int{7, 1, 2, 3, 4, 5, 6}

const baseYear = 2000

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

const numClockRegisters = 9

const writeProtectEnabled = byte(1 << 7) // in clockWriteProtect
const clockHalted = byte(1 << 7)         // in clockSecond
const mode12Hour = byte(1 << 7)          // in clockHour

// return true if the clock address is valid
func clockAddressValid(adr int) bool {
	return adr >= clockSecond && adr <= clockTrickleCharge
}

// mask the valid bits of the clock registers
var clockMask = [numClockRegisters]byte{0xff, 0x7f, 0xbf, 0x3f, 0x1f, 0x07, 0xff, 0x80, 0xff}

//-----------------------------------------------------------------------------
// ds1302 ram registers

// return true if the ram address is valid
func ramAddressValid(adr int) bool {
	return adr >= 0 && adr <= 30
}

const numRamRegisters = 31

//-----------------------------------------------------------------------------

type rtcState int

const (
	commandState   rtcState = iota
	byteReadState           // reading byte(s)
	byteWriteState          // writing byte(s)
)

type RTC struct {
	present bool                    // is the rtc present?
	clk     bool                    // clock state
	state   rtcState                // state
	command byte                    // command register
	data    byte                    // data register
	bits    int                     // count of shift bits
	address int                     // register address
	burst   bool                    // burst mode
	out     bool                    // output bit
	offset  time.Duration           // rtc time - utc time
	clock   [numClockRegisters]byte // clock registers
	ram     [numRamRegisters]byte   // ram registers
}

func New() (*RTC, error) {
	rtc := &RTC{
		present: true,
	}
	rtc.reset()
	// set power-on values
	rtc.clock[clockWriteProtect] = writeProtectEnabled
	rtc.clock[clockTrickleCharge] = 0x5c
	rtc.setClock(time.Now().UTC().Add(rtc.offset))
	return rtc, nil
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

// disable the rtc
func (rtc *RTC) Disable() {
	rtc.present = false
}

//-----------------------------------------------------------------------------

// is the clock running?
func (rtc *RTC) isClockRunning() bool {
	return rtc.clock[clockSecond]&clockHalted == 0
}

// is the write protect mode enabled?
func (rtc *RTC) isWriteProtected() bool {
	return rtc.clock[clockWriteProtect]&writeProtectEnabled != 0
}

// encode 0..23 to the clockHour byte
func (rtc *RTC) encodeHour(n int) byte {
	val := rtc.clock[clockHour] & mode12Hour
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
func (rtc *RTC) decodeHour(val byte) int {
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

// set the clock registers from a time value
func (rtc *RTC) setClock(t time.Time) {
	var n int
	// clockSecond
	n = t.Second()
	rtc.clock[clockSecond] &= clockHalted
	rtc.clock[clockSecond] |= intToBcd(n)
	// clockMinute
	n = t.Minute()
	rtc.clock[clockMinute] = intToBcd(n)
	// clockHour
	n = t.Hour()
	rtc.clock[clockHour] = rtc.encodeHour(n)
	// clockDayOfMonth
	n = t.Day()
	rtc.clock[clockDayOfMonth] = intToBcd(n)
	// clockMonthOfYear
	n = int(t.Month())
	rtc.clock[clockMonthOfYear] = intToBcd(n)
	// clockDayOfWeek
	n = weekdayMap[t.Weekday()]
	rtc.clock[clockDayOfWeek] = byte(n)
	//clockYear
	n = t.Year() - baseYear
	rtc.clock[clockYear] = intToBcd(n)
}

// get the time value in the clock registers
func (rtc *RTC) getClock() time.Time {
	var n byte
	// year
	n = rtc.clock[clockYear]
	year := baseYear + bcdToInt(n)
	// month
	n = rtc.clock[clockMonthOfYear]
	month := time.Month(bcdToInt(n))
	// day
	n = rtc.clock[clockDayOfMonth]
	day := bcdToInt(n)
	// hour
	n = rtc.clock[clockHour]
	hour := rtc.decodeHour(n)
	// minute
	n = rtc.clock[clockMinute]
	min := bcdToInt(n)
	// second
	n = rtc.clock[clockSecond] &^ clockHalted
	sec := bcdToInt(n)

	return time.Date(year, month, day, hour, min, sec, 0, time.UTC)
}

func (rtc *RTC) read(adr int) byte {
	if rtc.command&rcBit == 0 {
		//log.Printf("clock read [%d] (%s)", adr, rtc.mode())
		if rtc.isClockRunning() {
			// update the clock registers
			rtc.setClock(time.Now().UTC().Add(rtc.offset))
		}
		if clockAddressValid(adr) {
			return rtc.clock[adr]
		} else {
			log.Printf("rtc read: bad clock address %d", adr)
		}
	} else {
		//log.Printf("ram read [%d] (%s)", adr, rtc.mode())
		if ramAddressValid(adr) {
			return rtc.ram[adr]
		} else {
			log.Printf("rtc read: bad ram address %d", adr)
		}
	}
	return 0
}

func (rtc *RTC) write(adr int, data byte) {
	if rtc.command&rcBit == 0 {
		// check write protect
		if rtc.isWriteProtected() && adr != clockWriteProtect {
			log.Printf("rtc: write protect enabled")
			return
		}
		//log.Printf("clock write [%d]=0x%02x (%s)", adr, data, rtc.mode())
		if clockAddressValid(adr) {
			rtc.clock[adr] = data & clockMask[adr]
			if rtc.isClockRunning() {
				rtc.offset = rtc.getClock().Sub(time.Now().UTC())
			}
		} else {
			log.Printf("rtc write: bad clock address %d", adr)
		}
	} else {
		// check write protect
		if rtc.isWriteProtected() {
			log.Printf("rtc: write protect enabled")
			return
		}
		//log.Printf("ram write [%d]=0x%02x (%s)", adr, data, rtc.mode())
		if ramAddressValid(adr) {
			rtc.ram[adr] = data
		} else {
			log.Printf("rtc write: bad ram address %d", adr)
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

// read a value from the RTC board buffer
func (rtc *RTC) Read() byte {
	if !rtc.present {
		// no device present
		return 0xff
	}
	if rtc.out {
		return outBit
	}
	return 0
}

// write a value to the RTC board latch
func (rtc *RTC) Write(data byte) {
	if !rtc.present {
		// no device present
		return
	}

	if data&ceBit == 0 {
		// chip enable is low
		rtc.reset()
		return
	}

	// chip is enabled
	clk := (data & clkBit) != 0
	risingEdge := !rtc.clk && clk
	fallingEdge := rtc.clk && !clk
	rtc.clk = clk

	switch rtc.state {
	case commandState:
		if risingEdge {
			rtc.command >>= 1
			if (data & inBit) != 0 {
				rtc.command |= 0x80
			}
			rtc.bits += 1
			if rtc.bits == 8 {
				if rtc.command&topBit == 0 {
					log.Printf("bad command byte %02x\n", rtc.command)
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
			if (data & inBit) != 0 {
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
