package game

const SHIP_DESIGNS_PER_FACTION = 6

type faction struct {
	isPlayerControlled bool
	storedBc           int
	colorName          string

	hasTech                 [TECH_CATEGORIES][]bool
	canResearchTech         [TECH_CATEGORIES][]bool
	CurrentResearchingTech  [TECH_CATEGORIES]int // -1 means no tech selected
	bcSpentInTechCategories [TECH_CATEGORIES]int
	// It will be updated when any new tech is acquired
	// should be used ONLY in calculations
	currentCumulativeTech techStruct

	bcInReserve int // TODO: use

	notificationsForThisTurn []*notificationStruct

	shipsDesigns [SHIP_DESIGNS_PER_FACTION]*shipDesign
}

func (f *faction) GetDesignByIndex(ind int) *shipDesign {
	return f.shipsDesigns[ind]
}

func createFaction(colorName string) *faction {
	f := &faction{colorName: colorName}
	for cat := 0; cat < TECH_CATEGORIES; cat++ {
		f.hasTech[cat] = make([]bool, len(techTable[cat]))
		f.canResearchTech[cat] = make([]bool, len(techTable[cat]))
	}
	f.GenerateTechAllowances()
	//for i := range f.CurrentResearchingTech {
	//	f.CurrentResearchingTech[i] = -1
	//}
	return f
}

func (f *faction) GetColorName() string {
	return f.colorName
}

var FactionColors = [...]string{
	"BLUE",
	"RED",
	"GREEN",
	"YELLOW",
}
