//-----------------------------------------------------------------------------
/*

SD Card Driver

*/
//-----------------------------------------------------------------------------

package sdcard

import "errors"

//-----------------------------------------------------------------------------

type circularBuffer struct {
	buffer []byte // buffer
	rd, wr int    // read/write indices
	n      int    // current byte count
}

func newCircularBuffer(size int) *circularBuffer {
	return &circularBuffer{
		buffer: make([]byte, size),
	}
}

// Increment and wrap-around a read/write index.
func (c *circularBuffer) incMod(idx int) int {
	idx++
	if idx == len(c.buffer) {
		return 0
	}
	return idx
}

// isFull returns true if the buffer is full
func (c *circularBuffer) isFull() bool {
	return len(c.buffer) == c.n
}

// isEmpty returns true if the buffer is empty
func (c *circularBuffer) isEmpty() bool {
	return c.n == 0
}

// write a byte
func (c *circularBuffer) write(val byte) error {
	if c.isFull() {
		return errors.New("full")
	}
	c.buffer[c.wr] = val
	c.wr = c.incMod(c.wr)
	c.n += 1
	return nil
}

// read a byte
func (c *circularBuffer) read() (byte, error) {
	if c.isEmpty() {
		return 0, errors.New("empty")
	}
	val := c.buffer[c.rd]
	c.rd = c.incMod(c.rd)
	c.n -= 1
	return val, nil
}

// empty the circular buffer of all data
func (c *circularBuffer) empty() {
	c.n = 0
	c.rd = 0
	c.wr = 0
}

//-----------------------------------------------------------------------------
