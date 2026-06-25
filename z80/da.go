//-----------------------------------------------------------------------------
/*

Z80 Disassembler

*/
//-----------------------------------------------------------------------------

package z80

import (
	"fmt"
	"strings"
)

//-----------------------------------------------------------------------------

var _r = []string{"b", "c", "d", "e", "h", "l", "(hl)", "a"}
var _rp = []string{"bc", "de", "hl", "sp"}
var _rp2 = []string{"bc", "de", "hl", "af"}
var _cc = []string{"nz", "z", "nc", "c", "po", "pe", "p", "m"}
var _alu = []string{"add", "adc", "sub", "sbc", "and", "xor", "or", "cp"}
var _alux = []string{"a,", "a,", "", "a,", "", "", "", ""}
var _rot = []string{"rlc", "rrc", "rl", "rr", "sla", "sra", "sll", "srl"}
var _rota = []string{"rlca", "rrca", "rla", "rra", "daa", "cpl", "scf", "ccf"}
var _im = []string{"0", "0", "1", "2", "0", "0", "1", "2"}
var _bli = [][]string{
	{"ldi", "ldd", "ldir", "lddr"},
	{"cpi", "cpd", "cpir", "cpdr"},
	{"ini", "ind", "inir", "indr"},
	{"outi", "outd", "otir", "otdr"},
}

//-----------------------------------------------------------------------------

func getSign(d uint16) string {
	if int16(d) < 0 {
		// fmt.Sprintf("%02x") will add the "-"
		return ""
	}
	return "+"
}

//-----------------------------------------------------------------------------

// SymbolTable maps an address to a symbol.
type SymbolTable map[uint16]string

// Disassembly returns the result of the disassembler call.
type Disassembly struct {
	Dump        string  // address and memory bytes
	Symbol      string  // symbol for the address (if any)
	Instruction string  // instruction decode
	Comment     string  // useful comment
	Bytes       []uint8 // decoded bytes
}

func (da *Disassembly) String() string {
	s := make([]string, 2)
	s[0] = fmt.Sprintf("%-18s %8s %-13s", da.Dump, da.Symbol, da.Instruction)
	if da.Comment != "" {
		s[1] = fmt.Sprintf(" ; %s", da.Comment)
	}
	return strings.Join(s, "")
}

//-----------------------------------------------------------------------------

type Decode struct {
	Mnemonic string
	Operands string
	Length   int
}

func da_Index(mem Memory, pc uint16, ir string) Decode {
	m0 := mem.Read8(pc)
	m1 := mem.Read8(pc + 1)
	m2 := mem.Read8(pc + 2)

	x := (m0 >> 6) & 3
	y := (m0 >> 3) & 7
	z := m0 & 7
	p := (m0 >> 4) & 3
	q := (m0 >> 3) & 1

	n0 := m1
	n1 := m2
	nn := uint16(m2)<<8 | uint16(m1)

	d := offset16(m1)
	sign := getSign(d)

	dj := pc + d + 2

	// if using (hl) then: (hl)->(ix+d), h and l unchanged
	alt0r := make([]string, len(_r))
	copy(alt0r, _r)
	alt0r[6] = fmt.Sprintf("(%s%s%02x)", ir, sign, int8(d))

	// if not using (hl) then: hl->ix, h->ixh, l->ixl
	alt1r := make([]string, len(_r))
	copy(alt1r, _r)
	alt1r[4] = ir + "h"
	alt1r[5] = ir + "l"

	altRp := make([]string, len(_rp))
	copy(altRp, _rp)
	altRp[2] = ir

	altRp2 := make([]string, len(_rp2))
	copy(altRp2, _rp2)
	altRp2[2] = ir

	if x == 0 {
		if z == 0 {
			if y == 0 {
				return Decode{"nop", "", 2}
			} else if y == 1 {
				return Decode{"ex", "af,af'", 2}
			} else if y == 2 {
				return Decode{"djnz", fmt.Sprintf("%04x", dj), 3}
			} else if y == 3 {
				return Decode{"jr", fmt.Sprintf("%04x", dj), 3}
			}
			return Decode{"jr", fmt.Sprintf("%s,%04x", _cc[y-4], dj), 3}
		} else if z == 1 {
			if q == 0 {
				return Decode{"ld", fmt.Sprintf("%s,%04x", altRp[p], nn), 4}
			}
			return Decode{"add", fmt.Sprintf("%s,%s", ir, altRp[p]), 2}
		} else if z == 2 {
			if q == 0 {
				if p == 0 {
					return Decode{"ld", "(bc),a", 2}
				} else if p == 1 {
					return Decode{"ld", "(de),a", 2}
				} else if p == 2 {
					return Decode{"ld", fmt.Sprintf("(%04x),%s", nn, ir), 4}
				}
				return Decode{"ld", fmt.Sprintf("(%04x),a", nn), 4}
			}
			if p == 0 {
				return Decode{"ld", "a,(bc)", 2}
			} else if p == 1 {
				return Decode{"ld", "a,(de)", 2}
			} else if p == 2 {
				return Decode{"ld", fmt.Sprintf("%s,(%04x)", ir, nn), 4}
			}
			return Decode{"ld", fmt.Sprintf("a,(%04x)", nn), 4}
		} else if z == 3 {
			if q == 0 {
				return Decode{"inc", altRp[p], 2}
			}
			return Decode{"dec", altRp[p], 2}
		} else if z == 4 {
			if y == 6 {
				return Decode{"inc", alt0r[y], 3}
			}
			return Decode{"inc", alt1r[y], 2}
		} else if z == 5 {
			if y == 6 {
				return Decode{"dec", alt0r[y], 3}
			}
			return Decode{"dec", alt1r[y], 2}
		} else if z == 6 {
			if y == 6 {
				return Decode{"ld", fmt.Sprintf("%s,%02x", alt0r[y], n1), 4}
			}
			return Decode{"ld", fmt.Sprintf("%s,%02x", alt1r[y], n0), 3}
		}
		return Decode{_rota[y], "", 2}
	} else if x == 1 {
		if z == 6 && y == 6 {
			return Decode{"halt", "", 2}
		}
		if y == 6 || z == 6 {
			return Decode{"ld", fmt.Sprintf("%s,%s", alt0r[y], alt0r[z]), 3}
		}
		return Decode{"ld", fmt.Sprintf("%s,%s", alt1r[y], alt1r[z]), 2}
	} else if x == 2 {
		if z == 6 {
			return Decode{_alu[y], fmt.Sprintf("%s%s", _alux[y], alt0r[z]), 3}
		}
		return Decode{_alu[y], fmt.Sprintf("%s%s", _alux[y], alt1r[z]), 2}
	} else {
		if z == 0 {
			return Decode{"ret", _cc[y], 2}
		} else if z == 1 {
			if q == 0 {
				return Decode{"pop", altRp2[p], 2}
			}
			if p == 0 {
				return Decode{"ret", "", 2}
			} else if p == 1 {
				return Decode{"exx", "", 2}
			} else if p == 2 {
				return Decode{"jp", ir, 2}
			}
			return Decode{"ld", fmt.Sprintf("sp,%s", ir), 2}
		} else if z == 2 {
			return Decode{"jp", fmt.Sprintf("%s,%04x", _cc[y], nn), 4}
		} else if z == 3 {
			if y == 0 {
				return Decode{"jp", fmt.Sprintf("%04x", nn), 4}
			} else if y == 2 {
				return Decode{"out", fmt.Sprintf("(%02x),a", n0), 3}
			} else if y == 3 {
				return Decode{"in", fmt.Sprintf("a,(%02x)", n0), 3}
			} else if y == 4 {
				return Decode{"ex", fmt.Sprintf("(sp),%s", ir), 2}
			} else if y == 5 {
				return Decode{"ex", "de,hl", 2}
			} else if y == 6 {
				return Decode{"di", "", 2}
			}
			return Decode{"ei", "", 2}
		} else if z == 4 {
			return Decode{"call", fmt.Sprintf("%s,%04x", _cc[y], nn), 4}
		} else if z == 5 {
			if q == 0 {
				return Decode{"push", altRp2[p], 2}
			}
			if p == 0 {
				return Decode{"call", fmt.Sprintf("%04x", nn), 4}
			}
		} else if z == 6 {
			return Decode{_alu[y], fmt.Sprintf("%s%02x", _alux[y], n0), 3}
		} else {
			return Decode{"rst", fmt.Sprintf("%02x", y<<3), 2}
		}
	}
	panic("unreachable")
}

func da_ddcb_fdcb_Prefix(mem Memory, pc uint16, ir string) Decode {

	m0 := mem.Read8(pc)
	m1 := mem.Read8(pc + 1)

	x := (m1 >> 6) & 3
	y := (m1 >> 3) & 7
	z := m1 & 7

	d := offset16(m0)
	sign := getSign(d)

	disp := fmt.Sprintf("(%s%s%02x)", ir, sign, int8(d))

	if x == 0 {
		if z == 6 {
			return Decode{_rot[y], disp, 4}
		}
		return Decode{_rot[y], fmt.Sprintf("%s,%s", disp, _r[z]), 4}
	} else if x == 1 {
		return Decode{"bit", fmt.Sprintf("%d,%s", y, disp), 4}
	} else if x == 2 {
		if z == 6 {
			return Decode{"res", fmt.Sprintf("%d,%s", y, disp), 4}
		}
		return Decode{"res", fmt.Sprintf("%d,%s,%s", y, disp, _r[z]), 4}
	} else {
		if z == 6 {
			return Decode{"set", fmt.Sprintf("%d,%s", y, disp), 4}
		}
		return Decode{"set", fmt.Sprintf("%d,%s,%s", y, disp, _r[z]), 4}
	}
}

// 0xCB <opcode>
func da_cb_Prefix(mem Memory, pc uint16) Decode {
	m0 := mem.Read8(pc)

	x := (m0 >> 6) & 3
	y := (m0 >> 3) & 7
	z := m0 & 7

	if x == 0 {
		return Decode{_rot[y], _r[z], 2}
	} else if x == 1 {
		return Decode{"bit", fmt.Sprintf("%d,%s", y, _r[z]), 2}
	} else if x == 2 {
		return Decode{"res", fmt.Sprintf("%d,%s", y, _r[z]), 2}
	}
	return Decode{"set", fmt.Sprintf("%d,%s", y, _r[z]), 2}
}

// 0xDD <x>
// 0xFD <x>
func da_dd_fd_Prefix(mem Memory, pc uint16, ir string) Decode {
	m0 := mem.Read8(pc)

	if m0 == 0xdd || m0 == 0xed || m0 == 0xfd {
		return Decode{"nop", "", 1}
	} else if m0 == 0xcb {
		return da_ddcb_fdcb_Prefix(mem, pc+1, ir)
	} else {
		return da_Index(mem, pc, ir)
	}
}

// 0xED <opcode>
// 0xED <opcode> <nn>
func da_ed_Prefix(mem Memory, pc uint16) Decode {

	m0 := mem.Read8(pc)
	m1 := mem.Read8(pc + 1)
	m2 := mem.Read8(pc + 2)

	x := (m0 >> 6) & 3
	y := (m0 >> 3) & 7
	z := m0 & 7
	p := (m0 >> 4) & 3
	q := (m0 >> 3) & 1

	nn := uint16(m2)<<8 | uint16(m1)

	if x == 1 {
		if z == 0 {
			if y == 6 {
				return Decode{"in", "(c)", 2}
			}
			return Decode{"in", fmt.Sprintf("%s,(c)", _r[y]), 2}
		} else if z == 1 {
			if y == 6 {
				return Decode{"out", "(c)", 2}
			}
			return Decode{"out", fmt.Sprintf("(c),%s", _r[y]), 2}
		} else if z == 2 {
			if q == 0 {
				return Decode{"sbc", fmt.Sprintf("hl,%s", _rp[p]), 2}
			}
			return Decode{"adc", fmt.Sprintf("hl,%s", _rp[p]), 2}
		} else if z == 3 {
			if q == 0 {
				return Decode{"ld", fmt.Sprintf("(%04x),%s", nn, _rp[p]), 4}
			}
			return Decode{"ld", fmt.Sprintf("%s,(%04x)", _rp[p], nn), 4}
		} else if z == 4 {
			return Decode{"neg", "", 2}
		} else if z == 5 {
			if y == 1 {
				return Decode{"reti", "", 2}
			}
			return Decode{"retn", "", 2}
		} else if z == 6 {
			return Decode{"im", _im[y], 2}
		} else {
			if y == 0 {
				return Decode{"ld", "i,a", 2}
			} else if y == 1 {
				return Decode{"ld", "r,a", 2}
			} else if y == 2 {
				return Decode{"ld", "a,i", 2}
			} else if y == 3 {
				return Decode{"ld", "a,r", 2}
			} else if y == 4 {
				return Decode{"rrd", "", 2}
			} else if y == 5 {
				return Decode{"rld", "", 2}
			}
			return Decode{"nop", "", 2}
		}
	} else if x == 2 {
		if z <= 3 && y >= 4 {
			return Decode{_bli[z][y-4], "", 2}
		}
	}
	return Decode{"nop", "", 2}
}

// Normal decode with no prefixes
func da_Normal(mem Memory, pc uint16) Decode {

	m0 := mem.Read8(pc)
	m1 := mem.Read8(pc + 1)
	m2 := mem.Read8(pc + 2)

	x := (m0 >> 6) & 3
	y := (m0 >> 3) & 7
	z := m0 & 7
	p := (m0 >> 4) & 3
	q := (m0 >> 3) & 1

	n := m1
	nn := uint16(m2)<<8 | uint16(m1)
	d := pc + offset16(m1) + 2

	if x == 0 {
		if z == 0 {
			if y == 0 {
				return Decode{"nop", "", 1}
			} else if y == 1 {
				return Decode{"ex", "af,af'", 1}
			} else if y == 2 {
				return Decode{"djnz", fmt.Sprintf("%04x", d), 2}
			} else if y == 3 {
				return Decode{"jr", fmt.Sprintf("%04x", d), 2}
			} else {
				return Decode{"jr", fmt.Sprintf("%s,%04x", _cc[y-4], d), 2}
			}
		} else if z == 1 {
			if q == 0 {
				return Decode{"ld", fmt.Sprintf("%s,%04x", _rp[p], nn), 3}
			}
			return Decode{"add", fmt.Sprintf("hl,%s", _rp[p]), 1}
		} else if z == 2 {
			if q == 0 {
				if p == 0 {
					return Decode{"ld", "(bc),a", 1}
				} else if p == 1 {
					return Decode{"ld", "(de),a", 1}
				} else if p == 2 {
					return Decode{"ld", fmt.Sprintf("(%04x),hl", nn), 3}
				}
				return Decode{"ld", fmt.Sprintf("(%04x),a", nn), 3}
			}

			if p == 0 {
				return Decode{"ld", "a,(bc)", 1}
			} else if p == 1 {
				return Decode{"ld", "a,(de)", 1}
			} else if p == 2 {
				return Decode{"ld", fmt.Sprintf("hl,(%04x)", nn), 3}
			}
			return Decode{"ld", fmt.Sprintf("a,(%04x)", nn), 3}
		} else if z == 3 {
			if q == 0 {
				return Decode{"inc", _rp[p], 1}
			}
			return Decode{"dec", _rp[p], 1}
		} else if z == 4 {
			return Decode{"inc", _r[y], 1}
		} else if z == 5 {
			return Decode{"dec", _r[y], 1}
		} else if z == 6 {
			return Decode{"ld", fmt.Sprintf("%s,%02x", _r[y], n), 2}
		}
		return Decode{_rota[y], "", 1}
	} else if x == 1 {
		if z == 6 && y == 6 {
			return Decode{"halt", "", 1}
		}
		return Decode{"ld", fmt.Sprintf("%s,%s", _r[y], _r[z]), 1}
	} else if x == 2 {
		return Decode{_alu[y], fmt.Sprintf("%s%s", _alux[y], _r[z]), 1}
	} else {
		if z == 0 {
			return Decode{"ret", _cc[y], 1}
		} else if z == 1 {
			if q == 0 {
				return Decode{"pop", _rp2[p], 1}
			}
			if p == 0 {
				return Decode{"ret", "", 1}
			} else if p == 1 {
				return Decode{"exx", "", 1}
			} else if p == 2 {
				return Decode{"jp", "hl", 1}
			}
			return Decode{"ld", "sp,hl", 1}
		} else if z == 2 {
			return Decode{"jp", fmt.Sprintf("%s,%04x", _cc[y], nn), 3}
		} else if z == 3 {
			if y == 0 {
				return Decode{"jp", fmt.Sprintf("%04x", nn), 3}
			} else if y == 2 {
				return Decode{"out", fmt.Sprintf("(%02x),a", n), 2}
			} else if y == 3 {
				return Decode{"in", fmt.Sprintf("a,(%02x)", n), 2}
			} else if y == 4 {
				return Decode{"ex", "(sp),hl", 1}
			} else if y == 5 {
				return Decode{"ex", "de,hl", 1}
			} else if y == 6 {
				return Decode{"di", "", 1}
			}
			return Decode{"ei", "", 1}
		} else if z == 4 {
			return Decode{"call", fmt.Sprintf("%s,%04x", _cc[y], nn), 3}
		} else if z == 5 {
			if q == 0 {
				return Decode{"push", _rp2[p], 1}
			}
			if p == 0 {
				return Decode{"call", fmt.Sprintf("%04x", nn), 3}
			}
		} else if z == 6 {
			return Decode{_alu[y], fmt.Sprintf("%s%02x", _alux[y], n), 2}
		} else {
			return Decode{"rst", fmt.Sprintf("%02x", y<<3), 1}
		}
	}
	panic("")
}

// Disassemble z80 opcodes starting at mem[pc].
// Return a (operation, operands, nbytes) tuple.
func daInstruction(mem Memory, pc uint16) Decode {
	m0 := mem.Read8(pc)
	if m0 == 0xCB {
		return da_cb_Prefix(mem, pc+1)
	} else if m0 == 0xDD {
		return da_dd_fd_Prefix(mem, pc+1, "ix")
	} else if m0 == 0xED {
		return da_ed_Prefix(mem, pc+1)
	} else if m0 == 0xFD {
		return da_dd_fd_Prefix(mem, pc+1, "iy")
	}
	return da_Normal(mem, pc)
}

//-----------------------------------------------------------------------------

func daDump(adr uint16, mem []byte) string {
	s := make([]string, len(mem))
	for i, v := range mem {
		s[i] = fmt.Sprintf("%02x", v)
	}
	return fmt.Sprintf("%04x: %s", adr, strings.Join(s, " "))
}

func daSymbol(adr uint16, st SymbolTable) string {
	if st != nil {
		return st[adr]
	}
	return ""
}

// Disassemble a Z80 instruction from the memory at the address.
func Disassemble(mem Memory, pc uint16, st SymbolTable) *Disassembly {
	decode := daInstruction(mem, pc)
	bytes := make([]uint8, decode.Length)
	for i := 0; i < decode.Length; i++ {
		bytes[i] = mem.Read8(pc + uint16(i))
	}
	return &Disassembly{
		Dump:        daDump(pc, bytes),
		Symbol:      daSymbol(pc, st),
		Instruction: fmt.Sprintf("%s %s", decode.Mnemonic, decode.Operands),
		Comment:     "",
		Bytes:       bytes,
	}
}

//-----------------------------------------------------------------------------

// Disassemble returns the disassembly for a region of the CPU memory.
func (cpu *CPU) Disassemble(adr uint16, size int) string {
	s := make([]string, 0, 16)
	for size > 0 {
		da := Disassemble(cpu.mem, adr, nil)
		s = append(s, da.String())
		n := len(da.Bytes)
		size -= n
		adr += uint16(n)
	}
	return strings.Join(s, "\n")
}

//-----------------------------------------------------------------------------
