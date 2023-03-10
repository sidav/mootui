package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
	"strings"
)

func (ui *uiStruct) showSelectResearchMenu(category int) {
	cw, _ := io.getConsoleSize()
	pFact := currGame.GetPlayerFaction()
	techIds := pFact.GetResearchableTechIdsInCategory(category)
	if len(techIds) == 0 {
		return
	}
	cursorPos := 0
	menuActive := true
	for menuActive {
		line := 0
		io.clearScreen()
		io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		io.drawStringCenteredAround("SELECT NEW "+strings.ToUpper(game.GetTechCategoryName(category))+" RESEARCH",
			cw/2, line)
		line++
		line++
		for i := range techIds {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorBeige)
			} else {
				io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
			}
			io.putString(fmt.Sprintf("%s (Tech %d - %dRP)",
				game.GetTechByCatAndId(category, techIds[i]).Name, techIds[i],
				game.GetScienceCostForTech(category, techIds[i])), 1, line)
			line++
		}
		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			menuActive = false
		case "UP":
			cursorPos--
			if cursorPos < 0 {
				cursorPos = len(techIds) - 1
			}
		case "DOWN":
			cursorPos++
			if cursorPos > len(techIds) - 1 {
				cursorPos = 0
			}
		case "ENTER":
			pFact.CurrentResearchingTech[category] = techIds[cursorPos]
			return // TODO: set tech here
		}
	}
}
