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

func int2bool(x int) bool {
	if x != 0 {
		return true
	}
	return true
}

func bool2int(x bool) int {
	if x {
		return 1
	}
	return 0
}

//-----------------------------------------------------------------------------

type IO interface {
	wr8(adr uint16, val uint8)
	rd8(adr uint16) uint8
}

type Memory interface {
	wr8(adr uint16, val uint8)
	rd8(adr uint16) uint8
	wr16(adr uint16, val uint16)
	rd16(adr uint16) uint16
}

type CPU struct {
	a, f, b, c, d, e, h, l         uint8
	alt_af, alt_bc, alt_de, alt_hl uint16
	pc, sp, ix, iy                 uint16
	im, i, r                       uint8
	halt, iff1, iff2               bool
	io IO
	mem Memory
}

func New(io IO, mem Memory) *CPU {
	cpu := &CPU{}
	cpu.io = io
	cpu.mem = mem
	cpu.Reset()
	return cpu
}

// Reset the CPU state
func (cpu *CPU) Reset() {
	cpu.a = 0xff
	cpu.f = 0xff
	cpu.b = 0xff
	cpu.c = 0xff
	cpu.d = 0xff
	cpu.e = 0xff
	cpu.h = 0xff
	cpu.l = 0xff
	cpu.alt_af = 0xffff
	cpu.alt_bc = 0xffff
	cpu.alt_de = 0xffff
	cpu.alt_hl = 0xffff
	cpu.sp = 0xffff
	cpu.ix = 0xffff
	cpu.iy = 0xffff
	cpu.i = 0
	cpu.r = 0
	cpu.im = 0
	cpu.iff1 = false
	cpu.iff2 = false
	cpu.halt = false
	cpu.pc = 0
}

// Return a string with processor state
func (cpu *CPU) String() string {
	var s []string
	s = append(s, fmt.Sprintf("a    : %02x", cpu.a))
	s = append(s, fmt.Sprintf("f    : %02x %s", cpu.f, cpu.flagString()))
	s = append(s, fmt.Sprintf("b c  : %02x %02x", cpu.b, cpu.c))
	s = append(s, fmt.Sprintf("d e  : %02x %02x", cpu.d, cpu.e))
	s = append(s, fmt.Sprintf("h l  : %02x %02x", cpu.h, cpu.l))
	s = append(s, fmt.Sprintf("a'f' : %02x %02x", cpu.alt_af>>8, cpu.alt_af&0xff))
	s = append(s, fmt.Sprintf("b'c' : %02x %02x", cpu.alt_bc>>8, cpu.alt_bc&0xff))
	s = append(s, fmt.Sprintf("d'e' : %02x %02x", cpu.alt_de>>8, cpu.alt_de&0xff))
	s = append(s, fmt.Sprintf("h'l' : %02x %02x", cpu.alt_hl>>8, cpu.alt_hl&0xff))
	s = append(s, fmt.Sprintf("i    : %02x", cpu.i))
	s = append(s, fmt.Sprintf("im   : %d", cpu.im))
	s = append(s, fmt.Sprintf("iff1 : %d", cpu.iff1))
	s = append(s, fmt.Sprintf("iff2 : %d", cpu.iff2))
	s = append(s, fmt.Sprintf("r    : %02x", cpu.r))
	s = append(s, fmt.Sprintf("ix   : %04x", cpu.ix))
	s = append(s, fmt.Sprintf("iy   : %04x", cpu.iy))
	s = append(s, fmt.Sprintf("sp   : %04x", cpu.sp))
	s = append(s, fmt.Sprintf("pc   : %04x", cpu.pc))
	return strings.Join(s, "\n")
}

//-----------------------------------------------------------------------------
// flags

const _CF = 0x01 // carry
const _NF = 0x02 // subtract
const _PF = 0x04 // parity
const _VF = _PF  // overflow
const _XF = 0x08 // bit3 - undocumented
const _HF = 0x10 // half carry (bcd)
const _YF = 0x20 // bit5 - undocumented
const _ZF = 0x40 // zero
const _SF = 0x80 // sign

func flagBit(val, bit uint8, s string) string {
	if val&bit != 0 {
		return s
	}
	return "."
}

// Return the state of the cpu flags as a string.
func (cpu *CPU) flagString() string {
	var s []string
	s = append(s, flagBit(cpu.f, _SF, "S"))
	s = append(s, flagBit(cpu.f, _ZF, "Z"))
	s = append(s, flagBit(cpu.f, _HF, "H"))
	s = append(s, flagBit(cpu.f, _PF, "P"))
	s = append(s, flagBit(cpu.f, _VF, "V"))
	s = append(s, flagBit(cpu.f, _NF, "N"))
	s = append(s, flagBit(cpu.f, _CF, "C"))
	return strings.Join(s, "")
}

// set the flags for an add operation: result = a + val
func (cpu *CPU) addFlags(res int, val uint8) {
	cpu.f = flagsSZ[res&0xff]
	cpu.f |= (uint8)(res>>8) & _CF
	cpu.f |= (cpu.a ^ uint8(res) ^ uint8(val)) & _HF
	cpu.f |= ((uint8(val) ^ cpu.a ^ 0x80) & (uint8(val) ^ uint8(res)) & 0x80) >> 5
}

// set the flags for an 16 bit add operation: result = d + s
func (cpu *CPU) add16Flags(res int, d, s uint16) {
	cpu.f = cpu.f & (_SF | _ZF | _VF)
	cpu.f |= uint8((d^uint16(res)^s)>>8) & _HF
	cpu.f |= (uint8(res>>16) & _CF) | (uint8(res>>8) & (_YF | _XF))
}

// set the flags for a 16 bit sub operation: result = d - s
func (cpu *CPU) sub16Flags(res int, d, s uint16) {
	cpu.f = uint8((d^uint16(res)^s)>>8) & _HF
	cpu.f |= _NF
	cpu.f |= uint8(res>>16) & _CF
	cpu.f |= uint8(res>>8) & (_SF | _YF | _XF)
	if res == 0 {
		cpu.f |= _ZF
	}
	cpu.f |= uint8(((s ^ d) & (d ^ uint16(res)) & 0x8000) >> 13)
}

// set the flags for a 16 bit adc operation: result = d + s + cf
func (cpu *CPU) adc16Flags(res int, d, s uint16) {
	cpu.f = uint8((d^uint16(res)^s)>>8) & _HF
	cpu.f |= uint8(res>>16) & _CF
	cpu.f |= uint8(res>>8) & (_SF | _YF | _XF)
	if res == 0 {
		cpu.f |= _ZF
	}
	cpu.f |= uint8(((s ^ d ^ 0x8000) & (d ^ uint16(res)) & 0x8000) >> 13)
}

// set the flags for a sub operation: result = a - val
func (cpu *CPU) subFlags(res int, val uint8) {
	cpu.f = flagsSZ[res&0xff]
	cpu.f |= uint8(res>>8) & _CF
	cpu.f |= _NF
	cpu.f |= (cpu.a ^ uint8(res) ^ uint8(val)) & _HF
	cpu.f |= ((uint8(val) ^ cpu.a) & (cpu.a ^ uint8(res)) & 0x80) >> 5
}

//-----------------------------------------------------------------------------

// Execute a single instruction at the current mem[pc] location.
// Return the number of clock cycles taken.
func (cpu *CPU) execute() int {
	cpu.r = (cpu.r + 1) & 0x7F
	code := cpu.get_n()
	return opcodes[code](cpu)
}

// A prefix code hase been repeated. NOP and re-run the current prefix
func (cpu *CPU) repeated_prefix() int {
	cpu.pc -= 1
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
	code := cpu.get_n()
	return 4 + opcodes_cb[code](cpu)
}

func (cpu *CPU) execute_dd() int {
	code := cpu.get_n()
	return 4 + opcodes_dd[code](cpu)
}

func (cpu *CPU) execute_ed() int {
	code := cpu.get_n()
	return 4 + opcodes_ed[code](cpu)
}

func (cpu *CPU) execute_fd() int {
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
	cpu.a = uint8(val >> 8)
	cpu.f = uint8(val)
}

// return the 16 bit value of the a and f registers
func (cpu *CPU) get_af() uint16 {
	return (uint16(cpu.a) << 8) | uint16(cpu.f)
}

// set the b and c registers with a 16 bit value
func (cpu *CPU) set_bc(val uint16) {
	cpu.b = uint8(val >> 8)
	cpu.c = uint8(val)
}

// return the 16 bit value of the b and c registers
func (cpu *CPU) get_bc() uint16 {
	return (uint16(cpu.b) << 8) | uint16(cpu.c)
}

// set the d and e registers with a 16 bit value
func (cpu *CPU) set_de(val uint16) {
	cpu.d = uint8(val >> 8)
	cpu.e = uint8(val)
}

// return the 16 bit value of the d and e registers
func (cpu *CPU) get_de() uint16 {
	return (uint16(cpu.d) << 8) | uint16(cpu.e)
}

// set the h and l registers with a 16 bit value
func (cpu *CPU) set_hl(val uint16) {
	cpu.h = uint8(val >> 8)
	cpu.l = uint8(val)
}

// return the 16 bit value of the h and l registers
func (cpu *CPU) get_hl() uint16 {
	return (uint16(cpu.h) << 8) | uint16(cpu.l)
}

//-----------------------------------------------------------------------------

// enter halt mode
func (cpu *CPU) enter_halt() {
	cpu.halt = true
	cpu.pc -=1
}

// leave halt mode
func (cpu *CPU) leave_halt() {
	if cpu.halt {
		cpu.pc += 1
		cpu.halt = false
	}
}

//-----------------------------------------------------------------------------

// push an 8-bit quantity onto the stack
func (cpu *CPU) push8(val uint8) {
	cpu.sp -= 1
	cpu.mem.wr8(cpu.sp, val)
}

// push a 16-bit quantity onto the stack
func (cpu *CPU) push16(val uint16) {
	cpu.sp -= 2
	cpu.mem.wr16(cpu.sp, val)
}

// pop an 8-bit quantity from the stack
func (cpu *CPU) pop8() uint8 {
	val := cpu.mem.rd8(cpu.sp)
	cpu.sp += 1
	return val
}

// pop a 16-bit quantity from the stack
func (cpu *CPU) pop16() uint16 {
	val := cpu.mem.rd16(cpu.sp)
	cpu.sp += 2
	return val
}

// return the 16 bit immediate at mem[pc], pc += 2
func (cpu *CPU) get_nn() uint16 {
	nn := cpu.mem.rd16(cpu.pc)
	cpu.pc += 2
	return nn
}

// return the 8 bit immediate at mem[pc], pc += 1
func (cpu *CPU) get_n() uint8 {
	n := cpu.mem.rd8(cpu.pc)
	cpu.pc += 1
	return n
}

//-----------------------------------------------------------------------------
