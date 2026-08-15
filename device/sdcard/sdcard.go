//-----------------------------------------------------------------------------
/*

SD Card Driver

*/
//-----------------------------------------------------------------------------

package sdcard

//-----------------------------------------------------------------------------

import (
	"log"
)

//-----------------------------------------------------------------------------

func boolToByte(x bool) byte {
	if x {
		return 1
	}
	return 0
}

//-----------------------------------------------------------------------------
// command responses

const rsp1Idle = (1 << 0)               // in idle state
const rsp1EraseReset = (1 << 1)         // erase reset
const rsp1IllegalCommand = (1 << 2)     // illegal command
const rsp1ComCRCError = (1 << 3)        // com crc error
const rsp1EraseSequenceError = (1 << 4) // erase sequence error
const rsp1AddressError = (1 << 5)       // address error
const rsp1ParameterError = (1 << 6)     // parameter error

//-----------------------------------------------------------------------------

const (
	stateCommand  = iota // rx a command
	stateArgument        // rx the command argument
	stateCRC             // rx the command sequence crc value
)

type SDIO struct {
	enable bool // is the sdcard enabled?

	// spi state variables
	clock     bool // clock state
	dataIn    byte // data in register
	inCount   int  // count of input bits (clock rising edge)
	dataOut   byte // data out register
	outCount  int  // count of output bits (clock falling edge)
	outActive bool // are we actively outputting a bit?
	miso      bool // miso output bit

	// command state variables
	cmdState int    // current command state
	cmd      int    // current command
	arg      uint32 // command argument
	argBytes int    // argument bytes rx-ed
	crc      byte   // calculated crc7 value

	// response buffer
	rsp *circularBuffer
}

func New() (*SDIO, error) {
	return &SDIO{
		enable: true,
		rsp:    newCircularBuffer(1024),
	}, nil
}

//-----------------------------------------------------------------------------
// response buffer

// push a byte into the response FIFO
func (sd *SDIO) wrResponse(val byte) {
	if sd.outActive {
		// we already have a byte in chute
		// add this one to the buffer
		sd.rsp.write(val)
		return
	}
	// no active byte, so this byte is now active
	sd.outActive = true
	sd.dataOut = val
}

//-----------------------------------------------------------------------------

func (sd *SDIO) wrCommand(cmd int, arg uint32) {
	log.Printf("sdio.wrCommand cmd %d arg 0x%08x", cmd, arg)
	switch cmd {
	case 0: // GO_IDLE_STATE, R1 (1 byte)
		sd.wrResponse(rsp1Idle)
	case 8: // SEND_IF_COND, R7 (5 bytes)
		sd.wrResponse(rsp1Idle)              // r1
		sd.wrResponse(0)                     // command version (left as zero)
		sd.wrResponse(0)                     // reserved
		sd.wrResponse(byte((arg >> 8) & 15)) // voltage accepted
		sd.wrResponse(byte(arg & 0xff))      // pattern
	default:
		log.Printf("sdio.wrCommand unknown command %d", cmd)
		sd.wrResponse(rsp1IllegalCommand)
	}
}

//-----------------------------------------------------------------------------

func sdCommand(cmd byte) (int, bool) {
	// check the top 2 bits for 01
	if cmd&0xc0 == 0x40 {
		return int(cmd & 0x3f), true
	}
	return 0, false
}

func (sd *SDIO) wrData(val byte) {
	//log.Printf("sdio.wrData %02x", val)
	switch sd.cmdState {
	case stateCommand:
		if cmd, ok := sdCommand(val); ok {
			sd.crc = crcAdd(0, val)
			sd.cmd = cmd
			sd.arg = 0
			sd.argBytes = 0
			sd.cmdState = stateArgument
		}
	case stateArgument:
		sd.crc = crcAdd(sd.crc, val)
		sd.arg = (sd.arg << 8) | uint32(val)
		sd.argBytes += 1
		if sd.argBytes == 4 {
			sd.cmdState = stateCRC
		}
	case stateCRC:
		if val == (sd.crc<<1)|1 {
			sd.wrCommand(sd.cmd, sd.arg)
		} else {
			log.Printf("sdio.wrData bad crc")
			sd.wrResponse(rsp1ComCRCError)
		}
		sd.cmdState = stateCommand
	}
}

//-----------------------------------------------------------------------------

// read the card detect and miso bits
func (sd *SDIO) Read() (bool, bool) {
	//log.Printf("sdio.Read")
	return sd.enable, sd.miso
}

func (sd *SDIO) Write(chipSelect, mosi, clock bool) {
	if !sd.enable {
		// no device present
		return
	}

	if !chipSelect {
		// chip is not selected
		sd.clock = false
		sd.dataIn = 0
		sd.inCount = 0
		sd.outCount = 0
		// note: don't reset the response data
		return
	}

	//log.Printf("sdio.Write mosi %t clock %t", mosi, clock)

	risingEdge := !sd.clock && clock
	fallingEdge := sd.clock && !clock
	sd.clock = clock

	if risingEdge {
		sd.dataIn <<= 1
		sd.dataIn |= boolToByte(mosi)
		sd.inCount += 1
		if sd.inCount == 8 {
			sd.wrData(sd.dataIn)
			sd.inCount = 0
			sd.dataIn = 0
		}
	}

	if fallingEdge {
		if sd.outActive {
			sd.miso = (sd.dataOut & 0x80) != 0
			sd.dataOut <<= 1
			sd.outCount += 1
			if sd.outCount == 8 {
				sd.outCount = 0
				val, err := sd.rsp.read()
				sd.outActive = err == nil
				sd.dataOut = val
			}
		} else {
			sd.miso = false
		}
	}
}

//-----------------------------------------------------------------------------
