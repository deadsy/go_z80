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
	"log"
	"os"

	cli "github.com/deadsy/go-cli"
	"github.com/deadsy/go_z80/z80"
	"github.com/hajimehoshi/ebiten/v2"
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
	bus := newBus()
	cpu := z80.New(io, mem, bus)
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

func mainx() {
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

type Game struct {
	display       *Display
	width, height int
}

func (g *Game) Update() error {
	return g.display.update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)
	g.display.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func main() {

	d := newDisplay()

	d.set(0, 0x5e)
	d.set(1, 0x79)
	d.set(2, 0x77)
	d.set(3, 0x5e)
	d.set(4, 0x7f)
	d.set(5, 0x7f)

	g := &Game{
		display: d,
		width:   d.getWidth(),
		height:  d.getHeight(),
	}

	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle("TEC-1A")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

//-----------------------------------------------------------------------------
