//-----------------------------------------------------------------------------
/*

Test z80 opcodes

Single steps each opcode and check final machine, memory and IO state.
Uses JSON test vectors from https://github.com/SingleStepTests/z80

*/
//-----------------------------------------------------------------------------

package z80

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

//-----------------------------------------------------------------------------

const KiB = 1024

//-----------------------------------------------------------------------------

// ignore the undocumented XY flags
const ignore_XY_Flags = true

// ignore state errors for these tests
var ignoreStateErrors = []string{
	"ED A2",
	"ED A3",
	"ED AA",
	"ED AB",
	"ED B2",
	"ED B3",
	"ED BA",
	"ED BB",
}

func ignoreState(name string) bool {
	for _, v := range ignoreStateErrors {
		if strings.HasPrefix(name, v) {
			return true
		}
	}
	return false
}

//-----------------------------------------------------------------------------
// Opcode Test Memory

type otMemory struct {
	mem [64 * KiB]uint8
}

func newOpcodeTestMemory() Memory {
	return &otMemory{}
}

func (m *otMemory) Write8(addr uint16, val uint8) {
	//log.Printf("mem wr8 [%04x] = %02x", addr, val)
	m.mem[addr] = val
}

func (m *otMemory) Read8(addr uint16) uint8 {
	return m.mem[addr]
}

func (m *otMemory) Write16(addr uint16, val uint16) {
	//log.Printf("mem wr16 [%04x] = %04x", addr, val)
	m.mem[addr] = uint8(val)
	m.mem[addr+1] = uint8(val >> 8)
}

func (m *otMemory) Read16(addr uint16) uint16 {
	return uint16(m.mem[addr]) + (uint16(m.mem[addr+1]) << 8)
}

func (m *otMemory) Set(ram [][2]int) {
	for _, x := range ram {
		addr := uint16(x[0])
		val := uint8(x[1])
		m.Write8(addr, val)
	}
}

func (m *otMemory) Check(ram [][2]int) error {
	for _, x := range ram {
		addr := uint16(x[0])
		expected := uint8(x[1])
		actual := m.Read8(addr)
		if actual != expected {
			return fmt.Errorf("[%04x] is %02x, expected %02x", addr, actual, expected)
		}
	}
	return nil
}

//-----------------------------------------------------------------------------
// Opcode Test IO

const ioRead = uint8(1 << 0)
const ioWrite = uint8(1 << 1)

type otIO struct {
	port [64 * KiB]uint8
	op   [64 * KiB]uint8
}

func newOpcodeTestIO() IO {
	return &otIO{}
}

func (io *otIO) Write8(addr uint16, val uint8) {
	//log.Printf("io wr8 [%04x] = %02x", addr, val)
	if io.op[addr]&ioWrite != 0 {
		io.port[addr] = val
	} else {
		log.Printf("io wr8 [%04x] no write allowed", addr)
	}
}

func (io *otIO) Read8(addr uint16) uint8 {
	//log.Printf("io rd8 [%04x]", addr)
	if io.op[addr]&ioRead != 0 {
		return io.port[addr]
	}
	log.Printf("io rd8 [%04x] no read allowed", addr)
	return 0
}

func (io *otIO) Set(ports [][3]any) {
	for _, x := range ports {
		addr := uint16(x[0].(float64))
		val := uint8(x[1].(float64))
		op := x[2].(string)
		io.port[addr] = val
		switch op {
		case "r":
			io.op[addr] |= ioRead
		case "w":
			io.op[addr] |= ioWrite
		default:
			panic(fmt.Sprintf("unknown operation %s", op))
		}
	}
}

//-----------------------------------------------------------------------------
// Opcode Test Bus

type otBus struct {
}

func newOpcodeTestBus() Bus {
	return &otBus{}
}

func (bus *otBus) ReadIV() uint8 {
	return 0xff
}

//-----------------------------------------------------------------------------
// Opcode Test State

// State matches the register values before and after an instruction executes
type otState struct {
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

// opcodeTest represents a single standalone test vector
type opcodeTest struct {
	Name    string   `json:"name"`
	Initial otState  `json:"initial"`
	Final   otState  `json:"final"`
	Cycles  [][3]any `json:"cycles"`
	Ports   [][3]any `json:"ports"`
}

func (s *otState) set(cpu *CPU) {

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
	cpu.IFF1 = byte2bool(s.IFF1)
	cpu.IFF2 = byte2bool(s.IFF2)
}

func (s *otState) compare(cpu *CPU) error {

	if ignore_XY_Flags {
		// clear the undocumented flag bits
		cpu.F = cpu.F &^ (_XF | _YF)
		s.F = s.F &^ (_XF | _YF)
	}

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
		return fmt.Errorf("PC, expected 0x%04x(%d), actual 0x%04x(%d)", s.PC, s.PC, cpu.PC, cpu.PC)
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
		return fmt.Errorf("R, expected 0x%02x(%d), actual 0x%02x(%d)", s.R, s.R, cpu.R, cpu.R)
	}
	if cpu.IFF1 != byte2bool(s.IFF1) {
		return fmt.Errorf("IFF1, expected %t, actual %t", byte2bool(s.IFF1), cpu.IFF1)
	}
	if cpu.IFF2 != byte2bool(s.IFF2) {
		return fmt.Errorf("IFF2, expected %t, actual %t", byte2bool(s.IFF2), cpu.IFF2)
	}

	return nil
}

func runTest(t *testing.T, fname string) error {

	data, err := os.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("failed to load test file: %v", err)
	}

	var tests []opcodeTest
	if err := json.Unmarshal(data, &tests); err != nil {
		return fmt.Errorf("failed to parse test JSON: %v", err)
	}

	for _, v := range tests {
		t.Logf("%s, %s\n", fname, v.Name)

		io := newOpcodeTestIO()
		mem := newOpcodeTestMemory()
		bus := newOpcodeTestBus()
		cpu := New(io, mem, bus)

		// setup ram state
		mem.(*otMemory).Set(v.Initial.Ram)

		// setup io state
		io.(*otIO).Set(v.Ports)

		// setup cpu state
		v.Initial.set(cpu)

		cpu.Run()

		if cpu.cycles != len(v.Cycles) {
			return fmt.Errorf("cycles, expected %d, actual %d", len(v.Cycles), cpu.cycles)
		}

		// check cpu state
		err = v.Final.compare(cpu)
		if err != nil {
			if ignoreState(v.Name) {
				t.Logf("%s (ignored)\n", err.Error())
			} else {
				return err
			}
		}

		// check ram state
		err = mem.(*otMemory).Check(v.Final.Ram)
		if err != nil {
			return err
		}

	}

	return nil
}

//-----------------------------------------------------------------------------

func Test_Opcodes(t *testing.T) {

	pattern := filepath.Join("../ext/z80step/v1", "*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		t.Fatal(err)
	}
	for _, match := range matches {
		err = runTest(t, match)
		if err != nil {
			t.Fatalf("error: %s", err)
		}
	}
}

//-----------------------------------------------------------------------------
