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

func (g *Game) SplitFleet(f *Fleet, newFleetCounts []int) {
	f2 := Fleet{
		owner:         f.owner,
		x:             f.x,
		y:             f.y,
		shipsByDesign: [6]int{},
	}
	for i := range f.shipsByDesign {
		f.shipsByDesign[i] -= newFleetCounts[i]
		f2.shipsByDesign[i] += newFleetCounts[i]
		if f.shipsByDesign[i] < 0 {
			panic("Wrong fleet split!")
		}
	}
}

func (g *Game) MergeFleets(f1 *Fleet, f2 *Fleet) {
	if f1.owner != f2.owner {
		panic("Merging fleets of different factions")
	}
	for i := 0; i < SHIP_DESIGNS_PER_FACTION; i++ {
		f1.shipsByDesign[i] += f2.shipsByDesign[i]
	}
	g.Galaxy.removeFleet(f2)
}
