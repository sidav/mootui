package game

type galaxyStruct struct {
	W, H     int
	stars    []*StarStruct
	fleets   []*Fleet
	factions []*faction
}

func (gs *galaxyStruct) GetStarAt(x, y int) *StarStruct {
	for _, str := range gs.stars {
		if str.X == x && str.Y == y {
			return str
		}
	}
	return nil
}

func (gs *galaxyStruct) GetAllStars() []*StarStruct {
	return gs.stars
}

func (gs *galaxyStruct) GetAllFleets() []*Fleet {
	return gs.fleets
}

func (gs *galaxyStruct) CreateOrAppendFleetWithShip(x, y int, shipFaction *faction) {
	// TODO: implement
	gs.fleets = append(gs.fleets, &Fleet{
		x: x, y: y,
		owner: shipFaction},
	)
}
