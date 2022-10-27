package game

import (
	"moocli/graphic_primitives"
)

type galaxyStruct struct {
	W, H     int
	stars    []*StarStruct
	fleets   []*Fleet
	factions []*faction
}

func (gs *galaxyStruct) GetStarAt(x, y int) *StarStruct {
	for _, str := range gs.stars {
		if str.X == x && str.Y == y {
			return str
		}
	}
	return nil
}

func (gs *galaxyStruct) GetAllStars() []*StarStruct {
	return gs.stars
}

func (gs *galaxyStruct) GetAllFleets() []*Fleet {
	return gs.fleets
}

func (gs *galaxyStruct) GetDistanceToCoordsForEmpire(x, y int, e *faction) int {
	dist := 9999999
	for _, s := range gs.stars {
		if s.planet.colonizedBy == e {
			tDist := graphic_primitives.GetLineDist(s.X, s.Y, x, y)
			if tDist < dist {
				dist = tDist
			}
		}
	}
	return dist
}

func (gs *galaxyStruct) GetFleetsAt(x, y int) (result []*Fleet) {
	for _, f := range gs.fleets {
		if f.x == x && f.y == y {
			result = append(result, f)
		}
	}
	return
}

func (gs *galaxyStruct) GetFleetOfFactionAt(x, y int, fact *faction) *Fleet {
	for _, f := range gs.fleets {
		if f.x == x && f.y == y && f.owner == fact {
			return f
		}
	}
	return nil
}

func (gs *galaxyStruct) CreateOrAppendFleetWithShipOfDesign(x, y int, shipFaction *faction, designNumber int) {
	fleetToAppend := gs.GetFleetOfFactionAt(x, y, shipFaction)
	if fleetToAppend == nil {
		fleetToAppend = &Fleet{
			x: x, y: y,
			destX: x, destY: y,
			owner: shipFaction}
		fleetToAppend.ResetDestination()
		gs.fleets = append(gs.fleets, fleetToAppend)
	}
	fleetToAppend.shipsByDesign[designNumber]++
}
