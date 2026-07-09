//-----------------------------------------------------------------------------
/*

1-Bit Speaker Audio

*/
//-----------------------------------------------------------------------------

package speaker

import (
	"errors"
	"math"
	"sync"
)

//-----------------------------------------------------------------------------

// clip and convert samples to the -32767..32767 range.
func clipConvert(x float32) int16 {
	if x > 1.0 {
		x = 1.0
	}
	if x < -1.0 {
		x = -1.0
	}
	return int16(x * float32(32767))
}

//-----------------------------------------------------------------------------

type circularBuffer struct {
	lock   sync.Mutex // buffer is read from multiple threads
	buffer []byte     // sample buffer
	rd, wr int        // read/write indices
	n      int        // current byte count
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

// isFull tells you if there is no space for a sample write
func (c *circularBuffer) isFull() bool {
	return (len(c.buffer) - c.n) < 4
}

// isEmpty tells you if there is a no sample read available
func (c *circularBuffer) isEmpty() bool {
	return c.n < 4
}

// write a left/right sample as an atomic operation
func (c *circularBuffer) writeSample(l, r int16) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.isFull() {
		return errors.New("full")
	}
	c.buffer[c.wr] = byte(l)
	c.wr = c.incMod(c.wr)
	c.buffer[c.wr] = byte(l >> 8)
	c.wr = c.incMod(c.wr)
	c.buffer[c.wr] = byte(r)
	c.wr = c.incMod(c.wr)
	c.buffer[c.wr] = byte(r >> 8)
	c.wr = c.incMod(c.wr)
	c.n += 4
	return nil
}

// read a left/right sample (4 bytes) into the passed buffer
func (c *circularBuffer) readSample(b []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.isEmpty() {
		return errors.New("empty")
	}
	b[0] = c.buffer[c.rd]
	c.rd = c.incMod(c.rd)
	b[1] = c.buffer[c.rd]
	c.rd = c.incMod(c.rd)
	b[2] = c.buffer[c.rd]
	c.rd = c.incMod(c.rd)
	b[3] = c.buffer[c.rd]
	c.rd = c.incMod(c.rd)
	c.n -= 4
	return nil
}

// return the number of frames in the buffer
func (c *circularBuffer) frames() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.n / 4
}

//-----------------------------------------------------------------------------
// Low Pass Filter

type lowPassFilter struct {
	// Feedforward coefficients
	b0, b1, b2 float32
	// Feedback coefficients
	a1, a2 float32
	// State (Direct Form II Transposed)
	z1, z2 float32
}

// newLowPassFilter creates a Butterworth low-pass filter.
func newLowPassFilter(sampleRate, cutoff float64) *lowPassFilter {
	// Butterworth Q
	const Q = 0.7071067811865476

	w0 := 2.0 * math.Pi * cutoff / sampleRate
	cosw0 := math.Cos(w0)
	sinw0 := math.Sin(w0)

	alpha := sinw0 / (2.0 * Q)

	b0 := (1 - cosw0) / 2
	b1 := 1 - cosw0
	b2 := (1 - cosw0) / 2

	a0 := 1 + alpha
	a1 := -2 * cosw0
	a2 := 1 - alpha

	// Normalize by a0
	return &lowPassFilter{
		b0: float32(b0 / a0),
		b1: float32(b1 / a0),
		b2: float32(b2 / a0),
		a1: float32(a1 / a0),
		a2: float32(a2 / a0),
	}
}

// Process filters one sample.
func (f *lowPassFilter) process(x float32) float32 {
	y := f.b0*x + f.z1
	f.z1 = f.b1*x - f.a1*y + f.z2
	f.z2 = f.b2*x - f.a2*y
	return y
}

func (f *lowPassFilter) reset() {
	f.z1 = 0
	f.z2 = 0
}

//-----------------------------------------------------------------------------
// DC blocker

type blocker struct {
	r     float32
	lastX float32
	lastY float32
}

func newBlocker(sampleRate, cutoff float64) *blocker {
	r := math.Exp(-2.0 * math.Pi * cutoff / sampleRate)
	return &blocker{
		r: float32(r),
	}
}

func (d *blocker) process(x float32) float32 {
	y := x - d.lastX + d.r*d.lastY
	d.lastX = x
	d.lastY = y
	return y
}

func (d *blocker) reset() {
	d.lastX = 0
	d.lastY = 0
}

//-----------------------------------------------------------------------------

type Config struct {
	BitAmplitude float32
	BufferSize   int
	SampleRate   int
	HighCutoff   int
	LowCutoff    int
}

type Speaker struct {
	config *Config         // speaker configuration
	buffer *circularBuffer // circular buffer of sample values
	lpf    *lowPassFilter  // low pass filter - remove high frequency components
	block  *blocker        // dc block - give 0 average output
}

func New(k *Config) (*Speaker, error) {
	if k.BitAmplitude <= 0 {
		return nil, errors.New("bad bit amplitude")
	}
	// ensure the buffer size is a multiple of left/right audio samples (4 bytes)
	if (k.BufferSize <= 0) || (k.BufferSize%4 != 0) {
		return nil, errors.New("invalid buffer size")
	}
	if k.SampleRate <= 0 {
		return nil, errors.New("invalid sample rate")
	}
	if k.HighCutoff <= 0 {
		return nil, errors.New("invalid high cutoff frequency")
	}
	if k.LowCutoff <= 0 {
		return nil, errors.New("invalid low cutoff frequency")
	}
	return &Speaker{
		config: k,
		buffer: newCircularBuffer(k.BufferSize),
		lpf:    newLowPassFilter(float64(k.SampleRate), float64(k.HighCutoff)),
		block:  newBlocker(float64(k.SampleRate), float64(k.LowCutoff)),
	}, nil
}

// Read samples from the buffer (implements io.Reader)
func (s *Speaker) Read(b []byte) (n int, err error) {
	// read complete left/right samples (4 bytes) at a time
	var ofs int
	for ofs = 0; ofs+4 <= len(b); ofs += 4 {
		err := s.buffer.readSample(b[ofs:])
		if err != nil {
			// emptied the sample buffer
			return ofs, nil
		}
	}
	// filled the provided buffer
	return ofs, nil
}

// write a bit sample to the buffer
func (s *Speaker) WriteSample(bit bool) error {
	// start with a square wave
	sample := s.config.BitAmplitude
	if !bit {
		sample = -sample
	}
	// low pass
	sample = s.lpf.process(sample)
	// dc block
	sample = s.block.process(sample)
	// left and right channels are the same
	x := clipConvert(sample)
	// buffer the sample
	return s.buffer.writeSample(x, x)
}

// return the number of samples in the circular buffer
func (s *Speaker) Samples() int {
	return s.buffer.frames()
}

//-----------------------------------------------------------------------------
