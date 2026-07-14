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

import (
	"errors"
)

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
	stateWaitForIdle state = iota
	stateIdle
	stateStart
	stateData
	stateParity
	stateStop
)

type Config struct {
	SamplesPerBit int    // Fs / baud, e.g. 16 for 16x oversampling
	DataBits      int    // 5..9
	StopBits      int    // 1..2
	Parity        Parity // parity
	LsbLast       bool   // false for standard UART
	IdleLow       bool   // false for normal TTL/RS-232-logical polarity
}

type frameState struct {
	state     state
	counter   int  // sample counter within current bit
	bitIndex  int  // which data bit we're on
	shift     int  // assembled data
	stopCount int  // stop bits emitted (encode only)
	parityAcc bool // running parity
}

type UART struct {
	config      *Config
	decodeState frameState // sample decoding state
	encodeState frameState // sample encoding state
	frameError  bool       // framing error on decode
	parityError bool       // parity error on decode
}

func NewUART(k *Config) (*UART, error) {
	return &UART{
		config: k,
	}, nil
}

//-----------------------------------------------------------------------------

// WriteSample takes a sample of a serial line and turns it into a serial frame value.
func (u *UART) WriteSample(level bool) (int, error) {
	var rxFrame bool
	s := &u.decodeState

	if u.config.IdleLow {
		level = !level
	}

	//log.Printf("state %d sample %t\n", s.state, level)

	switch s.state {
	case stateWaitForIdle:
		// waiting for the line to be idle (mark)
		if level {
			// we have a mark - we are now idle
			s.state = stateIdle
		}
	case stateIdle:
		// line is idle, waiting for a start bit (space)
		if !level {
			s.state = stateStart
			s.counter = 0
		}
	case stateStart:
		// Sample the start bit at its center to reject noise/glitches.
		s.counter += 1
		if s.counter >= (u.config.SamplesPerBit >> 1) {
			if !level {
				// still low -> valid start
				s.state = stateData
				s.counter = 0
				s.bitIndex = 0
				s.shift = 0
				s.parityAcc = u.config.Parity == ParityOdd
				u.parityError = false
				u.frameError = false
			} else {
				// false start, glitch
				s.state = stateWaitForIdle
			}
		}
	case stateData:
		// Full bit period between successive center samples.
		s.counter += 1
		if s.counter >= u.config.SamplesPerBit {
			s.counter = 0
			s.parityAcc = s.parityAcc != level
			if u.config.LsbLast {
				s.shift = (s.shift << 1) | bool2int(level)
			} else {
				s.shift |= bool2int(level) << s.bitIndex
			}
			s.bitIndex += 1
			if s.bitIndex >= u.config.DataBits {
				if u.config.Parity == ParityNone {
					s.state = stateStop
				} else {
					s.state = stateParity
				}
			}
		}
	case stateParity:
		s.counter += 1
		if s.counter >= u.config.SamplesPerBit {
			s.counter = 0
			u.parityError = level != s.parityAcc
			s.state = stateStop
		}
	case stateStop:
		s.counter += 1
		if s.counter >= u.config.SamplesPerBit {
			s.counter = 0
			u.frameError = level != true // stop bit must be mark (1)
			rxFrame = true
			// Ignore a 2nd stop bit's exact timing
			// return to idle and resync on the next falling edge.
			s.state = stateWaitForIdle
		}
	}

	if !rxFrame {
		return 0, errors.New("no data")
	}
	if u.parityError {
		return 0, errors.New("parity error")
	}
	if u.frameError {
		return 0, errors.New("framing error")
	}

	return s.shift, nil
}

//-----------------------------------------------------------------------------

// ReadSample reads data from a PTY and returns a set of sample values.
func (u *UART) ReadSample(pty *PTY) (bool, error) {
	level := false
	s := &u.encodeState

	switch s.state {
	case stateIdle:
		word, err := pty.Read()
		if err != nil {
			return true, err
		}
		s.shift = int(word)
		s.parityAcc = u.config.Parity == ParityOdd
		s.counter = 0
		s.bitIndex = 0
		s.stopCount = 0
		s.state = stateStart
		level = true
	case stateStart:
		level = false // start bit is always space (0)
		s.counter += 1
		if s.counter >= u.config.SamplesPerBit {
			s.counter = 0
			s.state = stateData
		}
	case stateData:
		// Present the current data bit for the whole bit period.
		if u.config.LsbLast {
			level = (s.shift>>(u.config.DataBits-1-s.bitIndex))&1 != 0
		} else {
			level = (s.shift>>s.bitIndex)&1 != 0
		}
		s.counter += 1
		if s.counter >= u.config.SamplesPerBit {
			s.counter = 0
			s.parityAcc = s.parityAcc != level
			s.bitIndex += 1
			if s.bitIndex >= u.config.DataBits {

				if u.config.Parity == ParityNone {
					s.state = stateStop
				} else {
					s.state = stateParity
				}
			}
		}
	case stateParity:
		level = s.parityAcc
		s.counter += 1
		if s.counter >= u.config.SamplesPerBit {
			s.counter = 0
			s.state = stateStop
		}
	case stateStop:
		level = true // stop bit(s) are mark (1)
		s.counter += 1
		if s.counter >= u.config.SamplesPerBit {
			s.counter = 0
			if s.stopCount >= u.config.StopBits {
				s.state = stateIdle
			}
		}
	}

	if u.config.IdleLow {
		level = !level
	}

	return level, nil
}

//-----------------------------------------------------------------------------
