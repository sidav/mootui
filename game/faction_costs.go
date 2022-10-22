package game

func (f *faction) getFactoryCost() int {
	return 10
}

func (f *faction) getPopCost() int {
	return 20
}

func (f *faction) getActiveFactoriesPerPop() int {
	return 2
}

func (f *faction) getWasteRemovedFor1Bc() int {
	return 2
}
