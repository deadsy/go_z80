//-----------------------------------------------------------------------------
/*

Intel Hex File Handling

*/
//-----------------------------------------------------------------------------

package util

import (
	"errors"
	"os"

	"github.com/deadsy/go_z80/z80"
	"github.com/unixdj/ihex"
)

//-----------------------------------------------------------------------------

func LoadIntelHex(mem z80.Memory, path string) error {
	// open the file
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// parse the hex records
	ix := ihex.IHex{}
	err = ix.ReadFrom(f)
	if err != nil {
		return err
	}
	// write the chunks to memory
	for i, _ := range ix.Chunks {
		adr := ix.Chunks[i].Addr
		buf := ix.Chunks[i].Data
		//log.Printf("chunk %d adr 0x%04x %d bytes", i, adr, len(buf))
		if int(adr) > 0xffff || int(adr)+len(buf) > 0xffff {
			return errors.New("address out of range")
		}
		for ofs, val := range buf {
			mem.Write8(uint16(int(adr)+ofs), val)
		}
	}
	return nil
}

//-----------------------------------------------------------------------------
