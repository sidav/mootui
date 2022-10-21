package game

import "moocli/math"

func (g *Game) GetMaxFactoriesForPlanet(p *planet) int {
	return p.pop * 2 // TODO:
}

func (g *Game) GetPlanetProductionNetGross(p *planet) (int, int) {
	handLabor := p.pop/2 // TODO: race-specific hand labor factor
	activeFactoriesPerPopulation := 2 // TODO: tech-related modifier
	factoriesLabor := math.MinInt(p.pop * activeFactoriesPerPopulation, p.factories)
	gross := handLabor + factoriesLabor
	net := gross // TODO: taxes
	return net, gross
}

func (g *Game) GetPlanetProductionPerSlider(p *planet) (bcPerSlider [TOTAL_PLANET_SLIDERS]int) {
	netProduction, _ := g.GetPlanetProductionNetGross(p)
	for i := range p.prodSliders {
		bcPerSlider[i] = p.prodSliders[i].percent * netProduction / 100
	}
	return
}
