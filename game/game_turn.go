package game

import "moocli/math"

func (g *Game) ProcessTurn() {
	g.Turn++
	for _, star := range g.Galaxy.stars {
		if star.GetPlanet().IsColonized() {
			planetBC := g.GetPlanetProductionPerSlider(star.GetPlanet())

			g.buildFactories(star.planet, planetBC[PSLIDER_IND])
			g.growColonizedPlanetPop(star.planet)
		}
	}
}

func (g *Game) buildFactories(p *planet, spentBcs int) {
	factoryCost := 10 // TODO: consider tech
	maxFactories := g.GetMaxFactoriesForPlanet(p)
	p.factories += math.MinInt(spentBcs/factoryCost, maxFactories - p.factories)
}

func (g *Game) growColonizedPlanetPop(p *planet) {
	grPerc := 100 - (100 * p.pop / p.maxPop)
	switch p.growth {
	case PGROWTH_HOSTILE:
		grPerc /= 2
	case PGROWTH_FERTILE:
		grPerc = 3 * grPerc / 2
	case PGROWTH_GAIA:
		grPerc *= 2
	}
	naturalGrowth := grPerc * (p.pop + 5) / 1000
	if naturalGrowth == 0 && grPerc > 0 {
		naturalGrowth = 1
	}
	p.pop += naturalGrowth
}
