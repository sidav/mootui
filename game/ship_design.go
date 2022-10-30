package game

const (
	SSIZE_SMALL = iota
	SSIZE_MEDIUM
	SSIZE_LARGE
	SSIZE_GIANT
	SSIZES_COUNT
)

type ShipDesign struct {
	Name           string
	Size           int
	Weapons        [4]*WeaponInstallation
	Systems        [SDSLOT_COUNT]*ShipSystemStruct
	SpecialSystems [4]*ShipSystemStruct // holds indices
}

type WeaponInstallation struct {
	Weapon *ShipSystemStruct
	Count  int
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
		if sd.Weapons[i].Weapon != nil {
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
		Name:   "Scout",
		Systems: [SDSLOT_COUNT]*ShipSystemStruct{
			SDSLOT_FUEL: GetShipSystemByName("Basic fuel cells"),
			SDSLOT_PROPULSION: GetShipSystemByName("Nuclear engines"),
		},
	}
	f.shipsDesigns[1] = &ShipDesign{
		Name:           "Colony ship",
		Systems: [SDSLOT_COUNT]*ShipSystemStruct{
			SDSLOT_FUEL: GetShipSystemByName("Basic fuel cells"),
			SDSLOT_PROPULSION: GetShipSystemByName("Nuclear engines"),
		},
		SpecialSystems: [4]*ShipSystemStruct{0: GetShipSystemByName("Colony")},
	}
}
