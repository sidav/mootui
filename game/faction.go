package game

type faction struct {
	storedBc  int
	colorName string

	hasTech         [TECH_CATEGORIES][TECH_IN_CATEGORY]bool
	canResearchTech [TECH_CATEGORIES][TECH_IN_CATEGORY]bool
}

func createFaction(colorName string) *faction {
	f := &faction{colorName: colorName}
	f.GenerateTechAllowances()
	return f
}

func (f *faction) GetColorName() string {
	return f.colorName
}

func (f *faction) getFactoryCost() int {
	return 10
}

var FactionColors = [...]string{
	"BLUE",
	"RED",
	"GREEN",
	"YELLOW",
}
