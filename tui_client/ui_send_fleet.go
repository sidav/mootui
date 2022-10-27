package tui_client

import (
	"github.com/gdamore/tcell/v2"
	"moocli/game"
	"moocli/graphic_primitives"
)

func (ui *uiStruct) sendFleetScreen(sentFleet *game.Fleet) {
	cx, cy := ui.cursorX, ui.cursorY
	for {
		ui.centerScreenAroundCursorCoords()
		io.clearScreen()

		lineFromX, lineFromY := ui.realCoordsToScreenCoords(sentFleet.GetCoords())
		lineToX, lineToY := ui.realCoordsToScreenCoords(ui.cursorX, ui.cursorY)
		line := graphic_primitives.GetLine(lineFromX+1, lineFromY, lineToX+1, lineToY)
		io.setStyle(tcell.ColorDarkGreen, tcell.ColorBlack)
		for _, p := range line {
			io.putChar('*', p.X, p.Y)
		}
		ui.DrawGalaxy()
		ui.drawSendFleetCursor()
		ui.drawSidebarForCursorContents(sentFleet)
		io.screen.Show()
		keyPressed := io.readKey()
		ui.moveCursor(keyPressed)
		if keyPressed == "ESCAPE" {
			break
		}
		if keyPressed == "ENTER" &&
			currGame.Galaxy.GetDistanceToCoordsForEmpire(ui.cursorX, ui.cursorY, currGame.GetPlayerFaction()) <= sentFleet.GetMaxTravelingDistance() {

			currGame.SetFleetDestination(sentFleet, ui.cursorX, ui.cursorY)
			break
		}
	}
	ui.cursorX, ui.cursorY = cx, cy
}
