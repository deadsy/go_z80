//-----------------------------------------------------------------------------
/*

Jupiter ACE Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/deadsy/go_z80/cmd/jace/keyboard"
	"github.com/deadsy/go_z80/device/video"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

const keyboardPort = 0xfe // Matrix Keyboard Input

//-----------------------------------------------------------------------------

type ioDevices struct {
	keyboard *keyboard.Keyboard // matrix keyboard
	video    *video.Video       // video output
}

// System IO
type sysIO struct {
	dev     *ioDevices
	sys     *system // pointer back to system resources
	speaker bool    // latched speaker bit
}

func newIO(dev *ioDevices) *sysIO {
	return &sysIO{
		dev: dev,
	}
}

func (io *sysIO) setSystem(sys *system) {
	io.sys = sys
}

//-----------------------------------------------------------------------------

// Read8 reads a byte from an IO port.
func (io *sysIO) Read8(adr uint16) uint8 {
	dev := io.dev
	row := uint8(adr >> 8)
	adr &= 0xff
	switch adr {
	case keyboardPort:
		// a read on 0xfe drives the speaker output low
		io.speaker = false
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
	adr &= 0xff
	switch adr {
	case keyboardPort:
		// a write on 0xfe drives the speaker output high
		io.speaker = true
		return
	}
	log.Printf("io.Write8 [%02x] = %02x\n", adr, val)
}

//-----------------------------------------------------------------------------
// ebiten api

func (io *sysIO) Update() {
	if io.sys.mem.IsDirty() {
		// update the font atlas
		io.dev.video.Update()
		io.sys.mem.Clean()
	}
	io.dev.keyboard.Update()
}

func (io *sysIO) Draw(screen *ebiten.Image) {
	io.dev.video.Draw(screen)
}

//-----------------------------------------------------------------------------
