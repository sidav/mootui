package game

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
			Name:                   "Improved robotic controls",
			factoriesPerPopulation: 3,
		},
		{
			unused: true,
		},
		{
			Name:                   "Multi-adjusted robotic controls",
			factoriesPerPopulation: 4,
		},
		{
			unused: true,
		},
		{
			Name:                   "Neural robotic controls",
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
		{
			Name:                    "Basic shields",
			// givesShipSystemWithName: "Basic shield",
			alwaysAvailable:         true,
		},
		{
			Name: "Improved shields",
		},
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
			Name:                    "Basic terraforming +10",
			terraformingPopAddition: 10,
		},
		{
			Name:                "Further improved waste removal",
			wasteRemovedPerCost: 2,
		},
		{
			unused: true,
		},
		{
			Name:                    "Improved terraforming +20",
			terraformingPopAddition: 20,
		},
		{
			Name:                "Very improved waste removal",
			wasteRemovedPerCost: 3,
		},
		{
			unused: true,
		},
		{
			Name:                    "Deep terraforming +30",
			terraformingPopAddition: 30,
		},
		{
			Name:                "Extremely improved waste removal",
			wasteRemovedPerCost: 4,
		},
	},
	// PROPULSION
	{
		{
			Name:            "Basic engines",
			alwaysAvailable: true,
		},
		{
			Name: "Improved engines",
		},
	},
	// WEAPONS
	{
		{
			Name:            "Basic weapons",
			alwaysAvailable: true,
		},
		{
			Name: "Improved weapons",
		},
	},
}
