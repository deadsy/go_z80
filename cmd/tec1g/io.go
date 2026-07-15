//-----------------------------------------------------------------------------
/*

TEC-1G IO Devices

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/deadsy/go_z80/cmd/tec1g/keyboard"
	"github.com/deadsy/go_z80/cmd/tec1g/keypad"
	"github.com/deadsy/go_z80/device/ds1302"
	"github.com/deadsy/go_z80/device/hd44780"
	"github.com/deadsy/go_z80/device/led"
	"github.com/deadsy/go_z80/device/sixdigit"
	"github.com/hajimehoshi/ebiten/v2"
)

//-----------------------------------------------------------------------------

func boolToByte(x bool, val byte) byte {
	if x {
		return val
	}
	return 0
}

func boolToMsg(x bool, lo, hi string) string {
	if x {
		return hi
	}
	return lo
}

//-----------------------------------------------------------------------------
// ports

const keypadPort = 0x00   // keypad scan values
const digitPort = 0x01    // display digit enable
const segmentPort = 0x02  // display segment enable
const simpPort = 0x03     // General SIMP Input
const lcdCmdPort = 0x04   // LCD Display command
const x88Port = 0x05      // 8x8 X-axis display latch
const y88Port = 0x06      // 8x8 Y-axis display latch
const glcdPort0 = 0x07    // GLCD port
const lcdDataPort = 0x84  // LCD Display data
const glcdPort1 = 0x87    // GLCD port
const rtcPort = 0xfc      // GPIO Real Time Clock
const sdCardPort = 0xfd   // GPIO SD Card
const keyboardPort = 0xfe // Matrix Keyboard Input
const systemPort = 0xff   // System Latch

// digitPort
const digitMask = uint8(0x3f)      // D0..D5 digits
const serialTxMask = uint8(1 << 6) // D6, serialTx
const speakerMask = uint8(1 << 7)  // D7, speaker/led

// simpPort
const simpConfigK = byte(1 << 0)  // D0, 0 == encoder, 1 == matrix
const simpConfigP = byte(1 << 1)  // D1, 1 == protect memory
const simpConfigE = byte(1 << 2)  // D2
const simpExpand = byte(1 << 3)   // D3
const simpCart = byte(1 << 4)     // D4
const simpGimp = byte(1 << 5)     // D5
const simpKDA = byte(1 << 6)      // D6
const serialRxMask = byte(1 << 7) // D7

// rtcPort
const rtcOutMask = byte(1 << 0)    // D0
const rtcEnableMask = byte(1 << 4) // D4 active high
const rtcClockMask = byte(1 << 6)  // D6
const rtcInMask = byte(1 << 7)     // D7

//-----------------------------------------------------------------------------

type ioDevices struct {
	display    *sixdigit.Display  // 6 digit display
	ledSpeaker *led.LED           // speaker led
	ledHalt    *led.LED           // halt led
	lcd        *hd44780.LCD       // LCD
	keyboard   *keyboard.Keyboard // matrix keyboard
	keypad     *keypad.Keypad     // 74c923 keypad
	rtc        *ds1302.RTC        // realtime clock
}

type sysIO struct {
	dev      *ioDevices
	segment  uint8 // latched segment enable
	digit    uint8 // latched digit enable
	speaker  bool  // latched speaker/led enable
	serialTx bool  // serial tx line
	serialRx bool  // serial rx line
	kpe      byte  // dip switch KPE settings
}

func newIO(dev *ioDevices) *sysIO {
	return &sysIO{
		dev: dev,
	}
}

//-----------------------------------------------------------------------------
// DIP Switch KPE configuration

type dipSwitch struct {
	K bool `toml:"keyboard"`
	P bool `toml:"protect"`
	E bool `toml:"expansion"`
}

// set the values on the dip switch
func (io *sysIO) setDIP(cfg dipSwitch) {
	log.Printf("keyboard is %s", boolToMsg(cfg.K, "74c923", "matrix"))
	log.Printf("protect is %s", boolToMsg(cfg.P, "off", "on"))
	log.Printf("expansion is %s", boolToMsg(cfg.E, "low", "high"))
	io.kpe = boolToByte(cfg.K, simpConfigK) |
		boolToByte(cfg.P, simpConfigP) |
		boolToByte(cfg.E, simpConfigE)
}

//-----------------------------------------------------------------------------

// Read8 reads a byte from an IO port.
func (io *sysIO) Read8(adr uint16) uint8 {
	dev := io.dev
	row := uint8(adr >> 8)
	adr &= 0xff
	switch adr {
	case keypadPort:
		return 0
	case lcdCmdPort:
		return dev.lcd.ReadCommand()
	case simpPort:
		return io.kpe | boolToByte(io.serialRx, serialRxMask)
	case rtcPort:
		return boolToByte(dev.rtc.Read(), rtcOutMask)
	case sdCardPort:
		// TODO
		return 0
	case keyboardPort:
		code, err := dev.keyboard.Scan(row)
		if err != nil {
			log.Printf("keyboard scan error: %s\n", err)
		}
		return code
	}
	log.Printf("io.Read8 unknown port %02x\n", adr)
	return 0
}

// Write8 writes a byte to an IO port.
func (io *sysIO) Write8(adr uint16, val uint8) {
	dev := io.dev
	adr &= 0xff
	switch adr {
	case digitPort:
		io.digit = val & digitMask
		io.speaker = (val & speakerMask) != 0
		io.serialTx = (val & serialTxMask) != 0
		dev.display.Enable(io.digit, io.segment)
		dev.ledSpeaker.Control(io.speaker)
		return
	case segmentPort:
		io.segment = val
		dev.display.Enable(io.digit, io.segment)
		return
	case lcdCmdPort:
		dev.lcd.WriteCommand(val)
		return
	case x88Port:
		// TODO
		return
	case y88Port:
		// TODO
		return
	case glcdPort0, glcdPort1:
		// TODO
		return
	case lcdDataPort:
		dev.lcd.WriteData(val)
		return
	case rtcPort:
		ce := val&rtcEnableMask != 0
		clk := val&rtcClockMask != 0
		in := val&rtcInMask != 0
		dev.rtc.Write(ce, clk, in)
		return
	case sdCardPort:
		// TODO
		return
	case systemPort:
		// TODO
		return

	}
	log.Printf("io.Write8 [%02x] = %02x\n", adr, val)
}

//-----------------------------------------------------------------------------
// ebiten api

func (io *sysIO) Update() {
	io.dev.display.Update()
	io.dev.ledSpeaker.Update()
	io.dev.ledHalt.Update()
	io.dev.lcd.Update()
	io.dev.keyboard.Update()
	io.dev.keypad.Update()
}

func (io *sysIO) Draw(screen *ebiten.Image) {
	io.dev.display.Draw(screen)
	io.dev.ledSpeaker.Draw(screen)
	io.dev.ledHalt.Draw(screen)
	io.dev.lcd.Draw(screen)
}

//-----------------------------------------------------------------------------
