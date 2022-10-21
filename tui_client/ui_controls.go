package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
)

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

func (ui *uiStruct) colonyMenu(star *game.StarStruct) {
	cw, ch := io.getConsoleSize()
	planet := star.GetPlanet()

	cursorPos := 0
	menuActive := true
	for menuActive {

		io.setStyle(tcell.ColorBlack, tcell.ColorBlack)
		io.drawFilledRect(' ', 0, 0, cw, ch)

		io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
		line := 0
		io.putString("COLONY MENU", 0, line)
		line++
		pop, maxPop := planet.GetPopulation()
		io.putString(fmt.Sprintf("POPULATION %d/%d", pop, maxPop), 0, line)
		line++

		line++
		// slider control menu
		// scmStart := line
		for i := 0; i < game.TOTAL_PLANET_SLIDERS; i++ {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorWhite)
			} else {
				io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
			}
			ui.renderSlider(0, line+i, cw,
				game.GetSliderName(i),
				planet.GetSliderPercent(i), 100,
				fmt.Sprintf("%d%%", planet.GetSliderPercent(i)), "wat")
		}

		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			menuActive = false
		case "UP":
			cursorPos--
		case "DOWN":
			cursorPos++
		case "LEFT":
			planet.ChangeSliderPercent(-1, cursorPos)
		case "RIGHT":
			planet.ChangeSliderPercent(+1, cursorPos)
		}
	}
}
