package game

func (g *Game) GetFactionSciencePerBc(f *faction) int {
	return 10 // TODO: tech
}

func (g *Game) GetFactionScienceTotalProduced(f *faction) int {
	bc := 0
	for _, star := range g.Galaxy.stars {
		if star.GetPlanet().colonizedBy == f {
			bc += g.GetPlanetBCForSlider(star.GetPlanet(), PSLIDER_TECH) * g.GetFactionSciencePerBc(f)
		}
	}
	return bc
}

func (g *Game) AccumulateScienceForFaction(f *faction, rp int) {
	for cat := range f.bcSpentInTechCategories {
		if f.CurrentResearchingTech[cat] == -1 {
			continue
		}
		rpCurrent := rp / 6 // todo: science category sliders
		f.bcSpentInTechCategories[cat] += rpCurrent
	}
}

func (g *Game) PerformResearchForFaction(f *faction) {
	for cat := range f.bcSpentInTechCategories {
		resId := f.CurrentResearchingTech[cat]
		if resId == -1 {
			continue
		}
		if f.bcSpentInTechCategories[cat] >= GetScienceCostForTech(cat, resId) {
			f.bcSpentInTechCategories[cat] = 0
			f.hasTech[cat][resId] = true
			f.applyNewTech(cat, resId)
			f.addNotification(GetTechCategoryName(cat) + " research completed", GetTechByCatAndId(cat, resId).Name)
			f.CurrentResearchingTech[cat] = -1
			// TODO: create faction notification about completed research
		}
	}
}
