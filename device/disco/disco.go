//-----------------------------------------------------------------------------
/*

TEC-1G Disco LEDs Emulation

A driver for 2 RGB LEDs.

*/
//-----------------------------------------------------------------------------

package disco

import (
	"github.com/deadsy/go_z80/device/rgb"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

type Config struct {
	X, Y          float32 // xy position of display on screen
	YGap          float32 // y gap between leds
	Width, Height float32 // width/height size (rectangular only)
}

type Disco struct {
	led [2]*rgb.RGB
}

func New(cfg Config) (*Disco, error) {

	led0Config := rgb.Config{
		Type:   rgb.Rectangle,
		X:      cfg.X,
		Y:      cfg.Y,
		Width:  cfg.Width,
		Height: cfg.Height,
	}
	led0, err := rgb.New(led0Config)
	if err != nil {
		return nil, err
	}

	led1Config := rgb.Config{
		Type:   rgb.Rectangle,
		X:      cfg.X,
		Y:      cfg.Y + cfg.YGap + cfg.Height,
		Width:  cfg.Width,
		Height: cfg.Height,
	}
	led1, err := rgb.New(led1Config)
	if err != nil {
		return nil, err
	}

	return &Disco{
		led: [2]*rgb.RGB{led0, led1},
	}, nil
}

//-----------------------------------------------------------------------------

// Control the Disco lights (called from the IO layer)
func (d *Disco) Control(enable bool, val byte, cycles uint64) {
	if !enable {
		d.led[0].Control(false, false, false, cycles)
		d.led[1].Control(false, false, false, cycles)
		return
	}
	var r, g, b bool
	// led 0
	r = val&(1<<0) != 0
	g = val&(1<<1) != 0
	b = val&(1<<2) != 0
	d.led[0].Control(r, g, b, cycles)
	// led 1
	r = val&(1<<4) != 0
	g = val&(1<<5) != 0
	b = val&(1<<6) != 0
	d.led[1].Control(r, g, b, cycles)
}

// Draw the Disco lights (called from ebiten draw function)
func (d *Disco) Draw(screen *ebiten.Image) {
	d.led[0].Draw(screen)
	d.led[1].Draw(screen)
}

// Update the Disco lights (called from ebiten update)
func (d *Disco) Update(cycles uint64) {
	d.led[0].Update(cycles)
	d.led[1].Update(cycles)
}

//-----------------------------------------------------------------------------
