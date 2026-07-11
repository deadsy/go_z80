//-----------------------------------------------------------------------------
/*


TEC-1G Matrix Keyboard

Maps the pressed ebiten key values into the row/col scan codes required by
the tec1g matrix keyboard scanning code.

See: https://github.com/MarkJelic/TEC-1G/blob/main/Keyboards/Mechanical/TEC-1G_Matrix-Mechanical-Keyboard_Schematic-v10.pdf

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
		case ebiten.KeyA: // A12_D4 A
			k.set(4, 4)
		case ebiten.KeyB: // A12_D5 B
			k.set(4, 5)
		case ebiten.KeyC: // A12_D6 C
			k.set(4, 6)
		case ebiten.KeyD: // A12_D7 D
			k.set(4, 7)
		case ebiten.KeyE: // A13_D0 E
			k.set(5, 0)
		case ebiten.KeyF: // A13_D1 F
			k.set(5, 1)
		case ebiten.KeyG: // A13_D2 G
			k.set(5, 2)
		case ebiten.KeyH: // A13_D3 H
			k.set(5, 3)
		case ebiten.KeyI: // A13_D4 I
			k.set(5, 4)
		case ebiten.KeyJ: // A13_D5 J
			k.set(5, 5)
		case ebiten.KeyK: // A13_D6 K
			k.set(5, 6)
		case ebiten.KeyL: // A13_D7 L
			k.set(5, 7)
		case ebiten.KeyM: // A14_D0 M
			k.set(6, 0)
		case ebiten.KeyN: // A14_D1 N
			k.set(6, 1)
		case ebiten.KeyO: // A14_D2 O
			k.set(6, 2)
		case ebiten.KeyP: // A14_D3 P
			k.set(6, 3)
		case ebiten.KeyQ: // A14_D4 Q
			k.set(6, 4)
		case ebiten.KeyR: // A14_D5 R
			k.set(6, 5)
		case ebiten.KeyS: // A14_D6 S
			k.set(6, 6)
		case ebiten.KeyT: // A14_D7 T
			k.set(6, 7)
		case ebiten.KeyU: // A15_D0 U
			k.set(7, 0)
		case ebiten.KeyV: // A15_D1 V
			k.set(7, 1)
		case ebiten.KeyW: // A15_D2 W
			k.set(7, 2)
		case ebiten.KeyX: // A15_D3 X
			k.set(7, 3)
		case ebiten.KeyY: // A15_D4 Y
			k.set(7, 4)
		case ebiten.KeyZ: // A15_D5 Z
			k.set(7, 5)
		case ebiten.KeyDigit0: // A10_D3 0
			k.set(2, 3)
		case ebiten.KeyDigit1: // A10_D4 1
			k.set(2, 4)
		case ebiten.KeyDigit2: // A10_D5 2
			k.set(2, 5)
		case ebiten.KeyDigit3: // A10_D6 3
			k.set(2, 6)
		case ebiten.KeyDigit4: // A10_D7 4
			k.set(2, 7)
		case ebiten.KeyDigit5: // A11_D0 5
			k.set(3, 0)
		case ebiten.KeyDigit6: // A11_D1 6
			k.set(3, 1)
		case ebiten.KeyDigit7: // A11_D2 7
			k.set(3, 2)
		case ebiten.KeyDigit8: // A11_D3 8
			k.set(3, 3)
		case ebiten.KeyDigit9: // A11_D4 9
			k.set(3, 4)
		case ebiten.KeyBackslash: // A15_D7 \
			k.set(7, 7)
		case ebiten.KeyEqual: // A12_D0 =
			k.set(4, 0)
		case ebiten.KeySemicolon: // A11_D6 ;
			k.set(3, 6)
		case ebiten.KeyMinus: // A10_D0 -
			k.set(2, 0)
		case ebiten.KeyPeriod: // A10_D1 .
			k.set(2, 1)
		case ebiten.KeySlash: // A10_D2 /
			k.set(2, 2)
		case ebiten.KeyDelete: // A9_D0 del
			k.set(1, 0)
		case ebiten.KeyTab: // A9_D1 tab
			k.set(1, 1)
		case ebiten.KeyEnter: // A9_D2 enter/go
			k.set(1, 2)
		case ebiten.KeyEscape: // A9_D4 esc/addr
			k.set(1, 4)
		case ebiten.KeySpace: // A9_D5 space
			k.set(1, 5)
		case ebiten.KeyQuote: // A9_D6 '
			k.set(1, 6)
		case ebiten.KeyComma: // A9_D7 ,
			k.set(1, 7)
		case ebiten.KeyShift, ebiten.KeyShiftLeft, ebiten.KeyShiftRight: // A8_D0 shift
			k.set(0, 0)
		case ebiten.KeyControl, ebiten.KeyControlLeft, ebiten.KeyControlRight: // A8_D1 ctrl
			k.set(0, 1)
		case ebiten.KeyAlt, ebiten.KeyAltLeft, ebiten.KeyAltRight: // A8_D2 function
			k.set(0, 2)
		case ebiten.KeyArrowUp: // A8_D3 up
			k.set(0, 3)
		case ebiten.KeyArrowDown: // A8_D4 down
			k.set(0, 4)
		case ebiten.KeyArrowLeft: // A8_D5 left/-
			k.set(0, 5)
		case ebiten.KeyArrowRight: // A8_D6 right/+
			k.set(0, 6)
		case ebiten.KeyCapsLock: // A8_D7 caps
			k.set(0, 7)
		default:
			fmt.Printf("unmapped key %s\n", key)
		}
	}
}

//-----------------------------------------------------------------------------
