package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
)

func (ui *uiStruct) showSelectResearchMenu(category int) {
	cw, _ := io.getConsoleSize()
	pFact := currGame.GetPlayerFaction()
	teches, ids := pFact.GetResearchableTechesInCategory(category)
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
		// todo: write category
		line++
		line++
		for i := range teches {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorGray)
			} else {
				io.setStyle(tcell.ColorGray, tcell.ColorBlack)
			}
			io.putString(fmt.Sprintf("%s (%dRP)", teches[i].Name, game.GetScienceCostForTech(category, ids[i])), 0, line)
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
			pFact.CurrentResearchingTech[category] = ids[cursorPos]
			return // TODO: set tech here
		}
	}
}
