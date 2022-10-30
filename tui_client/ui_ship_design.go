package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"moocli/game"
	"strconv"
)

func (ui *uiStruct) SelectDesignToChange() {
	cw, _ := io.getConsoleSize()
	cursorPos := 0
	for {
		line := 0
		io.clearScreen()
		io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		io.drawStringCenteredAround("SELECT WHAT DESIGN TO CHANGE", cw/2, line)
		line++
		line++
		for i := 0; i < game.SHIP_DESIGNS_PER_FACTION; i++ {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorBeige)
			} else {
				io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
			}
			des := currGame.GetPlayerFaction().GetDesignByIndex(i)
			if des == nil {
				io.putString("[NEW DESIGN]", 1, line)
			} else {
				io.putString(des.Name, 1, line)
			}
			line++
		}
		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			return
		case "UP":
			cursorPos--
		case "DOWN":
			cursorPos++
		case "ENTER":
			ui.ChangeDesignNumber(cursorPos)
			return
		}
	}
}

func (ui *uiStruct) ChangeDesignNumber(num int) {
	initialDesign := currGame.GetPlayerFaction().GetDesignByIndex(num)
	var des game.ShipDesign
	if initialDesign != nil {
		des = *(initialDesign)
	}

	menuStrings := ui.GetStringsArrayForShipDesign(&des)
	cw, _ := io.getConsoleSize()
	cursorPos := 0
	for {
		line := 0
		io.clearScreen()
		io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		io.drawStringCenteredAround("Changing design", cw/2, line)
		line++
		line++
		for i := range menuStrings {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorBeige)
			} else {
				io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
			}
			io.putString(menuStrings[i], 1, line)
			line++
		}
		io.screen.Show()
		weapSlot := cursorPos-1-int(game.SDSLOT_COUNT)
		specSysSlot := cursorPos-5-int(game.SDSLOT_COUNT)
		key := io.readKey()
		switch key {
		case "ESCAPE":
			return
		case "UP":
			cursorPos--
		case "DOWN":
			cursorPos++
		case "LEFT":
			if cursorPos == 0 {
				des.Size = (des.Size -1)
				if des.Size < 0 {
					des.Size = game.SSIZES_COUNT-1
				}
			} else if weapSlot >= 0 && weapSlot < 4 {
				if des.Weapons[weapSlot] != nil {
					des.Weapons[weapSlot].Count--
					if des.Weapons[weapSlot].Count <= 0 {
						des.Weapons[weapSlot] = nil
					}
				}
			}
			menuStrings = ui.GetStringsArrayForShipDesign(&des)
		case "RIGHT":
			if cursorPos == 0 {
				des.Size = (des.Size + 1) % game.SSIZES_COUNT
			} else if weapSlot >= 0 && weapSlot < 4 {
				if des.Weapons[weapSlot] != nil {
					des.Weapons[weapSlot].Count++
				}
			}
			menuStrings = ui.GetStringsArrayForShipDesign(&des)
		case "ENTER":
			if cursorPos > 0 && cursorPos-1 < int(game.SDSLOT_COUNT) {
				des.Systems[cursorPos-1] = ui.selectShipSystemForSlot(cursorPos - 1)
			} else if weapSlot >= 0 && weapSlot < 4 {
				if des.Weapons[weapSlot] == nil {
					des.Weapons[weapSlot] = &game.WeaponInstallation{}
				}
				des.Weapons[weapSlot].Weapon = ui.selectShipSystemForSlot(int(game.SDSLOT_WEAPON))
				des.Weapons[weapSlot].Count = 1
			} else if specSysSlot >= 0 && specSysSlot < 4 {
				des.SpecialSystems[specSysSlot] = ui.selectShipSystemForSlot(int(game.SDSLOT_SPECIAL))
			}
			menuStrings = ui.GetStringsArrayForShipDesign(&des)
		}
	}
}

func (ui *uiStruct) GetStringsArrayForShipDesign(des *game.ShipDesign) (strs []string) {
	strs = append(strs, "Size:       "+game.GetShipSizeName(des.Size))
	strs = append(strs, "Armor:      "+ui.getShipSystemString(des.Systems[0]))
	strs = append(strs, "Shield:     "+ui.getShipSystemString(des.Systems[1]))
	strs = append(strs, "Engine:     "+ui.getShipSystemString(des.Systems[2]))
	strs = append(strs, "Fuel:       "+ui.getShipSystemString(des.Systems[3]))
	strs = append(strs, "Computer:   "+ui.getShipSystemString(des.Systems[4]))
	for i := range des.Weapons {
		strs = append(strs, "Weapon "+strconv.Itoa(i+1)+":   "+ui.getShipWeaponInstallationString(des.Weapons[i]))
	}
	for i := range des.SpecialSystems {
		strs = append(strs, "Special "+strconv.Itoa(i+1)+":  "+ui.getShipSystemString(des.SpecialSystems[i]))
	}

	return strs
}

func (ui *uiStruct) selectShipSystemForSlot(cat int) *game.ShipSystemStruct {
	availSystems := currGame.GetPlayerFaction().GetListOfAvailableShipSystemsInCategory(cat)
	if len(availSystems) == 0 {
		return nil
	}
	if len(availSystems) == 1 {
		return availSystems[0]
	}
	cw, _ := io.getConsoleSize()
	cursorPos := 0
	for {
		line := 0
		io.clearScreen()
		io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
		io.drawStringCenteredAround("Changing design", cw/2, line)
		line++
		line++
		for i := range availSystems {
			if i == cursorPos {
				io.setStyle(tcell.ColorBlack, tcell.ColorBeige)
			} else {
				io.setStyle(tcell.ColorBeige, tcell.ColorBlack)
			}
			io.putString(availSystems[i].GetName(), 1, line)
			line++
		}
		io.screen.Show()
		key := io.readKey()
		switch key {
		case "ESCAPE":
			return nil
		case "UP":
			cursorPos--
		case "DOWN":
			cursorPos++
		case "ENTER":
			if cursorPos == len(availSystems) {
				return nil
			}
			return availSystems[cursorPos]
		}
	}
}

func (ui *uiStruct) getShipWeaponInstallationString(inst *game.WeaponInstallation) string {
	if inst == nil || inst.Weapon == nil {
		return "-- no weapon --"
	}
	return fmt.Sprintf("%dx%s", inst.Count, inst.Weapon.GetName())
}

func (ui *uiStruct) getShipSystemString(s *game.ShipSystemStruct) string {
	if s == nil {
		return "-- no system --"
	}
	return s.GetName()
}
