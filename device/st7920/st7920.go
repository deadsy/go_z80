//-----------------------------------------------------------------------------
/*

ST7920 LCD Driver Emulation

Resources:

https://www.waveshare.com/datasheet/LCD_en_PDF/ST7920.pdf
https://github.com/teeminus/ST7920Emulator/tree/master

Note:

The st7920 can be used to drive LCD glass with various displayed resolutions.
In this case we are (only) emulating a 128 x 64 display.

The st7920 has an internal graphics memory (gdram) of 256 x 64.

The gdram to display mapping is as follows:

a b   maps to   a
c d             b

a = rows 0–31, left 128 bits -> display top half
b = rows 0–31, right 128 bits -> display bottom half
c and d (rows 32–63) are not displayed on a 128×64 panel — that GDRAM area exists but is off-screen.

*/
//-----------------------------------------------------------------------------

package st7920

import (
	"image/color"
	"log"
	"math/bits"

	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

const displayWidth = 128
const displayHeight = 64

//-----------------------------------------------------------------------------

type lcdSyncByteType int

const (
	sbtNone lcdSyncByteType = iota
	sbtCommand
	sbtData
)

type lcdCommandType int

const (
	ctNone lcdCommandType = iota
	ctDdramAddress
	ctCgramAddress
	ctFunctionSet
	ctCursorControl
	ctDisplayControl
	ctEntryMode
	ctHome
	ctClear
	ctGdramAddress
	ctIramAddress
	ctReverse
	ctScrollSelect
	ctStandBy
)

type lcdDataTarget int

const (
	dtNone lcdDataTarget = iota
	dtCGRAM
	dtDDRAM
	dtGDRAM
)

func (t lcdDataTarget) String() string {
	return []string{"None", "CGRAM", "DDRAM", "GDRAM"}[t]
}

//-----------------------------------------------------------------------------

const glyphWidth = 8
const glyphHeight = 16

// build an atlas image from font data
func buildImage(buf []byte, c color.RGBA) *ebiten.Image {
	nGlyphs := len(buf) >> 4 // 16 bytes per glyph
	img := ebiten.NewImage(nGlyphs*glyphWidth, glyphHeight)
	for i := 0; i < nGlyphs; i++ {
		for j := 0; j < glyphHeight; j++ {
			pixelData := buf[(i<<4)+j]
			for k := 0; k < glyphWidth; k++ {
				pixel := (pixelData & (1 << k)) != 0
				if pixel {
					img.Set((i*glyphWidth)+k, j, c)
				}
			}
		}
	}
	return img
}

//-----------------------------------------------------------------------------

type Config struct {
	Enable           bool    // is the lcd enabled?
	XBase, YBase     float64 // xy position
	XScale, YScale   float64 // xy scale
	XBorder, YBorder int     // border around graphics field
	BackgroundColor  color.RGBA
	PixelColor       color.RGBA
}

type LCD struct {
	cfg Config // lcd configuration

	// images
	font *ebiten.Image // font atlas
	img  *ebiten.Image // lcd image

	lastCommand lcdCommandType
	dataTarget  lcdDataTarget

	// serial interface
	syncByteType  lcdSyncByteType
	dataNibbleIdx byte
	dataByte      byte

	enableVerticalScroll bool
	extendedMode         bool
	graphicMode          bool
	displayOn            bool
	cursorOn             bool
	blinkOn              bool

	reverseMode byte

	addressX byte // column index 0..31 (bytes)
	addressY byte // row counter 0..63

	cgRam [64][2]byte  // character generator ram
	ddRam [4][32]byte  // display data ram
	gdRam [64][32]byte // graphics data ram (256x64)
}

func New(cfg Config) (*LCD, error) {
	lcd := &LCD{cfg: cfg}
	if !cfg.Enable {
		return lcd, nil
	}

	// load the 8x16 font atlas
	lcd.font = buildImage(font8x16, cfg.PixelColor)

	// build an lcd image
	width := (2 * cfg.XBorder) + displayWidth
	height := (2 * cfg.YBorder) + displayHeight
	lcd.img = ebiten.NewImage(width, height)

	return lcd, nil
}

//-----------------------------------------------------------------------------

// write command register
func (lcd *LCD) WriteCommand(cmd byte) {
	if !lcd.cfg.Enable {
		return
	}
	//log.Printf("st7920.WriteCommand 0x%02x", cmd)
	// Check highest byte set
	if (cmd & 0x80) != 0 { // Set DDRAM/Graphic RAM address
		// Check for extended mode
		if lcd.extendedMode {
			// Check if graphic mode is enabled
			if lcd.graphicMode {
				// Check if current byte is for Y
				if lcd.lastCommand != ctGdramAddress {
					// First byte is Y
					lcd.addressY = cmd & 0x3f // only 6 bits used 0..63
					// Store command
					lcd.lastCommand = ctGdramAddress
				} else {
					// Second byte is X
					// The cmd value is a 0..15 index of a 16-bit word in the horizontal gdram buffer (256 bits wide)
					// We adjust it so we have a 0..30 index into a byte buffer.
					lcd.addressX = (cmd & 0x0f) << 1 // 0..30
					//log.Printf("st7920: gdRam address (%d,%d)", lcd.addressX, lcd.addressY)
					// Clear command
					lcd.lastCommand = ctNone
					// Set data target
					lcd.dataTarget = dtGDRAM
				}
			}
		} else {
			// Set DDRAM address
			lcd.addressX = (cmd & 0b111) * 2 // Lower 3 bytes, organized in 16-bit blocks
			switch (cmd >> 3) & 0b11 {       // Upper 2 bytes, 16 bit high font
			case 0:
				lcd.addressY = 0
			case 1:
				lcd.addressY = 32
			case 2:
				lcd.addressY = 16
			case 3:
				lcd.addressY = 48
			}
			// Store command
			lcd.lastCommand = ctDdramAddress
			// Set data target
			lcd.dataTarget = dtDDRAM
		}
	} else if (cmd & 0x40) != 0 { // Set CGRAM/IRAM/SCROLL address

		log.Printf("set cgram address")

		// Check for extended mode
		if lcd.extendedMode {
			// Store command
			lcd.lastCommand = ctIramAddress
		} else {
			// Check if vertical scroll mode is disabled
			if !lcd.enableVerticalScroll {
				// X address is used for byte indexing
				lcd.addressX = 0
				// Y address is character index
				lcd.addressY = cmd & 0b111111
				// Write target is CGRAM
				lcd.dataTarget = dtCGRAM
			}
			// Store command
			lcd.lastCommand = ctCgramAddress
		}
	} else if (cmd & 0x20) != 0 { // (Extended) Function set
		// Check for extended instruction set
		lcd.extendedMode = cmd&0x04 != 0
		log.Printf("st7920: extended mode %t", lcd.extendedMode)
		// Check for graphic mode flag
		if lcd.extendedMode {
			// Get graphic mode flag
			lcd.graphicMode = cmd&0x02 != 0
			log.Printf("st7920: graphic mode %t", lcd.graphicMode)
		}
		// Store command
		lcd.lastCommand = ctFunctionSet
	} else if (cmd & 0x10) != 0 { // Cursor/Display control

		log.Printf("cursor/display control")

		// TODO
		// Store command
		lcd.lastCommand = ctCursorControl
	} else if (cmd & 0x08) != 0 { // Display on/off
		lcd.displayOn = cmd&0x04 != 0
		lcd.cursorOn = cmd&0x02 != 0
		lcd.blinkOn = cmd&0x01 != 0
		log.Printf("st7920: display on %t", lcd.displayOn)
		log.Printf("st7920: cursor on %t", lcd.cursorOn)
		log.Printf("st7920: blink on %t", lcd.blinkOn)
		// Store command
		lcd.lastCommand = ctDisplayControl
	} else if (cmd & 0x04) != 0 { // Entry mode / Reverse
		if lcd.extendedMode {
			lcd.reverseMode = cmd & 0x03
			// Store command
			lcd.lastCommand = ctReverse
		} else {
			// TODO cursor move to right
			// Store command
			lcd.lastCommand = ctEntryMode
		}
	} else if (cmd & 0x02) != 0 { // Home/Scroll or ram address select

		log.Printf("home, scroll, ram address select")

		// Check for extended mode
		if lcd.extendedMode {
			// Get vertical scroll enable flag
			lcd.enableVerticalScroll = cmd&0b1 != 0
			// Store command
			lcd.lastCommand = ctScrollSelect
		} else {
			// Reset cursor
			lcd.addressX = 0
			lcd.addressY = 0
			// Store command
			lcd.lastCommand = ctHome
			// Update data target
			lcd.dataTarget = dtDDRAM
		}
	} else if (cmd & 0x01) != 0 { // Clear/Stand by
		// Check for extended mode
		if lcd.extendedMode {
			log.Printf("st7920: standby")
			// Store command
			lcd.lastCommand = ctStandBy
		} else {
			log.Printf("st7920: clear")
			// Clear DDRAM
			for y := 0; y < 4; y++ {
				for x := 0; x < 32; x++ {
					lcd.ddRam[y][x] = 0x20
				}
			}
			// Reset cursor
			lcd.addressX = 0
			lcd.addressY = 0
			// Store command
			lcd.lastCommand = ctClear
			// Update data target
			lcd.dataTarget = dtDDRAM
		}
	}
}

// read command register
func (lcd *LCD) ReadCommand() byte {
	if !lcd.cfg.Enable {
		return 0
	}
	log.Printf("st7920.ReadCommand")
	// TODO
	return 0
}

//-----------------------------------------------------------------------------

// write data register
func (lcd *LCD) WriteData(val byte) {
	if !lcd.cfg.Enable {
		return
	}
	//log.Printf("st7920.WriteData 0x%02x to %s[%d,%d]", val, lcd.dataTarget, lcd.addressX, lcd.addressY)

	// Check for data target
	if lcd.dataTarget == dtCGRAM {
		// Write data to CGRAM
		lcd.cgRam[lcd.addressY][lcd.addressX] = bits.Reverse8(val)

		// Increase address
		if lcd.addressX == 0 {
			lcd.addressX = 1
		} else {
			lcd.addressX = 0
			if lcd.addressY >= 63 {
				lcd.addressY = 0
			} else {
				lcd.addressY += 1
			}
		}
	} else if lcd.dataTarget == dtGDRAM {
		lcd.gdRam[lcd.addressY][lcd.addressX] = bits.Reverse8(val)
		// increment x address
		lcd.addressX = (lcd.addressX + 1) & 0x1f
	} else if lcd.dataTarget == dtDDRAM {
		// Get current byte
		tmp := lcd.ddRam[lcd.addressY/16][lcd.addressX]

		// Update DDRAM
		lcd.ddRam[lcd.addressY/16][lcd.addressX] = val

		// Check if chargen char was requested
		if ((lcd.addressX & 1) > 0) && (lcd.ddRam[lcd.addressY/16][lcd.addressX&0b1110] == 0) {
			// Draw char
			for y := lcd.addressY; y < lcd.addressY+16; y++ {
				//showByte(lcd.addressX-1, y)
				//showByte(lcd.addressX, y)
			}
		} else if (val > 0) && (val <= 0x7F) { // Check for halfsize font
			// Draw char
			for y := lcd.addressY; y < lcd.addressY+16; y++ {
				//showByte(lcd.addressX, y)
			}

			// Check if chargen char has been overwritten
			if (tmp == 0) && ((lcd.addressX & 0b1) == 0) {
				// Clear second char of DDRAM
				x := lcd.addressX + 1
				lcd.ddRam[lcd.addressY/16][x] = 0x20
				// Update display
				for y := lcd.addressY; y < lcd.addressY+16; y++ {
					//showByte(x, y)
				}
			}
		}

		// Update cursor
		if lcd.addressX < 15 {
			lcd.addressX += 1
		} else {
			lcd.addressX = 0
			switch lcd.addressY {
			case 0:
				lcd.addressY = 32
			case 16:
				lcd.addressY = 48
			case 32:
				lcd.addressY = 16
			case 48:
				lcd.addressY = 0
			}
		}
	}
}

// read data register
func (lcd *LCD) ReadData() byte {
	if !lcd.cfg.Enable {
		return 0
	}
	log.Printf("st7920.ReadData")
	// TODO
	return 0
}

//-----------------------------------------------------------------------------
// serial interface

func (lcd *LCD) parseSyncByte(data byte) bool {
	// Check for sync byte pattern
	if (data & 0b11111000) == 0b11111000 {
		if (data & 0b100) > 0 { // Check for R/W bit
			// Invalid sync byte
			lcd.syncByteType = sbtNone
		} else if (data & 0b10) > 0 { // Check for RS bit
			// Switch to data mode
			lcd.syncByteType = sbtData
			// Clear last command
			lcd.lastCommand = ctNone
		} else { // Command mode
			// Switch to command mode
			lcd.syncByteType = sbtCommand
			// Clear data target
			lcd.dataTarget = dtNone
		}
		// Reset data nibble flag
		lcd.dataNibbleIdx = 0
		// Found sync byte
		return true
	}
	// Not a sync byte
	return false
}

func (lcd *LCD) SerialData(data byte) {
	if !lcd.cfg.Enable {
		return
	}
	// Check if byte is not a sync byte
	if lcd.parseSyncByte(data) {
		// Reconstruct byte
		if lcd.dataNibbleIdx == 0 {
			// Store higher nibble
			lcd.dataByte = data & 0xF0
			// Update flag
			lcd.dataNibbleIdx = 1
		} else {
			// Store lower byte
			lcd.dataByte |= (data >> 4) & 0x0F
			// Clear flag
			lcd.dataNibbleIdx = 0
			// Check how to parse the byte
			if lcd.syncByteType == sbtCommand {
				// Parse command
				lcd.WriteCommand(lcd.dataByte)
			} else if lcd.syncByteType == sbtData {
				// Parse data
				lcd.WriteData(lcd.dataByte)
			}
		}
	}
}

//-----------------------------------------------------------------------------

// Draw the display (called from ebiten draw function)
func (lcd *LCD) Draw(screen *ebiten.Image) {
	if !lcd.cfg.Enable {
		return
	}

	cfg := &lcd.cfg

	// clear the lcd image
	lcd.img.Clear()
	lcd.img.Fill(cfg.BackgroundColor)

	if lcd.displayOn {
		// render gdram
		for i := 0; i < displayWidth>>3; i++ {
			for j := 0; j < displayHeight; j++ {
				var pixelData byte
				if j > 31 {
					// bottom half of display
					pixelData = lcd.gdRam[j&0x1f][i+16]
				} else {
					// top half of display
					pixelData = lcd.gdRam[j][i]
				}
				for k := 0; k < glyphWidth; k++ {
					pixel := (pixelData & (1 << k)) != 0
					if pixel {
						x := cfg.XBorder + (i * glyphWidth) + k
						y := cfg.YBorder + j
						lcd.img.Set(x, y, cfg.PixelColor)
					}
				}
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
	if !lcd.cfg.Enable {
		return
	}
}

//-----------------------------------------------------------------------------
