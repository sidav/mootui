package tui_client

import (
	"github.com/gdamore/tcell/v2"
	"moocli/game"
)

func (ui *uiStruct) drawCursor(g *game.Game) {
	io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
	osx, osy := ui.realCoordsToScreenCoords(ui.cursorX, ui.cursorY)
	cursorW := GALAXY_CELL_W+1
	star := g.Galaxy.GetStarAt(ui.cursorX, ui.cursorY)
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

func (ui *uiStruct) drawSidebarForCursorContents() {
	cw, ch := io.getConsoleSize()
	io.drawFilledRect(' ', cw-SIDEBAR_W, 0, SIDEBAR_W, ch-1)
	io.setStyle(tcell.ColorBlue, tcell.ColorBlue)
	io.drawRect(cw-SIDEBAR_W, -1, SIDEBAR_W, ch+1)

	linesx := cw - SIDEBAR_W + 1
	liney := 1
	star := ui.game.Galaxy.GetStarAt(ui.cursorX, ui.cursorY)
	if star == nil {
		io.setStyle(tcell.ColorGray, tcell.ColorBlack)
		io.putString("Deep space", linesx, liney)
		return
	}
	io.setStyle(colorStringToTcell(star.GetStarTypeName()), tcell.ColorBlack)
	io.putString(star.Name, linesx, liney)
	liney++
	io.putString(star.GetStarTypeName() + " star", linesx, liney)
	liney++
	io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
	io.putString(star.GetPlanet().GetPlanetTypeName(), linesx, liney)
	liney++
	if star.GetPlanet().IsColonized() {
		io.setStyle(colorStringToTcell(star.GetPlanet().GetFaction().GetColorName()), tcell.ColorBlack)
		io.putString("Colony", linesx, liney)
	} else {
		io.setStyle(tcell.ColorGray, tcell.ColorBlack)
		io.putString("UNCOLONIZED", linesx, liney)
	}
	liney++
}
