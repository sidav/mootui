//go:build !gui
// +build !gui

package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
	"moocli/graphic_primitives"
	"strconv"
)

func (ui *uiStruct) galaxyScreen() {
	io.clearScreen()
	ui.DrawGalaxy()
	ui.drawSelectCursor()
	ui.drawSidebarForCursorContents()
	io.screen.Show()
}

func (ui *uiStruct) DrawGalaxy() {
	fleets := currGame.Galaxy.GetAllFleets()

	for _, f := range fleets {
		lineFromX, lineFromY := ui.realCoordsToScreenCoords(f.GetCoords())
		lineToX, lineToY := ui.realCoordsToScreenCoords(f.GetTargetCoords())
		line := graphic_primitives.GetLine(lineFromX+2, lineFromY, lineToX+1, lineToY)
		io.setStyle(tcell.ColorDarkGreen, tcell.ColorBlack)
		for _, p := range line {
			io.putChar('*', p.X, p.Y)
		}
	}

	stars := currGame.Galaxy.GetAllStars()
	for _, star := range stars {
		// fmt.Printf("STAR %d: %s at %d, %d\n", i, star.Name, star.X, star.Y)
		ui.drawStar(star)
	}
	for _, fleet := range fleets {
		ui.drawFleet(fleet)
	}
}

func (ui *uiStruct) drawFleet(fleet *game.Fleet) {
	fx, fy := fleet.GetCoords()
	onScreenX, onScreenY := ui.realCoordsToScreenCoords(fx, fy)
	onScreenX += GALAXY_CELL_W-1
	io.setStyle(colorStringToTcell(fleet.GetOwner().GetColorName()), tcell.ColorBlack)
	fleetStr := strconv.Itoa(fleet.GetTotalShipsNumber())
	if fleet.HasColonyShip() {
		fleetStr += "C"
	}
	io.putString(fleetStr, onScreenX, onScreenY)
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
