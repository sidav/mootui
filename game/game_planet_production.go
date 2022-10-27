package game

import (
	"fmt"
	"moocli/lib"
)

func (g *Game) GetMaxFactoriesForPlanet(p *planet) int {
	return p.pop * p.currentFactoriesPerPop
}

func (g *Game) GetActiveFactoriesForPlanet(p *planet) int {
	activeFactoriesPerPopulation := p.colonizedBy.getActiveFactoriesPerPop()
	return lib.MinInt(p.pop*activeFactoriesPerPopulation, p.factories)
}

func (g *Game) GetPlanetWaste(p *planet) int {
	return g.GetActiveFactoriesForPlanet(p) // todo: tech?
}

func (g *Game) GetPlanetFactoriesConstructedAndRemainderBC(p *planet, spentBcs int) (int, int) {
	maxFactories := g.GetMaxFactoriesForPlanet(p)
	if p.factories >= maxFactories {
		return 0, 0
	}
	spentBcs += p.bcSpentOnInd
	factoryCost := p.colonizedBy.getFactoryCost()
	builtFactories := lib.MinInt(spentBcs/factoryCost, maxFactories-p.factories)
	return builtFactories, spentBcs % factoryCost
}

func (g *Game) GetPlanetWasteRemoval(p *planet, gross bool) int {
	ecoBc := g.GetPlanetBCForSlider(p, PSLIDER_ECO)
	if gross {
		return ecoBc * p.colonizedBy.getWasteRemovedFor1Bc()
	}
	return lib.MinInt(ecoBc * p.colonizedBy.getWasteRemovedFor1Bc(), g.GetPlanetWaste(p))
}

func (g *Game) GetBcRequiredForPlanetWasteRemoval(p *planet) int {
	return g.GetPlanetWaste(p) / p.colonizedBy.getWasteRemovedFor1Bc()
}

func (g *Game) GetPlanetProductionNetGross(p *planet) (int, int) {
	handLabor := p.pop / 2 // TODO: race-specific hand labor factor
	factoriesLabor := g.GetActiveFactoriesForPlanet(p)
	switch p.special {
	case PSPECIAL_ULTRA_POOR:
		factoriesLabor = lib.DivideRoundingUp(50*factoriesLabor, 100)
	case PSPECIAL_POOR:
		factoriesLabor = lib.DivideRoundingUp(75*factoriesLabor, 100)
	case PSPECIAL_RICH:
		factoriesLabor = lib.DivideRoundingUp(125*factoriesLabor, 100)
	case PSPECIAL_ULTRA_RICH:
		factoriesLabor = lib.DivideRoundingUp(150*factoriesLabor, 100)
	}
	gross := handLabor + factoriesLabor
	net := gross // TODO: taxes
	return net, gross
}

func (g *Game) GetPlanetBCForSlider(p *planet, sNum int) int {
	netProduction, _ := g.GetPlanetProductionNetGross(p)
	return p.prodSliders[sNum].percent * netProduction / 100
}

// returns growth multiplied by 10!
func (g *Game) GetPopGrowthForBCs(p *planet, bcs int) int {
	popGrowthCost := p.colonizedBy.getPopCost()
	return (bcs * 10) / popGrowthCost
}

// multiplied by 10!
func (g *Game) GetNaturalGrowthForPlanet(p *planet) int {
	grPerc := 100 - (100 * (p.pop + (g.GetPlanetWaste(p) - g.GetPlanetWasteRemoval(p, false))) / p.GetMaxPop())
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
	return naturalGrowth
}

func (g *Game) GetSliderString(p *planet, snum int) string {
	switch snum {
	case PSLIDER_SHIP:
		spending := g.GetPlanetBCForSlider(p, PSLIDER_SHIP)
		if spending == 0 {
			return "None"
		}
		shipCost := p.colonizedBy.shipsDesigns[p.CurrentBuiltShipDesignIndex].GetBcCost()
		if spending >= shipCost {
			buildPer10Turns := 10*spending/shipCost
			return fmt.Sprintf("%d.%d/turn", buildPer10Turns/10, buildPer10Turns%10)
		} else {
			return fmt.Sprintf("%d turns", lib.DivideRoundingUp(shipCost, spending))
		}
	case PSLIDER_DEF:
		return "Not implemented"
	case PSLIDER_IND:
		spending := g.GetPlanetBCForSlider(p, PSLIDER_IND)
		if spending == 0 {
			return "None"
		}
		if p.factoriesUpgradeNeeded() {
			return "Upgrade"
		}
		factCost := p.colonizedBy.getFactoryCost()
		if spending > factCost {
			buildPer10Turns := 10*spending/factCost
			return fmt.Sprintf("%d.%d/turn", buildPer10Turns/10, buildPer10Turns%10)
		} else {
			return fmt.Sprintf("%d turns", lib.DivideRoundingUp(factCost, spending))
		}
	case PSLIDER_ECO:
		if g.GetPlanetWasteRemoval(p, true) >= g.GetPlanetWaste(p) {
			remEcoBc := g.GetPlanetBCForSlider(p, PSLIDER_ECO) - g.GetBcRequiredForPlanetWasteRemoval(p)
			if remEcoBc > 0 {
				if p.canBeTerraformed() {
					return "Terraform"
				}
				g := g.GetPopGrowthForBCs(p, remEcoBc)
				if g > 0 && p.pop < p.GetMaxPop() {
					return fmt.Sprintf("+%d.%d pop", g/10, g%10)
				}
			}
			return "Clean"
		}
		return "Waste"
	case PSLIDER_TECH:
		return fmt.Sprintf("%dRP", g.GetFactionSciencePerBc(p.GetFaction()) * g.GetPlanetBCForSlider(p, PSLIDER_TECH))
	}
	panic("Error")
}
