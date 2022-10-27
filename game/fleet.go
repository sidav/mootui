package game

type Fleet struct {
	owner         *faction
	x, y          int
	destX, destY  int

	shipsByDesign [SHIP_DESIGNS_PER_FACTION]int
}

func (f *Fleet) ResetDestination() {
	f.destX, f.destY = f.x, f.y
}

func (f *Fleet) HasColonyShip() bool {
	// todo: implement
	return true
}


func (f *Fleet) spendColonyShip() {
	// todo: implement
	for i := range f.shipsByDesign {
		if f.shipsByDesign[i] > 0 {
			f.shipsByDesign[i]--
			return
		}
	}
}

func (f *Fleet) IsUnderWay() bool {
	return f.destX != f.x || f.destY != f.y
}

func (f *Fleet) GetShipsNumber() int {
	return 3 // TODO: remove this stub
}

func (f *Fleet) GetMaxTravelingDistance() int {
	return 5 // TODO: remove this stub
}

func (f *Fleet) GetSpeed() int {
	return 1 // TODO: remove this stub
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

func (f *Fleet) SetTargetCoords(x, y int) {
	f.destX, f.destY = x, y
}
