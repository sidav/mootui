package game

import "strconv"

type planet struct {
	colonizedBy *faction
	planetType  int

	// sliders-related
	bcSpentOnShip               int
	CurrentBuiltShipDesignIndex int
	factories                   int
	currentFactoriesPerPop      int // stored so that upgrading will work
	bcSpentOnInd                int // remaining from previous turn, "unbuilt" factory, upgrade in progress etc
	pop                         int
	popTenths                   int // it's "3" in "4.3 pop"
	baseMaxPop                  int
	popGivenByTerraforming      int
	bcSpentOnTerraforming       int
	special, growth             int
	prodSliders                 [TOTAL_PLANET_SLIDERS]prodSliderStruct // hold production values
}

func (p *planet) setColonyFor(f *faction) {
	p.setSlidersToInitialValues()
	p.colonizedBy = f
	p.currentFactoriesPerPop = f.getActiveFactoriesPerPop()
	p.pop = 1
}

func (p *planet) GetMaxPop() int {
	return p.baseMaxPop + p.popGivenByTerraforming
}

func (p *planet) IsColonized() bool {
	return p.colonizedBy != nil
}

func (p *planet) GetFaction() *faction {
	return p.colonizedBy
}

func (p *planet) GetFactories() int {
	return p.factories
}

func (p *planet) GetPlanetTypeName() string {
	return sTablePlanets[p.planetType].name
}

func (p *planet) GetPopulationStrings() (string, string) {
	return strconv.Itoa(p.pop) + "." + strconv.Itoa(p.popTenths), strconv.Itoa(p.GetMaxPop())
}

func (p *planet) canBeTerraformed() bool {
	return p.popGivenByTerraforming < p.colonizedBy.currentCumulativeTech.terraformingPopAddition
}

func (p *planet) factoriesUpgradeNeeded() bool {
	return p.factories > 0 && p.currentFactoriesPerPop < p.colonizedBy.getActiveFactoriesPerPop()
}

const (
	PLANET_TYPE_NOT_HABITABLE = iota
	PLANET_TYPE_RADIATED
	PLANET_TYPE_TOXIC
	PLANET_TYPE_INFERNO
	PLANET_TYPE_DEAD
	PLANET_TYPE_TUNDRA
	PLANET_TYPE_BARREN
	PLANET_TYPE_MINIMAL
	PLANET_TYPE_DESERT
	PLANET_TYPE_STEPPE
	PLANET_TYPE_ARID
	PLANET_TYPE_OCEAN
	PLANET_TYPE_JUNGLE
	PLANET_TYPE_TERRAN
	PLANET_TYPE_GAIA
	PLANET_TYPE_NUM
)

type planetStatic struct {
	name              string
	baseMaxPopulation int
}

var sTablePlanets = map[int]*planetStatic{
	PLANET_TYPE_NOT_HABITABLE: {
		name:              "Not habitable",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_RADIATED: {
		name:              "Irradiated",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_TOXIC: {
		name:              "Toxic",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_INFERNO: {
		name:              "Inferno",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_DEAD: {
		name:              "Dead",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_TUNDRA: {
		name:              "Tundra",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_BARREN: {
		name:              "Barren",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_MINIMAL: {
		name:              "Minimal",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_DESERT: {
		name:              "Desert",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_STEPPE: {
		name:              "Steppe",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_ARID: {
		name:              "Arid",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_OCEAN: {
		name:              "Ocean",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_JUNGLE: {
		name:              "Jungle",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_TERRAN: {
		name:              "Terran",
		baseMaxPopulation: 10,
	},
	PLANET_TYPE_GAIA: {
		name:              "Gaia",
		baseMaxPopulation: 10,
	},
}
