package game

type sdsSlot int

const (
	SDSLOT_SPECIAL sdsSlot = iota
	SDSLOT_PROPULSION
	SDSLOT_COUNT
)

type sdsUniqueCode int

const (
	SSYSTEM_NONE sdsUniqueCode = iota
	SSYSTEM_COLONY
)

type ShipSystemStruct struct {
	name            string
	alwaysAvailable bool // true if no research required

	uniqCode sdsUniqueCode
}

var STableSystems = [SDSLOT_COUNT][]*ShipSystemStruct{
	// specials
	{
		{
			name:            "Colony",
			alwaysAvailable: true,
			uniqCode:        SSYSTEM_COLONY,
		},
	},
}
