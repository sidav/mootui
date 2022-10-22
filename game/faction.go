package game

type faction struct {
	storedBc  int
	colorName string
}

func (f *faction) GetColorName() string {
	return f.colorName
}

func (f *faction) getFactoryCost() int {
	return 10
}

var FactionColors = [...]string {
	"BLUE",
	"RED",
	"GREEN",
	"YELLOW",
}
