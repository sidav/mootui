package game

const TECH_CATEGORIES = 6
const TECH_IN_CATEGORY = 2

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
	// BIOLOGY
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
	//
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
}
