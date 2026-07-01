//-----------------------------------------------------------------------------
/*

HD44780 LCD Driver Emulation


display data ram:

There are 80 bytes of memory in a 128 byte address space

40 bytes, 0x00-0x27 : live
24 bytes, 0x28-0x3f : dead
40 bytes, 0x40-0x67 : live
24 bytes, 0x68-0x7f : dead

The ddram address is a 7 bit counter that auto increments/decrements over 0x00 to 0x7f.




*/
//-----------------------------------------------------------------------------

package hd44780

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

const (
	cmdClear        = byte(0x01)
	cmdHome         = byte(0x02)
	cmdEntryMode    = byte(0x04)
	cmdDisplay      = byte(0x08)
	cmdShift        = byte(0x10)
	cmdFunction     = byte(0x20)
	cmdSetCgramAddr = byte(0x40)
	cmdSetDramAddr  = byte(0x80)

	cmdEntryModeDec   = byte(0x00)
	cmdEntryModeShift = byte(0x01)
	cmdEntryModeInc   = byte(0x02)

	cmdDisplayCursorBlink = byte(0x01)
	cmdDisplayCursor      = byte(0x02)
	cmdDisplayOn          = byte(0x04)

	cmdShiftCursor  = byte(0x00)
	cmdShiftDisplay = byte(0x08)
	cmdShiftLeft    = byte(0x00)
	cmdShiftRight   = byte(0x04)

	cmdFunctionLcd1Line = byte(0x00)
	cmdFunctionLcd2Line = byte(0x08)
	cmdFunctionExtMode  = byte(0x04)
	cmdFunctionStdMode  = byte(0x00)
	cmdExtFunctionGfx   = byte(0x02)
	cmdExtFunctionStd   = byte(0x00)
)

//-----------------------------------------------------------------------------

type ramMode int

const (
	ddramMode ramMode = iota
	cgramMode
)

type Config struct {
	Rows, Cols int
}

type LCD struct {
	config        *Config
	ddram         []byte // display data ram
	cgram         []byte // character generator ram
	cgrom         []byte // character generator rom
	ddramAddress  int    // ddram address
	cgramAddress  int    // cgram address
	scrollOffset  int
	mode          ramMode // which ram are we working with?
	cursorBlink   bool    // is the cursor blinking?
	cursorEnable  bool    // is the cursor enabled?
	displayEnable bool    // is the display enabled?
	cursorState   bool    // current cursor state
}

func New(k *Config) (*LCD, error) {
	if (k.Rows != 1) || (k.Rows != 2) || (k.Rows != 4) {
		return nil, errors.New("invalid rows")
	}
	if (k.Cols < 8) || (k.Cols > 40) {
		return nil, errors.New("invalid cols")
	}

	return &LCD{
		config: k,
	}, nil
}

//-----------------------------------------------------------------------------

// read instruction (RS = 0, RW = 1)
func (lcd *LCD) ReadInstruction() byte {
	return 0
}

// write instruction (RS = 0, RW = 0)
func (lcd *LCD) WriteInstruction(cmd byte) {
	if cmd&cmdSetDramAddr != 0 {
		// ddram address is 7 bits
		lcd.ddramAddress = int(cmd) & 0x7f
		lcd.mode = ddramMode
	} else if cmd&cmdSetCgramAddr != 0 {
		// cgram address is 6 bits
		lcd.cgramAddress = int(cmd) & 0x3f
		lcd.mode = cgramMode
	} else if cmd&cmdFunction != 0 {

	} else if cmd&cmdShift != 0 {

	} else if cmd&cmdDisplay != 0 {
		lcd.cursorBlink = (cmd & cmdDisplayCursorBlink) != 0
		lcd.cursorEnable = (cmd & cmdDisplayCursor) != 0
		lcd.displayEnable = (cmd & cmdDisplayOn) != 0
		lcd.cursorState = false
	} else if cmd&cmdEntryMode != 0 {

	} else if cmd&cmdHome != 0 {
		lcd.ddramAddress = 0
		lcd.scrollOffset = 0
	} else if cmd&cmdClear != 0 {
		lcd.ddramAddress = 0
		lcd.scrollOffset = 0
		for i := 0; i < len(lcd.ddram); i++ {
			lcd.ddram[i] = 0x20 // space
		}
	}
}

// read data (RS = 1, RW = 1)
func (lcd *LCD) ReadData() byte {
	return 0
}

// write data (RS = 1, RW = 0)
func (lcd *LCD) WriteData(val byte) {
}

//-----------------------------------------------------------------------------

// Draw the display (called from ebiten draw function)
func (lcd *LCD) Draw(screen *ebiten.Image) {
}

// Update the display logic (called from ebiten update)
func (lcd *LCD) Update() {
}

//-----------------------------------------------------------------------------
