package game

import (
	"fmt"
	"moocli/lib"
)

func generateGalaxy(w, h, desiredStarsCount int) *galaxyStruct {
	fmt.Println("Generating galaxy:")
	totalStars := desiredStarsCount // (w/5) * (h/5)
	gs := &galaxyStruct{
		W: w,
		H: h,
	}

	fmt.Println("   Placing stars...")
	for sn := 0; sn < totalStars; sn++ {
		gs.stars = append(gs.stars, generateNewStar(gs))
	}

	fmt.Println("   Adding factions...")
	gs.factions = make([]*faction, 0)
	for i := 0; i < 4; i++ {
		gs.factions = append(gs.factions, createFaction(FactionColors[i]))
	}

	fmt.Println("   Setting initial designs...")
	for _, f := range gs.factions {
		SetDefaultShipsDesignToFaction(f)
	}

	fmt.Println("   Placing homeworlds...")
	for _, f := range gs.factions {
		placeHomeworldForFaction(gs, f)
	}

	fmt.Println("   Finding Orion...")
	placeOrionSystem(gs)

	fmt.Println("   Placing player...")
	gs.factions[0].isPlayerControlled = true

	return gs
}

func generateNewStar(g *galaxyStruct) *StarStruct {
	minDistFromAnotherStar := 3
	minDistFromPrevStar := 10
	coordsSet := false
	var x, y int
	try := 0
	for !coordsSet {
		try++
		coordsSet = true
		x = rnd.RandInRange(1, g.W-1)
		y = rnd.RandInRange(1, g.H-1)
		if len(g.stars) > 0 {
			lastStar := g.stars[len(g.stars)-1]
			if lib.SqDistInt(x, y, lastStar.X, lastStar.Y) < minDistFromPrevStar*minDistFromPrevStar {
				coordsSet = false
				continue
			}
		}
		for _, otherStar := range g.stars {
			if lib.SqDistInt(x, y, otherStar.X, otherStar.Y) < minDistFromAnotherStar*minDistFromAnotherStar {
				coordsSet = false
				continue
			}
		}
		if try > 10000 {
			panic(fmt.Sprintf("Can't place new star! (placed %d)", len(g.stars)))
		}
	}

	nameSet := false
	selectedName := ""
	for !nameSet {
		nameSet = true
		selectedName = starNamesList[rnd.Rand(len(starNamesList))]
		for _, otherStar := range g.stars {
			if otherStar.Name == selectedName {
				nameSet = false
				continue
			}
		}
	}
	starTypeIndex := rnd.SelectRandomIndexFromWeighted(len(starsDataTable),
		func(x int) int { return starsDataTable[x].frequencyForGeneration },
	)

	planetTypeRoll := rnd.Rand(20)
	star := StarStruct{
		staticData: starsDataTable[starTypeIndex],
		Name:       selectedName,
		X:          x,
		Y:          y,
		planet: &planet{
			planetType: starsDataTable[starTypeIndex].selectPlanetTypeByRoll(planetTypeRoll),
		},
	}
	star.planet.baseMaxPop = rnd.RandInRange(1, 4) * sTablePlanets[star.planet.planetType].baseMaxPopulation

	// set growth
	star.planet.growth = PGROWTH_NORMAL
	if star.planet.planetType < PLANET_TYPE_MINIMAL {
		star.planet.growth = PGROWTH_HOSTILE
	}
	if star.planet.planetType > PLANET_TYPE_DESERT && rnd.Rand(12) == 0 {
		star.planet.growth = PGROWTH_FERTILE
		star.planet.baseMaxPop += 25 * star.planet.baseMaxPop / 100
	}

	// set planet special
	star.planet.special = PSPECIAL_NORMAL
	spRoll := rnd.Rand(20) + star.staticData.poornessRollModifier
	if spRoll <= 2 {
		star.planet.special = PSPECIAL_POOR
		if rnd.Rand(20)+star.staticData.poornessRollModifier <= 2 {
			star.planet.special = PSPECIAL_ULTRA_POOR
		}
	}
	spRoll = rnd.Rand(20) + star.staticData.richnessRollModifier
	if PLANET_TYPE_STEPPE-star.planet.planetType > spRoll {
		star.planet.special = PSPECIAL_RICH
		if rnd.Rand(20)+star.staticData.richnessRollModifier <= 6 {
			star.planet.special = PSPECIAL_ULTRA_RICH
		}
	}
	if star.planet.special == PSPECIAL_NORMAL && star.planet.planetType >= PLANET_TYPE_MINIMAL &&
		star.planet.planetType <= PLANET_TYPE_OCEAN {

		if rnd.Rand(20) < 2 {
			star.planet.special = PSPECIAL_ARTIFACTS
		}
	}

	return &star
}

func placeHomeworldForFaction(g *galaxyStruct, f *faction) {
	minDist := 1 * g.H / 2
	currIndex := rnd.Rand(len(g.stars))
	currStar := g.stars[currIndex%len(g.stars)]
	selected := false
	for !selected {
		selected = true
		currStar = g.stars[currIndex%len(g.stars)]
		if currStar.planet.IsColonized() {
			selected = false
		}
		for _, otherStar := range g.stars {
			if otherStar.planet.IsColonized() {
				if lib.SqDistInt(currStar.X, currStar.Y, otherStar.X, otherStar.Y) < minDist*minDist {
					selected = false
					break
				}
			}
		}
		currIndex++
	}
	currStar.planet.setColonyFor(f)
	g.CreateOrAppendFleetWithShipOfDesign(currStar.X, currStar.Y, f, 0)
	g.CreateOrAppendFleetWithShipOfDesign(currStar.X, currStar.Y, f, 0)
	g.CreateOrAppendFleetWithShipOfDesign(currStar.X, currStar.Y, f, 1)
	currStar.planet.planetType = PLANET_TYPE_TERRAN
	currStar.planet.growth = PGROWTH_NORMAL
	currStar.planet.special = PSPECIAL_NORMAL
	currStar.planet.baseMaxPop = 80
	currStar.planet.pop = 10
}

func placeOrionSystem(gs *galaxyStruct) {
	offset := rnd.Rand(len(gs.GetAllStars()))
	for {
		for _, s := range gs.stars {
			if s.GetPlanet().IsColonized() {
				continue
			}
			if offset <= 0 {
				s.Name = "Orion"
				s.staticData = starsDataTable[0]
				s.planet.planetType = PLANET_TYPE_GAIA
				s.planet.growth = PGROWTH_GAIA
				s.planet.special = PSPECIAL_ULTRA_RICH
				s.planet.baseMaxPop = 150
				return
			}
			offset--
		}
	}
}
