//-----------------------------------------------------------------------------
/*

LED Array Emulation

Implements a 2-D array of LEDs

*/
//-----------------------------------------------------------------------------

package array

import (
	"errors"
	"image/color"

	"github.com/deadsy/go_z80/device/led"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

type Config struct {
	Rows, Cols    int        // number of rows and columns
	Type          led.Type   // led type
	X, Y          float32    // xy position of display on screen
	XStep, YStep  float32    // step distance between LEDs
	Radius        float32    // radius (round only)
	Width, Height float32    // width/height size (rectangular only)
	On, Off       color.RGBA // on/off colors
}

type Array struct {
	numCols int
	leds    []*led.LED
}

func New(cfg Config) (*Array, error) {
	if cfg.Rows < 1 {
		return nil, errors.New("bad number of rows")
	}
	if cfg.Cols < 1 {
		return nil, errors.New("bad number of columns")
	}
	// build the leds
	leds := make([]*led.LED, cfg.Rows*cfg.Cols)
	for j := 0; j < cfg.Rows; j++ {
		for i := 0; i < cfg.Cols; i++ {
			ledCfg := led.Config{
				Type:   cfg.Type,
				X:      cfg.X + float32(i)*cfg.XStep,
				Y:      cfg.Y + float32(j)*cfg.YStep,
				Radius: cfg.Radius,
				Width:  cfg.Width,
				Height: cfg.Height,
				On:     cfg.On,
				Off:    cfg.Off,
			}
			l, err := led.New(&ledCfg)
			if err != nil {
				return nil, err
			}
			leds[(j*cfg.Cols)+i] = l
		}
	}

	return &Array{
		numCols: cfg.Cols,
		leds:    leds,
	}, nil
}

//-----------------------------------------------------------------------------

// Control an Array LED
func (array *Array) Control(row, col int, state bool) {
	n := (row * array.numCols) + col
	if n < 0 || n >= len(array.leds) {
		return
	}
	array.leds[n].Control(state)
}

// Draw the Array LEDs (called from ebiten draw function)
func (array *Array) Draw(screen *ebiten.Image) {
	for _, l := range array.leds {
		l.Draw(screen)
	}
}

// Update the Array LED logic (called from ebiten update)
func (array *Array) Update() {
	for _, l := range array.leds {
		l.Update()
	}
}

//-----------------------------------------------------------------------------
