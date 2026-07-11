//-----------------------------------------------------------------------------
/*

Jupiter ACE Emulator

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
	"math"

	"github.com/deadsy/go_z80/device/keyboard"
	"github.com/deadsy/go_z80/device/sound"
	"github.com/deadsy/go_z80/device/speaker"
	"github.com/deadsy/go_z80/device/video"
	"github.com/deadsy/go_z80/util"
	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

//go:embed assets/jace.bin assets/keyboard.png
var assets embed.FS

//-----------------------------------------------------------------------------

const Hz = 1
const kHz = 1000 * Hz
const MHz = kHz * kHz

const cpuClock = 3.25 * MHz

// ebiten timing
const tickRate = 60 * Hz
const cpuCyclesPerTick = float32(cpuClock) / float32(tickRate) // cpu cycles per ebiten update tick

// audio timing
const audioSampleRate = 48000 // samples/sec
const cpuCyclesPerAudioSample = float32(cpuClock) / float32(audioSampleRate)

// periodic interrupt
const interruptRate = 50 * Hz
const cpuCyclesPerInterrupt = float32(cpuClock) / float32(interruptRate) // cpu cycles per interrupt tick

//-----------------------------------------------------------------------------

const videoWidth = 800.0
const videoBorder = 40.0
const videoHeight = videoWidth * (24.0 / 32.0)
const videoScale = videoWidth / (32 * 8)

func buildBackgroundImage() (*ebiten.Image, error) {

	const keyboardImageWidth = 552
	const keyboardScale = videoWidth / keyboardImageWidth
	const keyboardHeight = 224 * keyboardScale

	xSize := int(math.Floor(videoWidth + (2.0 * videoBorder)))
	ySize := int(math.Floor(videoHeight + (3.0 * videoBorder) + keyboardHeight))

	img := ebiten.NewImage(xSize, ySize)
	img.Fill(color.RGBA{0xf9, 0xf9, 0xf9, 255})

	// get the keyboard image
	imgData, err := assets.ReadFile("assets/keyboard.png")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded image: %w", err)
	}
	decodedImg, err := png.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, fmt.Errorf("failed to decode PNG image: %w", err)
	}
	kbd := ebiten.NewImageFromImage(decodedImg)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(keyboardScale, keyboardScale)
	op.GeoM.Translate(videoBorder, (2.0*videoBorder)+videoHeight)
	op.Filter = ebiten.FilterLinear
	img.DrawImage(kbd, op)

	return img, nil
}

//-----------------------------------------------------------------------------

type system struct {
	speaker           *speaker.Speaker // audio speaker
	sound             *sound.Sound     // ebiten audio
	keyboard          *keyboard.Jace   // matrix keyboard
	video             *video.Video     // video
	io                *sysIO           // system IO
	mem               *sysMemory       // system memory
	bus               *Bus             // system bus
	cpu               *z80.CPU         // z80 cpu
	background        *ebiten.Image    // background graphic
	width, height     int              // window dimensions
	tickCycles        float32          // ebiten tick cpu cycles
	audioSampleCycles float32          // audio sample cpu cycles
	interruptCycles   float32          // periodic interrupt
	soundStarted      bool             // has the sound been started?
}

func newSystem() (*system, error) {

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

	// setup the video
	kVideo := video.Config{
		XBase:  videoBorder,
		YBase:  videoBorder,
		XScale: videoScale,
		YScale: videoScale,
	}
	video, err := video.New(&kVideo, mem)
	if err != nil {
		return nil, err
	}

	// setup the bus
	bus := newBus()

	// setup the cpu
	cpu := z80.New(io, mem, bus)

	s := &system{
		speaker:         speaker,
		sound:           sound,
		keyboard:        keyboard,
		video:           video,
		io:              io,
		mem:             mem,
		bus:             bus,
		cpu:             cpu,
		interruptCycles: cpuCyclesPerInterrupt,
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
}

func (s *system) Update() error {

	// start the sound (once)
	if !s.soundStarted && s.sound.IsReady() && s.speaker.Samples() >= 800 {
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
				log.Printf("speaker.WriteSample: %s", err)
				s.speaker.Empty()
			}
			s.audioSampleCycles += cpuCyclesPerAudioSample
		}
	}

	if s.mem.IsDirty() {
		// update the font atlas
		s.video.Update()
		s.mem.Clean()
	}

	s.keyboard.Update()
	return nil
}

func (s *system) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.background, nil)
	s.video.Draw(screen)
}

func (s *system) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.width, s.height
}

//-----------------------------------------------------------------------------

func main() {
	log.Printf("%s\n", util.GetBuildInfo())
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
