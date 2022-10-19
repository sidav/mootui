package tui_client

import (
	"moocli/game"
)

var (
	currGame *game.Game
	currUi   *uiStruct
)

func StartGame() {
	currUi = &uiStruct{}
	currGame = game.InitNewGame()
	io.init()
	defer io.close()

	currUi.DrawGalaxy(currGame)

	io.readKey()
}
