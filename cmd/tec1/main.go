//-----------------------------------------------------------------------------
/*

TEC-1A Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"bytes"
	"embed"
	"fmt"
	"image/color"
	"image/png"
	"log"

	"github.com/deadsy/go_z80/cmd/tec1/keypad"
	"github.com/deadsy/go_z80/device/array"
	"github.com/deadsy/go_z80/device/array88"
	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/sevseg"
	"github.com/deadsy/go_z80/device/sixdigit"
	"github.com/deadsy/go_z80/device/sound"
	"github.com/deadsy/go_z80/device/speaker"
	"github.com/deadsy/go_z80/util"
	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

//go:embed assets/mon1B.bin assets/tec1a.png
var assets embed.FS

//-----------------------------------------------------------------------------

const Hz = 1
const kHz = 1000 * Hz
const MHz = kHz * kHz

const cpuClock = 500 * kHz

// ebiten timing
const tickRate = 60 * Hz
const cpuCyclesPerTick = float32(cpuClock) / float32(tickRate) // cpu cycles per ebiten update tick

// audio timing
const audioSampleRate = 48000 // samples/sec
const cpuCyclesPerAudioSample = float32(cpuClock) / float32(audioSampleRate)

//-----------------------------------------------------------------------------

func buildBackgroundImage() (*ebiten.Image, error) {
	data, err := assets.ReadFile("assets/tec1a.png")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded image: %w", err)
	}
	src, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode PNG image: %w", err)
	}
	b := src.Bounds()
	img := ebiten.NewImageWithOptions(b, &ebiten.NewImageOptions{
		Unmanaged: true, // don't put a big background on a texture atlas
	})
	img.DrawImage(ebiten.NewImageFromImage(src), nil)
	return img, nil
}

//-----------------------------------------------------------------------------

type system struct {
	cfg               *Config          // configuration
	speaker           *speaker.Speaker // audio speaker
	sound             *sound.Sound     // ebiten audio
	io                *sysIO           // system IO
	mem               *sysMemory       // system memory
	cpu               *z80.CPU         // z80 cpu
	background        *ebiten.Image    // background graphic
	width, height     int              // window dimensions
	tickCycles        float32          // ebiten tick cpu cycles
	audioSampleCycles float32          // audio sample cpu cycles
	soundStarted      bool             // has the sound been started?
	haltLogged        bool             // have we logged a cpu halt?
}

func newSystem(cfg *Config) (*system, error) {

	// setup the display
	const digitSize = float32(55.0)
	cfgDisplay := sixdigit.Config{
		XBase:  362,
		YBase:  665,
		XScale: digitSize,
		YScale: sevseg.XYScale(digitSize),
		XGap0:  24,
		XGap1:  14,
	}
	display := sixdigit.New(&cfgDisplay)

	// setup the LED
	cfgSpeakerLED := led.Config{
		Type:   led.Round,
		X:      589,
		Y:      600,
		Radius: 13,
		On:     color.RGBA{0, 255, 0, 128},
		Off:    color.RGBA{0, 0, 0, 0},
	}
	ledSpeaker, err := led.New(&cfgSpeakerLED)
	if err != nil {
		return nil, err
	}

	// setup the speaker
	cfgSpeaker := speaker.Config{
		Enable:       cfg.Sound.Enable,
		BitAmplitude: 0.1,
		BufferSize:   16384,
		SampleRate:   audioSampleRate,
		HighCutoff:   6 * kHz,
		LowCutoff:    40 * Hz,
	}
	speaker, err := speaker.New(cfgSpeaker)
	if err != nil {
		return nil, err
	}

	// setup the sound
	cfgSound := sound.Config{
		Enable:     cfg.Sound.Enable,
		SampleRate: audioSampleRate,
		Src:        speaker,
	}
	sound, err := sound.New(cfgSound)
	if err != nil {
		return nil, err
	}

	// setup the keypad
	keypad, err := keypad.New()
	if err != nil {
		return nil, err
	}

	// setup the 8x8 LED display
	cfgLedArray := array.Config{
		Enable:     cfg.Array88.Enable,
		Type:       led.Rectangle,
		X:          100,
		Y:          100,
		XGap:       1,
		YGap:       1,
		Width:      20,
		Height:     20,
		On:         color.RGBA{0, 0xff, 0, 255},
		Off:        color.RGBA{0x90, 0x90, 0x90, 255},
		Background: color.RGBA{0x80, 0x80, 0x80, 255},
		Border:     10,
	}
	ledArray, err := array88.New(cfgLedArray)
	if err != nil {
		return nil, err
	}

	// setup the IO
	devices := ioDevices{
		display:    display,
		ledSpeaker: ledSpeaker,
		ledArray:   ledArray,
		keypad:     keypad,
	}
	io := newIO(&devices)

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
		cfg:     cfg,
		speaker: speaker,
		sound:   sound,
		io:      io,
		mem:     mem,
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

	return s, nil
}

// exit cleans up system resources
func (s *system) Exit() {
	log.Printf("system exit")
	err := s.cfg.saveConfig(s, configFile)
	if err != nil {
		log.Printf("unable to save config: %s", err)
	} else {
		log.Printf("saved config to %s", configFile)
	}
}

func (s *system) Update() error {

	// start the sound (once)
	if !s.soundStarted && s.sound.IsReady() && s.speaker.Samples() >= 800 {
		s.speaker.Empty()
		log.Printf("starting sound\n")
		err := s.sound.Start()
		if err != nil {
			log.Printf("unable to start sound: %s\n", err)
		} else {
			s.soundStarted = true
		}
	}

	// run the cpu for as many cycles as are in an update tick
	s.tickCycles += cpuCyclesPerTick
	for s.tickCycles > 0 {
		cycles, err := s.cpu.Run()
		if err != nil {
			return err
		}
		s.tickCycles -= float32(cycles)

		// sample the audio output
		s.audioSampleCycles -= float32(cycles)
		for s.audioSampleCycles <= 0 {
			err := s.speaker.WriteSample(s.io.speaker)
			if err != nil {
				log.Printf("speaker.WriteSample: %s", err)
				s.speaker.Empty()
			}
			s.audioSampleCycles += cpuCyclesPerAudioSample
		}
	}

	// cpu halted?
	if s.cpu.IsHalted() && !s.haltLogged {
		log.Printf("cpu halted at pc=0x%04x", s.cpu.PC)
		s.haltLogged = true
	}

	// update the IO devices
	s.io.Update()

	if s.io.dev.keypad.Update() {
		if s.io.dev.keypad.Reset() {
			s.cpu.Reset()
		} else {
			// key presses are signalled with the NMI
			s.cpu.NMI()
		}
	}

	if ebiten.IsWindowBeingClosed() {
		s.Exit()
	}

	return nil
}

func (s *system) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.background, nil)
	// draw the IO devices
	s.io.Draw(screen)
}

func (s *system) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.width, s.height
}

//-----------------------------------------------------------------------------

func main() {
	log.Printf("%s\n", util.GetBuildInfo())

	// read the config
	cfg, err := loadConfig(configFile)
	if err != nil {
		log.Printf("unable to read %s, using defaults", configFile)
		cfg = defaultConfig()
	} else {
		log.Printf("read config from %s", configFile)
	}

	s, err := newSystem(cfg)
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
