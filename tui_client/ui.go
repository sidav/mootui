package tui_client

const (
	GALAXY_CELL_W = 3
	GALAXY_CELL_H = 2
	SIDEBAR_W     = 17
)

type uiStruct struct {
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
