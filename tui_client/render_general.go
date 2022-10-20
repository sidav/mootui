package tui_client

import (
	"github.com/gdamore/tcell/v2"
	"strings"
)

func colorStringToTcell(color string) tcell.Color {
	switch strings.ToLower(color) {
	case "yellow": return tcell.ColorYellow
	case "red": return tcell.ColorRed
	case "blue": return tcell.ColorBlue
	case "green": return tcell.ColorGreen
	default: return tcell.ColorDarkMagenta // panic("Unknown color " + color)
	}
}
