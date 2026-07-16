//-----------------------------------------------------------------------------
/*

RGB LED Emulation

Control an RGB led that is modulated with a digital PWM signal.

*/
//-----------------------------------------------------------------------------

package led

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//-----------------------------------------------------------------------------

type Type int

const (
	Round Type = iota
	Rectangle
)

type Config struct {
	Type          Type    // led type
	X, Y          float32 // xy position of display on screen
	Radius        float32 // radius (round only)
	Width, Height float32 // width/height size (rectangular only)
}

type RGB struct {
	config  *Config
	state   [3]bool // current state
	onTime  [3]uint // time spent in on state
	offTime [3]uint // time spent in off state
	duty    [3]byte // estimated duty cycle
}

func New(cfg *Config) (*RGB, error) {
	if cfg.Type == Round && cfg.Radius <= 0 {
		return nil, errors.New("bad radius")
	}
	if cfg.Type == Rectangle && cfg.Width <= 0 {
		return nil, errors.New("bad width")
	}
	if cfg.Type == Rectangle && cfg.Height <= 0 {
		return nil, errors.New("bad height")
	}
	return &RGB{
		config: cfg,
	}, nil
}

//-----------------------------------------------------------------------------

func (rgb *RGB) channel(channel int, state bool, cycles uint64) {
}

// Control the RGB (called from the IO layer)
func (rgb *RGB) Control(state byte, cycles uint64) {
	rgb.channel(0, state&(1<<0) != 0, cycles)
	rgb.channel(1, state&(1<<1) != 0, cycles)
	rgb.channel(2, state&(1<<2) != 0, cycles)
}

// Draw the LED (called from ebiten draw function)
func (rgb *RGB) Draw(screen *ebiten.Image) {
	cfg := rgb.config
	color := color.RGBA{rgb.duty[0], rgb.duty[1], rgb.duty[2], 255}
	switch cfg.Type {
	case Round:
		vector.FillCircle(screen, cfg.X, cfg.Y, cfg.Radius, color, true)
	case Rectangle:
		vector.FillRect(screen, cfg.X, cfg.Y, cfg.Width, cfg.Height, color, true)
	}
}

// Update the LED logic (called from ebiten update)
func (rgb *RGB) Update() {

	// Estimate PWM duty cycle for 3 channels

}

//-----------------------------------------------------------------------------
