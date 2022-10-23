package tui_client

func (ui *uiStruct) handleControls(keyPressed string) {
	ui.moveCursor(keyPressed)
	star := ui.getStarAtCursor()
	if keyPressed == "ENTER" && star != nil && star.GetPlanet().IsColonized() {
		ui.colonyMenu(star)
	}
	if keyPressed == "n" {
		currGame.ProcessTurn()
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
