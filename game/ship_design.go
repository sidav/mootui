package game

import (
	"fmt"
)

const (
	SSIZE_SMALL = iota
	SSIZE_MEDIUM
	SSIZE_LARGE
	SSIZE_GIANT
	SSIZES_COUNT
)

type ShipDesign struct {
	name           string
	Mark           int // mk1, mk2 etc
	Size           int
	Weapons        [4]*WeaponInstallation
	Systems        [SDSLOT_COUNT]*ShipSystemStruct
	SpecialSystems [4]*ShipSystemStruct // holds indices
}

type WeaponInstallation struct {
	Weapon *ShipSystemStruct
	Count  int
}

func (sd *ShipDesign) increaseMark() {
	sd.Mark++
}

func (sd *ShipDesign) SetName(n string) {
	sd.name = n
}

func (sd *ShipDesign) GetName() string {
	markString := ""
	switch sd.Mark {
	case 0, 1:
		break
	case 2:
		markString = " mk.II"
	case 3:
		markString = " mk.III"
	case 4:
		markString = " mk.IV"
	case 5:
		markString = " mk.V"
	case 6:
		markString = " mk.VI"
	default:
		markString = fmt.Sprintf(" mk.%d", sd.Mark)
	}
	return sd.name + markString
}

func (sd *ShipDesign) GetTotalSpace() int {
	switch sd.Size {
	case SSIZE_SMALL:
		return 100
	case SSIZE_MEDIUM:
		return 150
	case SSIZE_LARGE:
		return 250
	case SSIZE_GIANT:
		return 400
	}
	return 0
}

func (sd *ShipDesign) GetUsedSpace() int {
	space := 0
	for _, w := range sd.Weapons {
		if w != nil {
			space += w.Weapon.GetSize(sd.GetTotalSpace()) * w.Count
		}
	}
	for _, w := range sd.Systems {
		if w != nil {
			space += w.GetSize(sd.GetTotalSpace())
		}
	}
	for _, w := range sd.SpecialSystems {
		if w != nil {
			space += w.GetSize(sd.GetTotalSpace())
		}
	}
	return space
}

func GetShipSizeName(size int) string {
	switch size {
	case SSIZE_SMALL:
		return "Small"
	case SSIZE_MEDIUM:
		return "Medium"
	case SSIZE_LARGE:
		return "Large"
	case SSIZE_GIANT:
		return "Giant"
	}
	panic("No such ship size!")
}

func (sd *ShipDesign) GetBcCost() int {
	cost := 0
	// TODO: consider size
	for i := range sd.Systems {
		if sd.Systems[i] != nil {
			cost += sd.Systems[i].cost
		}
	}
	for i := range sd.Weapons {
		if sd.Weapons[i] != nil && sd.Weapons[i].Weapon != nil {
			cost += sd.Weapons[i].Weapon.cost
		}
	}
	for i := range sd.SpecialSystems {
		if sd.SpecialSystems[i] != nil {
			cost += sd.SpecialSystems[i].cost
		}
	}

	return cost
}

func (sd *ShipDesign) HasSpecialSystemWithCode(spec sdsUniqueCode) bool {
	for sys := range sd.SpecialSystems {
		if sd.SpecialSystems[sys] != nil && sd.SpecialSystems[sys].uniqCode == spec {
			return true
		}
	}
	return false
}

func SetDefaultShipsDesignToFaction(f *faction) {
	f.shipsDesigns[0] = &ShipDesign{
		name: "Scout",
		Systems: [SDSLOT_COUNT]*ShipSystemStruct{
			SDSLOT_FUEL:       GetShipSystemByName("Basic fuel cells"),
			SDSLOT_PROPULSION: GetShipSystemByName("Nuclear engines"),
		},
	}
	f.shipsDesigns[1] = &ShipDesign{
		name: "Colony ship",
		Systems: [SDSLOT_COUNT]*ShipSystemStruct{
			SDSLOT_FUEL:       GetShipSystemByName("Basic fuel cells"),
			SDSLOT_PROPULSION: GetShipSystemByName("Nuclear engines"),
		},
		SpecialSystems: [4]*ShipSystemStruct{0: GetShipSystemByName("Colony")},
	}
}
