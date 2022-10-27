package game

import "math"

const TECH_CATEGORIES = 6

var categoryNames = []string{
	"Computers",
	"Construction",
	"Force fields",
	"Planetology",
	"Propulsion",
	"Weapons",
}

func GetScienceCostForTech(cat, id int) int {
	const expFactor = 1.1
	id += 1
	idFloat := float64(id)
	return 220 * int(idFloat*math.Pow(expFactor, idFloat))
}

func GetTechByCatAndId(cat, id int) *techStruct {
	return techTable[cat][id]
}

func GetTechCategoryName(cat int) string {
	return categoryNames[cat]
}

type techStruct struct {
	Name                    string
	Description             string
	wasteRemovedPerCost     int
	factoryConstructionCost int
	factoriesPerPopulation  int
	terraformingPopAddition int
	givesShipSystemWithName string
	alwaysAvailable         bool
	unused                  bool
}
