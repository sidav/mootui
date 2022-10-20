package tui_client

import (
	"moocli/game"
)

var (
	currGame       *game.Game
	currUi         *uiStruct
	gameShouldExit bool
)

func StartGame() {
	currUi = &uiStruct{}
	currGame = game.InitNewGame()
	io.init()
	defer io.close()

	for !gameShouldExit {
		currUi.centerScreenAroundCursorCoords()
		currUi.DrawGalaxy(currGame)
		key := io.readKey()
		gameShouldExit = key == "EXIT"
		currUi.moveCursor(key)
	}
}
