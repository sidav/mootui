package game

import "strings"

type sdsSlot int

const (
	SDSLOT_WEAPON sdsSlot = iota
	SDSLOT_COMPUTER
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
	// return nil
	panic("System '" + name + "' not found!")
}

type ShipSystemStruct struct {
	name            string
	alwaysAvailable bool // true if no research required

	cost int

	// for weapons
	weaponRange, weaponAttackRating int
	// for computer
	toHitPercent int
	// for shield
	toDefendPercentBonus int
	// for armor
	armorRating int
	// for fuel
	maxTraveledDistance int
	// for engines
	speedOnGlobalMap int
	// for special
	uniqCode sdsUniqueCode
}

var ShipSystemsTable = map[sdsSlot][]*ShipSystemStruct{
	SDSLOT_WEAPON: {
		{
			name:               "Basic laser",
			weaponRange:        1,
			weaponAttackRating: 1,
			cost:               10,
			alwaysAvailable:    true,
		},
		{
			name:               "Gamma laser",
			weaponRange:        1,
			weaponAttackRating: 2,
			cost:               25,
			alwaysAvailable:    true,
		},
	},
	SDSLOT_COMPUTER: {
		{
			name:            "Basic laser",
			toHitPercent:    15,
			cost:            10,
			alwaysAvailable: true,
		},
	},
	SDSLOT_FUEL: {
		{
			name:                "Basic fuel cells",
			cost:                20,
			alwaysAvailable:     true,
			maxTraveledDistance: 3,
		},
		{
			name:                "Deuterium fuel cells",
			cost:                25,
			maxTraveledDistance: 4,
		},
		{
			name:                "Antimatter fuel cells",
			cost:                50,
			maxTraveledDistance: 5,
		},
	},
	SDSLOT_ARMOR: {
		{
			name:            "Duralloy armor",
			armorRating:     1,
			alwaysAvailable: true,
			cost:            15,
		},
	},
	SDSLOT_SHIELD: {
		{
			name:                 "Basic shield",
			toDefendPercentBonus: 5,
			cost:                 30,
		},
	},
	SDSLOT_PROPULSION: {
		{
			name:             "Nuclear engines",
			cost:             10,
			alwaysAvailable:  true,
			speedOnGlobalMap: 1,
		},
		{
			name:             "Thermonuclear engines",
			cost:             25,
			speedOnGlobalMap: 2,
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
