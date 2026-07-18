//-----------------------------------------------------------------------------
/*

Jupiter ACE Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/deadsy/go_z80/cmd/jace/keyboard"
)

//-----------------------------------------------------------------------------

const keyboardPort = 0xfe // Matrix Keyboard Input

// System IO
type sysIO struct {
	keyboard *keyboard.Keyboard // matrix keyboard
	speaker  bool               // latched speaker bit
}

// Read8 reads a byte from an IO port.
func (io *sysIO) Read8(adr uint16) uint8 {
	row := uint8(adr >> 8)
	adr &= 0xff
	switch adr {
	case keyboardPort:
		// a read on 0xfe drives the speaker output low
		io.speaker = false
		code, err := io.keyboard.Scan(row)
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

func newIO(keyboard *keyboard.Keyboard) *sysIO {
	return &sysIO{
		keyboard: keyboard,
	}
}

//-----------------------------------------------------------------------------
