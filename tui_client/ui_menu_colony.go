package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
)

func (ui *uiStruct) colonyMenu(star *game.StarStruct) {
	cw, ch := io.getConsoleSize()
	planet := star.GetPlanet()

	cursorPos := 0
	menuActive := true
	for menuActive {
		io.currentUiLine = 0
		io.setStyle(tcell.ColorBlack, tcell.ColorBlack)
		io.drawFilledRect(' ', 0, 0, cw, ch)

		io.setStyle(colorStringToTcell(star.GetStarTypeName()), tcell.ColorBlack)
		io.putStringAndIncrementLine(star.Name+" - "+star.GetStarTypeName()+" star", 0)
		io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
		//io.putString(fmt.Sprintf("Colony on %s planet:", planet.GetPlanetTypeName()), 0)
		//line++
		io.putStringAndIncrementLine(fmt.Sprintf("Colony on %s planet:", planet.GetGrowthAndSpecialString()), 0)
		pop, maxPop := planet.GetPopulationStrings()
		io.putStringAndIncrementLine(fmt.Sprintf("Pop. %s/%s bln.", pop, maxPop), 0)
		io.putStringAndIncrementLine(fmt.Sprintf("Fcts. %d/%d Waste +%d/-%d", planet.GetFactories(),
			currGame.GetMaxFactoriesForPlanet(planet),
			currGame.GetPlanetWaste(planet), currGame.GetPlanetWasteRemoval(planet, true)),
			0)
		net, gross := currGame.GetPlanetProductionNetGross(star.GetPlanet())
		io.putStringAndIncrementLine(fmt.Sprintf("Prod. %d (%d)", net, gross), 0)

		io.putStringAndIncrementLine(fmt.Sprintf("Built ship: %s",
			currGame.GetPlayerFaction().GetDesignByIndex(planet.CurrentBuiltShipDesignIndex).GetName()),
			0)
		io.putStringAndIncrementLine("  (Press s to change)", 0)
		io.currentUiLine++

		// slider control menu
		// scmStart := line
		for i := 0; i < game.TOTAL_PLANET_SLIDERS; i++ {
			sliderTextColor := tcell.ColorWhite
			sliderFillColor := tcell.ColorDarkGreen
			if planet.GetSliderLock(i) {
				sliderTextColor = tcell.ColorDarkRed
				sliderFillColor = tcell.ColorDarkRed
			}
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, sliderTextColor)
			} else {
				io.setStyle(sliderTextColor, tcell.ColorBlack)
			}
			ui.renderSlider(0, io.currentUiLine, cw,
				game.GetSliderName(i),
				planet.GetSliderPercent(i), 100,
				fmt.Sprintf("%d%% (%dBC)", planet.GetSliderPercent(i), currGame.GetPlanetBCForSlider(planet, i)),
				sliderFillColor, tcell.ColorGray,
				currGame.GetSliderString(planet, i))
			io.currentUiLine++
		}
		io.resetStyle()
		io.currentUiLine++
		io.putStringAndIncrementLine("Press SPACE to (un)lock slider", 0)
		io.putStringAndIncrementLine("Press E to equalize sliders", 0)

		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			menuActive = false
		case "UP":
			cursorPos--
			if cursorPos < 0 {
				cursorPos = 4
			}
		case "DOWN":
			cursorPos++
			if cursorPos > 4 {
				cursorPos = 0
			}
		case "LEFT":
			planet.ChangeSliderPercent(-1, cursorPos)
		case "RIGHT":
			planet.ChangeSliderPercent(+1, cursorPos)
		case "s":
			planet.CurrentBuiltShipDesignIndex = (planet.CurrentBuiltShipDesignIndex + 1) % game.SHIP_DESIGNS_PER_FACTION
			for currGame.GetPlayerFaction().GetDesignByIndex(planet.CurrentBuiltShipDesignIndex) == nil {
				planet.CurrentBuiltShipDesignIndex = (planet.CurrentBuiltShipDesignIndex + 1) % game.SHIP_DESIGNS_PER_FACTION
			}
		case "e":
			planet.EqualizeSliders(false)
		case " ":
			planet.FlipSliderLock(cursorPos)
		}
	}
}
