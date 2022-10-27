package game

func (f *faction) getMaxTerraformingPopIncrease() int {
	return f.currentCumulativeTech.terraformingPopAddition
}


func (f *faction) getFactoryCost() int {
	cost := f.currentCumulativeTech.factoryConstructionCost
	if cost == 0 {
		cost = 30
	}
	return cost
}

func (f *faction) getFactoryUpgradeCost() int {
	return (f.getFactoryCost() + 1) / 2
}

func (f *faction) getPopCost() int {
	return 20
}

func (f *faction) getActiveFactoriesPerPop() int {
	amount := f.currentCumulativeTech.factoriesPerPopulation
	if amount == 0 {
		amount = 1
	}
	return amount
}

func (f *faction) getWasteRemovedFor1Bc() int {
	amount := f.currentCumulativeTech.wasteRemovedPerCost
	if amount == 0 {
		amount = 1
	}
	return amount
}

func (f *faction) getSensorsRange() int {
	return 3
}
