//-----------------------------------------------------------------------------
/*

TEC-1G Matrix Keyboard

See: https://github.com/MarkJelic/TEC-1G/blob/main/Keyboards/Mechanical/TEC-1G_Matrix-Mechanical-Keyboard_Schematic-v10.pdf

*/
//-----------------------------------------------------------------------------

package keyboard

import "errors"

//-----------------------------------------------------------------------------

// A15_D0 U
// A15_D1 V
// A15_D2 W
// A15_D3 X
// A15_D4 Y
// A15_D5 Z
// A15_D6 unused
// A15_D7 \

// A14_D0 M
// A14_D1 N
// A14_D2 O
// A14_D3 P
// A14_D4 Q
// A14_D5 R
// A14_D6 S
// A14_D7 T

// A13_D0 E
// A13_D1 F
// A13_D2 G
// A13_D3 H
// A13_D4 I
// A13_D5 J
// A13_D6 K
// A13_D7 L

// A12_D0 =
// A12_D1 unused
// A12_D2 unused
// A12_D3 unused
// A12_D4 A
// A12_D5 B
// A12_D6 C
// A12_D7 D

// A11_D0 5
// A11_D1 6
// A11_D2 7
// A11_D3 8
// A11_D4 9
// A11_D5 unused
// A11_D6 ;
// A11_D7 unused

// A10_D0 -
// A10_D1 .
// A10_D2 /
// A10_D3 0
// A10_D4 1
// A10_D5 2
// A10_D6 3
// A10_D7 4

// A9_D0 del
// A9_D1 tab
// A9_D2 enter/go
// A9_D3 unused
// A9_D4 esc/addr
// A9_D5 space
// A9_D6 '
// A9_D7 ,

// A8_D0 shift
// A8_D1 ctrl
// A8_D2 function
// A8_D3 up
// A8_D4 down
// A8_D5 left/-
// A8_D6 right/+
// A8_D7 caps

//-----------------------------------------------------------------------------
// key row values (the low bit selects the row)

const keyRow0 = byte(0xfe) // A8
const keyRow1 = byte(0xfd) // A9
const keyRow2 = byte(0xfb) // A10
const keyRow3 = byte(0xf7) // A11
const keyRow4 = byte(0xef) // A12
const keyRow5 = byte(0xdf) // A13
const keyRow6 = byte(0xbf) // A14
const keyRow7 = byte(0x7f) // A15

//-----------------------------------------------------------------------------

const (
	keyShift = iota
	keyCtrl
	keyFunction
	keyUp
	keyDown
	keyLeft
	keyRight
	keyCaps
)

//-----------------------------------------------------------------------------

func (k *Keyboard) keyState(key, shift int) byte {
	if k.state[key] {
		return 1 << shift
	}
	return 0
}

//-----------------------------------------------------------------------------

type Keyboard struct {
	state []bool
}

func New() (*Keyboard, error) {
	return &Keyboard{}, nil
}

func (k *Keyboard) Scan(row byte) (byte, error) {
	var col byte
	switch row {
	case keyRow0:
		col = k.keyState(keyShift, 0) |
			k.keyState(keyCtrl, 1) |
			k.keyState(keyFunction, 2) |
			k.keyState(keyUp, 3) |
			k.keyState(keyDown, 4) |
			k.keyState(keyLeft, 5) |
			k.keyState(keyRight, 6) |
			k.keyState(keyCaps, 7)
	case keyRow1:
	case keyRow2:
	case keyRow3:
	case keyRow4:
	case keyRow5:
	case keyRow6:
	case keyRow7:
	default:
		return 0xff, errors.New("bad row value")

	}
	// return the column scan (inverted)
	return ^col, nil
}

//-----------------------------------------------------------------------------
