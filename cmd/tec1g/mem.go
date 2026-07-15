//-----------------------------------------------------------------------------
/*

TEC-1G Memory Emulation

Note:

The reset address for the CPU is 0, so typically ROM will be present at this
address upon power up. The original TEC-1 also had 2 KiB of ROM mapped to
the 0-0x7ff address range.

In the TEC-1G the 0 address is nominally RAM memory, but the first 2KiB of
the ROM can be "shadowed" to this location to allow the system to start.
A shadow control bit is provided in the system control latch to turn this
mapping on/off.

shadow = 0 (on) -> rom is mapped to 0-0x7ff
shadow = 1 (off) -> ram is mapped to 0-0x7ff

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"

	"github.com/deadsy/go_z80/memory"
)

//-----------------------------------------------------------------------------
// System Memory

const KiB = 1024
const chunkBits = 11 // 2 KiB chunks
const chunkSize = (1 << chunkBits)
const numChunks = (64 * KiB) / chunkSize

func chunkSelect(adr uint16) int { return int(adr >> chunkBits) }

type sysMemory struct {
	memmap     [numChunks]*memory.Memory
	ram0, ram1 *memory.Memory
	rom        *memory.Memory
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
	// Note: There's actually a single 32 KiB device but we break it into
	// two parts so we can write protect the second half independently
	// of the first half.
	ram0 := memory.New(14).RAM() // 16 KiB
	ram1 := memory.New(14).RAM() // 16 KiB

	// Empty
	empty := memory.New(11).Empty() // 2 KiB

	return &sysMemory{
		memmap: [numChunks]*memory.Memory{
			rom,   // 0x0000 - 0x07ff (shadow is ON)
			ram0,  // 0x0800
			ram0,  // 0x1000
			ram0,  // 0x1800
			ram0,  // 0x2000
			ram0,  // 0x2800
			ram0,  // 0x3000
			ram0,  // 0x3800
			ram1,  // 0x4000
			ram1,  // 0x4800
			ram1,  // 0x5000
			ram1,  // 0x5800
			ram1,  // 0x6000
			ram1,  // 0x6800
			ram1,  // 0x7000
			ram1,  // 0x7800
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
		ram0: ram0,
		ram1: ram1,
		rom:  rom,
	}, nil
}

//-----------------------------------------------------------------------------

// shadow controls the presence of rom in the first 2 KiB of the memory map.
func (m *sysMemory) Shadow(on bool) {
	if on {
		// rom is mapped from 0-0x7ff
		m.memmap[0] = m.rom
		return
	}
	// ram is mapped from 0-0x7ff
	m.memmap[0] = m.ram0
}

// write protect the second bank (16 KiB @ 0x4000) of ram.
func (m *sysMemory) WriteProtect(on bool) {
	if on {
		m.ram1.ROM()
		return
	}
	m.ram1.RAM()
}

// reset memory control to initial state
func (m *sysMemory) Reset() {
	m.Shadow(true)
	m.WriteProtect(false)
}

//-----------------------------------------------------------------------------

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
