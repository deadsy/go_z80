//-----------------------------------------------------------------------------
/*

TEC-1 Emulation

Keypad

*/
//-----------------------------------------------------------------------------

package keypad

import (
	"log"

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
	keys    []ebiten.Key
	current ebiten.Key // current key being pressed
	code    byte
	reset   bool
}

func New() (*Keypad, error) {
	return &Keypad{
		keys:    make([]ebiten.Key, 16),
		code:    keyNone,
		current: ebiten.KeyMax,
	}, nil
}

// return true if the reset button is pressed
func (k *Keypad) Reset() bool {
	if k.reset {
		k.reset = false
		return true
	}
	return false
}

// return the san code from the keypad
func (k *Keypad) Scan() byte {
	return k.code
}

func (k *Keypad) set(key ebiten.Key, code byte) bool {
	k.current = key
	k.code = code
	return true
}

// Update the keyboard logic (called from ebiten update).
// Return true if a keypress is recognized.
func (k *Keypad) Update() bool {

	k.keys = inpututil.AppendPressedKeys(k.keys[:0])

	// is the current key still being pressed?
	if k.current != ebiten.KeyMax {
		for _, key := range k.keys {
			if k.current == key {
				// yes... no new keypress
				return false
			}
		}
	}
	// no ... reset the current key
	k.current = ebiten.KeyMax
	k.code = keyNone

	for _, key := range k.keys {
		switch key {
		case ebiten.KeyA:
			return k.set(key, keyA)
		case ebiten.KeyB:
			return k.set(key, keyB)
		case ebiten.KeyC:
			return k.set(key, keyC)
		case ebiten.KeyD:
			return k.set(key, keyD)
		case ebiten.KeyE:
			return k.set(key, keyE)
		case ebiten.KeyF:
			return k.set(key, keyF)
		case ebiten.KeyDigit0:
			return k.set(key, key0)
		case ebiten.KeyDigit1:
			return k.set(key, key1)
		case ebiten.KeyDigit2:
			return k.set(key, key2)
		case ebiten.KeyDigit3:
			return k.set(key, key3)
		case ebiten.KeyDigit4:
			return k.set(key, key4)
		case ebiten.KeyDigit5:
			return k.set(key, key5)
		case ebiten.KeyDigit6:
			return k.set(key, key6)
		case ebiten.KeyDigit7:
			return k.set(key, key7)
		case ebiten.KeyDigit8:
			return k.set(key, key8)
		case ebiten.KeyDigit9:
			return k.set(key, key9)
		case ebiten.KeyArrowLeft: // -
			return k.set(key, keyMinus)
		case ebiten.KeyArrowRight: // +
			return k.set(key, keyPlus)
		case ebiten.KeyEnter: // go
			return k.set(key, keyGo)
		case ebiten.KeyEscape: // address
			return k.set(key, keyAddress)
		case ebiten.KeyR:
			k.reset = true
			return k.set(key, keyNone)
		default:
			log.Printf("unmapped key %s", key)
		}
	}
	return false
}

//-----------------------------------------------------------------------------
