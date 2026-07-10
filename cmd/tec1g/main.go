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
	"image/png"
	"log"

	"github.com/deadsy/go_z80/device/hd44780"
	"github.com/deadsy/go_z80/device/keyboard"
	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/rtc"
	"github.com/deadsy/go_z80/device/serial"
	"github.com/deadsy/go_z80/device/sevseg"
	"github.com/deadsy/go_z80/device/sixdigit"
	"github.com/deadsy/go_z80/device/sound"
	"github.com/deadsy/go_z80/device/speaker"
	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

//go:embed assets/mon3_2025BC_16.bin assets/tec1g.png
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
	display            *sixdigit.Display // 6 digit display
	led                *led.LED          // speaker activity LED
	speaker            *speaker.Speaker  // audio speaker
	sound              *sound.Sound      // ebiten audio
	lcd                *hd44780.LCD      // lcd
	keyboard           *keyboard.Tec1G   // matrix keyboard
	rtc                *rtc.RTC          // rtc board
	uart               *serial.UART      // serial uart
	pty                *serial.PTY       // pseudo tty
	io                 *sysIO            // system IO
	mem                *sysMemory        // system memory
	bus                *Bus              // system bus
	cpu                *z80.CPU          // z80 cpu
	background         *ebiten.Image     // background graphic
	width, height      int               // window dimensions
	tickCycles         float32           // ebiten tick cpu cycles
	audioSampleCycles  float32           // audio sample cpu cycles
	serialSampleCycles float32           // serial sample cpu cycles
	soundStarted       bool              // has the sound been started?
}

func newSystem() (*system, error) {

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

	// setup the LED
	cfgLED := led.Config{
		XBase:  926,
		YBase:  514,
		Radius: 15,
	}
	led, err := led.New(&cfgLED)
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
	keyboard, err := keyboard.NewTec1G()
	if err != nil {
		return nil, err
	}

	// setup the RTC
	rtc, err := rtc.New()
	if err != nil {
		return nil, err
	}
	rtc.Enable()

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
	fmt.Printf("serial port at %s\n", pty.Name())

	// setup the IO
	io := newIO(display, led, lcd, keyboard, rtc)

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
		display:  display,
		led:      led,
		speaker:  speaker,
		sound:    sound,
		lcd:      lcd,
		keyboard: keyboard,
		rtc:      rtc,
		uart:     uart,
		pty:      pty,
		io:       io,
		mem:      mem,
		bus:      bus,
		cpu:      cpu,
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
	s.pty.Close()
}

func (s *system) Update() error {

	// start the sound (once)
	if !s.soundStarted && s.speaker.Samples() >= 800 {
		s.sound.Start()
		s.soundStarted = true
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
				return fmt.Errorf("speaker.WriteSample: %s", err)
			}
			s.audioSampleCycles += cpuCyclesPerAudioSample
		}

		// sample the serial output
		s.serialSampleCycles -= float32(cycles)
		for s.serialSampleCycles < 0 {
			rx, err := s.uart.WriteSample(s.io.serialTx)
			if err == nil {
				s.pty.Write(byte(rx))
			}
			s.serialSampleCycles += cpuCyclesPerSerialSample
		}
	}

	s.display.Update()
	s.led.Update()
	s.lcd.Update()
	s.keyboard.Update()
	return nil
}

func (s *system) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.background, nil)
	s.display.Draw(screen)
	s.led.Draw(screen)
	s.lcd.Draw(screen)
}

func (s *system) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.width, s.height
}

func main() {
	s, err := newSystem()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer s.Exit()

	ebiten.SetWindowSize(s.width, s.height)
	ebiten.SetWindowTitle("TEC-1G")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatalf("error: %s", err)
	}
}

//-----------------------------------------------------------------------------
