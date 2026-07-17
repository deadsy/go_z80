//-----------------------------------------------------------------------------
/*

8x8 LED Array Emulation

Implements the TEC-1 8x8 LED Array
This wraps a generic 2D LED array and adds some write row/col byte control functions.

*/
//-----------------------------------------------------------------------------

package array88

import (
	"github.com/deadsy/go_z80/device/array"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

const numRows = 8
const numCols = 8

//-----------------------------------------------------------------------------

type Array88 struct {
	leds     *array.Array
	row, col byte // latched row/col values
}

func New(cfg array.Config) (*Array88, error) {
	cfg.Rows = numRows
	cfg.Cols = numCols
	a, err := array.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Array88{
		leds: a,
	}, nil
}

//-----------------------------------------------------------------------------

func (array *Array88) control() {
	for row := 0; row < numRows; row++ {
		if array.row&(1<<(numRows-row-1)) != 0 {
			// some of the leds in this row may be on.
			for col := 0; col < numCols; col++ {
				if array.col&(1<<col) != 0 {
					array.leds.Control(row, col, true)
				} else {
					array.leds.Control(row, col, false)
				}
			}
		} else {
			// all the leds in this row are off
			for col := 0; col < numCols; col++ {
				array.leds.Control(row, col, false)
			}
		}
	}
}

// WriteColumn controls the column latch
func (array *Array88) WriteColumn(val byte) {
	array.col = val
	array.control()
}

// WriteRow controls the row latch
func (array *Array88) WriteRow(val byte) {
	array.row = val
	array.control()
}

// Draw the Array LEDs (called from ebiten draw function)
func (array *Array88) Draw(screen *ebiten.Image) {
	array.leds.Draw(screen)
}

// Update the Array LED logic (called from ebiten update)
func (array *Array88) Update() {
	array.leds.Update()
}

//-----------------------------------------------------------------------------
