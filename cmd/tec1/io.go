//-----------------------------------------------------------------------------
/*

TEC-1 Emulation

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/deadsy/go_z80/cmd/tec1/keypad"
	"github.com/deadsy/go_z80/device/array88"
	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/sixdigit"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------
// ports

const keypadPort = 0x00  // keypad scan values
const digitPort = 0x01   // display digit enable
const segmentPort = 0x02 // display segment enable
const x88Port = 0x03     // 8x8 X-axis display latch
const y88Port = 0x04     // 8x8 Y-axis display latch

// digitPort
const digitMask = byte(0x3f)     // D0..D5, digits
const speakerMask = byte(1 << 7) // D7, speaker/led

//-----------------------------------------------------------------------------

type ioDevices struct {
	display    *sixdigit.Display // 6 digit display
	ledSpeaker *led.LED          // speaker led
	ledArray   *array88.Array88  // 8x8 led array
	keypad     *keypad.Keypad    // 74c923 keypad
}

type sysIO struct {
	dev     *ioDevices
	segment uint8 // latched segment enable
	digit   uint8 // latched digit enable
	speaker bool  // latched speaker/led enable
}

// Read8 reads a byte from an IO port.
func (io *sysIO) Read8(adr uint16) uint8 {
	dev := io.dev
	adr &= 0xff
	switch adr {
	case keypadPort:
		return dev.keypad.Scan()
	}
	log.Printf("io.Read8 unknown port %02x", adr)
	return 0
}

// Write8 writes a byte to an IO port.
func (io *sysIO) Write8(adr uint16, val uint8) {
	dev := io.dev
	adr &= 0xff
	switch adr {
	case digitPort:
		io.digit = val & digitMask
		io.speaker = (val & speakerMask) != 0
		dev.display.Enable(io.digit, io.segment)
		dev.ledSpeaker.Control(io.speaker)
		return
	case segmentPort:
		io.segment = val
		dev.display.Enable(io.digit, io.segment)
		return
	case x88Port:
		dev.ledArray.WriteColumn(val)
		return
	case y88Port:
		dev.ledArray.WriteRow(val)
		return
	}
	log.Printf("io.Write8 [%02x] = %02x", adr, val)
}

func newIO(dev *ioDevices) *sysIO {
	return &sysIO{
		dev: dev,
	}
}

//-----------------------------------------------------------------------------
// ebiten api

func (io *sysIO) Update() {
	io.dev.display.Update()
	io.dev.ledSpeaker.Update()
	io.dev.ledArray.Update()
}

func (io *sysIO) Draw(screen *ebiten.Image) {
	io.dev.display.Draw(screen)
	io.dev.ledSpeaker.Draw(screen)
	io.dev.ledArray.Draw(screen)
}

//-----------------------------------------------------------------------------
