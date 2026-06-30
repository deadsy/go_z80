//-----------------------------------------------------------------------------
/*

TEC-1 Emulation

6 digit, 7 segment display

*/
//-----------------------------------------------------------------------------

package main

import (
	"image/color"

	"github.com/deadsy/go_z80/device/seven_segment"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------
// Segment color definitions

var (
	segOff = color.RGBA{40, 10, 10, 255} // Dim red for unlit segments
	segOn  = color.RGBA{255, 0, 0, 255}  // Bright red for lit segments
)

//-----------------------------------------------------------------------------

func (d *Display) xMap(digit int) float32 {
	gap := float32(digit) * d.xGap0
	if digit >= 4 {
		gap += d.xGap1
	}
	return d.xBase + gap + d.xScale*float32(digit)
}

//-----------------------------------------------------------------------------

const numDigits = 6

type Display struct {
	digit [numDigits]*seven_segment.Display
	// position and scale
	xBase, yBase   float32 // xy position of display on screen
	xScale, yScale float32 // xy size of digit
	xGap0          float32 // gaps between digits
	xGap1          float32 // gap between address /data digits
}

const digitSize = float32(55.0)

func newDisplay() *Display {
	d := &Display{
		xBase:  362.0,
		yBase:  665.0,
		xScale: digitSize,
		yScale: seven_segment.XYScale(digitSize),
		xGap0:  24.0,
		xGap1:  14.0,
	}
	for i := 0; i < numDigits; i++ {
		k := seven_segment.Config{
			SegmentBit: [8]int{0, 3, 5, 7, 6, 1, 2, 4},
			ColorOn:    segOn,
			ColorOff:   segOff,
			XBase:      d.xMap(i),
			YBase:      d.yBase,
			XScale:     d.xScale,
			YScale:     d.yScale,
		}
		d.digit[i] = seven_segment.New(&k)
	}
	return d
}

//-----------------------------------------------------------------------------

func (d *Display) enable(digit, segment uint8) {
	if digit == 0 {
		// all digits are off
		for i := 0; i < numDigits; i++ {
			d.digit[i].Off()
		}
		return
	}
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

//-----------------------------------------------------------------------------

// display draw function (called in game draw)
func (d *Display) draw(screen *ebiten.Image) {
	for i := 0; i < numDigits; i++ {
		d.digit[i].Draw(screen)
	}
}

// periodic update function (called in game update)
func (d *Display) update() {
	for i := 0; i < numDigits; i++ {
		d.digit[i].Update()
	}
}

//-----------------------------------------------------------------------------
