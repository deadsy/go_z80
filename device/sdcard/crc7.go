//-----------------------------------------------------------------------------
/*

SD Card Driver

*/
//-----------------------------------------------------------------------------

package sdcard

//-----------------------------------------------------------------------------
// CRC-7, used to check command bytes

var crc7Table [256]byte

func generateCRC7Table() {
	const poly = 0x89 // the value of our CRC-7 polynomial
	// generate a table value for all 256 possible byte values
	for i := 0; i < 256; i++ {
		if i&0x80 != 0 {
			crc7Table[i] = byte(i ^ poly)
		} else {
			crc7Table[i] = byte(i)
		}
		for j := 1; j < 8; j++ {
			crc7Table[i] <<= 1
			if crc7Table[i]&0x80 != 0 {
				crc7Table[i] ^= poly
			}
		}
	}
}

func init() {
	generateCRC7Table()
}

// add a byte to the crc
func crcAdd(crc, val byte) byte {
	return crc7Table[(crc<<1)^val]
}

//-----------------------------------------------------------------------------
