package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"math"
	"moocli/game"
)

func (ui *uiStruct) renderSlider(x, y, w int, textOnLeft string,
	value, maxValue int, textOnSlider string, fillColor, emptyColor tcell.Color,
	textOnRight string) {

	fullTextOnLeft := fmt.Sprintf("%8s ", textOnLeft)
	if value > 0 {
		fullTextOnLeft += "<"
	} else {
		fullTextOnLeft += " "
	}
	fullTextOnRight := fmt.Sprintf(" %15s", textOnRight)
	if value < maxValue {
		fullTextOnRight = ">" + fullTextOnRight
	} else {
		fullTextOnRight = " " + fullTextOnRight
	}
	sliderW := w - len(fullTextOnLeft) - len(fullTextOnRight)
	if sliderW <= 0 {
		panic("UI failure")
	}
	io.putString(fullTextOnLeft, x, y)
	io.putString(fullTextOnRight, x+len(fullTextOnLeft)+sliderW, y)

	filledCells := int(math.Round(float64(value*sliderW) / float64(maxValue)))
	sliderBeginning := x + len(fullTextOnLeft)
	tosCell := (sliderW - len(textOnSlider)) / 2
	for i := 0; i < sliderW; i++ {
		if i < filledCells {
			io.setStyle(tcell.ColorBlack, fillColor)
		} else {
			io.setStyle(tcell.ColorBlack, emptyColor)
		}
		currTextOnSliderIndex := i - tosCell
		if currTextOnSliderIndex >= 0 && currTextOnSliderIndex < len(textOnSlider) {
			io.putChar(rune(textOnSlider[currTextOnSliderIndex]), i+sliderBeginning, y)
		} else {
			io.putChar(' ', i+sliderBeginning, y)
		}
	}
}

func (ui *uiStruct) drawSelectCursor() {
	io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
	osx, osy := ui.realCoordsToScreenCoords(ui.cursorX, ui.cursorY)
	cursorW := GALAXY_CELL_W + 1
	star := ui.getStarAtCursor()
	if star != nil {
		cursorW = len(star.Name)
		osx = osx - cursorW/2
		cursorW++
	} else {
		osx--
	}
	io.putChar('┏', osx, osy-1)
	io.putChar('┓', osx+cursorW, osy-1)
	io.putChar('┗', osx, osy+2)
	io.putChar('┛', osx+cursorW, osy+2)
}

func (ui *uiStruct) drawSendFleetCursor() {
	io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
	osx, osy := ui.realCoordsToScreenCoords(ui.cursorX, ui.cursorY)
	cursorW := GALAXY_CELL_W + 1
	star := ui.getStarAtCursor()
	if star != nil {
		cursorW = len(star.Name)
		osx = osx - cursorW/2
		cursorW++
		io.setStyle(tcell.ColorGreen, tcell.ColorBlack)
	} else {
		osx--
		io.setStyle(tcell.ColorRed, tcell.ColorBlack)
	}
	io.putChar('\\', osx, osy-1)
	io.putChar('/', osx+cursorW, osy-1)
	io.putChar('/', osx, osy+2)
	io.putChar('\\', osx+cursorW, osy+2)
}

func (ui *uiStruct) drawSidebarForCursorContents(selectedFleet *game.Fleet) {
	cw, ch := io.getConsoleSize()
	io.drawFilledRect(' ', cw-SIDEBAR_W, 0, SIDEBAR_W, ch-1)
	io.setStyle(tcell.ColorBlue, tcell.ColorBlue)
	io.drawRect(cw-SIDEBAR_W, -1, SIDEBAR_W, ch+1)

	linesx := cw - SIDEBAR_W + 1
	liney := 0
	io.resetStyle()
	io.putString(fmt.Sprintf("Turn %d", currGame.Turn), linesx, liney)
	liney++
	star := ui.getStarAtCursor()
	if star == nil {
		io.setStyle(tcell.ColorGray, tcell.ColorBlack)
		io.putString("Deep space", linesx, liney)
		return
	}
	io.setStyle(colorStringToTcell(star.GetStarTypeName()), tcell.ColorBlack)
	io.putString(star.Name, linesx, liney)
	liney++
	io.putString(star.GetStarTypeName()+" star", linesx, liney)
	liney++
	io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
	io.putString(star.GetPlanet().GetPlanetTypeName(), linesx, liney)
	liney++
	io.putString(star.GetPlanet().GetGrowthAndSpecialString(), linesx, liney)
	liney++
	if star.GetPlanet().IsColonized() {
		// FOR COLONIZED PLANETS
		io.setStyle(colorStringToTcell(star.GetPlanet().GetFaction().GetColorName()), tcell.ColorBlack)
		io.putString("Colony", linesx, liney)
		io.resetStyle()
		pop, maxPop := star.GetPlanet().GetPopulationStrings()
		io.putString(fmt.Sprintf("Pop. %s/%s", pop, maxPop), linesx, liney)
		liney++
		io.putString(fmt.Sprintf("Fcts. %d/%d", star.GetPlanet().GetFactories(), currGame.GetMaxFactoriesForPlanet(star.GetPlanet())), linesx, liney)
		liney++
		net, gross := currGame.GetPlanetProductionNetGross(star.GetPlanet())
		io.putString(fmt.Sprintf("Prod. %d (%d)", net, gross), linesx, liney)
		liney++
	} else {
		// FOR NON-COLONIZED PLANETS
		io.setStyle(tcell.ColorGray, tcell.ColorBlack)
		io.putString("UNCOLONIZED", linesx, liney)
		liney++
		_, mp := star.GetPlanet().GetPopulationStrings()
		io.putString(fmt.Sprintf("MAX POP %s", mp), linesx, liney)
	}
	liney++
	if selectedFleet != nil {
		distFromEmpire := currGame.Galaxy.GetDistanceToCoordsForEmpire(ui.cursorX, ui.cursorY, currGame.GetPlayerFaction())
		if selectedFleet.GetMaxTravelingDistance() >= distFromEmpire {
			io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		} else {
			io.setStyle(tcell.ColorRed, tcell.ColorBlack)
		}
		io.putString(fmt.Sprintf("%d parsecs",
			currGame.Galaxy.GetDistanceToCoordsForEmpire(ui.cursorX, ui.cursorY, currGame.GetPlayerFaction())), linesx, liney)
		liney++
		if selectedFleet.GetMaxTravelingDistance() >= distFromEmpire {
			io.putString(fmt.Sprintf("ETA %d turns",
				currGame.GetETAForFleetToCoords(selectedFleet, ui.cursorX, ui.cursorY)), linesx, liney)
		} else {
			io.putString("NOT ENOUGH FUEL", linesx, liney)
		}
		liney++
	}
}
