//-----------------------------------------------------------------------------
/*

LED Emulation

*/
//-----------------------------------------------------------------------------

package led

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

type Config struct {
	XBase, YBase float32 // xy position of display on screen
	Radius       float32 // radius of LED
}

type LED struct {
	config *Config
	state  ledState // current state
	fade   int
}

func New(k *Config) (*LED, error) {
	return &LED{
		config: k,
	}, nil
}

// Control the LED (called from the IO layer)
func (l *LED) Control(state bool) {
	if state {
		l.state = ledOn
	} else {
		if l.state == ledOn {
			l.state = ledFading
			l.fade = ledFadeCount
		}
	}
}

// Draw the LED (called from ebiten draw function)
func (l *LED) Draw(screen *ebiten.Image) {
	on := l.state != ledOff
	c := ledColor(on)
	vector.FillCircle(screen, l.config.XBase, l.config.YBase, l.config.Radius, c, true)
}

// Update the LED logic (called from ebiten update)
func (l *LED) Update() {
	// Fade the led to the off state.
	if (l.state == ledFading) && (l.fade > 0) {
		l.fade -= 1
		if l.fade == 0 {
			l.state = ledOff
		}
	}
}

//-----------------------------------------------------------------------------
