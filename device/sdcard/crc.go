//-----------------------------------------------------------------------------
/*

SD Card Driver

*/
//-----------------------------------------------------------------------------

package sdcard

//-----------------------------------------------------------------------------

func init() {
	generateCRC7Table()
	generateCRC16Table()
}

//-----------------------------------------------------------------------------
// CRC-7, used to check command bytes

var crc7Table [256]byte

func generateCRC7Table() {
	const poly = 0x89 // SD Card CRC7 (commands)
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

// add a byte to the crc
func addCRC7(crc, val byte) byte {
	return crc7Table[(crc<<1)^val]
}

//-----------------------------------------------------------------------------
// CRC-16 CCITT, used to check block contents

var crc16Table [256]uint16

func generateCRC16Table() {
	const polynomial = 0x1021 // SD Card CRC-16-CCITT polynomial
	for i := 0; i < 256; i++ {
		// Shift the index byte into the MSB position of our 16-bit register
		crc := uint16(i << 8)
		// Step through all 8 bits of the current byte
		for bit := 0; bit < 8; bit++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ polynomial
			} else {
				crc <<= 1
			}
		}
		crc16Table[i] = crc
	}
}

func addCRC16(crc uint16, val byte) uint16 {
	idx := byte(crc>>8) ^ val
	return (crc << 8) ^ crc16Table[idx]
}

func CRC16(buf []byte) uint16 {
	crc := uint16(0)
	for _, val := range buf {
		crc = addCRC16(crc, val)
	}
	return crc
}

//-----------------------------------------------------------------------------
