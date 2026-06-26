//-----------------------------------------------------------------------------
/*

Jupiter ACE Emulator

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/deadsy/go-cli"
	"github.com/deadsy/go_z80/z80"
)

//-----------------------------------------------------------------------------

const historyPath = "history.txt"

//-----------------------------------------------------------------------------

// userApp is state associated with the user application.
type userApp struct {
	io  *sysIO
	mem *sysMemory
	cpu *z80.CPU
}

// newUserApp returns a user application.
func newUserApp() (*userApp, error) {
	io := newIO()
	mem, err := newMemory()
	if err != nil {
		return nil, err
	}
	bus := newBus()
	cpu := z80.New(io, mem, bus)
	return &userApp{
		io:  io,
		mem: mem,
		cpu: cpu,
	}, nil
}

//-----------------------------------------------------------------------------

// Put outputs a string to the user application.
func (u *userApp) Put(s string) {
	fmt.Printf("%s", s)
}

//-----------------------------------------------------------------------------

func main() {

	// create the application
	app, err := newUserApp()
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	// create the cli
	c := cli.NewCLI(app)
	c.HistoryLoad(historyPath)
	c.SetRoot(menuRoot)
	c.SetPrompt("jace> ")

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
