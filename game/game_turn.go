package game

func (g *Game) ProcessTurn() {
	g.Turn++
	for _, star := range g.Galaxy.stars {
		if star.GetPlanet().IsColonized() {
			indBC := g.GetPlanetBCForSlider(star.GetPlanet(), PSLIDER_IND)
			// ecoBC := g.GetPlanetBCForSlider(star.GetPlanet(), PSLIDER_ECO)

			g.growColonizedPlanetPop(star.planet)
			g.buildFactories(star.planet, indBC)

			g.adjustEcoSliderToEliminatePollution(star.GetPlanet())
		}
	}
}

func (g *Game) buildFactories(p *planet, spentBc int) {
	fct, rem := g.GetPlanetFactoriesConstructedAndRemainderBC(p, spentBc)
	p.factories += fct
	p.bcRemainingForFactory = rem
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
