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

		io.setStyle(tcell.ColorBlack, tcell.ColorBlack)
		io.drawFilledRect(' ', 0, 0, cw, ch)

		line := 0
		io.setStyle(colorStringToTcell(star.GetStarTypeName()), tcell.ColorBlack)
		io.putString(star.Name + " - " + star.GetStarTypeName() + " star", 0, line)
		line++
		io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
		io.putString(fmt.Sprintf("Colony on %s planet:", planet.GetPlanetTypeName()), 0, line)
		line++
		pop, maxPop := planet.GetPopulation()
		io.putString(fmt.Sprintf("Pop. %d/%d billion", pop, maxPop), 0, line)
		line++
		io.putString(fmt.Sprintf("Fcts. %d/%d", star.GetPlanet().GetFactories(), ui.game.GetMaxFactoriesForPlanet(star.GetPlanet())), 0, line)
		line++
		net, gross := ui.game.GetPlanetProductionNetGross(star.GetPlanet())
		io.putString(fmt.Sprintf("Prod. %d (%d)", net, gross), 0, line)
		line++

		line++
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
			ui.renderSlider(0, line, cw,
				game.GetSliderName(i),
				planet.GetSliderPercent(i), 100,
				fmt.Sprintf("%d%%", planet.GetSliderPercent(i)), sliderFillColor, tcell.ColorGray,
				"not implemented")
			line++
		}
		io.resetStyle()
		io.putString("Press SPACE to (un)lock slider", 0, line)
		line++
		io.putString("Press E to equalize sliders", 0, line)
		line++

		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			menuActive = false
		case "UP":
			cursorPos--
		case "DOWN":
			cursorPos++
		case "LEFT":
			planet.ChangeSliderPercent(-1, cursorPos)
		case "RIGHT":
			planet.ChangeSliderPercent(+1, cursorPos)
		case "e":
			planet.EqualizeSliders(false)
		case " ":
			planet.FlipSliderLock(cursorPos)
		}
	}
}

