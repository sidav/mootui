package game

import (
	"moocli/math"
)

func (f *faction) GenerateTechAllowances() {
	for cat := range f.canResearchTech {
		for i := range f.canResearchTech[cat] {
			f.canResearchTech[cat][i] = true // techTable[cat][i].alwaysAvailable || rnd.OneChanceFrom(2)
			// !techTable[cat][i].unused
		}
	}
}

func (f *faction) GetMaxTechIdInCategory(cat int) int {
	highestLevel := -1
	for i := range f.hasTech[cat] {
		if f.hasTech[cat][i] {
			highestLevel = i
		}
	}
	return highestLevel
}

func (f *faction) GetResearchableTechIdsInCategory(cat int) ([]int) {
	retIds := make([]int, 0)
	maxTechGotten := f.GetMaxTechIdInCategory(cat)
	maxResearchableId := 0
	// todo: bugs here
	if maxTechGotten == 0 {
		maxResearchableId = 1
	}
	if maxTechGotten >= 0 {
		maxResearchableId = 5 * math.DivideRoundingUp(maxTechGotten+1, 5)
	}
	// todo: bugs end (?) here
	for i := 0; i <= maxResearchableId; i++ {
		if i >= len(techTable[cat]) {
			break
		}
		if f.canResearchTech[cat][i] && !f.hasTech[cat][i] {
			//ret = append(ret, techTable[cat][i])
			retIds = append(retIds, i)
		}
	}
	return retIds
}

func (f *faction) applyNewTech(cat, id int) {
	tech := GetTechByCatAndId(cat, id)
	if tech.wasteRemovedPerCost > 0 && tech.wasteRemovedPerCost > f.currentCumulativeTech.wasteRemovedPerCost {
		f.currentCumulativeTech.wasteRemovedPerCost = tech.wasteRemovedPerCost
	}
	if tech.factoryConstructionCost > 0 && tech.factoryConstructionCost < f.currentCumulativeTech.factoryConstructionCost {
		f.currentCumulativeTech.factoryConstructionCost = tech.factoryConstructionCost
	}
	if tech.factoriesPerPopulation > 0 && tech.factoriesPerPopulation > f.currentCumulativeTech.factoriesPerPopulation {
		f.currentCumulativeTech.factoriesPerPopulation = tech.factoriesPerPopulation
	}
}
