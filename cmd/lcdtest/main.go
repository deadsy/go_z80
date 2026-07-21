//-----------------------------------------------------------------------------
/*

LCD Emulation Testing

*/
//-----------------------------------------------------------------------------

package main

import (
	"image/color"
	"log"

	"github.com/deadsy/go_z80/cmd/lcdtest/lcd"
	"github.com/deadsy/go_z80/device/hd44780"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

const backgroundWidth = 1600
const backgroundHeight = 1200

func buildBackgroundImage() (*ebiten.Image, error) {
	img := ebiten.NewImage(backgroundWidth, backgroundHeight)
	img.Fill(color.RGBA{0x3e, 0x9b, 0x0a, 255})
	return img, nil
}

//-----------------------------------------------------------------------------

type system struct {
	dev0          *hd44780.LCD  // lcd device
	lcd0          *lcd.LCD      // character lcd
	background    *ebiten.Image // background graphic
	width, height int           // window dimensions
}

func newSystem() (*system, error) {

	// setup the LCD
	cfgLCD := hd44780.Config{
		Rows:   4,
		Cols:   20,
		XBase:  10,
		YBase:  10,
		XScale: 0.37,
		YScale: 0.37,
	}
	dev0, err := hd44780.New(&cfgLCD)
	if err != nil {
		return nil, err
	}
	lcd0, err := lcd.New(dev0, 4, 20)
	if err != nil {
		return nil, err
	}

	s := &system{
		dev0: dev0,
		lcd0: lcd0,
	}

	// build the background image
	img, err := buildBackgroundImage()
	if err != nil {
		return nil, err
	}
	s.background = img

	// set the background dimensions
	bounds := s.background.Bounds()
	s.width = bounds.Dx()
	s.height = bounds.Dy()

	return s, nil
}

func (s *system) Update() error {
	s.dev0.Update()
	return nil
}

func (s *system) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.background, nil)
	s.dev0.Draw(screen)
}

func (s *system) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.width, s.height
}

//-----------------------------------------------------------------------------

func main() {

	log.Printf("hd44780 lcd emulation test")

	s, err := newSystem()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	ebiten.SetWindowSize(s.width, s.height)
	ebiten.SetWindowTitle("HD44780 Emulation Test")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatalf("error: %s", err)
	}

}

//-----------------------------------------------------------------------------
