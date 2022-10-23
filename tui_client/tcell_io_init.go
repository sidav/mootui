package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
)

var io consoleIO

func (io *consoleIO) init() {
	var e error
	io.screen, e = tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = io.screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	io.clearScreen()
}

func (io *consoleIO) close() {
	io.screen.Fini()
}
