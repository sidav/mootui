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
	totalGrowth := g.GetNaturalGrowthForPlanet(p)
	remEcoBc := g.GetPlanetBCForSlider(p, PSLIDER_ECO) - g.GetBcRequiredForPlanetWasteRemoval(p)
	if remEcoBc > 0 {
		totalGrowth += g.GetPopGrowthForBCs(p, remEcoBc)
	}
	p.pop += totalGrowth / 10
	p.popTenths += totalGrowth % 10
	if p.popTenths >= 10 {
		p.pop++
		p.popTenths -= 10
	}
	if p.pop >= p.maxPop {
		p.pop = p.maxPop
		p.popTenths = 0
	}

}
