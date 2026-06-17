//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/deadsy/go_z80/z80"
)

//-----------------------------------------------------------------------------

type IO struct {
}

func NewIO() z80.IO {
	return &IO{}
}

func (io *IO) Wr8(addr uint16, val uint8) {
}

func (io *IO) Rd8(addr uint16) uint8 {
	return 0
}

//-----------------------------------------------------------------------------

const KiB = 1024

type Memory struct {
	mem [64 * KiB]uint8
}

func NewMemory() z80.Memory {
	return &Memory{}
}

func (m *Memory) Wr8(addr uint16, val uint8) {
	//log.Printf("mem wr8 [%04x] = %02x\n", addr, val)
	m.mem[addr] = val
}

func (m *Memory) Rd8(addr uint16) uint8 {
	return m.mem[addr]
}

func (m *Memory) Wr16(addr uint16, val uint16) {
	//log.Printf("mem wr16 [%04x] = %04x\n", addr, val)
	m.mem[addr] = uint8(val)
	m.mem[addr+1] = uint8(val >> 8)
}

func (m *Memory) Rd16(addr uint16) uint16 {
	return uint16(m.mem[addr]) + (uint16(m.mem[addr+1]) << 8)
}

func (m *Memory) Set(ram [][2]int) {
	for _, x := range ram {
		addr := uint16(x[0])
		val := uint8(x[1])
		m.Wr8(addr, val)
	}
}

func (m *Memory) Check(ram [][2]int) error {
	for _, x := range ram {
		addr := uint16(x[0])
		expected := uint8(x[1])
		actual := m.Rd8(addr)
		if actual != expected {
			return fmt.Errorf("[%04x] is %02x, expected %02x", addr, actual, expected)
		}
	}
	return nil
}

//-----------------------------------------------------------------------------

// State matches the register values before and after an instruction executes
type State struct {
	A byte `json:"a"`
	F byte `json:"f"`
	B byte `json:"b"`
	C byte `json:"c"`
	D byte `json:"d"`
	E byte `json:"e"`
	H byte `json:"h"`
	L byte `json:"l"`

	Alt_AF uint16 `json:"af_"`
	Alt_BC uint16 `json:"bc_"`
	Alt_DE uint16 `json:"de_"`
	Alt_HL uint16 `json:"hl_"`

	PC uint16 `json:"pc"`
	SP uint16 `json:"sp"`
	IX uint16 `json:"ix"`
	IY uint16 `json:"iy"`

	IM   byte `json:"im"`
	I    byte `json:"i"`
	R    byte `json:"r"`
	IFF1 byte `json:"iff1"`
	IFF2 byte `json:"iff2"`

	EI byte   `json:"ei"`
	P  byte   `json:"p"`
	Q  byte   `json:"q"`
	WZ uint16 `json:"wz"`

	Ram [][2]int `json:"ram"`
}

// Z80Test represents a single standalone test vector
type Z80Test struct {
	Name    string `json:"name"`
	Initial State  `json:"initial"`
	Final   State  `json:"final"`
}

func setState(cpu *z80.CPU, s *State) {

	cpu.A = s.A
	cpu.F = s.F
	cpu.B = s.B
	cpu.C = s.C
	cpu.D = s.D
	cpu.E = s.E
	cpu.H = s.H
	cpu.L = s.L

	cpu.Alt_AF = s.Alt_AF
	cpu.Alt_BC = s.Alt_BC
	cpu.Alt_DE = s.Alt_DE
	cpu.Alt_HL = s.Alt_HL

	cpu.PC = s.PC
	cpu.SP = s.SP
	cpu.IX = s.IX
	cpu.IY = s.IY

	cpu.IM = s.IM
	cpu.I = s.I
	cpu.R = s.R
	cpu.IFF1 = s.IFF1
	cpu.IFF2 = s.IFF2
}

func cmpState(cpu *z80.CPU, s *State) error {

	if cpu.A != s.A {
		return fmt.Errorf("A, expected 0x%02x(%d), actual 0x%02x(%d)", s.A, s.A, cpu.A, cpu.A)
	}
	if cpu.F != s.F {
		return fmt.Errorf("F, expected 0x%02x(%d), actual 0x%02x(%d)", s.F, s.F, cpu.F, cpu.F)
	}
	if cpu.B != s.B {
		return fmt.Errorf("B")
	}
	if cpu.C != s.C {
		return fmt.Errorf("C")
	}
	if cpu.D != s.D {
		return fmt.Errorf("D")
	}
	if cpu.E != s.E {
		return fmt.Errorf("E")
	}
	if cpu.H != s.H {
		return fmt.Errorf("H")
	}
	if cpu.L != s.L {
		return fmt.Errorf("L")
	}

	if cpu.Alt_AF != s.Alt_AF {
		return fmt.Errorf("Alt_AF")
	}
	if cpu.Alt_BC != s.Alt_BC {
		return fmt.Errorf("Alt_BC")
	}
	if cpu.Alt_DE != s.Alt_DE {
		return fmt.Errorf("Alt_DE")
	}
	if cpu.Alt_HL != s.Alt_HL {
		return fmt.Errorf("Alt_HL")
	}

	if cpu.PC != s.PC {
		return fmt.Errorf("PC")
	}
	if cpu.SP != s.SP {
		return fmt.Errorf("SP")
	}
	if cpu.IX != s.IX {
		return fmt.Errorf("IX")
	}
	if cpu.IY != s.IY {
		return fmt.Errorf("IY")
	}

	if cpu.IM != s.IM {
		return fmt.Errorf("IM")
	}
	if cpu.I != s.I {
		return fmt.Errorf("I")
	}
	if cpu.R != s.R {
		return fmt.Errorf("R")
	}
	if cpu.IFF1 != s.IFF1 {
		return fmt.Errorf("IFF1")
	}
	if cpu.IFF2 != s.IFF2 {
		return fmt.Errorf("IFF2")
	}

	return nil
}

//-----------------------------------------------------------------------------

func runTest(fname string) error {

	data, err := os.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("failed to load test file: %v", err)
	}

	var tests []Z80Test
	if err := json.Unmarshal(data, &tests); err != nil {
		return fmt.Errorf("failed to parse test JSON: %v", err)
	}

	for i, t := range tests {
		fmt.Printf("%s, test %d: %s\n", fname, i, t.Name)

		io := NewIO()
		mem := NewMemory()
		cpu := z80.New(io, mem)

		// setup ram state
		mem.(*Memory).Set(t.Initial.Ram)

		// setup cpu state
		setState(cpu, &t.Initial)

		cpu.Execute()

		// check cpu state
		err = cmpState(cpu, &t.Final)
		if err != nil {
			return err
		}

		// check ram state
		err = mem.(*Memory).Check(t.Final.Ram)
		if err != nil {
			return err
		}

	}

	return nil
}

//-----------------------------------------------------------------------------

func main() {

	pattern := filepath.Join("/home/jasonh/personal/z80_tests/z80/v1", "*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {

		err = runTest(match)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	}

}

//-----------------------------------------------------------------------------
