//-----------------------------------------------------------------------------
/*

TEC-1 (Z80) Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"bytes"
	"embed"
	"fmt"
	"image/png"
	"log"

	"github.com/deadsy/go_z80/cmd/tec1/keypad"
	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/sevseg"
	"github.com/deadsy/go_z80/device/sixdigit"
	"github.com/deadsy/go_z80/device/speaker"
	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

//-----------------------------------------------------------------------------

//go:embed assets/mon1B.bin assets/tec1a.png
var assets embed.FS

//-----------------------------------------------------------------------------

const Hz = 1
const kHz = 1000 * Hz
const MHz = kHz * kHz

const cpuClock = 500 * kHz
const tickRate = 60 * Hz
const sampleRate = 48000

const cpuCyclesPerTick = float32(cpuClock) / float32(tickRate)     // cpu cycles per ebiten update tick
const cpuCyclesPerSample = float32(cpuClock) / float32(sampleRate) // cpu cycles per audio sample

//-----------------------------------------------------------------------------

func buildBackgroundImage() (*ebiten.Image, error) {
	data, err := assets.ReadFile("assets/tec1a.png")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded image: %w", err)
	}
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode PNG image: %w", err)
	}
	return ebiten.NewImageFromImage(img), nil
}

//-----------------------------------------------------------------------------

type system struct {
	display       *sixdigit.Display // 6 digit display
	led           *led.LED          // speaker activity LED
	keypad        *keypad.Keypad    // keypad
	speaker       *speaker.Speaker  // audio speaker
	io            *sysIO            // system IO
	mem           *sysMemory        // system memory
	bus           *Bus              // system bus
	cpu           *z80.CPU          // z80 cpu
	background    *ebiten.Image     // background graphic
	width, height int               // window dimensions
	tickCycles    float32           // ebiten tick cpu cycles
	sampleCycles  float32           // audio sample cpu cycles
}

func newSystem() (*system, error) {

	// setup the display
	const digitSize = float32(55.0)
	kDisplay := &sixdigit.Config{
		XBase:  362.0,
		YBase:  665.0,
		XScale: digitSize,
		YScale: sevseg.XYScale(digitSize),
		XGap0:  24.0,
		XGap1:  14.0,
	}
	display := sixdigit.New(kDisplay)

	// setup the LED
	kLED := &led.Config{
		XBase:  589.0,
		YBase:  600.5,
		Radius: 13.0,
	}
	led, err := led.New(kLED)
	if err != nil {
		return nil, err
	}

	// setup the keypad
	keypad, err := keypad.New()
	if err != nil {
		return nil, err
	}

	// setup the speaker
	k := speaker.Config{
		BitAmplitude: 0.1,
		BufferSize:   16384,
		SampleRate:   sampleRate,
		HighCutoff:   6 * kHz,
		LowCutoff:    40 * Hz,
	}
	speaker, err := speaker.New(&k)
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
	io := newIO(display, led, keypad)

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
		keypad:  keypad,
		speaker: speaker,
		io:      io,
		mem:     mem,
		bus:     bus,
		cpu:     cpu,
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

	if s.keypad.Update() {
		if s.keypad.Reset() {
			s.cpu.Reset()
		} else {
			// key presses are signalled with the NMI
			s.cpu.NMI()
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
	ebiten.SetWindowTitle("TEC-1A")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatalf("error: %s", err)
	}
}

//-----------------------------------------------------------------------------
