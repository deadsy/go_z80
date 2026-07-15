//-----------------------------------------------------------------------------
/*

TEC-1G Keypad Emulation

"Keypad" refers to the keys on the TEC-1G PCB.

The key data comes from the 74C923 keypad encoder.
There are 5 bits coming from the encoder (D0..D4).
There is an extra bit (D5) used as a function/shift key.
There is a CPU reset button.

The key mapping should be standard for all platforms:

0..9, A..F = same keys
"-" = arrow left
"+" = arrow right
address = escape
go = enter
reset = "r"
function/shift = shift key

*/
//-----------------------------------------------------------------------------

package keypad

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//-----------------------------------------------------------------------------

const (
	key0       byte = 0x00
	key1       byte = 0x01
	key2       byte = 0x02
	key3       byte = 0x03
	key4       byte = 0x04
	key5       byte = 0x05
	key6       byte = 0x06
	key7       byte = 0x07
	key8       byte = 0x08
	key9       byte = 0x09
	keyA       byte = 0x0a
	keyB       byte = 0x0b
	keyC       byte = 0x0c
	keyD       byte = 0x0d
	keyE       byte = 0x0e
	keyF       byte = 0x0f
	keyPlus    byte = 0x10
	keyMinus   byte = 0x11
	keyGo      byte = 0x12
	keyAddress byte = 0x13
	keyNone    byte = 0xff
)

const shiftMask = byte(1 << 5) // D5

type Keypad struct {
	enable bool // is the keypad enabled
	keys   []ebiten.Key
	code   byte
	reset  bool
}

func New(enable bool) (*Keypad, error) {
	return &Keypad{
		keys: make([]ebiten.Key, 16),
		code: keyNone,
	}, nil
}

// return true if the reset button is pressed
func (k *Keypad) Reset() bool {
	if !k.enable {
		return false
	}
	if k.reset {
		k.reset = false
		return true
	}
	return false
}

// return the san code from the keypad
func (k *Keypad) Scan() byte {
	if !k.enable {
		return keyNone
	}
	return k.code
}

// get the current keypad code
func (k *Keypad) getCode() (reset bool, code byte) {

	reset = false
	shift := byte(0)
	code = keyNone

	k.keys = inpututil.AppendPressedKeys(k.keys[:0])

	// do we have a reset or shift key?
	for _, key := range k.keys {
		switch key {
		case ebiten.KeyShift, ebiten.KeyShiftLeft, ebiten.KeyShiftRight:
			shift = shiftMask
		case ebiten.KeyR:
			reset = true
		}
	}

	// do we have an encoder key?
	for _, key := range k.keys {
		switch key {
		case ebiten.KeyA:
			code = keyA
		case ebiten.KeyB:
			code = keyB
		case ebiten.KeyC:
			code = keyC
		case ebiten.KeyD:
			code = keyD
		case ebiten.KeyE:
			code = keyE
		case ebiten.KeyF:
			code = keyF
		case ebiten.KeyDigit0:
			code = key0
		case ebiten.KeyDigit1:
			code = key1
		case ebiten.KeyDigit2:
			code = key2
		case ebiten.KeyDigit3:
			code = key3
		case ebiten.KeyDigit4:
			code = key4
		case ebiten.KeyDigit5:
			code = key5
		case ebiten.KeyDigit6:
			code = key6
		case ebiten.KeyDigit7:
			code = key7
		case ebiten.KeyDigit8:
			code = key8
		case ebiten.KeyDigit9:
			code = key9
		case ebiten.KeyArrowLeft: // -
			code = keyMinus
		case ebiten.KeyArrowRight: // +
			code = keyPlus
		case ebiten.KeyEnter: // go
			code = keyGo
		case ebiten.KeyEscape: // address
			code = keyAddress
		default:
			//log.Printf("unmapped key %s", key)
		}
	}
	return reset, code | shift
}

// Update the keyboard logic (called from ebiten update).
// Return true if a keypress is recognized.
func (k *Keypad) Update() bool {
	if !k.enable {
		return false
	}
	reset, code := k.getCode()
	k.reset = reset
	// is the current key still being pressed?
	if k.code == code {
		// no new keypress
		return false
	}
	// set the current key
	k.code = code
	return true
}

//-----------------------------------------------------------------------------
