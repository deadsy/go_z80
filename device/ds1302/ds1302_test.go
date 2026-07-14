//-----------------------------------------------------------------------------
/*

DS1302 Real Time Clock Emulation

*/
//-----------------------------------------------------------------------------

package ds1302

import (
	"math/rand/v2"
	"testing"
	"time"
)

//-----------------------------------------------------------------------------

func Test_EncodeHour(t *testing.T) {

	testMode24 := []struct {
		n   int  // 0..23
		val byte // clockHour register value
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 9},
		{10, (1 << 4) | 0},
		{11, (1 << 4) | 1},
		{12, (1 << 4) | 2},
		{13, (1 << 4) | 3},
		{14, (1 << 4) | 4},
		{15, (1 << 4) | 5},
		{16, (1 << 4) | 6},
		{17, (1 << 4) | 7},
		{18, (1 << 4) | 8},
		{19, (1 << 4) | 9},
		{20, (2 << 4) | 0},
		{21, (2 << 4) | 1},
		{22, (2 << 4) | 2},
		{23, (2 << 4) | 3},
	}

	for i, v := range testMode24 {
		val := encodeHour(v.n, false)
		if val != v.val {
			t.Fatalf("case %d: bad value, expected 0x%02x, actual 0x%02x", i, v.val, val)
		}
	}

	testMode12 := []struct {
		n   int  // 0..23
		val byte // clockHour register value
	}{
		{0, 0x80 | 0},
		{1, 0x80 | 1},
		{2, 0x80 | 2},
		{3, 0x80 | 3},
		{4, 0x80 | 4},
		{5, 0x80 | 5},
		{6, 0x80 | 6},
		{7, 0x80 | 7},
		{8, 0x80 | 8},
		{9, 0x80 | 9},
		{10, 0x80 | (1 << 4) | 0},
		{11, 0x80 | (1 << 4) | 1},
		{12, 0x80 | (1 << 5) | (1 << 4) | 2},
		{13, 0x80 | (1 << 5) | (0 << 4) | 1},
		{14, 0x80 | (1 << 5) | (0 << 4) | 2},
		{15, 0x80 | (1 << 5) | (0 << 4) | 3},
		{16, 0x80 | (1 << 5) | (0 << 4) | 4},
		{17, 0x80 | (1 << 5) | (0 << 4) | 5},
		{18, 0x80 | (1 << 5) | (0 << 4) | 6},
		{19, 0x80 | (1 << 5) | (0 << 4) | 7},
		{20, 0x80 | (1 << 5) | (0 << 4) | 8},
		{21, 0x80 | (1 << 5) | (0 << 4) | 9},
		{22, 0x80 | (1 << 5) | (1 << 4) | 0},
		{23, 0x80 | (1 << 5) | (1 << 4) | 1},
	}

	for i, v := range testMode12 {
		val := encodeHour(v.n, true)
		if val != v.val {
			t.Fatalf("case %d: bad value, expected 0x%02x, actual 0x%02x", i, v.val, val)
		}
	}
}

//-----------------------------------------------------------------------------

func encodeDecode(t *testing.T, mode12 bool) {
	for i := 0; i <= 23; i++ {
		val := encodeHour(i, mode12)
		hour := decodeHour(val)
		if i != hour {
			t.Fatalf("case %d: bad value, expected %d, actual %d", i, i, hour)
		}
	}
}

func Test_DecodeHour(t *testing.T) {
	// 24 hour mode
	encodeDecode(t, false)
	// 12 hour mode
	encodeDecode(t, true)
}

//-----------------------------------------------------------------------------

func randInt(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func Test_SetAndGet(t *testing.T) {
	for i := 0; i < 2000; i++ {
		year := randInt(0, 99)
		month := time.Month(randInt(1, 12))
		day := randInt(1, 31)
		hour := randInt(0, 23)
		min := randInt(0, 59)
		sec := randInt(0, 59)

		t_in := time.Date(year, month, day, hour, min, sec, 0, time.UTC)
		rtcTime := newRtcTime(t_in, 2000, 0)
		t_out := rtcTime.getTime(2000)

		if !t_in.Equal(t_out) {
			t.Fatalf("case %d: in %s out %s", i, t_in, t_out)
		}
	}
}

//-----------------------------------------------------------------------------
