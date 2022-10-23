package tui_client

import "github.com/gdamore/tcell/v2"

func (ui *uiStruct) showSelectResearchMenu(category int) {
	cw, _ := io.getConsoleSize()
	pFact := currGame.GetPlayerFaction()
	teches := pFact.GetResearchableTechesInCategory(category)
	if len(teches) == 0 {
		return
	}
	cursorPos := 0
	menuActive := true
	for menuActive {
		line := 0
		io.clearScreen()
		io.setStyle(tcell.ColorGray, tcell.ColorBlack)
		io.drawStringCenteredAround("SELECT NEW RESEARCH", cw/2, line)
		line++
		line++
		for i := range teches {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorGray)
			} else {
				io.setStyle(tcell.ColorGray, tcell.ColorBlack)
			}
			io.putString(teches[i].Name, 0, line)
			line++
		}
		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			menuActive = false
		case "UP":
			cursorPos--
		case "DOWN":
			cursorPos++
		case "ENTER":
			return // TODO: set tech here
		}
	}
}
