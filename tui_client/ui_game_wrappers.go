package tui_client

import "moocli/game"

func (ui *uiStruct) getStarAtCursor() *game.StarStruct {
	return currGame.Galaxy.GetStarAt(ui.cursorX, ui.cursorY)
}
