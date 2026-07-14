//-----------------------------------------------------------------------------
/*

TEC-1G Emulation

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"
	"log"

	"github.com/deadsy/go_z80/cmd/tec1g/keyboard"
	"github.com/deadsy/go_z80/device/ds1302"
	"github.com/deadsy/go_z80/device/hd44780"
	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/sixdigit"
	"github.com/deadsy/go_z80/memory"
	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

func boolToByte(x bool, val byte) byte {
	if x {
		return val
	}
	return 0
}

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
	rom := memory.New(14).ROM() // 16 KiB
	data, err := assets.ReadFile("assets/mon3_2025BC_16.bin")
	//data, err := assets.ReadFile("assets/DIAG-1G_CH24-11.bin")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded ROM: %w", err)
	}
	if err := rom.Load(0, data); err != nil {
		return nil, fmt.Errorf("failed to load ROM: %w", err)
	}

	// RAM
	ram := memory.New(15).RAM() // 32 KiB

	// Empty
	empty := memory.New(11).Empty() // 2 KiB

	return &sysMemory{
		memmap: [numChunks]z80.Memory{
			rom,   // 0x0000 - 0x07ff (shadow)
			ram,   // 0x0800
			ram,   // 0x1000
			ram,   // 0x1800
			ram,   // 0x2000
			ram,   // 0x2800
			ram,   // 0x3000
			ram,   // 0x3800
			ram,   // 0x4000
			ram,   // 0x4800
			ram,   // 0x5000
			ram,   // 0x5800
			ram,   // 0x6000
			ram,   // 0x6800
			ram,   // 0x7000
			ram,   // 0x7800
			empty, // 0x8000
			empty, // 0x8800
			empty, // 0x9000
			empty, // 0x9800
			empty, // 0xa000
			empty, // 0xa800
			empty, // 0xb000
			empty, // 0xb800
			rom,   // 0xc000
			rom,   // 0xc800
			rom,   // 0xd000
			rom,   // 0xd800
			rom,   // 0xe000
			rom,   // 0xe800
			rom,   // 0xf000
			rom,   // 0xf800
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

const keypadPort = 0x00   // keypad scan values
const digitPort = 0x01    // display digit enable
const segmentPort = 0x02  // display segment enable
const simpPort = 0x03     // General SIMP Input
const lcdCmdPort = 0x04   // LCD Display command
const x88Port = 0x05      // 8x8 X-axis display latch
const y88Port = 0x06      // 8x8 Y-axis display latch
const glcdPort0 = 0x07    // GLCD port
const lcdDataPort = 0x84  // LCD Display data
const glcdPort1 = 0x87    // GLCD port
const rtcPort = 0xfc      // GPIO Real Time Clock
const sdCardPort = 0xfd   // GPIO SD Card
const keyboardPort = 0xfe // Matrix Keyboard Input
const systemPort = 0xff   // System Latch

// digitPort
const digitMask = uint8(0x3f)      // digits are bits 0..5
const serialTxMask = uint8(1 << 6) // serialTx is bit 6
const speakerMask = uint8(1 << 7)  // speaker/led is bit 7

// simpPort
const simpKeyboard = byte(1 << 0) // 0 == encoder, 1 == matrix
// const simpProtect = byte(1 << 1) // 1 == protect memory
// const simpExpansion = byte(1 << 2)
// const simpExpand = byte(1 << 3)
// const simpCart = byte(1 << 4)
// const simpGimp = byte(1 << 5)
// const simpKey = byte(1 << 6)

type ioDevices struct {
	display    *sixdigit.Display  // 6 digit display
	ledSpeaker *led.LED           // speaker led
	ledHalt    *led.LED           // halt led
	lcd        *hd44780.LCD       // LCD
	keyboard   *keyboard.Keyboard // matrix keyboard
	rtc        *ds1302.RTC        // realtime clock
}

type sysIO struct {
	dev      *ioDevices
	segment  uint8 // latched segment enable
	digit    uint8 // latched digit enable
	speaker  bool  // latched speaker/led enable
	serialTx bool  // serial tx line
	serialRx bool  // serial rx line
}

func newIO(dev *ioDevices) *sysIO {
	return &sysIO{
		dev: dev,
	}
}

// Read8 reads a byte from an IO port.
func (io *sysIO) Read8(adr uint16) uint8 {
	dev := io.dev
	row := uint8(adr >> 8)
	adr &= 0xff
	switch adr {
	case keypadPort:
		return 0
	case lcdCmdPort:
		return dev.lcd.ReadCommand()
	case simpPort:
		// TODO
		return simpKeyboard | boolToByte(io.serialRx, 1<<7 /*D7*/)
	case rtcPort:
		return boolToByte(dev.rtc.Read(), 1<<0 /*D0*/)
	case sdCardPort:
		// TODO
		return 0
	case keyboardPort:
		code, err := dev.keyboard.Scan(row)
		if err != nil {
			log.Printf("keyboard scan error: %s\n", err)
		}
		return code
	}
	log.Printf("io.Read8 unknown port %02x\n", adr)
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
		io.serialTx = (val & serialTxMask) != 0
		dev.display.Enable(io.digit, io.segment)
		dev.ledSpeaker.Control(io.speaker)
		return
	case segmentPort:
		io.segment = val
		dev.display.Enable(io.digit, io.segment)
		return
	case lcdCmdPort:
		dev.lcd.WriteCommand(val)
		return
	case x88Port:
		// TODO
		return
	case y88Port:
		// TODO
		return
	case glcdPort0, glcdPort1:
		// TODO
		return
	case lcdDataPort:
		dev.lcd.WriteData(val)
		return
	case rtcPort:
		ce := val&(1<<4) != 0  // D4 active high
		clk := val&(1<<6) != 0 // D6
		in := val&(1<<7) != 0  // D7
		dev.rtc.Write(ce, clk, in)
		return
	case sdCardPort:
		// TODO
		return
	case systemPort:
		// TODO
		return

	}
	log.Printf("io.Write8 [%02x] = %02x\n", adr, val)
}

func (io *sysIO) Update() {
	io.dev.display.Update()
	io.dev.ledSpeaker.Update()
	io.dev.ledHalt.Update()
	io.dev.lcd.Update()
	io.dev.keyboard.Update()
}

func (io *sysIO) Draw(screen *ebiten.Image) {
	io.dev.display.Draw(screen)
	io.dev.ledSpeaker.Draw(screen)
	io.dev.ledHalt.Draw(screen)
	io.dev.lcd.Draw(screen)
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
