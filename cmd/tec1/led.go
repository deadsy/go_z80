//-----------------------------------------------------------------------------
/*

TEC-1 Emulation

Speaker Activity LED

*/
//-----------------------------------------------------------------------------

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

type LED struct {
}

func newLED() *LED {
	l := &LED{}
	return l
}

//-----------------------------------------------------------------------------

func (l *LED) control(state bool) {
}

//-----------------------------------------------------------------------------

// display draw function (called in game draw)
func (l *LED) draw(screen *ebiten.Image) {
}

// periodic update function (called in game update)
func (l *LED) update() {
}

//-----------------------------------------------------------------------------
