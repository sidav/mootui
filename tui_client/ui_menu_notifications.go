package tui_client

import (
	"github.com/gdamore/tcell/v2"
	"strings"
)

func (ui *uiStruct) showThisTurnNotifications() {
	nots := currGame.GetPlayerFaction().GetNotifications()
	if len(nots) == 0 {
		return
	}
	menuActive := true
	cw, _ := io.getConsoleSize()
	for menuActive {
		io.clearScreen()
		io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		io.putStringAndIncrementLine("EVENTS:", cw/2)
		io.currentUiLine++
		for i := range nots {
			io.setStyle(tcell.ColorYellow, tcell.ColorBlack)
			io.putStringAndIncrementLine(strings.ToUpper(nots[i].Header) + ":", 0)
			io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
			io.putStringAndIncrementLine(nots[i].Text, 2)
		}
		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE", "ENTER":
			menuActive = false
		}
	}
}
