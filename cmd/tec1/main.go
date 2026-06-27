//-----------------------------------------------------------------------------
/*

Z80 Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//-----------------------------------------------------------------------------

type system struct {
	display       *Display      // 6 x 7 segment display
	led           *LED          // speaker activity LED
	io            *sysIO        // system IO
	mem           *sysMemory    // system memory
	bus           *Bus          // system bus
	cpu           *z80.CPU      // z80 cpu
	background    *ebiten.Image // background graphic
	width, height int           // window dimensions
	cycles        float32       // tick cpu cycles
}

func newSystem() (*system, error) {

	// setup the display
	display := newDisplay()

	// setup the LED
	led := newLED()

	// setup the IO
	io := newIO(display, led)

	// setup the memory
	mem, err := newMemory()
	if err != nil {
		return nil, err
	}

	// setup the bus
	bus := newBus()

	// setup the cpu
	cpu := z80.New(io, mem, bus)

	s := &system{
		display: display,
		led:     led,
		io:      io,
		mem:     mem,
		bus:     bus,
		cpu:     cpu,
	}

	// load background image
	img, _, err := ebitenutil.NewImageFromFile("../../images/tec1a.png")
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

const kHz = 1000
const MHz = kHz * kHz
const Hz = 1
const cpuClock = 500 * kHz
const tickRate = 60 * Hz
const cpuCyclesPerTick = float32(cpuClock) / float32(tickRate)

var updateCount int

func (s *system) Update() error {
	s.cycles += cpuCyclesPerTick
	for s.cycles > 0 {
		cycles, err := s.cpu.Run()
		if err != nil {
			return err
		}
		s.cycles -= float32(cycles)
	}
	s.display.update()
	s.led.update()

	updateCount += 1

	if updateCount == 30 {
		updateCount = 0
		s.cpu.NMI()
	}

	return nil
}

func (s *system) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)
	screen.DrawImage(s.background, nil)
	s.display.draw(screen)
	s.led.draw(screen)
}

func (s *system) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.width, s.height
}

func main() {
	s, err := newSystem()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	ebiten.SetWindowSize(s.width, s.height)
	ebiten.SetWindowTitle("TEC-1A")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatalf("error: %s", err)
	}
}

//-----------------------------------------------------------------------------
