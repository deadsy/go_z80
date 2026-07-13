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

//-----------------------------------------------------------------------------
// ds1302 ram registers

// return true if the ram address is valid
func ramAddressValid(adr int) bool {
	return adr >= 0 && adr <= 30
}

const numRamRegisters = 31

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
	Mode12        bool          `toml:"mode_12"`        // 12 hour (am/pm) OR 24 hour clock
	WeekDayOffset int           `toml:"weekday_offset"` // weekday offset 0..6
	TimeOffset    time.Duration `toml:"time_offset"`    // time offset from UTC
	RAM           []byte        `toml:"ram"`            // ram contents
}

type RTC struct {
	enable bool       // is the rtc enabled?
	cancel func()     // cancel the background ticker
	lock   sync.Mutex // lock for ticker access
	// serial bus state
	clk     bool        // clock state
	state   serialState // serial transfer state
	command byte        // command register
	data    byte        // data register
	bits    int         // count of shift bits
	address int         // register address
	burst   bool        // burst mode
	out     bool        // output bit
	// register state
	mode12        bool                  // 12 hour (am/pm) OR 24 hour clock
	clockHalted   bool                  // is the clock halted
	writeProtect  bool                  // write protect the registers
	weekDayOffset int                   // weekday offset 0..6
	clockTime     time.Time             // rtc time
	ram           [numRamRegisters]byte // ram registers
}

func New(cfg *Config) (*RTC, error) {
	if cfg == nil {
		return nil, errors.New("no configuration")
	}
	rtc := &RTC{}

	// defaults
	rtc.reset()

	// start the background second ticker
	ctx, cancel := context.WithCancel(context.Background())
	rtc.cancel = cancel
	rtc.clockTime = time.Now().UTC().Add(cfg.TimeOffset)
	go backgroundTick(ctx, rtc)

	return rtc, nil
}

// return a config based on the current state
func (rtc *RTC) GetConfig() Config {
	return Config{
		Enable:        rtc.enable,
		Mode12:        rtc.mode12,
		WeekDayOffset: rtc.weekDayOffset,
		TimeOffset:    rtc.clockTime.Sub(time.Now().UTC()),
		RAM:           rtc.ram[:],
	}
}

func (rtc *RTC) Close() {
	// stop the background ticker
	rtc.cancel()
}

//-----------------------------------------------------------------------------

// update the rtc time every second
func backgroundTick(ctx context.Context, rtc *RTC) {
	for {
		now := time.Now()
		next := now.Truncate(time.Second).Add(time.Second)
		timer := time.NewTimer(next.Sub(now))
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			rtc.lock.Lock()
			if !rtc.clockHalted {
				rtc.clockTime = rtc.clockTime.Add(time.Second)
			}
			rtc.lock.Unlock()
		}
	}
}

//-----------------------------------------------------------------------------

// read a clock or ram register
func (rtc *RTC) read(adr int) byte {
	if rtc.command&rcBit == 0 {
		//log.Printf("clock read [%d] (%s)", adr, rtc.mode())
		switch adr {
		case clockSecond:
		case clockMinute:
		case clockHour:
		case clockDayOfMonth:
		case clockMonthOfYear:
		case clockDayOfWeek:
		case clockYear:
		case clockWriteProtect:
		case clockTrickleCharge:
		default:
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

// write a clock or ram register
func (rtc *RTC) write(adr int, data byte) {
	if rtc.command&rcBit == 0 {
		// check write protect
		if rtc.writeProtect && adr != clockWriteProtect {
			log.Printf("rtc: write protect enabled")
			return
		}
		//log.Printf("clock write [%d]=0x%02x (%s)", adr, data, rtc.mode())
		switch adr {
		case clockSecond:
		case clockMinute:
		case clockHour:
		case clockDayOfMonth:
		case clockMonthOfYear:
		case clockDayOfWeek:
		case clockYear:
		case clockWriteProtect:
		case clockTrickleCharge:
		default:
			log.Printf("rtc write: bad clock address %d", adr)
		}
	} else {
		// check write protect
		if rtc.writeProtect {
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
