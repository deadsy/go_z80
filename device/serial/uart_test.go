//-----------------------------------------------------------------------------
/*

UART Tests

*/
//-----------------------------------------------------------------------------

package serial

import "testing"

//-----------------------------------------------------------------------------
// test helpers

func parityBit(cfg Config, value int) bool {
	parityAcc := cfg.Parity == ParityOdd
	for i := 0; i < cfg.DataBits; i++ {
		var level bool
		if cfg.LsbLast {
			level = (value>>(cfg.DataBits-1-i))&1 != 0
		} else {
			level = (value>>i)&1 != 0
		}
		parityAcc = parityAcc != level
	}
	return parityAcc
}

func emitFrameSamples(cfg Config, value int, parityOverride *bool, stopOverride *bool) []bool {
	bits := make([]bool, 0, 1+cfg.DataBits+1+cfg.StopBits)

	// start bit
	bits = append(bits, false)

	// data bits
	for i := 0; i < cfg.DataBits; i++ {
		if cfg.LsbLast {
			bits = append(bits, (value>>(cfg.DataBits-1-i))&1 != 0)
		} else {
			bits = append(bits, (value>>i)&1 != 0)
		}
	}

	// optional parity bit
	if cfg.Parity != ParityNone {
		p := parityBit(cfg, value)
		if parityOverride != nil {
			p = *parityOverride
		}
		bits = append(bits, p)
	}

	// stop bit(s)
	stopBit := true
	if stopOverride != nil {
		stopBit = *stopOverride
	}
	for i := 0; i < cfg.StopBits; i++ {
		bits = append(bits, stopBit)
	}

	// one leading and trailing idle bit period to establish/resync decode idle state
	idle := true
	samples := make([]bool, 0, (len(bits)+2)*cfg.SamplesPerBit)
	for i := 0; i < cfg.SamplesPerBit; i++ {
		samples = append(samples, idle)
	}
	for _, bit := range bits {
		for i := 0; i < cfg.SamplesPerBit; i++ {
			samples = append(samples, bit)
		}
	}
	for i := 0; i < cfg.SamplesPerBit; i++ {
		samples = append(samples, idle)
	}

	if cfg.IdleLow {
		for i := range samples {
			samples[i] = !samples[i]
		}
	}

	return samples
}

func feedSamples(u *UART, samples []bool) (int, bool, error) {
	for _, sample := range samples {
		value, gotFrame, err := u.WriteSample(sample)
		if err != nil || gotFrame {
			return value, gotFrame, err
		}
	}
	return 0, false, nil
}

//-----------------------------------------------------------------------------

func TestUARTRoundTripDecodeKnownBytes(t *testing.T) {
	values := []int{0x00, 0x01, 0x55, 0xA6, 0xFF}

	for _, dataBits := range []int{5, 7, 8} {
		for _, stopBits := range []int{1, 2} {
			for _, parity := range []Parity{ParityNone, ParityEven, ParityOdd} {
				for _, lsbLast := range []bool{false, true} {
					for _, idleLow := range []bool{false, true} {
						cfg := &Config{
							SamplesPerBit: 8,
							DataBits:      dataBits,
							StopBits:      stopBits,
							Parity:        parity,
							LsbLast:       lsbLast,
							IdleLow:       idleLow,
						}
						u, err := NewUART(cfg)
						if err != nil {
							t.Fatalf("NewUART() error: %v", err)
						}

						mask := (1 << dataBits) - 1
						for _, value := range values {
							want := value & mask
							samples := emitFrameSamples(*cfg, value, nil, nil)
							got, gotFrame, err := feedSamples(u, samples)
							if err != nil {
								t.Fatalf("cfg=%+v value=0x%x: unexpected error: %v", *cfg, value, err)
							}
							if !gotFrame {
								t.Fatalf("cfg=%+v value=0x%x: expected frame, got none", *cfg, value)
							}
							if got != want {
								t.Fatalf("cfg=%+v value=0x%x: got 0x%x want 0x%x", *cfg, value, got, want)
							}
						}
					}
				}
			}
		}
	}
}

func TestUARTParityError(t *testing.T) {
	cfg := &Config{
		SamplesPerBit: 8,
		DataBits:      8,
		StopBits:      1,
		Parity:        ParityEven,
		LsbLast:       false,
		IdleLow:       false,
	}
	u, err := NewUART(cfg)
	if err != nil {
		t.Fatalf("NewUART() error: %v", err)
	}

	value := 0x5A
	wrongParity := !parityBit(*cfg, value)
	samples := emitFrameSamples(*cfg, value, &wrongParity, nil)

	_, gotFrame, err := feedSamples(u, samples)
	if err == nil || err.Error() != "parity error" {
		t.Fatalf("expected parity error, got: %v", err)
	}
	if gotFrame {
		t.Fatalf("expected gotFrame=false on parity error")
	}
}

func TestUARTFramingError(t *testing.T) {
	cfg := &Config{
		SamplesPerBit: 8,
		DataBits:      8,
		StopBits:      1,
		Parity:        ParityNone,
		LsbLast:       false,
		IdleLow:       false,
	}
	u, err := NewUART(cfg)
	if err != nil {
		t.Fatalf("NewUART() error: %v", err)
	}

	badStop := false
	samples := emitFrameSamples(*cfg, 0x33, nil, &badStop)

	_, gotFrame, err := feedSamples(u, samples)
	if err == nil || err.Error() != "framing error" {
		t.Fatalf("expected framing error, got: %v", err)
	}
	if gotFrame {
		t.Fatalf("expected gotFrame=false on framing error")
	}
}

func TestUARTStartBitGlitchRejected(t *testing.T) {
	cfg := &Config{
		SamplesPerBit: 8,
		DataBits:      8,
		StopBits:      1,
		Parity:        ParityNone,
		LsbLast:       false,
		IdleLow:       false,
	}
	u, err := NewUART(cfg)
	if err != nil {
		t.Fatalf("NewUART() error: %v", err)
	}

	// idle sample, then a too-short low glitch that recovers before center sample
	samples := []bool{true, false, true, true, true, true, true, true, true, true}

	for i, sample := range samples {
		_, gotFrame, err := u.WriteSample(sample)
		if err != nil {
			t.Fatalf("sample %d: unexpected error: %v", i, err)
		}
		if gotFrame {
			t.Fatalf("sample %d: unexpected frame decoded from start-bit glitch", i)
		}
	}
}

func TestBool2Int(t *testing.T) {
	if got := bool2int(true); got != 1 {
		t.Fatalf("bool2int(true) = %d, want 1", got)
	}
	if got := bool2int(false); got != 0 {
		t.Fatalf("bool2int(false) = %d, want 0", got)
	}
}

//-----------------------------------------------------------------------------
