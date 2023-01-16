package game

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
