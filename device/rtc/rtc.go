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

// latched write bits
const ceBit = byte(1 << 4) // active high
const clkBit = byte(1 << 6)
const inBit = byte(1 << 7) // maps to D7

// buffered read bits
const outBit = byte(1 << 0) // maps to D0

// per the tec1-g monitor mapping
var weekdayMap = [7]int{7, 1, 2, 3, 4, 5, 6}

const baseYear = 2000

//-----------------------------------------------------------------------------
// ds1302 commmand byte

const rwBit = byte(1 << 0)  // read/write
const rcBit = byte(1 << 6)  // ram/clock
const topBit = byte(1 << 7) // should always be 1

const burstAddress = 0x1f

// extract an address value from the command
func cmdAddress(cmd byte) int {
	return (int(cmd) >> 1) & 0x1f
}

// return true if the ram address is valid
func ramAddressValid(adr int) bool {
	return adr >= 0 && adr <= 30
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

const writeProtectBit = 7

//-----------------------------------------------------------------------------

func bcd1(x int) int {
	return x % 10
}

func bcd10(x int) int {
	return x / 10
}

//-----------------------------------------------------------------------------

type rtcState int

const (
	commandState   rtcState = iota
	byteReadState           // reading byte(s)
	byteWriteState          // writing byte(s)
)

type RTC struct {
	present      bool     // is the rtc present?
	clk          bool     // clock state
	state        rtcState // state
	command      byte     // command register
	data         byte     // data register
	bits         int      // count of shift bits
	address      int      // register address
	burst        bool     // burst mode
	out          bool     // output bit
	clock        [9]byte  // clock registers
	ram          [31]byte // ram registers
	writeProtect bool     // write protected
}

func New() (*RTC, error) {
	return &RTC{
		writeProtect: true,
	}, nil
}

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

func (rtc *RTC) getTime() {
	t := time.Now().UTC()
	var n int
	// clockSecond
	n = t.Second()
	rtc.clock[clockSecond] = byte((bcd10(n) << 4) | (bcd1(n) << 0))
	// clockMinute
	n = t.Minute()
	rtc.clock[clockMinute] = byte((bcd10(n) << 4) | (bcd1(n) << 0))
	// clockHour
	n = t.Hour()
	rtc.clock[clockHour] = byte((bcd10(n) << 4) | (bcd1(n) << 0))
	// clockDayOfMonth
	n = t.Day()
	rtc.clock[clockDayOfMonth] = byte((bcd10(n) << 4) | (bcd1(n) << 0))
	// clockMonthOfYear
	n = int(t.Month())
	rtc.clock[clockMonthOfYear] = byte((bcd10(n) << 4) | (bcd1(n) << 0))
	// clockDayOfWeek
	n = weekdayMap[t.Weekday()]
	rtc.clock[clockDayOfWeek] = byte(n)
	//clockYear
	n = t.Year() - baseYear
	rtc.clock[clockYear] = byte((bcd10(n) << 4) | (bcd1(n) << 0))
}

func (rtc *RTC) read(adr int) byte {
	if rtc.command&rcBit == 0 {
		//log.Printf("clock read [%d] (%s)\n", adr, rtc.mode())
		switch adr {
		case clockSecond, clockMinute, clockHour, clockDayOfMonth, clockMonthOfYear, clockDayOfWeek, clockYear:
			rtc.getTime()
			return rtc.clock[adr]
		case clockWriteProtect:
			if rtc.writeProtect {
				return 1 << writeProtectBit
			}
			return 0
		case clockTrickleCharge:
			return 0x5c // power-on state
		default:
			log.Printf("bad clock address %d\n", adr)
		}
	} else {
		//log.Printf("ram read [%d] (%s)\n", adr, rtc.mode())
		if ramAddressValid(adr) {
			return rtc.ram[adr]
		} else {
			log.Printf("bad ram address %d\n", adr)
		}
	}
	return 0
}

func (rtc *RTC) write(adr int, data byte) {
	if rtc.command&rcBit == 0 {
		// check write protect
		if rtc.writeProtect && adr != clockWriteProtect {
			log.Printf("write protect enabled\n")
		}
		//log.Printf("clock write [%d]=0x%02x (%s)\n", adr, data, rtc.mode())
		switch adr {
		case clockSecond:
		case clockMinute:
		case clockHour:
		case clockDayOfMonth:
		case clockMonthOfYear:
		case clockDayOfWeek:
		case clockYear:
		case clockWriteProtect:
			rtc.writeProtect = data&(1<<writeProtectBit) != 0
		case clockTrickleCharge:
			// whatever...
		default:
			log.Printf("bad clock address %d\n", adr)
		}

	} else {
		// check write protect
		if rtc.writeProtect {
			log.Printf("write protect enabled\n")
		}
		log.Printf("ram write [%d]=0x%02x (%s)\n", adr, data, rtc.mode())
		if ramAddressValid(adr) {
			rtc.ram[adr] = data
		} else {
			log.Printf("bad ram address %d\n", adr)
		}
	}
}

// enable the rtc
func (rtc *RTC) Enable() {
	rtc.present = true
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
