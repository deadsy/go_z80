//-----------------------------------------------------------------------------
/*

ZexAll Testing

*/
//-----------------------------------------------------------------------------

package z80

import (
	"fmt"
	"os"
	"testing"
	"time"
)

//-----------------------------------------------------------------------------

// ram is a plain 64KB memory implementing the Memory interface.
type ram struct {
	b [0x10000]uint8
}

func (m *ram) Read8(a uint16) uint8     { return m.b[a] }
func (m *ram) Write8(a uint16, v uint8) { m.b[a] = v }
func (m *ram) Read16(a uint16) uint16 {
	return uint16(m.b[a]) | uint16(m.b[a+1])<<8
}
func (m *ram) Write16(a uint16, v uint16) {
	m.b[a] = uint8(v)
	m.b[a+1] = uint8(v >> 8)
}

//-----------------------------------------------------------------------------

// nullIO
type nullIO struct{}

func (nullIO) Read8(uint16) uint8   { return 0xff }
func (nullIO) Write8(uint16, uint8) {}

// nullBus
type nullBus struct{}

func (nullBus) ReadIV() uint8 { return 0xff }

//-----------------------------------------------------------------------------

const bdosEntry = uint16(0x0005)
const bdosExit = uint16(0x0000)
const loadAdr = uint16(0x0100)

// runCPM loads a CP/M .com at 0x0100 and runs it to completion,
// emulating the two BDOS calls zex needs. It returns the instruction count.
func runCPM(t *testing.T, path string) uint64 {

	// build and load the memory
	prog, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	mem := &ram{}
	copy(mem.b[0x0100:], prog)

	// Trap addresses:
	//  0x0000: warm boot / exit  -> HALT
	mem.Write8(bdosExit, 0x76) // HALT at exit vector
	//  0x0005: BDOS entry -> RET (we intercept before executing)
	mem.Write8(bdosEntry, 0xc9) // RET at BDOS entry (fallback)

	cpu := New(nullIO{}, mem, nullBus{})
	cpu.PC = loadAdr

	var count uint64
	for {
		// Intercept BDOS before the CALL target executes.
		if cpu.PC == bdosEntry {
			switch cpu.C {
			case 2: // C_WRITE: print char in E
				fmt.Printf("%c", rune(cpu.E))
			case 9: // C_WRITESTR: print $-terminated string at DE
				addr := cpu.get_de()

				for true {
					c := mem.Read8(addr)
					if c == '$' {
						break
					}
					fmt.Printf("%c", rune(c))
					addr += 1
				}
			}
			// return from BDOS
			cpu.PC = cpu.pop16()
			continue
		}
		if cpu.PC == 0x0000 || cpu.halt {
			break
		}
		if _, err := cpu.Run(); err != nil {
			t.Fatalf("run error: %v", err)
		}
		count++
	}
	return count
}

//-----------------------------------------------------------------------------

// go test -run Test_Zexdoc -v
func Test_Zexdoc(t *testing.T) {
	start := time.Now()
	n := runCPM(t, "../ext/zexall/zexdoc.com")
	elapsed := time.Since(start)
	t.Logf("\n%d instructions in %s (%.1f MIPS)",
		n, elapsed, float64(n)/elapsed.Seconds()/1e6)
}

// go test -run Test_Zexall -v
func Test_Zexall(t *testing.T) {
	start := time.Now()
	n := runCPM(t, "../ext/zexall/zexall.com")
	elapsed := time.Since(start)
	t.Logf("\n%d instructions in %s (%.1f MIPS)",
		n, elapsed, float64(n)/elapsed.Seconds()/1e6)
}

//-----------------------------------------------------------------------------
