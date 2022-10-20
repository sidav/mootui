//go:build !gui
// +build !gui

package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
)

func (ui *uiStruct) DrawGalaxy(g *game.Game) {
	io.screen.Clear()
	stars := g.Galaxy.GetAllStars()
	ui.drawCursor(g)
	for _, star := range stars {
		// fmt.Printf("STAR %d: %s at %d, %d\n", i, star.Name, star.X, star.Y)
		ui.drawStar(star)
	}
	io.screen.Show()
}

func (ui *uiStruct) drawCursor(g *game.Game) {
	io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
	osx, osy := ui.realCoordsToScreenCoords(ui.cursorX, ui.cursorY)
	cursorW := GALAXY_CELL_W
	star := g.Galaxy.GetStarAt(ui.cursorX, ui.cursorY)
	if star != nil {
		cursorW = len(star.Name)
		osx -= (cursorW - GALAXY_CELL_W) / 2
		if (cursorW-GALAXY_CELL_W)%2 == 0 {
			cursorW++
			osx--
		} else {
			osx -= 2
		}
	}
	io.putChar('┏', osx, osy-1)
	io.putChar('┓', osx+cursorW, osy-1)
	io.putChar('┗', osx, osy+2)
	io.putChar('┛', osx+cursorW, osy+2)
}

func (ui *uiStruct) drawStar(star *game.StarStruct) {
	onScreenX, onScreenY := ui.realCoordsToScreenCoords(star.X, star.Y)

	var starRune rune
	var starColor tcell.Color
	switch star.GetStarTypeName() {
	case "Yellow":
		starRune = '*'
		starColor = tcell.ColorYellow
	case "White":
		starRune = '*'
		starColor = tcell.ColorWhite
	case "Blue":
		starRune = '*'
		starColor = tcell.ColorBlue
	case "Red":
		starRune = 'o'
		starColor = tcell.ColorRed
	case "Green":
		starRune = 'o'
		starColor = tcell.ColorGreen
	case "Neutron":
		starRune = '*'
		starColor = tcell.ColorDarkRed
	default:
		panic(fmt.Sprintf("%s is unknown star type", star.GetStarTypeName()))
	}

	io.setStyle(starColor, tcell.ColorBlack)
	io.putChar(starRune, onScreenX+1, onScreenY)
	if star.GetColony() == nil {
		io.setStyle(tcell.ColorGray, tcell.ColorBlack)
	} else {
		io.setStyle(tcell.ColorBlack, strColorToTcell(star.GetColony().GetFaction().GetColorName()))
	}
	io.drawStringCenteredAround(star.Name, onScreenX+1, onScreenY+1)
}
