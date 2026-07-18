;-----------------------------------------------------------------------------
; HD44780U LCD Controller Conformance Test Suite
;
; Written against the Hitachi HD44780U datasheet (real hardware behaviour).
; It is EXPECTED that the current emulator does NOT pass every test —
; each failure documents a conformance gap to fix in device/hd44780/lcd.go.
;
; Use test_fail as a breakpoint; the offending value is left in A.
; Execution reaches tests_done (HALT) only when all assertions pass.
;
; TEC-1G port map (cmd/tec1g/io.go):
;   LCD_CMD  equ $04  RS=0  out=WriteCommand  in=ReadCommand (BF b7 | AC b6..0)
;   LCD_DAT  equ $84  RS=1  out=WriteData     in=ReadData
;
; Origin $0800, SP $0fd0  — consistent with asm/segments.asm
;
; See device/hd44780/README.md for the full list of emulator conformance gaps.
;-----------------------------------------------------------------------------

LCD_CMD  equ  $04          ; LCD command port (RS=0)
LCD_DAT  equ  $84          ; LCD data port    (RS=1)

; HD44780 command byte bases
CMD_CLEAR     equ  $01     ; Clear Display
CMD_HOME      equ  $02     ; Return Home
CMD_ENTRY     equ  $04     ; Entry Mode Set       (| I/D<<1 | S)
CMD_DISPLAY   equ  $08     ; Display On/Off       (| D<<2 | C<<1 | B)
CMD_SHIFT     equ  $10     ; Cursor/Display Shift (| S_C<<3 | R_L<<2)
CMD_FUNCTION  equ  $20     ; Function Set         (| DL<<4 | N<<3 | F<<2)
CMD_SET_CGRAM equ  $40     ; Set CGRAM Address    (6-bit addr)
CMD_SET_DDRAM equ  $80     ; Set DDRAM Address    (7-bit addr)

; DDRAM row start addresses (20-column, 4-row display — see lcd.go New())
ROW0  equ  $00
ROW1  equ  $40
ROW2  equ  $14
ROW3  equ  $54

; Busy flag
BF_MASK    equ  $80
; Bounded busy-wait retry count — prevents infinite loop on broken emulators
BF_RETRIES equ  $ff

;=============================================================================

TEST_PORT equ $08

test_entry MACRO arg1
    ld a,arg1
    out (TEST_PORT),a
    ENDM

;=============================================================================
; ENTRY POINT
;=============================================================================
         org  $4000

         ld   sp,$6fd0

;=============================================================================
; POWER-ON INITIALISATION  (HD44780U datasheet §Initializing by Instruction)
;
; 8-bit interface sequence (p.45):
;   1.  Wait >40 ms after VCC rises to 4.5 V
;   2.  Write Function Set ($30); wait >4.1 ms
;   3.  Write Function Set ($30); wait >100 µs
;   4.  Write Function Set ($30)  — internal reset now complete
; The busy flag MUST NOT be used during writes 2–4.
;=============================================================================
         call delay_50ms            ; >40 ms VCC stabilisation

         ld   a,CMD_FUNCTION|$10    ; $30  DL=1 (8-bit); N/F irrelevant here
         out  (LCD_CMD),a
         call delay_5ms             ; >4.1 ms

         ld   a,CMD_FUNCTION|$10    ; $30
         out  (LCD_CMD),a
         call delay_150us           ; >100 µs

         ld   a,CMD_FUNCTION|$10    ; $30 — internal reset complete
         out  (LCD_CMD),a

; Busy flag may now be used for all subsequent commands.

;=============================================================================
; TEST 1: Function Set — 8-bit bus, 2-line, 5×8 font
; Datasheet: DL=1, N=1, F=0 → $38
;=============================================================================
test_entry 1
;t_func_8b2l:
         ld   a,CMD_FUNCTION|$18    ; $38
         call wr_cmd

;=============================================================================
; TEST 2: Display On/Off Control — display off
; Datasheet: D=0, C=0, B=0 → $08
;=============================================================================
test_entry 2
;t_disp_off:
         ld   a,CMD_DISPLAY|$00     ; $08
         call wr_cmd

;=============================================================================
; TEST 3: Clear Display
; Fills entire DDRAM with $20 (space), sets AC=0, sets I/D=1 (increment).
; Real execution time ~1.52 ms.
;=============================================================================
test_entry 3
;t_clear:
         ld   a,CMD_CLEAR           ; $01
         call wr_cmd
         call delay_2ms             ; extra guard for emulators that ignore BF

;=============================================================================
; TEST 4: Return Home
; Sets AC=0, resets display shift; DDRAM content is unchanged.
; Real execution time ~1.52 ms.
; Verify: ReadCommand AC bits[6:0] = 0 after settling.
;=============================================================================
test_entry 4
;t_home:
         ; Write sentinel away from position 0 so Home is observable
         ld   a,CMD_SET_DDRAM|$06   ; position 6 on row 0
         call wr_cmd
         ld   a,'X'
         call wr_dat
         ; Issue Return Home
         ld   a,CMD_HOME            ; $02
         call wr_cmd
         call delay_2ms             ; >1.52 ms for Home to complete
         in   a,(LCD_CMD)
         and  $7f                   ; mask off BF
         cp   $00                   ; AC must be 0
         jr   z,home_ok
         jp   test_fail
home_ok:

;=============================================================================
; TEST 5: Entry Mode Set — increment, no display shift  (I/D=1, S=0)
; Datasheet: $06
;=============================================================================
test_entry 5
;t_entry_inc:
         ld   a,CMD_ENTRY|$02       ; $06
         call wr_cmd

;=============================================================================
; TEST 6: Display On/Off Control — display on, cursor on, blink on
; Datasheet: D=1, C=1, B=1 → $0F
;=============================================================================
test_entry 6
;t_disp_on:
         ld   a,CMD_DISPLAY|$07     ; $0F
         call wr_cmd

;=============================================================================
; TEST 7: Set DDRAM Address + Write Data on all four row bases
; Also verifies ReadCommand returns the correct AC after each address set.
;
; Row addresses for a 20×4 display: $00, $40, $14, $54
;=============================================================================
; --- Row 0 ($00) ---
test_entry 7
;t_ddram_r0:
         ld   a,CMD_SET_DDRAM|ROW0  ; $80
         call wr_cmd
         in   a,(LCD_CMD)
         and  $7f
         cp   ROW0                  ; AC should equal ROW0
         jr   z,ddram_r0_ok
         jp   test_fail
ddram_r0_ok:
         ld   hl,msg_r0
         call print_str

; --- Row 1 ($40) ---
;t_ddram_r1:
         ld   a,CMD_SET_DDRAM|ROW1  ; $c0
         call wr_cmd
         in   a,(LCD_CMD)
         and  $7f
         cp   ROW1
         jr   z,ddram_r1_ok
         jp   test_fail
ddram_r1_ok:
         ld   hl,msg_r1
         call print_str

; --- Row 2 ($14) ---
;t_ddram_r2:
         ld   a,CMD_SET_DDRAM|ROW2  ; $94
         call wr_cmd
         in   a,(LCD_CMD)
         and  $7f
         cp   ROW2
         jr   z,ddram_r2_ok
         jp   test_fail
ddram_r2_ok:
         ld   hl,msg_r2
         call print_str

; --- Row 3 ($54) ---
;t_ddram_r3:
         ld   a,CMD_SET_DDRAM|ROW3  ; $d4
         call wr_cmd
         in   a,(LCD_CMD)
         and  $7f
         cp   ROW3
         jr   z,ddram_r3_ok
         jp   test_fail
ddram_r3_ok:
         ld   hl,msg_r3
         call print_str

;=============================================================================
; TEST 8: Read Data — dummy read then real read after Set DDRAM Address
;
; Datasheet: after Set DDRAM Address, the output register is pre-loaded with
; the data at that address.  The first ReadData returns that pre-loaded value
; (treating it as a "dummy" to load the latch) and auto-increments AC by 1;
; the second ReadData returns the data at AC+1.
;
; This test verifies that AC auto-increments on ReadData (a known gap in the
; current emulator, which does not increment AC on reads).
;
; Expected (real hardware):
;   dummy read  → 'R' (msg_r0[0]), AC becomes 1
;   real read   → 'o' (msg_r0[1])
;
; Expected (current emulator — conformance gap):
;   both reads  → 'R' (AC never incremented), cp 'o' FAILS
;=============================================================================
test_entry 8
;t_read_data:
         ld   a,CMD_SET_DDRAM|ROW0  ; set AC = 0
         call wr_cmd
         in   a,(LCD_DAT)           ; dummy read: loads output latch, AC→1 (real HW)
         in   a,(LCD_DAT)           ; real read: should be msg_r0[1] = 'o' (real HW)
         cp   'o'
         jr   z,rdata_ok
         jp   test_fail
rdata_ok:

;=============================================================================
; TEST 9: Entry Mode Set — decrement  (I/D=0, S=0)
; Write three bytes starting at $0A; with decrement the address walks
; $0A → $09 → $08 → $07 (AC after three writes = $07).
;=============================================================================
test_entry 9
;t_entry_dec:
         ld   a,CMD_ENTRY|$00       ; $04  I/D=0, S=0
         call wr_cmd
         ld   a,CMD_SET_DDRAM|$0a   ; start at $0A
         call wr_cmd
         ld   a,'3'
         call wr_dat
         ld   a,'2'
         call wr_dat
         ld   a,'1'
         call wr_dat
         in   a,(LCD_CMD)
         and  $7f
         cp   $07                   ; AC must be $0A - 3 = $07
         jr   z,entry_dec_ok
         jp   test_fail
entry_dec_ok:
         ; Restore standard increment mode
         ld   a,CMD_ENTRY|$02       ; $06
         call wr_cmd

;=============================================================================
; TEST 10: Entry Mode — display shift on each write  (I/D=1, S=1)
; Datasheet: when S=1 the display shifts on each Write Data.
; The current emulator logs "TODO: entry mode shift" — conformance gap.
;=============================================================================
test_entry 10
;t_entry_shift:
         ld   a,CMD_ENTRY|$03       ; $07  I/D=1, S=1
         call wr_cmd
         ld   a,CMD_SET_DDRAM|ROW0
         call wr_cmd
         ld   a,'S'
         call wr_dat                ; should trigger one display shift
         ; Restore standard entry mode
         ld   a,CMD_ENTRY|$02       ; $06
         call wr_cmd

;=============================================================================
; TEST 11: Cursor/Display Shift — all four S/C × R/L combinations
;
; CMD_SHIFT bit layout:  b3=S/C (0=cursor, 1=display), b2=R/L (0=left, 1=right)
;   $10  cursor move left    $14  cursor move right
;   $18  display shift left  $1C  display shift right
;
; The current emulator only logs "shift" without updating AC or scroll —
; conformance gap for all four variants.
;=============================================================================
test_entry 11

         ld   a,CMD_SET_DDRAM|$05   ; place cursor at position $05
         call wr_cmd

; (i) Cursor shift left: AC should decrease from $05 to $04
;t_shift_cl:
         ld   a,CMD_SHIFT|$00       ; $10
         call wr_cmd
         in   a,(LCD_CMD)
         and  $7f
         cp   $04
         jr   z,shift_cl_ok
         jp   test_fail
shift_cl_ok:

; (ii) Cursor shift right: AC should return to $05
;t_shift_cr:
         ld   a,CMD_SHIFT|$04       ; $14
         call wr_cmd
         in   a,(LCD_CMD)
         and  $7f
         cp   $05
         jr   z,shift_cr_ok
         jp   test_fail
shift_cr_ok:

; (iii) Display shift left — AC should remain $05 (cursor position unchanged)
;t_shift_dl:
         ld   a,CMD_SHIFT|$08       ; $18
         call wr_cmd

; (iv) Display shift right
;t_shift_dr:
         ld   a,CMD_SHIFT|$0c       ; $1C
         call wr_cmd

;=============================================================================
; TEST 12: Set CGRAM Address + Write 8 custom glyphs (64 bytes)
; Custom character n uses CGRAM addresses n*8 .. n*8+7.
; Only bits[4:0] of each byte are significant (5 pixels wide).
;=============================================================================

test_entry 12
;t_cgram_write:
         ld   a,CMD_SET_CGRAM|$00   ; CGRAM address 0 (glyph 0, row 0)
         call wr_cmd
         ld   hl,cg_chars
         ld   b,64                  ; 8 glyphs × 8 rows
cgwr_lp:
         ld   a,(hl)
         call wr_dat
         inc  hl
         djnz cgwr_lp

;=============================================================================
; TEST 13: Read CGRAM back and verify all 64 bytes
; Includes dummy ReadData after Set CGRAM Address to load the output latch.
;
; cgramRead is a stub in the current emulator (returns 0) — expect failure.
;
; Expected (real hardware):
;   dummy read    → cg_chars[0], AC→1
;   62 more reads → cg_chars[1..63]   (total 63 data reads after dummy)
;
; NOTE: byte 0 of cg_chars is consumed by the dummy read and not re-verified
; in the loop; only bytes 1..63 are compared.  On the current emulator all
; reads return 0 so the first comparison fails at cg_chars[1].
;=============================================================================

test_entry 13
;t_cgram_read:
         ld   a,CMD_SET_CGRAM|$00   ; back to CGRAM start
         call wr_cmd
         in   a,(LCD_DAT)           ; dummy read: loads output latch, AC→1
         ld   hl,cg_chars
         inc  hl                    ; advance past cg_chars[0]: already returned by the dummy read above
         ld   b,63                  ; verify bytes 1..63
cgrd_lp:
         in   a,(LCD_DAT)
         cp   (hl)
         jr   z,cgrd_ok
         jp   test_fail
cgrd_ok:
         inc  hl
         djnz cgrd_lp

;=============================================================================
; TEST 14: Function Set — additional DL/N/F combinations
; (a) 8-bit, 1-line, 5×8 font   DL=1 N=0 F=0 → $30
; (b) 8-bit, 1-line, 5×10 font  DL=1 N=0 F=1 → $34
; (c) 8-bit, 2-line, 5×8 font   DL=1 N=1 F=0 → $38  (restore)
;
; Note: 4-bit mode (DL=0) would break Z80 I/O — not tested here.
;=============================================================================

test_entry 14
;t_func_1l_5x8:
         ld   a,CMD_FUNCTION|$10    ; $30  1-line, 5×8
         call wr_cmd

;t_func_1l_5x10:
         ld   a,CMD_FUNCTION|$14    ; $34  1-line, 5×10
         call wr_cmd

;t_func_restore:
         ld   a,CMD_FUNCTION|$18    ; $38  2-line, 5×8 (restore)
         call wr_cmd

;=============================================================================
; TEST 15: Display On/Off Control — all eight D/C/B permutations
;=============================================================================

test_entry 15
         ld   a,CMD_DISPLAY|$00     ; $08  D=0 C=0 B=0 — all off
         call wr_cmd
         ld   a,CMD_DISPLAY|$01     ; $09  B only
         call wr_cmd
         ld   a,CMD_DISPLAY|$02     ; $0A  C only
         call wr_cmd
         ld   a,CMD_DISPLAY|$03     ; $0B  C+B
         call wr_cmd
         ld   a,CMD_DISPLAY|$04     ; $0C  D only
         call wr_cmd
         ld   a,CMD_DISPLAY|$05     ; $0D  D+B
         call wr_cmd
         ld   a,CMD_DISPLAY|$06     ; $0E  D+C
         call wr_cmd
         ld   a,CMD_DISPLAY|$07     ; $0F  D+C+B — all on
         call wr_cmd

;=============================================================================
; TEST 16: Clear Display then verify DDRAM contains $20 (space)
; After Clear, DDRAM[0] should be $20.  Uses dummy-read idiom.
;=============================================================================

test_entry 16
;t_clear_verify:
         ld   a,CMD_CLEAR
         call wr_cmd
         call delay_2ms
         ld   a,CMD_SET_DDRAM|ROW0
         call wr_cmd
         in   a,(LCD_DAT)           ; dummy read (loads latch), AC→1 on real HW
         in   a,(LCD_DAT)           ; real read: DDRAM[1] should also be $20
         cp   $20
         jr   z,clrv_ok
         jp   test_fail
clrv_ok:

;=============================================================================
; TEST 17: Read Busy Flag — BF (bit 7) must be 0 in settled state
;=============================================================================

test_entry 17
;t_read_bf:
         call delay_2ms             ; ensure any lingering command is done
         in   a,(LCD_CMD)
         and  BF_MASK
         cp   $00                   ; BF must be 0
         jr   z,rbf_ok
         jp   test_fail
rbf_ok:

;=============================================================================
; ALL TESTS PASSED
;=============================================================================
test_entry 18
;tests_done:
         halt

;=============================================================================
; TEST FAILURE — offending value is in A; set a breakpoint here
;=============================================================================
test_fail:
         halt

;=============================================================================
; SUBROUTINES
;=============================================================================

;-----------------------------------------------------------------------------
; wr_cmd — write command byte in A to the LCD command port, then busy-wait.
; Modifies: A (clobbered by busy_wait's ReadCommand poll)
;-----------------------------------------------------------------------------
wr_cmd:
         out  (LCD_CMD),a
         call busy_wait
         ret

;-----------------------------------------------------------------------------
; busy_wait — poll BF (bit 7 of ReadCommand) with a bounded retry count.
; Exits as soon as BF=0 or after BF_RETRIES polls (avoids hang on emulators
; that never set BF).
; Preserves: BC  Modifies: A
;-----------------------------------------------------------------------------
busy_wait:
         push bc
         ld   b,BF_RETRIES
bw_lp:
         in   a,(LCD_CMD)
         and  BF_MASK
         jr   z,bw_done             ; BF=0 → LCD ready
         djnz bw_lp
bw_done:
         pop  bc
         ret

;-----------------------------------------------------------------------------
; wr_dat — write data byte in A to the LCD data port, then busy-wait.
; Preserves: A  Modifies: (A clobbered and then restored via push/pop AF)
;-----------------------------------------------------------------------------
wr_dat:
         push af
         out  (LCD_DAT),a
         call busy_wait
         pop  af
         ret

;-----------------------------------------------------------------------------
; print_str — write null-terminated string at (HL) to current DDRAM position
; Returns when the zero terminator is found.
; Modifies: A, HL
;-----------------------------------------------------------------------------
print_str:
         ld   a,(hl)
         or   a
         ret  z
         call wr_dat
         inc  hl
         jr   print_str

;-----------------------------------------------------------------------------
; Delay routines — cycle-counted busy loops, calibrated for ~4 MHz Z80 clock.
;
; Inner loop body: dec de (6T) + ld a,e (4T) + or d (4T) + jp nz (10T) = 24T
; At 4 MHz: 1 iteration ≈ 6 µs
;
; delay_50ms  : DE=$20D0  (~8400 iters × 6 µs ≈ 50 ms)  — power-on VCC wait
; delay_5ms   : DE=$0348  (~ 840 iters × 6 µs ≈  5 ms)  — init inter-write
; delay_2ms   : DE=$0154  (~ 340 iters × 6 µs ≈  2 ms)  — Clear / Home guard
; delay_150us : DE=$0019  (~  25 iters × 6 µs ≈150 µs)  — init inter-write
;-----------------------------------------------------------------------------

delay_50ms:
         ld   de,$20d0
d50_lp:
         dec  de
         ld   a,e
         or   d
         jp   nz,d50_lp
         ret

delay_5ms:
         ld   de,$0348
d5_lp:
         dec  de
         ld   a,e
         or   d
         jp   nz,d5_lp
         ret

delay_2ms:
         ld   de,$0154
d2_lp:
         dec  de
         ld   a,e
         or   d
         jp   nz,d2_lp
         ret

delay_150us:
         ld   de,$0019
d150_lp:
         dec  de
         ld   a,e
         or   d
         jp   nz,d150_lp
         ret

;=============================================================================
; DATA
;=============================================================================

; Row messages — null-terminated, ≤20 chars each (20-column display)
msg_r0:  defb "Row0: HD44780 test",0
msg_r1:  defb "Row1: line 1 ok",0
msg_r2:  defb "Row2: addr $14",0
msg_r3:  defb "Row3: addr $54",0

; Custom glyph data — 8 glyphs × 8 rows; only bits[4:0] used per row.
; Glyph 0: solid block
; Glyph 1: checkerboard
; Glyph 2: left-pointing arrow
; Glyph 3: right-pointing arrow
; Glyph 4: up arrow
; Glyph 5: down arrow
; Glyph 6: top border line
; Glyph 7: bottom border line
cg_chars:
         defb $1f,$1f,$1f,$1f,$1f,$1f,$1f,$1f  ; glyph 0: solid block
         defb $15,$0a,$15,$0a,$15,$0a,$15,$0a  ; glyph 1: checkerboard
         defb $04,$0c,$1f,$0c,$04,$00,$00,$00  ; glyph 2: left arrow
         defb $04,$06,$1f,$06,$04,$00,$00,$00  ; glyph 3: right arrow
         defb $04,$0e,$15,$04,$04,$04,$00,$00  ; glyph 4: up arrow
         defb $00,$00,$04,$04,$15,$0e,$04,$00  ; glyph 5: down arrow
         defb $1f,$00,$00,$00,$00,$00,$00,$00  ; glyph 6: top border
         defb $00,$00,$00,$00,$00,$00,$00,$1f  ; glyph 7: bottom border
