package game

func (g *Game) adjustEcoSliderToEliminatePollution(p *planet) {
	if p.GetSliderLock(PSLIDER_ECO) {
		return
	}
	for i := 0; i < 100; i++ {
		if g.GetPlanetWaste(p) - g.GetPlanetWasteRemoval(p, true) == 0 {
			return
		}
		if g.GetPlanetWaste(p) > g.GetPlanetWasteRemoval(p, false) {
			p.ChangeSliderPercent(+1, PSLIDER_ECO)
		}
		// TODO: don't lower it on T-FORM or pop growth
		if g.GetPlanetWasteRemoval(p, true) - g.GetPlanetWaste(p) > 1 {
			p.ChangeSliderPercent(-1, PSLIDER_ECO)
		}
	}
}
