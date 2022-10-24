package game

func (g *Game) adjustEcoSliderToEliminatePollution(p *planet) {
	if p.GetSliderLock(PSLIDER_ECO) {
		return
	}
	for i := 0; i < 100; i++ {
		if g.GetPlanetWaste(p) - g.GetPlanetWasteRemoval(p, true) == 0 {
			return
		}
		// TODO: don't lower it on T-FORM or pop growth
		if g.GetPlanetWasteRemoval(p, true) - g.GetPlanetWaste(p) > p.colonizedBy.getWasteRemovedFor1Bc() {
			p.ChangeSliderPercent(-1, PSLIDER_ECO)
		}
		if g.GetPlanetWaste(p) > g.GetPlanetWasteRemoval(p, false) {
			p.ChangeSliderPercent(+1, PSLIDER_ECO)
		}
	}
}
