package game

type faction struct {
	isPlayerControlled bool
	storedBc           int
	colorName          string

	hasTech                 [TECH_CATEGORIES][TECH_IN_CATEGORY]bool
	canResearchTech         [TECH_CATEGORIES][TECH_IN_CATEGORY]bool
	CurrentResearchingTech  [TECH_CATEGORIES]int // -1 means no tech selected
	bcSpentInTechCategories [TECH_CATEGORIES]int

	bcInReserve int
}

func createFaction(colorName string) *faction {
	f := &faction{colorName: colorName}
	f.GenerateTechAllowances()
	for i := range f.CurrentResearchingTech {
		f.CurrentResearchingTech[i] = -1
	}
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
