//-----------------------------------------------------------------------------
/*

Bit Banged Serial

On the Tx side the emulated device generates a bit-banged Tx-ed serial stream.
We convert this back into characters.

On the Rx side the emulated device samples an serial line.
We create this from characters.

*/
//-----------------------------------------------------------------------------

package serial

import "errors"

//-----------------------------------------------------------------------------

func bool2int(x bool) int {
	if x {
		return 1
	}
	return 0
}

//-----------------------------------------------------------------------------

type Parity int

const (
	ParityNone Parity = iota
	ParityEven
	ParityOdd
)

type state int

const (
	stateIdle state = iota
	stateStart
	stateData
	stateParity
	stateStop
)

type Config struct {
	SamplesPerBit int    // Fs / baud, e.g. 16 for 16x oversampling
	DataBits      int    // 5..9
	Parity        Parity // parity
	LsbLast       bool   // false for standard UART
	IdleLow       bool   // false for normal TTL/RS-232-logical polarity
}

type UART struct {
	config *Config
	// runtime state
	state     state
	counter   int  // sample counter within current bit
	bitIndex  int  // which data bit we're on
	shift     int  // assembled data
	parityAcc bool // running parity
	// results of the last completed frame
	frameError  bool
	parityError bool
}

func NewUART(k *Config) (*UART, error) {
	return &UART{
		config: k,
	}, nil
}

// Rx serial line samples and convert them back into frame values
func (s *UART) Sample(level bool) (int, error) {
	var rxFrame bool

	if s.config.IdleLow {
		level = !level
	}

	switch s.state {
	case stateIdle:
		// Wait for the falling edge into the start bit (mark -> space).
		if !level {
			s.state = stateStart
			s.counter = 0
		}
	case stateStart:
		// Sample the start bit at its center to reject noise/glitches.
		s.counter += 1
		if s.counter >= (s.config.SamplesPerBit >> 1) {
			if !level {
				// still low -> valid start
				s.state = stateData
				s.counter = 0
				s.bitIndex = 0
				s.shift = 0
				s.parityAcc = s.config.Parity == ParityOdd
				s.parityError = false
				s.frameError = false
			} else {
				// false start, glitch
				s.state = stateIdle
			}
		}
	case stateData:
		// Full bit period between successive center samples.
		s.counter += 1
		if s.counter >= s.config.SamplesPerBit {
			s.counter = 0
			s.parityAcc = s.parityAcc != level
			if s.config.LsbLast {
				s.shift = (s.shift << 1) | bool2int(level)
			} else {
				s.shift |= bool2int(level) << s.bitIndex
			}
			s.bitIndex += 1
			if s.bitIndex >= s.config.DataBits {
				if s.config.Parity == ParityNone {
					s.state = stateStop
				} else {
					s.state = stateParity
				}
			}
		}
	case stateParity:
		s.counter += 1
		if s.counter >= s.config.SamplesPerBit {
			s.counter = 0
			s.parityError = level != s.parityAcc
			s.state = stateStop
		}
	case stateStop:
		s.counter += 1
		if s.counter >= s.config.SamplesPerBit {
			s.counter = 0
			s.frameError = level != true // stop bit must be mark (1)
			rxFrame = true
			// Ignore a 2nd stop bit's exact timing
			// return to idle and resync on the next falling edge.
			s.state = stateIdle
		}
	}

	if !rxFrame {
		return 0, errors.New("no data")
	}
	if s.parityError {
		return 0, errors.New("parity error")
	}
	if s.frameError {
		return 0, errors.New("framing error")
	}
	return s.shift, nil
}

//-----------------------------------------------------------------------------
