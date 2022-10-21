package game

const (
	PSLIDER_SHIP = iota
	PSLIDER_DEF
	PSLIDER_IND
	PSLIDER_ECO
	PSLIDER_TECH
	TOTAL_PLANET_SLIDERS
)

type prodSliderStruct struct {
	percent int
	locked  bool
}

func (p *planet) EqualizeSliders() {
	perSlider := 100 / TOTAL_PLANET_SLIDERS
	for i := range p.prodSliders {
		p.prodSliders[i].percent = perSlider
	}
}

func (p *planet) GetSliderLock(ind int) bool {
	return p.prodSliders[ind].locked
}

func (p *planet) FlipSliderLock(ind int) {
	p.prodSliders[ind].locked = !p.prodSliders[ind].locked
}

func (p *planet) getSlidersSum() int {
	sum := 0
	for i := range p.prodSliders {
		sum += p.prodSliders[i].percent
	}
	return sum
}

func (p *planet) NormalizeSliders(excludedIndex int) {
	for p.getSlidersSum() < 100 {
		// increasing the lowest non-zero non-locked slider
		selectedSlider := -1
		for i := range p.prodSliders {
			if i == excludedIndex || p.prodSliders[i].locked {
				continue
			}
			if selectedSlider == -1 || p.prodSliders[i].percent < p.prodSliders[selectedSlider].percent {
				selectedSlider = i
			}
		}
		p.prodSliders[selectedSlider].percent++
	}
	for p.getSlidersSum() > 100 {
		// decreasing the highest non-locked slider
		selectedSlider := -1
		for i := range p.prodSliders {
			if i == excludedIndex || p.prodSliders[i].locked {
				continue
			}
			if selectedSlider == -1 || p.prodSliders[i].percent > p.prodSliders[selectedSlider].percent {
				selectedSlider = i
			}
		}
		p.prodSliders[selectedSlider].percent--
	}
}

func (p *planet) GetSliderPercent(num int) int {
	return p.prodSliders[num].percent
}

func (p *planet) ChangeSliderPercent(diff int, sliderNum int) {
	p.prodSliders[sliderNum].percent += diff
	if p.prodSliders[sliderNum].percent < 0 {
		p.prodSliders[sliderNum].percent = 0
	}
	if p.prodSliders[sliderNum].percent > 100 {
		p.prodSliders[sliderNum].percent = 100
	}
	p.NormalizeSliders(sliderNum)
}

func GetSliderName(s int) string {
	switch s {
	case PSLIDER_SHIP:
		return "SHIP"
	case PSLIDER_DEF:
		return "DEFENSE"
	case PSLIDER_IND:
		return "INDUSTRY"
	case PSLIDER_ECO:
		return "ECOLOGY"
	case PSLIDER_TECH:
		return "TECH"
	}
	return "ERROR"
}
