//-----------------------------------------------------------------------------
/*

ST7920 LCD Driver Emulation

Resources:

https://www.waveshare.com/datasheet/LCD_en_PDF/ST7920.pdf
https://github.com/teeminus/ST7920Emulator/tree/master

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

const pixelWidth = 128
const pixelHeight = 64

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

//-----------------------------------------------------------------------------

type Config struct {
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

	addressX byte
	addressY byte

	cgRam [64][2]byte  // character generator ram
	ddRam [4][32]byte  // display data ram
	gdRam [64][16]byte // graphics data ram (size?)
}

func New(cfg Config) (*LCD, error) {
	lcd := &LCD{cfg: cfg}

	// build an lcd image
	width := (2 * cfg.XBorder) + pixelWidth
	height := (2 * cfg.YBorder) + pixelHeight
	lcd.img = ebiten.NewImage(width, height)

	lcd.reset()

	return lcd, nil
}

//-----------------------------------------------------------------------------

func (lcd *LCD) reset() {

	// init enums
	lcd.lastCommand = ctNone
	lcd.dataTarget = dtNone

	// init serial byte decoding
	lcd.syncByteType = sbtNone
	lcd.dataNibbleIdx = 0
	lcd.dataByte = 0

	// Init variables set by commands
	lcd.enableVerticalScroll = false
	lcd.extendedMode = false
	lcd.graphicMode = false
	lcd.addressX = 0
	lcd.addressY = 0

	// Init rams
	for i := 0; i < 64; i++ {
		for j := 0; j < 2; j++ {
			lcd.cgRam[i][j] = 0
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 32; j++ {
			lcd.ddRam[i][j] = 0
		}
	}
	for i := 0; i < 64; i++ {
		for j := 0; j < 16; j++ {
			lcd.gdRam[i][j] = 0
		}
	}
}

//-----------------------------------------------------------------------------

// write command register
func (lcd *LCD) WriteCommand(cmd byte) {
	log.Printf("st7920.WriteCommand 0x%02x", cmd)
	// Check highest byte set
	if (cmd & 0b10000000) > 0 { // Set DDRAM/Grapic RAM address
		// Check for extended mode
		if lcd.extendedMode {
			// Check if graphic mode is enabled
			if lcd.graphicMode {
				// Check if current byte is for Y
				if lcd.lastCommand != ctGdramAddress {
					// First byte is Y
					lcd.addressY = cmd & 0b111111
					// Store command
					lcd.lastCommand = ctGdramAddress
				} else {
					// Second byte is X
					lcd.addressX = (cmd & 0b111) * 2
					// Check if we need to move the y cursor to the second half of the display
					if (cmd & 0b1000) > 0 {
						lcd.addressY = (lcd.addressY + 32) & 0b111111
					}
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
	} else if (cmd & 0b1000000) > 0 { // Set CGRAM/IRAM/SCROLL address
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
	} else if (cmd & 0b100000) > 0 { // (Extended) Function set
		// Check for extended instruction set
		lcd.extendedMode = cmd&0b100 != 0
		// Check for graphic mode flag
		if lcd.extendedMode {
			// Get graphic mode flag
			lcd.graphicMode = cmd&0b10 != 0
		}
		// Store command
		lcd.lastCommand = ctFunctionSet
	} else if (cmd & 0b10000) > 0 { // Cursor/Display control
		// TODO
		// Store command
		lcd.lastCommand = ctCursorControl
	} else if (cmd & 0b1000) > 0 { // Display on/off
		// Store command
		lcd.lastCommand = ctDisplayControl
	} else if (cmd & 0b100) > 0 { // Entry mode / Reverse
		// TODO Log ignoring of command
		if lcd.extendedMode {
			// Store command
			lcd.lastCommand = ctReverse
		} else {
			// Store command
			lcd.lastCommand = ctEntryMode
		}
	} else if (cmd & 0b10) > 0 { // Home / Scroll or ram address select
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
	} else if (cmd & 0b1) > 0 { // Clear/Stand by
		// Check for extended mode
		if lcd.extendedMode {
			// Store command
			lcd.lastCommand = ctStandBy
		} else {
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
	log.Printf("st7920.ReadCommand")
	// TODO
	return 0
}

//-----------------------------------------------------------------------------

func showByte(x, y byte) {
}

// write data register
func (lcd *LCD) WriteData(val byte) {
	log.Printf("st7920.WriteData 0x%02x", val)

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
		// Set byte
		lcd.gdRam[lcd.addressY][lcd.addressX] = bits.Reverse8(val)
		showByte(lcd.addressX, lcd.addressY)

		// Increase address
		if lcd.addressX < 15 {
			lcd.addressX += 1
		} else {
			lcd.addressX = 0
		}
	} else if lcd.dataTarget == dtDDRAM {
		// Get current byte
		tmp := lcd.ddRam[lcd.addressY/16][lcd.addressX]

		// Update DDRAM
		lcd.ddRam[lcd.addressY/16][lcd.addressX] = val

		// Check if chargen char was requested
		if ((lcd.addressX & 1) > 0) && (lcd.ddRam[lcd.addressY/16][lcd.addressX&0b1110] == 0) {
			// Draw char
			for y := lcd.addressY; y < lcd.addressY+16; y++ {
				showByte(lcd.addressX-1, y)
				showByte(lcd.addressX, y)
			}
		} else if (val > 0) && (val <= 0x7F) { // Check for halfsize font
			// Draw char
			for y := lcd.addressY; y < lcd.addressY+16; y++ {
				showByte(lcd.addressX, y)
			}

			// Check if chargen char has been overwritten
			if (tmp == 0) && ((lcd.addressX & 0b1) == 0) {
				// Clear second char of DDRAM
				x := lcd.addressX + 1
				lcd.ddRam[lcd.addressY/16][x] = 0x20
				// Update display
				for y := lcd.addressY; y < lcd.addressY+16; y++ {
					showByte(x, y)
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
	cfg := &lcd.cfg

	// clear the lcd image
	lcd.img.Clear()
	lcd.img.Fill(cfg.BackgroundColor)

	// render the lcd image to the screen (with scaling)
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
