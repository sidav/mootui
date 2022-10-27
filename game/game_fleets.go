package game

func (g *Game) SetFleetDestination(f *Fleet, x, y int) {
	if g.Galaxy.GetDistanceToCoordsForEmpire(x, y, f.owner) > f.GetMaxTravelingDistance() {
		panic("Fuel unsatisfied")
	}
	f.setTargetCoords(x, y)
}