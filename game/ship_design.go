package game

type shipDesign struct {
	Name string
	// Weapons        [4]int
	// Armor
	// Shield
	// Computer
	Fuel           *ShipSystemStruct
	Engine         *ShipSystemStruct
	SpecialSystems [4]*ShipSystemStruct // holds indices
}

func (sd *shipDesign) GetBcCost() int {
	cost := 0
	// TODO: consider size
	cost += sd.Fuel.cost
	cost += sd.Engine.cost
	for i := range sd.SpecialSystems {
		if sd.SpecialSystems[i] != nil {
			cost += sd.SpecialSystems[i].cost
		}
	}

	return cost
}

func (sd *shipDesign) HasSpecialSystemWithCode(spec sdsUniqueCode) bool {
	for sys := range sd.SpecialSystems {
		if sd.SpecialSystems[sys] != nil && sd.SpecialSystems[sys].uniqCode == spec {
			return true
		}
	}
	return false
}

func SetDefaultShipsDesignToFaction(f *faction) {
	f.shipsDesigns[0] = &shipDesign{
		Name:   "Scout",
		Fuel:   GetShipSystemByName("Basic fuel cells"),
		Engine: GetShipSystemByName("Nuclear engines"),
	}
	f.shipsDesigns[1] = &shipDesign{
		Name:           "Colony ship",
		Fuel:           GetShipSystemByName("Basic fuel cells"),
		Engine:         GetShipSystemByName("Nuclear engines"),
		SpecialSystems: [4]*ShipSystemStruct{GetShipSystemByName("Colony"), nil, nil, nil},
	}
}
