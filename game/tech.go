package game

import "moocli/math"

const TECH_CATEGORIES = 6
const TECH_IN_CATEGORY = 2

func GetScienceCostForTech(cat, id int) int {
	return 200 + math.PowInt(12, id+1)/math.PowInt(10, id+1)
}

type techStruct struct {
	Name                    string
	Description             string
	wasteRemovalCost        int
	wasteRemovedPerCost     int
	factoryConstructionCost int
	factoriesPerPopulation  int
	alwaysAvailable         bool
}

var techTable = [TECH_CATEGORIES][TECH_IN_CATEGORY]*techStruct{
	// COMPUTERS
	{
		{
			Name:                "Basic robotic controls",
			wasteRemovalCost:    2,
			wasteRemovedPerCost: 1,
			alwaysAvailable:     true,
		},
		{
			Name:                "Improved robotic controls",
			wasteRemovalCost:    3,
			wasteRemovedPerCost: 2,
		},
	},
	// CONSTRUCTION
	{
		{
			Name:                "Basic factories cost",
			wasteRemovalCost:    2,
			wasteRemovedPerCost: 1,
			alwaysAvailable:     true,
		},
		{
			Name:                "Improved factories cost",
			wasteRemovalCost:    3,
			wasteRemovedPerCost: 2,
		},
	},
	// FORCE FIELDS
	{
		{
			Name:                "Basic shields",
			alwaysAvailable:     true,
		},
		{
			Name:                "Improved shields",
		},
	},
	// PLANETOLOGY
	{
		{
			Name:                "Basic waste removal",
			wasteRemovalCost:    2,
			wasteRemovedPerCost: 1,
			alwaysAvailable:     true,
		},
		{
			Name:                "Improved waste removal",
			wasteRemovalCost:    3,
			wasteRemovedPerCost: 2,
		},
	},
	// PROPULSION
	{
		{
			Name:                "Basic engines",
			alwaysAvailable:     true,
		},
		{
			Name:                "Improved engines",
		},
	},
	// WEAPONS
	{
		{
			Name:                "Basic weapons",
			alwaysAvailable:     true,
		},
		{
			Name:                "Improved weapons",
		},
	},
}
