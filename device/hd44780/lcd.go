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
	"fmt"

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
// display modes

type DisplayMode int

const (
	Mode4x20 DisplayMode = iota
)

//-----------------------------------------------------------------------------

const cgramMode = true
const ddramMode = false

type Config struct {
	Mode DisplayMode
}

type LCD struct {
	config        *Config
	rows, cols    int       // displays rows and columns
	ddram         [128]byte // display data ram
	cgram         []byte    // character generator ram
	cgrom         []byte    // character generator rom
	ddAddr        int       // ddram address
	cgAddr        int       // cgram address
	scrollOffset  int
	ramMode       bool // which ram are we working with?
	cursorBlink   bool // is the cursor blinking?
	cursorEnable  bool // is the cursor enabled?
	displayEnable bool // is the display enabled?
	cursorState   bool // current cursor state
	dlFlag        bool // interface data width (false = 4, true = 8)
	nFlag         bool // number of display lines (false = 1, true = 2)
	fFlag         bool // font selection (false = 5x8, true = 5x10)
	incMode       bool // increment mode

}

func New(k *Config) (*LCD, error) {
	lcd := &LCD{
		config: k,
	}
	switch k.Mode {
	case Mode4x20:
		lcd.rows = 4
		lcd.cols = 20
	default:
		return nil, errors.New("unsupported mode")
	}
	return lcd, nil
}

// increment the ddram address
func inc_ddAddr(addr int) int {
	if (addr >= 0x27) && (addr <= 0x3f) {
		return 0x40
	}
	if (addr >= 0x67) && (addr <= 0x7f) {
		return 0
	}
	return addr + 1
}

// decrement the ddram address
func dec_ddAddr(addr int) int {
	if ((addr >= 0x28) && (addr <= 0x3f)) || (addr == 0x40) {
		return 0x27
	}
	if ((addr >= 0x68) && (addr <= 0x7f)) || (addr == 0) {
		return 0x67
	}
	return addr - 1
}

func (lcd *LCD) ddramWrite(val byte) {
	fmt.Printf("ddram write [0x%02x] = 0x%02x\n", lcd.ddAddr, val)

	lcd.ddram[lcd.ddAddr] = val

	if lcd.incMode {
		lcd.ddAddr = inc_ddAddr(lcd.ddAddr)

		/*

		   if (self->mcu.LCD_EntryMode & 0x01) {
		       if (self->mcu.DDRAM_display < 24)
		           self->mcu.DDRAM_display++;
		   }

		*/

	} else {
		lcd.ddAddr = dec_ddAddr(lcd.ddAddr)

		/*

		   if (self->mcu.LCD_EntryMode & 0x01) {
		       if (self->mcu.DDRAM_display > 0)
		           self->mcu.DDRAM_display--;
		   }

		*/

	}

}

func (lcd *LCD) cgramWrite(val byte) {
	fmt.Printf("cgram write\n")
	/*
	   n = self->mcu.CGRAM_counter / 8;
	   m = self->mcu.CGRAM_counter % 8;
	   self->mcu.CGROM[n][m] = instruction & 0xFF;
	   if (self->mcu.CGRAM_counter < 64)

	   	self->mcu.CGRAM_counter++;
	*/
}

//-----------------------------------------------------------------------------

// read command (RS = 0, RW = 1)
func (lcd *LCD) ReadCommand() byte {
	// Note: the busy flag is == 0
	if lcd.ramMode == ddramMode {
		return byte(lcd.ddAddr)
	}
	return byte(lcd.cgAddr)
}

// write command (RS = 0, RW = 0)
func (lcd *LCD) WriteCommand(cmd byte) {
	if cmd&cmdSetDramAddr != 0 {
		// ddram address is 7 bits
		lcd.ddAddr = int(cmd) & 0x7f
		lcd.ramMode = ddramMode
		fmt.Printf("ddAddr = 0x%02x\n", lcd.ddAddr)

	} else if cmd&cmdSetCgramAddr != 0 {
		// cgram address is 6 bits
		lcd.cgAddr = int(cmd) & 0x3f
		lcd.ramMode = cgramMode

	} else if cmd&cmdFunction != 0 {
		lcd.dlFlag = cmd&(1<<4) != 0
		lcd.nFlag = cmd&(1<<3) != 0
		lcd.fFlag = cmd&(1<<2) != 0

	} else if cmd&cmdShift != 0 {
		fmt.Printf("shift\n")

	} else if cmd&cmdDisplay != 0 {
		lcd.cursorBlink = (cmd & cmdDisplayCursorBlink) != 0
		lcd.cursorEnable = (cmd & cmdDisplayCursor) != 0
		lcd.displayEnable = (cmd & cmdDisplayOn) != 0
		lcd.cursorState = false

	} else if cmd&cmdEntryMode != 0 {
		lcd.incMode = cmd&(1<<1) != 0
		if cmd&(1<<0) != 0 {
			fmt.Printf("TODO: entry mode shift\n")
		}
	} else if cmd&cmdHome != 0 {
		lcd.ddAddr = 0
		lcd.scrollOffset = 0

	} else if cmd&cmdClear != 0 {
		lcd.ddAddr = 0
		lcd.scrollOffset = 0
		lcd.incMode = true
		for i := 0; i < len(lcd.ddram); i++ {
			lcd.ddram[i] = 0x20 // space
		}
	}
}

// read data (RS = 1, RW = 1)
func (lcd *LCD) ReadData() byte {
	fmt.Printf("lcd data read\n")
	return 0
}

// write data (RS = 1, RW = 0)
func (lcd *LCD) WriteData(val byte) {
	if lcd.ramMode == ddramMode {
		lcd.ddramWrite(val)
	} else {
		lcd.cgramWrite(val)
	}
}

//-----------------------------------------------------------------------------

// Draw the display (called from ebiten draw function)
func (lcd *LCD) Draw(screen *ebiten.Image) {
}

// Update the display logic (called from ebiten update)
func (lcd *LCD) Update() {
}

//-----------------------------------------------------------------------------
