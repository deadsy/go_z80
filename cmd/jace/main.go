//-----------------------------------------------------------------------------
/*

Jupiter ACE Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"
	"log"

	"github.com/deadsy/go_z80/device/keyboard"
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

const cpuClock = 3.25 * MHz

// ebiten timing
const tickRate = 60 * Hz
const cpuCyclesPerTick = float32(cpuClock) / float32(tickRate) // cpu cycles per ebiten update tick

// audio timing
const audioSampleRate = 48000
const cpuCyclesPerAudioSample = float32(cpuClock) / float32(audioSampleRate)

// periodic interrupt
const interruptRate = 50 * Hz
const cpuCyclesPerInterrupt = float32(cpuClock) / float32(interruptRate) // cpu cycles per interrupt tick

//-----------------------------------------------------------------------------

type system struct {
	speaker           *speaker.Speaker // audio speaker
	keyboard          *keyboard.Jace   // matrix keyboard
	io                *sysIO           // system IO
	mem               *sysMemory       // system memory
	bus               *Bus             // system bus
	cpu               *z80.CPU         // z80 cpu
	background        *ebiten.Image    // background graphic
	width, height     int              // window dimensions
	tickCycles        float32          // ebiten tick cpu cycles
	audioSampleCycles float32          // audio sample cpu cycles
	interruptCycles   float32          // periodic interrupt
}

func newSystem() (*system, error) {

	// setup the speaker
	kSpeaker := speaker.Config{
		BitAmplitude: 0.1,
		BufferSize:   16384,
		SampleRate:   audioSampleRate,
		HighCutoff:   6 * kHz,
		LowCutoff:    40 * Hz,
	}
	speaker, err := speaker.New(&kSpeaker)
	if err != nil {
		return nil, err
	}

	// setup the audio player
	ctx := audio.NewContext(audioSampleRate)
	player, err := ctx.NewPlayer(speaker)
	if err != nil {
		return nil, err
	}

	// setup the keyboard
	keyboard, err := keyboard.NewJace()
	if err != nil {
		return nil, err
	}

	// setup the IO
	io := newIO(keyboard)

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
		speaker:         speaker,
		io:              io,
		mem:             mem,
		bus:             bus,
		cpu:             cpu,
		interruptCycles: cpuCyclesPerInterrupt,
	}

	// load background image
	img, _, err := ebitenutil.NewImageFromFile("../../images/keyboard.png")
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

// exit cleans up system resources
func (s *system) Exit() {
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

		// periodic interrupts
		s.interruptCycles -= float32(cycles)
		for s.interruptCycles < 0 {
			s.cpu.IRQ()
			s.interruptCycles += cpuCyclesPerInterrupt
		}

		// sample the audio output
		s.audioSampleCycles -= float32(cycles)
		for s.audioSampleCycles < 0 {
			err := s.speaker.WriteSample(s.io.speaker)
			if err != nil {
				return fmt.Errorf("speaker.WriteSample: %s", err)
			}
			s.audioSampleCycles += cpuCyclesPerAudioSample
		}
	}

	return nil
}

func (s *system) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.background, nil)
}

func (s *system) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.width, s.height
}

//-----------------------------------------------------------------------------

func main() {

	s, err := newSystem()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer s.Exit()

	ebiten.SetWindowSize(s.width, s.height)
	ebiten.SetWindowTitle("Jupiter Ace")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatalf("error: %s", err)
	}

}

//-----------------------------------------------------------------------------
