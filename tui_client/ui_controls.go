package tui_client

func (ui *uiStruct) handleGalaxyScreenControls(keyPressed string) {
	ui.moveCursor(keyPressed)
	if keyPressed == "ENTER" {
		ui.selectEntityFromGalaxyScreen()
	}
	if keyPressed == "n" {
		currGame.ProcessTurn()
		ui.turnEnded = true
	}
	if keyPressed == "f" {
		ui.SelectDesignToChange()
		ui.turnEnded = true
	}
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
	if ui.cursorX > currGame.Galaxy.W-1 {
		ui.cursorX = currGame.Galaxy.W - 1
	}
	if ui.cursorY > currGame.Galaxy.H-1 {
		ui.cursorY = currGame.Galaxy.H - 1
	}
}
