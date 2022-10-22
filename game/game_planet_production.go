package game

import (
	"fmt"
	"moocli/math"
)

func (g *Game) GetMaxFactoriesForPlanet(p *planet) int {
	return p.pop * 2 // TODO: tech
}

func (g *Game) GetActiveFactoriesForPlanet(p *planet) int {
	activeFactoriesPerPopulation := 2 // TODO: tech-related modifier
	return math.MinInt(p.pop*activeFactoriesPerPopulation, p.factories)
}

func (g *Game) GetPlanetWaste(p *planet) int {
	return g.GetActiveFactoriesForPlanet(p)
}

func (g *Game) GetPlanetFactoriesConstructedAndRemainderBC(p *planet, spentBcs int) (int, int) {
	maxFactories := g.GetMaxFactoriesForPlanet(p)
	if p.factories >= maxFactories {
		return 0, 0
	}
	spentBcs += p.bcRemainingForFactory
	factoryCost := p.colonizedBy.getFactoryCost() // TODO: consider tech
	builtFactories := math.MinInt(spentBcs/factoryCost, maxFactories-p.factories)
	return builtFactories, spentBcs % factoryCost
}

func (g *Game) GetPlanetWasteRemoval(p *planet, gross bool) int {
	ecoBc := g.GetPlanetBCForSlider(p, PSLIDER_ECO)
	if gross {
		return ecoBc * 2
	}
	return math.MinInt(ecoBc*2, g.GetPlanetWaste(p))
}

func (g *Game) GetPlanetProductionNetGross(p *planet) (int, int) {
	handLabor := p.pop / 2 // TODO: race-specific hand labor factor
	factoriesLabor := g.GetActiveFactoriesForPlanet(p)
	gross := handLabor + factoriesLabor
	net := gross // TODO: taxes
	return net, gross
}

func (g *Game) GetPlanetBCForSlider(p *planet, sNum int) int {
	netProduction, _ := g.GetPlanetProductionNetGross(p)
	return p.prodSliders[sNum].percent * netProduction / 100
}

func (g *Game) GetSliderString(p *planet, snum int) string {
	switch snum {
	case PSLIDER_SHIP:
		return "Not implemented"
	case PSLIDER_DEF:
		return "Not implemented"
	case PSLIDER_IND:
		spending := g.GetPlanetBCForSlider(p, PSLIDER_IND)
		if spending == 0 {
			return "None"
		}
		factCost := p.colonizedBy.getFactoryCost()
		if spending > factCost {
			buildPer10Turns := 10*spending/factCost
			return fmt.Sprintf("%d.%d/turn", buildPer10Turns/10, buildPer10Turns%10)
		} else {
			//                                    it's int division rounding up
			return fmt.Sprintf("%d turns", (factCost+spending-1)/spending)
		}
	case PSLIDER_ECO:
		if g.GetPlanetWasteRemoval(p, true) >= g.GetPlanetWaste(p) {
			return "Clean"
		}
		return "Waste"
	case PSLIDER_TECH:
		return "Not implemented"
	}
	panic("Error")
}
