//-----------------------------------------------------------------------------
/*

TEC-1G (Z80) Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"
	"log"

	"github.com/deadsy/go_z80/device/hd44780"
	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/seven_segment"
	"github.com/deadsy/go_z80/device/six_digit"
	"github.com/deadsy/go_z80/device/speaker"
	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//-----------------------------------------------------------------------------

const Hz = 1
const kHz = 1000 * Hz
const MHz = kHz * kHz

const cpuClock = 4 * MHz
const tickRate = 60 * Hz
const sampleRate = 48000

const cpuCyclesPerTick = float32(cpuClock) / float32(tickRate)     // cpu cycles per ebiten update tick
const cpuCyclesPerSample = float32(cpuClock) / float32(sampleRate) // cpu cycles per audio sample

//-----------------------------------------------------------------------------

type system struct {
	display       *six_digit.Display // 6 digit display
	led           *led.LED           // speaker activity LED
	speaker       *speaker.Speaker   // audio speaker
	io            *sysIO             // system IO
	mem           *sysMemory         // system memory
	bus           *Bus               // system bus
	cpu           *z80.CPU           // z80 cpu
	background    *ebiten.Image      // background graphic
	width, height int                // window dimensions
	tickCycles    float32            // ebiten tick cpu cycles
	sampleCycles  float32            // audio sample cpu cycles
}

func newSystem() (*system, error) {

	// setup the display
	const digitSize = float32(70.0)
	kDisplay := &six_digit.Config{
		XBase:  195.0,
		YBase:  855.0,
		XScale: digitSize,
		YScale: seven_segment.XYScale(digitSize),
		XGap0:  15.8,
		XGap1:  28.0,
	}
	display := six_digit.New(kDisplay)

	// setup the LED
	kLED := &led.Config{
		XBase:  926,
		YBase:  514,
		Radius: 15,
	}
	led, err := led.New(kLED)
	if err != nil {
		return nil, err
	}

	// setup the speaker
	kSpeaker := speaker.Config{
		BitAmplitude: 0.1,
		BufferSize:   16384,
		SampleRate:   sampleRate,
		HighCutoff:   6 * kHz,
		LowCutoff:    40 * Hz,
	}
	speaker, err := speaker.New(&kSpeaker)
	if err != nil {
		return nil, err
	}

	// setup the LCD
	kLCD := hd44780.Config{
		Mode: hd44780.Mode20x4,
	}
	lcd, err := hd44780.New(&kLCD)
	if err != nil {
		return nil, err
	}

	// setup the audio player
	ctx := audio.NewContext(sampleRate)
	player, err := ctx.NewPlayer(speaker)
	if err != nil {
		return nil, err
	}

	// setup the IO
	io := newIO(display, led, lcd)

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
		speaker: speaker,
		io:      io,
		mem:     mem,
		bus:     bus,
		cpu:     cpu,
	}

	// load background image
	img, _, err := ebitenutil.NewImageFromFile("../../images/tec1g.png")
	if err != nil {
		return nil, err
	}
	s.background = img

	// set the background dimensions
	bounds := s.background.Bounds()
	s.width = bounds.Dx()
	s.height = bounds.Dy()

	// start the audio player
	player.Play()

	return s, nil
}

func (s *system) Update() error {
	// run the cpu for as many cycles as are in an update tick
	s.tickCycles += cpuCyclesPerTick
	for s.tickCycles > 0 {
		cycles, err := s.cpu.Run()
		if err != nil {
			return err
		}
		s.tickCycles -= float32(cycles)
		// sample the audio output
		s.sampleCycles -= float32(cycles)
		for s.sampleCycles <= 0 {
			err := s.speaker.WriteSample(s.io.speaker)
			if err != nil {
				return fmt.Errorf("speaker.WriteSample: %s", err)
			}
			s.sampleCycles += cpuCyclesPerSample
		}
	}

	s.display.Update()
	s.led.Update()
	return nil
}

func (s *system) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.background, nil)
	s.display.Draw(screen)
	s.led.Draw(screen)
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
	ebiten.SetWindowTitle("TEC-1G")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatalf("error: %s", err)
	}
}

//-----------------------------------------------------------------------------
