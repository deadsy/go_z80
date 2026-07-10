//-----------------------------------------------------------------------------
/*

Emulate the TEC-1G RTC Board (DS1302)

*/
//-----------------------------------------------------------------------------

package rtc

import "log"

//-----------------------------------------------------------------------------

// latched write bits
const ceBit = byte(1 << 4) // active high
const weBit = byte(1 << 5) // active low
const clkBit = byte(1 << 6)
const inBit = byte(1 << 7) // maps to D7

// buffered read bits
const outBit = byte(1 << 0) // maps to D0

//-----------------------------------------------------------------------------
// ds1302 commmand byte

const rwBit = byte(1 << 0)  // read/write
const rcBit = byte(1 << 6)  // ram/clock
const topBit = byte(1 << 7) // should always be 1

// extract an address value from the command
func cmdAddress(cmd byte) int {
	return (int(cmd) >> 1) & 0x1f
}

// return true if the ram address is valid
func ramAddressValid(adr int) bool {
	return adr >= 0 && adr <= 30
}

// return true if the clock address is valid
func clockAddressValid(adr int) bool {
	return adr >= 0 && adr <= 8
}

//-----------------------------------------------------------------------------

type rtcState int

const (
	commandState   rtcState = iota
	byteReadState           // reading byte(s)
	byteWriteState          // writing byte(s)
)

type RTC struct {
	present bool     // is the rtc present?
	clk     bool     // clock state
	state   rtcState // state
	command byte     // command register
	data    byte     // data register
	bits    int      // count of shift bits
	address int      // register address
	burst   bool     // burst mode
	out     bool     // output bit
	ram     [31]byte // ram
}

func New() (*RTC, error) {
	return &RTC{}, nil
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
		// clock 0..7
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

func (rtc *RTC) read(adr int) byte {
	if rtc.command&rcBit == 0 {
		// clock
		log.Printf("clock read [%d] (%s)\n", adr, rtc.mode())
	} else {
		log.Printf("ram read [%d] (%s)\n", adr, rtc.mode())
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
		// clock
		log.Printf("clock write [%d]=0x%02x (%s)\n", adr, data, rtc.mode())
	} else {
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
				rtc.address = cmdAddress(rtc.command)
				rtc.bits = 0
				if rtc.command&rwBit == 0 {
					// writing
					if rtc.address == 0x1f {
						// bursting write
						rtc.burst = true
						rtc.address = 0
					}
					rtc.state = byteWriteState
				} else {
					// reading
					if rtc.address == 0x1f {
						// bursting read
						rtc.burst = true
						rtc.address = 0
					}
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
				if rtc.burst {
					rtc.write(rtc.address, rtc.data)
					rtc.bits = 0
					rtc.address = rtc.burstIncrement()
				} else {
					rtc.write(rtc.address, rtc.data)
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
