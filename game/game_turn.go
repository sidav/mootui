package game

import (
	"fmt"
	"moocli/graphic_primitives"
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
			shipsBc := g.GetPlanetBCForSlider(star.GetPlanet(), PSLIDER_SHIP)
			indBC := g.GetPlanetBCForSlider(star.GetPlanet(), PSLIDER_IND)
			// ecoBC := g.GetPlanetBCForSlider(star.GetPlanet(), PSLIDER_ECO)

			g.buildShips(star, shipsBc)
			g.buildEco(star)
			g.buildIndustry(star, indBC)

			g.adjustEcoSliderToEliminatePollution(star.GetPlanet())
		}
	}
	for _, f := range g.Galaxy.factions {
		g.PerformResearchForFaction(f)
	}
	g.moveFleets()
}

func (g *Game) buildShips(star *StarStruct, spentBc int) {
	if spentBc == 0 {
		return
	}
	p := star.planet
	p.bcSpentOnShip += spentBc
	shipCost := p.colonizedBy.shipsDesigns[p.CurrentBuiltShipDesignIndex].GetBcCost()
	for p.bcSpentOnShip >= shipCost {
		p.bcSpentOnShip -= shipCost
		g.Galaxy.CreateOrAppendFleetWithShipOfDesign(star.X, star.Y, p.colonizedBy, p.CurrentBuiltShipDesignIndex)
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
	// grow pop
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
		if p.pop >= p.GetMaxPop() {
			p.pop = p.GetMaxPop()
			p.popTenths = 0
			//p.colonizedBy.addNotification(star.Name + " has grown to maximum",
			//	fmt.Sprintf("Reached maximum of %d population", p.pop))
		}
	}
}

func (g *Game) moveFleets() {
	for _, f := range g.Galaxy.fleets {
		if !f.IsUnderWay() {
			continue
		}
		movementLine := graphic_primitives.GetLine(f.x, f.y, f.destX, f.destY)
		for i, coord := range movementLine {
			if i > f.GetSpeed() {
				break
			}
			f.x, f.y = coord.GetCoords()
		}
	}
}

func (g *Game) IsStarColonizableByFleet(star *StarStruct, f *Fleet) bool {
	return star != nil && f.HasColonyShip() && !star.planet.IsColonized()
}

func (g *Game) OrderFleetToColonize(f *Fleet) {
	x, y := f.GetCoords()
	star := g.Galaxy.GetStarAt(x, y)
	if !g.IsStarColonizableByFleet(star, f) {
		panic("Broken colonize order requirements!")
	}
	f.spendColonyShip()
	star.planet.setColonyFor(f.owner)
}
