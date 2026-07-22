//-----------------------------------------------------------------------------
/*

HD44780 LCD Driver Emulation

*/
//-----------------------------------------------------------------------------

package hd44780

import (
	"errors"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//-----------------------------------------------------------------------------

// note: cgram has 8 glyphs from 0..7, but they repeat as 8..15
// rather than special case anything, we just treat it as 16.
const cgramGlyphs = 16

// from 16..255 we have the chosen font
const fontGlyphs = 240

// displaying 5x8 glyphs
const glyphPixelWidth = 5
const glyphPixelHeight = 8

// how the pixels are drawn
const dotSize = 10
const dotGap = 1
const glyphWidth = (glyphPixelWidth * (dotSize + dotGap)) - dotGap
const glyphHeight = (glyphPixelHeight * (dotSize + dotGap)) - dotGap

// border around the characters
const xBorder = 20
const yBorder = 20

// cursor blinking
const cursorBlinkPeriod = 24 // 1/60 second units

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
	cmdShiftDisplay = byte(0x08)
	//cmdShiftLeft    = byte(0x00)
	cmdShiftRight = byte(0x04)

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
func inc_ddAdr(x byte) byte {
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
func dec_ddAdr(x byte) byte {
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
/*

Display Modes

The hd44780 has two 40 byte "rows" of internal display data ram.
How these map onto the physical lcd row and column is a function
of the actual display being used. This code handles that mapping.

*/

// return true if this is a supported display mode
func isValid(row, col int, type2 bool) bool {
	switch row {
	case 1:
		if type2 {
			switch col {
			case 16, 20, 24, 32, 40:
				return true
			}
		} else {
			switch col {
			case 8:
				return true
			}
		}
	case 2:
		switch col {
		case 8, 16, 20, 24, 40:
			return true
		}
	case 4:
		switch col {
		case 16, 20:
			return true
		}
	}
	return false
}

// convert a row and column to a display data address
func (lcd *LCD) rowColToAdr(row, col int) byte {
	switch lcd.cfg.Rows {
	case 1:
		half := lcd.cfg.Cols / 2
		if lcd.cfg.Type2 && col >= half {
			// a type2 lcd splits it's single row of columms
			// over the two "rows" of the display data ram.
			return 0x40 + byte(col-half)
		}
		return 0x00 + byte(col)
	case 2:
		return []byte{0, 0x40}[row] + byte(col)
	case 4:
		return []byte{0, 0x40, 0x14, 0x54}[row] + byte(col)
	}
	return 0xff
}

// map a display data ram address to a row and column.
func (lcd *LCD) adrToRowCol(adr byte) (int, int, bool) {
	x, ok := lcd.toRowCol[adr]
	return x[0], x[1], ok
}

//-----------------------------------------------------------------------------

// build a font image from row column pixel data
func buildFontImage(font [fontGlyphs][glyphPixelWidth]byte, color color.RGBA) *ebiten.Image {
	img := ebiten.NewImage(fontGlyphs*glyphWidth, glyphHeight)
	for i := 0; i < fontGlyphs; i++ {
		for j := 0; j < glyphPixelWidth; j++ {
			pixelData := font[i][j]
			for k := 0; k < glyphPixelHeight; k++ {
				pixel := (pixelData & (1 << (glyphPixelHeight - 1 - k))) != 0
				if pixel {
					x := float32((i * glyphWidth) + (j * (dotSize + dotGap)))
					y := float32(k * (dotSize + dotGap))
					vector.FillRect(img, x, y, dotSize, dotSize, color, false)
				}
			}
		}
	}
	return img
}

// build an image from row ordered pixel data (cgram style)
func buildImage(buf []byte, color color.RGBA) *ebiten.Image {
	nGlyphs := len(buf) >> 3 // 8 bytes per glyph
	img := ebiten.NewImage(nGlyphs*glyphWidth, glyphHeight)
	for i := 0; i < nGlyphs; i++ {
		for j := 0; j < glyphPixelHeight; j++ {
			pixelData := buf[(i<<3)+j]
			for k := 0; k < glyphPixelWidth; k++ {
				pixel := (pixelData & (1 << (glyphPixelWidth - 1 - k))) != 0
				if pixel {
					x := float32((i * glyphWidth) + (k * (dotSize + dotGap)))
					y := float32(j * (dotSize + dotGap))
					vector.FillRect(img, x, y, dotSize, dotSize, color, false)
				}
			}
		}
	}
	return img
}

//-----------------------------------------------------------------------------

const cgRamMode = true
const ddRamMode = false

const ddRamSize = 128
const cgRamSize = 64

type Config struct {
	Rows, Cols      int     // rows and columns for the display
	Type2           bool    // is this a single row, type2 lcd?
	Font            int     // font selector
	XBase, YBase    float64 // xy position
	XScale, YScale  float64 // xy scale
	BackgroundColor color.RGBA
	CharacterColor  color.RGBA
}

type LCD struct {
	cfg Config // lcd configuration

	toRowCol map[byte][2]int // map a ddRam address to a (row,col) tuple

	// images
	font       *ebiten.Image // font atlas
	cgFont     *ebiten.Image // character generator glyph atlas
	cursorFont *ebiten.Image // cursor glyph atlas
	img        *ebiten.Image // lcd image

	// character generator ram
	cgRam   [cgRamSize]byte // character generator ram
	cgAdr   byte            // cgram address
	cgDirty bool            // do we need to rebuild the cg glyphs in the font atlas?

	// display data ram
	ddRam [ddRamSize]byte // display data ram
	ddAdr byte            // ddram address

	// cursor state
	cursorCount  int  // cursor flash counter
	cursorBlink  bool // is the cursor blinking?
	cursorEnable bool // is the cursor enabled?
	cursorState  bool // current cursor state

	// other variables...
	displayEnable bool // is the display enabled?
	scrollOffset  int
	ramMode       bool // which ram are we working with (cgRam/ddRam)?
	dlFlag        bool // interface data width (false = 4, true = 8)
	nFlag         bool // number of display lines (false = 1, true = 2)
	fFlag         bool // font selection (false = 5x8, true = 5x10)
	shiftMode     bool // shift mode (entry mode set command)
	incMode       bool // increment mode (entry mode set command)
}

func New(cfg Config) (*LCD, error) {
	// check the display mode
	if !isValid(cfg.Rows, cfg.Cols, cfg.Type2) {
		return nil, errors.New("unsupported display mode")
	}
	lcd := &LCD{
		cfg:      cfg,
		toRowCol: make(map[byte][2]int),
	}

	// work out the data display address to (row,col) map
	for row := 0; row < cfg.Rows; row++ {
		for col := 0; col < cfg.Cols; col++ {
			adr := lcd.rowColToAdr(row, col)
			lcd.toRowCol[adr] = [2]int{row, col}
		}
	}

	// load a font atlas
	switch cfg.Font {
	case 0:
		lcd.font = buildFontImage(fontA00, cfg.CharacterColor)
	case 1:
		lcd.font = buildFontImage(fontA02, cfg.CharacterColor)
	default:
		return nil, errors.New("invalid font selector")
	}

	// build an initial cgRam font
	lcd.cgFont = buildImage(lcd.cgRam[:], cfg.CharacterColor)

	// clear the ddRam
	for i := range lcd.ddRam {
		lcd.ddRam[i] = 0x20
	}

	// build the cursor font
	lcd.cursorFont = buildImage(cursorData, cfg.CharacterColor)

	// build an lcd image
	width := cfg.Cols*glyphWidth + (cfg.Cols-1)*xGap + (2 * xBorder)
	height := cfg.Rows*glyphHeight + (cfg.Rows-1)*yGap + (2 * yBorder)
	lcd.img = ebiten.NewImage(width, height)

	return lcd, nil
}

//-----------------------------------------------------------------------------

func (lcd *LCD) ddRamWrite(val byte) {
	//log.Printf("ddRamWrite [0x%02x] = 0x%02x", lcd.ddAddr, val)
	lcd.ddRam[lcd.ddAdr] = val
	if lcd.incMode {
		lcd.ddAdr = inc_ddAdr(lcd.ddAdr)
	} else {
		lcd.ddAdr = dec_ddAdr(lcd.ddAdr)
	}
}

func (lcd *LCD) ddRamRead() byte {
	//log.Printf("ddRamRead [0x%02x]", lcd.ddAdr)
	val := lcd.ddRam[lcd.ddAdr]
	// auto increment/decrement the address
	if lcd.incMode {
		lcd.ddAdr = inc_ddAdr(lcd.ddAdr)
	} else {
		lcd.ddAdr = dec_ddAdr(lcd.ddAdr)
	}
	return val
}

//-----------------------------------------------------------------------------

const cgAdrMask = byte(cgRamSize - 1)

func (lcd *LCD) cgRamWrite(val byte) {
	//log.Printf("cgRamWrite [0x%02x] = 0x%02x", lcd.cgAdr, val)
	if lcd.cgRam[lcd.cgAdr] != val {
		lcd.cgRam[lcd.cgAdr] = val
		// set the dirty flag, we need to rebuild the font atlas
		lcd.cgDirty = true
	}
	if lcd.incMode {
		lcd.cgAdr = (lcd.cgAdr + 1) & cgAdrMask
	} else {
		lcd.cgAdr = (lcd.cgAdr - 1) & cgAdrMask
	}
}

func (lcd *LCD) cgRamRead() byte {
	//log.Printf("cgRamRead [0x%02x]", lcd.cgAdr)
	val := lcd.cgRam[lcd.cgAdr]
	if lcd.incMode {
		lcd.cgAdr = (lcd.cgAdr + 1) & cgAdrMask
	} else {
		lcd.cgAdr = (lcd.cgAdr - 1) & cgAdrMask
	}
	return val
}

//-----------------------------------------------------------------------------

// read command (RS = 0, RW = 1)
func (lcd *LCD) ReadCommand() byte {
	// Note: the busy flag is == 0
	if lcd.ramMode == ddRamMode {
		return byte(lcd.ddAdr)
	}
	return byte(lcd.cgAdr)
}

// write command (RS = 0, RW = 0)
func (lcd *LCD) WriteCommand(cmd byte) {
	if cmd&cmdSetDramAddr != 0 {
		// ddRam address is 7 bits
		lcd.ddAdr = cmd & 0x7f
		lcd.ramMode = ddRamMode

	} else if cmd&cmdSetCgramAddr != 0 {
		// cgram address is 6 bits
		lcd.cgAdr = cmd & cgAdrMask
		lcd.ramMode = cgRamMode

	} else if cmd&cmdFunction != 0 {
		lcd.dlFlag = cmd&(1<<4) != 0
		lcd.nFlag = cmd&(1<<3) != 0
		lcd.fFlag = cmd&(1<<2) != 0

	} else if cmd&cmdShift != 0 {
		if cmd&cmdShiftDisplay != 0 {
			// shift display
			if cmd&cmdShiftRight != 0 {
				lcd.scrollOffset -= 1
			} else {
				lcd.scrollOffset += 1
			}
		} else {
			// shift cursor
			if cmd&cmdShiftRight != 0 {
				lcd.ddAdr = inc_ddAdr(lcd.ddAdr)
			} else {
				lcd.ddAdr = dec_ddAdr(lcd.ddAdr)
			}
		}

	} else if cmd&cmdDisplay != 0 {
		lcd.cursorBlink = (cmd & (1 << 0)) != 0
		lcd.cursorEnable = (cmd & (1 << 1)) != 0
		lcd.displayEnable = (cmd & (1 << 2)) != 0
		lcd.cursorState = false

	} else if cmd&cmdEntryMode != 0 {
		lcd.shiftMode = cmd&(1<<0) != 0
		lcd.incMode = cmd&(1<<1) != 0

	} else if cmd&cmdHome != 0 {
		lcd.ddAdr = 0
		lcd.scrollOffset = 0

	} else if cmd&cmdClear != 0 {
		lcd.ddAdr = 0
		lcd.scrollOffset = 0
		lcd.incMode = true
		for i := 0; i < len(lcd.ddRam); i++ {
			lcd.ddRam[i] = 0x20 // space
		}
	}
}

// read data (RS = 1, RW = 1)
func (lcd *LCD) ReadData() byte {
	if lcd.ramMode == ddRamMode {
		return lcd.ddRamRead()
	}
	return lcd.cgRamRead()
}

// write data (RS = 1, RW = 0)
func (lcd *LCD) WriteData(val byte) {
	if lcd.ramMode == ddRamMode {
		lcd.ddRamWrite(val)
	} else {
		lcd.cgRamWrite(val)
	}
}

//-----------------------------------------------------------------------------

// return the glyph for a given code
func (lcd *LCD) getGlyph(code byte) *ebiten.Image {
	var img *ebiten.Image
	x := 0
	if code < cgramGlyphs {
		// glyph is in cgram, 8..15 map onto 0..7
		code &= 7
		x = int(code) * glyphWidth
		img = lcd.cgFont
	} else {
		// glyph is in the chosen ROM font
		x = int(code-cgramGlyphs) * glyphWidth
		img = lcd.font
	}
	r := image.Rect(x, 0, x+glyphWidth, glyphHeight)
	return img.SubImage(r).(*ebiten.Image)
}

// return the image for the cursor
func (lcd *LCD) getCursor(state bool) *ebiten.Image {
	x := 0
	if state {
		x = glyphWidth
	}
	r := image.Rect(x, 0, x+glyphWidth, glyphHeight)
	return lcd.cursorFont.SubImage(r).(*ebiten.Image)
}

//-----------------------------------------------------------------------------

// intercharacter gaps
const xGap = dotSize + dotGap
const yGap = dotSize + dotGap

// character to character pitch
const pitchX = glyphWidth + xGap
const pitchY = glyphHeight + yGap

// Draw the display (called from ebiten draw function)
func (lcd *LCD) Draw(screen *ebiten.Image) {

	cfg := &lcd.cfg

	// clear the lcd image
	lcd.img.Clear()
	lcd.img.Fill(lcd.cfg.BackgroundColor)

	if lcd.displayEnable {
		// draw the glyphs
		for row := 0; row < cfg.Rows; row++ {
			for col := 0; col < cfg.Cols; col++ {
				// get the character
				code := lcd.ddRam[lcd.rowColToAdr(row, col)]
				glyph := lcd.getGlyph(code)
				// render the glyph to the lcd image
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(col*pitchX+xBorder), float64(row*pitchY+yBorder))
				lcd.img.DrawImage(glyph, op)
			}
		}
		// draw the cursor
		if lcd.cursorEnable {
			row, col, ok := lcd.adrToRowCol(lcd.ddAdr)
			if ok {
				// render the cursor to the lcd image
				cursor := lcd.getCursor(lcd.cursorBlink && lcd.cursorState)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(col*pitchX+xBorder), float64(row*pitchY+yBorder))
				lcd.img.DrawImage(cursor, op)
			}
		}
	}

	// render the lcd image to the screen (with scaling)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(cfg.XScale, cfg.YScale)
	op.GeoM.Translate(cfg.XBase, cfg.YBase)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(lcd.img, op)
}

// Update the display logic (called from ebiten update)
func (lcd *LCD) Update() {

	// update the cursor state
	if lcd.cursorCount == cursorBlinkPeriod {
		lcd.cursorState = !lcd.cursorState
		lcd.cursorCount = 0
	}
	lcd.cursorCount += 1

	// rebuild the cgram font if it has been changed
	if lcd.cgDirty {
		lcd.cgFont = buildImage(lcd.cgRam[:], lcd.cfg.CharacterColor)
		lcd.cgDirty = false
	}
}

//-----------------------------------------------------------------------------
