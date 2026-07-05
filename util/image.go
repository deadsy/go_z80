//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------

package util

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/bmp"
)

//-----------------------------------------------------------------------------

func SaveImage(ebitenImg *ebiten.Image, outputPath string) error {
	bounds := ebitenImg.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	ebitenImg.ReadPixels(rgbaImg.Pix)
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return bmp.Encode(file, rgbaImg)
}

//-----------------------------------------------------------------------------
