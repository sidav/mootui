package game

type galaxyStruct struct {
	W, H  int
	stars []*StarStruct
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
