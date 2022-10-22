package game

import "moocli/math"

func (g *Game) GetMaxFactoriesForPlanet(p *planet) int {
	return p.pop * 2 // TODO: tech
}

func (g *Game) GetActiveFactoriesForPlanet(p *planet) int {
	activeFactoriesPerPopulation := 2 // TODO: tech-related modifier
	return math.MinInt(p.pop * activeFactoriesPerPopulation, p.factories)
}

func (g *Game) GetPlanetWaste(p *planet) int {
	return g.GetActiveFactoriesForPlanet(p)
}

func (g *Game) GetPlanetWasteRemoval(p *planet, gross bool) int {
	ecoBc := g.GetPlanetBCPerSlider(p)[PSLIDER_ECO]
	if gross {
		return ecoBc * 2
	}
	return math.MinInt(ecoBc * 2, g.GetPlanetWaste(p))
}

func (g *Game) GetPlanetProductionNetGross(p *planet) (int, int) {
	handLabor := p.pop/2 // TODO: race-specific hand labor factor
	factoriesLabor := g.GetActiveFactoriesForPlanet(p)
	gross := handLabor + factoriesLabor
	net := gross // TODO: taxes
	return net, gross
}

func (g *Game) GetPlanetBCPerSlider(p *planet) (bcPerSlider [TOTAL_PLANET_SLIDERS]int) {
	netProduction, _ := g.GetPlanetProductionNetGross(p)
	for i := range p.prodSliders {
		bcPerSlider[i] = p.prodSliders[i].percent * netProduction / 100
	}
	return
}
