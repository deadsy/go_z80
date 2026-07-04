//-----------------------------------------------------------------------------
/*

Matrix Keyboards

Note:

When a Z80 performs an IN instruction, the lower 8 bits typically specifies
the port number, but the upper 8 bits is set to a register value. This can
be used by matrix scanning code to simultaneously set a row value while reading
the resulting column value.

*/
//-----------------------------------------------------------------------------

package keyboard

import "fmt"

//-----------------------------------------------------------------------------
// key row values (the low bit selects the row)

const numRows = 8

const keyRow0 = byte(0xfe) // A8
const keyRow1 = byte(0xfd) // A9
const keyRow2 = byte(0xfb) // A10
const keyRow3 = byte(0xf7) // A11
const keyRow4 = byte(0xef) // A12
const keyRow5 = byte(0xdf) // A13
const keyRow6 = byte(0xbf) // A14
const keyRow7 = byte(0x7f) // A15

func rowToInt(row byte) (int, error) {
	var n int
	switch row {
	case keyRow0:
		n = 0
	case keyRow1:
		n = 1
	case keyRow2:
		n = 2
	case keyRow3:
		n = 3
	case keyRow4:
		n = 4
	case keyRow5:
		n = 5
	case keyRow6:
		n = 6
	case keyRow7:
		n = 7
	default:
		return 0, fmt.Errorf("bad row 0x%02x", row)

	}
	return n, nil
}

//-----------------------------------------------------------------------------
