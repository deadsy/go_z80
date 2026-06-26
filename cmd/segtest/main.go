package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth    = 450
	screenHeight   = 200
	numDigits      = 4
	multiplexHertz = 5000 // Try 15 to watch the multiplexing loop slow down!
)

// Segment color definitions
var (
	bgColor = color.RGBA{15, 15, 15, 255} // Dark background
	segOff  = color.RGBA{40, 10, 10, 255} // Dim red for unlit segments
	segOn   = color.RGBA{255, 0, 0, 255}  // Bright red for lit segments
)

// SEGMENT_MAP: Maps numbers 0-9 to active segments [A, B, C, D, E, F, G]
var segmentMap = [10][7]int{
	{1, 1, 1, 1, 1, 1, 0}, // 0
	{0, 1, 1, 0, 0, 0, 0}, // 1
	{1, 1, 0, 1, 1, 0, 1}, // 2
	{1, 1, 1, 1, 0, 0, 1}, // 3
	{0, 1, 1, 0, 0, 1, 1}, // 4
	{1, 0, 1, 1, 0, 1, 1}, // 5
	{1, 0, 1, 1, 1, 1, 1}, // 6
	{1, 1, 1, 0, 0, 0, 0}, // 7
	{1, 1, 1, 1, 1, 1, 1}, // 8
	{1, 1, 1, 1, 0, 1, 1}, // 9
}

type Game struct {
	displayBuffer  [numDigits]int
	activeDigit    int
	lastSwitchTime time.Time
	texture        *ebiten.Image
	glow           [numDigits][7]float32
}

func (g *Game) Update() error {
	// Calculate elapsed time for hardware clock multiplexing
	now := time.Now()
	elapsed := now.Sub(g.lastSwitchTime)
	switchInterval := time.Second / time.Duration(multiplexHertz)

	if elapsed >= switchInterval {
		// Hardware multiplex switch: Enable the next sequential common pin
		g.activeDigit = (g.activeDigit + 1) % numDigits
		g.lastSwitchTime = now
	}

	// Apply persistence of vision decay & illumination rules
	for d := 0; d < numDigits; d++ {
		for seg := 0; seg < 7; seg++ {
			if d == g.activeDigit {
				// Check if the segment bit is active for the current number
				isActive := segmentMap[g.displayBuffer[d]][seg] != 0
				if isActive {
					g.glow[d][seg] = 1.0 // Fully illuminated
					continue
				}
			}
			// Simulate natural LED decay/cool-off over time
			g.glow[d][seg] *= 0.01
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(bgColor)

	// Hardware Emulation: Draw all 4 digit places
	for digitIndex := 0; digitIndex < numDigits; digitIndex++ {
		xPos := float32(50 + (digitIndex * 90))
		yPos := float32(50)
		val := g.displayBuffer[digitIndex]

		// Pin validation: Segment data is only visible if the active common pin matches
		isActive := (digitIndex == g.activeDigit)
		g.drawSegmentDigit(screen, xPos, yPos, val, isActive, g.glow[digitIndex][:])
	}
}

func (g *Game) drawSegmentDigit(screen *ebiten.Image, x, y float32, value int, isActive bool, glow []float32) {
	const (
		w float32 = 60  // Width
		h float32 = 100 // Height
		t float32 = 10  // Thickness
	)

	// Geometries for segments A through G
	states := segmentMap[value]

	// Segment A (Top)
	g.drawSegment(screen, isActive && states[0] == 1, glow[0], []float32{
		x + t, y,
		x + w - t, y,
		x + w - t - 3, y + t,
		x + t + 3, y + t,
	})
	// Segment B (Top-Right)
	g.drawSegment(screen, isActive && states[1] == 1, glow[1], []float32{
		x + w - t, y + t,
		x + w, y + t,
		x + w, y + h/2 - t/2,
		x + w - t, y + h/2 - t/2,
	})
	// Segment C (Bottom-Right)
	g.drawSegment(screen, isActive && states[2] == 1, glow[2], []float32{
		x + w - t, y + h/2 + t/2,
		x + w, y + h/2 + t/2,
		x + w, y + h - t,
		x + w - t, y + h - t,
	})
	// Segment D (Bottom)
	g.drawSegment(screen, isActive && states[3] == 1, glow[3], []float32{
		x + t, y + h - t,
		x + w - t, y + h - t,
		x + w - t - 3, y + h,
		x + t + 3, y + h,
	})
	// Segment E (Bottom-Left)
	g.drawSegment(screen, isActive && states[4] == 1, glow[4], []float32{
		x, y + h/2 + t/2,
		x + t, y + h/2 + t/2,
		x + t, y + h - t,
		x, y + h - t,
	})
	// Segment F (Top-Left)
	g.drawSegment(screen, isActive && states[5] == 1, glow[5], []float32{
		x, y + t,
		x + t, y + t,
		x + t, y + h/2 - t/2,
		x, y + h/2 - t/2,
	})
	// Segment G (Middle)
	g.drawSegment(screen, isActive && states[6] == 1, glow[6], []float32{
		x + t, y + h/2 - t/2,
		x + w - t, y + h/2 - t/2,
		x + w - t, y + h/2 + t/2,
		x + t, y + h/2 + t/2,
	})
}

// Helper to render the vector shapes onto the Ebiten canvas
func (g *Game) drawSegment(screen *ebiten.Image, isOn bool, glow float32, points []float32) {
	c := segOff
	if isOn {
		c = segOn
	}

	// Setup vertices for a 4-point polygon (quad)
	var vertices [4]ebiten.Vertex
	for i := 0; i < 4; i++ {
		vertices[i] = ebiten.Vertex{
			DstX:   points[i*2],
			DstY:   points[i*2+1],
			SrcX:   0,
			SrcY:   0,
			ColorR: glow * float32(c.R) / 255.0,
			ColorG: glow * float32(c.G) / 255.0,
			ColorB: glow * float32(c.B) / 255.0,
			ColorA: float32(c.A) / 255.0,
		}
	}

	indices := []uint16{0, 1, 2, 0, 2, 3}
	screen.DrawTriangles(vertices[:], indices, g.texture, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		displayBuffer:  [4]int{2, 0, 2, 6}, // Hardcoded target numbers to display
		lastSwitchTime: time.Now(),
		texture:        ebiten.NewImage(1, 1),
	}
	game.texture.Fill(color.White)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten Multiplexed 7-Segment")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
