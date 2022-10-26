//go:build !gui
// +build !gui

package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
	"strconv"
)

func (ui *uiStruct) DrawGalaxyScreen(g *game.Game) {
	currGame = g
	io.clearScreen()
	stars := g.Galaxy.GetAllStars()
	for _, star := range stars {
		// fmt.Printf("STAR %d: %s at %d, %d\n", i, star.Name, star.X, star.Y)
		ui.drawStar(star)
	}
	fleets := g.Galaxy.GetAllFleets()
	for _, fleet := range fleets {
		ui.drawFleet(fleet)
	}
	ui.drawCursor()
	ui.drawSidebarForCursorContents()
	io.screen.Show()
}

func (ui *uiStruct) drawFleet(fleet *game.Fleet) {
	fx, fy := fleet.GetCoords()
	onScreenX, onScreenY := ui.realCoordsToScreenCoords(fx, fy)
	onScreenX += GALAXY_CELL_W-1
	io.setStyle(colorStringToTcell(fleet.GetOwner().GetColorName()), tcell.ColorBlack)
	io.putString(strconv.Itoa(fleet.GetShipsNumber()), onScreenX, onScreenY)
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
	if star.GetPlanet().IsColonized() {
		io.setStyle(tcell.ColorBlack, colorStringToTcell(star.GetPlanet().GetFaction().GetColorName()))
	} else {
		io.setStyle(tcell.ColorGray, tcell.ColorBlack)
	}
	io.drawStringCenteredAround(star.Name, onScreenX+1, onScreenY+1)
}
