//-----------------------------------------------------------------------------
/*

BAR LED Emulation

Implements a 1-D linear array of LEDs

*/
//-----------------------------------------------------------------------------

package barled

import (
	"errors"
	"image/color"

	"github.com/deadsy/go_z80/device/led"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

type Config struct {
	N             int        // number of LEDs
	Type          led.Type   // led type
	X, Y          float32    // xy position of display on screen
	XStep, YStep  float32    // step distance between LEDs
	Radius        float32    // radius (round only)
	Width, Height float32    // width/height size (rectangular only)
	On, Off       color.RGBA // on/off colors
}

type BarLED struct {
	leds []*led.LED
}

func New(cfg *Config) (*BarLED, error) {
	if cfg.N < 2 || cfg.N > 32 {
		return nil, errors.New("bad number of leds")
	}
	// build the leds
	leds := make([]*led.LED, cfg.N)
	for i := range leds {
		ledCfg := led.Config{
			Type:   cfg.Type,
			X:      cfg.X + float32(i)*cfg.XStep,
			Y:      cfg.Y + float32(i)*cfg.YStep,
			Radius: cfg.Radius,
			Width:  cfg.Width,
			Height: cfg.Height,
			On:     cfg.On,
			Off:    cfg.Off,
		}
		led, err := led.New(&ledCfg)
		if err != nil {
			return nil, err
		}
		leds[i] = led
	}

	return &BarLED{
		leds: leds,
	}, nil
}

//-----------------------------------------------------------------------------

// Control the Bar LED (called from the IO layer)
func (bar *BarLED) Control(n int, state bool) {
	if n < 0 || n >= len(bar.leds) {
		return
	}
	bar.leds[n].Control(state)
}

// Draw the Bar LED (called from ebiten draw function)
func (bar *BarLED) Draw(screen *ebiten.Image) {
	for _, led := range bar.leds {
		led.Draw(screen)
	}
}

// Update the Bar LED logic (called from ebiten update)
func (bar *BarLED) Update() {
	for _, led := range bar.leds {
		led.Update()
	}
}

//-----------------------------------------------------------------------------
