//-----------------------------------------------------------------------------
/*

Jupiter ACE Video

*/
//-----------------------------------------------------------------------------

package video

import (
	"image"
	"image/color"

	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

const glyphWidth = 8
const glyphHeight = 8
const numGlyphs = 256

const bytesPerGlyph = glyphHeight

// screen dimensions
const numCols = 32
const numRows = 24

const pixelsH = numCols * glyphWidth
const pixelsV = numRows * glyphHeight

const charAdr = uint16(0x2800)
const videoAdr = uint16(0x2000)
const charMask = 0x7f

//-----------------------------------------------------------------------------
// Note: 128 glyphs are defined in memory. If bit7 is set it is displayed inverted.
// To keep things simple we pre-build a complete set of normal and inverted glyphs.

func glyphAddress(n int) uint16 {
	return charAdr + uint16((n&charMask)*bytesPerGlyph)
}

func buildFontImage(m z80.Memory) *ebiten.Image {
	img := ebiten.NewImage(numGlyphs*glyphWidth, glyphHeight)
	for i := 0; i <= numGlyphs; i++ {
		for j := 0; j < glyphHeight; j++ {
			pixelData := m.Read8(glyphAddress(i) + uint16(j))
			// invert if bit7 == 1
			if i&0x80 != 0 {
				pixelData = ^pixelData
			}
			for k := 0; k < glyphWidth; k++ {
				pixel := (pixelData & (1 << (glyphWidth - k - 1))) != 0
				if pixel {
					img.Set(i*glyphWidth+k, j, color.RGBA{255, 255, 255, 255})
				}
			}
		}
	}
	return img
}

//-----------------------------------------------------------------------------

type Config struct {
	XBase, YBase   float64 // xy position
	XScale, YScale float64 // xy scale
}

type Video struct {
	config *Config       // video configuration
	font   *ebiten.Image // font atlas
	img    *ebiten.Image // unscaled video image
	mem    z80.Memory    // video memory
}

func New(cfg *Config, mem z80.Memory) (*Video, error) {
	// build the display
	return &Video{
		config: cfg,
		img:    ebiten.NewImage(pixelsH, pixelsV),
		mem:    mem,
	}, nil
}

func (v *Video) getGlyph(code byte) *ebiten.Image {
	x := int(code) * glyphWidth
	r := image.Rect(x, 0, x+glyphWidth, glyphHeight)
	return v.font.SubImage(r).(*ebiten.Image)
}

// Draw the video (called from ebiten draw function)
func (v *Video) Draw(screen *ebiten.Image) {
	if v.font == nil {
		return
	}
	// create an unscaled video image
	v.img.Clear()
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			// get the character
			code := v.mem.Read8(videoAdr + uint16((row*numCols)+col))
			glyph := v.getGlyph(code)
			// render the glyph to the video image
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(col*glyphWidth), float64(row*glyphHeight))
			v.img.DrawImage(glyph, op)
		}
	}

	// render the video image to the screen (with scaling)
	cfg := v.config
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(cfg.XScale, cfg.YScale)
	op.GeoM.Translate(cfg.XBase, cfg.YBase)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(v.img, op)
}

// Update the video logic (called selectively from ebiten update)
func (v *Video) Update() {
	v.font = buildFontImage(v.mem)
}

//-----------------------------------------------------------------------------
