//-----------------------------------------------------------------------------
/*

HD44780 LCD Driver Emulation

*/
//-----------------------------------------------------------------------------

package hd44780

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/bmp"
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

	//cmdEntryModeDec   = byte(0x00)
	//cmdEntryModeShift = byte(0x01)
	//cmdEntryModeInc   = byte(0x02)

	//cmdDisplayCursorBlink = byte(0x01)
	//cmdDisplayCursor      = byte(0x02)
	//cmdDisplayOn          = byte(0x04)

	//cmdShiftCursor  = byte(0x00)
	//cmdShiftDisplay = byte(0x08)
	//cmdShiftLeft    = byte(0x00)
	//cmdShiftRight   = byte(0x04)

	//cmdFunctionLcd1Line = byte(0x00)
	//cmdFunctionLcd2Line = byte(0x08)
	//cmdFunctionExtMode  = byte(0x04)
	//cmdFunctionStdMode  = byte(0x00)
	//cmdExtFunctionGfx   = byte(0x02)
	//cmdExtFunctionStd   = byte(0x00)
)

//-----------------------------------------------------------------------------
/*

display data ram

There are 80 bytes of "live" memory in a 128 byte address space

40 bytes, 0x00-0x27 : live
24 bytes, 0x28-0x3f : dead
40 bytes, 0x40-0x67 : live
24 bytes, 0x68-0x7f : dead

The ddram address is a 7 bit counter that auto increments/decrements over 0x00 to 0x7f.

*/

func inDead0(x byte) bool { return (x >= 0x28) && (x <= 0x3f) }
func inDead1(x byte) bool { return (x >= 0x68) && (x <= 0x7f) }

// increment the ddram address (skip dead areas)
func inc_ddAddr(x byte) byte {
	x = (x + 1) & 0x7f
	if inDead0(x) {
		return 0x40
	}
	if inDead1(x) {
		return 0
	}
	return x
}

// decrement the ddram address (skip dead areas)
func dec_ddAddr(x byte) byte {
	x = (x - 1) & 0x7f
	if inDead0(x) {
		return 0x27
	}
	if inDead1(x) {
		return 0x67
	}
	return x
}

//-----------------------------------------------------------------------------
// display modes

type DisplayMode uint16

const (
	Mode16x2 DisplayMode = ((16 << 8) | 2)
	Mode20x2 DisplayMode = ((20 << 8) | 2)
	Mode20x4 DisplayMode = ((20 << 8) | 4)
	Mode8x1  DisplayMode = ((8 << 8) | 1)
	Mode8x2  DisplayMode = ((8 << 8) | 2)
	Mode12x2 DisplayMode = ((12 << 8) | 2)
	Mode16x1 DisplayMode = ((16 << 8) | 1)
	Mode16x4 DisplayMode = ((16 << 8) | 4)
	Mode20x1 DisplayMode = ((20 << 8) | 1)
	Mode24x2 DisplayMode = ((24 << 8) | 2)
	Mode40x2 DisplayMode = ((40 << 8) | 2)
)

//-----------------------------------------------------------------------------

func saveImageToFile(ebitenImg *ebiten.Image, outputPath string) error {
	bounds := ebitenImg.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	ebitenImg.ReadPixels(rgbaImg.Pix)
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return bmp.Encode(file, rgbaImg)
}

//-----------------------------------------------------------------------------

const dotSize = 10
const dotGap = 1
const glyphWidth = (glyphPixelWidth * (dotSize + dotGap)) - dotGap
const glyphHeight = (glyphPixelHeight * (dotSize + dotGap)) - dotGap

func buildFontImage(font [fontChars][glyphPixelWidth]byte) *ebiten.Image {
	img := ebiten.NewImage(numGlyphs*glyphWidth, glyphHeight)
	for i := cgramSize; i < numGlyphs; i++ {
		for j := 0; j < glyphPixelWidth; j++ {
			pixelData := font[i-cgramSize][j]
			for k := 0; k < glyphPixelHeight; k++ {
				pixel := (pixelData & (1 << (7 - k))) != 0
				if pixel {
					x := float32((i * glyphWidth) + (j * (dotSize + dotGap)))
					y := float32(k * (dotSize + dotGap))
					vector.FillRect(img, x, y, dotSize, dotSize, color.RGBA{0, 0, 0, 255}, false)
				}
			}
		}
	}
	return img
}

//-----------------------------------------------------------------------------

const cgramMode = true
const ddramMode = false

type Config struct {
	Mode           DisplayMode
	XBase, YBase   float64 // xy position
	XScale, YScale float64 // xy scale
}

type LCD struct {
	config        *Config          // lcd configuration
	rows, cols    int              // displays rows and columns
	font          [2]*ebiten.Image // font atlas images
	img           *ebiten.Image    // unscaled lcd image
	rowAddr       []byte           // address of row start in ddram
	ddram         [128]byte        // display data ram
	cgram         []byte           // character generator ram
	cgrom         []byte           // character generator rom
	ddAddr        byte             // ddram address
	cgAddr        byte             // cgram address
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

func New(cfg *Config) (*LCD, error) {
	lcd := &LCD{
		config: cfg,
		rows:   int(cfg.Mode & 0xff),
		cols:   int(cfg.Mode >> 8),
	}
	// pre-load font images
	lcd.font[0] = buildFontImage(fontA00)
	lcd.font[1] = buildFontImage(fontA02)

	// build an unscaled lcd image
	width := lcd.cols*glyphWidth + (lcd.cols-1)*xGap
	height := lcd.rows*glyphHeight + (lcd.rows-1)*yGap
	lcd.img = ebiten.NewImage(width, height)

	// work out the row addresses
	lcd.rowAddr = make([]byte, lcd.rows)
	if lcd.rows == 4 {
		lcd.rowAddr[0] = 0
		lcd.rowAddr[1] = 0x40
		lcd.rowAddr[2] = 0x14
		lcd.rowAddr[3] = 0x54
	} else {
		return nil, errors.New("TODO, rows != 4")
	}

	return lcd, nil
}

func (lcd *LCD) ddramWrite(val byte) {
	//fmt.Printf("ddram write [0x%02x] = 0x%02x\n", lcd.ddAddr, val)

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
		lcd.ddAddr = cmd & 0x7f
		lcd.ramMode = ddramMode
		//fmt.Printf("ddAddr = 0x%02x\n", lcd.ddAddr)

	} else if cmd&cmdSetCgramAddr != 0 {
		// cgram address is 6 bits
		lcd.cgAddr = cmd & 0x3f
		lcd.ramMode = cgramMode

	} else if cmd&cmdFunction != 0 {
		lcd.dlFlag = cmd&(1<<4) != 0
		lcd.nFlag = cmd&(1<<3) != 0
		lcd.fFlag = cmd&(1<<2) != 0

	} else if cmd&cmdShift != 0 {
		fmt.Printf("shift\n")

	} else if cmd&cmdDisplay != 0 {
		lcd.cursorBlink = (cmd & (1 << 0)) != 0
		lcd.cursorEnable = (cmd & (1 << 1)) != 0
		lcd.displayEnable = (cmd & (1 << 2)) != 0
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

// intercharacter gaps
const xGap = dotSize + dotGap
const yGap = dotSize + dotGap

// character to character pitch
const pitchX = glyphWidth + xGap
const pitchY = glyphHeight + yGap

func (lcd *LCD) getGlyph(set int, code byte) *ebiten.Image {
	x := int(code) * glyphWidth
	r := image.Rect(x, 0, x+glyphWidth, glyphHeight)
	return lcd.font[set].SubImage(r).(*ebiten.Image)
}

// Draw the display (called from ebiten draw function)
func (lcd *LCD) Draw(screen *ebiten.Image) {

	// create an unscaled lcd image
	lcd.img.Clear()
	for row := 0; row < lcd.rows; row++ {
		for col := 0; col < lcd.cols; col++ {
			// get the character
			code := lcd.ddram[lcd.rowAddr[row]+byte(col)]
			glyph := lcd.getGlyph(0, code)
			// render the glyph to the lcd image
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(col*pitchX), float64(row*pitchY))
			lcd.img.DrawImage(glyph, op)
		}
	}

	// render the lcd image to the screen (with scaling)
	cfg := lcd.config
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(cfg.XScale, cfg.YScale)
	op.GeoM.Translate(cfg.XBase, cfg.YBase)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(lcd.img, op)
}

// Update the display logic (called from ebiten update)
func (lcd *LCD) Update() {
}

//-----------------------------------------------------------------------------
