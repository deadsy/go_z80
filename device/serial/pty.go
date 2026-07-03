//-----------------------------------------------------------------------------
/*

Serial Port Pseudo-TTY

*/
//-----------------------------------------------------------------------------

package serial

import (
	"errors"
	"fmt"
	"os"

	"github.com/creack/pty"
)

//-----------------------------------------------------------------------------

// PTY is a pseudo-tty.
type PTY struct {
	master *os.File // our side of the pty (the ptmx)
	tty    *os.File // slave side, e.g. /dev/pts/7
	rxCh   chan byte
}

// NewPTY opens a pseudo-tty
func NewPTY() (*PTY, error) {
	master, tty, err := pty.Open()
	if err != nil {
		return nil, err
	}
	p := &PTY{
		master: master,
		tty:    tty,
		rxCh:   make(chan byte, 256),
	}
	// pump user keystrokes off the pty into a buffered channel so the
	// emulator's sample loop never blocks on a read.
	go p.readLoop()
	return p, nil
}

// Name returns the slave device path, e.g. "/dev/pts/7".
// Print this so the user knows where to point their terminal.
func (p *PTY) Name() string {
	return p.tty.Name()
}

// Close releases the pty file descriptors.
func (p *PTY) Close() error {
	err1 := p.master.Close()
	err2 := p.tty.Close()
	if err1 != nil {
		return err1
	}
	return err2
}

// Write a character to the pty.
func (p *PTY) Write(c byte) {
	buf := [1]byte{c}
	if _, err := p.master.Write(buf[:]); err != nil {
		fmt.Printf("serial pty write error: %s\n", err)
	}
}

// Read returns the next character typed by the user.
// Non-blocking so it fits an emulator tick loop.
func (p *PTY) Read() (byte, error) {
	select {
	case b := <-p.rxCh:
		return b, nil
	default:
		return 0, errors.New("empty")
	}
}

// readLoop moves bytes typed in the terminal into rxCh.
func (p *PTY) readLoop() {
	buf := make([]byte, 64)
	for {
		n, err := p.master.Read(buf)
		if err != nil {
			return // pty closed
		}
		for i := 0; i < n; i++ {
			p.rxCh <- buf[i]
		}
	}
}

//-----------------------------------------------------------------------------
