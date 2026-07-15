//-----------------------------------------------------------------------------
/*

TEC-1G Emulation

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"

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

type Bus struct {
}

func newBus() *Bus {
	return &Bus{}
}

func (bus *Bus) ReadIV() uint8 {
	return 0xff
}

//-----------------------------------------------------------------------------
