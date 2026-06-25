# -----------------------------------------------------------------------------
"""
Z80 Opcode Emulation Generator
"""
# -----------------------------------------------------------------------------

import sys
import getopt
import z80da
import memory

# -----------------------------------------------------------------------------
# format is (opcode prefix, prototype), (links to other prefixes), 'function preamble'

_preamble_0 = "() int"
_prototype_0 = "func(*CPU) int"

_preamble_1 = "(d uint8) int"
_prototype_1 = "func(*CPU, uint8) int"

_prefixes = (
    ((), (0xCB, 0xDD, 0xED, 0xFD), _preamble_0, _prototype_0),
    ((0xCB,), (), _preamble_0, _prototype_0),
    ((0xDD,), (0xCB, 0xDD, 0xFD), _preamble_0, _prototype_0),
    ((0xDD, 0xCB, 0x00), (), _preamble_1, _prototype_1),
    ((0xED,), (), _preamble_0, _prototype_0),
    ((0xFD,), (0xCB, 0xDD, 0xFD), _preamble_0, _prototype_0),
    ((0xFD, 0xCB, 0x00), (), _preamble_1, _prototype_1),
)

# -----------------------------------------------------------------------------

_indent = 4


class output:
    """class for handling file output with auto indenting"""

    def __init__(self, ofname):
        self.ofile = open("%s" % ofname, "w")
        self.lhs = 0
        self.col = 0

    def close(self):
        self.ofile.close()

    def put(self, s):
        for c in s:
            if c == "\n":
                self.ofile.write("\n")
                self.col = 0
            else:
                if self.col == 0:
                    self.ofile.write(" " * self.lhs)
                    self.col = self.lhs
                self.ofile.write(c)
                self.col += 1

    def pad(self, col):
        if self.col < col:
            self.ofile.write(" " * (col - self.col))
            self.col = col

    def indent(self, n):
        self.lhs += n * _indent

    def outdent(self, n):
        self.lhs -= n * _indent


# -----------------------------------------------------------------------------

_r = ("b", "c", "d", "e", "h", "l", "(hl)", "a")
_rp = ("bc", "de", "hl", "sp")
_rp2 = ("bc", "de", "hl", "af")
_direct_rp = ("sp", "ix", "iy")
_cc = ("nz", "z", "nc", "c", "po", "pe", "p", "m")
_alu = ("add", "adc", "sub", "sbc", "and", "xor", "or", "cp")
_alux = ("a", "a", "", "a", "", "", "", "")
_rot = ("rlc", "rrc", "rl", "rr", "sla", "sra", "sll", "srl")
_rota = ("rlca", "rrca", "rla", "rra", "daa", "cpl", "scf", "ccf")
_im = ("0", "0", "1", "2", "0", "0", "1", "2")
_bli = (
    ("ldi", "ldd", "ldir", "lddr"),
    ("cpi", "cpd", "cpir", "cpdr"),
    ("ini", "ind", "inir", "indr"),
    ("outi", "outd", "otir", "otdr"),
)

# -----------------------------------------------------------------------------
# 8-Bit Load Group


def emit_ld_r_n(out, r):
    """load immediate register n"""
    if r == "(hl)":
        out.put("cpu.mem.Write8(cpu.get_hl(), cpu.get_n())\n")
        out.put("return 10\n")
    else:
        out.put("cpu.%s = cpu.get_n()\n" % r.upper())
        out.put("return 7\n")


def emit_ld_hilo_immediate(out, r):
    hi = r[2] == "h"
    r = r[:-1].upper()
    if hi:
        out.put(f"cpu.{r} = (uint16(cpu.get_n()) << 8) | (cpu.{r} & 0xff)\n")
    else:
        out.put(f"cpu.{r} = (cpu.{r} & 0xff00) | uint16(cpu.get_n())\n")
    out.put("return 7\n")


def emit_ld_r_hilo(out, dst, src):
    if dst == src:
        out.put("return 4\n")
        return

    if dst in _r:
        dst = dst.upper()
        hi = src[2] == "h"
        src = src[:-1].upper()
        select = ("& 0xff", ">> 8")[hi]
        out.put(f"cpu.{dst} = uint8(cpu.{src} {select})\n")
        out.put("return 4\n")
        return

    if src in _r:
        src = src.upper()
        hi = dst[2] == "h"
        dst = dst[:-1].upper()
        if hi:
            out.put(f"cpu.{dst} = (uint16(cpu.{src}) << 8) | (cpu.{dst} & 0xff)\n")
        else:
            out.put(f"cpu.{dst} = (cpu.{dst} & 0xff00) | uint16(cpu.{src})\n")
        out.put("return 4\n")
        return

    lo2hi = (dst[2] == "h") and (src[2] == "l")
    dst = dst[:-1].upper()
    if lo2hi:
        out.put(f"cpu.{dst} = ((cpu.{dst} & 0xff) << 8) | (cpu.{dst} & 0xff)\n")
    else:
        out.put(f"cpu.{dst} = (cpu.{dst} & 0xff00) | (cpu.{dst} >> 8)\n")
    out.put("return 4\n")


def emit_ld_mem_xx_n(out, r):
    """ld (xx),n where xx is ix+d, iy+d"""
    out.put("d := offset16(cpu.get_n())\n")
    out.put("cpu.mem.Write8(cpu.%s + d, cpu.get_n())\n" % r.upper())
    out.put("return 15\n")


def emit_ld_r_r(out, rd, rs):
    """load register to register"""
    if rd == "(hl)":
        out.put("cpu.mem.Write8(cpu.get_hl(), cpu.%s)\n" % rs.upper())
        out.put("return 7\n")
    elif rd == "(ix+d)":
        out.put("d := offset16(cpu.get_n())\n")
        out.put("cpu.mem.Write8(cpu.IX + d, cpu.%s)\n" % rs.upper())
        out.put("return 15\n")
    elif rd == "(iy+d)":
        out.put("d := offset16(cpu.get_n())\n")
        out.put("cpu.mem.Write8(cpu.IY + d, cpu.%s)\n" % rs.upper())
        out.put("return 15\n")
    elif rs == "(hl)":
        out.put("cpu.%s = cpu.mem.Read8(cpu.get_hl())\n" % rd.upper())
        out.put("return 7\n")
    elif rs == "(ix+d)":
        out.put("d := offset16(cpu.get_n())\n")
        out.put("cpu.%s = cpu.mem.Read8(cpu.IX + d)\n" % rd.upper())
        out.put("return 15\n")
    elif rs == "(iy+d)":
        out.put("d := offset16(cpu.get_n())\n")
        out.put("cpu.%s = cpu.mem.Read8(cpu.IY + d)\n" % rd.upper())
        out.put("return 15\n")
    else:
        out.put("cpu.%s = cpu.%s\n" % (rd.upper(), rs.upper()))
        out.put("return 4\n")


def emit_ld_a_mem_xx(out, xx):
    """ld a,(xx) where xx in (bc,de,nn)"""
    out.put("cpu.A = cpu.mem.Read8(cpu.get_%s())\n" % xx)
    out.put("return %d\n" % (7, 13)[xx == "nn"])


def emit_ld_mem_xx_a(out, xx):
    """ld (xx),a - where xx in (bc,de,nn)"""
    out.put("cpu.mem.Write8(cpu.get_%s(), cpu.A)\n" % xx)
    out.put("return %d\n" % (7, 13)[xx == "nn"])


def emit_ld_ira(out, d, s):
    """ld i/r/a, i/r/a"""
    out.put("cpu.%s = cpu.%s\n" % (d.upper(), s.upper()))
    if d == "a":
        out.put("cpu.F =  (cpu.F & _CF) | (flagsSZ[cpu.A]) | (bool2byte(cpu.IFF2) << 2)\n")
    out.put("return 5\n")


# -----------------------------------------------------------------------------
# 16-Bit Load Group


def emit_ld_rp_nn(out, rp):
    """ld rp,nn"""
    if rp in _direct_rp:
        out.put("cpu.%s = cpu.get_nn()\n" % rp.upper())
    else:
        out.put("cpu.set_%s(cpu.get_nn())\n" % rp)
    out.put("return 10\n")


def emit_ld_mem_nn_rp(out, rp):
    """ld (nn), rp"""
    out.put("nn := cpu.get_nn()\n")
    if rp in _direct_rp:
        out.put("cpu.mem.Write8(nn, uint8(cpu.%s))\n" % rp.upper())
        out.put("cpu.mem.Write8(nn + 1, uint8(cpu.%s >> 8))\n" % rp.upper())
    else:
        out.put("cpu.mem.Write8(nn, cpu.%s)\n" % rp[1].upper())
        out.put("cpu.mem.Write8(nn + 1, cpu.%s)\n" % rp[0].upper())
    out.put("return 16\n")


def emit_ld_rp_mem_nn(out, rp):
    """ld rp,(nn)"""
    out.put("nn := cpu.get_nn()\n")
    if rp in _direct_rp:
        out.put("cpu.%s = uint16(cpu.mem.Read8(nn + 1)) << 8\n" % rp.upper())
        out.put("cpu.%s |= uint16(cpu.mem.Read8(nn))\n" % rp.upper())
    else:
        out.put("cpu.%s = cpu.mem.Read8(nn + 1)\n" % rp[0].upper())
        out.put("cpu.%s = cpu.mem.Read8(nn)\n" % rp[1].upper())
    out.put("return 16\n")


def emit_ld_sp_hl(out):
    """ld sp, hl"""
    out.put("cpu.SP = cpu.get_hl()\n")
    out.put("return 6\n")


def emit_ld_index(out, r):
    """ld sp, ix/iy"""
    out.put(f"cpu.SP = cpu.{r.upper()}\n")
    out.put("return 6\n")


def emit_pop_rp(out, rp):
    """pop rp"""
    if rp in _direct_rp:
        out.put("cpu.%s = cpu.pop16()\n" % rp.upper())
    else:
        out.put("cpu.%s = cpu.mem.Read8(cpu.SP + 1)\n" % rp[0].upper())
        out.put("cpu.%s = cpu.mem.Read8(cpu.SP)\n" % rp[1].upper())
        out.put("cpu.SP += 2\n")
    out.put("return 10\n")


def emit_push_rp(out, rp):
    """pop rp"""
    if rp in _direct_rp:
        out.put("cpu.push16(cpu.%s)\n" % rp.upper())
    else:
        out.put("cpu.mem.Write8(cpu.SP - 1, cpu.%s)\n" % rp[0].upper())
        out.put("cpu.mem.Write8(cpu.SP - 2, cpu.%s)\n" % rp[1].upper())
        out.put("cpu.SP -= 2\n")
    out.put("return 11\n")


# -----------------------------------------------------------------------------
# Exchange, Block Transfer, and Search Group


def emit_ex_de_hl(out):
    """ex de,hl"""
    out.put("cpu.D, cpu.H = cpu.H, cpu.D\n")
    out.put("cpu.E, cpu.L = cpu.L, cpu.E\n")
    out.put("return 4\n")


def emit_ex_af_af(out):
    """ex af,af'"""
    out.put("tmp := cpu.get_af()\n")
    out.put("cpu.set_af(cpu.Alt_AF)\n")
    out.put("cpu.Alt_AF = tmp\n")
    out.put("return 4\n")


def emit_ldxx(out, op):
    """ldi, ldir, ldd, lddr"""
    dirn = ("-", "+")[op in ("ldi", "ldir")]
    out.put("d := cpu.get_de()\n")
    out.put("s := cpu.get_hl()\n")
    out.put("n := cpu.get_bc() - 1\n")
    out.put("val := cpu.mem.Read8(s)\n")
    out.put("cpu.mem.Write8(d, val)\n")
    out.put("cpu.F &= (_SF | _ZF | _CF)\n")

    out.put("if ((cpu.A + val) & 0x02) != 0 {cpu.F |= _YF}\n")
    out.put("if ((cpu.A + val) & 0x08 ) != 0 {cpu.F |= _XF}\n")

    out.put("cpu.set_de(d %s 1)\n" % dirn)
    out.put("cpu.set_hl(s %s 1)\n" % dirn)
    out.put("cpu.set_bc(n)\n")

    out.put("if n != 0 {\n")
    out.put("cpu.F |= _VF\n")
    if op in ("ldir", "lddr"):
        out.put("cpu.PC -= 2\n")
        out.put("return 17\n")
    out.put("}\n")

    out.put("return 12\n")


def emit_cp_mem(out, inc, rep):
    delta = ("-1", "+1")[inc]
    out.put("src := cpu.get_hl()\n")
    out.put("val := cpu.mem.Read8(src)\n")
    out.put("res := int(cpu.A) - int(val)\n")
    out.put(f"cpu.set_hl(src {delta})\n")
    out.put("n := cpu.get_bc() - 1\n")
    out.put("cpu.set_bc(n)\n")
    out.put("cpu.F =  (cpu.F & _CF) | _NF\n")
    out.put("if cpu.A == val {cpu.F |= _ZF}\n")
    out.put("if n != 0 {cpu.F |= _PF}\n")
    out.put("if res & 0x80 != 0 {cpu.F |= _SF}\n")
    out.put("cpu.F |= (cpu.A ^ val ^ uint8(res)) & _HF\n")
    if rep:
        out.put("if (cpu.A != val) && (n != 0) {cpu.PC -= 2; return 17}\n")
    out.put("return 12\n")


def emit_in_io(out, inc, rep):
    delta = ("-1", "+ 1")[inc]
    out.put("dst := cpu.get_hl()\n")
    out.put("val := cpu.io.Read8(cpu.get_bc())\n")
    out.put("cpu.mem.Write8(dst, val)\n")
    out.put(f"cpu.set_hl(dst {delta})\n")
    out.put("cpu.B -= 1\n")
    out.put("cpu.F = (cpu.F & _CF) | _NF\n")
    out.put("if cpu.B == 0 {cpu.F |= _ZF}\n")
    if rep:
        out.put("if cpu.B != 0 {cpu.PC -= 2; return 17}\n")
    out.put("return 12\n")


def emit_out_io(out, inc, rep):
    delta = ("-1", "+ 1")[inc]
    out.put("src := cpu.get_hl()\n")
    out.put("val := cpu.mem.Read8(src)\n")
    out.put("cpu.B -= 1\n")
    out.put("cpu.io.Write8(cpu.get_bc(), val)\n")
    out.put(f"cpu.set_hl(src {delta})\n")
    out.put("cpu.F = (cpu.F & _CF) | _NF\n")
    out.put("if cpu.B == 0 {cpu.F |= _ZF}\n")
    if rep:
        out.put("if cpu.B != 0 {cpu.PC -= 2; return 17}\n")
    out.put("return 12\n")


def emit_bli(out, op):
    """block instructions"""
    if op in ("ldi", "ldir", "ldd", "lddr"):
        return emit_ldxx(out, op)

    if op == "cpi":
        emit_cp_mem(out, True, False)
    elif op == "cpir":
        emit_cp_mem(out, True, True)
    elif op == "cpd":
        emit_cp_mem(out, False, False)
    elif op == "cpdr":
        emit_cp_mem(out, False, True)
    elif op == "ini":
        emit_in_io(out, True, False)
    elif op == "inir":
        emit_in_io(out, True, True)
    elif op == "ind":
        emit_in_io(out, False, False)
    elif op == "indr":
        emit_in_io(out, False, True)
    elif op == "outi":
        emit_out_io(out, True, False)
    elif op == "otir":
        emit_out_io(out, True, True)
    elif op == "outd":
        emit_out_io(out, False, False)
    elif op == "otdr":
        emit_out_io(out, False, True)
    else:
        assert False


def emit_ex_mem_sp_r(out, r):
    """ex (sp),r"""
    out.put("tmp := cpu.mem.Read16(cpu.SP)\n")
    if r == "hl":
        out.put("cpu.mem.Write16(cpu.SP, cpu.get_hl())\n")
        out.put("cpu.set_hl(tmp)\n")
    else:
        out.put("cpu.mem.Write16(cpu.SP, cpu.%s)\n" % r.upper())
        out.put("cpu.%s = tmp\n" % r.upper())
    out.put("return 19\n")


def emit_exx(out):
    """exx"""
    out.put("tmp := cpu.get_bc()\n")
    out.put("cpu.set_bc(cpu.Alt_BC)\n")
    out.put("cpu.Alt_BC = tmp\n")
    out.put("tmp = cpu.get_de()\n")
    out.put("cpu.set_de(cpu.Alt_DE)\n")
    out.put("cpu.Alt_DE = tmp\n")
    out.put("tmp = cpu.get_hl()\n")
    out.put("cpu.set_hl(cpu.Alt_HL)\n")
    out.put("cpu.Alt_HL = tmp\n")
    out.put("return 4\n")


# -----------------------------------------------------------------------------
# 8-Bit Arithmetic Group


def emit_inc_dec_r(out, r, op):
    """inc/dec register"""
    delta = ("+ 1", "- 1")[op == "dec"]
    flags = ("flagsSZHVinc", "flagsSZHVdec")[op == "dec"]
    if r == "(hl)":
        out.put("hl := cpu.get_hl()\n")
        out.put("n := cpu.mem.Read8(hl) %s\n" % delta)
        out.put("cpu.mem.Write8(hl, n)\n")
        out.put("cpu.F = (cpu.F & _CF) | %s[n]\n" % flags)
        out.put("return 11\n")
    elif r == "(ix+d)":
        out.put("adr := cpu.IX + offset16(cpu.get_n())\n")
        out.put("n := cpu.mem.Read8(adr) %s\n" % delta)
        out.put("cpu.mem.Write8(adr, n)\n")
        out.put("cpu.F = (cpu.F & _CF) | %s[n]\n" % flags)
        out.put("return 19\n")
    elif r == "(iy+d)":
        out.put("adr := cpu.IY + offset16(cpu.get_n())\n")
        out.put("n := cpu.mem.Read8(adr) %s\n" % delta)
        out.put("cpu.mem.Write8(adr,n)\n")
        out.put("cpu.F = (cpu.F & _CF) | %s[n]\n" % flags)
        out.put("return 19\n")
    else:
        out.put("n := cpu.%s %s\n" % (r.upper(), delta))
        out.put("cpu.%s = n\n" % r.upper())
        out.put("cpu.F = (cpu.F & _CF) | %s[n]\n" % flags)
        out.put("return 4\n")


def emit_alu_r(out, op, r):
    """alu operation with register"""
    if r == "(ix+d)":
        out.put("val := cpu.mem.Read8(cpu.IX + offset16(cpu.get_n()))\n")
        tclks = 15
    elif r == "(iy+d)":
        out.put("val := cpu.mem.Read8(cpu.IY + offset16(cpu.get_n()))\n")
        tclks = 15
    elif r == "(hl)":
        out.put("val := cpu.mem.Read8(cpu.get_hl())\n")
        tclks = 7
    else:
        out.put("val := cpu.%s\n" % r.upper())
        tclks = 4
    if op == "add":
        out.put("result := int(cpu.A) + int(val)\n")
        out.put("cpu.addFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
        out.put("return %d\n" % tclks)
    elif op == "adc":
        out.put("result := int(cpu.A) + int(val) + int(cpu.F & _CF)\n")
        out.put("cpu.addFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
        out.put("return %d\n" % tclks)
    elif op == "sub":
        out.put("result := int(cpu.A) - int(val)\n")
        out.put("cpu.subFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
        out.put("return %d\n" % tclks)
    elif op == "sbc":
        out.put("result := int(cpu.A) - int(val) - int(cpu.F & _CF)\n")
        out.put("cpu.subFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
        out.put("return %d\n" % tclks)
    elif op == "and":
        out.put("cpu.A &= val\n")
        out.put("cpu.F = flagsSZP[cpu.A] | _HF\n")
        out.put("return %d\n" % tclks)
    elif op == "xor":
        out.put("cpu.A ^= val\n")
        out.put("cpu.F = flagsSZP[cpu.A]\n")
        out.put("return %d\n" % tclks)
    elif op == "or":
        out.put("cpu.A |= val\n")
        out.put("cpu.F = flagsSZP[cpu.A]\n")
        out.put("return %d\n" % tclks)
    elif op == "cp":
        out.put("result := int(cpu.A) - int(val)\n")
        out.put("cpu.subFlags(result, val)\n")
        out.put("return %d\n" % tclks)
    else:
        assert False


def emit_alu_n(out, op):
    """alu operation with immediate"""
    out.put("val := cpu.get_n()\n")
    if op == "add":
        out.put("result := int(cpu.A) + int(val)\n")
        out.put("cpu.addFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
    elif op == "adc":
        out.put("result := int(cpu.A) + int(val) + int(cpu.F & _CF)\n")
        out.put("cpu.addFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
    elif op == "sub":
        out.put("result := int(cpu.A) - int(val)\n")
        out.put("cpu.subFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
    elif op == "sbc":
        out.put("result := int(cpu.A) - int(val) - int(cpu.F & _CF)\n")
        out.put("cpu.subFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
    elif op == "and":
        out.put("cpu.A &= val\n")
        out.put("cpu.F =  flagsSZP[cpu.A] | _HF\n")
    elif op == "xor":
        out.put("cpu.A ^= val\n")
        out.put("cpu.F = flagsSZP[cpu.A]\n")
    elif op == "or":
        out.put("cpu.A |= val\n")
        out.put("cpu.F = flagsSZP[cpu.A]\n")
    elif op == "cp":
        out.put("result := int(cpu.A) - int(val)\n")
        out.put("cpu.subFlags(result, val)\n")
    else:
        assert False
    out.put("return 7\n")


def emit_alu_hilo(out, op, dst, src):

    hi = src[2] == "h"
    select = ("& 0xff", ">> 8")[hi]
    src = src[:-1].upper()
    out.put(f"val := uint8(cpu.{src} {select})\n")

    if op == "add":
        out.put("result := int(cpu.A) + int(val)\n")
        out.put("cpu.addFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
    elif op == "adc":
        out.put("result := int(cpu.A) + int(val) + int(cpu.F & _CF)\n")
        out.put("cpu.addFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
    elif op == "sub":
        out.put("result := int(cpu.A) - int(val)\n")
        out.put("cpu.subFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
    elif op == "sbc":
        out.put("result := int(cpu.A) - int(val) - int(cpu.F & _CF)\n")
        out.put("cpu.subFlags(result, val)\n")
        out.put("cpu.A = uint8(result)\n")
    elif op == "and":
        out.put("cpu.A &= val\n")
        out.put("cpu.F = flagsSZP[cpu.A] | _HF\n")
    elif op == "xor":
        out.put("cpu.A ^= val\n")
        out.put("cpu.F = flagsSZP[cpu.A]\n")
    elif op == "or":
        out.put("cpu.A |= val\n")
        out.put("cpu.F = flagsSZP[cpu.A]\n")
    elif op == "cp":
        out.put("result := int(cpu.A) - int(val)\n")
        out.put("cpu.subFlags(result, val)\n")
    else:
        assert False
    out.put("return 4\n")


# -----------------------------------------------------------------------------
# General-Purpose Arithmetic and CPU Control Groups


def emit_nop(out):
    """nop"""
    out.put("return 4\n")


def emit_di(out):
    """disable interrupts"""
    out.put("cpu.IFF1 = false\n")
    out.put("cpu.IFF2 = false\n")
    out.put("return 4\n")


def emit_ei(out):
    """enable interrupts"""
    out.put("cpu.IFF1 = true\n")
    out.put("cpu.IFF2 = true\n")
    out.put("return 4\n")


def emit_im(out, n):
    """im n"""
    out.put("cpu.IM = %s\n" % n)
    out.put("return 4\n")


def emit_halt(out):
    """halt"""
    out.put("cpu.halt = true\n")
    out.put("return 4\n")


def emit_daa(out):
    """daa"""
    out.put("cf := byte2bool(cpu.F & _CF)\n")
    out.put("nf := byte2bool(cpu.F & _NF)\n")
    out.put("hf := byte2bool(cpu.F & _HF)\n")
    out.put("lo := cpu.A & 0xf\n")

    out.put("var correction uint8\n")
    out.put("var flags uint8\n")

    out.put("if nf {\n")
    out.put("	flags |= _NF\n")
    out.put("}\n")

    out.put("if hf || (lo > 9) {\n")
    out.put("	correction |= 0x06\n")
    out.put("}\n")

    out.put("if cf || (cpu.A > 0x99) {\n")
    out.put("	correction |= 0x60\n")
    out.put("	flags |= _CF\n")
    out.put("}\n")

    out.put("if nf {\n")
    out.put("	if hf && (lo < 6) {\n")
    out.put("		flags |= _HF\n")
    out.put("	}\n")
    out.put("} else {\n")
    out.put("	if lo >= 0x0A {\n")
    out.put("		flags |= _HF\n")
    out.put("	}\n")
    out.put("}\n")

    out.put("if nf {\n")
    out.put("	cpu.A -= correction\n")
    out.put("} else {\n")
    out.put("	cpu.A += correction\n")
    out.put("}\n")

    # Undocumented Y flag (Bit 3)
    # if (cpu.A & 0x08) != 0 {
    # 	flags |= _YF
    # }

    # Undocumented X flag (Bit 5)
    # if (cpu.A & 0x20) != 0 {
    # 	flags |= _XF
    # }

    out.put("cpu.F = flagsSZP[cpu.A] | flags\n")
    out.put("return 4\n")


def emit_neg(out):
    """neg"""
    out.put("cpu.F = _NF\n")
    out.put("if cpu.A != 0 {cpu.F |= _CF}\n")
    out.put("if (cpu.A & 0x0f) != 0 {cpu.F |= _HF}\n")
    out.put("if cpu.A == 0x80 {cpu.F |= _VF}\n")
    out.put("cpu.A = -cpu.A\n")
    out.put("cpu.F |= flagsSZ[cpu.A]\n")
    out.put("return 4\n")


# -----------------------------------------------------------------------------
# 16-Bit Arithmetic Group


def emit_op_rp_rp(out, op, d, s):
    """add/adc/sub hl/ix/iy,rp"""
    if s in _direct_rp:
        out.put("s := cpu.%s\n" % s.upper())
    else:
        out.put("s := cpu.get_%s()\n" % s)
    if d in _direct_rp:
        out.put("d := cpu.%s\n" % d.upper())
    else:
        out.put("d := cpu.get_%s()\n" % d)
    if op == "add":
        out.put("res := int(d) + int(s)\n")
        out.put("cpu.add16Flags(res, d, s)\n")
    elif op == "sbc":
        out.put("res := int(d) - int(s) - int(cpu.F & _CF)\n")
        out.put("cpu.sub16Flags(res, d, s)\n")
    elif op == "adc":
        out.put("res := int(d) + int(s) + int(cpu.F & _CF)\n")
        out.put("cpu.adc16Flags(res, d, s)\n")
    if d in _direct_rp:
        out.put("cpu.%s = uint16(res)\n" % d.upper())
    else:
        out.put("cpu.set_%s(uint16(res))\n" % d)
    out.put("return 11\n")


def emit_dec_rp(out, rp):
    """dec ss"""
    if rp in _direct_rp:
        out.put("cpu.%s -= 1\n" % rp.upper())
    else:
        out.put("cpu.set_%s(cpu.get_%s() - 1)\n" % (rp, rp))
    out.put("return 6\n")


def emit_inc_rp(out, rp):
    """inc ss"""
    if rp in _direct_rp:
        out.put("cpu.%s += 1\n" % rp.upper())
    else:
        out.put("cpu.set_%s(cpu.get_%s() + 1)\n" % (rp, rp))
    out.put("return 6\n")


def emit_inc_dec_r_hilo(out, r, op):
    """inc/dec high/low byte of ix or iy register"""
    flags = ("flagsSZHVinc", "flagsSZHVdec")[op == "dec"]
    hi = r[2] == "h"
    select = ("& 0xff", ">> 8")[hi]
    r = r[:-1].upper()
    if hi:
        delta = ("+= 0x0100", "-= 0x0100")[op == "dec"]
        out.put(f"cpu.{r} {delta}\n")
    else:
        delta = ("+ 1", "- 1")[op == "dec"]
        out.put(f"cpu.{r} = (cpu.{r} & 0xff00) | ((cpu.{r} {delta}) & 0xff)\n")
    out.put(f"cpu.F = (cpu.F & _CF) | {flags}[cpu.{r} {select}]\n")
    out.put("return 4\n")


# -----------------------------------------------------------------------------
# Rotate and Shift Group


def emit_rota(out, op):
    """rotate a"""
    if op == "rlca":
        out.put("cpu.A = ((cpu.A << 1) | (cpu.A >> 7)) & 0xff\n")
        out.put("cpu.F =  (cpu.F & (_SF | _ZF | _PF)) | (cpu.A & (_YF | _XF | _CF))\n")
    elif op == "rrca":
        out.put("cpu.F =  (cpu.F & (_SF | _ZF | _PF)) | (cpu.A & _CF)\n")
        out.put("cpu.A = ((cpu.A >> 1) | (cpu.A << 7)) & 0xff\n")
        out.put("cpu.F |= (cpu.A & (_YF | _XF))\n")
    elif op == "rla":
        out.put("res := (cpu.A << 1) | (cpu.F & _CF)\n")
        out.put("var c uint8\n")
        out.put("if (cpu.A & 0x80) != 0 {c = _CF}\n")
        out.put("cpu.F =  (cpu.F & (_SF | _ZF | _PF)) | c | (res & (_YF | _XF))\n")
        out.put("cpu.A = res & 0xff\n")
    elif op == "rra":
        out.put("res := (cpu.A >> 1) | (cpu.F << 7)\n")
        out.put("var c uint8\n")
        out.put("if (cpu.A & 0x01) != 0 {c = _CF}\n")
        out.put("cpu.F =  (cpu.F & (_SF | _ZF | _PF)) | c | (res & (_YF | _XF))\n")
        out.put("cpu.A = res & 0xff\n")
    elif op == "daa":
        return emit_daa(out)
    elif op == "cpl":
        out.put("cpu.A ^= 0xff\n")
        out.put("cpu.F =  (cpu.F & (_SF | _ZF | _PF | _CF)) | _HF | _NF | (cpu.A & (_YF | _XF))\n")
    elif op == "scf":
        out.put("cpu.F |= _CF\n")
        out.put("cpu.F = cpu.F &^ (_NF | _HF)\n")
    elif op == "ccf":
        out.put("cpu.F =  ((cpu.F & (_SF | _ZF | _PF | _CF)) | ((cpu.F & _CF) << 4) | (cpu.A & (_YF | _XF))) ^ _CF\n")
    else:
        assert False
    out.put("return 4\n")


def emit_rot_r_x(out, op, r, x):
    """rotate operation on r - optionally store in x also"""
    if r == "(ix+d)":
        out.put("res := cpu.mem.Read8(cpu.IX + offset16(d))\n")
    elif r == "(iy+d)":
        out.put("res := cpu.mem.Read8(cpu.IY + offset16(d))\n")
    elif r == "(hl)":
        out.put("res := cpu.mem.Read8(cpu.get_hl())\n")
    else:
        out.put("res := cpu.%s\n" % r.upper())

    if op == "rlc":
        out.put("var cf uint8\n")
        out.put("if (res & 0x80) != 0 {cf = _CF}\n")
        out.put("res = ((res << 1) | (res >> 7)) & 0xff\n")
    elif op == "rrc":
        out.put("var cf uint8\n")
        out.put("if (res & 0x01) != 0 {cf = _CF}\n")
        out.put("res = ((res >> 1) | (res << 7)) & 0xff\n")
    elif op == "rl":
        out.put("var cf uint8\n")
        out.put("if (res & 0x80) != 0 {cf = _CF}\n")
        out.put("res = ((res << 1) | (cpu.F & _CF)) & 0xff\n")
    elif op == "rr":
        out.put("var cf uint8\n")
        out.put("if (res & 0x01) != 0 {cf = _CF}\n")
        out.put("res = ((res >> 1) | (cpu.F << 7)) & 0xff\n")
    elif op == "sla":
        out.put("var cf uint8\n")
        out.put("if (res & 0x80) != 0 {cf = _CF}\n")
        out.put("res = (res << 1) & 0xff\n")
    elif op == "sra":
        out.put("var cf uint8\n")
        out.put("if (res & 0x01) != 0 {cf = _CF}\n")
        out.put("res = ((res >> 1) | (res & 0x80)) & 0xff\n")
    elif op == "sll":
        out.put("var cf uint8\n")
        out.put("if (res & 0x80) != 0 {cf = _CF}\n")
        out.put("res = ((res << 1) | 0x01) & 0xff\n")
    elif op == "srl":
        out.put("var cf uint8\n")
        out.put("if (res & 0x01) != 0 {cf = _CF}\n")
        out.put("res = (res >> 1) & 0xff\n")
    else:
        assert False

    out.put("cpu.F =  flagsSZP[res] | cf\n")
    if x != "":
        out.put("cpu.%s = res\n" % x.upper())
    if r == "(ix+d)":
        out.put("cpu.mem.Write8(cpu.IX + offset16(d), res)\n")
        out.put("return 11\n")
    elif r == "(iy+d)":
        out.put("cpu.mem.Write8(cpu.IY + offset16(d), res)\n")
        out.put("return 11\n")
    elif r == "(hl)":
        out.put("cpu.mem.Write8(cpu.get_hl(), res)\n")
        out.put("return 11\n")
    else:
        out.put("cpu.%s = res\n" % r.upper())
        out.put("return 4\n")


def emit_rxd(out, op):
    """rld, rrd"""
    out.put("adr := cpu.get_hl()\n")
    out.put("n := cpu.mem.Read8(adr)\n")
    if op == "rrd":
        out.put("cpu.mem.Write8(adr, ((n >> 4) | (cpu.A << 4)) & 0xff)\n")
        out.put("cpu.A = (cpu.A & 0xf0) | (n & 0x0f)\n")
    elif op == "rld":
        out.put("cpu.mem.Write8(adr, ((n << 4) | (cpu.A & 0x0f)) & 0xff)\n")
        out.put("cpu.A = (cpu.A & 0xf0) | (n >> 4)\n")
    else:
        assert False
    out.put("cpu.F =  (cpu.F & _CF) | flagsSZP[cpu.A]\n")
    out.put("return 14\n")


# -----------------------------------------------------------------------------
# Bit Set, Reset, and Test Group


def emit_bit_b_r(out, b, r):
    """bit test operation on r"""
    if r == "(ix+d)":
        out.put("bit := cpu.mem.Read8(cpu.IX + offset16(d)) & (1 << %d)\n" % b)
        t = 8
    elif r == "(iy+d)":
        out.put("bit := cpu.mem.Read8(cpu.IY + offset16(d)) & (1 << %d)\n" % b)
        t = 8
    elif r == "(hl)":
        out.put("bit := cpu.mem.Read8(cpu.get_hl()) & (1 << %d)\n" % b)
        t = 8
    else:
        out.put("bit := cpu.%s & (1 << %d)\n" % (r.upper(), b))
        t = 4
    out.put("var flags uint8\n")
    out.put("if bit == 0 {flags = _ZF | _PF} ")
    if b == 7:
        out.put("else {flags = _SF}\n")
    else:
        out.put("\n")
    out.put("cpu.F =  (cpu.F & _CF) | _HF | flags\n")
    out.put("return %d\n" % t)


def emit_set_b_r(out, b, r, x):
    """bit set operation on r"""
    if r == "(ix+d)":
        out.put("n := cpu.IX + offset16(d)\n")
        out.put("val := cpu.mem.Read8(n) | (1 << %d)\n" % b)
        out.put("cpu.mem.Write8(n, val)\n")
        t = 11
    elif r == "(iy+d)":
        out.put("n := cpu.IY + offset16(d)\n")
        out.put("val := cpu.mem.Read8(n) | (1 << %d)\n" % b)
        out.put("cpu.mem.Write8(n, val)\n")
        t = 11
    elif r == "(hl)":
        out.put("n := cpu.get_hl()\n")
        out.put("val := cpu.mem.Read8(n) | (1 << %d)\n" % b)
        out.put("cpu.mem.Write8(n, val)\n")
        t = 11
    else:
        out.put("val := cpu.%s | (1 << %d)\n" % (r.upper(), b))
        out.put("cpu.%s = val\n" % r.upper())
        t = 4
    if x != "":
        out.put("cpu.%s = val\n" % x.upper())
    out.put("return %d\n" % t)


def emit_res_b_r(out, b, r, x):
    """bit reset operation on r"""
    if r == "(ix+d)":
        out.put("n := cpu.IX + offset16(d)\n")
        out.put("val := cpu.mem.Read8(n) &^ (1 << %d)\n" % b)
        out.put("cpu.mem.Write8(n, val)\n")
        t = 11
    elif r == "(iy+d)":
        out.put("n := cpu.IY + offset16(d)\n")
        out.put("val := cpu.mem.Read8(n) &^ (1 << %d)\n" % b)
        out.put("cpu.mem.Write8(n, val)\n")
        t = 11
    elif r == "(hl)":
        out.put("n := cpu.get_hl()\n")
        out.put("val := cpu.mem.Read8(n) &^ (1 << %d)\n" % b)
        out.put("cpu.mem.Write8(n, val)\n")
        t = 11
    else:
        out.put("cpu.%s = cpu.%s &^ (1 << %d)\n" % (r.upper(), r.upper(), b))
        t = 4
    if x != "":
        out.put("cpu.%s = val\n" % x.upper())
    out.put("return %d\n" % t)


# -----------------------------------------------------------------------------
# Jump Group


def emit_jr_e(out):
    """jump relative"""
    out.put("cpu.PC += offset16(cpu.get_n())\n")
    out.put("return 12\n")


def emit_jr_cc_d(out, cc):
    """jump relative on condition"""
    out.put("ofs := offset16(cpu.get_n())\n")
    if cc == "nz":
        out.put("if (cpu.F & _ZF) == 0 {\n")
    elif cc == "z":
        out.put("if (cpu.F & _ZF) != 0 {\n")
    elif cc == "nc":
        out.put("if (cpu.F & _CF) == 0 {\n")
    elif cc == "c":
        out.put("if (cpu.F & _CF) != 0 {\n")
    else:
        assert False

    out.put("    cpu.PC += ofs\n")
    out.put("    return 12\n")
    out.put("}\n")

    out.put("return 7\n")


def emit_jp_nn(out):
    """jp nn"""
    out.put("cpu.PC = cpu.get_nn()\n")
    out.put("return 10\n")


def emit_jp_cc_nn(out, cc):
    """jp cc,nn"""
    out.put("nn := cpu.get_nn()\n")
    if cc == "nz":
        out.put("if (cpu.F & _ZF) == 0 {\n")
    elif cc == "z":
        out.put("if (cpu.F & _ZF) != 0 {\n")
    elif cc == "nc":
        out.put("if (cpu.F & _CF) == 0 {\n")
    elif cc == "c":
        out.put("if (cpu.F & _CF) != 0 {\n")
    elif cc == "po":
        out.put("if (cpu.F & _PF) == 0 {\n")
    elif cc == "pe":
        out.put("if (cpu.F & _PF) != 0 {\n")
    elif cc == "p":
        out.put("if (cpu.F & _SF) == 0 {\n")
    elif cc == "m":
        out.put("if (cpu.F & _SF) != 0 {\n")
    else:
        assert False
    out.put("    cpu.PC = nn\n")
    out.put("}\n")
    out.put("return 10\n")


def emit_jp_rp(out, rp):
    """jp rp"""
    if rp in _direct_rp:
        out.put("cpu.PC = cpu.%s\n" % rp.upper())
    else:
        out.put("cpu.PC = cpu.get_%s()\n" % rp)
    out.put("return 4\n")


def emit_djnz(out):
    """djnz e"""
    out.put("d := offset16(cpu.get_n())\n")
    out.put("cpu.B -= 1\n")
    out.put("if cpu.B != 0 {\n")
    out.put("    cpu.PC += d\n")
    out.put("    return 13}\n")
    out.put("return 8\n")


# -----------------------------------------------------------------------------
# Call And Return Group


def emit_call_nn(out):
    """call nn"""
    out.put("nn := cpu.get_nn()\n")
    out.put("cpu.push16(cpu.PC)\n")
    out.put("cpu.PC = nn\n")
    out.put("return 17\n")


def emit_call_cc_nn(out, cc):
    """call cc,nn"""
    out.put("nn := cpu.get_nn()\n")
    if cc == "nz":
        out.put("if (cpu.F & _ZF) == 0 {\n")
    elif cc == "z":
        out.put("if (cpu.F & _ZF) != 0 {\n")
    elif cc == "nc":
        out.put("if (cpu.F & _CF) == 0 {\n")
    elif cc == "c":
        out.put("if (cpu.F & _CF) != 0 {\n")
    elif cc == "po":
        out.put("if (cpu.F & _PF) == 0 {\n")
    elif cc == "pe":
        out.put("if (cpu.F & _PF) != 0 {\n")
    elif cc == "p":
        out.put("if (cpu.F & _SF) == 0 {\n")
    elif cc == "m":
        out.put("if (cpu.F & _SF) != 0 {\n")
    else:
        assert False
    out.put("    cpu.push16(cpu.PC)\n")
    out.put("    cpu.PC = nn\n")
    out.put("    return 17\n")
    out.put("}\n")
    out.put("return 10\n")


def emit_rst(out, p):
    out.put("cpu.push16(cpu.PC)\n")
    out.put("cpu.PC = 0x%02x\n" % p)
    out.put("return 11\n")


def emit_ret_cc(out, cc):
    """ret cc"""
    if cc == "nz":
        out.put("if (cpu.F & _ZF) == 0 {\n")
    elif cc == "z":
        out.put("if (cpu.F & _ZF) != 0 {\n")
    elif cc == "nc":
        out.put("if (cpu.F & _CF) == 0 {\n")
    elif cc == "c":
        out.put("if (cpu.F & _CF) != 0 {\n")
    elif cc == "po":
        out.put("if (cpu.F & _PF) == 0 {\n")
    elif cc == "pe":
        out.put("if (cpu.F & _PF) != 0 {\n")
    elif cc == "p":
        out.put("if (cpu.F & _SF) == 0 {\n")
    elif cc == "m":
        out.put("if (cpu.F & _SF) != 0 {\n")
    else:
        assert False
    out.put("    cpu.PC = cpu.pop16()\n")
    out.put("    return 11\n")
    out.put("}\n")
    out.put("return 5\n")


def emit_ret(out):
    """ret"""
    out.put("cpu.PC = cpu.pop16()\n")
    out.put("return 10\n")


def emit_retn(out):
    """retn"""
    out.put("cpu.IFF1 = cpu.IFF2\n")
    out.put("cpu.PC = cpu.pop16()\n")
    out.put("return 10\n")


def emit_reti(out):
    """reti"""
    out.put("cpu.IFF1 = cpu.IFF2\n")
    out.put("cpu.PC = cpu.pop16()\n")
    out.put("return 10\n")


# -----------------------------------------------------------------------------
# Input and Output Group


def emit_in_r_c(out, r):
    """in r,(c)"""
    out.put("val := cpu.io.Read8(cpu.get_bc())\n")
    if r != "":
        out.put("cpu.%s = val\n" % r.upper())
    out.put("cpu.F =  (cpu.F & _CF) | flagsSZP[val]\n")
    out.put("return 8\n")


def emit_in_a_n(out):
    """in a,(n)"""
    out.put("cpu.A = cpu.io.Read8((uint16(cpu.A) << 8) | uint16(cpu.get_n()))   \n")
    out.put("return 11\n")


def emit_out_n_a(out):
    """out (n),a"""
    out.put("cpu.io.Write8((uint16(cpu.A) << 8) | uint16(cpu.get_n()), cpu.A)\n")
    out.put("return 11\n")


def emit_out_c_r(out, r):
    if r == "":
        out.put("cpu.io.Write8(cpu.get_bc(), 0)\n")
    else:
        out.put("cpu.io.Write8(cpu.get_bc(), cpu.%s)\n" % r.upper())
    out.put("return 8\n")


# -----------------------------------------------------------------------------


def emit_unimplemented(out):
    """unimplemented instruction - crash"""
    out.put('panic("unimplemented instruction")\n')


# -----------------------------------------------------------------------------


def emit_normal(out, code):
    """
    Normal decode with no prefixes
    """
    m0 = code[0]
    x = (m0 >> 6) & 3
    y = (m0 >> 3) & 7
    z = (m0 >> 0) & 7
    p = (m0 >> 4) & 3
    q = (m0 >> 3) & 1

    if x == 0:
        if z == 0:
            if y == 0:
                return emit_nop(out)
            elif y == 1:
                return emit_ex_af_af(out)
            elif y == 2:
                return emit_djnz(out)
            elif y == 3:
                return emit_jr_e(out)
            else:
                return emit_jr_cc_d(out, _cc[y - 4])
        elif z == 1:
            if q == 0:
                return emit_ld_rp_nn(out, _rp[p])
            elif q == 1:
                return emit_op_rp_rp(out, "add", "hl", _rp[p])
        elif z == 2:
            if q == 0:
                if p == 0:
                    return emit_ld_mem_xx_a(out, "bc")
                elif p == 1:
                    return emit_ld_mem_xx_a(out, "de")
                elif p == 2:
                    return emit_ld_mem_nn_rp(out, "hl")
                else:
                    return emit_ld_mem_xx_a(out, "nn")
            else:
                if p == 0:
                    return emit_ld_a_mem_xx(out, "bc")
                elif p == 1:
                    return emit_ld_a_mem_xx(out, "de")
                elif p == 2:
                    emit_ld_rp_mem_nn(out, "hl")
                else:
                    return emit_ld_a_mem_xx(out, "nn")
        elif z == 3:
            if q == 0:
                return emit_inc_rp(out, _rp[p])
            else:
                return emit_dec_rp(out, _rp[p])
        elif z == 4:
            return emit_inc_dec_r(out, _r[y], "inc")
        elif z == 5:
            return emit_inc_dec_r(out, _r[y], "dec")
        elif z == 6:
            return emit_ld_r_n(out, _r[y])
        else:
            return emit_rota(out, _rota[y])
    elif x == 1:
        if (z == 6) and (y == 6):
            return emit_halt(out)
        else:
            return emit_ld_r_r(out, _r[y], _r[z])
    elif x == 2:
        return emit_alu_r(out, _alu[y], _r[z])
    else:
        if z == 0:
            return emit_ret_cc(out, _cc[y])
        elif z == 1:
            if q == 0:
                return emit_pop_rp(out, _rp2[p])
            else:
                if p == 0:
                    return emit_ret(out)
                elif p == 1:
                    return emit_exx(out)
                elif p == 2:
                    return emit_jp_rp(out, "hl")
                else:
                    return emit_ld_sp_hl(out)
        elif z == 2:
            return emit_jp_cc_nn(out, _cc[y])
        elif z == 3:
            if y == 0:
                return emit_jp_nn(out)
            elif y == 2:
                return emit_out_n_a(out)
            elif y == 3:
                return emit_in_a_n(out)
            elif y == 4:
                return emit_ex_mem_sp_r(out, "hl")
            elif y == 5:
                return emit_ex_de_hl(out)
            elif y == 6:
                return emit_di(out)
            else:
                return emit_ei(out)
        elif z == 4:
            return emit_call_cc_nn(out, _cc[y])
        elif z == 5:
            if q == 0:
                return emit_push_rp(out, _rp2[p])
            else:
                if p == 0:
                    return emit_call_nn(out)
        elif z == 6:
            return emit_alu_n(out, _alu[y])
        else:
            return emit_rst(out, (y << 3))


# -----------------------------------------------------------------------------


def emit_index(out, code, ir):
    """
    Decode with index register substitutions
    """
    m0 = code[0]
    x = (m0 >> 6) & 3
    y = (m0 >> 3) & 7
    z = (m0 >> 0) & 7
    p = (m0 >> 4) & 3
    q = (m0 >> 3) & 1

    # if using (hl) then: (hl)->(ix+d), h and l are unaffected.
    alt0_r = list(_r)
    alt0_r[6] = "(%s+d)" % ir

    # if not using (hl) then: hl->ix, h->ixh, l->ixl
    alt1_r = list(_r)
    alt1_r[4] = "%sh" % ir
    alt1_r[5] = "%sl" % ir

    alt_rp = list(_rp)
    alt_rp[2] = ir
    alt_rp2 = list(_rp2)
    alt_rp2[2] = ir

    if x == 0:
        if z == 0:
            if y == 0:
                return emit_nop(out)
            elif y == 1:
                return emit_ex_af_af(out)
            elif y == 2:
                return emit_djnz(out)
            elif y == 3:
                return emit_jr_e(out)
            else:
                return emit_jr_cc_d(out, _cc[y - 4])
        elif z == 1:
            if q == 0:
                return emit_ld_rp_nn(out, alt_rp[p])
            elif q == 1:
                return emit_op_rp_rp(out, "add", ir, alt_rp[p])
        elif z == 2:
            if q == 0:
                if p == 0:
                    # return ('ld', '(bc),a', 2)
                    return emit_unimplemented(out)
                elif p == 1:
                    # return ('ld', '(de),a', 2)
                    return emit_unimplemented(out)
                elif p == 2:
                    return emit_ld_mem_nn_rp(out, ir)
                else:
                    # return ('ld', '(%04x),a' % nn, 4)
                    return emit_unimplemented(out)
            else:
                if p == 0:
                    # return ('ld', 'a,(bc)', 2)
                    return emit_unimplemented(out)
                elif p == 1:
                    # return ('ld', 'a,(de)', 2)
                    return emit_unimplemented(out)
                elif p == 2:
                    return emit_ld_rp_mem_nn(out, ir)
                else:
                    # return ('ld', 'a,(%04x)' % nn, 4)
                    return emit_unimplemented(out)
        elif z == 3:
            if q == 0:
                return emit_inc_rp(out, alt_rp[p])
            else:
                return emit_dec_rp(out, alt_rp[p])
        elif z == 4:
            if y == 6:
                return emit_inc_dec_r(out, alt0_r[y], "inc")
            else:
                return emit_inc_dec_r_hilo(out, alt1_r[y], "inc")
        elif z == 5:
            if y == 6:
                return emit_inc_dec_r(out, alt0_r[y], "dec")
            else:
                return emit_inc_dec_r_hilo(out, alt1_r[y], "dec")
        elif z == 6:
            if y == 6:
                return emit_ld_mem_xx_n(out, ir)
            else:
                return emit_ld_hilo_immediate(out, alt1_r[y])
        else:
            # return (_rota[y], '', 2)
            return emit_unimplemented(out)
    elif x == 1:
        if (z == 6) and (y == 6):
            # return ('halt', '', 2)
            return emit_unimplemented(out)
        else:
            if (y == 6) or (z == 6):
                return emit_ld_r_r(out, alt0_r[y], alt0_r[z])
            else:
                return emit_ld_r_hilo(out, alt1_r[y], alt1_r[z])
    elif x == 2:
        if z == 6:
            return emit_alu_r(out, _alu[y], alt0_r[z])
        else:
            return emit_alu_hilo(out, _alu[y], _alux[y], alt1_r[z])
    else:
        if z == 0:
            # return ('ret', _cc[y], 2)
            return emit_unimplemented(out)
        elif z == 1:
            if q == 0:
                return emit_pop_rp(out, alt_rp2[p])
            else:
                if p == 0:
                    # return ('ret', '', 2)
                    return emit_unimplemented(out)
                elif p == 1:
                    return emit_exx(out)
                elif p == 2:
                    return emit_jp_rp(out, ir)
                else:
                    return emit_ld_index(out, ir)
        elif z == 2:
            return emit_jp_cc_nn(out, _cc[y])
        elif z == 3:
            if y == 0:
                # return ('jp', '%04x' % nn, 4)
                return emit_unimplemented(out)
            elif y == 2:
                return emit_out_n_a(out)
            elif y == 3:
                return emit_in_a_n(out)
            elif y == 4:
                return emit_ex_mem_sp_r(out, ir)
            elif y == 5:
                # return ('ex', 'de,hl', 2)
                return emit_unimplemented(out)
            elif y == 6:
                # return ('di', '', 2)
                return emit_unimplemented(out)
            else:
                # return ('ei', '', 2)
                return emit_unimplemented(out)
        elif z == 4:
            return emit_call_cc_nn(out, _cc[y])
        elif z == 5:
            if q == 0:
                return emit_push_rp(out, alt_rp2[p])
            else:
                if p == 0:
                    return emit_call_nn(out)
        elif z == 6:
            # return (_alu[y], '%s%02x' % (_alux[y], n0), 3)
            return emit_unimplemented(out)
        else:
            return emit_rst(out, (y << 3))


# -----------------------------------------------------------------------------


def emit_cb_prefix(out, code):
    """
    0xCB <opcode>
    """
    m0 = code[0]
    x = (m0 >> 6) & 3
    y = (m0 >> 3) & 7
    z = (m0 >> 0) & 7

    if x == 0:
        return emit_rot_r_x(out, _rot[y], _r[z], "")
    elif x == 1:
        return emit_bit_b_r(out, y, _r[z])
    elif x == 2:
        return emit_res_b_r(out, y, _r[z], "")
    else:
        return emit_set_b_r(out, y, _r[z], "")


# -----------------------------------------------------------------------------


def emit_ddcb_fdcb_prefix(out, code, ir):
    """
    0xDDCB <d> <opcode>
    0xFDCB <d> <opcode>
    """
    m1 = code[1]
    x = (m1 >> 6) & 3
    y = (m1 >> 3) & 7
    z = (m1 >> 0) & 7

    if x == 0:
        if z == 6:
            return emit_rot_r_x(out, _rot[y], "(%s+d)" % ir, "")
        else:
            return emit_rot_r_x(out, _rot[y], "(%s+d)" % ir, _r[z])
    elif x == 1:
        return emit_bit_b_r(out, y, "(%s+d)" % ir)
    elif x == 2:
        if z == 6:
            return emit_res_b_r(out, y, "(%s+d)" % ir, "")
        else:
            return emit_res_b_r(out, y, "(%s+d)" % ir, _r[z])
    else:
        if z == 6:
            return emit_set_b_r(out, y, "(%s+d)" % ir, "")
        else:
            return emit_set_b_r(out, y, "(%s+d)" % ir, _r[z])


# -----------------------------------------------------------------------------


def emit_ed_prefix(out, code):
    """
    0xED <opcode>
    0xED <opcode> <nn>
    """
    m0 = code[0]
    # m1 = code[1]
    # m2 = code[2]
    x = (m0 >> 6) & 3
    y = (m0 >> 3) & 7
    z = (m0 >> 0) & 7
    p = (m0 >> 4) & 3
    q = (m0 >> 3) & 1
    # nn = (m2 << 8) + m1

    if x == 1:
        if z == 0:
            if y == 6:
                return emit_in_r_c(out, "")
            else:
                return emit_in_r_c(out, _r[y])
        elif z == 1:
            if y == 6:
                return emit_out_c_r(out, "")
            else:
                return emit_out_c_r(out, _r[y])
        elif z == 2:
            if q == 0:
                return emit_op_rp_rp(out, "sbc", "hl", _rp[p])
            else:
                return emit_op_rp_rp(out, "adc", "hl", _rp[p])
        elif z == 3:
            if q == 0:
                return emit_ld_mem_nn_rp(out, _rp[p])
            else:
                return emit_ld_rp_mem_nn(out, _rp[p])
        elif z == 4:
            return emit_neg(out)
        elif z == 5:
            if y == 1:
                return emit_reti(out)
            else:
                return emit_retn(out)
        elif z == 6:
            return emit_im(out, _im[y])
        else:
            if y == 0:
                return emit_ld_ira(out, "i", "a")
            elif y == 1:
                return emit_ld_ira(out, "r", "a")
            elif y == 2:
                return emit_ld_ira(out, "a", "i")
            elif y == 3:
                return emit_ld_ira(out, "a", "r")
            elif y == 4:
                return emit_rxd(out, "rrd")
            elif y == 5:
                return emit_rxd(out, "rld")
            else:
                return emit_nop(out)
    elif x == 2:
        if (z <= 3) and (y >= 4):
            return emit_bli(out, _bli[z][y - 4])
    return emit_nop(out)


# -----------------------------------------------------------------------------


def emit_dd_fd_prefix(out, code, ir):
    """
    0xDD <x>
    0xFD <x>
    """
    m0 = code[0]
    if m0 in (0xDD, 0xED, 0xFD):
        emit_nop(out)
    elif m0 == 0xCB:
        return emit_ddcb_fdcb_prefix(out, code[1:], ir)
    else:
        return emit_index(out, code, ir)


# -----------------------------------------------------------------------------


def emit_instruction_code(out, code):
    """emit the code for an instruction"""
    m0 = code[0]
    if m0 == 0xCB:
        return emit_cb_prefix(out, code[1:])
    elif m0 == 0xDD:
        return emit_dd_fd_prefix(out, code[1:], "ix")
    elif m0 == 0xED:
        return emit_ed_prefix(out, code[1:])
    elif m0 == 0xFD:
        return emit_dd_fd_prefix(out, code[1:], "iy")
    else:
        return emit_normal(out, code)


# -----------------------------------------------------------------------------


def emit_opcode_table(out, idic, prefix, links, preamble, prototype):
    """emit a function table for each opcode with this prefix"""
    label = "_%s" % "".join(["%02x" % byte for byte in prefix])

    out.put("var opcodes%s = [256]" % (label, "")[len(label) == 1])
    out.put("%s {\n" % prototype)

    for opcode in range(0x100):
        code = list(prefix)
        code.append(opcode)
        # disassemble the instruction
        mem = memory.ram(4)
        mem.load(0, code)
        (operation, operands, nbytes) = z80da.disassemble(mem, 0)
        # add the instruction to the dictionary
        inst = " ".join((operation, operands))
        label = "".join(["%02x" % byte for byte in code])

        if opcode in links:
            out.put("(*CPU).execute_%s, " % label)
            out.put("// 0x%02x execute %s prefix\n" % (opcode, label))
        else:
            # add the inst/label to the dictionary if it is unique
            if not inst in idic:
                idic[inst] = (label, code, preamble)
            out.put("(*CPU).ins_%s, " % idic[inst][0])
            out.put("// 0x%02x %s\n" % (opcode, inst))

    out.put("}\n")


# -----------------------------------------------------------------------------


def emit_instruction_function(out, instruction, x):
    """emit the functon header and code for an instruction"""
    (label, code, preamble) = x
    out.put("// %s\n" % instruction)
    out.put("func (cpu *CPU) ins_%s%s {\n" % (label, preamble))
    emit_instruction_code(out, code)
    out.put("}\n")


# -----------------------------------------------------------------------------
# flag lookup tables

_CF = 0x01
_NF = 0x02
_PF = 0x04
_VF = _PF
_XF = 0x08
_HF = 0x10
_YF = 0x20
_ZF = 0x40
_SF = 0x80


def pop(x):
    """return number of 1's in x"""
    p = 0
    while x != 0:
        p += x & 1
        x >>= 1
    return p


def emit_table(out, name, data):
    out.put("var %s = [256]uint8{\n" % name)
    out.indent(1)
    for x in range(16):
        for y in range(16):
            out.put("0x%02x, " % data[(x * 16) + y])
        out.put("\n")
    out.outdent(1)
    out.put("}\n")


def emit_flag_tables(out):

    SZ = []
    SZ_BIT = []
    SZP = []
    SZHV_inc = []
    SZHV_dec = []

    for i in range(0x100):
        p = pop(i)
        if i:
            SZ.append(i & _SF)
            SZ_BIT.append(i & _SF)
        else:
            SZ.append(_ZF)
            SZ_BIT.append(_ZF | _PF)
        # undocumented flag bits 5+3
        SZ[i] |= i & (_YF | _XF)
        SZ_BIT[i] |= i & (_YF | _XF)
        # parity
        SZP.append(SZ[i])
        if (p & 1) == 0:
            SZP[i] |= _PF
        # increment
        SZHV_inc.append(SZ[i])
        if i == 0x80:
            SZHV_inc[i] |= _VF
        if (i & 0x0F) == 0:
            SZHV_inc[i] |= _HF
        # decrement
        SZHV_dec.append(SZ[i] | _NF)
        if i == 0x7F:
            SZHV_dec[i] |= _VF
        if (i & 0x0F) == 0x0F:
            SZHV_dec[i] |= _HF

    out.indent(2)
    emit_table(out, "flagsSZ", SZ)
    emit_table(out, "flagsSZP", SZP)
    emit_table(out, "flagsSZHVinc", SZHV_inc)
    emit_table(out, "flagsSZHVdec", SZHV_dec)
    out.outdent(2)


# -----------------------------------------------------------------------------


def generate(ofname):
    """generate the opcode emulation file"""
    out = output(ofname)
    out.put("package z80\n")

    # generate flag tables
    emit_flag_tables(out)
    idic = {}
    # generate the opcode tables
    for prefix, links, preamble, prototype in _prefixes:
        emit_opcode_table(out, idic, prefix, links, preamble, prototype)
    # generate the instruction functions
    for k, v in idic.items():
        emit_instruction_function(out, k, v)
    out.close()


# -----------------------------------------------------------------------------


def usage():
    print("usage:")
    print("%s -o [OUTPUT]" % sys.argv[0])
    sys.exit(2)


# -----------------------------------------------------------------------------


def main():
    ofname = "opcodes.go"
    try:
        optlist, arglist = getopt.gnu_getopt(sys.argv[1:], "o:")
    except getopt.GetoptError:
        usage()
    for opt in optlist:
        if opt[0] == "-o":
            ofname = opt[1]
    if len(arglist) != 0:
        usage()
    generate(ofname)


# -----------------------------------------------------------------------------

if __name__ == "__main__":
    main()

# -----------------------------------------------------------------------------
