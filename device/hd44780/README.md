# HD44780 LCD Emulator

Package `hd44780` emulates the Hitachi HD44780U dot-matrix LCD controller
as used in the TEC-1G single-board computer.

## Port map (TEC-1G)

| Port  | Direction | Function                                     |
|-------|-----------|----------------------------------------------|
| `$04` | OUT       | Write Command (RS=0, RW=0)                   |
| `$04` | IN        | Read Command  (RS=0, RW=1) — busy flag bit 7, address counter bits 6:0 |
| `$84` | OUT       | Write Data    (RS=1, RW=0)                   |
| `$84` | IN        | Read Data     (RS=1, RW=1)                   |

## Conformance test suite

`asm/hd44780_test.asm` is a Z80 assembly conformance suite that exercises
every HD44780U instruction against the **real Hitachi datasheet behaviour**.
It is assembled alongside `asm/segments.asm` via `cd asm && make`.

The suite documents correct behaviour; the emulator is not yet fully
conformant.  Each known gap is listed below.

## Known conformance gaps

The items below are tested by `asm/hd44780_test.asm`.  A `test_fail` halt
(offending value in A) is the expected outcome for each gap when run against
the current emulator.

### 1. Busy flag always 0 (`ReadCommand`)

`ReadCommand` hard-codes the busy flag (bit 7) to 0.  A real HD44780 sets
bit 7 for the duration of command execution (typically 37 µs–1.52 ms
depending on the command).

**File:** `lcd.go` — `ReadCommand()`, comment: *"the busy flag is == 0"*

### 2. Address Counter not incremented on `ReadData`

`ddramRead` returns the current DDRAM byte but does not advance `ddAddr`.
On real hardware every `ReadData` auto-increments (or decrements) the
address counter and pre-fetches the next byte into the output register,
exactly as `WriteData` does.

**File:** `lcd.go` — `ddramRead()`

### 3. Dummy-read-after-address-set behaviour not modelled

After a Set DDRAM / CGRAM Address command the real device pre-fetches the
data at the new address into its output register.  The first `ReadData`
that follows returns that pre-fetched value and simultaneously initiates
the next pre-fetch.  This output-register pipeline is not implemented.

### 4. CGRAM read/write are stubs

`cgramWrite` and `cgramRead` are stubs; CGRAM writes are silently discarded
and reads always return 0.  Custom glyph programming and readback therefore
do not work.

**File:** `lcd.go` — `cgramWrite()`, `cgramRead()`

### 5. Entry-mode display shift (S=1) is a TODO

When the S bit of the Entry Mode Set command is 1, each `WriteData` should
shift the visible display by one position.  This is currently unimplemented.

**File:** `lcd.go` — `WriteCommand`, branch `cmd&cmdEntryMode != 0`,
log message: *"TODO: entry mode shift"*

### 6. Cursor/Display Shift instruction not implemented

The Cursor/Display Shift command (`$10`–`$1F`) is decoded but only logged;
neither the cursor position (AC) nor the scroll offset is updated.

**File:** `lcd.go` — `WriteCommand`, branch `cmd&cmdShift != 0`,
log message: *"shift"*

### 7. `New()` only supports 4-row displays

`New` returns an error for any display mode whose row count is not 4.  The
HD44780U supports 1-row and 2-row configurations as well.

**File:** `lcd.go` — `New()`, comment: *"TODO, rows != 4"*
