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
	img.Fill(color.RGBA{0xda, 0xe7, 0xe9, 255})
	return img, nil
}

//-----------------------------------------------------------------------------

func makeLCD(rows, cols int, x, y float64) (*lcd.LCD, error) {
	// setup the hd44780
	cfg := hd44780.Config{
		Rows:            rows,
		Cols:            cols,
		Type2:           (rows == 1) && (cols > 8),
		XBase:           x,
		YBase:           y,
		XScale:          0.37,
		YScale:          0.37,
		BackgroundColor: color.RGBA{0x3e, 0x9b, 0x0a, 255},
		CharacterColor:  color.RGBA{0, 0, 0, 255},
	}
	dev, err := hd44780.New(cfg)
	if err != nil {
		return nil, err
	}
	// return the "cooked" lcd device
	return lcd.New(dev, rows, cols)
}

//-----------------------------------------------------------------------------

type system struct {
	lcdSet        [6]*lcd.LCD
	background    *ebiten.Image // background graphic
	width, height int           // window dimensions
}

func newSystem() (*system, error) {
	s := &system{}

	display, err := makeLCD(1, 8, 10, 10)
	if err != nil {
		return nil, err
	}
	s.lcdSet[0] = display

	display, err = makeLCD(1, 32, 10, 70)
	if err != nil {
		return nil, err
	}
	s.lcdSet[1] = display

	display, err = makeLCD(2, 8, 10, 130)
	if err != nil {
		return nil, err
	}
	s.lcdSet[2] = display

	display, err = makeLCD(2, 40, 10, 226)
	if err != nil {
		return nil, err
	}
	s.lcdSet[3] = display

	display, err = makeLCD(4, 16, 10, 322)
	if err != nil {
		return nil, err
	}
	s.lcdSet[4] = display

	display, err = makeLCD(4, 20, 10, 491)
	if err != nil {
		return nil, err
	}
	s.lcdSet[5] = display

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

	for _, d := range s.lcdSet {
		d.Update()
	}
	return nil
}

func (s *system) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.background, nil)
	for _, d := range s.lcdSet {
		d.Draw(screen)
	}
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
