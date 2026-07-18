//-----------------------------------------------------------------------------
/*

TEC-1 Memory Emulation

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
	memmap [numChunks]*memory.Memory
}

func newMemory() (*sysMemory, error) {
	// ROM
	rom := memory.New(11).ROM() // 2 KiB
	data, err := assets.ReadFile("assets/mon1B.bin")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded ROM: %w", err)
	}
	if err := rom.Load(0, data); err != nil {
		return nil, fmt.Errorf("failed to load ROM: %w", err)
	}
	// RAM
	ram := memory.New(11).RAM() // 2 KiB

	// Empty
	empty := memory.New(11).Empty() // 2 KiB

	return &sysMemory{
		memmap: [numChunks]*memory.Memory{
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

//-----------------------------------------------------------------------------

func (m *sysMemory) Read8(adr uint16) uint8 {
	return m.memmap[chunkSelect(adr)].Read8(adr)
}

func (m *sysMemory) Write8(adr uint16, val uint8) {
	m.memmap[chunkSelect(adr)].Write8(adr, val)
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
