//-----------------------------------------------------------------------------
/*

Emulate the TEC-1G RTC Board (DS1302)

*/
//-----------------------------------------------------------------------------

package rtc

import "testing"

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
		rtc, _ := New()
		rtc.clock[clockHour] = 0
		val := rtc.encodeHour(v.n)
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
		rtc, _ := New()
		rtc.clock[clockHour] = (1 << mode12Bit)
		val := rtc.encodeHour(v.n)
		if val != v.val {
			t.Fatalf("case %d: bad value, expected 0x%02x, actual 0x%02x", i, v.val, val)
		}
	}

}

func Test_DecodeHour(t *testing.T) {

	// 24 hour mode
	for i := 0; i <= 23; i++ {
		rtc, _ := New()
		rtc.clock[clockHour] = 0
		val := rtc.encodeHour(i)
		hour := rtc.decodeHour(val)
		if i != hour {
			t.Fatalf("case %d: bad value, expected %d, actual %d", i, i, hour)
		}
	}

	// 12 hour mode
	for i := 0; i <= 23; i++ {
		rtc, _ := New()
		rtc.clock[clockHour] = (1 << mode12Bit)
		val := rtc.encodeHour(i)
		hour := rtc.decodeHour(val)
		if i != hour {
			t.Fatalf("case %d: bad value, expected %d, actual %d", i, i, hour)
		}
	}

}

//-----------------------------------------------------------------------------
