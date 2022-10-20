package game

import "fmt"

func generateGalaxy(w, h, desiredStarsCount int) *galaxyStruct {
	fmt.Println("Generating galaxy...")
	totalStars := desiredStarsCount // (w/5) * (h/5)
	gs := &galaxyStruct{
		W: w,
		H: h,
	}

	for sn := 0; sn < totalStars; sn++ {
		gs.stars = append(gs.stars, generateNewStar(gs))
	}

	fmt.Println("   Adding factions...")
	for i := 0; i < 4; i++ {
		gs.factions = append(gs.factions, &faction{
			colorName: FactionColors[i],
		})
	}

	fmt.Println("   Placing homeworlds...")
	for _, f := range gs.factions {
		placeHomeworldForFaction(gs, f)
	}

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
			if sqDistInt(x, y, lastStar.X, lastStar.Y) < minDistFromPrevStar*minDistFromPrevStar {
				coordsSet = false
				continue
			}
		}
		for _, otherStar := range g.stars {
			if sqDistInt(x, y, otherStar.X, otherStar.Y) < minDistFromAnotherStar*minDistFromAnotherStar {
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
			colonizedBy: nil,
			planetType:  starsDataTable[starTypeIndex].selectPlanetTypeByRoll(planetTypeRoll),
			pop:         0,
			maxPop:      0,
		},
	}
	return &star
}

func placeHomeworldForFaction(g *galaxyStruct, f *faction) {
	minDist := 1 * g.H / 2
	currInd :=  rnd.Rand(len(g.stars))
	currStar := g.stars[currInd % len(g.stars)]
	selected := false
	for !selected {
		selected = true
		currStar = g.stars[currInd % len(g.stars)]
		if currStar.planet.IsColonized() {
			selected = false
		}
		for _, otherStar := range g.stars {
			if otherStar.planet.IsColonized() {
				if sqDistInt(currStar.X, currStar.Y, otherStar.X, otherStar.Y) < minDist*minDist {
					selected = false
					break
				}
			}
		}
		currInd++
	}
	currStar.planet.colonizedBy = f
}
