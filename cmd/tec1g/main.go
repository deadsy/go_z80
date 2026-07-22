//-----------------------------------------------------------------------------
/*

TEC-1G Emulator

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
	"github.com/deadsy/go_z80/cmd/tec1g/keypad"
	"github.com/deadsy/go_z80/device/array"
	"github.com/deadsy/go_z80/device/array88"
	"github.com/deadsy/go_z80/device/disco"
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
	cpu                *z80.CPU         // z80 cpu
	background         *ebiten.Image    // background graphic
	width, height      int              // window dimensions
	totalCycles        uint64           // total cpu cycles
	tickCycles         float32          // ebiten tick cpu cycles
	audioSampleCycles  float32          // audio sample cpu cycles
	serialSampleCycles float32          // serial sample cpu cycles
	soundStarted       bool             // has the sound been started?
	haltLogged         bool             // have we logged a cpu halt?
}

func newSystem(cfg *Config) (*system, error) {

	// setup the memory
	mem, err := newMemory()
	if err != nil {
		return nil, err
	}

	// setup the bus
	bus := newBus()

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

	// setup the LCD
	cfgLCD := hd44780.Config{
		Rows:           4,
		Cols:           20,
		XBase:          224,
		YBase:          577,
		XScale:         0.34,
		YScale:         0.34,
		CharacterColor: color.RGBA{0xda, 0xe7, 0xe9, 255},
	}
	lcd, err := hd44780.New(cfgLCD)
	if err != nil {
		return nil, err
	}

	// setup the keyboard
	keyboard, err := keyboard.New(cfg.DIP.K)
	if err != nil {
		return nil, err
	}

	// setup the keypad
	keypad, err := keypad.New(!cfg.DIP.K)
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

	// setup the Bar LED
	cfgBarLed := array.Config{
		Enable: true,
		Rows:   1,
		Cols:   10,
		Type:   led.Rectangle,
		X:      759.5,
		Y:      794,
		XGap:   7,
		Width:  7.5,
		Height: 31,
		On:     color.RGBA{0, 0, 255, 128},
	}
	ledBar, err := array.New(cfgBarLed)
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

	// setup the disco lights
	cfgLedDisco := disco.Config{
		X:      537,
		Y:      852,
		YGap:   49,
		Width:  10.5,
		Height: 26,
	}
	ledDisco, err := disco.New(cfgLedDisco)
	if err != nil {
		return nil, err
	}

	// setup the IO
	devices := ioDevices{
		display:    display,
		ledSpeaker: ledSpeaker,
		ledHalt:    ledHalt,
		ledBar:     ledBar,
		ledArray:   ledArray,
		ledDisco:   ledDisco,
		lcd:        lcd,
		keyboard:   keyboard,
		keypad:     keypad,
		rtc:        rtc,
	}
	io := newIO(&devices)

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
		cpu:     cpu,
	}

	io.setSystem(s)
	io.setDIP(cfg.DIP)

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

func (s *system) GetCpuCycles() uint64 {
	return s.totalCycles
}

func (s *system) Update() error {

	// start the sound (once)
	if !s.soundStarted && s.sound.IsReady() && s.speaker.Samples() >= 800 {
		s.speaker.Empty()
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
		s.totalCycles += uint64(cycles)
		s.tickCycles -= float32(cycles)

		// sample the audio output
		s.audioSampleCycles -= float32(cycles)
		for s.audioSampleCycles < 0 {
			err := s.speaker.WriteSample(s.io.speaker)
			if err != nil {
				log.Printf("speaker.WriteSample %s", err)
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
				log.Printf("uart.WriteSample %s", err)
			}
			// drive the rx line *to* the tec-1g
			sample, err := s.uart.ReadSample(s.pty)
			if err != nil {
				log.Printf("uart.ReadSample %s", err)
			}
			s.io.serialRx = sample
			s.serialSampleCycles += cpuCyclesPerSerialSample
		}
	}

	// cpu halted?
	halted := s.cpu.IsHalted()
	s.io.dev.ledHalt.Control(halted)
	s.io.dev.ledBar.Control(0, 9, halted)
	if halted && !s.haltLogged {
		log.Printf("cpu halted at pc=0x%04x", s.cpu.PC)
		s.haltLogged = true
	}

	// update the IO devices
	s.io.Update()

	if s.io.dev.keyboard.Reset() || s.io.dev.keypad.Reset() {
		s.haltLogged = false
		s.mem.Reset()
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
	ebiten.SetWindowTitle("TEC-1G")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatalf("error: %s", err)
	}
}

//-----------------------------------------------------------------------------
