//-----------------------------------------------------------------------------
/*

TEC-1 Display Emulation

6 digit, 7 segment display

4 digits are used for the address (uint16), 2 are used for the value (uint8).
The display is controlled by displaying a single digit at a time (multiplexing).

*/
//-----------------------------------------------------------------------------

package sixdigit

import (
	"image/color"

	"github.com/deadsy/go_z80/device/sevseg"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------
// Segment color definitions

var (
	segOff = color.RGBA{40, 10, 10, 255} // Dim red for unlit segments
	segOn  = color.RGBA{255, 0, 0, 255}  // Bright red for lit segments
)

//-----------------------------------------------------------------------------

func (dc *Config) xMap(digit int) float32 {
	gap := float32(digit) * dc.XGap0
	if digit >= 4 {
		gap += dc.XGap1
	}
	return dc.XBase + gap + dc.XScale*float32(digit)
}

//-----------------------------------------------------------------------------

const numDigits = 6

type Config struct {
	XBase, YBase   float32 // xy position of display on screen
	XScale, YScale float32 // xy size of digit
	XGap0          float32 // gaps between digits
	XGap1          float32 // gap between address/data digits
}

type Display struct {
	config *Config                    // display configuration
	digit  [numDigits]*sevseg.Display // digit state
}

func New(k *Config) *Display {
	d := &Display{
		config: k,
	}
	for i := 0; i < numDigits; i++ {
		k := sevseg.Config{
			SegmentBit: [8]int{0, 3, 5, 7, 6, 1, 2, 4},
			ColorOn:    segOn,
			ColorOff:   segOff,
			XBase:      d.config.xMap(i),
			YBase:      d.config.YBase,
			XScale:     d.config.XScale,
			YScale:     d.config.YScale,
		}
		d.digit[i] = sevseg.New(&k)
	}
	return d
}

// Enable a given digit with a specific value.
func (d *Display) Enable(digit, segment uint8) {
	if digit == 0 {
		// all digits are off
		for i := 0; i < numDigits; i++ {
			d.digit[i].Off()
		}
		return
	}
	// Digit bit mapping is governed by TEC1 hardware.
	if (digit & 0x20) != 0 {
		d.digit[0].Set(segment)
	}
	if (digit & 0x10) != 0 {
		d.digit[1].Set(segment)
	}
	if (digit & 0x08) != 0 {
		d.digit[2].Set(segment)
	}
	if (digit & 0x04) != 0 {
		d.digit[3].Set(segment)
	}
	if (digit & 0x02) != 0 {
		d.digit[4].Set(segment)
	}
	if (digit & 0x01) != 0 {
		d.digit[5].Set(segment)
	}
}

// Draw the display (called from ebiten draw function)
func (d *Display) Draw(screen *ebiten.Image) {
	for i := 0; i < numDigits; i++ {
		d.digit[i].Draw(screen)
	}
}

// Update the display logic (called from ebiten update)
func (d *Display) Update() {
	for i := 0; i < numDigits; i++ {
		d.digit[i].Update()
	}
}

//-----------------------------------------------------------------------------
