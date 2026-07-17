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
	Enable        bool       // is the array enabled?
	Rows, Cols    int        // number of rows and columns
	Type          led.Type   // led type
	X, Y          float32    // xy position of display on screen
	XGap, YGap    float32    // gap distance between LEDs
	Radius        float32    // radius (round only)
	Width, Height float32    // width/height size (rectangular only)
	On, Off       color.RGBA // on/off colors
	Background    color.RGBA // background color
	Border        float32    // background border
}

type Array struct {
	cfg        Config
	numCols    int
	leds       []*led.LED
	background *ebiten.Image
}

func New(cfg Config) (*Array, error) {
	if !cfg.Enable {
		return &Array{}, nil
	}
	if cfg.Rows < 1 {
		return nil, errors.New("bad number of rows")
	}
	if cfg.Cols < 1 {
		return nil, errors.New("bad number of columns")
	}

	var xOfs, yOfs, xStep, yStep float32
	var width, height int
	if cfg.Type == led.Round {
		if cfg.Radius <= 0 {
			return nil, errors.New("bad radius")
		}
		xOfs = cfg.Border + cfg.Radius + cfg.X
		yOfs = cfg.Border + cfg.Radius + cfg.Y
		xStep = cfg.XGap + 2.0*cfg.Radius
		yStep = cfg.YGap + 2.0*cfg.Radius
		width = int(2*cfg.Border + cfg.XGap*float32(cfg.Cols-1) + 2*cfg.Radius*float32(cfg.Cols))
		height = int(2*cfg.Border + cfg.YGap*float32(cfg.Rows-1) + 2*cfg.Radius*float32(cfg.Rows))
	} else {
		if cfg.Width <= 0 {
			return nil, errors.New("bad width")
		}
		if cfg.Height <= 0 {
			return nil, errors.New("bad height")
		}
		xOfs = cfg.Border + cfg.X
		yOfs = cfg.Border + cfg.Y
		xStep = cfg.XGap + cfg.Width
		yStep = cfg.YGap + cfg.Height
		width = int(2*cfg.Border + cfg.XGap*float32(cfg.Cols-1) + cfg.Width*float32(cfg.Cols))
		height = int(2*cfg.Border + cfg.YGap*float32(cfg.Rows-1) + cfg.Height*float32(cfg.Rows))
	}

	// build the leds
	leds := make([]*led.LED, cfg.Rows*cfg.Cols)
	for j := 0; j < cfg.Rows; j++ {
		for i := 0; i < cfg.Cols; i++ {
			ledCfg := led.Config{
				Type:   cfg.Type,
				X:      xOfs + float32(i)*xStep,
				Y:      yOfs + float32(j)*yStep,
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

	// build the background image
	background := ebiten.NewImage(width, height)
	background.Fill(cfg.Background)

	return &Array{
		cfg:        cfg,
		leds:       leds,
		background: background,
	}, nil
}

//-----------------------------------------------------------------------------

// Control an Array LED
func (array *Array) Control(row, col int, state bool) {
	if !array.cfg.Enable {
		return
	}
	n := (row * array.cfg.Cols) + col
	if n < 0 || n >= len(array.leds) {
		return
	}
	array.leds[n].Control(state)
}

// Draw the Array LEDs (called from ebiten draw function)
func (array *Array) Draw(screen *ebiten.Image) {
	if !array.cfg.Enable {
		return
	}
	cfg := array.cfg
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cfg.X), float64(cfg.Y))
	screen.DrawImage(array.background, op)

	for _, l := range array.leds {
		l.Draw(screen)
	}
}

// Update the Array LED logic (called from ebiten update)
func (array *Array) Update() {
	if !array.cfg.Enable {
		return
	}
	for _, l := range array.leds {
		l.Update()
	}
}

//-----------------------------------------------------------------------------
