package game

type StarStruct struct {
	Name       string
	staticData *starStaticTable
	X, Y       int
	colony *colony
}

func (s *StarStruct) GetColony() *colony {
	return s.colony
}

func (s *StarStruct) GetStarTypeName() string {
	return s.staticData.starTypeName
}
