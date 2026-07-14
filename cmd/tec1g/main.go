//-----------------------------------------------------------------------------
/*

TEC-1G (Z80) Emulator

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

	"github.com/deadsy/go_z80/cmd/tec1g/keyboard"
	"github.com/deadsy/go_z80/device/ds1302"
	"github.com/deadsy/go_z80/device/hd44780"
	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/serial"
	"github.com/deadsy/go_z80/device/sevseg"
	"github.com/deadsy/go_z80/device/sixdigit"
	"github.com/deadsy/go_z80/device/sound"
	"github.com/deadsy/go_z80/device/speaker"
	"github.com/deadsy/go_z80/util"
	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

//go:embed assets/mon3_2025BC_16.bin assets/tec1g.png assets/DIAG-1G_CH24-11.bin
var assets embed.FS

//-----------------------------------------------------------------------------

const Hz = 1
const kHz = 1000 * Hz
const MHz = kHz * kHz

const cpuClock = 4 * MHz

// ebiten timing
const tickRate = 60 * Hz
const cpuCyclesPerTick = float32(cpuClock) / float32(tickRate) // cpu cycles per ebiten update tick

// audio timing
const audioSampleRate = 48000 // samples/sec
const cpuCyclesPerAudioSample = float32(cpuClock) / float32(audioSampleRate)

// serial timing
const serialSamplesPerBit = 5
const serialBaudRate = 4800 // this is the default MON3 rate
const cpuCyclesPerSerialSample = float32(cpuClock) / (float32(serialBaudRate) * float32(serialSamplesPerBit))

//-----------------------------------------------------------------------------

func buildBackgroundImage() (*ebiten.Image, error) {
	data, err := assets.ReadFile("assets/tec1g.png")
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
	cfg                *Config          // configuration
	speaker            *speaker.Speaker // audio speaker
	sound              *sound.Sound     // ebiten audio
	uart               *serial.UART     // serial uart
	pty                *serial.PTY      // pseudo tty
	io                 *sysIO           // system IO
	mem                *sysMemory       // system memory
	bus                *Bus             // system bus
	cpu                *z80.CPU         // z80 cpu
	background         *ebiten.Image    // background graphic
	width, height      int              // window dimensions
	tickCycles         float32          // ebiten tick cpu cycles
	audioSampleCycles  float32          // audio sample cpu cycles
	serialSampleCycles float32          // serial sample cpu cycles
	soundStarted       bool             // has the sound been started?
}

func newSystem(cfg *Config) (*system, error) {

	// setup the display
	const digitSize = float32(70.0)
	cfgDisplay := sixdigit.Config{
		XBase:  195.0,
		YBase:  855.0,
		XScale: digitSize,
		YScale: sevseg.XYScale(digitSize),
		XGap0:  15.8,
		XGap1:  28.0,
	}
	display := sixdigit.New(&cfgDisplay)

	// setup the Speaker LED
	cfgSpeakerLed := led.Config{
		Type:   led.Round,
		X:      931,
		Y:      514,
		Radius: 12,
		On:     color.RGBA{255, 255, 255, 128},
		Off:    color.RGBA{0, 0, 0, 0},
	}
	ledSpeaker, err := led.New(&cfgSpeakerLed)
	if err != nil {
		return nil, err
	}

	// setup the speaker
	cfgSpeaker := speaker.Config{
		BitAmplitude: 0.1,
		BufferSize:   16384,
		SampleRate:   audioSampleRate,
		HighCutoff:   6 * kHz,
		LowCutoff:    40 * Hz,
	}
	speaker, err := speaker.New(&cfgSpeaker)
	if err != nil {
		return nil, err
	}

	// setup the sound
	cfgSound := sound.Config{
		SampleRate: audioSampleRate,
		Src:        speaker,
	}
	sound, err := sound.New(&cfgSound)
	if err != nil {
		return nil, err
	}

	// setup the LCD
	cfgLCD := hd44780.Config{
		Mode:   hd44780.Mode20x4,
		XBase:  233,
		YBase:  584,
		XScale: 0.34,
		YScale: 0.34,
	}
	lcd, err := hd44780.New(&cfgLCD)
	if err != nil {
		return nil, err
	}

	// setup the keyboard
	keyboard, err := keyboard.New()
	if err != nil {
		return nil, err
	}

	// setup the RTC
	rtc, err := ds1302.New(&cfg.RTC)
	if err != nil {
		return nil, err
	}

	// setup the serial
	cfgSerial := serial.Config{
		SamplesPerBit: serialSamplesPerBit,
		DataBits:      8,
		StopBits:      1,
	}
	uart, err := serial.NewUART(&cfgSerial)
	if err != nil {
		return nil, err
	}

	// setup the pseudo-tty
	pty, err := serial.NewPTY()
	if err != nil {
		return nil, err
	}
	log.Printf("serial port at %s\n", pty.Name())

	// setup the Halt LED
	cfgHaltLed := led.Config{
		Type:   led.Round,
		X:      1271,
		Y:      526,
		Radius: 12,
		On:     color.RGBA{255, 0, 0, 128},
		Off:    color.RGBA{0, 0, 0, 0},
	}
	ledHalt, err := led.New(&cfgHaltLed)
	if err != nil {
		return nil, err
	}

	// setup the IO
	devices := ioDevices{
		display:    display,
		ledSpeaker: ledSpeaker,
		ledHalt:    ledHalt,
		lcd:        lcd,
		keyboard:   keyboard,
		rtc:        rtc,
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
		uart:    uart,
		pty:     pty,
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
	s.io.dev.rtc.Close()
	s.pty.Close()
}

func (s *system) Update() error {

	// start the sound (once)
	if !s.soundStarted && s.sound.IsReady() && s.speaker.Samples() >= 800 {
		log.Printf("starting sound")
		err := s.sound.Start()
		if err != nil {
			log.Printf("unable to start sound: %s", err)
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
		for s.audioSampleCycles < 0 {
			err := s.speaker.WriteSample(s.io.speaker)
			if err != nil {
				log.Printf("speaker.WriteSample: %s", err)
				s.speaker.Empty()
			}
			s.audioSampleCycles += cpuCyclesPerAudioSample
		}

		// sample the serial tx/rx
		s.serialSampleCycles -= float32(cycles)
		for s.serialSampleCycles < 0 {
			// sample the tx line *from* the tec-1g
			rx, ok, err := s.uart.WriteSample(s.io.serialTx)
			if ok {
				s.pty.Write(byte(rx))
			} else if err != nil {
				log.Printf("uart.WriteSample: %s", err)
			}
			// drive the rx line *to* the tec-1g
			sample, err := s.uart.ReadSample(s.pty)
			if err != nil {
				log.Printf("uart.ReadSample: %s", err)
			}
			s.io.serialRx = sample
			s.serialSampleCycles += cpuCyclesPerSerialSample
		}
	}

	// halt LED
	s.io.dev.ledHalt.Control(s.cpu.IsHalted())

	// update the IO devices
	s.io.Update()

	if s.io.dev.keyboard.Reset() {
		s.cpu.Reset()
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
	ebiten.SetWindowTitle("TEC-1G")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatalf("error: %s", err)
	}
}

//-----------------------------------------------------------------------------
