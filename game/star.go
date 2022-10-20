package game

type StarStruct struct {
	Name       string
	staticData *starStaticTable
	X, Y       int
	planet     *planet
}

func (s *StarStruct) GetPlanet() *planet {
	return s.planet
}

func (s *StarStruct) GetStarTypeName() string {
	return s.staticData.starTypeName
}
