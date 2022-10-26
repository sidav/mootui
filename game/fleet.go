package game

type Fleet struct {
	owner        *faction
	x, y         int
	destX, destY int
}

func (f *Fleet) GetShipsNumber() int {
	return 3 // TODO: remove this stub
}

func (f *Fleet) GetMaxTravelingDistance() int {
	return 5 // TODO: remove this stub
}

func (f *Fleet) GetOwner() *faction {
	return f.owner
}

func (f *Fleet) GetCoords() (int, int) {
	return f.x, f.y
}

func (f *Fleet) GetTargetCoords() (int, int) {
	return f.destX, f.destY
}
