//-----------------------------------------------------------------------------
/*

TEC-1 Emulation

6 digit, 8 segment display

*/
//-----------------------------------------------------------------------------

package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//-----------------------------------------------------------------------------
// Segment color definitions

var (
	bgColor = color.RGBA{15, 15, 15, 255} // Dark background
	segOff  = color.RGBA{40, 10, 10, 255} // Dim red for unlit segments
	segOn   = color.RGBA{255, 0, 0, 255}  // Bright red for lit segments
)

func segColor(on bool) color.RGBA {
	if on {
		return segOn
	}
	return segOff
}

//-----------------------------------------------------------------------------

// see https://share.google/zEdJE5XEsoYcyDqK4
// nominally 90:122 for x:y proportions
// scaled to 0..1 for both dimensions

var p0 = [2]float32{0.3738764264697234, 0.9163934426229509}
var p1 = [2]float32{0.30766789354782287, 0.8754098360655738}
var p2 = [2]float32{0.24145936062592246, 0.8344262295081968}
var p3 = [2]float32{0.35428453972433843, 0.8344262295081968}
var p4 = [2]float32{0.825177142863388, 0.9163934426229509}
var p5 = [2]float32{0.8717937890399037, 0.8754098360655738}
var p6 = [2]float32{0.805585256118003, 0.8344262295081968}
var p7 = [2]float32{0.9184104352164192, 0.8344262295081968}
var p8 = [2]float32{0.7354463015695245, 0.5409836065573771}
var p9 = [2]float32{0.8482714806679408, 0.5409836065573771}
var p10 = [2]float32{0.7820629477460402, 0.5}
var p11 = [2]float32{0.7158544148241397, 0.45901639344262296}
var p12 = [2]float32{0.828679593922556, 0.45901639344262296}
var p13 = [2]float32{0.28414558517586025, 0.5409836065573771}
var p14 = [2]float32{0.17132040607744412, 0.5409836065573771}
var p15 = [2]float32{0.21793705225395957, 0.5}
var p16 = [2]float32{0.1517285193320591, 0.45901639344262296}
var p17 = [2]float32{0.2645536984304752, 0.4590163934426229}
var p18 = [2]float32{0.08158956478358079, 0.16557377049180327}
var p19 = [2]float32{0.1944147438819969, 0.16557377049180327}
var p20 = [2]float32{0.12820621096009635, 0.12459016393442623}
var p21 = [2]float32{0.6457154602756613, 0.16557377049180327}
var p22 = [2]float32{0.7585406393740776, 0.16557377049180327}
var p23 = [2]float32{0.6923321064521769, 0.12459016393442623}
var p24 = [2]float32{0.6261235735302764, 0.08360655737704918}
var p25 = [2]float32{0.1748228571366119, 0.08360655737704918}
var p26 = [2]float32{0.8692110012368004, 0.12459016393442623}

var segA = [][2]float32{p1, p0, p4, p5, p6, p3}
var segB = [][2]float32{p5, p7, p9, p10, p8, p6}
var segC = [][2]float32{p10, p12, p22, p23, p21, p11}
var segD = [][2]float32{p20, p19, p21, p23, p24, p25}
var segE = [][2]float32{p15, p17, p19, p20, p18, p16}
var segF = [][2]float32{p1, p3, p13, p15, p14, p2}
var segG = [][2]float32{p15, p13, p8, p10, p11, p17}
var segDP = [][2]float32{p26}

const dpRadius = (5.0 / 122.0)

func xyScale(x float32) float32 {
	return x * (122.0 / 90.0)
}

//-----------------------------------------------------------------------------
// digits

const fadeCount = 2

type digitState int

const (
	digitOff digitState = iota
	digitOn
	digitFading
)

type digit struct {
	state digitState // current state
	fade  int        // fade count
	val   byte       // segment value
}

func (d *digit) set(val byte) {
	if val == 0 {
		d.state = digitOff
	} else {
		d.state = digitOn
	}
	d.val = val
}

func (d *digit) off() {
	if d.state == digitOn {
		d.state = digitFading
		d.fade = fadeCount
	}
}

// Is the digit on (or fading)?
func (d *digit) isOn() bool {
	return d.state != digitOff
}

// Digit update routine
func (d *digit) update() {
	// Fade the digit to the off state.
	if (d.state == digitFading) && (d.fade > 0) {
		d.fade -= 1
		if d.fade == 0 {
			d.state = digitOff
		}
	}
}

//-----------------------------------------------------------------------------

const numDigits = 6

// bit to segment mapping
func aOn(val byte) bool  { return val&(1<<0) != 0 }
func fOn(val byte) bool  { return val&(1<<1) != 0 }
func gOn(val byte) bool  { return val&(1<<2) != 0 }
func bOn(val byte) bool  { return val&(1<<3) != 0 }
func dpOn(val byte) bool { return val&(1<<4) != 0 }
func cOn(val byte) bool  { return val&(1<<5) != 0 }
func eOn(val byte) bool  { return val&(1<<6) != 0 }
func dOn(val byte) bool  { return val&(1<<7) != 0 }

type Display struct {
	digits  [numDigits]digit // digit states
	texture *ebiten.Image
	// position and scale
	xBase, yBase   float32 // xy position of display on screen
	xScale, yScale float32 // xy size of digit
	xGap0          float32 // gaps between digits
	xGap1          float32 // gap between address /data digits

}

const digitSize = float32(55.0)

func newDisplay() *Display {
	d := &Display{
		texture: ebiten.NewImage(1, 1),
		// position and scale
		xBase:  362.0,
		yBase:  665.0,
		xScale: digitSize,
		yScale: xyScale(digitSize),
		xGap0:  24.0,
		xGap1:  14.0,
	}
	d.texture.Fill(color.White)
	return d
}

func (d *Display) xMap(digit int, x float32) float32 {
	gap := float32(digit) * d.xGap0
	if digit >= 4 {
		gap += d.xGap1
	}
	return d.xBase + gap + d.xScale*(float32(digit)+x)
}

func (d *Display) yMap(y float32) float32 {
	return d.yBase + d.yScale*(1.0-y)
}

func (d *Display) drawSegment(screen *ebiten.Image, digit int, on bool, point [][2]float32) {
	c := segColor(on)
	const cScale = float32(1.0 / 255.0)
	// Setup vertices for the segment polygon (6-sided)
	var vertices [6]ebiten.Vertex
	for i := range vertices {
		vertices[i] = ebiten.Vertex{
			DstX:   d.xMap(digit, point[i][0]),
			DstY:   d.yMap(point[i][1]),
			SrcX:   0,
			SrcY:   0,
			ColorR: cScale * float32(c.R),
			ColorG: cScale * float32(c.G),
			ColorB: cScale * float32(c.B),
			ColorA: cScale * float32(c.A),
		}
	}
	indices := []uint16{0, 2, 1, 0, 3, 2, 0, 4, 3, 0, 5, 4}
	screen.DrawTriangles(vertices[:], indices, d.texture, nil)
}

func (d *Display) drawSegmentDP(screen *ebiten.Image, digit int, on bool) {
	c := segColor(on)
	x := d.xMap(digit, segDP[0][0])
	y := d.yMap(segDP[0][1])
	r := d.yScale * dpRadius
	vector.FillCircle(screen, x, y, r, c, true)
}

func (d *Display) drawDigit(screen *ebiten.Image, digit int) {
	on := d.digits[digit].isOn()
	val := d.digits[digit].val
	d.drawSegment(screen, digit, on && aOn(val), segA)
	d.drawSegment(screen, digit, on && bOn(val), segB)
	d.drawSegment(screen, digit, on && cOn(val), segC)
	d.drawSegment(screen, digit, on && dOn(val), segD)
	d.drawSegment(screen, digit, on && eOn(val), segE)
	d.drawSegment(screen, digit, on && fOn(val), segF)
	d.drawSegment(screen, digit, on && gOn(val), segG)
	d.drawSegmentDP(screen, digit, on && dpOn(val))
}

//-----------------------------------------------------------------------------

func (d *Display) getWidth() int {
	return int(math.Ceil(float64(d.xMap(numDigits, 0))))
}

func (d *Display) getHeight() int {
	return int(math.Ceil(float64(d.yMap(0))))
}

func (d *Display) enable(digit, segment uint8) {
	if digit == 0 {
		// all digits are off
		for i := 0; i < numDigits; i++ {
			d.digits[i].off()
		}
		return
	}
	if (digit & 0x20) != 0 {
		d.digits[0].set(segment)
	}
	if (digit & 0x10) != 0 {
		d.digits[1].set(segment)
	}
	if (digit & 0x08) != 0 {
		d.digits[2].set(segment)
	}
	if (digit & 0x04) != 0 {
		d.digits[3].set(segment)
	}
	if (digit & 0x02) != 0 {
		d.digits[4].set(segment)
	}
	if (digit & 0x01) != 0 {
		d.digits[5].set(segment)
	}
}

//-----------------------------------------------------------------------------

// display draw function (called in game draw)
func (d *Display) draw(screen *ebiten.Image) {
	for i := 0; i < numDigits; i++ {
		d.drawDigit(screen, i)
	}
}

// periodic update function (called in game update)
func (d *Display) update() {
	for i := 0; i < numDigits; i++ {
		d.digits[i].update()
	}
}

//-----------------------------------------------------------------------------
