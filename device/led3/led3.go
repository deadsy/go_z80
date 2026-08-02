//-----------------------------------------------------------------------------
/*

3-channel LED emulation

This is an RGB LED, but the internal LEDs are controlled as on or off rather
than a PWM signal. This gives a total of 2x2x2 = 8 colors.

*/
//-----------------------------------------------------------------------------

package led3

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//-----------------------------------------------------------------------------

const fadeCount = 2

const (
	stateOff    = iota // led is OFF
	stateOn            // led is ON
	stateFading        // led to turn off after fadeCount updates
)

type ledState struct {
	state int
	fade  int
}

func (s ledState) isOn() bool {
	return s.state != stateOff
}

func (s *ledState) control(on bool) {
	if on {
		s.state = stateOn
	} else {
		if s.state == stateOn {
			s.state = stateFading
			s.fade = fadeCount
		}
	}
}

//-----------------------------------------------------------------------------

type Type int

const (
	Round Type = iota
	Rectangle
)

type Config struct {
	Type          Type          // led type
	X, Y          float32       // xy position of display on screen
	Radius        float32       // radius (round only)
	Width, Height float32       // width/height size (rectangular only)
	Colors        [4]color.RGBA // r,g,b,off color
}

type LED3 struct {
	config Config
	state  [3]ledState // current state
}

//-----------------------------------------------------------------------------

func New(cfg Config) (*LED3, error) {

	if cfg.Type == Round && cfg.Radius <= 0 {
		return nil, errors.New("bad radius")
	}
	if cfg.Type == Rectangle && cfg.Width <= 0 {
		return nil, errors.New("bad width")
	}
	if cfg.Type == Rectangle && cfg.Height <= 0 {
		return nil, errors.New("bad height")
	}
	return &LED3{
		config: cfg,
	}, nil
}

// Control the LED (called from the IO layer)
func (led *LED3) Control(r, g, b bool) {
	led.state[0].control(r)
	led.state[1].control(g)
	led.state[2].control(b)
}

//-----------------------------------------------------------------------------
// ebiten functions

func colorAdd(a, b color.RGBA) color.RGBA {
	return color.RGBA{
		R: min(255, a.R+b.R),
		G: min(255, a.G+b.G),
		B: min(255, a.B+b.B),
		A: min(255, a.A+b.A),
	}
}

// Draw the LED (called from ebiten draw function)
func (led *LED3) Draw(screen *ebiten.Image) {
	cfg := &led.config

	// work out the color
	c := color.RGBA{}
	allOff := true
	for i := range led.state {
		if led.state[i].isOn() {
			c = colorAdd(c, cfg.Colors[i])
			allOff = false
		}
	}
	if allOff {
		// give it the off color
		c = cfg.Colors[3]
	}

	switch cfg.Type {
	case Round:
		vector.FillCircle(screen, cfg.X, cfg.Y, cfg.Radius, c, true)
	case Rectangle:
		vector.FillRect(screen, cfg.X, cfg.Y, cfg.Width, cfg.Height, c, true)
	}
}

// Update the LED logic (called from ebiten update)
func (led *LED3) Update() {
	for i := range led.state {
		s := &led.state[i]
		// Fade the led to the off state.
		if (s.state == stateFading) && (s.fade > 0) {
			s.fade -= 1
			if s.fade == 0 {
				s.state = stateOff
			}
		}
	}
}

//-----------------------------------------------------------------------------
