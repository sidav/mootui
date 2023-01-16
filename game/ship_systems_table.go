package game

type sdsSlot int

const (
	SDSLOT_ARMOR sdsSlot = iota
	SDSLOT_SHIELD
	SDSLOT_PROPULSION
	SDSLOT_FUEL
	SDSLOT_COMPUTER
	SDSLOT_COUNT
	SDSLOT_WEAPON
	SDSLOT_SPECIAL
)

var ShipSystemsTable = map[sdsSlot][]*ShipSystemStruct{
	SDSLOT_WEAPON: {
		{
			name:               "Basic laser",
			weaponRange:        1,
			weaponAttackRating: 1,
			cost:               10,
			sizeAbsolute:       1,
			hullSpacePer1Size:  10,
			alwaysAvailable:    true,
		},
		{
			name:               "Gamma laser",
			weaponRange:        1,
			weaponAttackRating: 2,
			cost:               25,
			sizeAbsolute:       1,
			hullSpacePer1Size:  10,
			alwaysAvailable:    true,
		},
	},
	SDSLOT_COMPUTER: {
		{
			name:              "Electronic computer",
			toHitPercent:      15,
			cost:              10,
			sizeAbsolute:      0,
			hullSpacePer1Size: 4,
			alwaysAvailable:   true,
		},
	},
	SDSLOT_FUEL: {
		{
			name:                "Basic fuel cells",
			cost:                20,
			alwaysAvailable:     true,
			sizeAbsolute:        10,
			maxTraveledDistance: 3,
		},
		{
			name:                "Deuterium fuel cells",
			cost:                25,
			sizeAbsolute:        10,
			maxTraveledDistance: 4,
		},
		{
			name:                "Antimatter fuel cells",
			cost:                50,
			sizeAbsolute:        10,
			maxTraveledDistance: 5,
		},
	},
	SDSLOT_ARMOR: {
		{
			name:              "Duralloy armor",
			armorRating:       1,
			alwaysAvailable:   true,
			cost:              15,
			hullSpacePer1Size: 10,
		},
	},
	SDSLOT_SHIELD: {
		{
			name:                 "Basic shield",
			toDefendPercentBonus: 5,
			cost:                 30,
			hullSpacePer1Size:    5,
		},
	},
	SDSLOT_PROPULSION: {
		{
			name:              "Nuclear engines",
			cost:              10,
			hullSpacePer1Size: 5,
			alwaysAvailable:   true,
			speedOnGlobalMap:  1,
		},
		{
			name:              "Thermonuclear engines",
			cost:              25,
			hullSpacePer1Size: 5,
			speedOnGlobalMap:  2,
		},
	},
	SDSLOT_SPECIAL: {
		{
			name:            "Colony",
			alwaysAvailable: true,
			cost:            100,
			sizeAbsolute:    100,
			uniqCode:        SSYSTEM_COLONY,
		},
	},
}
