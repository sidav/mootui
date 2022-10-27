package game

import "strings"

type sdsSlot int

const (
	SDSLOT_WEAPON sdsSlot = iota
	SDSLOT_PROPULSION
	SDSLOT_SHIELD
	SDSLOT_ARMOR
	SDSLOT_FUEL
	SDSLOT_SPECIAL
	SDSLOT_COUNT
)

type sdsUniqueCode int

const (
	SSYSTEM_NONE sdsUniqueCode = iota
	SSYSTEM_COLONY
)

func GetShipSystemByName(name string) *ShipSystemStruct {
	name = strings.ToLower(name)
	for slot := range ShipSystemsTable {
		for _, system := range ShipSystemsTable[slot] {
			if strings.ToLower(system.name) == name {
				return system
			}
		}
	}
	panic("System '" + name + "' not found!")
}

type ShipSystemStruct struct {
	name            string
	alwaysAvailable bool // true if no research required

	cost int

	maxTraveledDistance int

	speedOnGlobalMap int

	uniqCode sdsUniqueCode
}

var ShipSystemsTable = map[sdsSlot][]*ShipSystemStruct{
	SDSLOT_FUEL: {
		{
			name:                "Basic fuel cells",
			cost:                10,
			alwaysAvailable:     true,
			maxTraveledDistance: 3,
		},
	},
	SDSLOT_PROPULSION: {
		{
			name:             "Nuclear engines",
			cost:             10,
			alwaysAvailable:  true,
			speedOnGlobalMap: 1,
		},
	},
	SDSLOT_SPECIAL: {
		{
			name:            "Colony",
			alwaysAvailable: true,
			cost:            100,
			uniqCode:        SSYSTEM_COLONY,
		},
	},
}
