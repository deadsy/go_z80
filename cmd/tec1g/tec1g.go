//-----------------------------------------------------------------------------
/*

TEC-1G Emulation

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"

	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/six_digit"
	"github.com/deadsy/go_z80/memory"
	"github.com/deadsy/go_z80/z80"
)

//-----------------------------------------------------------------------------
// System Memory

const KiB = 1024
const chunkBits = 11 // 2 KiB chunks
const chunkSize = (1 << chunkBits)
const numChunks = (64 * KiB) / chunkSize

func chunkSelect(adr uint16) int { return int(adr >> chunkBits) }

type sysMemory struct {
	memmap [numChunks]z80.Memory
}

func newMemory() (*sysMemory, error) {
	// ROM
	rom := memory.New(11).ROM() // 2 KiB
	err := rom.LoadFile(0, "../../roms/mon1B.bin")
	if err != nil {
		return nil, err
	}
	// RAM
	ram := memory.New(11).RAM() // 2 KiB
	ram.Write8(0, 0xef)

	// Empty
	empty := memory.New(11).Empty() // 2 KiB

	return &sysMemory{
		memmap: [numChunks]z80.Memory{
			rom,   // 0x0000 - 0x07ff
			ram,   // 0x0800 - 0x0fff
			empty, // 0x1000
			empty, // 0x1800
			empty, // 0x2000
			empty, // 0x2800
			empty, // 0x3000
			empty, // 0x3800
			empty, // 0x4000
			empty, // 0x4800
			empty, // 0x5000
			empty, // 0x5800
			empty, // 0x6000
			empty, // 0x6800
			empty, // 0x7000
			empty, // 0x7800
			empty, // 0x8000
			empty, // 0x8800
			empty, // 0x9000
			empty, // 0x9800
			empty, // 0xa000
			empty, // 0xa800
			empty, // 0xb000
			empty, // 0xb800
			empty, // 0xc000
			empty, // 0xc800
			empty, // 0xd000
			empty, // 0xd800
			empty, // 0xe000
			empty, // 0xe800
			empty, // 0xf000
			empty, // 0xf800
		},
	}, nil
}

func (m *sysMemory) Read8(adr uint16) uint8 {
	return m.memmap[chunkSelect(adr)].Read8(adr)
}

func (m *sysMemory) Write8(adr uint16, val uint8) {
	m.memmap[chunkSelect(adr)].Write8(adr, val)
}

func (m *sysMemory) Read16(adr uint16) uint16 {
	return m.memmap[chunkSelect(adr)].Read16(adr)
}

func (m *sysMemory) Write16(adr uint16, val uint16) {
	m.memmap[chunkSelect(adr)].Write16(adr, val)
}

//-----------------------------------------------------------------------------

const keypadPort = 0x00  // keypad scan values
const digitPort = 0x01   // display digit enable
const segmentPort = 0x02 // display segment enable

const digitMask = uint8(0x3f)   // digits are bits 0..5
const speakerMask = uint8(0x80) // speaker/led is bit 7

type sysIO struct {
	display *six_digit.Display // 6 digit display
	led     *led.LED           // speaker led
	segment uint8              // latched segment enable
	digit   uint8              // latched digit enable
	speaker bool               // latched speaker/led enable
}

// Read8 reads a byte from an IO port.
func (io *sysIO) Read8(adr uint16) uint8 {
	adr &= 0xff
	switch adr {
	case keypadPort:
		//return keyAddress
		return 0 // keyGo
	}
	fmt.Printf("io.Read8 [%02x]\n", adr)
	return 0
}

// Write8 writes a byte to an IO port.
func (io *sysIO) Write8(adr uint16, val uint8) {
	adr &= 0xff
	switch adr {
	case digitPort:
		io.digit = val & digitMask
		io.speaker = (val & speakerMask) != 0
		io.display.Enable(io.digit, io.segment)
		io.led.Control(io.speaker)
		return
	case segmentPort:
		io.segment = val
		io.display.Enable(io.digit, io.segment)
		return
	}
	fmt.Printf("io.Write8 [%02x] = %02x\n", adr, val)
}

func newIO(display *six_digit.Display, led *led.LED) *sysIO {
	return &sysIO{
		display: display,
		led:     led,
	}
}

//-----------------------------------------------------------------------------

type Bus struct {
}

func newBus() *Bus {
	return &Bus{}
}

func (bus *Bus) ReadIV() uint8 {
	return 0xff
}

//-----------------------------------------------------------------------------
