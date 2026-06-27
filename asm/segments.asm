
org $0800

	ld sp,0x0fd0
	ld a, $20
	out ($01), a

loop0:
	ld a, $01
	call delay
	ld a, $08
	call delay
	ld a, $20
	call delay
	ld a, $80
	call delay
	ld a, $40
	call delay
	ld a, $02
	call delay
	jp loop0

delay:
	out ($02), a
	ld de, $07ff
loop1:
	dec de
	ld a, e
	or d
	jp nz, loop1
	ret
