//-----------------------------------------------------------------------------
/*

SD Card Driver

*/
//-----------------------------------------------------------------------------

package sdcard

//-----------------------------------------------------------------------------

import (
	"io"
	"log"
	"os"
)

//-----------------------------------------------------------------------------

func boolToByte(x bool, val byte) byte {
	if x {
		return val
	}
	return 0
}

//-----------------------------------------------------------------------------
// command responses

const rsp1Success = 0                   // success/ready
const rsp1Idle = (1 << 0)               // in idle state
const rsp1EraseReset = (1 << 1)         // erase reset
const rsp1IllegalCommand = (1 << 2)     // illegal command
const rsp1ComCRCError = (1 << 3)        // com crc error
const rsp1EraseSequenceError = (1 << 4) // erase sequence error
const rsp1AddressError = (1 << 5)       // address error
const rsp1ParameterError = (1 << 6)     // parameter error

//-----------------------------------------------------------------------------

const defaultBlockLength = 512

const blockBusy = 0xff           // getting the block data ...
const blockStart = 0xfe          // starting the block data
const blockCardLocked = (1 << 3) // error: the card is locked
const blockCcError = (1 << 2)    // error: internal failure
const blockEccFail = (1 << 1)    // error: ecc failure on block data
const blockOutOfRange = (1 << 0) // error: address is out of range

//-----------------------------------------------------------------------------

// data response tokens
const dataAccepted = (0xe0 | (2 << 1) | 1)
const dataCrcError = (0xe0 | (5 << 1) | 1)
const dataWriteError = (0xe0 | (6 << 1) | 1)

//-----------------------------------------------------------------------------

type rxState int

const (
	stateCommand         rxState = iota // rx a command
	stateCommandArgument                // rx the command argument
	stateCommandCRC                     // rx the command sequence crc value
	stateBlockStart                     // rx a block state
	stateBlockData                      // rx the block data
	stateBlockCRC                       // rx a block CRC
)

type Config struct {
	Enable bool   `toml:"enable"`    // is the sdcard enabled
	SDSC   bool   `toml:"sdsc_mode"` // identify as an SDSC device
	Image  string `toml:"image"`     // path to fat32 file system image
}

type SDIO struct {
	cfg Config // sdcard configuration

	// fat32 file system image
	image *os.File

	// spi state variables
	clock     bool // clock state
	dataIn    byte // data in register
	inCount   int  // count of input bits (clock rising edge)
	dataOut   byte // data out register
	outCount  int  // count of output bits (clock falling edge)
	outActive bool // are we actively outputting a bit?
	miso      bool // miso output bit

	// rx state variables
	rxState    rxState // current receive state
	cmd        int     // current command
	arg        uint32  // command argument
	rxBytes    int     // bytes rx-ed
	crcOn      bool    // are we checking CRCs?
	cmdCRC     byte    // command CRC7 value
	blockCRC   uint16  // block CRC16 value
	appCommand bool    // is the next command an app command?

	// rx-ed block state
	blockLength int    // block length
	blockBuffer []byte // rx-ed block buffer
	rxCRC       uint16 // rx-ed CRC for block
	wrOffset    int64  // write offset for rx-ed block

	// response buffer
	rsp *circularBuffer
}

func New(cfg Config) (*SDIO, error) {

	if !cfg.Enable {
		return &SDIO{cfg: cfg}, nil
	}

	// Open the image file
	file, err := os.OpenFile(cfg.Image, os.O_RDWR, 0644)
	if err != nil {
		log.Printf("sdcard: disabled, %s", err.Error())
		cfg.Enable = false
		return &SDIO{cfg: cfg}, nil
	}

	log.Printf("sdcard: image %s", cfg.Image)
	return &SDIO{
		cfg:         cfg,
		image:       file,
		miso:        true,
		blockLength: defaultBlockLength,
		blockBuffer: make([]byte, defaultBlockLength),
		rsp:         newCircularBuffer(1024),
	}, nil
}

//-----------------------------------------------------------------------------
// file system image operations

// read a block in the file system image
func (sd *SDIO) readBlock(offset int64, buf []byte) error {
	log.Printf("sdcard: sd.readBlock %d bytes @ 0x%08x", len(buf), offset)
	n, err := sd.image.ReadAt(buf, offset)
	if err != nil {
		// io.EOF is expected if you try to read exactly up to the end of the file
		if err == io.EOF && n == len(buf) {
			return nil
		}
		return err
	}
	return nil
}

// write a block in the file system image
func (sd *SDIO) writeBlock(offset int64, buf []byte) error {
	log.Printf("sdcard: sd.writeBlock %d bytes @ 0x%08x", len(buf), offset)
	_, err := sd.image.WriteAt(buf, offset)
	return err
}

//-----------------------------------------------------------------------------
// response buffer

// push a byte into the response FIFO
func (sd *SDIO) wrResponse(val byte) {
	if sd.outActive {
		// we already have a byte in the chute
		// add this one to the buffer
		sd.rsp.write(val)
		return
	}
	// no active byte, so this byte is now active
	sd.outActive = true
	sd.dataOut = val
	sd.outCount = 0
}

//-----------------------------------------------------------------------------

func (sd *SDIO) wrCommand(cmd int, arg uint32) rxState {

	if sd.appCommand {
		// ACMDx processing
		log.Printf("sdio.wrCommand acmd%d arg 0x%08x", cmd, arg)
		sd.appCommand = false
		switch cmd {
		case 41: // SD_SEND_OP_COND, R1
			hcs := arg&(1<<30) != 0
			_ = hcs // and do what?
			sd.wrResponse(rsp1Success)

		default:
			log.Printf("sdio.wrCommand unknown acmd%d", cmd)
			sd.wrResponse(rsp1IllegalCommand)
		}
	} else {
		// CMDx processing
		log.Printf("sdio.wrCommand cmd%d arg 0x%08x", cmd, arg)
		switch cmd {
		case 0: // GO_IDLE_STATE, R1
			sd.wrResponse(rsp1Idle)

		case 8: // SEND_IF_COND, R7 (5 bytes)
			sd.wrResponse(rsp1Success)
			sd.wrResponse(0)                     // command version (left as zero)
			sd.wrResponse(0)                     // reserved
			sd.wrResponse(byte((arg >> 8) & 15)) // voltage accepted
			sd.wrResponse(byte(arg & 0xff))      // pattern

		case 16: // SET_BLOCKLEN, R1 (SDSC only)
			if !sd.cfg.SDSC {
				log.Printf("sdcard: cmd16 for non SDSC, ignoring")
				sd.wrResponse(rsp1IllegalCommand)
				break
			}
			sd.blockLength = int(arg)
			// re-allocate the block buffer, if we need to
			if sd.blockLength != len(sd.blockBuffer) {
				sd.blockBuffer = make([]byte, sd.blockLength)
			}
			log.Printf("sdcard: cmd16 blockLength %d", sd.blockLength)
			sd.wrResponse(rsp1Success)

		case 17: // READ_SINGLE_BLOCK, R1
			sd.wrResponse(rsp1Success)
			sd.wrResponse(blockBusy)
			// work out the image offset
			var ofs int64
			if sd.cfg.SDSC {
				// SDSC has byte addressing
				ofs = int64(arg)
			} else {
				// SDHC has block addressing
				ofs = int64(sd.blockLength) * int64(arg)
			}
			// get the block
			buf := make([]byte, sd.blockLength)
			err := sd.readBlock(ofs, buf)
			if err != nil {
				log.Printf("sdcard: %s", err.Error())
				// Host code *should* handle this error response
				// The mon3 code does not :-(
				sd.wrResponse(blockCcError)
				break
			}
			// write out the block and crc
			sd.wrResponse(blockStart)
			var crc uint16
			for _, b := range buf {
				sd.wrResponse(b)
				crc = addCRC16(crc, b)
			}
			// big endian order
			sd.wrResponse(byte(crc >> 8))
			sd.wrResponse(byte(crc))

		case 24: // WRITE_BLOCK, R1
			sd.wrResponse(rsp1Success)
			if sd.cfg.SDSC {
				// SDSC has byte addressing
				sd.wrOffset = int64(arg)
			} else {
				// SDHC has block addressing
				sd.wrOffset = int64(sd.blockLength) * int64(arg)
			}
			// get the block data
			return stateBlockStart

		case 55: // APP_CMD, R1
			sd.appCommand = true
			sd.wrResponse(rsp1Success)

		case 58: // READ_OCR, R3 (5 bytes)
			sd.wrResponse(rsp1Success)
			sd.wrResponse(0x80 | boolToByte(!sd.cfg.SDSC, 0x40)) // not busy, sdhc/sdsc
			sd.wrResponse(0xff)                                  // 2.8-3.6v
			sd.wrResponse(0)
			sd.wrResponse(0)

		case 59: // CRC_ON_OFF, R1
			sd.crcOn = arg&1 != 0
			sd.wrResponse(rsp1Success)

		default:
			log.Printf("sdio.wrCommand unknown cmd%d", cmd)
			sd.wrResponse(rsp1IllegalCommand)
		}
	}

	// wait for the next command
	return stateCommand
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
	switch sd.rxState {
	case stateCommand:
		// wait for a command byte
		if cmd, ok := sdCommand(val); ok {
			sd.cmdCRC = addCRC7(0, val)
			sd.cmd = cmd
			sd.arg = 0
			sd.rxBytes = 0
			sd.rxState = stateCommandArgument
		}
	case stateCommandArgument:
		// receive the command argument (4 bytes)
		sd.cmdCRC = addCRC7(sd.cmdCRC, val)
		sd.arg = (sd.arg << 8) | uint32(val)
		sd.rxBytes += 1
		if sd.rxBytes == 4 {
			sd.rxState = stateCommandCRC
		}
	case stateCommandCRC:
		// receive and check the command crc7
		// note: cmd0 and cmd8 always have correct crc values
		if !(sd.crcOn || (sd.cmd == 0) || (sd.cmd == 8)) {
			// no crc check, just check for 1 in the lsb.
			val = (sd.cmdCRC << 1) | (val & 1)
		}
		// check the CRC
		if val != (sd.cmdCRC<<1)|1 {
			log.Printf("sdio.wrData bad crc %02x %08x", sd.cmd, sd.arg)
			sd.wrResponse(rsp1ComCRCError)
			sd.rxState = stateCommand
			break
		}
		// process the command
		sd.rxState = sd.wrCommand(sd.cmd, sd.arg)
	case stateBlockStart:
		// wait for the block start byte
		if val == blockStart {
			sd.rxState = stateBlockData
			sd.rxBytes = 0
			sd.blockCRC = 0
		}
	case stateBlockData:
		// receive the block data
		sd.blockCRC = addCRC16(sd.blockCRC, val)
		sd.blockBuffer[sd.rxBytes] = val
		sd.rxBytes += 1
		if sd.rxBytes == len(sd.blockBuffer) {
			sd.rxState = stateBlockCRC
			sd.rxCRC = 0
			sd.rxBytes = 0
		}
	case stateBlockCRC:
		// receive and check the block crc16
		sd.rxCRC = (sd.rxCRC << 8) | uint16(val)
		sd.rxBytes += 1
		if sd.rxBytes == 2 {
			if sd.crcOn && (sd.rxCRC != sd.blockCRC) {
				sd.wrResponse(dataCrcError)
				sd.rxState = stateCommand
				break
			}
			// write the block
			err := sd.writeBlock(sd.wrOffset, sd.blockBuffer)
			if err != nil {
				log.Printf("sdcard: %s", err.Error())
				sd.wrResponse(dataWriteError)
			} else {
				sd.wrResponse(dataAccepted)
			}
			sd.wrResponse(0xff)
			sd.rxState = stateCommand
		}
	}
}

//-----------------------------------------------------------------------------

// read the card detect and miso bits
func (sd *SDIO) Read() (bool, bool) {
	if !sd.cfg.Enable {
		return false, true
	}
	//log.Printf("sdio.Read %t", sd.miso)
	return true, sd.miso
}

func (sd *SDIO) Write(cs, mosi, clock bool) {
	if !sd.cfg.Enable {
		// no device present
		return
	}

	if !cs {
		// chip is not selected
		// reset spi state
		// note: don't reset any response data
		sd.clock = false
		sd.dataIn = 0
		sd.inCount = 0
		// abandon any rx in-progress
		//sd.rxState = stateCommand
		return
	}

	risingEdge := !sd.clock && clock
	fallingEdge := sd.clock && !clock
	sd.clock = clock

	if risingEdge {
		sd.dataIn <<= 1
		sd.dataIn |= boolToByte(mosi, 1)
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
				val, err := sd.rsp.read()
				if err != nil {
					sd.outActive = false
					return
				}
				sd.dataOut = val
				sd.outCount = 0
			}
		} else {
			sd.miso = true
		}
	}
}

//-----------------------------------------------------------------------------
