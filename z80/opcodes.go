package z80

var flagsSZ = [256]uint8{
	0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88,
	0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88,
	0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8,
	0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8,
	0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88,
	0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88,
	0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8,
	0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8,
}
var flagsSZP = [256]uint8{
	0x44, 0x00, 0x00, 0x04, 0x00, 0x04, 0x04, 0x00, 0x08, 0x0c, 0x0c, 0x08, 0x0c, 0x08, 0x08, 0x0c,
	0x00, 0x04, 0x04, 0x00, 0x04, 0x00, 0x00, 0x04, 0x0c, 0x08, 0x08, 0x0c, 0x08, 0x0c, 0x0c, 0x08,
	0x20, 0x24, 0x24, 0x20, 0x24, 0x20, 0x20, 0x24, 0x2c, 0x28, 0x28, 0x2c, 0x28, 0x2c, 0x2c, 0x28,
	0x24, 0x20, 0x20, 0x24, 0x20, 0x24, 0x24, 0x20, 0x28, 0x2c, 0x2c, 0x28, 0x2c, 0x28, 0x28, 0x2c,
	0x00, 0x04, 0x04, 0x00, 0x04, 0x00, 0x00, 0x04, 0x0c, 0x08, 0x08, 0x0c, 0x08, 0x0c, 0x0c, 0x08,
	0x04, 0x00, 0x00, 0x04, 0x00, 0x04, 0x04, 0x00, 0x08, 0x0c, 0x0c, 0x08, 0x0c, 0x08, 0x08, 0x0c,
	0x24, 0x20, 0x20, 0x24, 0x20, 0x24, 0x24, 0x20, 0x28, 0x2c, 0x2c, 0x28, 0x2c, 0x28, 0x28, 0x2c,
	0x20, 0x24, 0x24, 0x20, 0x24, 0x20, 0x20, 0x24, 0x2c, 0x28, 0x28, 0x2c, 0x28, 0x2c, 0x2c, 0x28,
	0x80, 0x84, 0x84, 0x80, 0x84, 0x80, 0x80, 0x84, 0x8c, 0x88, 0x88, 0x8c, 0x88, 0x8c, 0x8c, 0x88,
	0x84, 0x80, 0x80, 0x84, 0x80, 0x84, 0x84, 0x80, 0x88, 0x8c, 0x8c, 0x88, 0x8c, 0x88, 0x88, 0x8c,
	0xa4, 0xa0, 0xa0, 0xa4, 0xa0, 0xa4, 0xa4, 0xa0, 0xa8, 0xac, 0xac, 0xa8, 0xac, 0xa8, 0xa8, 0xac,
	0xa0, 0xa4, 0xa4, 0xa0, 0xa4, 0xa0, 0xa0, 0xa4, 0xac, 0xa8, 0xa8, 0xac, 0xa8, 0xac, 0xac, 0xa8,
	0x84, 0x80, 0x80, 0x84, 0x80, 0x84, 0x84, 0x80, 0x88, 0x8c, 0x8c, 0x88, 0x8c, 0x88, 0x88, 0x8c,
	0x80, 0x84, 0x84, 0x80, 0x84, 0x80, 0x80, 0x84, 0x8c, 0x88, 0x88, 0x8c, 0x88, 0x8c, 0x8c, 0x88,
	0xa0, 0xa4, 0xa4, 0xa0, 0xa4, 0xa0, 0xa0, 0xa4, 0xac, 0xa8, 0xa8, 0xac, 0xa8, 0xac, 0xac, 0xa8,
	0xa4, 0xa0, 0xa0, 0xa4, 0xa0, 0xa4, 0xa4, 0xa0, 0xa8, 0xac, 0xac, 0xa8, 0xac, 0xa8, 0xa8, 0xac,
}
var flagsSZHVinc = [256]uint8{
	0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x94, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88,
	0x90, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88,
	0xb0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8,
	0xb0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8,
	0x90, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88,
	0x90, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88,
	0xb0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8,
	0xb0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8, 0xa8,
}
var flagsSZHVdec = [256]uint8{
	0x42, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x1a,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x1a,
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x3a,
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x3a,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x1a,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x1a,
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x3a,
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x2a, 0x3e,
	0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x9a,
	0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x9a,
	0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xba,
	0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xba,
	0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x9a,
	0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x82, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x8a, 0x9a,
	0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xba,
	0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xa2, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xba,
}
var opcodes = [256]func(*CPU) int{
	(*CPU).ins_00,     // 0x00 nop
	(*CPU).ins_01,     // 0x01 ld bc,0000
	(*CPU).ins_02,     // 0x02 ld (bc),a
	(*CPU).ins_03,     // 0x03 inc bc
	(*CPU).ins_04,     // 0x04 inc b
	(*CPU).ins_05,     // 0x05 dec b
	(*CPU).ins_06,     // 0x06 ld b,00
	(*CPU).ins_07,     // 0x07 rlca
	(*CPU).ins_08,     // 0x08 ex af,af'
	(*CPU).ins_09,     // 0x09 add hl,bc
	(*CPU).ins_0a,     // 0x0a ld a,(bc)
	(*CPU).ins_0b,     // 0x0b dec bc
	(*CPU).ins_0c,     // 0x0c inc c
	(*CPU).ins_0d,     // 0x0d dec c
	(*CPU).ins_0e,     // 0x0e ld c,00
	(*CPU).ins_0f,     // 0x0f rrca
	(*CPU).ins_10,     // 0x10 djnz 0002
	(*CPU).ins_11,     // 0x11 ld de,0000
	(*CPU).ins_12,     // 0x12 ld (de),a
	(*CPU).ins_13,     // 0x13 inc de
	(*CPU).ins_14,     // 0x14 inc d
	(*CPU).ins_15,     // 0x15 dec d
	(*CPU).ins_16,     // 0x16 ld d,00
	(*CPU).ins_17,     // 0x17 rla
	(*CPU).ins_18,     // 0x18 jr 0002
	(*CPU).ins_19,     // 0x19 add hl,de
	(*CPU).ins_1a,     // 0x1a ld a,(de)
	(*CPU).ins_1b,     // 0x1b dec de
	(*CPU).ins_1c,     // 0x1c inc e
	(*CPU).ins_1d,     // 0x1d dec e
	(*CPU).ins_1e,     // 0x1e ld e,00
	(*CPU).ins_1f,     // 0x1f rra
	(*CPU).ins_20,     // 0x20 jr nz,0002
	(*CPU).ins_21,     // 0x21 ld hl,0000
	(*CPU).ins_22,     // 0x22 ld (0000),hl
	(*CPU).ins_23,     // 0x23 inc hl
	(*CPU).ins_24,     // 0x24 inc h
	(*CPU).ins_25,     // 0x25 dec h
	(*CPU).ins_26,     // 0x26 ld h,00
	(*CPU).ins_27,     // 0x27 daa
	(*CPU).ins_28,     // 0x28 jr z,0002
	(*CPU).ins_29,     // 0x29 add hl,hl
	(*CPU).ins_2a,     // 0x2a ld hl,(0000)
	(*CPU).ins_2b,     // 0x2b dec hl
	(*CPU).ins_2c,     // 0x2c inc l
	(*CPU).ins_2d,     // 0x2d dec l
	(*CPU).ins_2e,     // 0x2e ld l,00
	(*CPU).ins_2f,     // 0x2f cpl
	(*CPU).ins_30,     // 0x30 jr nc,0002
	(*CPU).ins_31,     // 0x31 ld sp,0000
	(*CPU).ins_32,     // 0x32 ld (0000),a
	(*CPU).ins_33,     // 0x33 inc sp
	(*CPU).ins_34,     // 0x34 inc (hl)
	(*CPU).ins_35,     // 0x35 dec (hl)
	(*CPU).ins_36,     // 0x36 ld (hl),00
	(*CPU).ins_37,     // 0x37 scf
	(*CPU).ins_38,     // 0x38 jr c,0002
	(*CPU).ins_39,     // 0x39 add hl,sp
	(*CPU).ins_3a,     // 0x3a ld a,(0000)
	(*CPU).ins_3b,     // 0x3b dec sp
	(*CPU).ins_3c,     // 0x3c inc a
	(*CPU).ins_3d,     // 0x3d dec a
	(*CPU).ins_3e,     // 0x3e ld a,00
	(*CPU).ins_3f,     // 0x3f ccf
	(*CPU).ins_40,     // 0x40 ld b,b
	(*CPU).ins_41,     // 0x41 ld b,c
	(*CPU).ins_42,     // 0x42 ld b,d
	(*CPU).ins_43,     // 0x43 ld b,e
	(*CPU).ins_44,     // 0x44 ld b,h
	(*CPU).ins_45,     // 0x45 ld b,l
	(*CPU).ins_46,     // 0x46 ld b,(hl)
	(*CPU).ins_47,     // 0x47 ld b,a
	(*CPU).ins_48,     // 0x48 ld c,b
	(*CPU).ins_49,     // 0x49 ld c,c
	(*CPU).ins_4a,     // 0x4a ld c,d
	(*CPU).ins_4b,     // 0x4b ld c,e
	(*CPU).ins_4c,     // 0x4c ld c,h
	(*CPU).ins_4d,     // 0x4d ld c,l
	(*CPU).ins_4e,     // 0x4e ld c,(hl)
	(*CPU).ins_4f,     // 0x4f ld c,a
	(*CPU).ins_50,     // 0x50 ld d,b
	(*CPU).ins_51,     // 0x51 ld d,c
	(*CPU).ins_52,     // 0x52 ld d,d
	(*CPU).ins_53,     // 0x53 ld d,e
	(*CPU).ins_54,     // 0x54 ld d,h
	(*CPU).ins_55,     // 0x55 ld d,l
	(*CPU).ins_56,     // 0x56 ld d,(hl)
	(*CPU).ins_57,     // 0x57 ld d,a
	(*CPU).ins_58,     // 0x58 ld e,b
	(*CPU).ins_59,     // 0x59 ld e,c
	(*CPU).ins_5a,     // 0x5a ld e,d
	(*CPU).ins_5b,     // 0x5b ld e,e
	(*CPU).ins_5c,     // 0x5c ld e,h
	(*CPU).ins_5d,     // 0x5d ld e,l
	(*CPU).ins_5e,     // 0x5e ld e,(hl)
	(*CPU).ins_5f,     // 0x5f ld e,a
	(*CPU).ins_60,     // 0x60 ld h,b
	(*CPU).ins_61,     // 0x61 ld h,c
	(*CPU).ins_62,     // 0x62 ld h,d
	(*CPU).ins_63,     // 0x63 ld h,e
	(*CPU).ins_64,     // 0x64 ld h,h
	(*CPU).ins_65,     // 0x65 ld h,l
	(*CPU).ins_66,     // 0x66 ld h,(hl)
	(*CPU).ins_67,     // 0x67 ld h,a
	(*CPU).ins_68,     // 0x68 ld l,b
	(*CPU).ins_69,     // 0x69 ld l,c
	(*CPU).ins_6a,     // 0x6a ld l,d
	(*CPU).ins_6b,     // 0x6b ld l,e
	(*CPU).ins_6c,     // 0x6c ld l,h
	(*CPU).ins_6d,     // 0x6d ld l,l
	(*CPU).ins_6e,     // 0x6e ld l,(hl)
	(*CPU).ins_6f,     // 0x6f ld l,a
	(*CPU).ins_70,     // 0x70 ld (hl),b
	(*CPU).ins_71,     // 0x71 ld (hl),c
	(*CPU).ins_72,     // 0x72 ld (hl),d
	(*CPU).ins_73,     // 0x73 ld (hl),e
	(*CPU).ins_74,     // 0x74 ld (hl),h
	(*CPU).ins_75,     // 0x75 ld (hl),l
	(*CPU).ins_76,     // 0x76 halt
	(*CPU).ins_77,     // 0x77 ld (hl),a
	(*CPU).ins_78,     // 0x78 ld a,b
	(*CPU).ins_79,     // 0x79 ld a,c
	(*CPU).ins_7a,     // 0x7a ld a,d
	(*CPU).ins_7b,     // 0x7b ld a,e
	(*CPU).ins_7c,     // 0x7c ld a,h
	(*CPU).ins_7d,     // 0x7d ld a,l
	(*CPU).ins_7e,     // 0x7e ld a,(hl)
	(*CPU).ins_7f,     // 0x7f ld a,a
	(*CPU).ins_80,     // 0x80 add a,b
	(*CPU).ins_81,     // 0x81 add a,c
	(*CPU).ins_82,     // 0x82 add a,d
	(*CPU).ins_83,     // 0x83 add a,e
	(*CPU).ins_84,     // 0x84 add a,h
	(*CPU).ins_85,     // 0x85 add a,l
	(*CPU).ins_86,     // 0x86 add a,(hl)
	(*CPU).ins_87,     // 0x87 add a,a
	(*CPU).ins_88,     // 0x88 adc a,b
	(*CPU).ins_89,     // 0x89 adc a,c
	(*CPU).ins_8a,     // 0x8a adc a,d
	(*CPU).ins_8b,     // 0x8b adc a,e
	(*CPU).ins_8c,     // 0x8c adc a,h
	(*CPU).ins_8d,     // 0x8d adc a,l
	(*CPU).ins_8e,     // 0x8e adc a,(hl)
	(*CPU).ins_8f,     // 0x8f adc a,a
	(*CPU).ins_90,     // 0x90 sub b
	(*CPU).ins_91,     // 0x91 sub c
	(*CPU).ins_92,     // 0x92 sub d
	(*CPU).ins_93,     // 0x93 sub e
	(*CPU).ins_94,     // 0x94 sub h
	(*CPU).ins_95,     // 0x95 sub l
	(*CPU).ins_96,     // 0x96 sub (hl)
	(*CPU).ins_97,     // 0x97 sub a
	(*CPU).ins_98,     // 0x98 sbc a,b
	(*CPU).ins_99,     // 0x99 sbc a,c
	(*CPU).ins_9a,     // 0x9a sbc a,d
	(*CPU).ins_9b,     // 0x9b sbc a,e
	(*CPU).ins_9c,     // 0x9c sbc a,h
	(*CPU).ins_9d,     // 0x9d sbc a,l
	(*CPU).ins_9e,     // 0x9e sbc a,(hl)
	(*CPU).ins_9f,     // 0x9f sbc a,a
	(*CPU).ins_a0,     // 0xa0 and b
	(*CPU).ins_a1,     // 0xa1 and c
	(*CPU).ins_a2,     // 0xa2 and d
	(*CPU).ins_a3,     // 0xa3 and e
	(*CPU).ins_a4,     // 0xa4 and h
	(*CPU).ins_a5,     // 0xa5 and l
	(*CPU).ins_a6,     // 0xa6 and (hl)
	(*CPU).ins_a7,     // 0xa7 and a
	(*CPU).ins_a8,     // 0xa8 xor b
	(*CPU).ins_a9,     // 0xa9 xor c
	(*CPU).ins_aa,     // 0xaa xor d
	(*CPU).ins_ab,     // 0xab xor e
	(*CPU).ins_ac,     // 0xac xor h
	(*CPU).ins_ad,     // 0xad xor l
	(*CPU).ins_ae,     // 0xae xor (hl)
	(*CPU).ins_af,     // 0xaf xor a
	(*CPU).ins_b0,     // 0xb0 or b
	(*CPU).ins_b1,     // 0xb1 or c
	(*CPU).ins_b2,     // 0xb2 or d
	(*CPU).ins_b3,     // 0xb3 or e
	(*CPU).ins_b4,     // 0xb4 or h
	(*CPU).ins_b5,     // 0xb5 or l
	(*CPU).ins_b6,     // 0xb6 or (hl)
	(*CPU).ins_b7,     // 0xb7 or a
	(*CPU).ins_b8,     // 0xb8 cp b
	(*CPU).ins_b9,     // 0xb9 cp c
	(*CPU).ins_ba,     // 0xba cp d
	(*CPU).ins_bb,     // 0xbb cp e
	(*CPU).ins_bc,     // 0xbc cp h
	(*CPU).ins_bd,     // 0xbd cp l
	(*CPU).ins_be,     // 0xbe cp (hl)
	(*CPU).ins_bf,     // 0xbf cp a
	(*CPU).ins_c0,     // 0xc0 ret nz
	(*CPU).ins_c1,     // 0xc1 pop bc
	(*CPU).ins_c2,     // 0xc2 jp nz,0000
	(*CPU).ins_c3,     // 0xc3 jp 0000
	(*CPU).ins_c4,     // 0xc4 call nz,0000
	(*CPU).ins_c5,     // 0xc5 push bc
	(*CPU).ins_c6,     // 0xc6 add a,00
	(*CPU).ins_c7,     // 0xc7 rst 00
	(*CPU).ins_c8,     // 0xc8 ret z
	(*CPU).ins_c9,     // 0xc9 ret
	(*CPU).ins_ca,     // 0xca jp z,0000
	(*CPU).execute_cb, // 0xcb execute cb prefix
	(*CPU).ins_cc,     // 0xcc call z,0000
	(*CPU).ins_cd,     // 0xcd call 0000
	(*CPU).ins_ce,     // 0xce adc a,00
	(*CPU).ins_cf,     // 0xcf rst 08
	(*CPU).ins_d0,     // 0xd0 ret nc
	(*CPU).ins_d1,     // 0xd1 pop de
	(*CPU).ins_d2,     // 0xd2 jp nc,0000
	(*CPU).ins_d3,     // 0xd3 out (00),a
	(*CPU).ins_d4,     // 0xd4 call nc,0000
	(*CPU).ins_d5,     // 0xd5 push de
	(*CPU).ins_d6,     // 0xd6 sub 00
	(*CPU).ins_d7,     // 0xd7 rst 10
	(*CPU).ins_d8,     // 0xd8 ret c
	(*CPU).ins_d9,     // 0xd9 exx
	(*CPU).ins_da,     // 0xda jp c,0000
	(*CPU).ins_db,     // 0xdb in a,(00)
	(*CPU).ins_dc,     // 0xdc call c,0000
	(*CPU).execute_dd, // 0xdd execute dd prefix
	(*CPU).ins_de,     // 0xde sbc a,00
	(*CPU).ins_df,     // 0xdf rst 18
	(*CPU).ins_e0,     // 0xe0 ret po
	(*CPU).ins_e1,     // 0xe1 pop hl
	(*CPU).ins_e2,     // 0xe2 jp po,0000
	(*CPU).ins_e3,     // 0xe3 ex (sp),hl
	(*CPU).ins_e4,     // 0xe4 call po,0000
	(*CPU).ins_e5,     // 0xe5 push hl
	(*CPU).ins_e6,     // 0xe6 and 00
	(*CPU).ins_e7,     // 0xe7 rst 20
	(*CPU).ins_e8,     // 0xe8 ret pe
	(*CPU).ins_e9,     // 0xe9 jp hl
	(*CPU).ins_ea,     // 0xea jp pe,0000
	(*CPU).ins_eb,     // 0xeb ex de,hl
	(*CPU).ins_ec,     // 0xec call pe,0000
	(*CPU).execute_ed, // 0xed execute ed prefix
	(*CPU).ins_ee,     // 0xee xor 00
	(*CPU).ins_ef,     // 0xef rst 28
	(*CPU).ins_f0,     // 0xf0 ret p
	(*CPU).ins_f1,     // 0xf1 pop af
	(*CPU).ins_f2,     // 0xf2 jp p,0000
	(*CPU).ins_f3,     // 0xf3 di
	(*CPU).ins_f4,     // 0xf4 call p,0000
	(*CPU).ins_f5,     // 0xf5 push af
	(*CPU).ins_f6,     // 0xf6 or 00
	(*CPU).ins_f7,     // 0xf7 rst 30
	(*CPU).ins_f8,     // 0xf8 ret m
	(*CPU).ins_f9,     // 0xf9 ld sp,hl
	(*CPU).ins_fa,     // 0xfa jp m,0000
	(*CPU).ins_fb,     // 0xfb ei
	(*CPU).ins_fc,     // 0xfc call m,0000
	(*CPU).execute_fd, // 0xfd execute fd prefix
	(*CPU).ins_fe,     // 0xfe cp 00
	(*CPU).ins_ff,     // 0xff rst 38
}
var opcodes_cb = [256]func(*CPU) int{
	(*CPU).ins_cb00, // 0x00 rlc b
	(*CPU).ins_cb01, // 0x01 rlc c
	(*CPU).ins_cb02, // 0x02 rlc d
	(*CPU).ins_cb03, // 0x03 rlc e
	(*CPU).ins_cb04, // 0x04 rlc h
	(*CPU).ins_cb05, // 0x05 rlc l
	(*CPU).ins_cb06, // 0x06 rlc (hl)
	(*CPU).ins_cb07, // 0x07 rlc a
	(*CPU).ins_cb08, // 0x08 rrc b
	(*CPU).ins_cb09, // 0x09 rrc c
	(*CPU).ins_cb0a, // 0x0a rrc d
	(*CPU).ins_cb0b, // 0x0b rrc e
	(*CPU).ins_cb0c, // 0x0c rrc h
	(*CPU).ins_cb0d, // 0x0d rrc l
	(*CPU).ins_cb0e, // 0x0e rrc (hl)
	(*CPU).ins_cb0f, // 0x0f rrc a
	(*CPU).ins_cb10, // 0x10 rl b
	(*CPU).ins_cb11, // 0x11 rl c
	(*CPU).ins_cb12, // 0x12 rl d
	(*CPU).ins_cb13, // 0x13 rl e
	(*CPU).ins_cb14, // 0x14 rl h
	(*CPU).ins_cb15, // 0x15 rl l
	(*CPU).ins_cb16, // 0x16 rl (hl)
	(*CPU).ins_cb17, // 0x17 rl a
	(*CPU).ins_cb18, // 0x18 rr b
	(*CPU).ins_cb19, // 0x19 rr c
	(*CPU).ins_cb1a, // 0x1a rr d
	(*CPU).ins_cb1b, // 0x1b rr e
	(*CPU).ins_cb1c, // 0x1c rr h
	(*CPU).ins_cb1d, // 0x1d rr l
	(*CPU).ins_cb1e, // 0x1e rr (hl)
	(*CPU).ins_cb1f, // 0x1f rr a
	(*CPU).ins_cb20, // 0x20 sla b
	(*CPU).ins_cb21, // 0x21 sla c
	(*CPU).ins_cb22, // 0x22 sla d
	(*CPU).ins_cb23, // 0x23 sla e
	(*CPU).ins_cb24, // 0x24 sla h
	(*CPU).ins_cb25, // 0x25 sla l
	(*CPU).ins_cb26, // 0x26 sla (hl)
	(*CPU).ins_cb27, // 0x27 sla a
	(*CPU).ins_cb28, // 0x28 sra b
	(*CPU).ins_cb29, // 0x29 sra c
	(*CPU).ins_cb2a, // 0x2a sra d
	(*CPU).ins_cb2b, // 0x2b sra e
	(*CPU).ins_cb2c, // 0x2c sra h
	(*CPU).ins_cb2d, // 0x2d sra l
	(*CPU).ins_cb2e, // 0x2e sra (hl)
	(*CPU).ins_cb2f, // 0x2f sra a
	(*CPU).ins_cb30, // 0x30 sll b
	(*CPU).ins_cb31, // 0x31 sll c
	(*CPU).ins_cb32, // 0x32 sll d
	(*CPU).ins_cb33, // 0x33 sll e
	(*CPU).ins_cb34, // 0x34 sll h
	(*CPU).ins_cb35, // 0x35 sll l
	(*CPU).ins_cb36, // 0x36 sll (hl)
	(*CPU).ins_cb37, // 0x37 sll a
	(*CPU).ins_cb38, // 0x38 srl b
	(*CPU).ins_cb39, // 0x39 srl c
	(*CPU).ins_cb3a, // 0x3a srl d
	(*CPU).ins_cb3b, // 0x3b srl e
	(*CPU).ins_cb3c, // 0x3c srl h
	(*CPU).ins_cb3d, // 0x3d srl l
	(*CPU).ins_cb3e, // 0x3e srl (hl)
	(*CPU).ins_cb3f, // 0x3f srl a
	(*CPU).ins_cb40, // 0x40 bit 0,b
	(*CPU).ins_cb41, // 0x41 bit 0,c
	(*CPU).ins_cb42, // 0x42 bit 0,d
	(*CPU).ins_cb43, // 0x43 bit 0,e
	(*CPU).ins_cb44, // 0x44 bit 0,h
	(*CPU).ins_cb45, // 0x45 bit 0,l
	(*CPU).ins_cb46, // 0x46 bit 0,(hl)
	(*CPU).ins_cb47, // 0x47 bit 0,a
	(*CPU).ins_cb48, // 0x48 bit 1,b
	(*CPU).ins_cb49, // 0x49 bit 1,c
	(*CPU).ins_cb4a, // 0x4a bit 1,d
	(*CPU).ins_cb4b, // 0x4b bit 1,e
	(*CPU).ins_cb4c, // 0x4c bit 1,h
	(*CPU).ins_cb4d, // 0x4d bit 1,l
	(*CPU).ins_cb4e, // 0x4e bit 1,(hl)
	(*CPU).ins_cb4f, // 0x4f bit 1,a
	(*CPU).ins_cb50, // 0x50 bit 2,b
	(*CPU).ins_cb51, // 0x51 bit 2,c
	(*CPU).ins_cb52, // 0x52 bit 2,d
	(*CPU).ins_cb53, // 0x53 bit 2,e
	(*CPU).ins_cb54, // 0x54 bit 2,h
	(*CPU).ins_cb55, // 0x55 bit 2,l
	(*CPU).ins_cb56, // 0x56 bit 2,(hl)
	(*CPU).ins_cb57, // 0x57 bit 2,a
	(*CPU).ins_cb58, // 0x58 bit 3,b
	(*CPU).ins_cb59, // 0x59 bit 3,c
	(*CPU).ins_cb5a, // 0x5a bit 3,d
	(*CPU).ins_cb5b, // 0x5b bit 3,e
	(*CPU).ins_cb5c, // 0x5c bit 3,h
	(*CPU).ins_cb5d, // 0x5d bit 3,l
	(*CPU).ins_cb5e, // 0x5e bit 3,(hl)
	(*CPU).ins_cb5f, // 0x5f bit 3,a
	(*CPU).ins_cb60, // 0x60 bit 4,b
	(*CPU).ins_cb61, // 0x61 bit 4,c
	(*CPU).ins_cb62, // 0x62 bit 4,d
	(*CPU).ins_cb63, // 0x63 bit 4,e
	(*CPU).ins_cb64, // 0x64 bit 4,h
	(*CPU).ins_cb65, // 0x65 bit 4,l
	(*CPU).ins_cb66, // 0x66 bit 4,(hl)
	(*CPU).ins_cb67, // 0x67 bit 4,a
	(*CPU).ins_cb68, // 0x68 bit 5,b
	(*CPU).ins_cb69, // 0x69 bit 5,c
	(*CPU).ins_cb6a, // 0x6a bit 5,d
	(*CPU).ins_cb6b, // 0x6b bit 5,e
	(*CPU).ins_cb6c, // 0x6c bit 5,h
	(*CPU).ins_cb6d, // 0x6d bit 5,l
	(*CPU).ins_cb6e, // 0x6e bit 5,(hl)
	(*CPU).ins_cb6f, // 0x6f bit 5,a
	(*CPU).ins_cb70, // 0x70 bit 6,b
	(*CPU).ins_cb71, // 0x71 bit 6,c
	(*CPU).ins_cb72, // 0x72 bit 6,d
	(*CPU).ins_cb73, // 0x73 bit 6,e
	(*CPU).ins_cb74, // 0x74 bit 6,h
	(*CPU).ins_cb75, // 0x75 bit 6,l
	(*CPU).ins_cb76, // 0x76 bit 6,(hl)
	(*CPU).ins_cb77, // 0x77 bit 6,a
	(*CPU).ins_cb78, // 0x78 bit 7,b
	(*CPU).ins_cb79, // 0x79 bit 7,c
	(*CPU).ins_cb7a, // 0x7a bit 7,d
	(*CPU).ins_cb7b, // 0x7b bit 7,e
	(*CPU).ins_cb7c, // 0x7c bit 7,h
	(*CPU).ins_cb7d, // 0x7d bit 7,l
	(*CPU).ins_cb7e, // 0x7e bit 7,(hl)
	(*CPU).ins_cb7f, // 0x7f bit 7,a
	(*CPU).ins_cb80, // 0x80 res 0,b
	(*CPU).ins_cb81, // 0x81 res 0,c
	(*CPU).ins_cb82, // 0x82 res 0,d
	(*CPU).ins_cb83, // 0x83 res 0,e
	(*CPU).ins_cb84, // 0x84 res 0,h
	(*CPU).ins_cb85, // 0x85 res 0,l
	(*CPU).ins_cb86, // 0x86 res 0,(hl)
	(*CPU).ins_cb87, // 0x87 res 0,a
	(*CPU).ins_cb88, // 0x88 res 1,b
	(*CPU).ins_cb89, // 0x89 res 1,c
	(*CPU).ins_cb8a, // 0x8a res 1,d
	(*CPU).ins_cb8b, // 0x8b res 1,e
	(*CPU).ins_cb8c, // 0x8c res 1,h
	(*CPU).ins_cb8d, // 0x8d res 1,l
	(*CPU).ins_cb8e, // 0x8e res 1,(hl)
	(*CPU).ins_cb8f, // 0x8f res 1,a
	(*CPU).ins_cb90, // 0x90 res 2,b
	(*CPU).ins_cb91, // 0x91 res 2,c
	(*CPU).ins_cb92, // 0x92 res 2,d
	(*CPU).ins_cb93, // 0x93 res 2,e
	(*CPU).ins_cb94, // 0x94 res 2,h
	(*CPU).ins_cb95, // 0x95 res 2,l
	(*CPU).ins_cb96, // 0x96 res 2,(hl)
	(*CPU).ins_cb97, // 0x97 res 2,a
	(*CPU).ins_cb98, // 0x98 res 3,b
	(*CPU).ins_cb99, // 0x99 res 3,c
	(*CPU).ins_cb9a, // 0x9a res 3,d
	(*CPU).ins_cb9b, // 0x9b res 3,e
	(*CPU).ins_cb9c, // 0x9c res 3,h
	(*CPU).ins_cb9d, // 0x9d res 3,l
	(*CPU).ins_cb9e, // 0x9e res 3,(hl)
	(*CPU).ins_cb9f, // 0x9f res 3,a
	(*CPU).ins_cba0, // 0xa0 res 4,b
	(*CPU).ins_cba1, // 0xa1 res 4,c
	(*CPU).ins_cba2, // 0xa2 res 4,d
	(*CPU).ins_cba3, // 0xa3 res 4,e
	(*CPU).ins_cba4, // 0xa4 res 4,h
	(*CPU).ins_cba5, // 0xa5 res 4,l
	(*CPU).ins_cba6, // 0xa6 res 4,(hl)
	(*CPU).ins_cba7, // 0xa7 res 4,a
	(*CPU).ins_cba8, // 0xa8 res 5,b
	(*CPU).ins_cba9, // 0xa9 res 5,c
	(*CPU).ins_cbaa, // 0xaa res 5,d
	(*CPU).ins_cbab, // 0xab res 5,e
	(*CPU).ins_cbac, // 0xac res 5,h
	(*CPU).ins_cbad, // 0xad res 5,l
	(*CPU).ins_cbae, // 0xae res 5,(hl)
	(*CPU).ins_cbaf, // 0xaf res 5,a
	(*CPU).ins_cbb0, // 0xb0 res 6,b
	(*CPU).ins_cbb1, // 0xb1 res 6,c
	(*CPU).ins_cbb2, // 0xb2 res 6,d
	(*CPU).ins_cbb3, // 0xb3 res 6,e
	(*CPU).ins_cbb4, // 0xb4 res 6,h
	(*CPU).ins_cbb5, // 0xb5 res 6,l
	(*CPU).ins_cbb6, // 0xb6 res 6,(hl)
	(*CPU).ins_cbb7, // 0xb7 res 6,a
	(*CPU).ins_cbb8, // 0xb8 res 7,b
	(*CPU).ins_cbb9, // 0xb9 res 7,c
	(*CPU).ins_cbba, // 0xba res 7,d
	(*CPU).ins_cbbb, // 0xbb res 7,e
	(*CPU).ins_cbbc, // 0xbc res 7,h
	(*CPU).ins_cbbd, // 0xbd res 7,l
	(*CPU).ins_cbbe, // 0xbe res 7,(hl)
	(*CPU).ins_cbbf, // 0xbf res 7,a
	(*CPU).ins_cbc0, // 0xc0 set 0,b
	(*CPU).ins_cbc1, // 0xc1 set 0,c
	(*CPU).ins_cbc2, // 0xc2 set 0,d
	(*CPU).ins_cbc3, // 0xc3 set 0,e
	(*CPU).ins_cbc4, // 0xc4 set 0,h
	(*CPU).ins_cbc5, // 0xc5 set 0,l
	(*CPU).ins_cbc6, // 0xc6 set 0,(hl)
	(*CPU).ins_cbc7, // 0xc7 set 0,a
	(*CPU).ins_cbc8, // 0xc8 set 1,b
	(*CPU).ins_cbc9, // 0xc9 set 1,c
	(*CPU).ins_cbca, // 0xca set 1,d
	(*CPU).ins_cbcb, // 0xcb set 1,e
	(*CPU).ins_cbcc, // 0xcc set 1,h
	(*CPU).ins_cbcd, // 0xcd set 1,l
	(*CPU).ins_cbce, // 0xce set 1,(hl)
	(*CPU).ins_cbcf, // 0xcf set 1,a
	(*CPU).ins_cbd0, // 0xd0 set 2,b
	(*CPU).ins_cbd1, // 0xd1 set 2,c
	(*CPU).ins_cbd2, // 0xd2 set 2,d
	(*CPU).ins_cbd3, // 0xd3 set 2,e
	(*CPU).ins_cbd4, // 0xd4 set 2,h
	(*CPU).ins_cbd5, // 0xd5 set 2,l
	(*CPU).ins_cbd6, // 0xd6 set 2,(hl)
	(*CPU).ins_cbd7, // 0xd7 set 2,a
	(*CPU).ins_cbd8, // 0xd8 set 3,b
	(*CPU).ins_cbd9, // 0xd9 set 3,c
	(*CPU).ins_cbda, // 0xda set 3,d
	(*CPU).ins_cbdb, // 0xdb set 3,e
	(*CPU).ins_cbdc, // 0xdc set 3,h
	(*CPU).ins_cbdd, // 0xdd set 3,l
	(*CPU).ins_cbde, // 0xde set 3,(hl)
	(*CPU).ins_cbdf, // 0xdf set 3,a
	(*CPU).ins_cbe0, // 0xe0 set 4,b
	(*CPU).ins_cbe1, // 0xe1 set 4,c
	(*CPU).ins_cbe2, // 0xe2 set 4,d
	(*CPU).ins_cbe3, // 0xe3 set 4,e
	(*CPU).ins_cbe4, // 0xe4 set 4,h
	(*CPU).ins_cbe5, // 0xe5 set 4,l
	(*CPU).ins_cbe6, // 0xe6 set 4,(hl)
	(*CPU).ins_cbe7, // 0xe7 set 4,a
	(*CPU).ins_cbe8, // 0xe8 set 5,b
	(*CPU).ins_cbe9, // 0xe9 set 5,c
	(*CPU).ins_cbea, // 0xea set 5,d
	(*CPU).ins_cbeb, // 0xeb set 5,e
	(*CPU).ins_cbec, // 0xec set 5,h
	(*CPU).ins_cbed, // 0xed set 5,l
	(*CPU).ins_cbee, // 0xee set 5,(hl)
	(*CPU).ins_cbef, // 0xef set 5,a
	(*CPU).ins_cbf0, // 0xf0 set 6,b
	(*CPU).ins_cbf1, // 0xf1 set 6,c
	(*CPU).ins_cbf2, // 0xf2 set 6,d
	(*CPU).ins_cbf3, // 0xf3 set 6,e
	(*CPU).ins_cbf4, // 0xf4 set 6,h
	(*CPU).ins_cbf5, // 0xf5 set 6,l
	(*CPU).ins_cbf6, // 0xf6 set 6,(hl)
	(*CPU).ins_cbf7, // 0xf7 set 6,a
	(*CPU).ins_cbf8, // 0xf8 set 7,b
	(*CPU).ins_cbf9, // 0xf9 set 7,c
	(*CPU).ins_cbfa, // 0xfa set 7,d
	(*CPU).ins_cbfb, // 0xfb set 7,e
	(*CPU).ins_cbfc, // 0xfc set 7,h
	(*CPU).ins_cbfd, // 0xfd set 7,l
	(*CPU).ins_cbfe, // 0xfe set 7,(hl)
	(*CPU).ins_cbff, // 0xff set 7,a
}
var opcodes_dd = [256]func(*CPU) int{
	(*CPU).ins_00,       // 0x00 nop
	(*CPU).ins_01,       // 0x01 ld bc,0000
	(*CPU).ins_02,       // 0x02 ld (bc),a
	(*CPU).ins_03,       // 0x03 inc bc
	(*CPU).ins_04,       // 0x04 inc b
	(*CPU).ins_05,       // 0x05 dec b
	(*CPU).ins_06,       // 0x06 ld b,00
	(*CPU).ins_07,       // 0x07 rlca
	(*CPU).ins_08,       // 0x08 ex af,af'
	(*CPU).ins_dd09,     // 0x09 add ix,bc
	(*CPU).ins_0a,       // 0x0a ld a,(bc)
	(*CPU).ins_0b,       // 0x0b dec bc
	(*CPU).ins_0c,       // 0x0c inc c
	(*CPU).ins_0d,       // 0x0d dec c
	(*CPU).ins_0e,       // 0x0e ld c,00
	(*CPU).ins_0f,       // 0x0f rrca
	(*CPU).ins_dd10,     // 0x10 djnz 0003
	(*CPU).ins_11,       // 0x11 ld de,0000
	(*CPU).ins_12,       // 0x12 ld (de),a
	(*CPU).ins_13,       // 0x13 inc de
	(*CPU).ins_14,       // 0x14 inc d
	(*CPU).ins_15,       // 0x15 dec d
	(*CPU).ins_16,       // 0x16 ld d,00
	(*CPU).ins_17,       // 0x17 rla
	(*CPU).ins_dd18,     // 0x18 jr 0003
	(*CPU).ins_dd19,     // 0x19 add ix,de
	(*CPU).ins_1a,       // 0x1a ld a,(de)
	(*CPU).ins_1b,       // 0x1b dec de
	(*CPU).ins_1c,       // 0x1c inc e
	(*CPU).ins_1d,       // 0x1d dec e
	(*CPU).ins_1e,       // 0x1e ld e,00
	(*CPU).ins_1f,       // 0x1f rra
	(*CPU).ins_dd20,     // 0x20 jr nz,0003
	(*CPU).ins_dd21,     // 0x21 ld ix,0000
	(*CPU).ins_dd22,     // 0x22 ld (0000),ix
	(*CPU).ins_dd23,     // 0x23 inc ix
	(*CPU).ins_dd24,     // 0x24 inc ixh
	(*CPU).ins_dd25,     // 0x25 dec ixh
	(*CPU).ins_dd26,     // 0x26 ld ixh,00
	(*CPU).ins_27,       // 0x27 daa
	(*CPU).ins_dd28,     // 0x28 jr z,0003
	(*CPU).ins_dd29,     // 0x29 add ix,ix
	(*CPU).ins_dd2a,     // 0x2a ld ix,(0000)
	(*CPU).ins_dd2b,     // 0x2b dec ix
	(*CPU).ins_dd2c,     // 0x2c inc ixl
	(*CPU).ins_dd2d,     // 0x2d dec ixl
	(*CPU).ins_dd2e,     // 0x2e ld ixl,00
	(*CPU).ins_2f,       // 0x2f cpl
	(*CPU).ins_dd30,     // 0x30 jr nc,0003
	(*CPU).ins_31,       // 0x31 ld sp,0000
	(*CPU).ins_32,       // 0x32 ld (0000),a
	(*CPU).ins_33,       // 0x33 inc sp
	(*CPU).ins_dd34,     // 0x34 inc (ix+00)
	(*CPU).ins_dd35,     // 0x35 dec (ix+00)
	(*CPU).ins_dd36,     // 0x36 ld (ix+00),00
	(*CPU).ins_37,       // 0x37 scf
	(*CPU).ins_dd38,     // 0x38 jr c,0003
	(*CPU).ins_dd39,     // 0x39 add ix,sp
	(*CPU).ins_3a,       // 0x3a ld a,(0000)
	(*CPU).ins_3b,       // 0x3b dec sp
	(*CPU).ins_3c,       // 0x3c inc a
	(*CPU).ins_3d,       // 0x3d dec a
	(*CPU).ins_3e,       // 0x3e ld a,00
	(*CPU).ins_3f,       // 0x3f ccf
	(*CPU).ins_40,       // 0x40 ld b,b
	(*CPU).ins_41,       // 0x41 ld b,c
	(*CPU).ins_42,       // 0x42 ld b,d
	(*CPU).ins_43,       // 0x43 ld b,e
	(*CPU).ins_dd44,     // 0x44 ld b,ixh
	(*CPU).ins_dd45,     // 0x45 ld b,ixl
	(*CPU).ins_dd46,     // 0x46 ld b,(ix+00)
	(*CPU).ins_47,       // 0x47 ld b,a
	(*CPU).ins_48,       // 0x48 ld c,b
	(*CPU).ins_49,       // 0x49 ld c,c
	(*CPU).ins_4a,       // 0x4a ld c,d
	(*CPU).ins_4b,       // 0x4b ld c,e
	(*CPU).ins_dd4c,     // 0x4c ld c,ixh
	(*CPU).ins_dd4d,     // 0x4d ld c,ixl
	(*CPU).ins_dd4e,     // 0x4e ld c,(ix+00)
	(*CPU).ins_4f,       // 0x4f ld c,a
	(*CPU).ins_50,       // 0x50 ld d,b
	(*CPU).ins_51,       // 0x51 ld d,c
	(*CPU).ins_52,       // 0x52 ld d,d
	(*CPU).ins_53,       // 0x53 ld d,e
	(*CPU).ins_dd54,     // 0x54 ld d,ixh
	(*CPU).ins_dd55,     // 0x55 ld d,ixl
	(*CPU).ins_dd56,     // 0x56 ld d,(ix+00)
	(*CPU).ins_57,       // 0x57 ld d,a
	(*CPU).ins_58,       // 0x58 ld e,b
	(*CPU).ins_59,       // 0x59 ld e,c
	(*CPU).ins_5a,       // 0x5a ld e,d
	(*CPU).ins_5b,       // 0x5b ld e,e
	(*CPU).ins_dd5c,     // 0x5c ld e,ixh
	(*CPU).ins_dd5d,     // 0x5d ld e,ixl
	(*CPU).ins_dd5e,     // 0x5e ld e,(ix+00)
	(*CPU).ins_5f,       // 0x5f ld e,a
	(*CPU).ins_dd60,     // 0x60 ld ixh,b
	(*CPU).ins_dd61,     // 0x61 ld ixh,c
	(*CPU).ins_dd62,     // 0x62 ld ixh,d
	(*CPU).ins_dd63,     // 0x63 ld ixh,e
	(*CPU).ins_dd64,     // 0x64 ld ixh,ixh
	(*CPU).ins_dd65,     // 0x65 ld ixh,ixl
	(*CPU).ins_dd66,     // 0x66 ld h,(ix+00)
	(*CPU).ins_dd67,     // 0x67 ld ixh,a
	(*CPU).ins_dd68,     // 0x68 ld ixl,b
	(*CPU).ins_dd69,     // 0x69 ld ixl,c
	(*CPU).ins_dd6a,     // 0x6a ld ixl,d
	(*CPU).ins_dd6b,     // 0x6b ld ixl,e
	(*CPU).ins_dd6c,     // 0x6c ld ixl,ixh
	(*CPU).ins_dd6d,     // 0x6d ld ixl,ixl
	(*CPU).ins_dd6e,     // 0x6e ld l,(ix+00)
	(*CPU).ins_dd6f,     // 0x6f ld ixl,a
	(*CPU).ins_dd70,     // 0x70 ld (ix+00),b
	(*CPU).ins_dd71,     // 0x71 ld (ix+00),c
	(*CPU).ins_dd72,     // 0x72 ld (ix+00),d
	(*CPU).ins_dd73,     // 0x73 ld (ix+00),e
	(*CPU).ins_dd74,     // 0x74 ld (ix+00),h
	(*CPU).ins_dd75,     // 0x75 ld (ix+00),l
	(*CPU).ins_76,       // 0x76 halt
	(*CPU).ins_dd77,     // 0x77 ld (ix+00),a
	(*CPU).ins_78,       // 0x78 ld a,b
	(*CPU).ins_79,       // 0x79 ld a,c
	(*CPU).ins_7a,       // 0x7a ld a,d
	(*CPU).ins_7b,       // 0x7b ld a,e
	(*CPU).ins_dd7c,     // 0x7c ld a,ixh
	(*CPU).ins_dd7d,     // 0x7d ld a,ixl
	(*CPU).ins_dd7e,     // 0x7e ld a,(ix+00)
	(*CPU).ins_7f,       // 0x7f ld a,a
	(*CPU).ins_80,       // 0x80 add a,b
	(*CPU).ins_81,       // 0x81 add a,c
	(*CPU).ins_82,       // 0x82 add a,d
	(*CPU).ins_83,       // 0x83 add a,e
	(*CPU).ins_dd84,     // 0x84 add a,ixh
	(*CPU).ins_dd85,     // 0x85 add a,ixl
	(*CPU).ins_dd86,     // 0x86 add a,(ix+00)
	(*CPU).ins_87,       // 0x87 add a,a
	(*CPU).ins_88,       // 0x88 adc a,b
	(*CPU).ins_89,       // 0x89 adc a,c
	(*CPU).ins_8a,       // 0x8a adc a,d
	(*CPU).ins_8b,       // 0x8b adc a,e
	(*CPU).ins_dd8c,     // 0x8c adc a,ixh
	(*CPU).ins_dd8d,     // 0x8d adc a,ixl
	(*CPU).ins_dd8e,     // 0x8e adc a,(ix+00)
	(*CPU).ins_8f,       // 0x8f adc a,a
	(*CPU).ins_90,       // 0x90 sub b
	(*CPU).ins_91,       // 0x91 sub c
	(*CPU).ins_92,       // 0x92 sub d
	(*CPU).ins_93,       // 0x93 sub e
	(*CPU).ins_dd94,     // 0x94 sub ixh
	(*CPU).ins_dd95,     // 0x95 sub ixl
	(*CPU).ins_dd96,     // 0x96 sub (ix+00)
	(*CPU).ins_97,       // 0x97 sub a
	(*CPU).ins_98,       // 0x98 sbc a,b
	(*CPU).ins_99,       // 0x99 sbc a,c
	(*CPU).ins_9a,       // 0x9a sbc a,d
	(*CPU).ins_9b,       // 0x9b sbc a,e
	(*CPU).ins_dd9c,     // 0x9c sbc a,ixh
	(*CPU).ins_dd9d,     // 0x9d sbc a,ixl
	(*CPU).ins_dd9e,     // 0x9e sbc a,(ix+00)
	(*CPU).ins_9f,       // 0x9f sbc a,a
	(*CPU).ins_a0,       // 0xa0 and b
	(*CPU).ins_a1,       // 0xa1 and c
	(*CPU).ins_a2,       // 0xa2 and d
	(*CPU).ins_a3,       // 0xa3 and e
	(*CPU).ins_dda4,     // 0xa4 and ixh
	(*CPU).ins_dda5,     // 0xa5 and ixl
	(*CPU).ins_dda6,     // 0xa6 and (ix+00)
	(*CPU).ins_a7,       // 0xa7 and a
	(*CPU).ins_a8,       // 0xa8 xor b
	(*CPU).ins_a9,       // 0xa9 xor c
	(*CPU).ins_aa,       // 0xaa xor d
	(*CPU).ins_ab,       // 0xab xor e
	(*CPU).ins_ddac,     // 0xac xor ixh
	(*CPU).ins_ddad,     // 0xad xor ixl
	(*CPU).ins_ddae,     // 0xae xor (ix+00)
	(*CPU).ins_af,       // 0xaf xor a
	(*CPU).ins_b0,       // 0xb0 or b
	(*CPU).ins_b1,       // 0xb1 or c
	(*CPU).ins_b2,       // 0xb2 or d
	(*CPU).ins_b3,       // 0xb3 or e
	(*CPU).ins_ddb4,     // 0xb4 or ixh
	(*CPU).ins_ddb5,     // 0xb5 or ixl
	(*CPU).ins_ddb6,     // 0xb6 or (ix+00)
	(*CPU).ins_b7,       // 0xb7 or a
	(*CPU).ins_b8,       // 0xb8 cp b
	(*CPU).ins_b9,       // 0xb9 cp c
	(*CPU).ins_ba,       // 0xba cp d
	(*CPU).ins_bb,       // 0xbb cp e
	(*CPU).ins_ddbc,     // 0xbc cp ixh
	(*CPU).ins_ddbd,     // 0xbd cp ixl
	(*CPU).ins_ddbe,     // 0xbe cp (ix+00)
	(*CPU).ins_bf,       // 0xbf cp a
	(*CPU).ins_c0,       // 0xc0 ret nz
	(*CPU).ins_c1,       // 0xc1 pop bc
	(*CPU).ins_c2,       // 0xc2 jp nz,0000
	(*CPU).ins_c3,       // 0xc3 jp 0000
	(*CPU).ins_c4,       // 0xc4 call nz,0000
	(*CPU).ins_c5,       // 0xc5 push bc
	(*CPU).ins_c6,       // 0xc6 add a,00
	(*CPU).ins_c7,       // 0xc7 rst 00
	(*CPU).ins_c8,       // 0xc8 ret z
	(*CPU).ins_c9,       // 0xc9 ret
	(*CPU).ins_ca,       // 0xca jp z,0000
	(*CPU).execute_ddcb, // 0xcb execute ddcb prefix
	(*CPU).ins_cc,       // 0xcc call z,0000
	(*CPU).ins_cd,       // 0xcd call 0000
	(*CPU).ins_ce,       // 0xce adc a,00
	(*CPU).ins_cf,       // 0xcf rst 08
	(*CPU).ins_d0,       // 0xd0 ret nc
	(*CPU).ins_d1,       // 0xd1 pop de
	(*CPU).ins_d2,       // 0xd2 jp nc,0000
	(*CPU).ins_d3,       // 0xd3 out (00),a
	(*CPU).ins_d4,       // 0xd4 call nc,0000
	(*CPU).ins_d5,       // 0xd5 push de
	(*CPU).ins_d6,       // 0xd6 sub 00
	(*CPU).ins_d7,       // 0xd7 rst 10
	(*CPU).ins_d8,       // 0xd8 ret c
	(*CPU).ins_d9,       // 0xd9 exx
	(*CPU).ins_da,       // 0xda jp c,0000
	(*CPU).ins_db,       // 0xdb in a,(00)
	(*CPU).ins_dc,       // 0xdc call c,0000
	(*CPU).execute_dddd, // 0xdd execute dddd prefix
	(*CPU).ins_de,       // 0xde sbc a,00
	(*CPU).ins_df,       // 0xdf rst 18
	(*CPU).ins_e0,       // 0xe0 ret po
	(*CPU).ins_dde1,     // 0xe1 pop ix
	(*CPU).ins_e2,       // 0xe2 jp po,0000
	(*CPU).ins_dde3,     // 0xe3 ex (sp),ix
	(*CPU).ins_e4,       // 0xe4 call po,0000
	(*CPU).ins_dde5,     // 0xe5 push ix
	(*CPU).ins_e6,       // 0xe6 and 00
	(*CPU).ins_e7,       // 0xe7 rst 20
	(*CPU).ins_e8,       // 0xe8 ret pe
	(*CPU).ins_dde9,     // 0xe9 jp ix
	(*CPU).ins_ea,       // 0xea jp pe,0000
	(*CPU).ins_eb,       // 0xeb ex de,hl
	(*CPU).ins_ec,       // 0xec call pe,0000
	(*CPU).ins_00,       // 0xed nop
	(*CPU).ins_ee,       // 0xee xor 00
	(*CPU).ins_ef,       // 0xef rst 28
	(*CPU).ins_f0,       // 0xf0 ret p
	(*CPU).ins_f1,       // 0xf1 pop af
	(*CPU).ins_f2,       // 0xf2 jp p,0000
	(*CPU).ins_f3,       // 0xf3 di
	(*CPU).ins_f4,       // 0xf4 call p,0000
	(*CPU).ins_f5,       // 0xf5 push af
	(*CPU).ins_f6,       // 0xf6 or 00
	(*CPU).ins_f7,       // 0xf7 rst 30
	(*CPU).ins_f8,       // 0xf8 ret m
	(*CPU).ins_ddf9,     // 0xf9 ld sp,ix
	(*CPU).ins_fa,       // 0xfa jp m,0000
	(*CPU).ins_fb,       // 0xfb ei
	(*CPU).ins_fc,       // 0xfc call m,0000
	(*CPU).execute_ddfd, // 0xfd execute ddfd prefix
	(*CPU).ins_fe,       // 0xfe cp 00
	(*CPU).ins_ff,       // 0xff rst 38
}
var opcodes_ddcb00 = [256]func(*CPU, uint8) int{
	(*CPU).ins_ddcb0000, // 0x00 rlc (ix+00),b
	(*CPU).ins_ddcb0001, // 0x01 rlc (ix+00),c
	(*CPU).ins_ddcb0002, // 0x02 rlc (ix+00),d
	(*CPU).ins_ddcb0003, // 0x03 rlc (ix+00),e
	(*CPU).ins_ddcb0004, // 0x04 rlc (ix+00),h
	(*CPU).ins_ddcb0005, // 0x05 rlc (ix+00),l
	(*CPU).ins_ddcb0006, // 0x06 rlc (ix+00)
	(*CPU).ins_ddcb0007, // 0x07 rlc (ix+00),a
	(*CPU).ins_ddcb0008, // 0x08 rrc (ix+00),b
	(*CPU).ins_ddcb0009, // 0x09 rrc (ix+00),c
	(*CPU).ins_ddcb000a, // 0x0a rrc (ix+00),d
	(*CPU).ins_ddcb000b, // 0x0b rrc (ix+00),e
	(*CPU).ins_ddcb000c, // 0x0c rrc (ix+00),h
	(*CPU).ins_ddcb000d, // 0x0d rrc (ix+00),l
	(*CPU).ins_ddcb000e, // 0x0e rrc (ix+00)
	(*CPU).ins_ddcb000f, // 0x0f rrc (ix+00),a
	(*CPU).ins_ddcb0010, // 0x10 rl (ix+00),b
	(*CPU).ins_ddcb0011, // 0x11 rl (ix+00),c
	(*CPU).ins_ddcb0012, // 0x12 rl (ix+00),d
	(*CPU).ins_ddcb0013, // 0x13 rl (ix+00),e
	(*CPU).ins_ddcb0014, // 0x14 rl (ix+00),h
	(*CPU).ins_ddcb0015, // 0x15 rl (ix+00),l
	(*CPU).ins_ddcb0016, // 0x16 rl (ix+00)
	(*CPU).ins_ddcb0017, // 0x17 rl (ix+00),a
	(*CPU).ins_ddcb0018, // 0x18 rr (ix+00),b
	(*CPU).ins_ddcb0019, // 0x19 rr (ix+00),c
	(*CPU).ins_ddcb001a, // 0x1a rr (ix+00),d
	(*CPU).ins_ddcb001b, // 0x1b rr (ix+00),e
	(*CPU).ins_ddcb001c, // 0x1c rr (ix+00),h
	(*CPU).ins_ddcb001d, // 0x1d rr (ix+00),l
	(*CPU).ins_ddcb001e, // 0x1e rr (ix+00)
	(*CPU).ins_ddcb001f, // 0x1f rr (ix+00),a
	(*CPU).ins_ddcb0020, // 0x20 sla (ix+00),b
	(*CPU).ins_ddcb0021, // 0x21 sla (ix+00),c
	(*CPU).ins_ddcb0022, // 0x22 sla (ix+00),d
	(*CPU).ins_ddcb0023, // 0x23 sla (ix+00),e
	(*CPU).ins_ddcb0024, // 0x24 sla (ix+00),h
	(*CPU).ins_ddcb0025, // 0x25 sla (ix+00),l
	(*CPU).ins_ddcb0026, // 0x26 sla (ix+00)
	(*CPU).ins_ddcb0027, // 0x27 sla (ix+00),a
	(*CPU).ins_ddcb0028, // 0x28 sra (ix+00),b
	(*CPU).ins_ddcb0029, // 0x29 sra (ix+00),c
	(*CPU).ins_ddcb002a, // 0x2a sra (ix+00),d
	(*CPU).ins_ddcb002b, // 0x2b sra (ix+00),e
	(*CPU).ins_ddcb002c, // 0x2c sra (ix+00),h
	(*CPU).ins_ddcb002d, // 0x2d sra (ix+00),l
	(*CPU).ins_ddcb002e, // 0x2e sra (ix+00)
	(*CPU).ins_ddcb002f, // 0x2f sra (ix+00),a
	(*CPU).ins_ddcb0030, // 0x30 sll (ix+00),b
	(*CPU).ins_ddcb0031, // 0x31 sll (ix+00),c
	(*CPU).ins_ddcb0032, // 0x32 sll (ix+00),d
	(*CPU).ins_ddcb0033, // 0x33 sll (ix+00),e
	(*CPU).ins_ddcb0034, // 0x34 sll (ix+00),h
	(*CPU).ins_ddcb0035, // 0x35 sll (ix+00),l
	(*CPU).ins_ddcb0036, // 0x36 sll (ix+00)
	(*CPU).ins_ddcb0037, // 0x37 sll (ix+00),a
	(*CPU).ins_ddcb0038, // 0x38 srl (ix+00),b
	(*CPU).ins_ddcb0039, // 0x39 srl (ix+00),c
	(*CPU).ins_ddcb003a, // 0x3a srl (ix+00),d
	(*CPU).ins_ddcb003b, // 0x3b srl (ix+00),e
	(*CPU).ins_ddcb003c, // 0x3c srl (ix+00),h
	(*CPU).ins_ddcb003d, // 0x3d srl (ix+00),l
	(*CPU).ins_ddcb003e, // 0x3e srl (ix+00)
	(*CPU).ins_ddcb003f, // 0x3f srl (ix+00),a
	(*CPU).ins_ddcb0040, // 0x40 bit 0,(ix+00)
	(*CPU).ins_ddcb0040, // 0x41 bit 0,(ix+00)
	(*CPU).ins_ddcb0040, // 0x42 bit 0,(ix+00)
	(*CPU).ins_ddcb0040, // 0x43 bit 0,(ix+00)
	(*CPU).ins_ddcb0040, // 0x44 bit 0,(ix+00)
	(*CPU).ins_ddcb0040, // 0x45 bit 0,(ix+00)
	(*CPU).ins_ddcb0040, // 0x46 bit 0,(ix+00)
	(*CPU).ins_ddcb0040, // 0x47 bit 0,(ix+00)
	(*CPU).ins_ddcb0048, // 0x48 bit 1,(ix+00)
	(*CPU).ins_ddcb0048, // 0x49 bit 1,(ix+00)
	(*CPU).ins_ddcb0048, // 0x4a bit 1,(ix+00)
	(*CPU).ins_ddcb0048, // 0x4b bit 1,(ix+00)
	(*CPU).ins_ddcb0048, // 0x4c bit 1,(ix+00)
	(*CPU).ins_ddcb0048, // 0x4d bit 1,(ix+00)
	(*CPU).ins_ddcb0048, // 0x4e bit 1,(ix+00)
	(*CPU).ins_ddcb0048, // 0x4f bit 1,(ix+00)
	(*CPU).ins_ddcb0050, // 0x50 bit 2,(ix+00)
	(*CPU).ins_ddcb0050, // 0x51 bit 2,(ix+00)
	(*CPU).ins_ddcb0050, // 0x52 bit 2,(ix+00)
	(*CPU).ins_ddcb0050, // 0x53 bit 2,(ix+00)
	(*CPU).ins_ddcb0050, // 0x54 bit 2,(ix+00)
	(*CPU).ins_ddcb0050, // 0x55 bit 2,(ix+00)
	(*CPU).ins_ddcb0050, // 0x56 bit 2,(ix+00)
	(*CPU).ins_ddcb0050, // 0x57 bit 2,(ix+00)
	(*CPU).ins_ddcb0058, // 0x58 bit 3,(ix+00)
	(*CPU).ins_ddcb0058, // 0x59 bit 3,(ix+00)
	(*CPU).ins_ddcb0058, // 0x5a bit 3,(ix+00)
	(*CPU).ins_ddcb0058, // 0x5b bit 3,(ix+00)
	(*CPU).ins_ddcb0058, // 0x5c bit 3,(ix+00)
	(*CPU).ins_ddcb0058, // 0x5d bit 3,(ix+00)
	(*CPU).ins_ddcb0058, // 0x5e bit 3,(ix+00)
	(*CPU).ins_ddcb0058, // 0x5f bit 3,(ix+00)
	(*CPU).ins_ddcb0060, // 0x60 bit 4,(ix+00)
	(*CPU).ins_ddcb0060, // 0x61 bit 4,(ix+00)
	(*CPU).ins_ddcb0060, // 0x62 bit 4,(ix+00)
	(*CPU).ins_ddcb0060, // 0x63 bit 4,(ix+00)
	(*CPU).ins_ddcb0060, // 0x64 bit 4,(ix+00)
	(*CPU).ins_ddcb0060, // 0x65 bit 4,(ix+00)
	(*CPU).ins_ddcb0060, // 0x66 bit 4,(ix+00)
	(*CPU).ins_ddcb0060, // 0x67 bit 4,(ix+00)
	(*CPU).ins_ddcb0068, // 0x68 bit 5,(ix+00)
	(*CPU).ins_ddcb0068, // 0x69 bit 5,(ix+00)
	(*CPU).ins_ddcb0068, // 0x6a bit 5,(ix+00)
	(*CPU).ins_ddcb0068, // 0x6b bit 5,(ix+00)
	(*CPU).ins_ddcb0068, // 0x6c bit 5,(ix+00)
	(*CPU).ins_ddcb0068, // 0x6d bit 5,(ix+00)
	(*CPU).ins_ddcb0068, // 0x6e bit 5,(ix+00)
	(*CPU).ins_ddcb0068, // 0x6f bit 5,(ix+00)
	(*CPU).ins_ddcb0070, // 0x70 bit 6,(ix+00)
	(*CPU).ins_ddcb0070, // 0x71 bit 6,(ix+00)
	(*CPU).ins_ddcb0070, // 0x72 bit 6,(ix+00)
	(*CPU).ins_ddcb0070, // 0x73 bit 6,(ix+00)
	(*CPU).ins_ddcb0070, // 0x74 bit 6,(ix+00)
	(*CPU).ins_ddcb0070, // 0x75 bit 6,(ix+00)
	(*CPU).ins_ddcb0070, // 0x76 bit 6,(ix+00)
	(*CPU).ins_ddcb0070, // 0x77 bit 6,(ix+00)
	(*CPU).ins_ddcb0078, // 0x78 bit 7,(ix+00)
	(*CPU).ins_ddcb0078, // 0x79 bit 7,(ix+00)
	(*CPU).ins_ddcb0078, // 0x7a bit 7,(ix+00)
	(*CPU).ins_ddcb0078, // 0x7b bit 7,(ix+00)
	(*CPU).ins_ddcb0078, // 0x7c bit 7,(ix+00)
	(*CPU).ins_ddcb0078, // 0x7d bit 7,(ix+00)
	(*CPU).ins_ddcb0078, // 0x7e bit 7,(ix+00)
	(*CPU).ins_ddcb0078, // 0x7f bit 7,(ix+00)
	(*CPU).ins_ddcb0080, // 0x80 res 0,(ix+00),b
	(*CPU).ins_ddcb0081, // 0x81 res 0,(ix+00),c
	(*CPU).ins_ddcb0082, // 0x82 res 0,(ix+00),d
	(*CPU).ins_ddcb0083, // 0x83 res 0,(ix+00),e
	(*CPU).ins_ddcb0084, // 0x84 res 0,(ix+00),h
	(*CPU).ins_ddcb0085, // 0x85 res 0,(ix+00),l
	(*CPU).ins_ddcb0086, // 0x86 res 0,(ix+00)
	(*CPU).ins_ddcb0087, // 0x87 res 0,(ix+00),a
	(*CPU).ins_ddcb0088, // 0x88 res 1,(ix+00),b
	(*CPU).ins_ddcb0089, // 0x89 res 1,(ix+00),c
	(*CPU).ins_ddcb008a, // 0x8a res 1,(ix+00),d
	(*CPU).ins_ddcb008b, // 0x8b res 1,(ix+00),e
	(*CPU).ins_ddcb008c, // 0x8c res 1,(ix+00),h
	(*CPU).ins_ddcb008d, // 0x8d res 1,(ix+00),l
	(*CPU).ins_ddcb008e, // 0x8e res 1,(ix+00)
	(*CPU).ins_ddcb008f, // 0x8f res 1,(ix+00),a
	(*CPU).ins_ddcb0090, // 0x90 res 2,(ix+00),b
	(*CPU).ins_ddcb0091, // 0x91 res 2,(ix+00),c
	(*CPU).ins_ddcb0092, // 0x92 res 2,(ix+00),d
	(*CPU).ins_ddcb0093, // 0x93 res 2,(ix+00),e
	(*CPU).ins_ddcb0094, // 0x94 res 2,(ix+00),h
	(*CPU).ins_ddcb0095, // 0x95 res 2,(ix+00),l
	(*CPU).ins_ddcb0096, // 0x96 res 2,(ix+00)
	(*CPU).ins_ddcb0097, // 0x97 res 2,(ix+00),a
	(*CPU).ins_ddcb0098, // 0x98 res 3,(ix+00),b
	(*CPU).ins_ddcb0099, // 0x99 res 3,(ix+00),c
	(*CPU).ins_ddcb009a, // 0x9a res 3,(ix+00),d
	(*CPU).ins_ddcb009b, // 0x9b res 3,(ix+00),e
	(*CPU).ins_ddcb009c, // 0x9c res 3,(ix+00),h
	(*CPU).ins_ddcb009d, // 0x9d res 3,(ix+00),l
	(*CPU).ins_ddcb009e, // 0x9e res 3,(ix+00)
	(*CPU).ins_ddcb009f, // 0x9f res 3,(ix+00),a
	(*CPU).ins_ddcb00a0, // 0xa0 res 4,(ix+00),b
	(*CPU).ins_ddcb00a1, // 0xa1 res 4,(ix+00),c
	(*CPU).ins_ddcb00a2, // 0xa2 res 4,(ix+00),d
	(*CPU).ins_ddcb00a3, // 0xa3 res 4,(ix+00),e
	(*CPU).ins_ddcb00a4, // 0xa4 res 4,(ix+00),h
	(*CPU).ins_ddcb00a5, // 0xa5 res 4,(ix+00),l
	(*CPU).ins_ddcb00a6, // 0xa6 res 4,(ix+00)
	(*CPU).ins_ddcb00a7, // 0xa7 res 4,(ix+00),a
	(*CPU).ins_ddcb00a8, // 0xa8 res 5,(ix+00),b
	(*CPU).ins_ddcb00a9, // 0xa9 res 5,(ix+00),c
	(*CPU).ins_ddcb00aa, // 0xaa res 5,(ix+00),d
	(*CPU).ins_ddcb00ab, // 0xab res 5,(ix+00),e
	(*CPU).ins_ddcb00ac, // 0xac res 5,(ix+00),h
	(*CPU).ins_ddcb00ad, // 0xad res 5,(ix+00),l
	(*CPU).ins_ddcb00ae, // 0xae res 5,(ix+00)
	(*CPU).ins_ddcb00af, // 0xaf res 5,(ix+00),a
	(*CPU).ins_ddcb00b0, // 0xb0 res 6,(ix+00),b
	(*CPU).ins_ddcb00b1, // 0xb1 res 6,(ix+00),c
	(*CPU).ins_ddcb00b2, // 0xb2 res 6,(ix+00),d
	(*CPU).ins_ddcb00b3, // 0xb3 res 6,(ix+00),e
	(*CPU).ins_ddcb00b4, // 0xb4 res 6,(ix+00),h
	(*CPU).ins_ddcb00b5, // 0xb5 res 6,(ix+00),l
	(*CPU).ins_ddcb00b6, // 0xb6 res 6,(ix+00)
	(*CPU).ins_ddcb00b7, // 0xb7 res 6,(ix+00),a
	(*CPU).ins_ddcb00b8, // 0xb8 res 7,(ix+00),b
	(*CPU).ins_ddcb00b9, // 0xb9 res 7,(ix+00),c
	(*CPU).ins_ddcb00ba, // 0xba res 7,(ix+00),d
	(*CPU).ins_ddcb00bb, // 0xbb res 7,(ix+00),e
	(*CPU).ins_ddcb00bc, // 0xbc res 7,(ix+00),h
	(*CPU).ins_ddcb00bd, // 0xbd res 7,(ix+00),l
	(*CPU).ins_ddcb00be, // 0xbe res 7,(ix+00)
	(*CPU).ins_ddcb00bf, // 0xbf res 7,(ix+00),a
	(*CPU).ins_ddcb00c0, // 0xc0 set 0,(ix+00),b
	(*CPU).ins_ddcb00c1, // 0xc1 set 0,(ix+00),c
	(*CPU).ins_ddcb00c2, // 0xc2 set 0,(ix+00),d
	(*CPU).ins_ddcb00c3, // 0xc3 set 0,(ix+00),e
	(*CPU).ins_ddcb00c4, // 0xc4 set 0,(ix+00),h
	(*CPU).ins_ddcb00c5, // 0xc5 set 0,(ix+00),l
	(*CPU).ins_ddcb00c6, // 0xc6 set 0,(ix+00)
	(*CPU).ins_ddcb00c7, // 0xc7 set 0,(ix+00),a
	(*CPU).ins_ddcb00c8, // 0xc8 set 1,(ix+00),b
	(*CPU).ins_ddcb00c9, // 0xc9 set 1,(ix+00),c
	(*CPU).ins_ddcb00ca, // 0xca set 1,(ix+00),d
	(*CPU).ins_ddcb00cb, // 0xcb set 1,(ix+00),e
	(*CPU).ins_ddcb00cc, // 0xcc set 1,(ix+00),h
	(*CPU).ins_ddcb00cd, // 0xcd set 1,(ix+00),l
	(*CPU).ins_ddcb00ce, // 0xce set 1,(ix+00)
	(*CPU).ins_ddcb00cf, // 0xcf set 1,(ix+00),a
	(*CPU).ins_ddcb00d0, // 0xd0 set 2,(ix+00),b
	(*CPU).ins_ddcb00d1, // 0xd1 set 2,(ix+00),c
	(*CPU).ins_ddcb00d2, // 0xd2 set 2,(ix+00),d
	(*CPU).ins_ddcb00d3, // 0xd3 set 2,(ix+00),e
	(*CPU).ins_ddcb00d4, // 0xd4 set 2,(ix+00),h
	(*CPU).ins_ddcb00d5, // 0xd5 set 2,(ix+00),l
	(*CPU).ins_ddcb00d6, // 0xd6 set 2,(ix+00)
	(*CPU).ins_ddcb00d7, // 0xd7 set 2,(ix+00),a
	(*CPU).ins_ddcb00d8, // 0xd8 set 3,(ix+00),b
	(*CPU).ins_ddcb00d9, // 0xd9 set 3,(ix+00),c
	(*CPU).ins_ddcb00da, // 0xda set 3,(ix+00),d
	(*CPU).ins_ddcb00db, // 0xdb set 3,(ix+00),e
	(*CPU).ins_ddcb00dc, // 0xdc set 3,(ix+00),h
	(*CPU).ins_ddcb00dd, // 0xdd set 3,(ix+00),l
	(*CPU).ins_ddcb00de, // 0xde set 3,(ix+00)
	(*CPU).ins_ddcb00df, // 0xdf set 3,(ix+00),a
	(*CPU).ins_ddcb00e0, // 0xe0 set 4,(ix+00),b
	(*CPU).ins_ddcb00e1, // 0xe1 set 4,(ix+00),c
	(*CPU).ins_ddcb00e2, // 0xe2 set 4,(ix+00),d
	(*CPU).ins_ddcb00e3, // 0xe3 set 4,(ix+00),e
	(*CPU).ins_ddcb00e4, // 0xe4 set 4,(ix+00),h
	(*CPU).ins_ddcb00e5, // 0xe5 set 4,(ix+00),l
	(*CPU).ins_ddcb00e6, // 0xe6 set 4,(ix+00)
	(*CPU).ins_ddcb00e7, // 0xe7 set 4,(ix+00),a
	(*CPU).ins_ddcb00e8, // 0xe8 set 5,(ix+00),b
	(*CPU).ins_ddcb00e9, // 0xe9 set 5,(ix+00),c
	(*CPU).ins_ddcb00ea, // 0xea set 5,(ix+00),d
	(*CPU).ins_ddcb00eb, // 0xeb set 5,(ix+00),e
	(*CPU).ins_ddcb00ec, // 0xec set 5,(ix+00),h
	(*CPU).ins_ddcb00ed, // 0xed set 5,(ix+00),l
	(*CPU).ins_ddcb00ee, // 0xee set 5,(ix+00)
	(*CPU).ins_ddcb00ef, // 0xef set 5,(ix+00),a
	(*CPU).ins_ddcb00f0, // 0xf0 set 6,(ix+00),b
	(*CPU).ins_ddcb00f1, // 0xf1 set 6,(ix+00),c
	(*CPU).ins_ddcb00f2, // 0xf2 set 6,(ix+00),d
	(*CPU).ins_ddcb00f3, // 0xf3 set 6,(ix+00),e
	(*CPU).ins_ddcb00f4, // 0xf4 set 6,(ix+00),h
	(*CPU).ins_ddcb00f5, // 0xf5 set 6,(ix+00),l
	(*CPU).ins_ddcb00f6, // 0xf6 set 6,(ix+00)
	(*CPU).ins_ddcb00f7, // 0xf7 set 6,(ix+00),a
	(*CPU).ins_ddcb00f8, // 0xf8 set 7,(ix+00),b
	(*CPU).ins_ddcb00f9, // 0xf9 set 7,(ix+00),c
	(*CPU).ins_ddcb00fa, // 0xfa set 7,(ix+00),d
	(*CPU).ins_ddcb00fb, // 0xfb set 7,(ix+00),e
	(*CPU).ins_ddcb00fc, // 0xfc set 7,(ix+00),h
	(*CPU).ins_ddcb00fd, // 0xfd set 7,(ix+00),l
	(*CPU).ins_ddcb00fe, // 0xfe set 7,(ix+00)
	(*CPU).ins_ddcb00ff, // 0xff set 7,(ix+00),a
}
var opcodes_ed = [256]func(*CPU) int{
	(*CPU).ins_00,   // 0x00 nop
	(*CPU).ins_00,   // 0x01 nop
	(*CPU).ins_00,   // 0x02 nop
	(*CPU).ins_00,   // 0x03 nop
	(*CPU).ins_00,   // 0x04 nop
	(*CPU).ins_00,   // 0x05 nop
	(*CPU).ins_00,   // 0x06 nop
	(*CPU).ins_00,   // 0x07 nop
	(*CPU).ins_00,   // 0x08 nop
	(*CPU).ins_00,   // 0x09 nop
	(*CPU).ins_00,   // 0x0a nop
	(*CPU).ins_00,   // 0x0b nop
	(*CPU).ins_00,   // 0x0c nop
	(*CPU).ins_00,   // 0x0d nop
	(*CPU).ins_00,   // 0x0e nop
	(*CPU).ins_00,   // 0x0f nop
	(*CPU).ins_00,   // 0x10 nop
	(*CPU).ins_00,   // 0x11 nop
	(*CPU).ins_00,   // 0x12 nop
	(*CPU).ins_00,   // 0x13 nop
	(*CPU).ins_00,   // 0x14 nop
	(*CPU).ins_00,   // 0x15 nop
	(*CPU).ins_00,   // 0x16 nop
	(*CPU).ins_00,   // 0x17 nop
	(*CPU).ins_00,   // 0x18 nop
	(*CPU).ins_00,   // 0x19 nop
	(*CPU).ins_00,   // 0x1a nop
	(*CPU).ins_00,   // 0x1b nop
	(*CPU).ins_00,   // 0x1c nop
	(*CPU).ins_00,   // 0x1d nop
	(*CPU).ins_00,   // 0x1e nop
	(*CPU).ins_00,   // 0x1f nop
	(*CPU).ins_00,   // 0x20 nop
	(*CPU).ins_00,   // 0x21 nop
	(*CPU).ins_00,   // 0x22 nop
	(*CPU).ins_00,   // 0x23 nop
	(*CPU).ins_00,   // 0x24 nop
	(*CPU).ins_00,   // 0x25 nop
	(*CPU).ins_00,   // 0x26 nop
	(*CPU).ins_00,   // 0x27 nop
	(*CPU).ins_00,   // 0x28 nop
	(*CPU).ins_00,   // 0x29 nop
	(*CPU).ins_00,   // 0x2a nop
	(*CPU).ins_00,   // 0x2b nop
	(*CPU).ins_00,   // 0x2c nop
	(*CPU).ins_00,   // 0x2d nop
	(*CPU).ins_00,   // 0x2e nop
	(*CPU).ins_00,   // 0x2f nop
	(*CPU).ins_00,   // 0x30 nop
	(*CPU).ins_00,   // 0x31 nop
	(*CPU).ins_00,   // 0x32 nop
	(*CPU).ins_00,   // 0x33 nop
	(*CPU).ins_00,   // 0x34 nop
	(*CPU).ins_00,   // 0x35 nop
	(*CPU).ins_00,   // 0x36 nop
	(*CPU).ins_00,   // 0x37 nop
	(*CPU).ins_00,   // 0x38 nop
	(*CPU).ins_00,   // 0x39 nop
	(*CPU).ins_00,   // 0x3a nop
	(*CPU).ins_00,   // 0x3b nop
	(*CPU).ins_00,   // 0x3c nop
	(*CPU).ins_00,   // 0x3d nop
	(*CPU).ins_00,   // 0x3e nop
	(*CPU).ins_00,   // 0x3f nop
	(*CPU).ins_ed40, // 0x40 in b,(c)
	(*CPU).ins_ed41, // 0x41 out (c),b
	(*CPU).ins_ed42, // 0x42 sbc hl,bc
	(*CPU).ins_ed43, // 0x43 ld (0000),bc
	(*CPU).ins_ed44, // 0x44 neg
	(*CPU).ins_ed45, // 0x45 retn
	(*CPU).ins_ed46, // 0x46 im 0
	(*CPU).ins_ed47, // 0x47 ld i,a
	(*CPU).ins_ed48, // 0x48 in c,(c)
	(*CPU).ins_ed49, // 0x49 out (c),c
	(*CPU).ins_ed4a, // 0x4a adc hl,bc
	(*CPU).ins_ed4b, // 0x4b ld bc,(0000)
	(*CPU).ins_ed44, // 0x4c neg
	(*CPU).ins_ed4d, // 0x4d reti
	(*CPU).ins_ed46, // 0x4e im 0
	(*CPU).ins_ed4f, // 0x4f ld r,a
	(*CPU).ins_ed50, // 0x50 in d,(c)
	(*CPU).ins_ed51, // 0x51 out (c),d
	(*CPU).ins_ed52, // 0x52 sbc hl,de
	(*CPU).ins_ed53, // 0x53 ld (0000),de
	(*CPU).ins_ed44, // 0x54 neg
	(*CPU).ins_ed45, // 0x55 retn
	(*CPU).ins_ed56, // 0x56 im 1
	(*CPU).ins_ed57, // 0x57 ld a,i
	(*CPU).ins_ed58, // 0x58 in e,(c)
	(*CPU).ins_ed59, // 0x59 out (c),e
	(*CPU).ins_ed5a, // 0x5a adc hl,de
	(*CPU).ins_ed5b, // 0x5b ld de,(0000)
	(*CPU).ins_ed44, // 0x5c neg
	(*CPU).ins_ed45, // 0x5d retn
	(*CPU).ins_ed5e, // 0x5e im 2
	(*CPU).ins_ed5f, // 0x5f ld a,r
	(*CPU).ins_ed60, // 0x60 in h,(c)
	(*CPU).ins_ed61, // 0x61 out (c),h
	(*CPU).ins_ed62, // 0x62 sbc hl,hl
	(*CPU).ins_22,   // 0x63 ld (0000),hl
	(*CPU).ins_ed44, // 0x64 neg
	(*CPU).ins_ed45, // 0x65 retn
	(*CPU).ins_ed46, // 0x66 im 0
	(*CPU).ins_ed67, // 0x67 rrd
	(*CPU).ins_ed68, // 0x68 in l,(c)
	(*CPU).ins_ed69, // 0x69 out (c),l
	(*CPU).ins_ed6a, // 0x6a adc hl,hl
	(*CPU).ins_2a,   // 0x6b ld hl,(0000)
	(*CPU).ins_ed44, // 0x6c neg
	(*CPU).ins_ed45, // 0x6d retn
	(*CPU).ins_ed46, // 0x6e im 0
	(*CPU).ins_ed6f, // 0x6f rld
	(*CPU).ins_ed70, // 0x70 in (c)
	(*CPU).ins_ed71, // 0x71 out (c)
	(*CPU).ins_ed72, // 0x72 sbc hl,sp
	(*CPU).ins_ed73, // 0x73 ld (0000),sp
	(*CPU).ins_ed44, // 0x74 neg
	(*CPU).ins_ed45, // 0x75 retn
	(*CPU).ins_ed56, // 0x76 im 1
	(*CPU).ins_00,   // 0x77 nop
	(*CPU).ins_ed78, // 0x78 in a,(c)
	(*CPU).ins_ed79, // 0x79 out (c),a
	(*CPU).ins_ed7a, // 0x7a adc hl,sp
	(*CPU).ins_ed7b, // 0x7b ld sp,(0000)
	(*CPU).ins_ed44, // 0x7c neg
	(*CPU).ins_ed45, // 0x7d retn
	(*CPU).ins_ed5e, // 0x7e im 2
	(*CPU).ins_00,   // 0x7f nop
	(*CPU).ins_00,   // 0x80 nop
	(*CPU).ins_00,   // 0x81 nop
	(*CPU).ins_00,   // 0x82 nop
	(*CPU).ins_00,   // 0x83 nop
	(*CPU).ins_00,   // 0x84 nop
	(*CPU).ins_00,   // 0x85 nop
	(*CPU).ins_00,   // 0x86 nop
	(*CPU).ins_00,   // 0x87 nop
	(*CPU).ins_00,   // 0x88 nop
	(*CPU).ins_00,   // 0x89 nop
	(*CPU).ins_00,   // 0x8a nop
	(*CPU).ins_00,   // 0x8b nop
	(*CPU).ins_00,   // 0x8c nop
	(*CPU).ins_00,   // 0x8d nop
	(*CPU).ins_00,   // 0x8e nop
	(*CPU).ins_00,   // 0x8f nop
	(*CPU).ins_00,   // 0x90 nop
	(*CPU).ins_00,   // 0x91 nop
	(*CPU).ins_00,   // 0x92 nop
	(*CPU).ins_00,   // 0x93 nop
	(*CPU).ins_00,   // 0x94 nop
	(*CPU).ins_00,   // 0x95 nop
	(*CPU).ins_00,   // 0x96 nop
	(*CPU).ins_00,   // 0x97 nop
	(*CPU).ins_00,   // 0x98 nop
	(*CPU).ins_00,   // 0x99 nop
	(*CPU).ins_00,   // 0x9a nop
	(*CPU).ins_00,   // 0x9b nop
	(*CPU).ins_00,   // 0x9c nop
	(*CPU).ins_00,   // 0x9d nop
	(*CPU).ins_00,   // 0x9e nop
	(*CPU).ins_00,   // 0x9f nop
	(*CPU).ins_eda0, // 0xa0 ldi
	(*CPU).ins_eda1, // 0xa1 cpi
	(*CPU).ins_eda2, // 0xa2 ini
	(*CPU).ins_eda3, // 0xa3 outi
	(*CPU).ins_00,   // 0xa4 nop
	(*CPU).ins_00,   // 0xa5 nop
	(*CPU).ins_00,   // 0xa6 nop
	(*CPU).ins_00,   // 0xa7 nop
	(*CPU).ins_eda8, // 0xa8 ldd
	(*CPU).ins_eda9, // 0xa9 cpd
	(*CPU).ins_edaa, // 0xaa ind
	(*CPU).ins_edab, // 0xab outd
	(*CPU).ins_00,   // 0xac nop
	(*CPU).ins_00,   // 0xad nop
	(*CPU).ins_00,   // 0xae nop
	(*CPU).ins_00,   // 0xaf nop
	(*CPU).ins_edb0, // 0xb0 ldir
	(*CPU).ins_edb1, // 0xb1 cpir
	(*CPU).ins_edb2, // 0xb2 inir
	(*CPU).ins_edb3, // 0xb3 otir
	(*CPU).ins_00,   // 0xb4 nop
	(*CPU).ins_00,   // 0xb5 nop
	(*CPU).ins_00,   // 0xb6 nop
	(*CPU).ins_00,   // 0xb7 nop
	(*CPU).ins_edb8, // 0xb8 lddr
	(*CPU).ins_edb9, // 0xb9 cpdr
	(*CPU).ins_edba, // 0xba indr
	(*CPU).ins_edbb, // 0xbb otdr
	(*CPU).ins_00,   // 0xbc nop
	(*CPU).ins_00,   // 0xbd nop
	(*CPU).ins_00,   // 0xbe nop
	(*CPU).ins_00,   // 0xbf nop
	(*CPU).ins_00,   // 0xc0 nop
	(*CPU).ins_00,   // 0xc1 nop
	(*CPU).ins_00,   // 0xc2 nop
	(*CPU).ins_00,   // 0xc3 nop
	(*CPU).ins_00,   // 0xc4 nop
	(*CPU).ins_00,   // 0xc5 nop
	(*CPU).ins_00,   // 0xc6 nop
	(*CPU).ins_00,   // 0xc7 nop
	(*CPU).ins_00,   // 0xc8 nop
	(*CPU).ins_00,   // 0xc9 nop
	(*CPU).ins_00,   // 0xca nop
	(*CPU).ins_00,   // 0xcb nop
	(*CPU).ins_00,   // 0xcc nop
	(*CPU).ins_00,   // 0xcd nop
	(*CPU).ins_00,   // 0xce nop
	(*CPU).ins_00,   // 0xcf nop
	(*CPU).ins_00,   // 0xd0 nop
	(*CPU).ins_00,   // 0xd1 nop
	(*CPU).ins_00,   // 0xd2 nop
	(*CPU).ins_00,   // 0xd3 nop
	(*CPU).ins_00,   // 0xd4 nop
	(*CPU).ins_00,   // 0xd5 nop
	(*CPU).ins_00,   // 0xd6 nop
	(*CPU).ins_00,   // 0xd7 nop
	(*CPU).ins_00,   // 0xd8 nop
	(*CPU).ins_00,   // 0xd9 nop
	(*CPU).ins_00,   // 0xda nop
	(*CPU).ins_00,   // 0xdb nop
	(*CPU).ins_00,   // 0xdc nop
	(*CPU).ins_00,   // 0xdd nop
	(*CPU).ins_00,   // 0xde nop
	(*CPU).ins_00,   // 0xdf nop
	(*CPU).ins_00,   // 0xe0 nop
	(*CPU).ins_00,   // 0xe1 nop
	(*CPU).ins_00,   // 0xe2 nop
	(*CPU).ins_00,   // 0xe3 nop
	(*CPU).ins_00,   // 0xe4 nop
	(*CPU).ins_00,   // 0xe5 nop
	(*CPU).ins_00,   // 0xe6 nop
	(*CPU).ins_00,   // 0xe7 nop
	(*CPU).ins_00,   // 0xe8 nop
	(*CPU).ins_00,   // 0xe9 nop
	(*CPU).ins_00,   // 0xea nop
	(*CPU).ins_00,   // 0xeb nop
	(*CPU).ins_00,   // 0xec nop
	(*CPU).ins_00,   // 0xed nop
	(*CPU).ins_00,   // 0xee nop
	(*CPU).ins_00,   // 0xef nop
	(*CPU).ins_00,   // 0xf0 nop
	(*CPU).ins_00,   // 0xf1 nop
	(*CPU).ins_00,   // 0xf2 nop
	(*CPU).ins_00,   // 0xf3 nop
	(*CPU).ins_00,   // 0xf4 nop
	(*CPU).ins_00,   // 0xf5 nop
	(*CPU).ins_00,   // 0xf6 nop
	(*CPU).ins_00,   // 0xf7 nop
	(*CPU).ins_00,   // 0xf8 nop
	(*CPU).ins_00,   // 0xf9 nop
	(*CPU).ins_00,   // 0xfa nop
	(*CPU).ins_00,   // 0xfb nop
	(*CPU).ins_00,   // 0xfc nop
	(*CPU).ins_00,   // 0xfd nop
	(*CPU).ins_00,   // 0xfe nop
	(*CPU).ins_00,   // 0xff nop
}
var opcodes_fd = [256]func(*CPU) int{
	(*CPU).ins_00,       // 0x00 nop
	(*CPU).ins_01,       // 0x01 ld bc,0000
	(*CPU).ins_02,       // 0x02 ld (bc),a
	(*CPU).ins_03,       // 0x03 inc bc
	(*CPU).ins_04,       // 0x04 inc b
	(*CPU).ins_05,       // 0x05 dec b
	(*CPU).ins_06,       // 0x06 ld b,00
	(*CPU).ins_07,       // 0x07 rlca
	(*CPU).ins_08,       // 0x08 ex af,af'
	(*CPU).ins_fd09,     // 0x09 add iy,bc
	(*CPU).ins_0a,       // 0x0a ld a,(bc)
	(*CPU).ins_0b,       // 0x0b dec bc
	(*CPU).ins_0c,       // 0x0c inc c
	(*CPU).ins_0d,       // 0x0d dec c
	(*CPU).ins_0e,       // 0x0e ld c,00
	(*CPU).ins_0f,       // 0x0f rrca
	(*CPU).ins_dd10,     // 0x10 djnz 0003
	(*CPU).ins_11,       // 0x11 ld de,0000
	(*CPU).ins_12,       // 0x12 ld (de),a
	(*CPU).ins_13,       // 0x13 inc de
	(*CPU).ins_14,       // 0x14 inc d
	(*CPU).ins_15,       // 0x15 dec d
	(*CPU).ins_16,       // 0x16 ld d,00
	(*CPU).ins_17,       // 0x17 rla
	(*CPU).ins_dd18,     // 0x18 jr 0003
	(*CPU).ins_fd19,     // 0x19 add iy,de
	(*CPU).ins_1a,       // 0x1a ld a,(de)
	(*CPU).ins_1b,       // 0x1b dec de
	(*CPU).ins_1c,       // 0x1c inc e
	(*CPU).ins_1d,       // 0x1d dec e
	(*CPU).ins_1e,       // 0x1e ld e,00
	(*CPU).ins_1f,       // 0x1f rra
	(*CPU).ins_dd20,     // 0x20 jr nz,0003
	(*CPU).ins_fd21,     // 0x21 ld iy,0000
	(*CPU).ins_fd22,     // 0x22 ld (0000),iy
	(*CPU).ins_fd23,     // 0x23 inc iy
	(*CPU).ins_fd24,     // 0x24 inc iyh
	(*CPU).ins_fd25,     // 0x25 dec iyh
	(*CPU).ins_fd26,     // 0x26 ld iyh,00
	(*CPU).ins_27,       // 0x27 daa
	(*CPU).ins_dd28,     // 0x28 jr z,0003
	(*CPU).ins_fd29,     // 0x29 add iy,iy
	(*CPU).ins_fd2a,     // 0x2a ld iy,(0000)
	(*CPU).ins_fd2b,     // 0x2b dec iy
	(*CPU).ins_fd2c,     // 0x2c inc iyl
	(*CPU).ins_fd2d,     // 0x2d dec iyl
	(*CPU).ins_fd2e,     // 0x2e ld iyl,00
	(*CPU).ins_2f,       // 0x2f cpl
	(*CPU).ins_dd30,     // 0x30 jr nc,0003
	(*CPU).ins_31,       // 0x31 ld sp,0000
	(*CPU).ins_32,       // 0x32 ld (0000),a
	(*CPU).ins_33,       // 0x33 inc sp
	(*CPU).ins_fd34,     // 0x34 inc (iy+00)
	(*CPU).ins_fd35,     // 0x35 dec (iy+00)
	(*CPU).ins_fd36,     // 0x36 ld (iy+00),00
	(*CPU).ins_37,       // 0x37 scf
	(*CPU).ins_dd38,     // 0x38 jr c,0003
	(*CPU).ins_fd39,     // 0x39 add iy,sp
	(*CPU).ins_3a,       // 0x3a ld a,(0000)
	(*CPU).ins_3b,       // 0x3b dec sp
	(*CPU).ins_3c,       // 0x3c inc a
	(*CPU).ins_3d,       // 0x3d dec a
	(*CPU).ins_3e,       // 0x3e ld a,00
	(*CPU).ins_3f,       // 0x3f ccf
	(*CPU).ins_40,       // 0x40 ld b,b
	(*CPU).ins_41,       // 0x41 ld b,c
	(*CPU).ins_42,       // 0x42 ld b,d
	(*CPU).ins_43,       // 0x43 ld b,e
	(*CPU).ins_fd44,     // 0x44 ld b,iyh
	(*CPU).ins_fd45,     // 0x45 ld b,iyl
	(*CPU).ins_fd46,     // 0x46 ld b,(iy+00)
	(*CPU).ins_47,       // 0x47 ld b,a
	(*CPU).ins_48,       // 0x48 ld c,b
	(*CPU).ins_49,       // 0x49 ld c,c
	(*CPU).ins_4a,       // 0x4a ld c,d
	(*CPU).ins_4b,       // 0x4b ld c,e
	(*CPU).ins_fd4c,     // 0x4c ld c,iyh
	(*CPU).ins_fd4d,     // 0x4d ld c,iyl
	(*CPU).ins_fd4e,     // 0x4e ld c,(iy+00)
	(*CPU).ins_4f,       // 0x4f ld c,a
	(*CPU).ins_50,       // 0x50 ld d,b
	(*CPU).ins_51,       // 0x51 ld d,c
	(*CPU).ins_52,       // 0x52 ld d,d
	(*CPU).ins_53,       // 0x53 ld d,e
	(*CPU).ins_fd54,     // 0x54 ld d,iyh
	(*CPU).ins_fd55,     // 0x55 ld d,iyl
	(*CPU).ins_fd56,     // 0x56 ld d,(iy+00)
	(*CPU).ins_57,       // 0x57 ld d,a
	(*CPU).ins_58,       // 0x58 ld e,b
	(*CPU).ins_59,       // 0x59 ld e,c
	(*CPU).ins_5a,       // 0x5a ld e,d
	(*CPU).ins_5b,       // 0x5b ld e,e
	(*CPU).ins_fd5c,     // 0x5c ld e,iyh
	(*CPU).ins_fd5d,     // 0x5d ld e,iyl
	(*CPU).ins_fd5e,     // 0x5e ld e,(iy+00)
	(*CPU).ins_5f,       // 0x5f ld e,a
	(*CPU).ins_fd60,     // 0x60 ld iyh,b
	(*CPU).ins_fd61,     // 0x61 ld iyh,c
	(*CPU).ins_fd62,     // 0x62 ld iyh,d
	(*CPU).ins_fd63,     // 0x63 ld iyh,e
	(*CPU).ins_fd64,     // 0x64 ld iyh,iyh
	(*CPU).ins_fd65,     // 0x65 ld iyh,iyl
	(*CPU).ins_fd66,     // 0x66 ld h,(iy+00)
	(*CPU).ins_fd67,     // 0x67 ld iyh,a
	(*CPU).ins_fd68,     // 0x68 ld iyl,b
	(*CPU).ins_fd69,     // 0x69 ld iyl,c
	(*CPU).ins_fd6a,     // 0x6a ld iyl,d
	(*CPU).ins_fd6b,     // 0x6b ld iyl,e
	(*CPU).ins_fd6c,     // 0x6c ld iyl,iyh
	(*CPU).ins_fd6d,     // 0x6d ld iyl,iyl
	(*CPU).ins_fd6e,     // 0x6e ld l,(iy+00)
	(*CPU).ins_fd6f,     // 0x6f ld iyl,a
	(*CPU).ins_fd70,     // 0x70 ld (iy+00),b
	(*CPU).ins_fd71,     // 0x71 ld (iy+00),c
	(*CPU).ins_fd72,     // 0x72 ld (iy+00),d
	(*CPU).ins_fd73,     // 0x73 ld (iy+00),e
	(*CPU).ins_fd74,     // 0x74 ld (iy+00),h
	(*CPU).ins_fd75,     // 0x75 ld (iy+00),l
	(*CPU).ins_76,       // 0x76 halt
	(*CPU).ins_fd77,     // 0x77 ld (iy+00),a
	(*CPU).ins_78,       // 0x78 ld a,b
	(*CPU).ins_79,       // 0x79 ld a,c
	(*CPU).ins_7a,       // 0x7a ld a,d
	(*CPU).ins_7b,       // 0x7b ld a,e
	(*CPU).ins_fd7c,     // 0x7c ld a,iyh
	(*CPU).ins_fd7d,     // 0x7d ld a,iyl
	(*CPU).ins_fd7e,     // 0x7e ld a,(iy+00)
	(*CPU).ins_7f,       // 0x7f ld a,a
	(*CPU).ins_80,       // 0x80 add a,b
	(*CPU).ins_81,       // 0x81 add a,c
	(*CPU).ins_82,       // 0x82 add a,d
	(*CPU).ins_83,       // 0x83 add a,e
	(*CPU).ins_fd84,     // 0x84 add a,iyh
	(*CPU).ins_fd85,     // 0x85 add a,iyl
	(*CPU).ins_fd86,     // 0x86 add a,(iy+00)
	(*CPU).ins_87,       // 0x87 add a,a
	(*CPU).ins_88,       // 0x88 adc a,b
	(*CPU).ins_89,       // 0x89 adc a,c
	(*CPU).ins_8a,       // 0x8a adc a,d
	(*CPU).ins_8b,       // 0x8b adc a,e
	(*CPU).ins_fd8c,     // 0x8c adc a,iyh
	(*CPU).ins_fd8d,     // 0x8d adc a,iyl
	(*CPU).ins_fd8e,     // 0x8e adc a,(iy+00)
	(*CPU).ins_8f,       // 0x8f adc a,a
	(*CPU).ins_90,       // 0x90 sub b
	(*CPU).ins_91,       // 0x91 sub c
	(*CPU).ins_92,       // 0x92 sub d
	(*CPU).ins_93,       // 0x93 sub e
	(*CPU).ins_fd94,     // 0x94 sub iyh
	(*CPU).ins_fd95,     // 0x95 sub iyl
	(*CPU).ins_fd96,     // 0x96 sub (iy+00)
	(*CPU).ins_97,       // 0x97 sub a
	(*CPU).ins_98,       // 0x98 sbc a,b
	(*CPU).ins_99,       // 0x99 sbc a,c
	(*CPU).ins_9a,       // 0x9a sbc a,d
	(*CPU).ins_9b,       // 0x9b sbc a,e
	(*CPU).ins_fd9c,     // 0x9c sbc a,iyh
	(*CPU).ins_fd9d,     // 0x9d sbc a,iyl
	(*CPU).ins_fd9e,     // 0x9e sbc a,(iy+00)
	(*CPU).ins_9f,       // 0x9f sbc a,a
	(*CPU).ins_a0,       // 0xa0 and b
	(*CPU).ins_a1,       // 0xa1 and c
	(*CPU).ins_a2,       // 0xa2 and d
	(*CPU).ins_a3,       // 0xa3 and e
	(*CPU).ins_fda4,     // 0xa4 and iyh
	(*CPU).ins_fda5,     // 0xa5 and iyl
	(*CPU).ins_fda6,     // 0xa6 and (iy+00)
	(*CPU).ins_a7,       // 0xa7 and a
	(*CPU).ins_a8,       // 0xa8 xor b
	(*CPU).ins_a9,       // 0xa9 xor c
	(*CPU).ins_aa,       // 0xaa xor d
	(*CPU).ins_ab,       // 0xab xor e
	(*CPU).ins_fdac,     // 0xac xor iyh
	(*CPU).ins_fdad,     // 0xad xor iyl
	(*CPU).ins_fdae,     // 0xae xor (iy+00)
	(*CPU).ins_af,       // 0xaf xor a
	(*CPU).ins_b0,       // 0xb0 or b
	(*CPU).ins_b1,       // 0xb1 or c
	(*CPU).ins_b2,       // 0xb2 or d
	(*CPU).ins_b3,       // 0xb3 or e
	(*CPU).ins_fdb4,     // 0xb4 or iyh
	(*CPU).ins_fdb5,     // 0xb5 or iyl
	(*CPU).ins_fdb6,     // 0xb6 or (iy+00)
	(*CPU).ins_b7,       // 0xb7 or a
	(*CPU).ins_b8,       // 0xb8 cp b
	(*CPU).ins_b9,       // 0xb9 cp c
	(*CPU).ins_ba,       // 0xba cp d
	(*CPU).ins_bb,       // 0xbb cp e
	(*CPU).ins_fdbc,     // 0xbc cp iyh
	(*CPU).ins_fdbd,     // 0xbd cp iyl
	(*CPU).ins_fdbe,     // 0xbe cp (iy+00)
	(*CPU).ins_bf,       // 0xbf cp a
	(*CPU).ins_c0,       // 0xc0 ret nz
	(*CPU).ins_c1,       // 0xc1 pop bc
	(*CPU).ins_c2,       // 0xc2 jp nz,0000
	(*CPU).ins_c3,       // 0xc3 jp 0000
	(*CPU).ins_c4,       // 0xc4 call nz,0000
	(*CPU).ins_c5,       // 0xc5 push bc
	(*CPU).ins_c6,       // 0xc6 add a,00
	(*CPU).ins_c7,       // 0xc7 rst 00
	(*CPU).ins_c8,       // 0xc8 ret z
	(*CPU).ins_c9,       // 0xc9 ret
	(*CPU).ins_ca,       // 0xca jp z,0000
	(*CPU).execute_fdcb, // 0xcb execute fdcb prefix
	(*CPU).ins_cc,       // 0xcc call z,0000
	(*CPU).ins_cd,       // 0xcd call 0000
	(*CPU).ins_ce,       // 0xce adc a,00
	(*CPU).ins_cf,       // 0xcf rst 08
	(*CPU).ins_d0,       // 0xd0 ret nc
	(*CPU).ins_d1,       // 0xd1 pop de
	(*CPU).ins_d2,       // 0xd2 jp nc,0000
	(*CPU).ins_d3,       // 0xd3 out (00),a
	(*CPU).ins_d4,       // 0xd4 call nc,0000
	(*CPU).ins_d5,       // 0xd5 push de
	(*CPU).ins_d6,       // 0xd6 sub 00
	(*CPU).ins_d7,       // 0xd7 rst 10
	(*CPU).ins_d8,       // 0xd8 ret c
	(*CPU).ins_d9,       // 0xd9 exx
	(*CPU).ins_da,       // 0xda jp c,0000
	(*CPU).ins_db,       // 0xdb in a,(00)
	(*CPU).ins_dc,       // 0xdc call c,0000
	(*CPU).execute_fddd, // 0xdd execute fddd prefix
	(*CPU).ins_de,       // 0xde sbc a,00
	(*CPU).ins_df,       // 0xdf rst 18
	(*CPU).ins_e0,       // 0xe0 ret po
	(*CPU).ins_fde1,     // 0xe1 pop iy
	(*CPU).ins_e2,       // 0xe2 jp po,0000
	(*CPU).ins_fde3,     // 0xe3 ex (sp),iy
	(*CPU).ins_e4,       // 0xe4 call po,0000
	(*CPU).ins_fde5,     // 0xe5 push iy
	(*CPU).ins_e6,       // 0xe6 and 00
	(*CPU).ins_e7,       // 0xe7 rst 20
	(*CPU).ins_e8,       // 0xe8 ret pe
	(*CPU).ins_fde9,     // 0xe9 jp iy
	(*CPU).ins_ea,       // 0xea jp pe,0000
	(*CPU).ins_eb,       // 0xeb ex de,hl
	(*CPU).ins_ec,       // 0xec call pe,0000
	(*CPU).ins_00,       // 0xed nop
	(*CPU).ins_ee,       // 0xee xor 00
	(*CPU).ins_ef,       // 0xef rst 28
	(*CPU).ins_f0,       // 0xf0 ret p
	(*CPU).ins_f1,       // 0xf1 pop af
	(*CPU).ins_f2,       // 0xf2 jp p,0000
	(*CPU).ins_f3,       // 0xf3 di
	(*CPU).ins_f4,       // 0xf4 call p,0000
	(*CPU).ins_f5,       // 0xf5 push af
	(*CPU).ins_f6,       // 0xf6 or 00
	(*CPU).ins_f7,       // 0xf7 rst 30
	(*CPU).ins_f8,       // 0xf8 ret m
	(*CPU).ins_fdf9,     // 0xf9 ld sp,iy
	(*CPU).ins_fa,       // 0xfa jp m,0000
	(*CPU).ins_fb,       // 0xfb ei
	(*CPU).ins_fc,       // 0xfc call m,0000
	(*CPU).execute_fdfd, // 0xfd execute fdfd prefix
	(*CPU).ins_fe,       // 0xfe cp 00
	(*CPU).ins_ff,       // 0xff rst 38
}
var opcodes_fdcb00 = [256]func(*CPU, uint8) int{
	(*CPU).ins_fdcb0000, // 0x00 rlc (iy+00),b
	(*CPU).ins_fdcb0001, // 0x01 rlc (iy+00),c
	(*CPU).ins_fdcb0002, // 0x02 rlc (iy+00),d
	(*CPU).ins_fdcb0003, // 0x03 rlc (iy+00),e
	(*CPU).ins_fdcb0004, // 0x04 rlc (iy+00),h
	(*CPU).ins_fdcb0005, // 0x05 rlc (iy+00),l
	(*CPU).ins_fdcb0006, // 0x06 rlc (iy+00)
	(*CPU).ins_fdcb0007, // 0x07 rlc (iy+00),a
	(*CPU).ins_fdcb0008, // 0x08 rrc (iy+00),b
	(*CPU).ins_fdcb0009, // 0x09 rrc (iy+00),c
	(*CPU).ins_fdcb000a, // 0x0a rrc (iy+00),d
	(*CPU).ins_fdcb000b, // 0x0b rrc (iy+00),e
	(*CPU).ins_fdcb000c, // 0x0c rrc (iy+00),h
	(*CPU).ins_fdcb000d, // 0x0d rrc (iy+00),l
	(*CPU).ins_fdcb000e, // 0x0e rrc (iy+00)
	(*CPU).ins_fdcb000f, // 0x0f rrc (iy+00),a
	(*CPU).ins_fdcb0010, // 0x10 rl (iy+00),b
	(*CPU).ins_fdcb0011, // 0x11 rl (iy+00),c
	(*CPU).ins_fdcb0012, // 0x12 rl (iy+00),d
	(*CPU).ins_fdcb0013, // 0x13 rl (iy+00),e
	(*CPU).ins_fdcb0014, // 0x14 rl (iy+00),h
	(*CPU).ins_fdcb0015, // 0x15 rl (iy+00),l
	(*CPU).ins_fdcb0016, // 0x16 rl (iy+00)
	(*CPU).ins_fdcb0017, // 0x17 rl (iy+00),a
	(*CPU).ins_fdcb0018, // 0x18 rr (iy+00),b
	(*CPU).ins_fdcb0019, // 0x19 rr (iy+00),c
	(*CPU).ins_fdcb001a, // 0x1a rr (iy+00),d
	(*CPU).ins_fdcb001b, // 0x1b rr (iy+00),e
	(*CPU).ins_fdcb001c, // 0x1c rr (iy+00),h
	(*CPU).ins_fdcb001d, // 0x1d rr (iy+00),l
	(*CPU).ins_fdcb001e, // 0x1e rr (iy+00)
	(*CPU).ins_fdcb001f, // 0x1f rr (iy+00),a
	(*CPU).ins_fdcb0020, // 0x20 sla (iy+00),b
	(*CPU).ins_fdcb0021, // 0x21 sla (iy+00),c
	(*CPU).ins_fdcb0022, // 0x22 sla (iy+00),d
	(*CPU).ins_fdcb0023, // 0x23 sla (iy+00),e
	(*CPU).ins_fdcb0024, // 0x24 sla (iy+00),h
	(*CPU).ins_fdcb0025, // 0x25 sla (iy+00),l
	(*CPU).ins_fdcb0026, // 0x26 sla (iy+00)
	(*CPU).ins_fdcb0027, // 0x27 sla (iy+00),a
	(*CPU).ins_fdcb0028, // 0x28 sra (iy+00),b
	(*CPU).ins_fdcb0029, // 0x29 sra (iy+00),c
	(*CPU).ins_fdcb002a, // 0x2a sra (iy+00),d
	(*CPU).ins_fdcb002b, // 0x2b sra (iy+00),e
	(*CPU).ins_fdcb002c, // 0x2c sra (iy+00),h
	(*CPU).ins_fdcb002d, // 0x2d sra (iy+00),l
	(*CPU).ins_fdcb002e, // 0x2e sra (iy+00)
	(*CPU).ins_fdcb002f, // 0x2f sra (iy+00),a
	(*CPU).ins_fdcb0030, // 0x30 sll (iy+00),b
	(*CPU).ins_fdcb0031, // 0x31 sll (iy+00),c
	(*CPU).ins_fdcb0032, // 0x32 sll (iy+00),d
	(*CPU).ins_fdcb0033, // 0x33 sll (iy+00),e
	(*CPU).ins_fdcb0034, // 0x34 sll (iy+00),h
	(*CPU).ins_fdcb0035, // 0x35 sll (iy+00),l
	(*CPU).ins_fdcb0036, // 0x36 sll (iy+00)
	(*CPU).ins_fdcb0037, // 0x37 sll (iy+00),a
	(*CPU).ins_fdcb0038, // 0x38 srl (iy+00),b
	(*CPU).ins_fdcb0039, // 0x39 srl (iy+00),c
	(*CPU).ins_fdcb003a, // 0x3a srl (iy+00),d
	(*CPU).ins_fdcb003b, // 0x3b srl (iy+00),e
	(*CPU).ins_fdcb003c, // 0x3c srl (iy+00),h
	(*CPU).ins_fdcb003d, // 0x3d srl (iy+00),l
	(*CPU).ins_fdcb003e, // 0x3e srl (iy+00)
	(*CPU).ins_fdcb003f, // 0x3f srl (iy+00),a
	(*CPU).ins_fdcb0040, // 0x40 bit 0,(iy+00)
	(*CPU).ins_fdcb0040, // 0x41 bit 0,(iy+00)
	(*CPU).ins_fdcb0040, // 0x42 bit 0,(iy+00)
	(*CPU).ins_fdcb0040, // 0x43 bit 0,(iy+00)
	(*CPU).ins_fdcb0040, // 0x44 bit 0,(iy+00)
	(*CPU).ins_fdcb0040, // 0x45 bit 0,(iy+00)
	(*CPU).ins_fdcb0040, // 0x46 bit 0,(iy+00)
	(*CPU).ins_fdcb0040, // 0x47 bit 0,(iy+00)
	(*CPU).ins_fdcb0048, // 0x48 bit 1,(iy+00)
	(*CPU).ins_fdcb0048, // 0x49 bit 1,(iy+00)
	(*CPU).ins_fdcb0048, // 0x4a bit 1,(iy+00)
	(*CPU).ins_fdcb0048, // 0x4b bit 1,(iy+00)
	(*CPU).ins_fdcb0048, // 0x4c bit 1,(iy+00)
	(*CPU).ins_fdcb0048, // 0x4d bit 1,(iy+00)
	(*CPU).ins_fdcb0048, // 0x4e bit 1,(iy+00)
	(*CPU).ins_fdcb0048, // 0x4f bit 1,(iy+00)
	(*CPU).ins_fdcb0050, // 0x50 bit 2,(iy+00)
	(*CPU).ins_fdcb0050, // 0x51 bit 2,(iy+00)
	(*CPU).ins_fdcb0050, // 0x52 bit 2,(iy+00)
	(*CPU).ins_fdcb0050, // 0x53 bit 2,(iy+00)
	(*CPU).ins_fdcb0050, // 0x54 bit 2,(iy+00)
	(*CPU).ins_fdcb0050, // 0x55 bit 2,(iy+00)
	(*CPU).ins_fdcb0050, // 0x56 bit 2,(iy+00)
	(*CPU).ins_fdcb0050, // 0x57 bit 2,(iy+00)
	(*CPU).ins_fdcb0058, // 0x58 bit 3,(iy+00)
	(*CPU).ins_fdcb0058, // 0x59 bit 3,(iy+00)
	(*CPU).ins_fdcb0058, // 0x5a bit 3,(iy+00)
	(*CPU).ins_fdcb0058, // 0x5b bit 3,(iy+00)
	(*CPU).ins_fdcb0058, // 0x5c bit 3,(iy+00)
	(*CPU).ins_fdcb0058, // 0x5d bit 3,(iy+00)
	(*CPU).ins_fdcb0058, // 0x5e bit 3,(iy+00)
	(*CPU).ins_fdcb0058, // 0x5f bit 3,(iy+00)
	(*CPU).ins_fdcb0060, // 0x60 bit 4,(iy+00)
	(*CPU).ins_fdcb0060, // 0x61 bit 4,(iy+00)
	(*CPU).ins_fdcb0060, // 0x62 bit 4,(iy+00)
	(*CPU).ins_fdcb0060, // 0x63 bit 4,(iy+00)
	(*CPU).ins_fdcb0060, // 0x64 bit 4,(iy+00)
	(*CPU).ins_fdcb0060, // 0x65 bit 4,(iy+00)
	(*CPU).ins_fdcb0060, // 0x66 bit 4,(iy+00)
	(*CPU).ins_fdcb0060, // 0x67 bit 4,(iy+00)
	(*CPU).ins_fdcb0068, // 0x68 bit 5,(iy+00)
	(*CPU).ins_fdcb0068, // 0x69 bit 5,(iy+00)
	(*CPU).ins_fdcb0068, // 0x6a bit 5,(iy+00)
	(*CPU).ins_fdcb0068, // 0x6b bit 5,(iy+00)
	(*CPU).ins_fdcb0068, // 0x6c bit 5,(iy+00)
	(*CPU).ins_fdcb0068, // 0x6d bit 5,(iy+00)
	(*CPU).ins_fdcb0068, // 0x6e bit 5,(iy+00)
	(*CPU).ins_fdcb0068, // 0x6f bit 5,(iy+00)
	(*CPU).ins_fdcb0070, // 0x70 bit 6,(iy+00)
	(*CPU).ins_fdcb0070, // 0x71 bit 6,(iy+00)
	(*CPU).ins_fdcb0070, // 0x72 bit 6,(iy+00)
	(*CPU).ins_fdcb0070, // 0x73 bit 6,(iy+00)
	(*CPU).ins_fdcb0070, // 0x74 bit 6,(iy+00)
	(*CPU).ins_fdcb0070, // 0x75 bit 6,(iy+00)
	(*CPU).ins_fdcb0070, // 0x76 bit 6,(iy+00)
	(*CPU).ins_fdcb0070, // 0x77 bit 6,(iy+00)
	(*CPU).ins_fdcb0078, // 0x78 bit 7,(iy+00)
	(*CPU).ins_fdcb0078, // 0x79 bit 7,(iy+00)
	(*CPU).ins_fdcb0078, // 0x7a bit 7,(iy+00)
	(*CPU).ins_fdcb0078, // 0x7b bit 7,(iy+00)
	(*CPU).ins_fdcb0078, // 0x7c bit 7,(iy+00)
	(*CPU).ins_fdcb0078, // 0x7d bit 7,(iy+00)
	(*CPU).ins_fdcb0078, // 0x7e bit 7,(iy+00)
	(*CPU).ins_fdcb0078, // 0x7f bit 7,(iy+00)
	(*CPU).ins_fdcb0080, // 0x80 res 0,(iy+00),b
	(*CPU).ins_fdcb0081, // 0x81 res 0,(iy+00),c
	(*CPU).ins_fdcb0082, // 0x82 res 0,(iy+00),d
	(*CPU).ins_fdcb0083, // 0x83 res 0,(iy+00),e
	(*CPU).ins_fdcb0084, // 0x84 res 0,(iy+00),h
	(*CPU).ins_fdcb0085, // 0x85 res 0,(iy+00),l
	(*CPU).ins_fdcb0086, // 0x86 res 0,(iy+00)
	(*CPU).ins_fdcb0087, // 0x87 res 0,(iy+00),a
	(*CPU).ins_fdcb0088, // 0x88 res 1,(iy+00),b
	(*CPU).ins_fdcb0089, // 0x89 res 1,(iy+00),c
	(*CPU).ins_fdcb008a, // 0x8a res 1,(iy+00),d
	(*CPU).ins_fdcb008b, // 0x8b res 1,(iy+00),e
	(*CPU).ins_fdcb008c, // 0x8c res 1,(iy+00),h
	(*CPU).ins_fdcb008d, // 0x8d res 1,(iy+00),l
	(*CPU).ins_fdcb008e, // 0x8e res 1,(iy+00)
	(*CPU).ins_fdcb008f, // 0x8f res 1,(iy+00),a
	(*CPU).ins_fdcb0090, // 0x90 res 2,(iy+00),b
	(*CPU).ins_fdcb0091, // 0x91 res 2,(iy+00),c
	(*CPU).ins_fdcb0092, // 0x92 res 2,(iy+00),d
	(*CPU).ins_fdcb0093, // 0x93 res 2,(iy+00),e
	(*CPU).ins_fdcb0094, // 0x94 res 2,(iy+00),h
	(*CPU).ins_fdcb0095, // 0x95 res 2,(iy+00),l
	(*CPU).ins_fdcb0096, // 0x96 res 2,(iy+00)
	(*CPU).ins_fdcb0097, // 0x97 res 2,(iy+00),a
	(*CPU).ins_fdcb0098, // 0x98 res 3,(iy+00),b
	(*CPU).ins_fdcb0099, // 0x99 res 3,(iy+00),c
	(*CPU).ins_fdcb009a, // 0x9a res 3,(iy+00),d
	(*CPU).ins_fdcb009b, // 0x9b res 3,(iy+00),e
	(*CPU).ins_fdcb009c, // 0x9c res 3,(iy+00),h
	(*CPU).ins_fdcb009d, // 0x9d res 3,(iy+00),l
	(*CPU).ins_fdcb009e, // 0x9e res 3,(iy+00)
	(*CPU).ins_fdcb009f, // 0x9f res 3,(iy+00),a
	(*CPU).ins_fdcb00a0, // 0xa0 res 4,(iy+00),b
	(*CPU).ins_fdcb00a1, // 0xa1 res 4,(iy+00),c
	(*CPU).ins_fdcb00a2, // 0xa2 res 4,(iy+00),d
	(*CPU).ins_fdcb00a3, // 0xa3 res 4,(iy+00),e
	(*CPU).ins_fdcb00a4, // 0xa4 res 4,(iy+00),h
	(*CPU).ins_fdcb00a5, // 0xa5 res 4,(iy+00),l
	(*CPU).ins_fdcb00a6, // 0xa6 res 4,(iy+00)
	(*CPU).ins_fdcb00a7, // 0xa7 res 4,(iy+00),a
	(*CPU).ins_fdcb00a8, // 0xa8 res 5,(iy+00),b
	(*CPU).ins_fdcb00a9, // 0xa9 res 5,(iy+00),c
	(*CPU).ins_fdcb00aa, // 0xaa res 5,(iy+00),d
	(*CPU).ins_fdcb00ab, // 0xab res 5,(iy+00),e
	(*CPU).ins_fdcb00ac, // 0xac res 5,(iy+00),h
	(*CPU).ins_fdcb00ad, // 0xad res 5,(iy+00),l
	(*CPU).ins_fdcb00ae, // 0xae res 5,(iy+00)
	(*CPU).ins_fdcb00af, // 0xaf res 5,(iy+00),a
	(*CPU).ins_fdcb00b0, // 0xb0 res 6,(iy+00),b
	(*CPU).ins_fdcb00b1, // 0xb1 res 6,(iy+00),c
	(*CPU).ins_fdcb00b2, // 0xb2 res 6,(iy+00),d
	(*CPU).ins_fdcb00b3, // 0xb3 res 6,(iy+00),e
	(*CPU).ins_fdcb00b4, // 0xb4 res 6,(iy+00),h
	(*CPU).ins_fdcb00b5, // 0xb5 res 6,(iy+00),l
	(*CPU).ins_fdcb00b6, // 0xb6 res 6,(iy+00)
	(*CPU).ins_fdcb00b7, // 0xb7 res 6,(iy+00),a
	(*CPU).ins_fdcb00b8, // 0xb8 res 7,(iy+00),b
	(*CPU).ins_fdcb00b9, // 0xb9 res 7,(iy+00),c
	(*CPU).ins_fdcb00ba, // 0xba res 7,(iy+00),d
	(*CPU).ins_fdcb00bb, // 0xbb res 7,(iy+00),e
	(*CPU).ins_fdcb00bc, // 0xbc res 7,(iy+00),h
	(*CPU).ins_fdcb00bd, // 0xbd res 7,(iy+00),l
	(*CPU).ins_fdcb00be, // 0xbe res 7,(iy+00)
	(*CPU).ins_fdcb00bf, // 0xbf res 7,(iy+00),a
	(*CPU).ins_fdcb00c0, // 0xc0 set 0,(iy+00),b
	(*CPU).ins_fdcb00c1, // 0xc1 set 0,(iy+00),c
	(*CPU).ins_fdcb00c2, // 0xc2 set 0,(iy+00),d
	(*CPU).ins_fdcb00c3, // 0xc3 set 0,(iy+00),e
	(*CPU).ins_fdcb00c4, // 0xc4 set 0,(iy+00),h
	(*CPU).ins_fdcb00c5, // 0xc5 set 0,(iy+00),l
	(*CPU).ins_fdcb00c6, // 0xc6 set 0,(iy+00)
	(*CPU).ins_fdcb00c7, // 0xc7 set 0,(iy+00),a
	(*CPU).ins_fdcb00c8, // 0xc8 set 1,(iy+00),b
	(*CPU).ins_fdcb00c9, // 0xc9 set 1,(iy+00),c
	(*CPU).ins_fdcb00ca, // 0xca set 1,(iy+00),d
	(*CPU).ins_fdcb00cb, // 0xcb set 1,(iy+00),e
	(*CPU).ins_fdcb00cc, // 0xcc set 1,(iy+00),h
	(*CPU).ins_fdcb00cd, // 0xcd set 1,(iy+00),l
	(*CPU).ins_fdcb00ce, // 0xce set 1,(iy+00)
	(*CPU).ins_fdcb00cf, // 0xcf set 1,(iy+00),a
	(*CPU).ins_fdcb00d0, // 0xd0 set 2,(iy+00),b
	(*CPU).ins_fdcb00d1, // 0xd1 set 2,(iy+00),c
	(*CPU).ins_fdcb00d2, // 0xd2 set 2,(iy+00),d
	(*CPU).ins_fdcb00d3, // 0xd3 set 2,(iy+00),e
	(*CPU).ins_fdcb00d4, // 0xd4 set 2,(iy+00),h
	(*CPU).ins_fdcb00d5, // 0xd5 set 2,(iy+00),l
	(*CPU).ins_fdcb00d6, // 0xd6 set 2,(iy+00)
	(*CPU).ins_fdcb00d7, // 0xd7 set 2,(iy+00),a
	(*CPU).ins_fdcb00d8, // 0xd8 set 3,(iy+00),b
	(*CPU).ins_fdcb00d9, // 0xd9 set 3,(iy+00),c
	(*CPU).ins_fdcb00da, // 0xda set 3,(iy+00),d
	(*CPU).ins_fdcb00db, // 0xdb set 3,(iy+00),e
	(*CPU).ins_fdcb00dc, // 0xdc set 3,(iy+00),h
	(*CPU).ins_fdcb00dd, // 0xdd set 3,(iy+00),l
	(*CPU).ins_fdcb00de, // 0xde set 3,(iy+00)
	(*CPU).ins_fdcb00df, // 0xdf set 3,(iy+00),a
	(*CPU).ins_fdcb00e0, // 0xe0 set 4,(iy+00),b
	(*CPU).ins_fdcb00e1, // 0xe1 set 4,(iy+00),c
	(*CPU).ins_fdcb00e2, // 0xe2 set 4,(iy+00),d
	(*CPU).ins_fdcb00e3, // 0xe3 set 4,(iy+00),e
	(*CPU).ins_fdcb00e4, // 0xe4 set 4,(iy+00),h
	(*CPU).ins_fdcb00e5, // 0xe5 set 4,(iy+00),l
	(*CPU).ins_fdcb00e6, // 0xe6 set 4,(iy+00)
	(*CPU).ins_fdcb00e7, // 0xe7 set 4,(iy+00),a
	(*CPU).ins_fdcb00e8, // 0xe8 set 5,(iy+00),b
	(*CPU).ins_fdcb00e9, // 0xe9 set 5,(iy+00),c
	(*CPU).ins_fdcb00ea, // 0xea set 5,(iy+00),d
	(*CPU).ins_fdcb00eb, // 0xeb set 5,(iy+00),e
	(*CPU).ins_fdcb00ec, // 0xec set 5,(iy+00),h
	(*CPU).ins_fdcb00ed, // 0xed set 5,(iy+00),l
	(*CPU).ins_fdcb00ee, // 0xee set 5,(iy+00)
	(*CPU).ins_fdcb00ef, // 0xef set 5,(iy+00),a
	(*CPU).ins_fdcb00f0, // 0xf0 set 6,(iy+00),b
	(*CPU).ins_fdcb00f1, // 0xf1 set 6,(iy+00),c
	(*CPU).ins_fdcb00f2, // 0xf2 set 6,(iy+00),d
	(*CPU).ins_fdcb00f3, // 0xf3 set 6,(iy+00),e
	(*CPU).ins_fdcb00f4, // 0xf4 set 6,(iy+00),h
	(*CPU).ins_fdcb00f5, // 0xf5 set 6,(iy+00),l
	(*CPU).ins_fdcb00f6, // 0xf6 set 6,(iy+00)
	(*CPU).ins_fdcb00f7, // 0xf7 set 6,(iy+00),a
	(*CPU).ins_fdcb00f8, // 0xf8 set 7,(iy+00),b
	(*CPU).ins_fdcb00f9, // 0xf9 set 7,(iy+00),c
	(*CPU).ins_fdcb00fa, // 0xfa set 7,(iy+00),d
	(*CPU).ins_fdcb00fb, // 0xfb set 7,(iy+00),e
	(*CPU).ins_fdcb00fc, // 0xfc set 7,(iy+00),h
	(*CPU).ins_fdcb00fd, // 0xfd set 7,(iy+00),l
	(*CPU).ins_fdcb00fe, // 0xfe set 7,(iy+00)
	(*CPU).ins_fdcb00ff, // 0xff set 7,(iy+00),a
}

// nop
func (cpu *CPU) ins_00() int {
	return 4
}

// ld bc,0000
func (cpu *CPU) ins_01() int {
	cpu.set_bc(cpu.get_nn())
	return 10
}

// ld (bc),a
func (cpu *CPU) ins_02() int {
	cpu.mem.Wr8(cpu.get_bc(), cpu.A)
	return 7
}

// inc bc
func (cpu *CPU) ins_03() int {
	cpu.set_bc(cpu.get_bc() + 1)
	return 6
}

// inc b
func (cpu *CPU) ins_04() int {
	n := cpu.B + 1
	cpu.B = n
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 4
}

// dec b
func (cpu *CPU) ins_05() int {
	n := cpu.B - 1
	cpu.B = n
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 4
}

// ld b,00
func (cpu *CPU) ins_06() int {
	cpu.B = cpu.get_n()
	return 7
}

// rlca
func (cpu *CPU) ins_07() int {
	cpu.A = ((cpu.A << 1) | (cpu.A >> 7)) & 0xff
	cpu.F = (cpu.F & (_SF | _ZF | _PF)) | (cpu.A & (_YF | _XF | _CF))
	return 4
}

// ex af,af'
func (cpu *CPU) ins_08() int {
	tmp := cpu.get_af()
	cpu.set_af(cpu.Alt_AF)
	cpu.Alt_AF = tmp
	return 4
}

// add hl,bc
func (cpu *CPU) ins_09() int {
	s := cpu.get_bc()
	d := cpu.get_hl()
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld a,(bc)
func (cpu *CPU) ins_0a() int {
	cpu.A = cpu.mem.Rd8(cpu.get_bc())
	return 7
}

// dec bc
func (cpu *CPU) ins_0b() int {
	cpu.set_bc(cpu.get_bc() - 1)
	return 6
}

// inc c
func (cpu *CPU) ins_0c() int {
	n := cpu.C + 1
	cpu.C = n
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 4
}

// dec c
func (cpu *CPU) ins_0d() int {
	n := cpu.C - 1
	cpu.C = n
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 4
}

// ld c,00
func (cpu *CPU) ins_0e() int {
	cpu.C = cpu.get_n()
	return 7
}

// rrca
func (cpu *CPU) ins_0f() int {
	cpu.F = (cpu.F & (_SF | _ZF | _PF)) | (cpu.A & _CF)
	cpu.A = ((cpu.A >> 1) | (cpu.A << 7)) & 0xff
	cpu.F |= (cpu.A & (_YF | _XF))
	return 4
}

// djnz 0002
func (cpu *CPU) ins_10() int {
	d := offset16(cpu.get_n())
	cpu.B -= 1
	if cpu.B != 0 {
		cpu.PC += d
		return 13
	}
	return 8
}

// ld de,0000
func (cpu *CPU) ins_11() int {
	cpu.set_de(cpu.get_nn())
	return 10
}

// ld (de),a
func (cpu *CPU) ins_12() int {
	cpu.mem.Wr8(cpu.get_de(), cpu.A)
	return 7
}

// inc de
func (cpu *CPU) ins_13() int {
	cpu.set_de(cpu.get_de() + 1)
	return 6
}

// inc d
func (cpu *CPU) ins_14() int {
	n := cpu.D + 1
	cpu.D = n
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 4
}

// dec d
func (cpu *CPU) ins_15() int {
	n := cpu.D - 1
	cpu.D = n
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 4
}

// ld d,00
func (cpu *CPU) ins_16() int {
	cpu.D = cpu.get_n()
	return 7
}

// rla
func (cpu *CPU) ins_17() int {
	res := (cpu.A << 1) | (cpu.F & _CF)
	var c uint8
	if (cpu.A & 0x80) != 0 {
		c = _CF
	}
	cpu.F = (cpu.F & (_SF | _ZF | _PF)) | c | (res & (_YF | _XF))
	cpu.A = res & 0xff
	return 4
}

// jr 0002
func (cpu *CPU) ins_18() int {
	cpu.PC += offset16(cpu.get_n())
	return 12
}

// add hl,de
func (cpu *CPU) ins_19() int {
	s := cpu.get_de()
	d := cpu.get_hl()
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld a,(de)
func (cpu *CPU) ins_1a() int {
	cpu.A = cpu.mem.Rd8(cpu.get_de())
	return 7
}

// dec de
func (cpu *CPU) ins_1b() int {
	cpu.set_de(cpu.get_de() - 1)
	return 6
}

// inc e
func (cpu *CPU) ins_1c() int {
	n := cpu.E + 1
	cpu.E = n
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 4
}

// dec e
func (cpu *CPU) ins_1d() int {
	n := cpu.E - 1
	cpu.E = n
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 4
}

// ld e,00
func (cpu *CPU) ins_1e() int {
	cpu.E = cpu.get_n()
	return 7
}

// rra
func (cpu *CPU) ins_1f() int {
	res := (cpu.A >> 1) | (cpu.F << 7)
	var c uint8
	if (cpu.A & 0x01) != 0 {
		c = _CF
	}
	cpu.F = (cpu.F & (_SF | _ZF | _PF)) | c | (res & (_YF | _XF))
	cpu.A = res & 0xff
	return 4
}

// jr nz,0002
func (cpu *CPU) ins_20() int {
	ofs := offset16(cpu.get_n())
	if (cpu.F & _ZF) == 0 {
		cpu.PC += ofs
		return 12
	}
	return 7
}

// ld hl,0000
func (cpu *CPU) ins_21() int {
	cpu.set_hl(cpu.get_nn())
	return 10
}

// ld (0000),hl
func (cpu *CPU) ins_22() int {
	nn := cpu.get_nn()
	cpu.mem.Wr8(nn, cpu.L)
	cpu.mem.Wr8(nn+1, cpu.H)
	return 16
}

// inc hl
func (cpu *CPU) ins_23() int {
	cpu.set_hl(cpu.get_hl() + 1)
	return 6
}

// inc h
func (cpu *CPU) ins_24() int {
	n := cpu.H + 1
	cpu.H = n
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 4
}

// dec h
func (cpu *CPU) ins_25() int {
	n := cpu.H - 1
	cpu.H = n
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 4
}

// ld h,00
func (cpu *CPU) ins_26() int {
	cpu.H = cpu.get_n()
	return 7
}

// daa
func (cpu *CPU) ins_27() int {
	cf := int2bool(int(cpu.F & _CF))
	nf := int2bool(int(cpu.F & _NF))
	hf := int2bool(int(cpu.F & _HF))
	lo := cpu.A & 0xf
	var correction uint8
	var flags uint8
	if nf {
		flags |= _NF
	}
	if hf || (lo > 9) {
		correction |= 0x06
	}
	if cf || (cpu.A > 0x99) {
		correction |= 0x60
		flags |= _CF
	}
	if nf {
		if hf && (lo < 6) {
			flags |= _HF
		}
	} else {
		if lo >= 0x0A {
			flags |= _HF
		}
	}
	if nf {
		cpu.A -= correction
	} else {
		cpu.A += correction
	}
	cpu.F = flagsSZP[cpu.A] | flags
	return 4
}

// jr z,0002
func (cpu *CPU) ins_28() int {
	ofs := offset16(cpu.get_n())
	if (cpu.F & _ZF) != 0 {
		cpu.PC += ofs
		return 12
	}
	return 7
}

// add hl,hl
func (cpu *CPU) ins_29() int {
	s := cpu.get_hl()
	d := cpu.get_hl()
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld hl,(0000)
func (cpu *CPU) ins_2a() int {
	nn := cpu.get_nn()
	cpu.H = cpu.mem.Rd8(nn + 1)
	cpu.L = cpu.mem.Rd8(nn)
	return 16
}

// dec hl
func (cpu *CPU) ins_2b() int {
	cpu.set_hl(cpu.get_hl() - 1)
	return 6
}

// inc l
func (cpu *CPU) ins_2c() int {
	n := cpu.L + 1
	cpu.L = n
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 4
}

// dec l
func (cpu *CPU) ins_2d() int {
	n := cpu.L - 1
	cpu.L = n
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 4
}

// ld l,00
func (cpu *CPU) ins_2e() int {
	cpu.L = cpu.get_n()
	return 7
}

// cpl
func (cpu *CPU) ins_2f() int {
	cpu.A ^= 0xff
	cpu.F = (cpu.F & (_SF | _ZF | _PF | _CF)) | _HF | _NF | (cpu.A & (_YF | _XF))
	return 4
}

// jr nc,0002
func (cpu *CPU) ins_30() int {
	ofs := offset16(cpu.get_n())
	if (cpu.F & _CF) == 0 {
		cpu.PC += ofs
		return 12
	}
	return 7
}

// ld sp,0000
func (cpu *CPU) ins_31() int {
	cpu.SP = cpu.get_nn()
	return 10
}

// ld (0000),a
func (cpu *CPU) ins_32() int {
	cpu.mem.Wr8(cpu.get_nn(), cpu.A)
	return 13
}

// inc sp
func (cpu *CPU) ins_33() int {
	cpu.SP += 1
	return 6
}

// inc (hl)
func (cpu *CPU) ins_34() int {
	hl := cpu.get_hl()
	n := cpu.mem.Rd8(hl) + 1
	cpu.mem.Wr8(hl, n)
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 11
}

// dec (hl)
func (cpu *CPU) ins_35() int {
	hl := cpu.get_hl()
	n := cpu.mem.Rd8(hl) - 1
	cpu.mem.Wr8(hl, n)
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 11
}

// ld (hl),00
func (cpu *CPU) ins_36() int {
	cpu.mem.Wr8(cpu.get_hl(), cpu.get_n())
	return 10
}

// scf
func (cpu *CPU) ins_37() int {
	cpu.F = (cpu.F & (_SF | _ZF | _PF)) | _CF | (cpu.A & (_YF | _XF))
	return 4
}

// jr c,0002
func (cpu *CPU) ins_38() int {
	ofs := offset16(cpu.get_n())
	if (cpu.F & _CF) != 0 {
		cpu.PC += ofs
		return 12
	}
	return 7
}

// add hl,sp
func (cpu *CPU) ins_39() int {
	s := cpu.SP
	d := cpu.get_hl()
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld a,(0000)
func (cpu *CPU) ins_3a() int {
	cpu.A = cpu.mem.Rd8(cpu.get_nn())
	return 13
}

// dec sp
func (cpu *CPU) ins_3b() int {
	cpu.SP -= 1
	return 6
}

// inc a
func (cpu *CPU) ins_3c() int {
	n := cpu.A + 1
	cpu.A = n
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 4
}

// dec a
func (cpu *CPU) ins_3d() int {
	n := cpu.A - 1
	cpu.A = n
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 4
}

// ld a,00
func (cpu *CPU) ins_3e() int {
	cpu.A = cpu.get_n()
	return 7
}

// ccf
func (cpu *CPU) ins_3f() int {
	cpu.F = ((cpu.F & (_SF | _ZF | _PF | _CF)) | ((cpu.F & _CF) << 4) | (cpu.A & (_YF | _XF))) ^ _CF
	return 4
}

// ld b,b
func (cpu *CPU) ins_40() int {
	cpu.B = cpu.B
	return 4
}

// ld b,c
func (cpu *CPU) ins_41() int {
	cpu.B = cpu.C
	return 4
}

// ld b,d
func (cpu *CPU) ins_42() int {
	cpu.B = cpu.D
	return 4
}

// ld b,e
func (cpu *CPU) ins_43() int {
	cpu.B = cpu.E
	return 4
}

// ld b,h
func (cpu *CPU) ins_44() int {
	cpu.B = cpu.H
	return 4
}

// ld b,l
func (cpu *CPU) ins_45() int {
	cpu.B = cpu.L
	return 4
}

// ld b,(hl)
func (cpu *CPU) ins_46() int {
	cpu.B = cpu.mem.Rd8(cpu.get_hl())
	return 7
}

// ld b,a
func (cpu *CPU) ins_47() int {
	cpu.B = cpu.A
	return 4
}

// ld c,b
func (cpu *CPU) ins_48() int {
	cpu.C = cpu.B
	return 4
}

// ld c,c
func (cpu *CPU) ins_49() int {
	cpu.C = cpu.C
	return 4
}

// ld c,d
func (cpu *CPU) ins_4a() int {
	cpu.C = cpu.D
	return 4
}

// ld c,e
func (cpu *CPU) ins_4b() int {
	cpu.C = cpu.E
	return 4
}

// ld c,h
func (cpu *CPU) ins_4c() int {
	cpu.C = cpu.H
	return 4
}

// ld c,l
func (cpu *CPU) ins_4d() int {
	cpu.C = cpu.L
	return 4
}

// ld c,(hl)
func (cpu *CPU) ins_4e() int {
	cpu.C = cpu.mem.Rd8(cpu.get_hl())
	return 7
}

// ld c,a
func (cpu *CPU) ins_4f() int {
	cpu.C = cpu.A
	return 4
}

// ld d,b
func (cpu *CPU) ins_50() int {
	cpu.D = cpu.B
	return 4
}

// ld d,c
func (cpu *CPU) ins_51() int {
	cpu.D = cpu.C
	return 4
}

// ld d,d
func (cpu *CPU) ins_52() int {
	cpu.D = cpu.D
	return 4
}

// ld d,e
func (cpu *CPU) ins_53() int {
	cpu.D = cpu.E
	return 4
}

// ld d,h
func (cpu *CPU) ins_54() int {
	cpu.D = cpu.H
	return 4
}

// ld d,l
func (cpu *CPU) ins_55() int {
	cpu.D = cpu.L
	return 4
}

// ld d,(hl)
func (cpu *CPU) ins_56() int {
	cpu.D = cpu.mem.Rd8(cpu.get_hl())
	return 7
}

// ld d,a
func (cpu *CPU) ins_57() int {
	cpu.D = cpu.A
	return 4
}

// ld e,b
func (cpu *CPU) ins_58() int {
	cpu.E = cpu.B
	return 4
}

// ld e,c
func (cpu *CPU) ins_59() int {
	cpu.E = cpu.C
	return 4
}

// ld e,d
func (cpu *CPU) ins_5a() int {
	cpu.E = cpu.D
	return 4
}

// ld e,e
func (cpu *CPU) ins_5b() int {
	cpu.E = cpu.E
	return 4
}

// ld e,h
func (cpu *CPU) ins_5c() int {
	cpu.E = cpu.H
	return 4
}

// ld e,l
func (cpu *CPU) ins_5d() int {
	cpu.E = cpu.L
	return 4
}

// ld e,(hl)
func (cpu *CPU) ins_5e() int {
	cpu.E = cpu.mem.Rd8(cpu.get_hl())
	return 7
}

// ld e,a
func (cpu *CPU) ins_5f() int {
	cpu.E = cpu.A
	return 4
}

// ld h,b
func (cpu *CPU) ins_60() int {
	cpu.H = cpu.B
	return 4
}

// ld h,c
func (cpu *CPU) ins_61() int {
	cpu.H = cpu.C
	return 4
}

// ld h,d
func (cpu *CPU) ins_62() int {
	cpu.H = cpu.D
	return 4
}

// ld h,e
func (cpu *CPU) ins_63() int {
	cpu.H = cpu.E
	return 4
}

// ld h,h
func (cpu *CPU) ins_64() int {
	cpu.H = cpu.H
	return 4
}

// ld h,l
func (cpu *CPU) ins_65() int {
	cpu.H = cpu.L
	return 4
}

// ld h,(hl)
func (cpu *CPU) ins_66() int {
	cpu.H = cpu.mem.Rd8(cpu.get_hl())
	return 7
}

// ld h,a
func (cpu *CPU) ins_67() int {
	cpu.H = cpu.A
	return 4
}

// ld l,b
func (cpu *CPU) ins_68() int {
	cpu.L = cpu.B
	return 4
}

// ld l,c
func (cpu *CPU) ins_69() int {
	cpu.L = cpu.C
	return 4
}

// ld l,d
func (cpu *CPU) ins_6a() int {
	cpu.L = cpu.D
	return 4
}

// ld l,e
func (cpu *CPU) ins_6b() int {
	cpu.L = cpu.E
	return 4
}

// ld l,h
func (cpu *CPU) ins_6c() int {
	cpu.L = cpu.H
	return 4
}

// ld l,l
func (cpu *CPU) ins_6d() int {
	cpu.L = cpu.L
	return 4
}

// ld l,(hl)
func (cpu *CPU) ins_6e() int {
	cpu.L = cpu.mem.Rd8(cpu.get_hl())
	return 7
}

// ld l,a
func (cpu *CPU) ins_6f() int {
	cpu.L = cpu.A
	return 4
}

// ld (hl),b
func (cpu *CPU) ins_70() int {
	cpu.mem.Wr8(cpu.get_hl(), cpu.B)
	return 7
}

// ld (hl),c
func (cpu *CPU) ins_71() int {
	cpu.mem.Wr8(cpu.get_hl(), cpu.C)
	return 7
}

// ld (hl),d
func (cpu *CPU) ins_72() int {
	cpu.mem.Wr8(cpu.get_hl(), cpu.D)
	return 7
}

// ld (hl),e
func (cpu *CPU) ins_73() int {
	cpu.mem.Wr8(cpu.get_hl(), cpu.E)
	return 7
}

// ld (hl),h
func (cpu *CPU) ins_74() int {
	cpu.mem.Wr8(cpu.get_hl(), cpu.H)
	return 7
}

// ld (hl),l
func (cpu *CPU) ins_75() int {
	cpu.mem.Wr8(cpu.get_hl(), cpu.L)
	return 7
}

// halt
func (cpu *CPU) ins_76() int {
	cpu.enter_halt()
	return 4
}

// ld (hl),a
func (cpu *CPU) ins_77() int {
	cpu.mem.Wr8(cpu.get_hl(), cpu.A)
	return 7
}

// ld a,b
func (cpu *CPU) ins_78() int {
	cpu.A = cpu.B
	return 4
}

// ld a,c
func (cpu *CPU) ins_79() int {
	cpu.A = cpu.C
	return 4
}

// ld a,d
func (cpu *CPU) ins_7a() int {
	cpu.A = cpu.D
	return 4
}

// ld a,e
func (cpu *CPU) ins_7b() int {
	cpu.A = cpu.E
	return 4
}

// ld a,h
func (cpu *CPU) ins_7c() int {
	cpu.A = cpu.H
	return 4
}

// ld a,l
func (cpu *CPU) ins_7d() int {
	cpu.A = cpu.L
	return 4
}

// ld a,(hl)
func (cpu *CPU) ins_7e() int {
	cpu.A = cpu.mem.Rd8(cpu.get_hl())
	return 7
}

// ld a,a
func (cpu *CPU) ins_7f() int {
	cpu.A = cpu.A
	return 4
}

// add a,b
func (cpu *CPU) ins_80() int {
	val := cpu.B
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// add a,c
func (cpu *CPU) ins_81() int {
	val := cpu.C
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// add a,d
func (cpu *CPU) ins_82() int {
	val := cpu.D
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// add a,e
func (cpu *CPU) ins_83() int {
	val := cpu.E
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// add a,h
func (cpu *CPU) ins_84() int {
	val := cpu.H
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// add a,l
func (cpu *CPU) ins_85() int {
	val := cpu.L
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// add a,(hl)
func (cpu *CPU) ins_86() int {
	val := cpu.mem.Rd8(cpu.get_hl())
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 7
}

// add a,a
func (cpu *CPU) ins_87() int {
	val := cpu.A
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// adc a,b
func (cpu *CPU) ins_88() int {
	val := cpu.B
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// adc a,c
func (cpu *CPU) ins_89() int {
	val := cpu.C
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// adc a,d
func (cpu *CPU) ins_8a() int {
	val := cpu.D
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// adc a,e
func (cpu *CPU) ins_8b() int {
	val := cpu.E
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// adc a,h
func (cpu *CPU) ins_8c() int {
	val := cpu.H
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// adc a,l
func (cpu *CPU) ins_8d() int {
	val := cpu.L
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// adc a,(hl)
func (cpu *CPU) ins_8e() int {
	val := cpu.mem.Rd8(cpu.get_hl())
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 7
}

// adc a,a
func (cpu *CPU) ins_8f() int {
	val := cpu.A
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sub b
func (cpu *CPU) ins_90() int {
	val := cpu.B
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sub c
func (cpu *CPU) ins_91() int {
	val := cpu.C
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sub d
func (cpu *CPU) ins_92() int {
	val := cpu.D
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sub e
func (cpu *CPU) ins_93() int {
	val := cpu.E
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sub h
func (cpu *CPU) ins_94() int {
	val := cpu.H
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sub l
func (cpu *CPU) ins_95() int {
	val := cpu.L
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sub (hl)
func (cpu *CPU) ins_96() int {
	val := cpu.mem.Rd8(cpu.get_hl())
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 7
}

// sub a
func (cpu *CPU) ins_97() int {
	val := cpu.A
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sbc a,b
func (cpu *CPU) ins_98() int {
	val := cpu.B
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sbc a,c
func (cpu *CPU) ins_99() int {
	val := cpu.C
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sbc a,d
func (cpu *CPU) ins_9a() int {
	val := cpu.D
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sbc a,e
func (cpu *CPU) ins_9b() int {
	val := cpu.E
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sbc a,h
func (cpu *CPU) ins_9c() int {
	val := cpu.H
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sbc a,l
func (cpu *CPU) ins_9d() int {
	val := cpu.L
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// sbc a,(hl)
func (cpu *CPU) ins_9e() int {
	val := cpu.mem.Rd8(cpu.get_hl())
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 7
}

// sbc a,a
func (cpu *CPU) ins_9f() int {
	val := cpu.A
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 4
}

// and b
func (cpu *CPU) ins_a0() int {
	val := cpu.B
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 4
}

// and c
func (cpu *CPU) ins_a1() int {
	val := cpu.C
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 4
}

// and d
func (cpu *CPU) ins_a2() int {
	val := cpu.D
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 4
}

// and e
func (cpu *CPU) ins_a3() int {
	val := cpu.E
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 4
}

// and h
func (cpu *CPU) ins_a4() int {
	val := cpu.H
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 4
}

// and l
func (cpu *CPU) ins_a5() int {
	val := cpu.L
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 4
}

// and (hl)
func (cpu *CPU) ins_a6() int {
	val := cpu.mem.Rd8(cpu.get_hl())
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 7
}

// and a
func (cpu *CPU) ins_a7() int {
	val := cpu.A
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 4
}

// xor b
func (cpu *CPU) ins_a8() int {
	val := cpu.B
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// xor c
func (cpu *CPU) ins_a9() int {
	val := cpu.C
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// xor d
func (cpu *CPU) ins_aa() int {
	val := cpu.D
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// xor e
func (cpu *CPU) ins_ab() int {
	val := cpu.E
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// xor h
func (cpu *CPU) ins_ac() int {
	val := cpu.H
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// xor l
func (cpu *CPU) ins_ad() int {
	val := cpu.L
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// xor (hl)
func (cpu *CPU) ins_ae() int {
	val := cpu.mem.Rd8(cpu.get_hl())
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 7
}

// xor a
func (cpu *CPU) ins_af() int {
	val := cpu.A
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// or b
func (cpu *CPU) ins_b0() int {
	val := cpu.B
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// or c
func (cpu *CPU) ins_b1() int {
	val := cpu.C
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// or d
func (cpu *CPU) ins_b2() int {
	val := cpu.D
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// or e
func (cpu *CPU) ins_b3() int {
	val := cpu.E
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// or h
func (cpu *CPU) ins_b4() int {
	val := cpu.H
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// or l
func (cpu *CPU) ins_b5() int {
	val := cpu.L
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// or (hl)
func (cpu *CPU) ins_b6() int {
	val := cpu.mem.Rd8(cpu.get_hl())
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 7
}

// or a
func (cpu *CPU) ins_b7() int {
	val := cpu.A
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 4
}

// cp b
func (cpu *CPU) ins_b8() int {
	val := cpu.B
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 4
}

// cp c
func (cpu *CPU) ins_b9() int {
	val := cpu.C
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 4
}

// cp d
func (cpu *CPU) ins_ba() int {
	val := cpu.D
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 4
}

// cp e
func (cpu *CPU) ins_bb() int {
	val := cpu.E
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 4
}

// cp h
func (cpu *CPU) ins_bc() int {
	val := cpu.H
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 4
}

// cp l
func (cpu *CPU) ins_bd() int {
	val := cpu.L
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 4
}

// cp (hl)
func (cpu *CPU) ins_be() int {
	val := cpu.mem.Rd8(cpu.get_hl())
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 7
}

// cp a
func (cpu *CPU) ins_bf() int {
	val := cpu.A
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 4
}

// ret nz
func (cpu *CPU) ins_c0() int {
	if (cpu.F & _ZF) == 0 {
		cpu.PC = cpu.pop16()
		return 11
	}
	return 5
}

// pop bc
func (cpu *CPU) ins_c1() int {
	cpu.B = cpu.mem.Rd8(cpu.SP + 1)
	cpu.C = cpu.mem.Rd8(cpu.SP)
	cpu.SP += 2
	return 10
}

// jp nz,0000
func (cpu *CPU) ins_c2() int {
	nn := cpu.get_nn()
	if (cpu.F & _ZF) == 0 {
		cpu.PC = nn
	}
	return 10
}

// jp 0000
func (cpu *CPU) ins_c3() int {
	cpu.PC = cpu.get_nn()
	return 10
}

// call nz,0000
func (cpu *CPU) ins_c4() int {
	nn := cpu.get_nn()
	if (cpu.F & _ZF) == 0 {
		cpu.push16(cpu.PC)
		cpu.PC = nn
		return 17
	}
	return 10
}

// push bc
func (cpu *CPU) ins_c5() int {
	cpu.mem.Wr8(cpu.SP-1, cpu.B)
	cpu.mem.Wr8(cpu.SP-2, cpu.C)
	cpu.SP -= 2
	return 11
}

// add a,00
func (cpu *CPU) ins_c6() int {
	val := cpu.get_n()
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 7
}

// rst 00
func (cpu *CPU) ins_c7() int {
	cpu.push16(cpu.PC)
	cpu.PC = 0x00
	return 11
}

// ret z
func (cpu *CPU) ins_c8() int {
	if (cpu.F & _ZF) != 0 {
		cpu.PC = cpu.pop16()
		return 11
	}
	return 5
}

// ret
func (cpu *CPU) ins_c9() int {
	cpu.PC = cpu.pop16()
	return 10
}

// jp z,0000
func (cpu *CPU) ins_ca() int {
	nn := cpu.get_nn()
	if (cpu.F & _ZF) != 0 {
		cpu.PC = nn
	}
	return 10
}

// call z,0000
func (cpu *CPU) ins_cc() int {
	nn := cpu.get_nn()
	if (cpu.F & _ZF) != 0 {
		cpu.push16(cpu.PC)
		cpu.PC = nn
		return 17
	}
	return 10
}

// call 0000
func (cpu *CPU) ins_cd() int {
	nn := cpu.get_nn()
	cpu.push16(cpu.PC)
	cpu.PC = nn
	return 17
}

// adc a,00
func (cpu *CPU) ins_ce() int {
	val := cpu.get_n()
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 7
}

// rst 08
func (cpu *CPU) ins_cf() int {
	cpu.push16(cpu.PC)
	cpu.PC = 0x08
	return 11
}

// ret nc
func (cpu *CPU) ins_d0() int {
	if (cpu.F & _CF) == 0 {
		cpu.PC = cpu.pop16()
		return 11
	}
	return 5
}

// pop de
func (cpu *CPU) ins_d1() int {
	cpu.D = cpu.mem.Rd8(cpu.SP + 1)
	cpu.E = cpu.mem.Rd8(cpu.SP)
	cpu.SP += 2
	return 10
}

// jp nc,0000
func (cpu *CPU) ins_d2() int {
	nn := cpu.get_nn()
	if (cpu.F & _CF) == 0 {
		cpu.PC = nn
	}
	return 10
}

// out (00),a
func (cpu *CPU) ins_d3() int {
	cpu.io.Wr8((uint16(cpu.A)<<8)|uint16(cpu.get_n()), cpu.A)
	return 7
}

// call nc,0000
func (cpu *CPU) ins_d4() int {
	nn := cpu.get_nn()
	if (cpu.F & _CF) == 0 {
		cpu.push16(cpu.PC)
		cpu.PC = nn
		return 17
	}
	return 10
}

// push de
func (cpu *CPU) ins_d5() int {
	cpu.mem.Wr8(cpu.SP-1, cpu.D)
	cpu.mem.Wr8(cpu.SP-2, cpu.E)
	cpu.SP -= 2
	return 11
}

// sub 00
func (cpu *CPU) ins_d6() int {
	val := cpu.get_n()
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 7
}

// rst 10
func (cpu *CPU) ins_d7() int {
	cpu.push16(cpu.PC)
	cpu.PC = 0x10
	return 11
}

// ret c
func (cpu *CPU) ins_d8() int {
	if (cpu.F & _CF) != 0 {
		cpu.PC = cpu.pop16()
		return 11
	}
	return 5
}

// exx
func (cpu *CPU) ins_d9() int {
	tmp := cpu.get_bc()
	cpu.set_bc(cpu.Alt_BC)
	cpu.Alt_BC = tmp
	tmp = cpu.get_de()
	cpu.set_de(cpu.Alt_DE)
	cpu.Alt_DE = tmp
	tmp = cpu.get_hl()
	cpu.set_hl(cpu.Alt_HL)
	cpu.Alt_HL = tmp
	return 4
}

// jp c,0000
func (cpu *CPU) ins_da() int {
	nn := cpu.get_nn()
	if (cpu.F & _CF) != 0 {
		cpu.PC = nn
	}
	return 10
}

// in a,(00)
func (cpu *CPU) ins_db() int {
	cpu.A = cpu.io.Rd8((uint16(cpu.A) << 8) | uint16(cpu.get_n()))
	return 7
}

// call c,0000
func (cpu *CPU) ins_dc() int {
	nn := cpu.get_nn()
	if (cpu.F & _CF) != 0 {
		cpu.push16(cpu.PC)
		cpu.PC = nn
		return 17
	}
	return 10
}

// sbc a,00
func (cpu *CPU) ins_de() int {
	val := cpu.get_n()
	result := int(cpu.A) - int(val) - int(cpu.F&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 7
}

// rst 18
func (cpu *CPU) ins_df() int {
	cpu.push16(cpu.PC)
	cpu.PC = 0x18
	return 11
}

// ret po
func (cpu *CPU) ins_e0() int {
	if (cpu.F & _PF) == 0 {
		cpu.PC = cpu.pop16()
		return 11
	}
	return 5
}

// pop hl
func (cpu *CPU) ins_e1() int {
	cpu.H = cpu.mem.Rd8(cpu.SP + 1)
	cpu.L = cpu.mem.Rd8(cpu.SP)
	cpu.SP += 2
	return 10
}

// jp po,0000
func (cpu *CPU) ins_e2() int {
	nn := cpu.get_nn()
	if (cpu.F & _PF) == 0 {
		cpu.PC = nn
	}
	return 10
}

// ex (sp),hl
func (cpu *CPU) ins_e3() int {
	tmp := cpu.mem.Rd16(cpu.SP)
	cpu.mem.Wr16(cpu.SP, cpu.get_hl())
	cpu.set_hl(tmp)
	return 19
}

// call po,0000
func (cpu *CPU) ins_e4() int {
	nn := cpu.get_nn()
	if (cpu.F & _PF) == 0 {
		cpu.push16(cpu.PC)
		cpu.PC = nn
		return 17
	}
	return 10
}

// push hl
func (cpu *CPU) ins_e5() int {
	cpu.mem.Wr8(cpu.SP-1, cpu.H)
	cpu.mem.Wr8(cpu.SP-2, cpu.L)
	cpu.SP -= 2
	return 11
}

// and 00
func (cpu *CPU) ins_e6() int {
	val := cpu.get_n()
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 7
}

// rst 20
func (cpu *CPU) ins_e7() int {
	cpu.push16(cpu.PC)
	cpu.PC = 0x20
	return 11
}

// ret pe
func (cpu *CPU) ins_e8() int {
	if (cpu.F & _PF) != 0 {
		cpu.PC = cpu.pop16()
		return 11
	}
	return 5
}

// jp hl
func (cpu *CPU) ins_e9() int {
	cpu.PC = cpu.get_hl()
	return 4
}

// jp pe,0000
func (cpu *CPU) ins_ea() int {
	nn := cpu.get_nn()
	if (cpu.F & _PF) != 0 {
		cpu.PC = nn
	}
	return 10
}

// ex de,hl
func (cpu *CPU) ins_eb() int {
	cpu.D, cpu.H = cpu.H, cpu.D
	cpu.E, cpu.L = cpu.L, cpu.E
	return 6
}

// call pe,0000
func (cpu *CPU) ins_ec() int {
	nn := cpu.get_nn()
	if (cpu.F & _PF) != 0 {
		cpu.push16(cpu.PC)
		cpu.PC = nn
		return 17
	}
	return 10
}

// xor 00
func (cpu *CPU) ins_ee() int {
	val := cpu.get_n()
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 7
}

// rst 28
func (cpu *CPU) ins_ef() int {
	cpu.push16(cpu.PC)
	cpu.PC = 0x28
	return 11
}

// ret p
func (cpu *CPU) ins_f0() int {
	if (cpu.F & _SF) == 0 {
		cpu.PC = cpu.pop16()
		return 11
	}
	return 5
}

// pop af
func (cpu *CPU) ins_f1() int {
	cpu.A = cpu.mem.Rd8(cpu.SP + 1)
	cpu.F = cpu.mem.Rd8(cpu.SP)
	cpu.SP += 2
	return 10
}

// jp p,0000
func (cpu *CPU) ins_f2() int {
	nn := cpu.get_nn()
	if (cpu.F & _SF) == 0 {
		cpu.PC = nn
	}
	return 10
}

// di
func (cpu *CPU) ins_f3() int {
	cpu.IFF1 = 0
	cpu.IFF2 = 0
	return 4
}

// call p,0000
func (cpu *CPU) ins_f4() int {
	nn := cpu.get_nn()
	if (cpu.F & _SF) == 0 {
		cpu.push16(cpu.PC)
		cpu.PC = nn
		return 17
	}
	return 10
}

// push af
func (cpu *CPU) ins_f5() int {
	cpu.mem.Wr8(cpu.SP-1, cpu.A)
	cpu.mem.Wr8(cpu.SP-2, cpu.F)
	cpu.SP -= 2
	return 11
}

// or 00
func (cpu *CPU) ins_f6() int {
	val := cpu.get_n()
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 7
}

// rst 30
func (cpu *CPU) ins_f7() int {
	cpu.push16(cpu.PC)
	cpu.PC = 0x30
	return 11
}

// ret m
func (cpu *CPU) ins_f8() int {
	if (cpu.F & _SF) != 0 {
		cpu.PC = cpu.pop16()
		return 11
	}
	return 5
}

// ld sp,hl
func (cpu *CPU) ins_f9() int {
	cpu.SP = cpu.get_hl()
	return 6
}

// jp m,0000
func (cpu *CPU) ins_fa() int {
	nn := cpu.get_nn()
	if (cpu.F & _SF) != 0 {
		cpu.PC = nn
	}
	return 10
}

// ei
func (cpu *CPU) ins_fb() int {
	cpu.IFF1 = 1
	cpu.IFF2 = 1
	return 4
}

// call m,0000
func (cpu *CPU) ins_fc() int {
	nn := cpu.get_nn()
	if (cpu.F & _SF) != 0 {
		cpu.push16(cpu.PC)
		cpu.PC = nn
		return 17
	}
	return 10
}

// cp 00
func (cpu *CPU) ins_fe() int {
	val := cpu.get_n()
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 7
}

// rst 38
func (cpu *CPU) ins_ff() int {
	cpu.push16(cpu.PC)
	cpu.PC = 0x38
	return 11
}

// rlc b
func (cpu *CPU) ins_cb00() int {
	res := cpu.B
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	return 4
}

// rlc c
func (cpu *CPU) ins_cb01() int {
	res := cpu.C
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	return 4
}

// rlc d
func (cpu *CPU) ins_cb02() int {
	res := cpu.D
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	return 4
}

// rlc e
func (cpu *CPU) ins_cb03() int {
	res := cpu.E
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	return 4
}

// rlc h
func (cpu *CPU) ins_cb04() int {
	res := cpu.H
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	return 4
}

// rlc l
func (cpu *CPU) ins_cb05() int {
	res := cpu.L
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	return 4
}

// rlc (hl)
func (cpu *CPU) ins_cb06() int {
	res := cpu.mem.Rd8(cpu.get_hl())
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.get_hl(), res)
	return 11
}

// rlc a
func (cpu *CPU) ins_cb07() int {
	res := cpu.A
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	return 4
}

// rrc b
func (cpu *CPU) ins_cb08() int {
	res := cpu.B
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	return 4
}

// rrc c
func (cpu *CPU) ins_cb09() int {
	res := cpu.C
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	return 4
}

// rrc d
func (cpu *CPU) ins_cb0a() int {
	res := cpu.D
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	return 4
}

// rrc e
func (cpu *CPU) ins_cb0b() int {
	res := cpu.E
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	return 4
}

// rrc h
func (cpu *CPU) ins_cb0c() int {
	res := cpu.H
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	return 4
}

// rrc l
func (cpu *CPU) ins_cb0d() int {
	res := cpu.L
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	return 4
}

// rrc (hl)
func (cpu *CPU) ins_cb0e() int {
	res := cpu.mem.Rd8(cpu.get_hl())
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.get_hl(), res)
	return 11
}

// rrc a
func (cpu *CPU) ins_cb0f() int {
	res := cpu.A
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	return 4
}

// rl b
func (cpu *CPU) ins_cb10() int {
	res := cpu.B
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	return 4
}

// rl c
func (cpu *CPU) ins_cb11() int {
	res := cpu.C
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	return 4
}

// rl d
func (cpu *CPU) ins_cb12() int {
	res := cpu.D
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	return 4
}

// rl e
func (cpu *CPU) ins_cb13() int {
	res := cpu.E
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	return 4
}

// rl h
func (cpu *CPU) ins_cb14() int {
	res := cpu.H
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	return 4
}

// rl l
func (cpu *CPU) ins_cb15() int {
	res := cpu.L
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	return 4
}

// rl (hl)
func (cpu *CPU) ins_cb16() int {
	res := cpu.mem.Rd8(cpu.get_hl())
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.get_hl(), res)
	return 11
}

// rl a
func (cpu *CPU) ins_cb17() int {
	res := cpu.A
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	return 4
}

// rr b
func (cpu *CPU) ins_cb18() int {
	res := cpu.B
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	return 4
}

// rr c
func (cpu *CPU) ins_cb19() int {
	res := cpu.C
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	return 4
}

// rr d
func (cpu *CPU) ins_cb1a() int {
	res := cpu.D
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	return 4
}

// rr e
func (cpu *CPU) ins_cb1b() int {
	res := cpu.E
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	return 4
}

// rr h
func (cpu *CPU) ins_cb1c() int {
	res := cpu.H
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	return 4
}

// rr l
func (cpu *CPU) ins_cb1d() int {
	res := cpu.L
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	return 4
}

// rr (hl)
func (cpu *CPU) ins_cb1e() int {
	res := cpu.mem.Rd8(cpu.get_hl())
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.get_hl(), res)
	return 11
}

// rr a
func (cpu *CPU) ins_cb1f() int {
	res := cpu.A
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	return 4
}

// sla b
func (cpu *CPU) ins_cb20() int {
	res := cpu.B
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	return 4
}

// sla c
func (cpu *CPU) ins_cb21() int {
	res := cpu.C
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	return 4
}

// sla d
func (cpu *CPU) ins_cb22() int {
	res := cpu.D
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	return 4
}

// sla e
func (cpu *CPU) ins_cb23() int {
	res := cpu.E
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	return 4
}

// sla h
func (cpu *CPU) ins_cb24() int {
	res := cpu.H
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	return 4
}

// sla l
func (cpu *CPU) ins_cb25() int {
	res := cpu.L
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	return 4
}

// sla (hl)
func (cpu *CPU) ins_cb26() int {
	res := cpu.mem.Rd8(cpu.get_hl())
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.get_hl(), res)
	return 11
}

// sla a
func (cpu *CPU) ins_cb27() int {
	res := cpu.A
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	return 4
}

// sra b
func (cpu *CPU) ins_cb28() int {
	res := cpu.B
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	return 4
}

// sra c
func (cpu *CPU) ins_cb29() int {
	res := cpu.C
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	return 4
}

// sra d
func (cpu *CPU) ins_cb2a() int {
	res := cpu.D
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	return 4
}

// sra e
func (cpu *CPU) ins_cb2b() int {
	res := cpu.E
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	return 4
}

// sra h
func (cpu *CPU) ins_cb2c() int {
	res := cpu.H
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	return 4
}

// sra l
func (cpu *CPU) ins_cb2d() int {
	res := cpu.L
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	return 4
}

// sra (hl)
func (cpu *CPU) ins_cb2e() int {
	res := cpu.mem.Rd8(cpu.get_hl())
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.get_hl(), res)
	return 11
}

// sra a
func (cpu *CPU) ins_cb2f() int {
	res := cpu.A
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	return 4
}

// sll b
func (cpu *CPU) ins_cb30() int {
	res := cpu.B
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	return 4
}

// sll c
func (cpu *CPU) ins_cb31() int {
	res := cpu.C
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	return 4
}

// sll d
func (cpu *CPU) ins_cb32() int {
	res := cpu.D
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	return 4
}

// sll e
func (cpu *CPU) ins_cb33() int {
	res := cpu.E
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	return 4
}

// sll h
func (cpu *CPU) ins_cb34() int {
	res := cpu.H
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	return 4
}

// sll l
func (cpu *CPU) ins_cb35() int {
	res := cpu.L
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	return 4
}

// sll (hl)
func (cpu *CPU) ins_cb36() int {
	res := cpu.mem.Rd8(cpu.get_hl())
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.get_hl(), res)
	return 11
}

// sll a
func (cpu *CPU) ins_cb37() int {
	res := cpu.A
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	return 4
}

// srl b
func (cpu *CPU) ins_cb38() int {
	res := cpu.B
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	return 4
}

// srl c
func (cpu *CPU) ins_cb39() int {
	res := cpu.C
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	return 4
}

// srl d
func (cpu *CPU) ins_cb3a() int {
	res := cpu.D
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	return 4
}

// srl e
func (cpu *CPU) ins_cb3b() int {
	res := cpu.E
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	return 4
}

// srl h
func (cpu *CPU) ins_cb3c() int {
	res := cpu.H
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	return 4
}

// srl l
func (cpu *CPU) ins_cb3d() int {
	res := cpu.L
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	return 4
}

// srl (hl)
func (cpu *CPU) ins_cb3e() int {
	res := cpu.mem.Rd8(cpu.get_hl())
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.get_hl(), res)
	return 11
}

// srl a
func (cpu *CPU) ins_cb3f() int {
	res := cpu.A
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	return 4
}

// bit 0,b
func (cpu *CPU) ins_cb40() int {
	bit := cpu.B & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 0,c
func (cpu *CPU) ins_cb41() int {
	bit := cpu.C & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 0,d
func (cpu *CPU) ins_cb42() int {
	bit := cpu.D & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 0,e
func (cpu *CPU) ins_cb43() int {
	bit := cpu.E & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 0,h
func (cpu *CPU) ins_cb44() int {
	bit := cpu.H & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 0,l
func (cpu *CPU) ins_cb45() int {
	bit := cpu.L & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 0,(hl)
func (cpu *CPU) ins_cb46() int {
	bit := cpu.mem.Rd8(cpu.get_hl()) & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 0,a
func (cpu *CPU) ins_cb47() int {
	bit := cpu.A & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 1,b
func (cpu *CPU) ins_cb48() int {
	bit := cpu.B & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 1,c
func (cpu *CPU) ins_cb49() int {
	bit := cpu.C & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 1,d
func (cpu *CPU) ins_cb4a() int {
	bit := cpu.D & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 1,e
func (cpu *CPU) ins_cb4b() int {
	bit := cpu.E & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 1,h
func (cpu *CPU) ins_cb4c() int {
	bit := cpu.H & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 1,l
func (cpu *CPU) ins_cb4d() int {
	bit := cpu.L & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 1,(hl)
func (cpu *CPU) ins_cb4e() int {
	bit := cpu.mem.Rd8(cpu.get_hl()) & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 1,a
func (cpu *CPU) ins_cb4f() int {
	bit := cpu.A & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 2,b
func (cpu *CPU) ins_cb50() int {
	bit := cpu.B & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 2,c
func (cpu *CPU) ins_cb51() int {
	bit := cpu.C & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 2,d
func (cpu *CPU) ins_cb52() int {
	bit := cpu.D & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 2,e
func (cpu *CPU) ins_cb53() int {
	bit := cpu.E & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 2,h
func (cpu *CPU) ins_cb54() int {
	bit := cpu.H & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 2,l
func (cpu *CPU) ins_cb55() int {
	bit := cpu.L & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 2,(hl)
func (cpu *CPU) ins_cb56() int {
	bit := cpu.mem.Rd8(cpu.get_hl()) & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 2,a
func (cpu *CPU) ins_cb57() int {
	bit := cpu.A & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 3,b
func (cpu *CPU) ins_cb58() int {
	bit := cpu.B & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 3,c
func (cpu *CPU) ins_cb59() int {
	bit := cpu.C & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 3,d
func (cpu *CPU) ins_cb5a() int {
	bit := cpu.D & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 3,e
func (cpu *CPU) ins_cb5b() int {
	bit := cpu.E & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 3,h
func (cpu *CPU) ins_cb5c() int {
	bit := cpu.H & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 3,l
func (cpu *CPU) ins_cb5d() int {
	bit := cpu.L & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 3,(hl)
func (cpu *CPU) ins_cb5e() int {
	bit := cpu.mem.Rd8(cpu.get_hl()) & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 3,a
func (cpu *CPU) ins_cb5f() int {
	bit := cpu.A & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 4,b
func (cpu *CPU) ins_cb60() int {
	bit := cpu.B & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 4,c
func (cpu *CPU) ins_cb61() int {
	bit := cpu.C & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 4,d
func (cpu *CPU) ins_cb62() int {
	bit := cpu.D & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 4,e
func (cpu *CPU) ins_cb63() int {
	bit := cpu.E & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 4,h
func (cpu *CPU) ins_cb64() int {
	bit := cpu.H & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 4,l
func (cpu *CPU) ins_cb65() int {
	bit := cpu.L & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 4,(hl)
func (cpu *CPU) ins_cb66() int {
	bit := cpu.mem.Rd8(cpu.get_hl()) & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 4,a
func (cpu *CPU) ins_cb67() int {
	bit := cpu.A & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 5,b
func (cpu *CPU) ins_cb68() int {
	bit := cpu.B & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 5,c
func (cpu *CPU) ins_cb69() int {
	bit := cpu.C & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 5,d
func (cpu *CPU) ins_cb6a() int {
	bit := cpu.D & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 5,e
func (cpu *CPU) ins_cb6b() int {
	bit := cpu.E & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 5,h
func (cpu *CPU) ins_cb6c() int {
	bit := cpu.H & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 5,l
func (cpu *CPU) ins_cb6d() int {
	bit := cpu.L & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 5,(hl)
func (cpu *CPU) ins_cb6e() int {
	bit := cpu.mem.Rd8(cpu.get_hl()) & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 5,a
func (cpu *CPU) ins_cb6f() int {
	bit := cpu.A & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 6,b
func (cpu *CPU) ins_cb70() int {
	bit := cpu.B & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 6,c
func (cpu *CPU) ins_cb71() int {
	bit := cpu.C & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 6,d
func (cpu *CPU) ins_cb72() int {
	bit := cpu.D & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 6,e
func (cpu *CPU) ins_cb73() int {
	bit := cpu.E & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 6,h
func (cpu *CPU) ins_cb74() int {
	bit := cpu.H & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 6,l
func (cpu *CPU) ins_cb75() int {
	bit := cpu.L & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 6,(hl)
func (cpu *CPU) ins_cb76() int {
	bit := cpu.mem.Rd8(cpu.get_hl()) & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 6,a
func (cpu *CPU) ins_cb77() int {
	bit := cpu.A & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 7,b
func (cpu *CPU) ins_cb78() int {
	bit := cpu.B & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 7,c
func (cpu *CPU) ins_cb79() int {
	bit := cpu.C & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 7,d
func (cpu *CPU) ins_cb7a() int {
	bit := cpu.D & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 7,e
func (cpu *CPU) ins_cb7b() int {
	bit := cpu.E & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 7,h
func (cpu *CPU) ins_cb7c() int {
	bit := cpu.H & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 7,l
func (cpu *CPU) ins_cb7d() int {
	bit := cpu.L & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// bit 7,(hl)
func (cpu *CPU) ins_cb7e() int {
	bit := cpu.mem.Rd8(cpu.get_hl()) & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 7,a
func (cpu *CPU) ins_cb7f() int {
	bit := cpu.A & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 4
}

// res 0,b
func (cpu *CPU) ins_cb80() int {
	cpu.B = cpu.B &^ (1 << 0)
	return 4
}

// res 0,c
func (cpu *CPU) ins_cb81() int {
	cpu.C = cpu.C &^ (1 << 0)
	return 4
}

// res 0,d
func (cpu *CPU) ins_cb82() int {
	cpu.D = cpu.D &^ (1 << 0)
	return 4
}

// res 0,e
func (cpu *CPU) ins_cb83() int {
	cpu.E = cpu.E &^ (1 << 0)
	return 4
}

// res 0,h
func (cpu *CPU) ins_cb84() int {
	cpu.H = cpu.H &^ (1 << 0)
	return 4
}

// res 0,l
func (cpu *CPU) ins_cb85() int {
	cpu.L = cpu.L &^ (1 << 0)
	return 4
}

// res 0,(hl)
func (cpu *CPU) ins_cb86() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 0,a
func (cpu *CPU) ins_cb87() int {
	cpu.A = cpu.A &^ (1 << 0)
	return 4
}

// res 1,b
func (cpu *CPU) ins_cb88() int {
	cpu.B = cpu.B &^ (1 << 1)
	return 4
}

// res 1,c
func (cpu *CPU) ins_cb89() int {
	cpu.C = cpu.C &^ (1 << 1)
	return 4
}

// res 1,d
func (cpu *CPU) ins_cb8a() int {
	cpu.D = cpu.D &^ (1 << 1)
	return 4
}

// res 1,e
func (cpu *CPU) ins_cb8b() int {
	cpu.E = cpu.E &^ (1 << 1)
	return 4
}

// res 1,h
func (cpu *CPU) ins_cb8c() int {
	cpu.H = cpu.H &^ (1 << 1)
	return 4
}

// res 1,l
func (cpu *CPU) ins_cb8d() int {
	cpu.L = cpu.L &^ (1 << 1)
	return 4
}

// res 1,(hl)
func (cpu *CPU) ins_cb8e() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 1,a
func (cpu *CPU) ins_cb8f() int {
	cpu.A = cpu.A &^ (1 << 1)
	return 4
}

// res 2,b
func (cpu *CPU) ins_cb90() int {
	cpu.B = cpu.B &^ (1 << 2)
	return 4
}

// res 2,c
func (cpu *CPU) ins_cb91() int {
	cpu.C = cpu.C &^ (1 << 2)
	return 4
}

// res 2,d
func (cpu *CPU) ins_cb92() int {
	cpu.D = cpu.D &^ (1 << 2)
	return 4
}

// res 2,e
func (cpu *CPU) ins_cb93() int {
	cpu.E = cpu.E &^ (1 << 2)
	return 4
}

// res 2,h
func (cpu *CPU) ins_cb94() int {
	cpu.H = cpu.H &^ (1 << 2)
	return 4
}

// res 2,l
func (cpu *CPU) ins_cb95() int {
	cpu.L = cpu.L &^ (1 << 2)
	return 4
}

// res 2,(hl)
func (cpu *CPU) ins_cb96() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 2,a
func (cpu *CPU) ins_cb97() int {
	cpu.A = cpu.A &^ (1 << 2)
	return 4
}

// res 3,b
func (cpu *CPU) ins_cb98() int {
	cpu.B = cpu.B &^ (1 << 3)
	return 4
}

// res 3,c
func (cpu *CPU) ins_cb99() int {
	cpu.C = cpu.C &^ (1 << 3)
	return 4
}

// res 3,d
func (cpu *CPU) ins_cb9a() int {
	cpu.D = cpu.D &^ (1 << 3)
	return 4
}

// res 3,e
func (cpu *CPU) ins_cb9b() int {
	cpu.E = cpu.E &^ (1 << 3)
	return 4
}

// res 3,h
func (cpu *CPU) ins_cb9c() int {
	cpu.H = cpu.H &^ (1 << 3)
	return 4
}

// res 3,l
func (cpu *CPU) ins_cb9d() int {
	cpu.L = cpu.L &^ (1 << 3)
	return 4
}

// res 3,(hl)
func (cpu *CPU) ins_cb9e() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 3,a
func (cpu *CPU) ins_cb9f() int {
	cpu.A = cpu.A &^ (1 << 3)
	return 4
}

// res 4,b
func (cpu *CPU) ins_cba0() int {
	cpu.B = cpu.B &^ (1 << 4)
	return 4
}

// res 4,c
func (cpu *CPU) ins_cba1() int {
	cpu.C = cpu.C &^ (1 << 4)
	return 4
}

// res 4,d
func (cpu *CPU) ins_cba2() int {
	cpu.D = cpu.D &^ (1 << 4)
	return 4
}

// res 4,e
func (cpu *CPU) ins_cba3() int {
	cpu.E = cpu.E &^ (1 << 4)
	return 4
}

// res 4,h
func (cpu *CPU) ins_cba4() int {
	cpu.H = cpu.H &^ (1 << 4)
	return 4
}

// res 4,l
func (cpu *CPU) ins_cba5() int {
	cpu.L = cpu.L &^ (1 << 4)
	return 4
}

// res 4,(hl)
func (cpu *CPU) ins_cba6() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 4,a
func (cpu *CPU) ins_cba7() int {
	cpu.A = cpu.A &^ (1 << 4)
	return 4
}

// res 5,b
func (cpu *CPU) ins_cba8() int {
	cpu.B = cpu.B &^ (1 << 5)
	return 4
}

// res 5,c
func (cpu *CPU) ins_cba9() int {
	cpu.C = cpu.C &^ (1 << 5)
	return 4
}

// res 5,d
func (cpu *CPU) ins_cbaa() int {
	cpu.D = cpu.D &^ (1 << 5)
	return 4
}

// res 5,e
func (cpu *CPU) ins_cbab() int {
	cpu.E = cpu.E &^ (1 << 5)
	return 4
}

// res 5,h
func (cpu *CPU) ins_cbac() int {
	cpu.H = cpu.H &^ (1 << 5)
	return 4
}

// res 5,l
func (cpu *CPU) ins_cbad() int {
	cpu.L = cpu.L &^ (1 << 5)
	return 4
}

// res 5,(hl)
func (cpu *CPU) ins_cbae() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 5,a
func (cpu *CPU) ins_cbaf() int {
	cpu.A = cpu.A &^ (1 << 5)
	return 4
}

// res 6,b
func (cpu *CPU) ins_cbb0() int {
	cpu.B = cpu.B &^ (1 << 6)
	return 4
}

// res 6,c
func (cpu *CPU) ins_cbb1() int {
	cpu.C = cpu.C &^ (1 << 6)
	return 4
}

// res 6,d
func (cpu *CPU) ins_cbb2() int {
	cpu.D = cpu.D &^ (1 << 6)
	return 4
}

// res 6,e
func (cpu *CPU) ins_cbb3() int {
	cpu.E = cpu.E &^ (1 << 6)
	return 4
}

// res 6,h
func (cpu *CPU) ins_cbb4() int {
	cpu.H = cpu.H &^ (1 << 6)
	return 4
}

// res 6,l
func (cpu *CPU) ins_cbb5() int {
	cpu.L = cpu.L &^ (1 << 6)
	return 4
}

// res 6,(hl)
func (cpu *CPU) ins_cbb6() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 6,a
func (cpu *CPU) ins_cbb7() int {
	cpu.A = cpu.A &^ (1 << 6)
	return 4
}

// res 7,b
func (cpu *CPU) ins_cbb8() int {
	cpu.B = cpu.B &^ (1 << 7)
	return 4
}

// res 7,c
func (cpu *CPU) ins_cbb9() int {
	cpu.C = cpu.C &^ (1 << 7)
	return 4
}

// res 7,d
func (cpu *CPU) ins_cbba() int {
	cpu.D = cpu.D &^ (1 << 7)
	return 4
}

// res 7,e
func (cpu *CPU) ins_cbbb() int {
	cpu.E = cpu.E &^ (1 << 7)
	return 4
}

// res 7,h
func (cpu *CPU) ins_cbbc() int {
	cpu.H = cpu.H &^ (1 << 7)
	return 4
}

// res 7,l
func (cpu *CPU) ins_cbbd() int {
	cpu.L = cpu.L &^ (1 << 7)
	return 4
}

// res 7,(hl)
func (cpu *CPU) ins_cbbe() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 7,a
func (cpu *CPU) ins_cbbf() int {
	cpu.A = cpu.A &^ (1 << 7)
	return 4
}

// set 0,b
func (cpu *CPU) ins_cbc0() int {
	val := cpu.B | (1 << 0)
	cpu.B = val
	return 4
}

// set 0,c
func (cpu *CPU) ins_cbc1() int {
	val := cpu.C | (1 << 0)
	cpu.C = val
	return 4
}

// set 0,d
func (cpu *CPU) ins_cbc2() int {
	val := cpu.D | (1 << 0)
	cpu.D = val
	return 4
}

// set 0,e
func (cpu *CPU) ins_cbc3() int {
	val := cpu.E | (1 << 0)
	cpu.E = val
	return 4
}

// set 0,h
func (cpu *CPU) ins_cbc4() int {
	val := cpu.H | (1 << 0)
	cpu.H = val
	return 4
}

// set 0,l
func (cpu *CPU) ins_cbc5() int {
	val := cpu.L | (1 << 0)
	cpu.L = val
	return 4
}

// set 0,(hl)
func (cpu *CPU) ins_cbc6() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 0,a
func (cpu *CPU) ins_cbc7() int {
	val := cpu.A | (1 << 0)
	cpu.A = val
	return 4
}

// set 1,b
func (cpu *CPU) ins_cbc8() int {
	val := cpu.B | (1 << 1)
	cpu.B = val
	return 4
}

// set 1,c
func (cpu *CPU) ins_cbc9() int {
	val := cpu.C | (1 << 1)
	cpu.C = val
	return 4
}

// set 1,d
func (cpu *CPU) ins_cbca() int {
	val := cpu.D | (1 << 1)
	cpu.D = val
	return 4
}

// set 1,e
func (cpu *CPU) ins_cbcb() int {
	val := cpu.E | (1 << 1)
	cpu.E = val
	return 4
}

// set 1,h
func (cpu *CPU) ins_cbcc() int {
	val := cpu.H | (1 << 1)
	cpu.H = val
	return 4
}

// set 1,l
func (cpu *CPU) ins_cbcd() int {
	val := cpu.L | (1 << 1)
	cpu.L = val
	return 4
}

// set 1,(hl)
func (cpu *CPU) ins_cbce() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 1,a
func (cpu *CPU) ins_cbcf() int {
	val := cpu.A | (1 << 1)
	cpu.A = val
	return 4
}

// set 2,b
func (cpu *CPU) ins_cbd0() int {
	val := cpu.B | (1 << 2)
	cpu.B = val
	return 4
}

// set 2,c
func (cpu *CPU) ins_cbd1() int {
	val := cpu.C | (1 << 2)
	cpu.C = val
	return 4
}

// set 2,d
func (cpu *CPU) ins_cbd2() int {
	val := cpu.D | (1 << 2)
	cpu.D = val
	return 4
}

// set 2,e
func (cpu *CPU) ins_cbd3() int {
	val := cpu.E | (1 << 2)
	cpu.E = val
	return 4
}

// set 2,h
func (cpu *CPU) ins_cbd4() int {
	val := cpu.H | (1 << 2)
	cpu.H = val
	return 4
}

// set 2,l
func (cpu *CPU) ins_cbd5() int {
	val := cpu.L | (1 << 2)
	cpu.L = val
	return 4
}

// set 2,(hl)
func (cpu *CPU) ins_cbd6() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 2,a
func (cpu *CPU) ins_cbd7() int {
	val := cpu.A | (1 << 2)
	cpu.A = val
	return 4
}

// set 3,b
func (cpu *CPU) ins_cbd8() int {
	val := cpu.B | (1 << 3)
	cpu.B = val
	return 4
}

// set 3,c
func (cpu *CPU) ins_cbd9() int {
	val := cpu.C | (1 << 3)
	cpu.C = val
	return 4
}

// set 3,d
func (cpu *CPU) ins_cbda() int {
	val := cpu.D | (1 << 3)
	cpu.D = val
	return 4
}

// set 3,e
func (cpu *CPU) ins_cbdb() int {
	val := cpu.E | (1 << 3)
	cpu.E = val
	return 4
}

// set 3,h
func (cpu *CPU) ins_cbdc() int {
	val := cpu.H | (1 << 3)
	cpu.H = val
	return 4
}

// set 3,l
func (cpu *CPU) ins_cbdd() int {
	val := cpu.L | (1 << 3)
	cpu.L = val
	return 4
}

// set 3,(hl)
func (cpu *CPU) ins_cbde() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 3,a
func (cpu *CPU) ins_cbdf() int {
	val := cpu.A | (1 << 3)
	cpu.A = val
	return 4
}

// set 4,b
func (cpu *CPU) ins_cbe0() int {
	val := cpu.B | (1 << 4)
	cpu.B = val
	return 4
}

// set 4,c
func (cpu *CPU) ins_cbe1() int {
	val := cpu.C | (1 << 4)
	cpu.C = val
	return 4
}

// set 4,d
func (cpu *CPU) ins_cbe2() int {
	val := cpu.D | (1 << 4)
	cpu.D = val
	return 4
}

// set 4,e
func (cpu *CPU) ins_cbe3() int {
	val := cpu.E | (1 << 4)
	cpu.E = val
	return 4
}

// set 4,h
func (cpu *CPU) ins_cbe4() int {
	val := cpu.H | (1 << 4)
	cpu.H = val
	return 4
}

// set 4,l
func (cpu *CPU) ins_cbe5() int {
	val := cpu.L | (1 << 4)
	cpu.L = val
	return 4
}

// set 4,(hl)
func (cpu *CPU) ins_cbe6() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 4,a
func (cpu *CPU) ins_cbe7() int {
	val := cpu.A | (1 << 4)
	cpu.A = val
	return 4
}

// set 5,b
func (cpu *CPU) ins_cbe8() int {
	val := cpu.B | (1 << 5)
	cpu.B = val
	return 4
}

// set 5,c
func (cpu *CPU) ins_cbe9() int {
	val := cpu.C | (1 << 5)
	cpu.C = val
	return 4
}

// set 5,d
func (cpu *CPU) ins_cbea() int {
	val := cpu.D | (1 << 5)
	cpu.D = val
	return 4
}

// set 5,e
func (cpu *CPU) ins_cbeb() int {
	val := cpu.E | (1 << 5)
	cpu.E = val
	return 4
}

// set 5,h
func (cpu *CPU) ins_cbec() int {
	val := cpu.H | (1 << 5)
	cpu.H = val
	return 4
}

// set 5,l
func (cpu *CPU) ins_cbed() int {
	val := cpu.L | (1 << 5)
	cpu.L = val
	return 4
}

// set 5,(hl)
func (cpu *CPU) ins_cbee() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 5,a
func (cpu *CPU) ins_cbef() int {
	val := cpu.A | (1 << 5)
	cpu.A = val
	return 4
}

// set 6,b
func (cpu *CPU) ins_cbf0() int {
	val := cpu.B | (1 << 6)
	cpu.B = val
	return 4
}

// set 6,c
func (cpu *CPU) ins_cbf1() int {
	val := cpu.C | (1 << 6)
	cpu.C = val
	return 4
}

// set 6,d
func (cpu *CPU) ins_cbf2() int {
	val := cpu.D | (1 << 6)
	cpu.D = val
	return 4
}

// set 6,e
func (cpu *CPU) ins_cbf3() int {
	val := cpu.E | (1 << 6)
	cpu.E = val
	return 4
}

// set 6,h
func (cpu *CPU) ins_cbf4() int {
	val := cpu.H | (1 << 6)
	cpu.H = val
	return 4
}

// set 6,l
func (cpu *CPU) ins_cbf5() int {
	val := cpu.L | (1 << 6)
	cpu.L = val
	return 4
}

// set 6,(hl)
func (cpu *CPU) ins_cbf6() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 6,a
func (cpu *CPU) ins_cbf7() int {
	val := cpu.A | (1 << 6)
	cpu.A = val
	return 4
}

// set 7,b
func (cpu *CPU) ins_cbf8() int {
	val := cpu.B | (1 << 7)
	cpu.B = val
	return 4
}

// set 7,c
func (cpu *CPU) ins_cbf9() int {
	val := cpu.C | (1 << 7)
	cpu.C = val
	return 4
}

// set 7,d
func (cpu *CPU) ins_cbfa() int {
	val := cpu.D | (1 << 7)
	cpu.D = val
	return 4
}

// set 7,e
func (cpu *CPU) ins_cbfb() int {
	val := cpu.E | (1 << 7)
	cpu.E = val
	return 4
}

// set 7,h
func (cpu *CPU) ins_cbfc() int {
	val := cpu.H | (1 << 7)
	cpu.H = val
	return 4
}

// set 7,l
func (cpu *CPU) ins_cbfd() int {
	val := cpu.L | (1 << 7)
	cpu.L = val
	return 4
}

// set 7,(hl)
func (cpu *CPU) ins_cbfe() int {
	n := cpu.get_hl()
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 7,a
func (cpu *CPU) ins_cbff() int {
	val := cpu.A | (1 << 7)
	cpu.A = val
	return 4
}

// add ix,bc
func (cpu *CPU) ins_dd09() int {
	s := cpu.get_bc()
	d := cpu.IX
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.IX = uint16(res)
	return 11
}

// djnz 0003
func (cpu *CPU) ins_dd10() int {
	d := offset16(cpu.get_n())
	cpu.B -= 1
	if cpu.B != 0 {
		cpu.PC += d
		return 13
	}
	return 8
}

// jr 0003
func (cpu *CPU) ins_dd18() int {
	panic("unimplemented instruction")
}

// add ix,de
func (cpu *CPU) ins_dd19() int {
	s := cpu.get_de()
	d := cpu.IX
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.IX = uint16(res)
	return 11
}

// jr nz,0003
func (cpu *CPU) ins_dd20() int {
	ofs := offset16(cpu.get_n())
	if (cpu.F & _ZF) == 0 {
		cpu.PC += ofs
		return 12
	}
	return 7
}

// ld ix,0000
func (cpu *CPU) ins_dd21() int {
	cpu.IX = cpu.get_nn()
	return 10
}

// ld (0000),ix
func (cpu *CPU) ins_dd22() int {
	nn := cpu.get_nn()
	cpu.mem.Wr8(nn, uint8(cpu.IX))
	cpu.mem.Wr8(nn+1, uint8(cpu.IX>>8))
	return 16
}

// inc ix
func (cpu *CPU) ins_dd23() int {
	cpu.IX += 1
	return 6
}

// inc ixh
func (cpu *CPU) ins_dd24() int {
	panic("unimplemented instruction")
}

// dec ixh
func (cpu *CPU) ins_dd25() int {
	panic("unimplemented instruction")
}

// ld ixh,00
func (cpu *CPU) ins_dd26() int {
	panic("unimplemented instruction")
}

// jr z,0003
func (cpu *CPU) ins_dd28() int {
	ofs := offset16(cpu.get_n())
	if (cpu.F & _ZF) != 0 {
		cpu.PC += ofs
		return 12
	}
	return 7
}

// add ix,ix
func (cpu *CPU) ins_dd29() int {
	s := cpu.IX
	d := cpu.IX
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.IX = uint16(res)
	return 11
}

// ld ix,(0000)
func (cpu *CPU) ins_dd2a() int {
	nn := cpu.get_nn()
	cpu.IX = uint16(cpu.mem.Rd8(nn+1)) << 8
	cpu.IX |= uint16(cpu.mem.Rd8(nn))
	return 16
}

// dec ix
func (cpu *CPU) ins_dd2b() int {
	cpu.IX -= 1
	return 6
}

// inc ixl
func (cpu *CPU) ins_dd2c() int {
	panic("unimplemented instruction")
}

// dec ixl
func (cpu *CPU) ins_dd2d() int {
	panic("unimplemented instruction")
}

// ld ixl,00
func (cpu *CPU) ins_dd2e() int {
	panic("unimplemented instruction")
}

// jr nc,0003
func (cpu *CPU) ins_dd30() int {
	ofs := offset16(cpu.get_n())
	if (cpu.F & _CF) == 0 {
		cpu.PC += ofs
		return 12
	}
	return 7
}

// inc (ix+00)
func (cpu *CPU) ins_dd34() int {
	adr := cpu.IX + offset16(cpu.get_n())
	n := cpu.mem.Rd8(adr) + 1
	cpu.mem.Wr8(adr, n)
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 19
}

// dec (ix+00)
func (cpu *CPU) ins_dd35() int {
	adr := cpu.IX + offset16(cpu.get_n())
	n := cpu.mem.Rd8(adr) - 1
	cpu.mem.Wr8(adr, n)
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 19
}

// ld (ix+00),00
func (cpu *CPU) ins_dd36() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IX+d, cpu.get_n())
	return 15
}

// jr c,0003
func (cpu *CPU) ins_dd38() int {
	ofs := offset16(cpu.get_n())
	if (cpu.F & _CF) != 0 {
		cpu.PC += ofs
		return 12
	}
	return 7
}

// add ix,sp
func (cpu *CPU) ins_dd39() int {
	s := cpu.SP
	d := cpu.IX
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.IX = uint16(res)
	return 11
}

// ld b,ixh
func (cpu *CPU) ins_dd44() int {
	panic("unimplemented instruction")
}

// ld b,ixl
func (cpu *CPU) ins_dd45() int {
	panic("unimplemented instruction")
}

// ld b,(ix+00)
func (cpu *CPU) ins_dd46() int {
	d := offset16(cpu.get_n())
	cpu.B = cpu.mem.Rd8(cpu.IX + d)
	return 15
}

// ld c,ixh
func (cpu *CPU) ins_dd4c() int {
	panic("unimplemented instruction")
}

// ld c,ixl
func (cpu *CPU) ins_dd4d() int {
	panic("unimplemented instruction")
}

// ld c,(ix+00)
func (cpu *CPU) ins_dd4e() int {
	d := offset16(cpu.get_n())
	cpu.C = cpu.mem.Rd8(cpu.IX + d)
	return 15
}

// ld d,ixh
func (cpu *CPU) ins_dd54() int {
	panic("unimplemented instruction")
}

// ld d,ixl
func (cpu *CPU) ins_dd55() int {
	panic("unimplemented instruction")
}

// ld d,(ix+00)
func (cpu *CPU) ins_dd56() int {
	d := offset16(cpu.get_n())
	cpu.D = cpu.mem.Rd8(cpu.IX + d)
	return 15
}

// ld e,ixh
func (cpu *CPU) ins_dd5c() int {
	panic("unimplemented instruction")
}

// ld e,ixl
func (cpu *CPU) ins_dd5d() int {
	panic("unimplemented instruction")
}

// ld e,(ix+00)
func (cpu *CPU) ins_dd5e() int {
	d := offset16(cpu.get_n())
	cpu.E = cpu.mem.Rd8(cpu.IX + d)
	return 15
}

// ld ixh,b
func (cpu *CPU) ins_dd60() int {
	panic("unimplemented instruction")
}

// ld ixh,c
func (cpu *CPU) ins_dd61() int {
	panic("unimplemented instruction")
}

// ld ixh,d
func (cpu *CPU) ins_dd62() int {
	panic("unimplemented instruction")
}

// ld ixh,e
func (cpu *CPU) ins_dd63() int {
	panic("unimplemented instruction")
}

// ld ixh,ixh
func (cpu *CPU) ins_dd64() int {
	panic("unimplemented instruction")
}

// ld ixh,ixl
func (cpu *CPU) ins_dd65() int {
	panic("unimplemented instruction")
}

// ld h,(ix+00)
func (cpu *CPU) ins_dd66() int {
	d := offset16(cpu.get_n())
	cpu.H = cpu.mem.Rd8(cpu.IX + d)
	return 15
}

// ld ixh,a
func (cpu *CPU) ins_dd67() int {
	panic("unimplemented instruction")
}

// ld ixl,b
func (cpu *CPU) ins_dd68() int {
	panic("unimplemented instruction")
}

// ld ixl,c
func (cpu *CPU) ins_dd69() int {
	panic("unimplemented instruction")
}

// ld ixl,d
func (cpu *CPU) ins_dd6a() int {
	panic("unimplemented instruction")
}

// ld ixl,e
func (cpu *CPU) ins_dd6b() int {
	panic("unimplemented instruction")
}

// ld ixl,ixh
func (cpu *CPU) ins_dd6c() int {
	panic("unimplemented instruction")
}

// ld ixl,ixl
func (cpu *CPU) ins_dd6d() int {
	panic("unimplemented instruction")
}

// ld l,(ix+00)
func (cpu *CPU) ins_dd6e() int {
	d := offset16(cpu.get_n())
	cpu.L = cpu.mem.Rd8(cpu.IX + d)
	return 15
}

// ld ixl,a
func (cpu *CPU) ins_dd6f() int {
	panic("unimplemented instruction")
}

// ld (ix+00),b
func (cpu *CPU) ins_dd70() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IX+d, cpu.B)
	return 15
}

// ld (ix+00),c
func (cpu *CPU) ins_dd71() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IX+d, cpu.C)
	return 15
}

// ld (ix+00),d
func (cpu *CPU) ins_dd72() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IX+d, cpu.D)
	return 15
}

// ld (ix+00),e
func (cpu *CPU) ins_dd73() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IX+d, cpu.E)
	return 15
}

// ld (ix+00),h
func (cpu *CPU) ins_dd74() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IX+d, cpu.H)
	return 15
}

// ld (ix+00),l
func (cpu *CPU) ins_dd75() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IX+d, cpu.L)
	return 15
}

// ld (ix+00),a
func (cpu *CPU) ins_dd77() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IX+d, cpu.A)
	return 15
}

// ld a,ixh
func (cpu *CPU) ins_dd7c() int {
	panic("unimplemented instruction")
}

// ld a,ixl
func (cpu *CPU) ins_dd7d() int {
	panic("unimplemented instruction")
}

// ld a,(ix+00)
func (cpu *CPU) ins_dd7e() int {
	d := offset16(cpu.get_n())
	cpu.A = cpu.mem.Rd8(cpu.IX + d)
	return 15
}

// add a,ixh
func (cpu *CPU) ins_dd84() int {
	panic("unimplemented instruction")
}

// add a,ixl
func (cpu *CPU) ins_dd85() int {
	panic("unimplemented instruction")
}

// add a,(ix+00)
func (cpu *CPU) ins_dd86() int {
	val := cpu.mem.Rd8(cpu.IX + offset16(cpu.get_n()))
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 15
}

// adc a,ixh
func (cpu *CPU) ins_dd8c() int {
	panic("unimplemented instruction")
}

// adc a,ixl
func (cpu *CPU) ins_dd8d() int {
	panic("unimplemented instruction")
}

// adc a,(ix+00)
func (cpu *CPU) ins_dd8e() int {
	val := cpu.mem.Rd8(cpu.IX + offset16(cpu.get_n()))
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 15
}

// sub ixh
func (cpu *CPU) ins_dd94() int {
	panic("unimplemented instruction")
}

// sub ixl
func (cpu *CPU) ins_dd95() int {
	panic("unimplemented instruction")
}

// sub (ix+00)
func (cpu *CPU) ins_dd96() int {
	val := cpu.mem.Rd8(cpu.IX + offset16(cpu.get_n()))
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 15
}

// sbc a,ixh
func (cpu *CPU) ins_dd9c() int {
	panic("unimplemented instruction")
}

// sbc a,ixl
func (cpu *CPU) ins_dd9d() int {
	panic("unimplemented instruction")
}

// sbc a,(ix+00)
func (cpu *CPU) ins_dd9e() int {
	val := cpu.mem.Rd8(cpu.IX + offset16(cpu.get_n()))
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 15
}

// and ixh
func (cpu *CPU) ins_dda4() int {
	panic("unimplemented instruction")
}

// and ixl
func (cpu *CPU) ins_dda5() int {
	panic("unimplemented instruction")
}

// and (ix+00)
func (cpu *CPU) ins_dda6() int {
	val := cpu.mem.Rd8(cpu.IX + offset16(cpu.get_n()))
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 15
}

// xor ixh
func (cpu *CPU) ins_ddac() int {
	panic("unimplemented instruction")
}

// xor ixl
func (cpu *CPU) ins_ddad() int {
	panic("unimplemented instruction")
}

// xor (ix+00)
func (cpu *CPU) ins_ddae() int {
	val := cpu.mem.Rd8(cpu.IX + offset16(cpu.get_n()))
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 15
}

// or ixh
func (cpu *CPU) ins_ddb4() int {
	panic("unimplemented instruction")
}

// or ixl
func (cpu *CPU) ins_ddb5() int {
	panic("unimplemented instruction")
}

// or (ix+00)
func (cpu *CPU) ins_ddb6() int {
	val := cpu.mem.Rd8(cpu.IX + offset16(cpu.get_n()))
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 15
}

// cp ixh
func (cpu *CPU) ins_ddbc() int {
	panic("unimplemented instruction")
}

// cp ixl
func (cpu *CPU) ins_ddbd() int {
	panic("unimplemented instruction")
}

// cp (ix+00)
func (cpu *CPU) ins_ddbe() int {
	val := cpu.mem.Rd8(cpu.IX + offset16(cpu.get_n()))
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 15
}

// pop ix
func (cpu *CPU) ins_dde1() int {
	cpu.IX = cpu.pop16()
	return 10
}

// ex (sp),ix
func (cpu *CPU) ins_dde3() int {
	tmp := cpu.mem.Rd16(cpu.SP)
	cpu.mem.Wr16(cpu.SP, cpu.IX)
	cpu.IX = tmp
	return 19
}

// push ix
func (cpu *CPU) ins_dde5() int {
	cpu.push16(cpu.IX)
	return 11
}

// jp ix
func (cpu *CPU) ins_dde9() int {
	cpu.PC = cpu.IX
	return 4
}

// ld sp,ix
func (cpu *CPU) ins_ddf9() int {
	panic("unimplemented instruction")
}

// rlc (ix+00),b
func (cpu *CPU) ins_ddcb0000(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rlc (ix+00),c
func (cpu *CPU) ins_ddcb0001(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rlc (ix+00),d
func (cpu *CPU) ins_ddcb0002(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rlc (ix+00),e
func (cpu *CPU) ins_ddcb0003(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rlc (ix+00),h
func (cpu *CPU) ins_ddcb0004(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rlc (ix+00),l
func (cpu *CPU) ins_ddcb0005(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rlc (ix+00)
func (cpu *CPU) ins_ddcb0006(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rlc (ix+00),a
func (cpu *CPU) ins_ddcb0007(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rrc (ix+00),b
func (cpu *CPU) ins_ddcb0008(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rrc (ix+00),c
func (cpu *CPU) ins_ddcb0009(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rrc (ix+00),d
func (cpu *CPU) ins_ddcb000a(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rrc (ix+00),e
func (cpu *CPU) ins_ddcb000b(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rrc (ix+00),h
func (cpu *CPU) ins_ddcb000c(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rrc (ix+00),l
func (cpu *CPU) ins_ddcb000d(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rrc (ix+00)
func (cpu *CPU) ins_ddcb000e(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rrc (ix+00),a
func (cpu *CPU) ins_ddcb000f(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rl (ix+00),b
func (cpu *CPU) ins_ddcb0010(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rl (ix+00),c
func (cpu *CPU) ins_ddcb0011(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rl (ix+00),d
func (cpu *CPU) ins_ddcb0012(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rl (ix+00),e
func (cpu *CPU) ins_ddcb0013(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rl (ix+00),h
func (cpu *CPU) ins_ddcb0014(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rl (ix+00),l
func (cpu *CPU) ins_ddcb0015(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rl (ix+00)
func (cpu *CPU) ins_ddcb0016(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rl (ix+00),a
func (cpu *CPU) ins_ddcb0017(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rr (ix+00),b
func (cpu *CPU) ins_ddcb0018(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rr (ix+00),c
func (cpu *CPU) ins_ddcb0019(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rr (ix+00),d
func (cpu *CPU) ins_ddcb001a(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rr (ix+00),e
func (cpu *CPU) ins_ddcb001b(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rr (ix+00),h
func (cpu *CPU) ins_ddcb001c(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rr (ix+00),l
func (cpu *CPU) ins_ddcb001d(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rr (ix+00)
func (cpu *CPU) ins_ddcb001e(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// rr (ix+00),a
func (cpu *CPU) ins_ddcb001f(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sla (ix+00),b
func (cpu *CPU) ins_ddcb0020(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sla (ix+00),c
func (cpu *CPU) ins_ddcb0021(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sla (ix+00),d
func (cpu *CPU) ins_ddcb0022(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sla (ix+00),e
func (cpu *CPU) ins_ddcb0023(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sla (ix+00),h
func (cpu *CPU) ins_ddcb0024(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sla (ix+00),l
func (cpu *CPU) ins_ddcb0025(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sla (ix+00)
func (cpu *CPU) ins_ddcb0026(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sla (ix+00),a
func (cpu *CPU) ins_ddcb0027(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sra (ix+00),b
func (cpu *CPU) ins_ddcb0028(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sra (ix+00),c
func (cpu *CPU) ins_ddcb0029(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sra (ix+00),d
func (cpu *CPU) ins_ddcb002a(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sra (ix+00),e
func (cpu *CPU) ins_ddcb002b(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sra (ix+00),h
func (cpu *CPU) ins_ddcb002c(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sra (ix+00),l
func (cpu *CPU) ins_ddcb002d(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sra (ix+00)
func (cpu *CPU) ins_ddcb002e(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sra (ix+00),a
func (cpu *CPU) ins_ddcb002f(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sll (ix+00),b
func (cpu *CPU) ins_ddcb0030(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sll (ix+00),c
func (cpu *CPU) ins_ddcb0031(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sll (ix+00),d
func (cpu *CPU) ins_ddcb0032(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sll (ix+00),e
func (cpu *CPU) ins_ddcb0033(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sll (ix+00),h
func (cpu *CPU) ins_ddcb0034(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sll (ix+00),l
func (cpu *CPU) ins_ddcb0035(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sll (ix+00)
func (cpu *CPU) ins_ddcb0036(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// sll (ix+00),a
func (cpu *CPU) ins_ddcb0037(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// srl (ix+00),b
func (cpu *CPU) ins_ddcb0038(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// srl (ix+00),c
func (cpu *CPU) ins_ddcb0039(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// srl (ix+00),d
func (cpu *CPU) ins_ddcb003a(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// srl (ix+00),e
func (cpu *CPU) ins_ddcb003b(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// srl (ix+00),h
func (cpu *CPU) ins_ddcb003c(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// srl (ix+00),l
func (cpu *CPU) ins_ddcb003d(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// srl (ix+00)
func (cpu *CPU) ins_ddcb003e(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// srl (ix+00),a
func (cpu *CPU) ins_ddcb003f(d uint8) int {
	res := cpu.mem.Rd8(cpu.IX + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IX+offset16(d), res)
	return 11
}

// bit 0,(ix+00)
func (cpu *CPU) ins_ddcb0040(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IX+offset16(d)) & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 1,(ix+00)
func (cpu *CPU) ins_ddcb0048(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IX+offset16(d)) & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 2,(ix+00)
func (cpu *CPU) ins_ddcb0050(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IX+offset16(d)) & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 3,(ix+00)
func (cpu *CPU) ins_ddcb0058(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IX+offset16(d)) & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 4,(ix+00)
func (cpu *CPU) ins_ddcb0060(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IX+offset16(d)) & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 5,(ix+00)
func (cpu *CPU) ins_ddcb0068(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IX+offset16(d)) & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 6,(ix+00)
func (cpu *CPU) ins_ddcb0070(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IX+offset16(d)) & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 7,(ix+00)
func (cpu *CPU) ins_ddcb0078(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IX+offset16(d)) & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// res 0,(ix+00),b
func (cpu *CPU) ins_ddcb0080(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 0,(ix+00),c
func (cpu *CPU) ins_ddcb0081(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 0,(ix+00),d
func (cpu *CPU) ins_ddcb0082(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 0,(ix+00),e
func (cpu *CPU) ins_ddcb0083(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 0,(ix+00),h
func (cpu *CPU) ins_ddcb0084(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 0,(ix+00),l
func (cpu *CPU) ins_ddcb0085(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 0,(ix+00)
func (cpu *CPU) ins_ddcb0086(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 0,(ix+00),a
func (cpu *CPU) ins_ddcb0087(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 1,(ix+00),b
func (cpu *CPU) ins_ddcb0088(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 1,(ix+00),c
func (cpu *CPU) ins_ddcb0089(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 1,(ix+00),d
func (cpu *CPU) ins_ddcb008a(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 1,(ix+00),e
func (cpu *CPU) ins_ddcb008b(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 1,(ix+00),h
func (cpu *CPU) ins_ddcb008c(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 1,(ix+00),l
func (cpu *CPU) ins_ddcb008d(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 1,(ix+00)
func (cpu *CPU) ins_ddcb008e(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 1,(ix+00),a
func (cpu *CPU) ins_ddcb008f(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 2,(ix+00),b
func (cpu *CPU) ins_ddcb0090(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 2,(ix+00),c
func (cpu *CPU) ins_ddcb0091(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 2,(ix+00),d
func (cpu *CPU) ins_ddcb0092(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 2,(ix+00),e
func (cpu *CPU) ins_ddcb0093(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 2,(ix+00),h
func (cpu *CPU) ins_ddcb0094(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 2,(ix+00),l
func (cpu *CPU) ins_ddcb0095(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 2,(ix+00)
func (cpu *CPU) ins_ddcb0096(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 2,(ix+00),a
func (cpu *CPU) ins_ddcb0097(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 3,(ix+00),b
func (cpu *CPU) ins_ddcb0098(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 3,(ix+00),c
func (cpu *CPU) ins_ddcb0099(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 3,(ix+00),d
func (cpu *CPU) ins_ddcb009a(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 3,(ix+00),e
func (cpu *CPU) ins_ddcb009b(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 3,(ix+00),h
func (cpu *CPU) ins_ddcb009c(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 3,(ix+00),l
func (cpu *CPU) ins_ddcb009d(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 3,(ix+00)
func (cpu *CPU) ins_ddcb009e(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 3,(ix+00),a
func (cpu *CPU) ins_ddcb009f(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 4,(ix+00),b
func (cpu *CPU) ins_ddcb00a0(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 4,(ix+00),c
func (cpu *CPU) ins_ddcb00a1(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 4,(ix+00),d
func (cpu *CPU) ins_ddcb00a2(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 4,(ix+00),e
func (cpu *CPU) ins_ddcb00a3(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 4,(ix+00),h
func (cpu *CPU) ins_ddcb00a4(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 4,(ix+00),l
func (cpu *CPU) ins_ddcb00a5(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 4,(ix+00)
func (cpu *CPU) ins_ddcb00a6(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 4,(ix+00),a
func (cpu *CPU) ins_ddcb00a7(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 5,(ix+00),b
func (cpu *CPU) ins_ddcb00a8(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 5,(ix+00),c
func (cpu *CPU) ins_ddcb00a9(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 5,(ix+00),d
func (cpu *CPU) ins_ddcb00aa(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 5,(ix+00),e
func (cpu *CPU) ins_ddcb00ab(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 5,(ix+00),h
func (cpu *CPU) ins_ddcb00ac(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 5,(ix+00),l
func (cpu *CPU) ins_ddcb00ad(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 5,(ix+00)
func (cpu *CPU) ins_ddcb00ae(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 5,(ix+00),a
func (cpu *CPU) ins_ddcb00af(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 6,(ix+00),b
func (cpu *CPU) ins_ddcb00b0(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 6,(ix+00),c
func (cpu *CPU) ins_ddcb00b1(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 6,(ix+00),d
func (cpu *CPU) ins_ddcb00b2(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 6,(ix+00),e
func (cpu *CPU) ins_ddcb00b3(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 6,(ix+00),h
func (cpu *CPU) ins_ddcb00b4(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 6,(ix+00),l
func (cpu *CPU) ins_ddcb00b5(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 6,(ix+00)
func (cpu *CPU) ins_ddcb00b6(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 6,(ix+00),a
func (cpu *CPU) ins_ddcb00b7(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 7,(ix+00),b
func (cpu *CPU) ins_ddcb00b8(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 7,(ix+00),c
func (cpu *CPU) ins_ddcb00b9(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 7,(ix+00),d
func (cpu *CPU) ins_ddcb00ba(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 7,(ix+00),e
func (cpu *CPU) ins_ddcb00bb(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 7,(ix+00),h
func (cpu *CPU) ins_ddcb00bc(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 7,(ix+00),l
func (cpu *CPU) ins_ddcb00bd(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 7,(ix+00)
func (cpu *CPU) ins_ddcb00be(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 7,(ix+00),a
func (cpu *CPU) ins_ddcb00bf(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 0,(ix+00),b
func (cpu *CPU) ins_ddcb00c0(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 0,(ix+00),c
func (cpu *CPU) ins_ddcb00c1(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 0,(ix+00),d
func (cpu *CPU) ins_ddcb00c2(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 0,(ix+00),e
func (cpu *CPU) ins_ddcb00c3(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 0,(ix+00),h
func (cpu *CPU) ins_ddcb00c4(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 0,(ix+00),l
func (cpu *CPU) ins_ddcb00c5(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 0,(ix+00)
func (cpu *CPU) ins_ddcb00c6(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 0,(ix+00),a
func (cpu *CPU) ins_ddcb00c7(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 1,(ix+00),b
func (cpu *CPU) ins_ddcb00c8(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 1,(ix+00),c
func (cpu *CPU) ins_ddcb00c9(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 1,(ix+00),d
func (cpu *CPU) ins_ddcb00ca(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 1,(ix+00),e
func (cpu *CPU) ins_ddcb00cb(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 1,(ix+00),h
func (cpu *CPU) ins_ddcb00cc(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 1,(ix+00),l
func (cpu *CPU) ins_ddcb00cd(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 1,(ix+00)
func (cpu *CPU) ins_ddcb00ce(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 1,(ix+00),a
func (cpu *CPU) ins_ddcb00cf(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 2,(ix+00),b
func (cpu *CPU) ins_ddcb00d0(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 2,(ix+00),c
func (cpu *CPU) ins_ddcb00d1(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 2,(ix+00),d
func (cpu *CPU) ins_ddcb00d2(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 2,(ix+00),e
func (cpu *CPU) ins_ddcb00d3(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 2,(ix+00),h
func (cpu *CPU) ins_ddcb00d4(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 2,(ix+00),l
func (cpu *CPU) ins_ddcb00d5(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 2,(ix+00)
func (cpu *CPU) ins_ddcb00d6(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 2,(ix+00),a
func (cpu *CPU) ins_ddcb00d7(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 3,(ix+00),b
func (cpu *CPU) ins_ddcb00d8(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 3,(ix+00),c
func (cpu *CPU) ins_ddcb00d9(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 3,(ix+00),d
func (cpu *CPU) ins_ddcb00da(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 3,(ix+00),e
func (cpu *CPU) ins_ddcb00db(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 3,(ix+00),h
func (cpu *CPU) ins_ddcb00dc(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 3,(ix+00),l
func (cpu *CPU) ins_ddcb00dd(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 3,(ix+00)
func (cpu *CPU) ins_ddcb00de(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 3,(ix+00),a
func (cpu *CPU) ins_ddcb00df(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 4,(ix+00),b
func (cpu *CPU) ins_ddcb00e0(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 4,(ix+00),c
func (cpu *CPU) ins_ddcb00e1(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 4,(ix+00),d
func (cpu *CPU) ins_ddcb00e2(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 4,(ix+00),e
func (cpu *CPU) ins_ddcb00e3(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 4,(ix+00),h
func (cpu *CPU) ins_ddcb00e4(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 4,(ix+00),l
func (cpu *CPU) ins_ddcb00e5(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 4,(ix+00)
func (cpu *CPU) ins_ddcb00e6(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 4,(ix+00),a
func (cpu *CPU) ins_ddcb00e7(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 5,(ix+00),b
func (cpu *CPU) ins_ddcb00e8(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 5,(ix+00),c
func (cpu *CPU) ins_ddcb00e9(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 5,(ix+00),d
func (cpu *CPU) ins_ddcb00ea(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 5,(ix+00),e
func (cpu *CPU) ins_ddcb00eb(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 5,(ix+00),h
func (cpu *CPU) ins_ddcb00ec(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 5,(ix+00),l
func (cpu *CPU) ins_ddcb00ed(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 5,(ix+00)
func (cpu *CPU) ins_ddcb00ee(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 5,(ix+00),a
func (cpu *CPU) ins_ddcb00ef(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 6,(ix+00),b
func (cpu *CPU) ins_ddcb00f0(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 6,(ix+00),c
func (cpu *CPU) ins_ddcb00f1(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 6,(ix+00),d
func (cpu *CPU) ins_ddcb00f2(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 6,(ix+00),e
func (cpu *CPU) ins_ddcb00f3(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 6,(ix+00),h
func (cpu *CPU) ins_ddcb00f4(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 6,(ix+00),l
func (cpu *CPU) ins_ddcb00f5(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 6,(ix+00)
func (cpu *CPU) ins_ddcb00f6(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 6,(ix+00),a
func (cpu *CPU) ins_ddcb00f7(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 7,(ix+00),b
func (cpu *CPU) ins_ddcb00f8(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 7,(ix+00),c
func (cpu *CPU) ins_ddcb00f9(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 7,(ix+00),d
func (cpu *CPU) ins_ddcb00fa(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 7,(ix+00),e
func (cpu *CPU) ins_ddcb00fb(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 7,(ix+00),h
func (cpu *CPU) ins_ddcb00fc(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 7,(ix+00),l
func (cpu *CPU) ins_ddcb00fd(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 7,(ix+00)
func (cpu *CPU) ins_ddcb00fe(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 7,(ix+00),a
func (cpu *CPU) ins_ddcb00ff(d uint8) int {
	n := cpu.IX + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// in b,(c)
func (cpu *CPU) ins_ed40() int {
	val := cpu.io.Rd8(cpu.get_bc())
	cpu.B = val
	cpu.F = (cpu.F & _CF) | flagsSZP[val]
	return 8
}

// out (c),b
func (cpu *CPU) ins_ed41() int {
	cpu.io.Wr8(cpu.get_bc(), cpu.B)
	return 8
}

// sbc hl,bc
func (cpu *CPU) ins_ed42() int {
	s := cpu.get_bc()
	d := cpu.get_hl()
	res := int(d) - int(s) - int(cpu.F&_CF)
	cpu.sub16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld (0000),bc
func (cpu *CPU) ins_ed43() int {
	nn := cpu.get_nn()
	cpu.mem.Wr8(nn, cpu.C)
	cpu.mem.Wr8(nn+1, cpu.B)
	return 16
}

// neg
func (cpu *CPU) ins_ed44() int {
	result := -int(cpu.A)
	cpu.subFlags(result, cpu.A)
	cpu.A = uint8(result)
	return 4
}

// retn
func (cpu *CPU) ins_ed45() int {
	panic("unimplemented instruction")
}

// im 0
func (cpu *CPU) ins_ed46() int {
	cpu.IM = 0
	return 4
}

// ld i,a
func (cpu *CPU) ins_ed47() int {
	cpu.I = cpu.A
	return 9
}

// in c,(c)
func (cpu *CPU) ins_ed48() int {
	val := cpu.io.Rd8(cpu.get_bc())
	cpu.C = val
	cpu.F = (cpu.F & _CF) | flagsSZP[val]
	return 8
}

// out (c),c
func (cpu *CPU) ins_ed49() int {
	cpu.io.Wr8(cpu.get_bc(), cpu.C)
	return 8
}

// adc hl,bc
func (cpu *CPU) ins_ed4a() int {
	s := cpu.get_bc()
	d := cpu.get_hl()
	res := int(d) + int(s) + int(cpu.F&_CF)
	cpu.adc16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld bc,(0000)
func (cpu *CPU) ins_ed4b() int {
	nn := cpu.get_nn()
	cpu.B = cpu.mem.Rd8(nn + 1)
	cpu.C = cpu.mem.Rd8(nn)
	return 16
}

// reti
func (cpu *CPU) ins_ed4d() int {
	panic("unimplemented instruction")
}

// ld r,a
func (cpu *CPU) ins_ed4f() int {
	cpu.R = cpu.A
	return 9
}

// in d,(c)
func (cpu *CPU) ins_ed50() int {
	val := cpu.io.Rd8(cpu.get_bc())
	cpu.D = val
	cpu.F = (cpu.F & _CF) | flagsSZP[val]
	return 8
}

// out (c),d
func (cpu *CPU) ins_ed51() int {
	cpu.io.Wr8(cpu.get_bc(), cpu.D)
	return 8
}

// sbc hl,de
func (cpu *CPU) ins_ed52() int {
	s := cpu.get_de()
	d := cpu.get_hl()
	res := int(d) - int(s) - int(cpu.F&_CF)
	cpu.sub16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld (0000),de
func (cpu *CPU) ins_ed53() int {
	nn := cpu.get_nn()
	cpu.mem.Wr8(nn, cpu.E)
	cpu.mem.Wr8(nn+1, cpu.D)
	return 16
}

// im 1
func (cpu *CPU) ins_ed56() int {
	cpu.IM = 1
	return 4
}

// ld a,i
func (cpu *CPU) ins_ed57() int {
	cpu.A = cpu.I
	cpu.F = (cpu.F & _CF) | (flagsSZ[cpu.A]) | (cpu.IFF2 << 2)
	return 9
}

// in e,(c)
func (cpu *CPU) ins_ed58() int {
	val := cpu.io.Rd8(cpu.get_bc())
	cpu.E = val
	cpu.F = (cpu.F & _CF) | flagsSZP[val]
	return 8
}

// out (c),e
func (cpu *CPU) ins_ed59() int {
	cpu.io.Wr8(cpu.get_bc(), cpu.E)
	return 8
}

// adc hl,de
func (cpu *CPU) ins_ed5a() int {
	s := cpu.get_de()
	d := cpu.get_hl()
	res := int(d) + int(s) + int(cpu.F&_CF)
	cpu.adc16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld de,(0000)
func (cpu *CPU) ins_ed5b() int {
	nn := cpu.get_nn()
	cpu.D = cpu.mem.Rd8(nn + 1)
	cpu.E = cpu.mem.Rd8(nn)
	return 16
}

// im 2
func (cpu *CPU) ins_ed5e() int {
	cpu.IM = 2
	return 4
}

// ld a,r
func (cpu *CPU) ins_ed5f() int {
	cpu.A = cpu.R
	cpu.F = (cpu.F & _CF) | (flagsSZ[cpu.A]) | (cpu.IFF2 << 2)
	return 9
}

// in h,(c)
func (cpu *CPU) ins_ed60() int {
	val := cpu.io.Rd8(cpu.get_bc())
	cpu.H = val
	cpu.F = (cpu.F & _CF) | flagsSZP[val]
	return 8
}

// out (c),h
func (cpu *CPU) ins_ed61() int {
	cpu.io.Wr8(cpu.get_bc(), cpu.H)
	return 8
}

// sbc hl,hl
func (cpu *CPU) ins_ed62() int {
	s := cpu.get_hl()
	d := cpu.get_hl()
	res := int(d) - int(s) - int(cpu.F&_CF)
	cpu.sub16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// rrd
func (cpu *CPU) ins_ed67() int {
	adr := cpu.get_hl()
	n := cpu.mem.Rd8(adr)
	cpu.mem.Wr8(adr, ((n>>4)|(cpu.A<<4))&0xff)
	cpu.A = (cpu.A & 0xf0) | (n & 0x0f)
	cpu.F = (cpu.F & _CF) | flagsSZP[cpu.A]
	return 14
}

// in l,(c)
func (cpu *CPU) ins_ed68() int {
	val := cpu.io.Rd8(cpu.get_bc())
	cpu.L = val
	cpu.F = (cpu.F & _CF) | flagsSZP[val]
	return 8
}

// out (c),l
func (cpu *CPU) ins_ed69() int {
	cpu.io.Wr8(cpu.get_bc(), cpu.L)
	return 8
}

// adc hl,hl
func (cpu *CPU) ins_ed6a() int {
	s := cpu.get_hl()
	d := cpu.get_hl()
	res := int(d) + int(s) + int(cpu.F&_CF)
	cpu.adc16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// rld
func (cpu *CPU) ins_ed6f() int {
	adr := cpu.get_hl()
	n := cpu.mem.Rd8(adr)
	cpu.mem.Wr8(adr, ((n<<4)|(cpu.A&0x0f))&0xff)
	cpu.A = (cpu.A & 0xf0) | (n >> 4)
	cpu.F = (cpu.F & _CF) | flagsSZP[cpu.A]
	return 14
}

// in (c)
func (cpu *CPU) ins_ed70() int {
	val := cpu.io.Rd8(cpu.get_bc())
	cpu.F = (cpu.F & _CF) | flagsSZP[val]
	return 8
}

// out (c)
func (cpu *CPU) ins_ed71() int {
	cpu.io.Wr8(cpu.get_bc(), 0)
	return 8
}

// sbc hl,sp
func (cpu *CPU) ins_ed72() int {
	s := cpu.SP
	d := cpu.get_hl()
	res := int(d) - int(s) - int(cpu.F&_CF)
	cpu.sub16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld (0000),sp
func (cpu *CPU) ins_ed73() int {
	nn := cpu.get_nn()
	cpu.mem.Wr8(nn, uint8(cpu.SP))
	cpu.mem.Wr8(nn+1, uint8(cpu.SP>>8))
	return 16
}

// in a,(c)
func (cpu *CPU) ins_ed78() int {
	val := cpu.io.Rd8(cpu.get_bc())
	cpu.A = val
	cpu.F = (cpu.F & _CF) | flagsSZP[val]
	return 8
}

// out (c),a
func (cpu *CPU) ins_ed79() int {
	cpu.io.Wr8(cpu.get_bc(), cpu.A)
	return 8
}

// adc hl,sp
func (cpu *CPU) ins_ed7a() int {
	s := cpu.SP
	d := cpu.get_hl()
	res := int(d) + int(s) + int(cpu.F&_CF)
	cpu.adc16Flags(res, d, s)
	cpu.set_hl(uint16(res))
	return 11
}

// ld sp,(0000)
func (cpu *CPU) ins_ed7b() int {
	nn := cpu.get_nn()
	cpu.SP = uint16(cpu.mem.Rd8(nn+1)) << 8
	cpu.SP |= uint16(cpu.mem.Rd8(nn))
	return 16
}

// ldi
func (cpu *CPU) ins_eda0() int {
	d := cpu.get_de()
	s := cpu.get_hl()
	n := cpu.get_bc() - 1
	val := cpu.mem.Rd8(s)
	cpu.mem.Wr8(d, val)
	cpu.F &= (_SF | _ZF | _CF)
	if ((cpu.A + val) & 0x02) != 0 {
		cpu.F |= _YF
	}
	if ((cpu.A + val) & 0x08) != 0 {
		cpu.F |= _XF
	}
	cpu.set_de(d + 1)
	cpu.set_hl(s + 1)
	cpu.set_bc(n)
	if n != 0 {
		cpu.F |= _VF
	}
	return 12
}

// cpi
func (cpu *CPU) ins_eda1() int {
	s := cpu.get_hl()
	n := cpu.get_bc() - 1
	val := cpu.mem.Rd8(s)
	res := cpu.A - val
	cpu.F = (cpu.F & _CF) | _NF
	cpu.F |= (flagsSZ[res] &^ (_YF | _XF))
	cpu.F |= ((cpu.A ^ val ^ res) & _HF)
	if (cpu.F & _HF) != 0 {
		res -= 1
	}
	if (res & 0x02) != 0 {
		cpu.F |= _YF
	}
	if (res & 0x08) != 0 {
		cpu.F |= _XF
	}
	cpu.set_hl(s + 1)
	cpu.set_bc(n)
	if n != 0 {
		cpu.F |= _VF
	}
	return 12
}

// ini
func (cpu *CPU) ins_eda2() int {
	panic("unimplemented instruction")
}

// outi
func (cpu *CPU) ins_eda3() int {
	panic("unimplemented instruction")
}

// ldd
func (cpu *CPU) ins_eda8() int {
	d := cpu.get_de()
	s := cpu.get_hl()
	n := cpu.get_bc() - 1
	val := cpu.mem.Rd8(s)
	cpu.mem.Wr8(d, val)
	cpu.F &= (_SF | _ZF | _CF)
	if ((cpu.A + val) & 0x02) != 0 {
		cpu.F |= _YF
	}
	if ((cpu.A + val) & 0x08) != 0 {
		cpu.F |= _XF
	}
	cpu.set_de(d - 1)
	cpu.set_hl(s - 1)
	cpu.set_bc(n)
	if n != 0 {
		cpu.F |= _VF
	}
	return 12
}

// cpd
func (cpu *CPU) ins_eda9() int {
	s := cpu.get_hl()
	n := cpu.get_bc() - 1
	val := cpu.mem.Rd8(s)
	res := cpu.A - val
	cpu.F = (cpu.F & _CF) | _NF
	cpu.F |= (flagsSZ[res] &^ (_YF | _XF))
	cpu.F |= ((cpu.A ^ val ^ res) & _HF)
	if (cpu.F & _HF) != 0 {
		res -= 1
	}
	if (res & 0x02) != 0 {
		cpu.F |= _YF
	}
	if (res & 0x08) != 0 {
		cpu.F |= _XF
	}
	cpu.set_hl(s - 1)
	cpu.set_bc(n)
	if n != 0 {
		cpu.F |= _VF
	}
	return 12
}

// ind
func (cpu *CPU) ins_edaa() int {
	panic("unimplemented instruction")
}

// outd
func (cpu *CPU) ins_edab() int {
	panic("unimplemented instruction")
}

// ldir
func (cpu *CPU) ins_edb0() int {
	d := cpu.get_de()
	s := cpu.get_hl()
	n := cpu.get_bc() - 1
	val := cpu.mem.Rd8(s)
	cpu.mem.Wr8(d, val)
	cpu.F &= (_SF | _ZF | _CF)
	if ((cpu.A + val) & 0x02) != 0 {
		cpu.F |= _YF
	}
	if ((cpu.A + val) & 0x08) != 0 {
		cpu.F |= _XF
	}
	cpu.set_de(d + 1)
	cpu.set_hl(s + 1)
	cpu.set_bc(n)
	if n != 0 {
		cpu.F |= _VF
		cpu.PC -= 2
		return 17
	}
	return 12
}

// cpir
func (cpu *CPU) ins_edb1() int {
	s := cpu.get_hl()
	n := cpu.get_bc() - 1
	val := cpu.mem.Rd8(s)
	res := cpu.A - val
	cpu.F = (cpu.F & _CF) | _NF
	cpu.F |= (flagsSZ[res] &^ (_YF | _XF))
	cpu.F |= ((cpu.A ^ val ^ res) & _HF)
	if (cpu.F & _HF) != 0 {
		res -= 1
	}
	if (res & 0x02) != 0 {
		cpu.F |= _YF
	}
	if (res & 0x08) != 0 {
		cpu.F |= _XF
	}
	cpu.set_hl(s + 1)
	cpu.set_bc(n)
	if (n != 0) && ((cpu.F & _ZF) == 0) {
		cpu.PC -= 2
		return 17
	}
	return 12
}

// inir
func (cpu *CPU) ins_edb2() int {
	panic("unimplemented instruction")
}

// otir
func (cpu *CPU) ins_edb3() int {
	panic("unimplemented instruction")
}

// lddr
func (cpu *CPU) ins_edb8() int {
	d := cpu.get_de()
	s := cpu.get_hl()
	n := cpu.get_bc() - 1
	val := cpu.mem.Rd8(s)
	cpu.mem.Wr8(d, val)
	cpu.F &= (_SF | _ZF | _CF)
	if ((cpu.A + val) & 0x02) != 0 {
		cpu.F |= _YF
	}
	if ((cpu.A + val) & 0x08) != 0 {
		cpu.F |= _XF
	}
	cpu.set_de(d - 1)
	cpu.set_hl(s - 1)
	cpu.set_bc(n)
	if n != 0 {
		cpu.F |= _VF
		cpu.PC -= 2
		return 17
	}
	return 12
}

// cpdr
func (cpu *CPU) ins_edb9() int {
	s := cpu.get_hl()
	n := cpu.get_bc() - 1
	val := cpu.mem.Rd8(s)
	res := cpu.A - val
	cpu.F = (cpu.F & _CF) | _NF
	cpu.F |= (flagsSZ[res] &^ (_YF | _XF))
	cpu.F |= ((cpu.A ^ val ^ res) & _HF)
	if (cpu.F & _HF) != 0 {
		res -= 1
	}
	if (res & 0x02) != 0 {
		cpu.F |= _YF
	}
	if (res & 0x08) != 0 {
		cpu.F |= _XF
	}
	cpu.set_hl(s - 1)
	cpu.set_bc(n)
	if (n != 0) && ((cpu.F & _ZF) == 0) {
		cpu.PC -= 2
		return 17
	}
	return 12
}

// indr
func (cpu *CPU) ins_edba() int {
	panic("unimplemented instruction")
}

// otdr
func (cpu *CPU) ins_edbb() int {
	panic("unimplemented instruction")
}

// add iy,bc
func (cpu *CPU) ins_fd09() int {
	s := cpu.get_bc()
	d := cpu.IY
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.IY = uint16(res)
	return 11
}

// add iy,de
func (cpu *CPU) ins_fd19() int {
	s := cpu.get_de()
	d := cpu.IY
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.IY = uint16(res)
	return 11
}

// ld iy,0000
func (cpu *CPU) ins_fd21() int {
	cpu.IY = cpu.get_nn()
	return 10
}

// ld (0000),iy
func (cpu *CPU) ins_fd22() int {
	nn := cpu.get_nn()
	cpu.mem.Wr8(nn, uint8(cpu.IY))
	cpu.mem.Wr8(nn+1, uint8(cpu.IY>>8))
	return 16
}

// inc iy
func (cpu *CPU) ins_fd23() int {
	cpu.IY += 1
	return 6
}

// inc iyh
func (cpu *CPU) ins_fd24() int {
	panic("unimplemented instruction")
}

// dec iyh
func (cpu *CPU) ins_fd25() int {
	panic("unimplemented instruction")
}

// ld iyh,00
func (cpu *CPU) ins_fd26() int {
	panic("unimplemented instruction")
}

// add iy,iy
func (cpu *CPU) ins_fd29() int {
	s := cpu.IY
	d := cpu.IY
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.IY = uint16(res)
	return 11
}

// ld iy,(0000)
func (cpu *CPU) ins_fd2a() int {
	nn := cpu.get_nn()
	cpu.IY = uint16(cpu.mem.Rd8(nn+1)) << 8
	cpu.IY |= uint16(cpu.mem.Rd8(nn))
	return 16
}

// dec iy
func (cpu *CPU) ins_fd2b() int {
	cpu.IY -= 1
	return 6
}

// inc iyl
func (cpu *CPU) ins_fd2c() int {
	panic("unimplemented instruction")
}

// dec iyl
func (cpu *CPU) ins_fd2d() int {
	panic("unimplemented instruction")
}

// ld iyl,00
func (cpu *CPU) ins_fd2e() int {
	panic("unimplemented instruction")
}

// inc (iy+00)
func (cpu *CPU) ins_fd34() int {
	adr := cpu.IY + offset16(cpu.get_n())
	n := cpu.mem.Rd8(adr) + 1
	cpu.mem.Wr8(adr, n)
	cpu.F = (cpu.F & _CF) | flagsSZHVinc[n]
	return 19
}

// dec (iy+00)
func (cpu *CPU) ins_fd35() int {
	adr := cpu.IY + offset16(cpu.get_n())
	n := cpu.mem.Rd8(adr) - 1
	cpu.mem.Wr8(adr, n)
	cpu.F = (cpu.F & _CF) | flagsSZHVdec[n]
	return 19
}

// ld (iy+00),00
func (cpu *CPU) ins_fd36() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IY+d, cpu.get_n())
	return 15
}

// add iy,sp
func (cpu *CPU) ins_fd39() int {
	s := cpu.SP
	d := cpu.IY
	res := int(d) + int(s)
	cpu.add16Flags(res, d, s)
	cpu.IY = uint16(res)
	return 11
}

// ld b,iyh
func (cpu *CPU) ins_fd44() int {
	panic("unimplemented instruction")
}

// ld b,iyl
func (cpu *CPU) ins_fd45() int {
	panic("unimplemented instruction")
}

// ld b,(iy+00)
func (cpu *CPU) ins_fd46() int {
	d := offset16(cpu.get_n())
	cpu.B = cpu.mem.Rd8(cpu.IY + d)
	return 15
}

// ld c,iyh
func (cpu *CPU) ins_fd4c() int {
	panic("unimplemented instruction")
}

// ld c,iyl
func (cpu *CPU) ins_fd4d() int {
	panic("unimplemented instruction")
}

// ld c,(iy+00)
func (cpu *CPU) ins_fd4e() int {
	d := offset16(cpu.get_n())
	cpu.C = cpu.mem.Rd8(cpu.IY + d)
	return 15
}

// ld d,iyh
func (cpu *CPU) ins_fd54() int {
	panic("unimplemented instruction")
}

// ld d,iyl
func (cpu *CPU) ins_fd55() int {
	panic("unimplemented instruction")
}

// ld d,(iy+00)
func (cpu *CPU) ins_fd56() int {
	d := offset16(cpu.get_n())
	cpu.D = cpu.mem.Rd8(cpu.IY + d)
	return 15
}

// ld e,iyh
func (cpu *CPU) ins_fd5c() int {
	panic("unimplemented instruction")
}

// ld e,iyl
func (cpu *CPU) ins_fd5d() int {
	panic("unimplemented instruction")
}

// ld e,(iy+00)
func (cpu *CPU) ins_fd5e() int {
	d := offset16(cpu.get_n())
	cpu.E = cpu.mem.Rd8(cpu.IY + d)
	return 15
}

// ld iyh,b
func (cpu *CPU) ins_fd60() int {
	panic("unimplemented instruction")
}

// ld iyh,c
func (cpu *CPU) ins_fd61() int {
	panic("unimplemented instruction")
}

// ld iyh,d
func (cpu *CPU) ins_fd62() int {
	panic("unimplemented instruction")
}

// ld iyh,e
func (cpu *CPU) ins_fd63() int {
	panic("unimplemented instruction")
}

// ld iyh,iyh
func (cpu *CPU) ins_fd64() int {
	panic("unimplemented instruction")
}

// ld iyh,iyl
func (cpu *CPU) ins_fd65() int {
	panic("unimplemented instruction")
}

// ld h,(iy+00)
func (cpu *CPU) ins_fd66() int {
	d := offset16(cpu.get_n())
	cpu.H = cpu.mem.Rd8(cpu.IY + d)
	return 15
}

// ld iyh,a
func (cpu *CPU) ins_fd67() int {
	panic("unimplemented instruction")
}

// ld iyl,b
func (cpu *CPU) ins_fd68() int {
	panic("unimplemented instruction")
}

// ld iyl,c
func (cpu *CPU) ins_fd69() int {
	panic("unimplemented instruction")
}

// ld iyl,d
func (cpu *CPU) ins_fd6a() int {
	panic("unimplemented instruction")
}

// ld iyl,e
func (cpu *CPU) ins_fd6b() int {
	panic("unimplemented instruction")
}

// ld iyl,iyh
func (cpu *CPU) ins_fd6c() int {
	panic("unimplemented instruction")
}

// ld iyl,iyl
func (cpu *CPU) ins_fd6d() int {
	panic("unimplemented instruction")
}

// ld l,(iy+00)
func (cpu *CPU) ins_fd6e() int {
	d := offset16(cpu.get_n())
	cpu.L = cpu.mem.Rd8(cpu.IY + d)
	return 15
}

// ld iyl,a
func (cpu *CPU) ins_fd6f() int {
	panic("unimplemented instruction")
}

// ld (iy+00),b
func (cpu *CPU) ins_fd70() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IY+d, cpu.B)
	return 15
}

// ld (iy+00),c
func (cpu *CPU) ins_fd71() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IY+d, cpu.C)
	return 15
}

// ld (iy+00),d
func (cpu *CPU) ins_fd72() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IY+d, cpu.D)
	return 15
}

// ld (iy+00),e
func (cpu *CPU) ins_fd73() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IY+d, cpu.E)
	return 15
}

// ld (iy+00),h
func (cpu *CPU) ins_fd74() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IY+d, cpu.H)
	return 15
}

// ld (iy+00),l
func (cpu *CPU) ins_fd75() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IY+d, cpu.L)
	return 15
}

// ld (iy+00),a
func (cpu *CPU) ins_fd77() int {
	d := offset16(cpu.get_n())
	cpu.mem.Wr8(cpu.IY+d, cpu.A)
	return 15
}

// ld a,iyh
func (cpu *CPU) ins_fd7c() int {
	panic("unimplemented instruction")
}

// ld a,iyl
func (cpu *CPU) ins_fd7d() int {
	panic("unimplemented instruction")
}

// ld a,(iy+00)
func (cpu *CPU) ins_fd7e() int {
	d := offset16(cpu.get_n())
	cpu.A = cpu.mem.Rd8(cpu.IY + d)
	return 15
}

// add a,iyh
func (cpu *CPU) ins_fd84() int {
	panic("unimplemented instruction")
}

// add a,iyl
func (cpu *CPU) ins_fd85() int {
	panic("unimplemented instruction")
}

// add a,(iy+00)
func (cpu *CPU) ins_fd86() int {
	val := cpu.mem.Rd8(cpu.IY + offset16(cpu.get_n()))
	result := int(cpu.A) + int(val)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 15
}

// adc a,iyh
func (cpu *CPU) ins_fd8c() int {
	panic("unimplemented instruction")
}

// adc a,iyl
func (cpu *CPU) ins_fd8d() int {
	panic("unimplemented instruction")
}

// adc a,(iy+00)
func (cpu *CPU) ins_fd8e() int {
	val := cpu.mem.Rd8(cpu.IY + offset16(cpu.get_n()))
	result := int(cpu.A) + int(val) + int(cpu.F&_CF)
	cpu.addFlags(result, val)
	cpu.A = uint8(result)
	return 15
}

// sub iyh
func (cpu *CPU) ins_fd94() int {
	panic("unimplemented instruction")
}

// sub iyl
func (cpu *CPU) ins_fd95() int {
	panic("unimplemented instruction")
}

// sub (iy+00)
func (cpu *CPU) ins_fd96() int {
	val := cpu.mem.Rd8(cpu.IY + offset16(cpu.get_n()))
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 15
}

// sbc a,iyh
func (cpu *CPU) ins_fd9c() int {
	panic("unimplemented instruction")
}

// sbc a,iyl
func (cpu *CPU) ins_fd9d() int {
	panic("unimplemented instruction")
}

// sbc a,(iy+00)
func (cpu *CPU) ins_fd9e() int {
	val := cpu.mem.Rd8(cpu.IY + offset16(cpu.get_n()))
	result := int(cpu.A) - int(val) - int(cpu.A&_CF)
	cpu.subFlags(result, val)
	cpu.A = uint8(result)
	return 15
}

// and iyh
func (cpu *CPU) ins_fda4() int {
	panic("unimplemented instruction")
}

// and iyl
func (cpu *CPU) ins_fda5() int {
	panic("unimplemented instruction")
}

// and (iy+00)
func (cpu *CPU) ins_fda6() int {
	val := cpu.mem.Rd8(cpu.IY + offset16(cpu.get_n()))
	cpu.A &= val
	cpu.F = flagsSZP[cpu.A] | _HF
	return 15
}

// xor iyh
func (cpu *CPU) ins_fdac() int {
	panic("unimplemented instruction")
}

// xor iyl
func (cpu *CPU) ins_fdad() int {
	panic("unimplemented instruction")
}

// xor (iy+00)
func (cpu *CPU) ins_fdae() int {
	val := cpu.mem.Rd8(cpu.IY + offset16(cpu.get_n()))
	cpu.A ^= val
	cpu.F = flagsSZP[cpu.A]
	return 15
}

// or iyh
func (cpu *CPU) ins_fdb4() int {
	panic("unimplemented instruction")
}

// or iyl
func (cpu *CPU) ins_fdb5() int {
	panic("unimplemented instruction")
}

// or (iy+00)
func (cpu *CPU) ins_fdb6() int {
	val := cpu.mem.Rd8(cpu.IY + offset16(cpu.get_n()))
	cpu.A |= val
	cpu.F = flagsSZP[cpu.A]
	return 15
}

// cp iyh
func (cpu *CPU) ins_fdbc() int {
	panic("unimplemented instruction")
}

// cp iyl
func (cpu *CPU) ins_fdbd() int {
	panic("unimplemented instruction")
}

// cp (iy+00)
func (cpu *CPU) ins_fdbe() int {
	val := cpu.mem.Rd8(cpu.IY + offset16(cpu.get_n()))
	result := int(cpu.A) - int(val)
	cpu.subFlags(result, val)
	return 15
}

// pop iy
func (cpu *CPU) ins_fde1() int {
	cpu.IY = cpu.pop16()
	return 10
}

// ex (sp),iy
func (cpu *CPU) ins_fde3() int {
	tmp := cpu.mem.Rd16(cpu.SP)
	cpu.mem.Wr16(cpu.SP, cpu.IY)
	cpu.IY = tmp
	return 19
}

// push iy
func (cpu *CPU) ins_fde5() int {
	cpu.push16(cpu.IY)
	return 11
}

// jp iy
func (cpu *CPU) ins_fde9() int {
	cpu.PC = cpu.IY
	return 4
}

// ld sp,iy
func (cpu *CPU) ins_fdf9() int {
	panic("unimplemented instruction")
}

// rlc (iy+00),b
func (cpu *CPU) ins_fdcb0000(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rlc (iy+00),c
func (cpu *CPU) ins_fdcb0001(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rlc (iy+00),d
func (cpu *CPU) ins_fdcb0002(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rlc (iy+00),e
func (cpu *CPU) ins_fdcb0003(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rlc (iy+00),h
func (cpu *CPU) ins_fdcb0004(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rlc (iy+00),l
func (cpu *CPU) ins_fdcb0005(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rlc (iy+00)
func (cpu *CPU) ins_fdcb0006(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rlc (iy+00),a
func (cpu *CPU) ins_fdcb0007(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (res >> 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rrc (iy+00),b
func (cpu *CPU) ins_fdcb0008(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rrc (iy+00),c
func (cpu *CPU) ins_fdcb0009(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rrc (iy+00),d
func (cpu *CPU) ins_fdcb000a(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rrc (iy+00),e
func (cpu *CPU) ins_fdcb000b(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rrc (iy+00),h
func (cpu *CPU) ins_fdcb000c(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rrc (iy+00),l
func (cpu *CPU) ins_fdcb000d(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rrc (iy+00)
func (cpu *CPU) ins_fdcb000e(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rrc (iy+00),a
func (cpu *CPU) ins_fdcb000f(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rl (iy+00),b
func (cpu *CPU) ins_fdcb0010(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rl (iy+00),c
func (cpu *CPU) ins_fdcb0011(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rl (iy+00),d
func (cpu *CPU) ins_fdcb0012(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rl (iy+00),e
func (cpu *CPU) ins_fdcb0013(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rl (iy+00),h
func (cpu *CPU) ins_fdcb0014(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rl (iy+00),l
func (cpu *CPU) ins_fdcb0015(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rl (iy+00)
func (cpu *CPU) ins_fdcb0016(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rl (iy+00),a
func (cpu *CPU) ins_fdcb0017(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | (cpu.F & _CF)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rr (iy+00),b
func (cpu *CPU) ins_fdcb0018(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rr (iy+00),c
func (cpu *CPU) ins_fdcb0019(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rr (iy+00),d
func (cpu *CPU) ins_fdcb001a(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rr (iy+00),e
func (cpu *CPU) ins_fdcb001b(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rr (iy+00),h
func (cpu *CPU) ins_fdcb001c(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rr (iy+00),l
func (cpu *CPU) ins_fdcb001d(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rr (iy+00)
func (cpu *CPU) ins_fdcb001e(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// rr (iy+00),a
func (cpu *CPU) ins_fdcb001f(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (cpu.F << 7)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sla (iy+00),b
func (cpu *CPU) ins_fdcb0020(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sla (iy+00),c
func (cpu *CPU) ins_fdcb0021(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sla (iy+00),d
func (cpu *CPU) ins_fdcb0022(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sla (iy+00),e
func (cpu *CPU) ins_fdcb0023(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sla (iy+00),h
func (cpu *CPU) ins_fdcb0024(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sla (iy+00),l
func (cpu *CPU) ins_fdcb0025(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sla (iy+00)
func (cpu *CPU) ins_fdcb0026(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sla (iy+00),a
func (cpu *CPU) ins_fdcb0027(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = (res << 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sra (iy+00),b
func (cpu *CPU) ins_fdcb0028(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sra (iy+00),c
func (cpu *CPU) ins_fdcb0029(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sra (iy+00),d
func (cpu *CPU) ins_fdcb002a(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sra (iy+00),e
func (cpu *CPU) ins_fdcb002b(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sra (iy+00),h
func (cpu *CPU) ins_fdcb002c(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sra (iy+00),l
func (cpu *CPU) ins_fdcb002d(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sra (iy+00)
func (cpu *CPU) ins_fdcb002e(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sra (iy+00),a
func (cpu *CPU) ins_fdcb002f(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = ((res >> 1) | (res & 0x80)) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sll (iy+00),b
func (cpu *CPU) ins_fdcb0030(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sll (iy+00),c
func (cpu *CPU) ins_fdcb0031(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sll (iy+00),d
func (cpu *CPU) ins_fdcb0032(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sll (iy+00),e
func (cpu *CPU) ins_fdcb0033(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sll (iy+00),h
func (cpu *CPU) ins_fdcb0034(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sll (iy+00),l
func (cpu *CPU) ins_fdcb0035(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sll (iy+00)
func (cpu *CPU) ins_fdcb0036(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// sll (iy+00),a
func (cpu *CPU) ins_fdcb0037(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x80) != 0 {
		cf = _CF
	}
	res = ((res << 1) | 0x01) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// srl (iy+00),b
func (cpu *CPU) ins_fdcb0038(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.B = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// srl (iy+00),c
func (cpu *CPU) ins_fdcb0039(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.C = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// srl (iy+00),d
func (cpu *CPU) ins_fdcb003a(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.D = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// srl (iy+00),e
func (cpu *CPU) ins_fdcb003b(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.E = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// srl (iy+00),h
func (cpu *CPU) ins_fdcb003c(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.H = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// srl (iy+00),l
func (cpu *CPU) ins_fdcb003d(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.L = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// srl (iy+00)
func (cpu *CPU) ins_fdcb003e(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// srl (iy+00),a
func (cpu *CPU) ins_fdcb003f(d uint8) int {
	res := cpu.mem.Rd8(cpu.IY + offset16(d))
	var cf uint8
	if (res & 0x01) != 0 {
		cf = _CF
	}
	res = (res >> 1) & 0xff
	cpu.F = flagsSZP[res] | cf
	cpu.A = res
	cpu.mem.Wr8(cpu.IY+offset16(d), res)
	return 11
}

// bit 0,(iy+00)
func (cpu *CPU) ins_fdcb0040(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IY+offset16(d)) & (1 << 0)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 1,(iy+00)
func (cpu *CPU) ins_fdcb0048(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IY+offset16(d)) & (1 << 1)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 2,(iy+00)
func (cpu *CPU) ins_fdcb0050(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IY+offset16(d)) & (1 << 2)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 3,(iy+00)
func (cpu *CPU) ins_fdcb0058(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IY+offset16(d)) & (1 << 3)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 4,(iy+00)
func (cpu *CPU) ins_fdcb0060(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IY+offset16(d)) & (1 << 4)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 5,(iy+00)
func (cpu *CPU) ins_fdcb0068(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IY+offset16(d)) & (1 << 5)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 6,(iy+00)
func (cpu *CPU) ins_fdcb0070(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IY+offset16(d)) & (1 << 6)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// bit 7,(iy+00)
func (cpu *CPU) ins_fdcb0078(d uint8) int {
	bit := cpu.mem.Rd8(cpu.IY+offset16(d)) & (1 << 7)
	var zf uint8
	if bit == 0 {
		zf = _ZF
	}
	cpu.F = (cpu.F & _CF) | _HF | zf
	return 8
}

// res 0,(iy+00),b
func (cpu *CPU) ins_fdcb0080(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 0,(iy+00),c
func (cpu *CPU) ins_fdcb0081(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 0,(iy+00),d
func (cpu *CPU) ins_fdcb0082(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 0,(iy+00),e
func (cpu *CPU) ins_fdcb0083(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 0,(iy+00),h
func (cpu *CPU) ins_fdcb0084(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 0,(iy+00),l
func (cpu *CPU) ins_fdcb0085(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 0,(iy+00)
func (cpu *CPU) ins_fdcb0086(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 0,(iy+00),a
func (cpu *CPU) ins_fdcb0087(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 1,(iy+00),b
func (cpu *CPU) ins_fdcb0088(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 1,(iy+00),c
func (cpu *CPU) ins_fdcb0089(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 1,(iy+00),d
func (cpu *CPU) ins_fdcb008a(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 1,(iy+00),e
func (cpu *CPU) ins_fdcb008b(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 1,(iy+00),h
func (cpu *CPU) ins_fdcb008c(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 1,(iy+00),l
func (cpu *CPU) ins_fdcb008d(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 1,(iy+00)
func (cpu *CPU) ins_fdcb008e(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 1,(iy+00),a
func (cpu *CPU) ins_fdcb008f(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 2,(iy+00),b
func (cpu *CPU) ins_fdcb0090(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 2,(iy+00),c
func (cpu *CPU) ins_fdcb0091(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 2,(iy+00),d
func (cpu *CPU) ins_fdcb0092(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 2,(iy+00),e
func (cpu *CPU) ins_fdcb0093(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 2,(iy+00),h
func (cpu *CPU) ins_fdcb0094(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 2,(iy+00),l
func (cpu *CPU) ins_fdcb0095(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 2,(iy+00)
func (cpu *CPU) ins_fdcb0096(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 2,(iy+00),a
func (cpu *CPU) ins_fdcb0097(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 3,(iy+00),b
func (cpu *CPU) ins_fdcb0098(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 3,(iy+00),c
func (cpu *CPU) ins_fdcb0099(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 3,(iy+00),d
func (cpu *CPU) ins_fdcb009a(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 3,(iy+00),e
func (cpu *CPU) ins_fdcb009b(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 3,(iy+00),h
func (cpu *CPU) ins_fdcb009c(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 3,(iy+00),l
func (cpu *CPU) ins_fdcb009d(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 3,(iy+00)
func (cpu *CPU) ins_fdcb009e(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 3,(iy+00),a
func (cpu *CPU) ins_fdcb009f(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 4,(iy+00),b
func (cpu *CPU) ins_fdcb00a0(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 4,(iy+00),c
func (cpu *CPU) ins_fdcb00a1(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 4,(iy+00),d
func (cpu *CPU) ins_fdcb00a2(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 4,(iy+00),e
func (cpu *CPU) ins_fdcb00a3(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 4,(iy+00),h
func (cpu *CPU) ins_fdcb00a4(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 4,(iy+00),l
func (cpu *CPU) ins_fdcb00a5(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 4,(iy+00)
func (cpu *CPU) ins_fdcb00a6(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 4,(iy+00),a
func (cpu *CPU) ins_fdcb00a7(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 5,(iy+00),b
func (cpu *CPU) ins_fdcb00a8(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 5,(iy+00),c
func (cpu *CPU) ins_fdcb00a9(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 5,(iy+00),d
func (cpu *CPU) ins_fdcb00aa(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 5,(iy+00),e
func (cpu *CPU) ins_fdcb00ab(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 5,(iy+00),h
func (cpu *CPU) ins_fdcb00ac(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 5,(iy+00),l
func (cpu *CPU) ins_fdcb00ad(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 5,(iy+00)
func (cpu *CPU) ins_fdcb00ae(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 5,(iy+00),a
func (cpu *CPU) ins_fdcb00af(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 6,(iy+00),b
func (cpu *CPU) ins_fdcb00b0(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 6,(iy+00),c
func (cpu *CPU) ins_fdcb00b1(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 6,(iy+00),d
func (cpu *CPU) ins_fdcb00b2(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 6,(iy+00),e
func (cpu *CPU) ins_fdcb00b3(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 6,(iy+00),h
func (cpu *CPU) ins_fdcb00b4(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 6,(iy+00),l
func (cpu *CPU) ins_fdcb00b5(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 6,(iy+00)
func (cpu *CPU) ins_fdcb00b6(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 6,(iy+00),a
func (cpu *CPU) ins_fdcb00b7(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// res 7,(iy+00),b
func (cpu *CPU) ins_fdcb00b8(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// res 7,(iy+00),c
func (cpu *CPU) ins_fdcb00b9(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// res 7,(iy+00),d
func (cpu *CPU) ins_fdcb00ba(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// res 7,(iy+00),e
func (cpu *CPU) ins_fdcb00bb(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// res 7,(iy+00),h
func (cpu *CPU) ins_fdcb00bc(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// res 7,(iy+00),l
func (cpu *CPU) ins_fdcb00bd(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// res 7,(iy+00)
func (cpu *CPU) ins_fdcb00be(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	return 11
}

// res 7,(iy+00),a
func (cpu *CPU) ins_fdcb00bf(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) &^ (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 0,(iy+00),b
func (cpu *CPU) ins_fdcb00c0(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 0,(iy+00),c
func (cpu *CPU) ins_fdcb00c1(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 0,(iy+00),d
func (cpu *CPU) ins_fdcb00c2(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 0,(iy+00),e
func (cpu *CPU) ins_fdcb00c3(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 0,(iy+00),h
func (cpu *CPU) ins_fdcb00c4(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 0,(iy+00),l
func (cpu *CPU) ins_fdcb00c5(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 0,(iy+00)
func (cpu *CPU) ins_fdcb00c6(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 0,(iy+00),a
func (cpu *CPU) ins_fdcb00c7(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 0)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 1,(iy+00),b
func (cpu *CPU) ins_fdcb00c8(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 1,(iy+00),c
func (cpu *CPU) ins_fdcb00c9(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 1,(iy+00),d
func (cpu *CPU) ins_fdcb00ca(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 1,(iy+00),e
func (cpu *CPU) ins_fdcb00cb(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 1,(iy+00),h
func (cpu *CPU) ins_fdcb00cc(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 1,(iy+00),l
func (cpu *CPU) ins_fdcb00cd(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 1,(iy+00)
func (cpu *CPU) ins_fdcb00ce(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 1,(iy+00),a
func (cpu *CPU) ins_fdcb00cf(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 1)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 2,(iy+00),b
func (cpu *CPU) ins_fdcb00d0(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 2,(iy+00),c
func (cpu *CPU) ins_fdcb00d1(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 2,(iy+00),d
func (cpu *CPU) ins_fdcb00d2(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 2,(iy+00),e
func (cpu *CPU) ins_fdcb00d3(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 2,(iy+00),h
func (cpu *CPU) ins_fdcb00d4(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 2,(iy+00),l
func (cpu *CPU) ins_fdcb00d5(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 2,(iy+00)
func (cpu *CPU) ins_fdcb00d6(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 2,(iy+00),a
func (cpu *CPU) ins_fdcb00d7(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 2)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 3,(iy+00),b
func (cpu *CPU) ins_fdcb00d8(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 3,(iy+00),c
func (cpu *CPU) ins_fdcb00d9(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 3,(iy+00),d
func (cpu *CPU) ins_fdcb00da(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 3,(iy+00),e
func (cpu *CPU) ins_fdcb00db(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 3,(iy+00),h
func (cpu *CPU) ins_fdcb00dc(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 3,(iy+00),l
func (cpu *CPU) ins_fdcb00dd(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 3,(iy+00)
func (cpu *CPU) ins_fdcb00de(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 3,(iy+00),a
func (cpu *CPU) ins_fdcb00df(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 3)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 4,(iy+00),b
func (cpu *CPU) ins_fdcb00e0(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 4,(iy+00),c
func (cpu *CPU) ins_fdcb00e1(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 4,(iy+00),d
func (cpu *CPU) ins_fdcb00e2(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 4,(iy+00),e
func (cpu *CPU) ins_fdcb00e3(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 4,(iy+00),h
func (cpu *CPU) ins_fdcb00e4(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 4,(iy+00),l
func (cpu *CPU) ins_fdcb00e5(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 4,(iy+00)
func (cpu *CPU) ins_fdcb00e6(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 4,(iy+00),a
func (cpu *CPU) ins_fdcb00e7(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 4)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 5,(iy+00),b
func (cpu *CPU) ins_fdcb00e8(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 5,(iy+00),c
func (cpu *CPU) ins_fdcb00e9(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 5,(iy+00),d
func (cpu *CPU) ins_fdcb00ea(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 5,(iy+00),e
func (cpu *CPU) ins_fdcb00eb(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 5,(iy+00),h
func (cpu *CPU) ins_fdcb00ec(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 5,(iy+00),l
func (cpu *CPU) ins_fdcb00ed(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 5,(iy+00)
func (cpu *CPU) ins_fdcb00ee(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 5,(iy+00),a
func (cpu *CPU) ins_fdcb00ef(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 5)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 6,(iy+00),b
func (cpu *CPU) ins_fdcb00f0(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 6,(iy+00),c
func (cpu *CPU) ins_fdcb00f1(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 6,(iy+00),d
func (cpu *CPU) ins_fdcb00f2(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 6,(iy+00),e
func (cpu *CPU) ins_fdcb00f3(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 6,(iy+00),h
func (cpu *CPU) ins_fdcb00f4(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 6,(iy+00),l
func (cpu *CPU) ins_fdcb00f5(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 6,(iy+00)
func (cpu *CPU) ins_fdcb00f6(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 6,(iy+00),a
func (cpu *CPU) ins_fdcb00f7(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 6)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}

// set 7,(iy+00),b
func (cpu *CPU) ins_fdcb00f8(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.B = val
	return 11
}

// set 7,(iy+00),c
func (cpu *CPU) ins_fdcb00f9(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.C = val
	return 11
}

// set 7,(iy+00),d
func (cpu *CPU) ins_fdcb00fa(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.D = val
	return 11
}

// set 7,(iy+00),e
func (cpu *CPU) ins_fdcb00fb(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.E = val
	return 11
}

// set 7,(iy+00),h
func (cpu *CPU) ins_fdcb00fc(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.H = val
	return 11
}

// set 7,(iy+00),l
func (cpu *CPU) ins_fdcb00fd(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.L = val
	return 11
}

// set 7,(iy+00)
func (cpu *CPU) ins_fdcb00fe(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	return 11
}

// set 7,(iy+00),a
func (cpu *CPU) ins_fdcb00ff(d uint8) int {
	n := cpu.IY + offset16(d)
	val := cpu.mem.Rd8(n) | (1 << 7)
	cpu.mem.Wr8(n, val)
	cpu.A = val
	return 11
}
