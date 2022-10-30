package game

type Fleet struct {
	owner        *faction
	x, y         int
	destX, destY int

	shipsByDesign [SHIP_DESIGNS_PER_FACTION]int
}

func (f *Fleet) ResetDestination() {
	f.destX, f.destY = f.x, f.y
}

func (f *Fleet) HasColonyShip() bool {
	for des := range f.shipsByDesign {
		if f.shipsByDesign[des] > 0 {
			if f.owner.shipsDesigns[des].HasSpecialSystemWithCode(SSYSTEM_COLONY) {
				return true
			}
		}
	}
	return false
}

func (f *Fleet) spendColonyShip() {
	for des := range f.shipsByDesign {
		if f.shipsByDesign[des] > 0 {
			if f.owner.shipsDesigns[des].HasSpecialSystemWithCode(SSYSTEM_COLONY) {
				f.shipsByDesign[des]--
				return
			}
		}
	}
	panic("Nothing to spend!")
}

func (f *Fleet) IsUnderWay() bool {
	return f.destX != f.x || f.destY != f.y
}

func (f *Fleet) GetTotalShipsNumber() int {
	sum := 0
	for i := range f.shipsByDesign {
		sum += f.shipsByDesign[i]
	}
	return sum
}

func (f *Fleet) GetMaxTravelingDistance() int {
	maxDist := 99999
	for des := range f.shipsByDesign {
		if f.shipsByDesign[des] > 0 {
			if f.owner.shipsDesigns[des].Systems[SDSLOT_FUEL].maxTraveledDistance < maxDist {
				maxDist = f.owner.shipsDesigns[des].Systems[SDSLOT_FUEL].maxTraveledDistance
			}
		}
	}
	return maxDist
}

func (f *Fleet) GetSpeed() int {
	minSpeed := 999999
	for des := range f.shipsByDesign {
		if f.shipsByDesign[des] > 0 {
			if f.owner.shipsDesigns[des].Systems[SDSLOT_PROPULSION].speedOnGlobalMap < minSpeed {
				minSpeed = f.owner.shipsDesigns[des].Systems[SDSLOT_PROPULSION].speedOnGlobalMap
			}
		}
	}
	return minSpeed
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

func (f *Fleet) setTargetCoords(x, y int) {
	f.destX, f.destY = x, y
}
