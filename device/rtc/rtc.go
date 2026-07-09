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

type RTC struct {
	enabled bool // is the rtc enabled?
}

func New() (*RTC, error) {
	return &RTC{}, nil
}

// enable the rtc
func (rtc *RTC) Enable() {
	rtc.enabled = true
}

// disbale the rtc
func (rtc *RTC) Disable() {
	rtc.enabled = false
}

// read a value from the RTC board buffer
func (rtc *RTC) Read() byte {
	if !rtc.enabled {
		// no device present
		return 0xff
	}
	log.Printf("Read\n")
	return 0
}

// write a value to the RTC board latch
func (rtc *RTC) Write(data byte) {
	if !rtc.enabled {
		// no device present
		return
	}
	chipEnable := (data & ceBit) != 0
	writeEnable := (data & weBit) != 0
	clock := (data & clkBit) != 0
	in := (data & inBit) != 0

	log.Printf("CE %t WE %t CLK %t IN %t\n", chipEnable, writeEnable, clock, in)
}

//-----------------------------------------------------------------------------
