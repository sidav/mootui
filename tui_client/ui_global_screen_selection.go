package tui_client

import (
	"github.com/gdamore/tcell/v2"
	"moocli/game"
	"strconv"
)

func (ui *uiStruct) isStarSelectable(s *game.StarStruct) bool {
	return s.GetPlanet().IsColonized()
}

func (ui *uiStruct) selectEntityFromGalaxyScreen() {
	starHere := currGame.Galaxy.GetStarAt(ui.cursorX, ui.cursorY)
	fleetsHere := currGame.Galaxy.GetFleetsAt(ui.cursorX, ui.cursorY)
	// if the star is colonized, AND there are no fleets, select the colony right away
	if !ui.isStarSelectable(starHere) && fleetsHere == nil {
		return
	}
	if ui.isStarSelectable(starHere) && len(fleetsHere) == 0 {
		ui.colonyMenu(starHere)
		return
	}
	// if the star is NOT colonized, AND there is only one fleet, select the fleet right away
	if !ui.isStarSelectable(starHere) && len(fleetsHere) == 1 {
		ui.sendFleetScreen(fleetsHere[0])
		return
	}
	// else, go to "what to select?" menu...
	cursorPos := 0
	menuActive := true
	cw, _ := io.getConsoleSize()
	menuStrs := make([]string, 0)
	if starHere != nil {
		menuStrs = append(menuStrs, "Colony at " + starHere.Name)
	}
	for i := range fleetsHere {
		menuStrs = append(menuStrs, "Fleet " + strconv.Itoa(i+1))
	}
	for menuActive {
		for menuActive {
			line := 0
			io.clearScreen()
			io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
			io.drawStringCenteredAround("WHAT SHOULD BE ORDERED?",
				cw/2, line)
			line++
			line++
			for i, s := range menuStrs {
				if i == cursorPos {
					io.setStyle(tcell.ColorBlack, tcell.ColorBeige)
				} else {
					io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
				}
				io.putString(s,1, line)
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
				if starHere != nil {
					if cursorPos == 0 {
						ui.colonyMenu(starHere)
						return
					}
					ui.sendFleetScreen(fleetsHere[cursorPos-1])
					return
				}
				ui.sendFleetScreen(fleetsHere[cursorPos])
				return
			}
		}
	}
}
