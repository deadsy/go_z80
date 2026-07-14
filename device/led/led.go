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

const fadeCount = 2

type ledState int

const (
	stateOff    ledState = iota // led is OFF
	stateOn                     // led is ON
	stateFading                 // led to turn off N updates
)

type Type int

const (
	Round Type = iota
	Rectangle
)

type Config struct {
	Type          Type       // led type
	X, Y          float32    // xy position of display on screen
	Radius        float32    // radius (round only)
	Width, Height float32    // width/height size (rectangular only)
	On, Off       color.RGBA // on/off colors
}

type LED struct {
	config *Config
	state  ledState // current state
	fade   int
}

func New(cfg *Config) (*LED, error) {
	return &LED{
		config: cfg,
	}, nil
}

//-----------------------------------------------------------------------------

// Control the LED (called from the IO layer)
func (led *LED) Control(state bool) {
	if state {
		led.state = stateOn
	} else {
		if led.state == stateOn {
			led.state = stateFading
			led.fade = fadeCount
		}
	}
}

// Draw the LED (called from ebiten draw function)
func (led *LED) Draw(screen *ebiten.Image) {
	cfg := led.config
	// work out the color
	var color color.RGBA
	if led.state != stateOff {
		color = cfg.On
	} else {
		color = cfg.Off
	}
	switch cfg.Type {
	case Round:
		vector.FillCircle(screen, cfg.X, cfg.Y, cfg.Radius, color, true)
	case Rectangle:
		vector.FillRect(screen, cfg.X, cfg.Y, cfg.Width, cfg.Height, color, true)
	}
}

// Update the LED logic (called from ebiten update)
func (led *LED) Update() {
	// Fade the led to the off state.
	if (led.state == stateFading) && (led.fade > 0) {
		led.fade -= 1
		if led.fade == 0 {
			led.state = stateOff
		}
	}
}

//-----------------------------------------------------------------------------
