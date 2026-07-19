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
// display modes

type DisplayMode uint16

const (
	Mode16x2 DisplayMode = ((16 << 8) | 2)
	Mode20x2 DisplayMode = ((20 << 8) | 2)
	Mode20x4 DisplayMode = ((20 << 8) | 4)
	Mode8x1  DisplayMode = ((8 << 8) | 1)
	Mode8x2  DisplayMode = ((8 << 8) | 2)
	Mode12x2 DisplayMode = ((12 << 8) | 2)
	Mode16x1 DisplayMode = ((16 << 8) | 1) // type1 only, type2 handling is TODO
	Mode16x4 DisplayMode = ((16 << 8) | 4)
	Mode20x1 DisplayMode = ((20 << 8) | 1)
	Mode24x2 DisplayMode = ((24 << 8) | 2)
	Mode40x2 DisplayMode = ((40 << 8) | 2)
)

//-----------------------------------------------------------------------------

// build a font image from row column pixel data
func buildFontImage(font [fontGlyphs][glyphPixelWidth]byte) *ebiten.Image {
	img := ebiten.NewImage(fontGlyphs*glyphWidth, glyphHeight)
	for i := 0; i < fontGlyphs; i++ {
		for j := 0; j < glyphPixelWidth; j++ {
			pixelData := font[i][j]
			for k := 0; k < glyphPixelHeight; k++ {
				pixel := (pixelData & (1 << (glyphPixelHeight - 1 - k))) != 0
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

// build an image from row ordered pixel data (cgram style)
func buildImage(buf []byte) *ebiten.Image {
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
					vector.FillRect(img, x, y, dotSize, dotSize, color.RGBA{0, 0, 0, 255}, false)
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
const cgAdrMask = byte(cgRamSize - 1)

type Config struct {
	Mode           DisplayMode
	XBase, YBase   float64 // xy position
	XScale, YScale float64 // xy scale
}

type LCD struct {
	config        *Config // lcd configuration
	rows, cols    int     // displays rows and columns
	scrollOffset  int
	ramMode       bool             // which ram are we working with?
	displayEnable bool             // is the display enabled?
	dlFlag        bool             // interface data width (false = 4, true = 8)
	nFlag         bool             // number of display lines (false = 1, true = 2)
	fFlag         bool             // font selection (false = 5x8, true = 5x10)
	shiftMode     bool             // shift mode (entry mode set command)
	incMode       bool             // increment mode (entry mode set command)
	font          [2]*ebiten.Image // rom font atlas images
	cgFont        *ebiten.Image    // character generator font atlas image
	cgRam         [cgRamSize]byte  // character generator ram
	cgAdr         byte             // cgram address
	cgDirty       bool             // do we need to rebuild the cg glyphs in the font atlas?
	cursorFont    *ebiten.Image    // cursor font atlas
	cursorCount   int              // cursor flash counter
	cursorBlink   bool             // is the cursor blinking?
	cursorEnable  bool             // is the cursor enabled?
	cursorState   bool             // current cursor state
	rowAdr        []byte           // address of row start in ddram
	ddRam         [ddRamSize]byte  // display data ram
	ddAdr         byte             // ddram address
	img           *ebiten.Image    // unscaled lcd image
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

	// build an initial cgram font
	for i := range lcd.cgRam {
		lcd.cgRam[i] = 0x1f
	}
	lcd.cgFont = buildImage(lcd.cgRam[:])

	// build the cursor font
	lcd.cursorFont = buildImage(cursorData)

	// build an unscaled lcd image
	width := lcd.cols*glyphWidth + (lcd.cols-1)*xGap
	height := lcd.rows*glyphHeight + (lcd.rows-1)*yGap
	lcd.img = ebiten.NewImage(width, height)

	// work out the row addresses
	lcd.rowAdr = make([]byte, lcd.rows)
	if lcd.rows == 1 {
		lcd.rowAdr[0] = 0
	} else if lcd.rows == 2 {
		lcd.rowAdr[0] = 0
		lcd.rowAdr[1] = 0x40
	} else if lcd.rows == 4 {
		lcd.rowAdr[0] = 0
		lcd.rowAdr[1] = 0x40
		lcd.rowAdr[2] = 0x14
		lcd.rowAdr[3] = 0x54
	} else {
		return nil, errors.New("rows != 1,2,4")
	}

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
func (lcd *LCD) getGlyph(set int, code byte) *ebiten.Image {
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
		img = lcd.font[set]
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

// given a display address, work out the row and column.
func (lcd *LCD) getRowCol(adr byte) (int, int, bool) {
	for row := 0; row < lcd.rows; row++ {
		start := lcd.rowAdr[row]
		if adr >= start && adr < start+byte(lcd.cols) {
			return row, int(adr - start), true
		}
	}
	return 0, 0, false
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

	// create an unscaled lcd image
	lcd.img.Clear()
	for row := 0; row < lcd.rows; row++ {
		for col := 0; col < lcd.cols; col++ {
			// get the character
			code := lcd.ddRam[lcd.rowAdr[row]+byte(col)]
			glyph := lcd.getGlyph(0, code)
			// render the glyph to the lcd image
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(col*pitchX), float64(row*pitchY))
			lcd.img.DrawImage(glyph, op)
		}
	}

	// draw the cursor
	if lcd.cursorEnable {
		row, col, ok := lcd.getRowCol(lcd.ddAdr)
		if ok {
			// render the cursor to the lcd image
			cursor := lcd.getCursor(lcd.cursorBlink && lcd.cursorState)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(col*pitchX), float64(row*pitchY))
			lcd.img.DrawImage(cursor, op)
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

	// update the cursor state
	if lcd.cursorCount == cursorBlinkPeriod {
		lcd.cursorState = !lcd.cursorState
		lcd.cursorCount = 0
	}
	lcd.cursorCount += 1

	// rebuild the cgram font if it has been changed
	if lcd.cgDirty {
		lcd.cgFont = buildImage(lcd.cgRam[:])
		lcd.cgDirty = false
	}
}

//-----------------------------------------------------------------------------
