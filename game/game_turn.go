package game

import (
	"fmt"
	"moocli/math"
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

			g.growColonizedPlanetPop(star)
			g.buildFactories(star, indBC)

			g.adjustEcoSliderToEliminatePollution(star.GetPlanet())
		}
	}
	for _, f := range g.Galaxy.factions {
		g.PerformResearchForFaction(f)
	}
}

func (g *Game) buildFactories(star *StarStruct, spentBc int) {
	p := star.planet
	if p.factories == g.GetMaxFactoriesForPlanet(p) {
		return
	}
	fct, rem := g.GetPlanetFactoriesConstructedAndRemainderBC(p, spentBc)
	p.factories += fct
	p.bcRemainingForFactory = rem
	if p.factories == g.GetMaxFactoriesForPlanet(p) {
		p.colonizedBy.addNotification(star.Name + " industry reached peak",
			fmt.Sprintf("Reached maximum of %d factories", p.factories))
	}
}

func (g *Game) growColonizedPlanetPop(star *StarStruct) {
	p := star.planet
	if p.pop == p.maxPop {
		return
	}
	totalGrowth := g.GetNaturalGrowthForPlanet(p)
	remEcoBc := g.GetPlanetBCForSlider(p, PSLIDER_ECO) - g.GetBcRequiredForPlanetWasteRemoval(p)
	if remEcoBc > 0 {
		totalGrowth += g.GetPopGrowthForBCs(p, remEcoBc)
	}
	if totalGrowth < 0 {
		negGrowth := math.AbsInt(totalGrowth)
		p.pop -= negGrowth / 10
		p.popTenths -= negGrowth % 10
		if p.popTenths < 0 {
			p.pop--
			p.popTenths += 10
		}
		p.colonizedBy.addNotification(star.Name + " has ecological disaster",
			fmt.Sprintf("Population reduces by %d.%d from industrial waste", negGrowth/10, negGrowth%10))
	} else {
		p.pop += totalGrowth / 10
		p.popTenths += totalGrowth % 10
		if p.popTenths >= 10 {
			p.pop++
			p.popTenths -= 10
		}
	}
	if p.pop >= p.maxPop {
		p.pop = p.maxPop
		p.popTenths = 0
		p.colonizedBy.addNotification(star.Name + " has grown to maximum",
			fmt.Sprintf("Reached maximum of %d population", p.pop))
	}

}
