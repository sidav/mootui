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
	playerFaction := currGame.GetPlayerFaction()
	currUi = &uiStruct{}
	for _, star := range currGame.Galaxy.GetAllStars() {
		if star.GetPlanet().GetFaction() == playerFaction {
			currUi.cursorX, currUi.cursorY = star.X, star.Y
			break
		}
	}
	io.init()
	defer io.close()

	for !gameShouldExit {
		// show tech setup if needed
		for cat := range playerFaction.CurrentResearchingTech {
			if playerFaction.CurrentResearchingTech[cat] == -1 {
//				currUi.showSelectResearchMenu(cat)
			}
		}
		currUi.centerScreenAroundCursorCoords()
		currUi.DrawGalaxyScreen(currGame)
		key := io.readKey()
		gameShouldExit = key == "EXIT"
		currUi.handleControls(key)
	}
}
