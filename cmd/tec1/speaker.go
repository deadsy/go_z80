//-----------------------------------------------------------------------------
/*

TEC-1 Emulation

Speaker Audio

*/
//-----------------------------------------------------------------------------

package main

import "errors"

//-----------------------------------------------------------------------------

type circularBuffer struct {
	buffer []byte
	rd, wr int
}

func newCircularBuffer(size int) *circularBuffer {
	return &circularBuffer{
		buffer: make([]byte, size),
	}
}

func (c *circularBuffer) write(val byte) error {
	wr := (c.wr + 1) % len(c.buffer)
	if wr == c.rd {
		return errors.New("full")
	}
	c.buffer[c.wr] = val
	c.wr = wr
	return nil
}

func (c *circularBuffer) read() (byte, error) {
	if c.rd == c.wr {
		return 0, errors.New("empty")
	}
	val := c.buffer[c.rd]
	c.rd = (c.rd + 1) % len(c.buffer)
	return val, nil
}

//-----------------------------------------------------------------------------

const bitAmplitude = 3000

func bitToSample(bit bool) (int16, int16) {
	if bit {
		return bitAmplitude, bitAmplitude
	}
	return -bitAmplitude, -bitAmplitude
}

//-----------------------------------------------------------------------------

type speaker struct {
	buffer *circularBuffer
}

func newSpeaker() *speaker {
	return &speaker{
		buffer: newCircularBuffer(4096),
	}
}

// Read samples from the buffer (implements io.Reader)
func (s *speaker) Read(b []byte) (n int, err error) {
	for i := 0; i < len(b); i++ {
		val, err := s.buffer.read()
		if err != nil {
			// emptied the sample buffer
			return i, nil
		}
		b[i] = val
	}
	// filled the provided buffer
	return len(b), nil
}

// write a bit sample to the buffer
func (s *speaker) writeSample(bit bool) error {
	l, r := bitToSample(bit)
	err := s.buffer.write(byte(l))
	if err != nil {
		return err
	}
	err = s.buffer.write(byte(l >> 8))
	if err != nil {
		return err
	}
	err = s.buffer.write(byte(r))
	if err != nil {
		return err
	}
	return s.buffer.write(byte(r >> 8))
}

//-----------------------------------------------------------------------------
