package game

import "moocli/lib"

func (g *Game) SetFleetDestination(f *Fleet, x, y int) {
	if g.Galaxy.GetDistanceToCoordsForEmpire(x, y, f.owner) > f.GetMaxTravelingDistance() {
		panic("Fuel unsatisfied")
	}
	f.setTargetCoords(x, y)
}

func (g *Game) GetETAForFleetToCoords(f *Fleet, x, y int) int {
	return lib.DivideRoundingUp(Distance(f.x, f.y, x, y), f.GetSpeed())
}
