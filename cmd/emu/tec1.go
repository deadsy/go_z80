//-----------------------------------------------------------------------------
/*

TEC-1 Emulation

*/
//-----------------------------------------------------------------------------

package main

import "fmt"

//-----------------------------------------------------------------------------

const KiB = 1024

const romStart = uint16(0)
const ramStart = uint16(0x800)
const ramSize = 2 * KiB
const romSize = 2 * KiB

func addressIsWithin(adr, start, size uint16) bool {
	return (adr >= start) && (adr < (start + size))
}

type Memory struct {
	rom          [romSize]byte
	ram          [ramSize]byte
	romWriteable bool
}

// Rd8 reads a byte from memory.
func (m *Memory) Read8(adr uint16) uint8 {
	if addressIsWithin(adr, romStart, romSize) {
		return m.rom[adr-romStart]
	}
	if addressIsWithin(adr, ramStart, ramSize) {
		return m.ram[adr-ramStart]
	}
	fmt.Printf("mem.Rd8 address %04x is out of range\n", adr)
	return 0xff
}

func (m *Memory) Write8(adr uint16, val uint8) {
	if addressIsWithin(adr, romStart, romSize) {
		if m.romWriteable {
			m.rom[adr-romStart] = val
		} else {
			fmt.Printf("mem.Wr8 address %04x is ROM - can't write\n", adr)
		}
		return
	}
	if addressIsWithin(adr, ramStart, ramSize) {
		m.ram[adr-ramStart] = val
		return
	}
	fmt.Printf("mem.Wr8 address %04x is out of range\n", adr)
}

func (m *Memory) WriteROM(flag bool) {
	m.romWriteable = flag
}

func (m *Memory) Read16(adr uint16) uint16 {
	l := uint16(m.Read8(adr))
	h := uint16(m.Read8(adr + 1))
	return (h << 8) | l
}

func (m *Memory) Write16(adr uint16, val uint16) {
	m.Write8(adr, uint8(val))
	m.Write8(adr+1, uint8(val>>8))
}

func newMemory() *Memory {
	m := Memory{}
	// all 0xffs
	for i := range m.rom {
		m.rom[i] = 0xff
	}
	for i := range m.ram {
		m.ram[i] = 0xff
	}
	m.romWriteable = true
	return &m
}

//-----------------------------------------------------------------------------

type IO struct {
}

// Rd8 reads a byte from an IO port.
func (io *IO) Read8(adr uint16) uint8 {
	adr &= 0xff
	fmt.Printf("io.Read8 [%02x]\n", adr)
	return 0
}

// Wr8 writes a byte to an IO port.
func (io *IO) Write8(adr uint16, val uint8) {
	adr &= 0xff
	fmt.Printf("io.Write8 [%02x] = %02x\n", adr, val)
}

func newIO() *IO {
	return &IO{}
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
