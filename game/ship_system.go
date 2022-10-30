package game

import "strings"

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

func (s *ShipSystemStruct) GetName() string {
	return s.name
}
