//-----------------------------------------------------------------------------
/*

Memory Emulation

16-bit address, 8-bit data

*/
//-----------------------------------------------------------------------------

package memory

import (
	"fmt"
	"io/ioutil"
)

//-----------------------------------------------------------------------------

type Memory struct {
	data       []byte // memory data
	mask       uint16 // address mask
	read       bool   // can the memory be read?
	write      bool   // can the memory be written?
	totalRead  int    // total attempted reads
	totalWrite int    // total attempted writes
	badRead    int    // bad read count
	badWrite   int    // bad write count
}

// New returns memory with storage allocated
func New(bits int) *Memory {
	size := 1 << bits
	return &Memory{
		data: make([]byte, size),
		mask: uint16(size) - 1,
	}
}

func (m *Memory) Write8(adr uint16, val uint8) {
	m.totalWrite += 1
	adr &= m.mask
	if !m.write || int(adr) >= len(m.data) {
		m.badWrite += 1
		return
	}
	m.data[adr] = val
}

func (m *Memory) Read8(adr uint16) uint8 {
	m.totalRead += 1
	adr &= m.mask
	if !m.read || int(adr) >= len(m.data) {
		m.badRead += 1
		return 0xff
	}
	return m.data[adr]
}

func (m *Memory) Write16(adr uint16, val uint16) {
	m.Write8(adr, uint8(val))
	m.Write8(adr+1, uint8(val>>8))
}

func (m *Memory) Read16(adr uint16) uint16 {
	l := uint16(m.Read8(adr))
	h := uint16(m.Read8(adr + 1))
	return (h << 8) | l
}

// RAM has read/write permissions
func (m *Memory) RAM() *Memory {
	m.read = true
	m.write = true
	return m
}

// ROM is read only
func (m *Memory) ROM() *Memory {
	m.read = true
	m.write = false
	return m
}

// WOM is write only
func (m *Memory) WOM() *Memory {
	m.read = false
	m.write = true
	return m
}

// Empty memory has no storage
func (m *Memory) Empty() *Memory {
	m.read = false
	m.write = false
	m.data = nil
	return m
}

//-----------------------------------------------------------------------------

// Load memory with provided data
func (m *Memory) Load(adr uint16, data []byte) error {
	if int(adr)+len(data) > len(m.data) {
		return fmt.Errorf("data is too long to load into memory")
	}
	for i, v := range data {
		m.data[int(adr)+i] = v
	}
	return nil
}

func (m *Memory) LoadFile(adr uint16, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return m.Load(adr, data)
}

//-----------------------------------------------------------------------------
