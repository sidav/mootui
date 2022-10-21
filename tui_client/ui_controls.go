package tui_client

func (ui *uiStruct) handleControls(keyPressed string) {
	ui.moveCursor(keyPressed)
	star := ui.getStarAtCursor()
	if keyPressed == "ENTER" && star != nil && star.GetPlanet().IsColonized() {
		ui.colonyMenu(star)
	}
	if keyPressed == "n" {
		ui.game.ProcessTurn()
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
	if ui.cursorX > ui.game.Galaxy.W-1 {
		ui.cursorX = ui.game.Galaxy.W - 1
	}
	if ui.cursorY > ui.game.Galaxy.H-1 {
		ui.cursorY = ui.game.Galaxy.H - 1
	}
}
