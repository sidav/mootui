package game

type faction struct {
	storedBc  int
	colorName string

	hasTech         [TECH_CATEGORIES][TECH_IN_CATEGORY]bool
	canResearchTech [TECH_CATEGORIES][TECH_IN_CATEGORY]bool

	bcInTechCategories [TECH_CATEGORIES]int
	bcInReserve        int
}

func createFaction(colorName string) *faction {
	f := &faction{colorName: colorName}
	f.GenerateTechAllowances()
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
