//-----------------------------------------------------------------------------
/*

Serial Protocol Tests for DS1302 RTC Emulation

These tests exercise the bit-level serial protocol implemented in Write/Read.
They are in the same package so unexported fields and helpers are accessible.

*/
//-----------------------------------------------------------------------------

package ds1302

import (
	"testing"
)

//-----------------------------------------------------------------------------
// test helpers

// newTestRTC creates an enabled RTC suitable for serial protocol testing.
// The caller must call Close() when done to stop the background goroutine.
func newTestRTC(t *testing.T) *RTC {
	t.Helper()
	rtc, err := New(&Config{
		Enable:   true,
		BaseYear: 2000,
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return rtc
}

// ceDeassert deasserts chip enable, which resets the device serial state.
func ceDeassert(rtc *RTC) {
	rtc.Write(false, false, false)
}

// clockCmd clocks 8 command bits into the device (LSB first).
// CE must be asserted via each Write call (chipEnable=true throughout).
// CLK starts low and ends HIGH — the trailing CLK-low is intentionally
// omitted: in byteReadState a falling edge would shift out the first data
// bit prematurely before clockByteOut can sample it.
func clockCmd(rtc *RTC, cmd byte) {
	for i := 0; i < 8; i++ {
		bit := (cmd>>i)&1 != 0
		rtc.Write(true, false, bit) // CLK low: set input bit
		rtc.Write(true, true, bit)  // CLK high: rising edge latches the bit
	}
	// CLK is now HIGH; caller is responsible for the next transition.
}

// clockByteIn clocks 8 data bits into the device (LSB first) during a write.
// Device must be in byteWriteState. CLK may start high or low (falling edges
// are ignored in byteWriteState). CLK ends low.
func clockByteIn(rtc *RTC, data byte) {
	for i := 0; i < 8; i++ {
		bit := (data>>i)&1 != 0
		rtc.Write(true, false, bit) // CLK low: set input (falling edge ignored)
		rtc.Write(true, true, bit)  // CLK high: rising edge latches the bit
	}
	rtc.Write(true, false, false) // CLK low to complete
}

// clockByteOut clocks 8 data bits out of the device (LSB first) during a read.
// Device must be in byteReadState; CLK must be HIGH at entry (left there by
// clockCmd). Each bit is sampled after the falling edge via Read().
//
// Implementation note: for the last bit of a non-burst read, the implementation
// calls reset() inside Write() after setting rtc.out, which clears rtc.out
// before Read() can sample it. The bit is instead read directly from rtc.data
// (which holds original_byte >> 7 at that point) before triggering the final
// falling edge.
func clockByteOut(rtc *RTC) byte {
	var result byte
	for i := 0; i < 8; i++ {
		if i == 7 && !rtc.burst {
			// Non-burst last bit: after 7 falling edges rtc.data == original >> 7,
			// so bit 0 is bit 7 of the original byte. Read it before the reset.
			if rtc.data&1 != 0 {
				result |= 1 << 7
			}
			rtc.Write(true, false, false) // 8th falling edge: sets out, triggers reset
			return result
		}
		rtc.Write(true, false, false) // CLK high→low: falling edge shifts out bit i
		if rtc.Read() {
			result |= 1 << i
		}
		rtc.Write(true, true, false) // CLK low→high: rising edge (ignored in byteReadState)
	}
	// CLK is HIGH; consistent for the next clockByteOut call in burst mode.
	return result
}

// writeReg performs a complete single write transaction:
// clock in cmd then data byte, then deassert CE.
func writeReg(rtc *RTC, cmd, data byte) {
	clockCmd(rtc, cmd)
	clockByteIn(rtc, data)
	ceDeassert(rtc)
}

// readReg performs a complete single read transaction:
// clock in cmd, read out data byte, then deassert CE.
func readReg(rtc *RTC, cmd byte) byte {
	clockCmd(rtc, cmd)
	data := clockByteOut(rtc)
	ceDeassert(rtc)
	return data
}

// cmdClockWrite returns the command byte for writing clock register adr.
func cmdClockWrite(adr int) byte {
	return topBit | byte(adr<<1)
}

// cmdClockRead returns the command byte for reading clock register adr.
func cmdClockRead(adr int) byte {
	return topBit | byte(adr<<1) | rwBit
}

// cmdRAMWrite returns the command byte for writing RAM register adr.
func cmdRAMWrite(adr int) byte {
	return topBit | rcBit | byte(adr<<1)
}

// cmdRAMRead returns the command byte for reading RAM register adr.
func cmdRAMRead(adr int) byte {
	return topBit | rcBit | byte(adr<<1) | rwBit
}

// disableWriteProtect clears the write-protect flag via the serial bus.
// Writing to the clockWriteProtect register is always permitted even when
// write-protect is currently enabled.
func disableWriteProtect(rtc *RTC) {
	writeReg(rtc, cmdClockWrite(clockWriteProtect), 0x00)
}

//-----------------------------------------------------------------------------

// Test_SerialRamReadWrite verifies single-byte RAM write/read-back at several
// addresses including boundaries (0 and 30). RAM never changes on its own so
// no clock-halting is required.
func Test_SerialRamReadWrite(t *testing.T) {
	rtc := newTestRTC(t)
	defer rtc.Close()

	disableWriteProtect(rtc)

	tests := []struct {
		adr  int
		data byte
	}{
		{0, 0xAA},  // lower boundary
		{1, 0x55},  // adjacent to lower boundary
		{15, 0x42}, // middle address
		{29, 0x80}, // near upper boundary, MSB only
		{30, 0xFF}, // upper boundary, all bits set
		{0, 0x00},  // zero value
		{30, 0x01}, // LSB only
	}

	for i, tc := range tests {
		writeReg(rtc, cmdRAMWrite(tc.adr), tc.data)
		got := readReg(rtc, cmdRAMRead(tc.adr))
		if got != tc.data {
			t.Fatalf("case %d: RAM[%d]: wrote 0x%02x, read back 0x%02x", i, tc.adr, tc.data, got)
		}
	}
}

//-----------------------------------------------------------------------------

// Test_SerialClockReadWrite verifies clock register write/read-back.
// The clock is halted first so time does not advance during the test.
func Test_SerialClockReadWrite(t *testing.T) {
	rtc := newTestRTC(t)
	defer rtc.Close()

	// Disable write protect before any clock writes.
	disableWriteProtect(rtc)

	// Halt the clock at 30 seconds to keep time stable.
	writeReg(rtc, cmdClockWrite(clockSecond), clockHalted|intToBcd(30))

	// Verify halt bit and seconds value round-trip.
	sec := readReg(rtc, cmdClockRead(clockSecond))
	if sec&clockHalted == 0 {
		t.Fatalf("halt bit not set after writing halt: got 0x%02x", sec)
	}
	if bcdToInt(sec&^clockHalted) != 30 {
		t.Fatalf("seconds: expected 30, got %d", bcdToInt(sec&^clockHalted))
	}

	// Each case writes rawVal to a clock register and checks the decoded read-back.
	tests := []struct {
		name     string
		writeCmd byte
		readCmd  byte
		rawVal   byte
		expected int
		decode   func(byte) int
	}{
		{
			"minute",
			cmdClockWrite(clockMinute), cmdClockRead(clockMinute),
			intToBcd(45), 45,
			func(b byte) int { return bcdToInt(b) },
		},
		{
			"hour24",
			cmdClockWrite(clockHour), cmdClockRead(clockHour),
			encodeHour(14, false), 14,
			func(b byte) int { return decodeHour(b) },
		},
		{
			"hour12-pm",
			cmdClockWrite(clockHour), cmdClockRead(clockHour),
			encodeHour(22, true), 22,
			func(b byte) int { return decodeHour(b) },
		},
		{
			"hour12-midnight",
			cmdClockWrite(clockHour), cmdClockRead(clockHour),
			encodeHour(0, true), 0,
			func(b byte) int { return decodeHour(b) },
		},
		{
			"dayOfMonth",
			cmdClockWrite(clockDayOfMonth), cmdClockRead(clockDayOfMonth),
			intToBcd(15), 15,
			func(b byte) int { return bcdToInt(b) },
		},
		{
			"month",
			cmdClockWrite(clockMonthOfYear), cmdClockRead(clockMonthOfYear),
			intToBcd(6), 6,
			func(b byte) int { return bcdToInt(b) },
		},
		{
			"year",
			cmdClockWrite(clockYear), cmdClockRead(clockYear),
			intToBcd(24), 24,
			func(b byte) int { return bcdToInt(b) },
		},
	}

	for i, tc := range tests {
		writeReg(rtc, tc.writeCmd, tc.rawVal)
		got := readReg(rtc, tc.readCmd)
		if tc.decode(got) != tc.expected {
			t.Fatalf("case %d (%s): expected %d, got %d (raw 0x%02x)", i, tc.name, tc.expected, tc.decode(got), got)
		}
	}
}

//-----------------------------------------------------------------------------

// Test_SerialWriteProtect verifies that writes are blocked when write-protect
// is enabled (the power-on default) and succeed after it is disabled.
func Test_SerialWriteProtect(t *testing.T) {
	rtc := newTestRTC(t)
	defer rtc.Close()

	// Write-protect is enabled at power-on. A RAM write must be silently
	// ignored and the register must remain at its zero-initialised value.
	writeReg(rtc, cmdRAMWrite(5), 0xAB)
	got := readReg(rtc, cmdRAMRead(5))
	if got != 0x00 {
		t.Fatalf("write-protect on: expected RAM[5]=0x00, got 0x%02x", got)
	}

	// The clockWriteProtect register must report bit 7 set.
	wp := readReg(rtc, cmdClockRead(clockWriteProtect))
	if wp&writeProtectEnabled == 0 {
		t.Fatalf("write-protect register: expected bit 7 set, got 0x%02x", wp)
	}

	// Disable write-protect and confirm the write now succeeds.
	disableWriteProtect(rtc)
	writeReg(rtc, cmdRAMWrite(5), 0xAB)
	got = readReg(rtc, cmdRAMRead(5))
	if got != 0xAB {
		t.Fatalf("write-protect off: expected RAM[5]=0xAB, got 0x%02x", got)
	}

	// clockWriteProtect register must now read back with bit 7 clear.
	wp = readReg(rtc, cmdClockRead(clockWriteProtect))
	if wp&writeProtectEnabled != 0 {
		t.Fatalf("write-protect disabled: expected bit 7 clear, got 0x%02x", wp)
	}
}

//-----------------------------------------------------------------------------

// Test_SerialBurstRAM verifies RAM burst mode: write all 31 registers in one
// CE transaction, then read them all back in another burst transaction.
func Test_SerialBurstRAM(t *testing.T) {
	rtc := newTestRTC(t)
	defer rtc.Close()

	disableWriteProtect(rtc)

	// Deterministic test pattern across all 31 RAM registers.
	var pattern [numRamRegisters]byte
	for i := range pattern {
		pattern[i] = byte(i*7 + 0x55)
	}

	// RAM burst write: topBit | rcBit | (burstAddress << 1) | write(0)
	const ramBurstWrite = topBit | rcBit | byte(burstAddress<<1)
	clockCmd(rtc, ramBurstWrite)
	for _, b := range pattern {
		clockByteIn(rtc, b)
	}
	ceDeassert(rtc)

	// RAM burst read: topBit | rcBit | (burstAddress << 1) | rwBit
	const ramBurstRead = topBit | rcBit | byte(burstAddress<<1) | rwBit
	clockCmd(rtc, ramBurstRead)
	for i := 0; i < numRamRegisters; i++ {
		got := clockByteOut(rtc)
		if got != pattern[i] {
			t.Fatalf("case %d: burst read: expected 0x%02x, got 0x%02x", i, pattern[i], got)
		}
	}
	ceDeassert(rtc)
}

//-----------------------------------------------------------------------------

// Test_SerialBadCommand verifies that a command byte with topBit (0x80) clear
// causes the device to reset without performing any operation.
func Test_SerialBadCommand(t *testing.T) {
	rtc := newTestRTC(t)
	defer rtc.Close()

	disableWriteProtect(rtc)

	// Pre-write a known value so we can confirm no spurious write occurred.
	writeReg(rtc, cmdRAMWrite(0), 0xBB)

	// Clock in a bad command: topBit is not set.
	clockCmd(rtc, 0x7E)
	if rtc.state != commandState {
		t.Fatalf("bad command: expected commandState, got %d", rtc.state)
	}
	ceDeassert(rtc)

	// Pre-written value must be unchanged (no spurious write).
	got := readReg(rtc, cmdRAMRead(0))
	if got != 0xBB {
		t.Fatalf("bad command: RAM[0] expected 0xBB, got 0x%02x", got)
	}

	// A subsequent well-formed transaction must still work correctly.
	writeReg(rtc, cmdRAMWrite(1), 0xCC)
	got = readReg(rtc, cmdRAMRead(1))
	if got != 0xCC {
		t.Fatalf("bad command recovery: RAM[1] expected 0xCC, got 0x%02x", got)
	}
}

//-----------------------------------------------------------------------------

// Test_SerialCEResetMidTransaction verifies that deasserting CE mid-command
// resets the device, and that a fresh transaction then works correctly.
func Test_SerialCEResetMidTransaction(t *testing.T) {
	rtc := newTestRTC(t)
	defer rtc.Close()

	disableWriteProtect(rtc)

	// Clock in 2 bits of a command byte, then deassert CE partway through.
	rtc.Write(true, false, true)
	rtc.Write(true, true, true)  // bit 0 latched
	rtc.Write(true, false, false)
	rtc.Write(true, true, false) // bit 1 latched

	ceDeassert(rtc)

	if rtc.state != commandState {
		t.Fatalf("CE mid-command: expected commandState, got %d", rtc.state)
	}
	if rtc.bits != 0 {
		t.Fatalf("CE mid-command: expected bits=0, got %d", rtc.bits)
	}

	// A fresh transaction after the reset must work correctly.
	writeReg(rtc, cmdRAMWrite(10), 0x99)
	got := readReg(rtc, cmdRAMRead(10))
	if got != 0x99 {
		t.Fatalf("CE mid-command recovery: RAM[10] expected 0x99, got 0x%02x", got)
	}
}

//-----------------------------------------------------------------------------
