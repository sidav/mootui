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
	currGame = game.InitNewGame()
	currUi = &uiStruct{}
	for _, star := range currGame.Galaxy.GetAllStars() {
		if star.GetPlanet().IsColonized() {
			currUi.cursorX, currUi.cursorY = star.X, star.Y
			break
		}
	}
	io.init()
	defer io.close()

	for !gameShouldExit {
		currUi.centerScreenAroundCursorCoords()
		currUi.DrawGalaxyScreen(currGame)
		key := io.readKey()
		gameShouldExit = key == "EXIT"
		currUi.handleControls(key)
	}
}
