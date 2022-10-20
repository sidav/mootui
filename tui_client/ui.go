package tui_client

import "moocli/game"

const (
	GALAXY_CELL_W = 3
	GALAXY_CELL_H = 2
	SIDEBAR_W     = 17
)

type uiStruct struct {
	game             *game.Game
	cursorX, cursorY int // REAL coordinates
	camtlX, camtlY   int // camera REAL top left coord
}

func (ui *uiStruct) centerScreenAroundCursorCoords() {
	sw, sh := io.getConsoleSize()
	sw -= SIDEBAR_W
	ui.camtlX = ui.cursorX - (sw/GALAXY_CELL_W)/2
	ui.camtlY = ui.cursorY - (sh/GALAXY_CELL_H)/2
}

func (ui *uiStruct) realCoordsToScreenCoords(rx, ry int) (int, int) {
	return GALAXY_CELL_W * (rx - ui.camtlX), GALAXY_CELL_H * (ry - ui.camtlY)
}

func (ui *uiStruct) moveCursor(keyPressed string) {
	switch keyPressed {
	case "UP":
		ui.cursorY--
	case "DOWN":
		ui.cursorY++
	case "LEFT":
		ui.cursorX--
	case "RIGHT":
		ui.cursorX++
	default:
		io.debugPrint(keyPressed)
	}
	if ui.cursorX < 0 {
		ui.cursorX = 0
	}
	if ui.cursorY < 0 {
		ui.cursorY = 0
	}
	if ui.cursorX > ui.game.Galaxy.W-1 {
		ui.cursorX = ui.game.Galaxy.W-1
	}
	if ui.cursorY > ui.game.Galaxy.H-1 {
		ui.cursorY = ui.game.Galaxy.H-1
	}
}
