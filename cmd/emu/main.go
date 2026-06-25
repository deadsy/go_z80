//-----------------------------------------------------------------------------
/*

Z80 Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	cli "github.com/deadsy/go-cli"
	"github.com/deadsy/go_z80/z80"
)

//-----------------------------------------------------------------------------

const historyPath = "history.txt"

//-----------------------------------------------------------------------------

// userApp is state associated with the user application.
type userApp struct {
	io  *IO
	mem *Memory
	cpu *z80.CPU
}

// newUserApp returns a user application.
func newUserApp() *userApp {
	io := newIO()
	mem := newMemory()
	cpu := z80.New(io, mem)
	return &userApp{
		io:  io,
		mem: mem,
		cpu: cpu,
	}
}

//-----------------------------------------------------------------------------
// file loading

// loadRaw loads a raw binary file.
func (u *userApp) loadRaw(filename string, x []uint8) (string, error) {

	// copy the code to the load address
	var loadAdr uint16
	for i, v := range x {
		u.mem.Write8(loadAdr+uint16(i), v)
	}
	endAdr := loadAdr + uint16(len(x)) - 1

	return fmt.Sprintf("%s code %04x-%04x", filename, loadAdr, endAdr), nil
}

func (u *userApp) loadFile(filename string) (string, error) {
	// get the file contents
	x, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return u.loadRaw(filename, x)
}

//-----------------------------------------------------------------------------

// Put outputs a string to the user application.
func (u *userApp) Put(s string) {
	fmt.Printf("%s", s)
}

//-----------------------------------------------------------------------------

func main() {
	// command line flags
	fname := flag.String("f", "out.bin", "file to load (sim6502 or raw)")
	flag.Parse()

	// create the application
	app := newUserApp()

	// load the file
	status, err := app.loadFile(*fname)
	app.mem.WriteROM(false)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", status)
	}

	// create the cli
	c := cli.NewCLI(app)
	c.HistoryLoad(historyPath)
	c.SetRoot(menuRoot)
	c.SetPrompt("emu> ")

	// reset the cpu
	app.cpu.Reset()

	// run the cli
	for c.Running() {
		c.Run()
	}

	// exit
	c.HistorySave(historyPath)
	os.Exit(0)
}

//-----------------------------------------------------------------------------
