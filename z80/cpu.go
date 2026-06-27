//-----------------------------------------------------------------------------

package z80

import (
	"fmt"
	"strings"
)

//-----------------------------------------------------------------------------

func offset16(ofs uint8) uint16 {
	return uint16(int8(ofs))
}

func byte2bool(x byte) bool {
	return x != 0
}

func bool2byte(x bool) byte {
	if x {
		return 1
	}
	return 0
}

//-----------------------------------------------------------------------------

type IO interface {
	Write8(adr uint16, val uint8)
	Read8(adr uint16) uint8
}

type Memory interface {
	Write8(adr uint16, val uint8)
	Read8(adr uint16) uint8
	Write16(adr uint16, val uint16)
	Read16(adr uint16) uint16
}

type Bus interface {
	ReadIV() uint8
}

// nmiAddress is the non-maskable interrupt address
const nmiAddress = uint16(0x0066)

// im1Address is the IM1 interrupt address
const im1Address = uint16(0x0038)

// rstAddress is the reset address
const rstAddress = uint16(0)

type CPU struct {
	A, F, B, C, D, E, H, L         uint8
	Alt_AF, Alt_BC, Alt_DE, Alt_HL uint16
	PC, SP, IX, IY                 uint16
	IM, I, R                       uint8
	IFF1, IFF2                     bool

	halt bool // halt state
	nmi  bool // nmi state
	irq  bool // irq state

	io  IO     // io of target system
	mem Memory // memory of target system
	bus Bus    // bus of target system

	//cycles int // total number of cpu cycles run
}

func New(io IO, mem Memory, bus Bus) *CPU {
	cpu := &CPU{}
	cpu.io = io
	cpu.mem = mem
	cpu.bus = bus
	cpu.Reset()
	return cpu
}

// Reset the CPU state
func (cpu *CPU) Reset() {
	cpu.A = 0xff
	cpu.F = 0xff
	cpu.B = 0xff
	cpu.C = 0xff
	cpu.D = 0xff
	cpu.E = 0xff
	cpu.H = 0xff
	cpu.L = 0xff

	cpu.Alt_AF = 0xffff
	cpu.Alt_BC = 0xffff
	cpu.Alt_DE = 0xffff
	cpu.Alt_HL = 0xffff

	cpu.PC = rstAddress
	cpu.SP = 0xffff
	cpu.IX = 0xffff
	cpu.IY = 0xffff

	cpu.IM = 0
	cpu.I = 0
	cpu.R = 0
	cpu.IFF1 = false
	cpu.IFF2 = false

	cpu.halt = false
	cpu.nmi = false
	cpu.irq = false
}

// Return a string with processor state
func (cpu *CPU) String() string {
	var s []string
	s = append(s, fmt.Sprintf("a    : %02x", cpu.A))
	s = append(s, fmt.Sprintf("f    : %02x %s", cpu.F, cpu.flagString()))
	s = append(s, fmt.Sprintf("b c  : %02x %02x", cpu.B, cpu.C))
	s = append(s, fmt.Sprintf("d e  : %02x %02x", cpu.D, cpu.E))
	s = append(s, fmt.Sprintf("h l  : %02x %02x", cpu.H, cpu.L))
	s = append(s, fmt.Sprintf("a'f' : %02x %02x", cpu.Alt_AF>>8, cpu.Alt_AF&0xff))
	s = append(s, fmt.Sprintf("b'c' : %02x %02x", cpu.Alt_BC>>8, cpu.Alt_BC&0xff))
	s = append(s, fmt.Sprintf("d'e' : %02x %02x", cpu.Alt_DE>>8, cpu.Alt_DE&0xff))
	s = append(s, fmt.Sprintf("h'l' : %02x %02x", cpu.Alt_HL>>8, cpu.Alt_HL&0xff))
	s = append(s, fmt.Sprintf("i    : %02x", cpu.I))
	s = append(s, fmt.Sprintf("im   : %d", cpu.IM))
	s = append(s, fmt.Sprintf("iff1 : %t", cpu.IFF1))
	s = append(s, fmt.Sprintf("iff2 : %t", cpu.IFF2))
	s = append(s, fmt.Sprintf("r    : %02x", cpu.R))
	s = append(s, fmt.Sprintf("ix   : %04x", cpu.IX))
	s = append(s, fmt.Sprintf("iy   : %04x", cpu.IY))
	s = append(s, fmt.Sprintf("sp   : %04x", cpu.SP))
	s = append(s, fmt.Sprintf("pc   : %04x", cpu.PC))
	return strings.Join(s, "\n")
}

//-----------------------------------------------------------------------------
// flags

const _CF = uint8(0x01) // carry
const _NF = uint8(0x02) // subtract
const _PF = uint8(0x04) // parity
const _VF = _PF         // overflow
const _XF = uint8(0x08) // bit3 - undocumented
const _HF = uint8(0x10) // half carry (bcd)
const _YF = uint8(0x20) // bit5 - undocumented
const _ZF = uint8(0x40) // zero
const _SF = uint8(0x80) // sign

func flagBit(val, bit uint8, s string) string {
	if val&bit != 0 {
		return s
	}
	return "."
}

// Return the state of the cpu flags as a string.
func (cpu *CPU) flagString() string {
	var s []string
	s = append(s, flagBit(cpu.F, _SF, "S"))
	s = append(s, flagBit(cpu.F, _ZF, "Z"))
	s = append(s, flagBit(cpu.F, _HF, "H"))
	s = append(s, flagBit(cpu.F, _PF, "P"))
	s = append(s, flagBit(cpu.F, _VF, "V"))
	s = append(s, flagBit(cpu.F, _NF, "N"))
	s = append(s, flagBit(cpu.F, _CF, "C"))
	return strings.Join(s, "")
}

// set the flags for an add operation: result = a + val
func (cpu *CPU) addFlags(res int, val uint8) {
	cpu.F = flagsSZ[res&0xff]
	cpu.F |= (uint8)(res>>8) & _CF
	cpu.F |= (cpu.A ^ uint8(res) ^ uint8(val)) & _HF
	cpu.F |= ((uint8(val) ^ cpu.A ^ 0x80) & (uint8(val) ^ uint8(res)) & 0x80) >> 5
}

// set the flags for an 16 bit add operation: result = d + s
func (cpu *CPU) add16Flags(res int, d, s uint16) {
	cpu.F = cpu.F & (_SF | _ZF | _VF)
	cpu.F |= uint8((d^uint16(res)^s)>>8) & _HF
	cpu.F |= (uint8(res>>16) & _CF) | (uint8(res>>8) & (_YF | _XF))
}

// set the flags for a 16 bit sub operation: result = d - s
func (cpu *CPU) sub16Flags(res int, d, s uint16) {
	cpu.F = uint8((d^uint16(res)^s)>>8) & _HF
	cpu.F |= _NF
	cpu.F |= uint8(res>>16) & _CF
	cpu.F |= uint8(res>>8) & (_SF | _YF | _XF)
	if res == 0 {
		cpu.F |= _ZF
	}
	cpu.F |= uint8(((s ^ d) & (d ^ uint16(res)) & 0x8000) >> 13)
}

// set the flags for a 16 bit adc operation: result = d + s + cf
func (cpu *CPU) adc16Flags(res int, d, s uint16) {
	cpu.F = uint8((d^uint16(res)^s)>>8) & _HF
	cpu.F |= uint8(res>>16) & _CF
	cpu.F |= uint8(res>>8) & (_SF | _YF | _XF)
	if res == 0 {
		cpu.F |= _ZF
	}
	cpu.F |= uint8(((s ^ d ^ 0x8000) & (d ^ uint16(res)) & 0x8000) >> 13)
}

// set the flags for a sub operation: result = a - val
func (cpu *CPU) subFlags(res int, val uint8) {
	cpu.F = flagsSZ[res&0xff]
	cpu.F |= uint8(res>>8) & _CF
	cpu.F |= _NF
	cpu.F |= (cpu.A ^ uint8(res) ^ uint8(val)) & _HF
	cpu.F |= ((uint8(val) ^ cpu.A) & (cpu.A ^ uint8(res)) & 0x80) >> 5
}

//-----------------------------------------------------------------------------

func (cpu *CPU) inc_r() {
	cpu.R = (cpu.R + 1) & 0x7F
}

// A prefix code hase been repeated. NOP and re-run the current prefix
func (cpu *CPU) repeated_prefix() int {
	cpu.inc_r()
	cpu.PC -= 1
	return 0
}

func (cpu *CPU) execute_dddd() int {
	return cpu.repeated_prefix()
}

func (cpu *CPU) execute_ddfd() int {
	return cpu.repeated_prefix()
}

func (cpu *CPU) execute_fddd() int {
	return cpu.repeated_prefix()
}

func (cpu *CPU) execute_fdfd() int {
	return cpu.repeated_prefix()
}

func (cpu *CPU) execute_cb() int {
	cpu.inc_r()
	code := cpu.get_n()
	return 4 + opcodes_cb[code](cpu)
}

func (cpu *CPU) execute_dd() int {
	cpu.inc_r()
	code := cpu.get_n()
	return 4 + opcodes_dd[code](cpu)
}

func (cpu *CPU) execute_ed() int {
	cpu.inc_r()
	code := cpu.get_n()
	return 4 + opcodes_ed[code](cpu)
}

func (cpu *CPU) execute_fd() int {
	cpu.inc_r()
	code := cpu.get_n()
	return 4 + opcodes_fd[code](cpu)
}

func (cpu *CPU) execute_ddcb() int {
	d := cpu.get_n()
	code := cpu.get_n()
	return 8 + opcodes_ddcb00[code](cpu, d)
}

func (cpu *CPU) execute_fdcb() int {
	d := cpu.get_n()
	code := cpu.get_n()
	return 8 + opcodes_fdcb00[code](cpu, d)
}

//-----------------------------------------------------------------------------

// set the a and f registers with a 16 bit value
func (cpu *CPU) set_af(val uint16) {
	cpu.A = uint8(val >> 8)
	cpu.F = uint8(val)
}

// return the 16 bit value of the a and f registers
func (cpu *CPU) get_af() uint16 {
	return (uint16(cpu.A) << 8) | uint16(cpu.F)
}

// set the b and c registers with a 16 bit value
func (cpu *CPU) set_bc(val uint16) {
	cpu.B = uint8(val >> 8)
	cpu.C = uint8(val)
}

// return the 16 bit value of the b and c registers
func (cpu *CPU) get_bc() uint16 {
	return (uint16(cpu.B) << 8) | uint16(cpu.C)
}

// set the d and e registers with a 16 bit value
func (cpu *CPU) set_de(val uint16) {
	cpu.D = uint8(val >> 8)
	cpu.E = uint8(val)
}

// return the 16 bit value of the d and e registers
func (cpu *CPU) get_de() uint16 {
	return (uint16(cpu.D) << 8) | uint16(cpu.E)
}

// set the h and l registers with a 16 bit value
func (cpu *CPU) set_hl(val uint16) {
	cpu.H = uint8(val >> 8)
	cpu.L = uint8(val)
}

// return the 16 bit value of the h and l registers
func (cpu *CPU) get_hl() uint16 {
	return (uint16(cpu.H) << 8) | uint16(cpu.L)
}

//-----------------------------------------------------------------------------

// push an 8-bit quantity onto the stack
func (cpu *CPU) push8(val uint8) {
	cpu.SP -= 1
	cpu.mem.Write8(cpu.SP, val)
}

// push a 16-bit quantity onto the stack
func (cpu *CPU) push16(val uint16) {
	cpu.SP -= 2
	cpu.mem.Write16(cpu.SP, val)
}

// pop an 8-bit quantity from the stack
func (cpu *CPU) pop8() uint8 {
	val := cpu.mem.Read8(cpu.SP)
	cpu.SP += 1
	return val
}

// pop a 16-bit quantity from the stack
func (cpu *CPU) pop16() uint16 {
	val := cpu.mem.Read16(cpu.SP)
	cpu.SP += 2
	return val
}

// return the 16 bit immediate at mem[pc], pc += 2
func (cpu *CPU) get_nn() uint16 {
	nn := cpu.mem.Read16(cpu.PC)
	cpu.PC += 2
	return nn
}

// return the 8 bit immediate at mem[pc], pc += 1
func (cpu *CPU) get_n() uint8 {
	n := cpu.mem.Read8(cpu.PC)
	cpu.PC += 1
	return n
}

//-----------------------------------------------------------------------------

// NMI starts the non-maskable interrupt processing.
func (cpu *CPU) NMI() {
	cpu.nmi = true
}

// IRQ starts the maskable interrupt processing.
func (cpu *CPU) IRQ() {
	cpu.irq = true
}

// IM0 interrupt mode handling
func (cpu *CPU) handleIM0() (int, error) {
	// read an opcode from the bus (single byte only)
	code := cpu.bus.ReadIV()
	// execute the opcode
	cycles := opcodes[code](cpu)
	return cycles + 2, nil
}

// IM1 interrupt mode handling
func (cpu *CPU) handleIM1() (int, error) {
	cpu.push16(cpu.PC)
	// Jump to fixed vector
	cpu.PC = im1Address
	return 13, nil
}

// IM2 interrupt mode handling
func (cpu *CPU) handleIM2() (int, error) {
	cpu.push16(cpu.PC)
	// Get the low byte from the hardware bus (16-bit aligned)
	vec := cpu.bus.ReadIV() & 0xfe
	// work out the table address
	adr := (uint16(cpu.I) << 8) | uint16(vec)
	// jump to the pc stored in the lookup table
	cpu.PC = cpu.mem.Read16(adr)
	return 19, nil
}

func (cpu *CPU) irqHandling() (int, error) {
	// taken out of halt
	cpu.halt = false
	// disable interrupts
	cpu.IFF1 = false
	cpu.IFF2 = false
	// Service the specific interrupt mode
	switch cpu.IM {
	case 0:
		return cpu.handleIM0()
	case 1:
		return cpu.handleIM1()
	case 2:
		return cpu.handleIM2()
	}
	return 0, fmt.Errorf("invalid IM mode %d", cpu.IM)
}

func (cpu *CPU) nmiHandling() (int, error) {
	// taken out of halt
	cpu.halt = false
	// Backup the maskable interrupt state
	cpu.IFF2 = cpu.IFF1
	// Disable future maskable interrupts during NMI execution
	cpu.IFF1 = false
	// Push current Program Counter to the stack
	cpu.push16(cpu.PC)
	// Jump to the hardcoded NMI vector address
	cpu.PC = nmiAddress
	// Account for internal CPU clock cycles (NMI takes 11 T-states)
	return 11, nil
}

// Run the Z80 CPU for a single instruction.
func (cpu *CPU) Run() (int, error) {

	// increment r
	cpu.inc_r()

	if cpu.nmi {
		cpu.nmi = false
		return cpu.nmiHandling()
	}

	if cpu.irq && cpu.IFF1 {
		cpu.irq = false
		return cpu.irqHandling()
	}

	if cpu.halt {
		// nop
		return 4, nil
	}

	// execute the next opcode
	code := cpu.get_n()
	return opcodes[code](cpu), nil
}

//-----------------------------------------------------------------------------
