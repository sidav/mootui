package tui_client

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
		ui.cursorX = ui.game.Galaxy.W - 1
	}
	if ui.cursorY > ui.game.Galaxy.H-1 {
		ui.cursorY = ui.game.Galaxy.H - 1
	}
}
