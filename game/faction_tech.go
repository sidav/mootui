package game

func (f *faction) GenerateTechAllowances() {
	for cat := range f.canResearchTech {
		for i := range f.canResearchTech[cat] {
			f.canResearchTech[cat][i] = techTable[cat][i].alwaysAvailable || rnd.OneChanceFrom(2)
		}
	}
}

func (f *faction) GetGeneralTechLevel() int {
	highestLevel := 0
	totalLevel := 0
	for cat := range f.canResearchTech {
		levelInCat := 0
		for i := range f.canResearchTech[cat] {
			if f.hasTech[cat][i] {
				levelInCat = i
			}
		}
		if levelInCat > highestLevel {
			highestLevel = levelInCat
		}
		totalLevel += levelInCat
	}
	return totalLevel - (2*highestLevel/10)
}

func (f *faction) GetResearchableTechesInCategory(cat int) ([]*techStruct, []int) {
	ret := make([]*techStruct, 0)
	retIds := make([]int, 0)
	ftl := f.GetGeneralTechLevel()
	for i := 0; i <= ftl; i++ {
		if f.canResearchTech[cat][i] && !f.hasTech[cat][i] {
			ret = append(ret, techTable[cat][i])
			retIds = append(retIds, i)
		}
	}
	return ret, retIds
}
