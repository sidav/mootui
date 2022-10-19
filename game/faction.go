package game

type faction struct {
	storedBc  int
	colorName string
}

func (f *faction) GetColorName() string {
	return f.colorName
}

var FactionColors = [...]string {
	"BLUE",
	"RED",
	"GREEN",
	"YELLOW",
}
