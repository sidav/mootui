package game

import (
	"fmt"
	"moocli/lib"
)

func (g *Game) ProcessTurn() {
	g.Turn++
	for _, f := range g.Galaxy.factions {
		f.clearNotifications()
		g.AccumulateScienceForFaction(f, g.GetFactionScienceTotalProduced(f))
	}
	for _, star := range g.Galaxy.stars {
		if star.GetPlanet().IsColonized() {
			indBC := g.GetPlanetBCForSlider(star.GetPlanet(), PSLIDER_IND)
			// ecoBC := g.GetPlanetBCForSlider(star.GetPlanet(), PSLIDER_ECO)

			g.buildEco(star)
			g.buildIndustry(star, indBC)

			g.adjustEcoSliderToEliminatePollution(star.GetPlanet())
		}
	}
	for _, f := range g.Galaxy.factions {
		g.PerformResearchForFaction(f)
	}
}

func (g *Game) buildIndustry(star *StarStruct, spentBc int) {
	if spentBc == 0 {
		return
	}
	p := star.planet
	// first, look if any upgrade is needed
	if p.factoriesUpgradeNeeded() {
		p.bcSpentOnInd += spentBc
		if p.bcSpentOnInd >= p.factories * p.colonizedBy.getFactoryUpgradeCost() {
			p.currentFactoriesPerPop++
			p.bcSpentOnInd = 0
		}
		return
	}

	if p.factories == g.GetMaxFactoriesForPlanet(p) {
		return
	}
	fct, rem := g.GetPlanetFactoriesConstructedAndRemainderBC(p, spentBc)
	p.factories += fct
	p.bcSpentOnInd = rem
	if p.factories == g.GetMaxFactoriesForPlanet(p) {
		p.colonizedBy.addNotification(star.Name + " industry reached peak",
			fmt.Sprintf("Reached maximum of %d factories", p.factories))
	}
}

func (g *Game) buildEco(star *StarStruct) {
	p := star.planet
	totalGrowth := g.GetNaturalGrowthForPlanet(p)
	remainingEcoBc := g.GetPlanetBCForSlider(p, PSLIDER_ECO) - g.GetBcRequiredForPlanetWasteRemoval(p)
	if remainingEcoBc > 0 {
		if p.canBeTerraformed() {
			p.bcSpentOnTerraforming += remainingEcoBc
			if p.bcSpentOnTerraforming >= TERRAFORMING_COST_PER_POP {
				p.popGivenByTerraforming++
				p.bcSpentOnTerraforming = 0
			}
		} else {
			totalGrowth += g.GetPopGrowthForBCs(p, remainingEcoBc)
		}
	}
	if totalGrowth < 0 {
		negGrowth := lib.AbsInt(totalGrowth)
		p.pop -= negGrowth / 10
		p.popTenths -= negGrowth % 10
		if p.popTenths < 0 {
			p.pop--
			p.popTenths += 10
		}
		p.colonizedBy.addNotification(star.Name + " has ecological disaster",
			fmt.Sprintf("Population reduces by %d.%d from industrial waste", negGrowth/10, negGrowth%10))
	} else if p.pop < p.GetMaxPop() {
		p.pop += totalGrowth / 10
		p.popTenths += totalGrowth % 10
		if p.popTenths >= 10 {
			p.pop++
			p.popTenths -= 10
		}
	}
	if p.pop >= p.GetMaxPop() {
		p.pop = p.GetMaxPop()
		p.popTenths = 0
		p.colonizedBy.addNotification(star.Name + " has grown to maximum",
			fmt.Sprintf("Reached maximum of %d population", p.pop))
	}
}
