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
	data  []byte // memory data
	mask  uint16 // address mask
	read  bool   // can the memory be read?
	write bool   // can the memory be written?
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
	adr &= m.mask
	if !m.write || int(adr) >= len(m.data) {
		return
	}
	m.data[adr] = val
}

func (m *Memory) Read8(adr uint16) uint8 {
	adr &= m.mask
	if !m.read || int(adr) >= len(m.data) {
		return 0xff
	}
	return m.data[adr]
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
