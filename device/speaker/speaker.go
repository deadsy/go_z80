//-----------------------------------------------------------------------------
/*

1 Bit Speaker Audio

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
	lock   sync.Mutex
	buffer []byte
	rd, wr int
}

func newCircularBuffer(size int) *circularBuffer {
	return &circularBuffer{
		buffer: make([]byte, size),
	}
}

// Increment and wrap-around an index value.
func incMod(idx, size int) int {
	idx++
	if idx == size {
		return 0
	}
	return idx
}

func (c *circularBuffer) write(val byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	wrInc := incMod(c.wr, len(c.buffer))
	if wrInc == c.rd {
		return errors.New("full")
	}
	c.buffer[c.wr] = val
	c.wr = wrInc
	return nil
}

func (c *circularBuffer) read() (byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.rd != c.wr {
		val := c.buffer[c.rd]
		c.rd = incMod(c.rd, len(c.buffer))
		return val, nil
	}
	return 0, errors.New("empty")
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

func New(k *Config) *Speaker {
	return &Speaker{
		config: k,
		buffer: newCircularBuffer(k.BufferSize),
		lpf:    newLowPassFilter(float64(k.SampleRate), float64(k.HighCutoff)),
		block:  newBlocker(float64(k.SampleRate), float64(k.LowCutoff)),
	}
}

// Read samples from the buffer (implements io.Reader)
func (s *Speaker) Read(b []byte) (n int, err error) {
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
func (s *Speaker) WriteSample(bit bool) error {
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
	l, r := x, x

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
