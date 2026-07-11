//-----------------------------------------------------------------------------
/*

Jupiter ACE Matrix Keyboard

Maps the pressed ebiten key values into the row/col scan codes required by
the jupiter ace matrix keyboard scanning code.

# -----------------------------------------------------------------------------
# ;                          LOGICAL VIEW OF KEYBOARD
# ;
# ;         0     1     2     3     4 -Bits-  4     3     2     1     0
# ; PORT                                                                    PORT
# ;
# ; F7FE  [ 1 ] [ 2 ] [ 3 ] [ 4 ] [ 5 ]  |  [ 6 ] [ 7 ] [ 8 ] [ 9 ] [ 0 ]   EFFE
# ;  ^                                   |                                   v
# ; FBFE  [ Q ] [ W ] [ E ] [ R ] [ T ]  |  [ Y ] [ U ] [ I ] [ O ] [ P ]   DFFE
# ;  ^                                   |                                   v
# ; FDFE  [ A ] [ S ] [ D ] [ F ] [ G ]  |  [ H ] [ J ] [ K ] [ L ] [ ENT ] BFFE
# ;  ^                                   |                                   v
# ; FEFE  [SHI] [SYM] [ Z ] [ X ] [ C ]  |  [ V ] [ B ] [ N ] [ M ] [ SPC ] 7FFE
# ;  ^            v                                                ^         v
# ; Start         +------------>--------------------->-------------+        End
# ;

Note:

When a Z80 performs an IN instruction, the lower 8 bits typically specifies
the port number, but the upper 8 bits is set to a register value. This can
be used by matrix scanning code to simultaneously set a row value while reading
the resulting column value.

*/
//-----------------------------------------------------------------------------

package keyboard

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

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

type Keyboard struct {
	keys []ebiten.Key
	row  [numRows]byte // row scan values
}

func New() (*Keyboard, error) {
	return &Keyboard{
		keys: make([]ebiten.Key, 16),
	}, nil
}

// return the scan code for a row
func (k *Keyboard) Scan(row byte) (byte, error) {
	n, err := rowToInt(row)
	if err != nil {
		return 0xff, err
	}
	// return the scan code (inverted)
	return ^k.row[n], nil
}

// clear all keys
func (k *Keyboard) clear() {
	for i := 0; i < numRows; i++ {
		k.row[i] = 0
	}
}

// set a key down (1) at the row/col
func (k *Keyboard) set(row, col int) {
	k.row[row] |= (1 << col)
}

// Update the keyboard logic (called from ebiten update)
func (k *Keyboard) Update() {
	k.keys = inpututil.AppendPressedKeys(k.keys[:0])
	k.clear()
	for _, key := range k.keys {
		switch key {
		case ebiten.KeyA:
			k.set(1, 0)
		case ebiten.KeyB:
			k.set(7, 3)
		case ebiten.KeyC:
			k.set(0, 4)
		case ebiten.KeyD:
			k.set(1, 2)
		case ebiten.KeyE:
			k.set(2, 2)
		case ebiten.KeyF:
			k.set(1, 3)
		case ebiten.KeyG:
			k.set(1, 4)
		case ebiten.KeyH:
			k.set(6, 4)
		case ebiten.KeyI:
			k.set(5, 2)
		case ebiten.KeyJ:
			k.set(6, 3)
		case ebiten.KeyK:
			k.set(6, 2)
		case ebiten.KeyL:
			k.set(6, 1)
		case ebiten.KeyM:
			k.set(7, 1)
		case ebiten.KeyN:
			k.set(7, 2)
		case ebiten.KeyO:
			k.set(5, 1)
		case ebiten.KeyP:
			k.set(5, 0)
		case ebiten.KeyQ:
			k.set(2, 0)
		case ebiten.KeyR:
			k.set(2, 3)
		case ebiten.KeyS:
			k.set(1, 1)
		case ebiten.KeyT:
			k.set(2, 4)
		case ebiten.KeyU:
			k.set(5, 3)
		case ebiten.KeyV:
			k.set(7, 4)
		case ebiten.KeyW:
			k.set(2, 1)
		case ebiten.KeyX:
			k.set(0, 3)
		case ebiten.KeyY:
			k.set(5, 4)
		case ebiten.KeyZ:
			k.set(0, 2)
		case ebiten.KeyDigit0:
			k.set(4, 0)
		case ebiten.KeyDigit1:
			k.set(3, 0)
		case ebiten.KeyDigit2:
			k.set(3, 1)
		case ebiten.KeyDigit3:
			k.set(3, 2)
		case ebiten.KeyDigit4:
			k.set(3, 3)
		case ebiten.KeyDigit5:
			k.set(3, 4)
		case ebiten.KeyDigit6:
			k.set(4, 4)
		case ebiten.KeyDigit7:
			k.set(4, 3)
		case ebiten.KeyDigit8:
			k.set(4, 2)
		case ebiten.KeyDigit9:
			k.set(4, 1)
		case ebiten.KeyShiftLeft:
			k.set(0, 0)
		case ebiten.KeyShiftRight:
			k.set(0, 1)
		case ebiten.KeySpace:
			k.set(7, 0)
		case ebiten.KeyEnter:
			k.set(6, 0)
		case ebiten.KeyBackspace:
			k.set(0, 0) // left shift
			k.set(4, 0) // 0
		case ebiten.KeyPeriod:
			k.set(0, 1) // right shift
			k.set(7, 1) // m
		case ebiten.KeyShift:
			// ignore
		case ebiten.KeyAltRight, ebiten.KeyAlt, ebiten.KeyAltLeft:
			// ignore
		default:
			fmt.Printf("unmapped key %s\n", key)
		}
	}
}

//-----------------------------------------------------------------------------
