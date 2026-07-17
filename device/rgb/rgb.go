//-----------------------------------------------------------------------------
/*

RGB LED Emulation

Control an RGB led modulated with a digital PWM signal.

*/
//-----------------------------------------------------------------------------

package rgb

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//-----------------------------------------------------------------------------

type pwmChannel struct {
	state   bool   // current state
	datum   uint64 // time datum
	onTime  uint64 // accumulated on time
	offTime uint64 // accumulated on time
	duty    byte   // estimated duty cycle
}

func (ch *pwmChannel) set(state bool, cycles uint64) {
	delta := cycles - ch.datum
	if ch.state {
		// accumulate on time
		ch.onTime += delta
	} else {
		// accumulate off time
		ch.offTime += delta
	}
	ch.datum = cycles
	ch.state = state
}

func (ch *pwmChannel) update(cycles uint64) {
	// final accumulation of time
	ch.set(ch.state, cycles)
	// calculate the duty cycle
	totalTime := ch.onTime + ch.offTime
	if totalTime > 0 {
		ch.duty = byte(255.0 * float32(ch.onTime) / float32(totalTime))
	} else {
		// should never happen...
		ch.duty = 0
	}
	// reset the accumulation times
	ch.onTime = 0
	ch.offTime = 0
}

//-----------------------------------------------------------------------------

type Type int

const (
	Round Type = iota
	Rectangle
)

type Config struct {
	Type          Type    // led type
	X, Y          float32 // xy position of led on screen
	Radius        float32 // radius (round only)
	Width, Height float32 // width/height size (rectangular only)
}

type RGB struct {
	cfg         Config
	channel     [3]pwmChannel
	updateCount int
}

func New(cfg Config) (*RGB, error) {
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
		cfg: cfg,
	}, nil
}

//-----------------------------------------------------------------------------

// Control the RGB (called from the IO layer)
func (rgb *RGB) Control(r, g, b bool, cycles uint64) {
	rgb.channel[0].set(r, cycles)
	rgb.channel[1].set(g, cycles)
	rgb.channel[2].set(b, cycles)
}

// Draw the LED (called from ebiten draw function)
func (rgb *RGB) Draw(screen *ebiten.Image) {

	// get the color
	r := rgb.channel[0].duty
	g := rgb.channel[1].duty
	b := rgb.channel[2].duty
	a := byte(200)
	if r == 0 && g == 0 && b == 0 {
		// off (transparent)
		a = 0
	}
	color := color.RGBA{r, g, b, a}

	cfg := rgb.cfg
	switch cfg.Type {
	case Round:
		vector.FillCircle(screen, cfg.X, cfg.Y, cfg.Radius, color, true)
	case Rectangle:
		vector.FillRect(screen, cfg.X, cfg.Y, cfg.Width, cfg.Height, color, true)
	}
}

// Update the LED logic (called from ebiten update)
func (rgb *RGB) Update(cycles uint64) {
	// note: The PWM period is around 120 Hz, giving
	// around 2 PWM cycles per update time (60Hz).
	// We wait 2 update times to get a more stable
	// duty cycle measurement.
	rgb.updateCount += 1
	if rgb.updateCount&1 == 0 {
		rgb.channel[0].update(cycles)
		rgb.channel[1].update(cycles)
		rgb.channel[2].update(cycles)
	}
}

//-----------------------------------------------------------------------------
