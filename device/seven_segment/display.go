//-----------------------------------------------------------------------------
/*

Seven Segment Display Emulation

*/
//-----------------------------------------------------------------------------

package seven_segment

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//-----------------------------------------------------------------------------

// see https://share.google/zEdJE5XEsoYcyDqK4
// Use ./dump.py to get the coordinates from 7seg.dxf cad file.
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

//-----------------------------------------------------------------------------

func (d *Display) aOn() bool {
	return d.val&(1<<d.config.SegmentBit[0]) != 0
}
func (d *Display) bOn() bool {
	return d.val&(1<<d.config.SegmentBit[1]) != 0
}
func (d *Display) cOn() bool {
	return d.val&(1<<d.config.SegmentBit[2]) != 0
}
func (d *Display) dOn() bool {
	return d.val&(1<<d.config.SegmentBit[3]) != 0
}
func (d *Display) eOn() bool {
	return d.val&(1<<d.config.SegmentBit[4]) != 0
}
func (d *Display) fOn() bool {
	return d.val&(1<<d.config.SegmentBit[5]) != 0
}
func (d *Display) gOn() bool {
	return d.val&(1<<d.config.SegmentBit[6]) != 0
}
func (d *Display) dpOn() bool {
	return d.val&(1<<d.config.SegmentBit[7]) != 0
}

//-----------------------------------------------------------------------------

func (d *Display) xMap(x float32) float32 {
	return d.config.xBase + d.config.xScale*x
}

func (d *Display) yMap(y float32) float32 {
	return d.config.yBase + d.config.yScale*(1.0-y)
}

//-----------------------------------------------------------------------------

func (d *Display) segColor(on bool) color.RGBA {
	if on {
		return d.config.ColorOn
	}
	return d.config.ColorOff
}

//-----------------------------------------------------------------------------

func (d *Display) drawSegment(screen *ebiten.Image, on bool, point [][2]float32) {
	c := d.segColor(on)
	const cScale = float32(1.0 / 255.0)
	// Setup vertices for the segment polygon (6-sided)
	var vertices [6]ebiten.Vertex
	for i := range vertices {
		vertices[i] = ebiten.Vertex{
			DstX:   d.xMap(point[i][0]),
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

func (d *Display) drawSegmentDP(screen *ebiten.Image, on bool) {
	c := d.segColor(on)
	x := d.xMap(segDP[0][0])
	y := d.yMap(segDP[0][1])
	r := d.config.yScale * dpRadius
	vector.FillCircle(screen, x, y, r, c, true)
}

//-----------------------------------------------------------------------------

type digitState int

const (
	digitOff digitState = iota
	digitOn
	digitFading
)

// Is the digit on (or fading)?
func (d *Display) isOn() bool {
	return d.state != digitOff
}

//-----------------------------------------------------------------------------

type Config struct {
	SegmentBit     [8]int     // map a segment to a value bit (0..7)
	ColorOn        color.RGBA // segment on color
	ColorOff       color.RGBA // segment off color
	xBase, yBase   float32    // position
	xScale, yScale float32    // scale
}

type Display struct {
	config  *Config
	texture *ebiten.Image
	val     byte // segment bits
	state   digitState
	fade    int
}

func New(k *Config) *Display {
	d := &Display{
		config:  k,
		texture: ebiten.NewImage(1, 1),
	}
	d.texture.Fill(color.White)
	return d
}

// Draw the display (called from ebiten draw function)
func (d *Display) Draw(screen *ebiten.Image) {
	on := d.isOn()
	d.drawSegment(screen, on && d.aOn(), segA)
	d.drawSegment(screen, on && d.bOn(), segB)
	d.drawSegment(screen, on && d.cOn(), segC)
	d.drawSegment(screen, on && d.dOn(), segD)
	d.drawSegment(screen, on && d.eOn(), segE)
	d.drawSegment(screen, on && d.fOn(), segF)
	d.drawSegment(screen, on && d.gOn(), segG)
	d.drawSegmentDP(screen, on && d.dpOn())
}

// Update the display logic (called from ebiten update)
func (d *Display) Update() {
	// Fade the digit to the off state.
	if (d.state == digitFading) && (d.fade > 0) {
		d.fade -= 1
		if d.fade == 0 {
			d.state = digitOff
		}
	}
}

//-----------------------------------------------------------------------------
