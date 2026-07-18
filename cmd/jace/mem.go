//-----------------------------------------------------------------------------
/*

Jupiter ACE Emulator

Memory Map

0000-1FFF 8k ACE O/S and Forth Interpreter/Compiler ROM
2000-22FF (Unused)
2300-23FF 256 bytes used for Cassette Header Information
2400-26FF 32 x 24 Display. Bit 7 indicates inverse (black text on white), Bits 6-0 indicate a character in the Character Definitions Table
2700-27FF Edit Buffer end ? Startup writes '0' here but little other obvious function.
2800-2BFF (Unused)
2C00-2FFF 128 x 8 byte character definitions, representing display characters 0..127.
3000-3BFF (Unused)
3C00-FFFF Program

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
const charChunk = 5 // character memory

func chunkSelect(adr uint16) int { return int(adr >> chunkBits) }

type sysMemory struct {
	memmap    [numChunks]*memory.Memory
	charDirty bool // has the character memory been written?
}

func newMemory() (*sysMemory, error) {
	// ROM
	rom := memory.New(13).ROM() // 8KiB
	data, err := assets.ReadFile("assets/jace.bin")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded ROM: %w", err)
	}
	if err := rom.Load(0, data); err != nil {
		return nil, fmt.Errorf("failed to load ROM: %w", err)
	}
	// Video RAM
	video := memory.New(10).RAM() // 1 KiB
	// Character RAM
	char := memory.New(10).RAM() // 1 KiB
	// RAM
	ram := memory.New(10).RAM() // 1 KiB
	// Empty
	empty := memory.New(11).Empty() // 2 KiB

	return &sysMemory{
		memmap: [numChunks]*memory.Memory{
			rom,   // 0x0000 - 0x07ff
			rom,   // 0x0800 - 0x0fff
			rom,   // 0x1000 - 0x17ff
			rom,   // 0x1800 - 0x1fff
			video, // 0x2000 - 0x27ff - 1K repeats 2 times
			char,  // 0x2800 - 0x2fff - 1K repeats 2 times
			ram,   // 0x3000 - 0x37ff - 1K repeats 4 times
			ram,   // 0x3800 - 0x3fff
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

func (m *sysMemory) IsDirty() bool {
	return m.charDirty
}

func (m *sysMemory) Clean() {
	m.charDirty = false
}

//-----------------------------------------------------------------------------

func (m *sysMemory) Read8(adr uint16) uint8 {
	return m.memmap[chunkSelect(adr)].Read8(adr)
}

func (m *sysMemory) Write8(adr uint16, val uint8) {
	n := chunkSelect(adr)
	// If character memory changes the video device needs to update the font atlas.
	if n == charChunk {
		m.charDirty = true
	}
	m.memmap[n].Write8(adr, val)
}

func (m *sysMemory) Read16(adr uint16) uint16 {
	l := uint16(m.memmap[chunkSelect(adr)].Read8(adr))
	h := uint16(m.memmap[chunkSelect(adr+1)].Read8(adr + 1))
	return (h << 8) | l
}

func (m *sysMemory) Write16(adr uint16, val uint16) {
	m.memmap[chunkSelect(adr)].Write8(adr, uint8(val))
	m.memmap[chunkSelect(adr+1)].Write8(adr+1, uint8(val>>8))
}

//-----------------------------------------------------------------------------
