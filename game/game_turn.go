package game

import "moocli/math"

func (g *Game) ProcessTurn() {
	g.Turn++
	for _, star := range g.Galaxy.stars {
		if star.GetPlanet().IsColonized() {
			planetBC := g.GetPlanetBCPerSlider(star.GetPlanet())

			g.buildFactories(star.planet, planetBC[PSLIDER_IND])
			g.growColonizedPlanetPop(star.planet)

			g.adjustEcoSliderToEliminatePollution(star.GetPlanet())
		}
	}
}

func (g *Game) buildFactories(p *planet, spentBcs int) {
	maxFactories := g.GetMaxFactoriesForPlanet(p)
	if p.factories >= maxFactories {
		return
	}
	spentBcs += p.bcRemainingForFactory
	p.bcRemainingForFactory = 0
	factoryCost := 10 // TODO: consider tech
	builtFactories := math.MinInt(spentBcs/factoryCost, maxFactories-p.factories)
	p.bcRemainingForFactory = spentBcs % factoryCost
	p.factories += builtFactories
	// TODO: partial factories building
}

func (g *Game) growColonizedPlanetPop(p *planet) {
	grPerc := 100 - (100 * (p.pop + (g.GetPlanetWaste(p) - g.GetPlanetWasteRemoval(p, false))) / p.maxPop)
	switch p.growth {
	case PGROWTH_HOSTILE:
		grPerc /= 2
	case PGROWTH_FERTILE:
		grPerc = 3 * grPerc / 2
	case PGROWTH_GAIA:
		grPerc *= 2
	}
	naturalGrowth := grPerc*(p.pop+5)/100 + p.popTenths
	if naturalGrowth == 0 && grPerc > 0 {
		naturalGrowth = 1
	}
	p.pop += naturalGrowth / 10
	if p.pop < p.maxPop {
		p.popTenths += naturalGrowth % 10
		if p.popTenths >= 10 {
			p.pop++
			p.popTenths -= 10
		}
	}
	if p.pop >= p.maxPop {
		p.popTenths = 0
	}
	// TODO: partial growth
}
