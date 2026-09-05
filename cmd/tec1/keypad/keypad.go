//-----------------------------------------------------------------------------
/*

TEC-1 Keypad Emulation

"Keypad" refers to the keys on the TEC-1 PCB.

The key data comes from the 74C923 keypad encoder.
There are 5 bits coming from the encoder (D0..D4).
There is a CPU reset button.

The encoder providers a data available line that asserts when a
key is pressed. It de-asserts when the key is released.

The 5 bit key code is latched to the output on key down.

The shift/function key is not provided by the encoder and so does not
assert data available.

The real 74C923 has a 2 key rollover, that is not emulated here.

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

type Keypad struct {
	keys  []ebiten.Key
	code  byte // current key code
	latch byte // latched '923 key code
	reset bool // current reset key state
}

func New() (*Keypad, error) {
	return &Keypad{
		keys:  make([]ebiten.Key, 16),
		code:  keyNone,
		latch: keyNone,
	}, nil
}

// return true if the reset button is pressed
func (k *Keypad) Reset() bool {
	return k.reset
}

// return the key code amd shift state from the keypad
func (k *Keypad) Scan() byte {
	return k.latch
}

// Does the '923 have data available?
func (k *Keypad) DataAvailable() bool {
	return k.code != keyNone
}

// get the current key state
func (k *Keypad) getState() (reset bool, code byte) {

	reset = false
	code = keyNone

	k.keys = inpututil.AppendPressedKeys(k.keys[:0])

	// do we have a reset key?
	for _, key := range k.keys {
		if key == ebiten.KeyDelete {
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

	return reset, code
}

// Update the keypad logic (called from ebiten update).
func (k *Keypad) Update() {
	reset, code := k.getState()
	k.reset = reset
	if k.code == keyNone && code != keyNone {
		// key down, latch the key code
		k.latch = code
	}
	k.code = code
}

//-----------------------------------------------------------------------------
