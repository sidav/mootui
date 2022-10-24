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
		line := 0
		io.clearScreen()
		io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		io.drawStringCenteredAround("EVENTS:", cw/2, line)
		line++
		line++
		for i := range nots {
			io.setStyle(tcell.ColorYellow, tcell.ColorBlack)
			io.putString(strings.ToUpper(nots[i].Header) + ":", 0, line)
			line++
			io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
			io.putString(nots[i].Text, 2, line)
			line++
		}
		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE", "ENTER":
			menuActive = false
		}
	}
}
