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
	alwaysAvailable         bool
	unused                  bool
}

var techTable = [TECH_CATEGORIES][]*techStruct{
	// COMPUTERS
	{
		{
			Name:                   "Basic robotic controls",
			factoriesPerPopulation: 2,
			alwaysAvailable:        true,
		},
		{
			unused: true,
		},
		{
			Name:                "Improved robotic controls",
			factoriesPerPopulation: 3,
		},
		{
			unused: true,
		},
		{
			Name:                "Multi-adjusted robotic controls",
			factoriesPerPopulation: 4,
		},
		{
			unused: true,
		},
		{
			Name:                "Neural robotic controls",
			factoriesPerPopulation: 5,
		},
	},
	// CONSTRUCTION
	{
		{
			Name:                    "Basic factories cost",
			factoryConstructionCost: 20,
			alwaysAvailable:         true,
		},
		{
			unused: true,
		},
		{
			Name:                    "Improved factories cost",
			factoryConstructionCost: 15,
		},
		{
			unused: true,
		},
		{
			Name:                    "Unified factories components",
			factoryConstructionCost: 12,
		},
		{
			unused: true,
		},
		{
			Name:                    "Factories 3D-printing",
			factoryConstructionCost: 10,
		},
	},
	// FORCE FIELDS
	{
		//{
		//	Name:            "Basic shields",
		//	alwaysAvailable: true,
		//},
		//{
		//	Name: "Improved shields",
		//},
	},
	// PLANETOLOGY
	{
		{
			Name:                "Basic waste removal",
			wasteRemovedPerCost: 1,
			alwaysAvailable:     true,
		},
		{
			unused: true,
		},
		{
			unused: true,
		},
		{
			Name:                "Improved waste removal",
			wasteRemovedPerCost: 2,
		},
		{
			unused: true,
		},
		{
			unused: true,
		},
		{
			Name:                "Further improved waste removal",
			wasteRemovedPerCost: 2,
		},
		{
			unused: true,
		},
		{
			unused: true,
		},
		{
			Name:                "Very improved waste removal",
			wasteRemovedPerCost: 3,
		},
		{
			unused: true,
		},
		{
			unused: true,
		},
		{
			Name:                "Extremely improved waste removal",
			wasteRemovedPerCost: 4,
		},
	},
	// PROPULSION
	{
		//{
		//	Name:            "Basic engines",
		//	alwaysAvailable: true,
		//},
		//{
		//	Name: "Improved engines",
		//},
	},
	// WEAPONS
	{
		//{
		//	Name:            "Basic weapons",
		//	alwaysAvailable: true,
		//},
		//{
		//	Name: "Improved weapons",
		//},
	},
}
