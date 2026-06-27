//-----------------------------------------------------------------------------
/*

TEC-1 Emulation

Speaker Activity LED

*/
//-----------------------------------------------------------------------------

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//-----------------------------------------------------------------------------
// LED color definitions

var (
	ledColorOff = color.RGBA{0, 0, 0, 0}     // transparent
	ledColorOn  = color.RGBA{0, 255, 0, 255} // bright green for lit led
)

func ledColor(on bool) color.RGBA {
	if on {
		return ledColorOn
	}
	return ledColorOff
}

//-----------------------------------------------------------------------------

const ledFadeCount = 2

type ledState int

const (
	ledOff ledState = iota
	ledOn
	ledFading
)

type LED struct {
	state ledState // curretn state
	fade  int
	// psition and size
	xBase, yBase float32 // xy position of display on screen
	radius       float32
}

func newLED() *LED {
	l := &LED{
		xBase:  589.0,
		yBase:  600.5,
		radius: 13.0,
	}
	return l
}

//-----------------------------------------------------------------------------

func (l *LED) control(state bool) {
	if state {
		l.state = ledOn
	} else {
		if l.state == ledOn {
			l.state = ledFading
			l.fade = ledFadeCount
		}
	}
}

//-----------------------------------------------------------------------------

// display draw function (called in game draw)
func (l *LED) draw(screen *ebiten.Image) {
	on := l.state != ledOff
	c := ledColor(on)
	vector.FillCircle(screen, l.xBase, l.yBase, l.radius, c, true)
}

// periodic update function (called in game update)
func (l *LED) update() {
	// Fade the led to the off state.
	if (l.state == ledFading) && (l.fade > 0) {
		l.fade -= 1
		if l.fade == 0 {
			l.state = ledOff
		}
	}
}

//-----------------------------------------------------------------------------
