package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
	"strconv"
)

func (ui *uiStruct) isStarSelectable(s *game.StarStruct) bool {
	return s != nil && s.GetPlanet().IsColonized()
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
		ui.selectOrderForFleet(fleetsHere[0])
		return
	}
	// else, go to "what to select?" menu...
	cursorPos := 0
	cw, _ := io.getConsoleSize()
	menuStrs := make([]string, 0)
	if starHere != nil {
		menuStrs = append(menuStrs, "Colony at "+starHere.Name)
	}
	for i := range fleetsHere {
		menuStrs = append(menuStrs, "Fleet "+strconv.Itoa(i+1))
	}
	for {
		io.clearScreen()
		io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		io.drawStringCenteredAndIncrementLine("WHAT SHOULD BE ORDERED?", cw/2)
		io.currentUiLine++
		for i, s := range menuStrs {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorBeige)
			} else {
				io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
			}
			io.putStringAndIncrementLine(s, 1)
		}
		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			return
		case "UP":
			cursorPos--
			if cursorPos < 0 {
				cursorPos = len(menuStrs) - 1
			}
		case "DOWN":
			cursorPos++
			if cursorPos > len(menuStrs) - 1 {
				cursorPos = 0
			}
		case "ENTER":
			if starHere != nil {
				if cursorPos == 0 {
					ui.colonyMenu(starHere)
					return
				}
				ui.selectOrderForFleet(fleetsHere[cursorPos-1])
				return
			}
			ui.selectOrderForFleet(fleetsHere[cursorPos])
			return
		}
	}
}

func (ui *uiStruct) selectOrderForFleet(f *game.Fleet) {
	cw, _ := io.getConsoleSize()

	menuStrs := []string{
		"Colonize",
		"Move",
		"Split",
	}
	cursorPos := 0
	star := ui.getStarAtCursor()
	for {
		if cursorPos == 0 && !currGame.IsStarColonizableByFleet(star, f) {
			cursorPos++
		}
		io.clearScreen()
		io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		io.drawStringCenteredAndIncrementLine("ORDER FLEET", cw/2)
		ships := f.GetShipsByDesign()
		io.putStringAndIncrementLine(fmt.Sprintf("Current fleet (%d ships total):", f.GetTotalShipsNumber()), 0)
		for designIndex, count := range ships {
			if count > 0 {
				io.putStringAndIncrementLine(
					fmt.Sprintf("%dx %s", count, currGame.GetPlayerFaction().GetDesignByIndex(designIndex).GetName()),
					2)
			}
		}
		io.putStringAndIncrementLine("What is your order?", 0)
		io.currentUiLine++
		for i, s := range menuStrs {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorBeige)
			} else {
				io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
			}
			io.putStringAndIncrementLine(s, 1)
		}

		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			return
		case "UP":
			cursorPos--
			if cursorPos < 0 {
				cursorPos = 2
			}
		case "DOWN":
			cursorPos++
			if cursorPos > 2 {
				cursorPos = 0
			}
		case "ENTER":
			switch cursorPos {
			case 0:
				if currGame.IsStarColonizableByFleet(star, f) {
					currGame.OrderFleetToColonize(f)
				}
				return
			case 1:
				ui.sendFleetScreen(f)
				return
			case 2:
				// todo: implement splitting fleets
			}
		}
	}
}
