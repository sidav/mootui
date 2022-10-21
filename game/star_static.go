package game

type starStaticTable struct {
	starTypeName           string
	frequencyForGeneration int      // opposite of rarity
	planetTypesRoll        [][2]int // SHOULD BE ORDERED. first int - planet type, second is MAX INCLUSIVE roll in 0-20 (roll can be negative!)
}

func (sst *starStaticTable) selectPlanetTypeByRoll(roll int) int {
	for _, data := range sst.planetTypesRoll {
		if roll <= data[1] {
			return data[0]
		}
	}
	panic("wat")
}

var starsDataTable = [...]*starStaticTable{
	{
		starTypeName:           "Yellow",
		frequencyForGeneration: 3,
		planetTypesRoll: [][2]int{
			{PLANET_TYPE_NOT_HABITABLE, -1},
			{PLANET_TYPE_RADIATED, 1},
			{PLANET_TYPE_TOXIC, 2},
			{PLANET_TYPE_INFERNO, 3},
			{PLANET_TYPE_DEAD, 4},
			{PLANET_TYPE_TUNDRA, 5},
			{PLANET_TYPE_BARREN, 6},
			{PLANET_TYPE_MINIMAL, 7},
			{PLANET_TYPE_DESERT, 8},
			{PLANET_TYPE_STEPPE, 9},
			{PLANET_TYPE_ARID, 10},
			{PLANET_TYPE_OCEAN, 11},
			{PLANET_TYPE_JUNGLE, 12},
			{PLANET_TYPE_TERRAN, 16},
			{PLANET_TYPE_GAIA, 20},
		},
	},
	{
		starTypeName:           "White",
		frequencyForGeneration: 2,
		planetTypesRoll: [][2]int{
			{PLANET_TYPE_NOT_HABITABLE, -1},
			{PLANET_TYPE_RADIATED, 1},
			{PLANET_TYPE_TOXIC, 2},
			{PLANET_TYPE_INFERNO, 3},
			{PLANET_TYPE_DEAD, 4},
			{PLANET_TYPE_TUNDRA, 5},
			{PLANET_TYPE_BARREN, 6},
			{PLANET_TYPE_MINIMAL, 7},
			{PLANET_TYPE_DESERT, 8},
			{PLANET_TYPE_STEPPE, 9},
			{PLANET_TYPE_ARID, 10},
			{PLANET_TYPE_OCEAN, 11},
			{PLANET_TYPE_JUNGLE, 12},
			{PLANET_TYPE_TERRAN, 16},
			{PLANET_TYPE_GAIA, 20},
		},
	},
	{
		starTypeName:           "Blue",
		frequencyForGeneration: 3,
		planetTypesRoll: [][2]int{
			{PLANET_TYPE_NOT_HABITABLE, -1},
			{PLANET_TYPE_RADIATED, 1},
			{PLANET_TYPE_TOXIC, 2},
			{PLANET_TYPE_INFERNO, 3},
			{PLANET_TYPE_DEAD, 4},
			{PLANET_TYPE_TUNDRA, 5},
			{PLANET_TYPE_BARREN, 6},
			{PLANET_TYPE_MINIMAL, 7},
			{PLANET_TYPE_DESERT, 8},
			{PLANET_TYPE_STEPPE, 9},
			{PLANET_TYPE_ARID, 10},
			{PLANET_TYPE_OCEAN, 11},
			{PLANET_TYPE_JUNGLE, 12},
			{PLANET_TYPE_TERRAN, 16},
			{PLANET_TYPE_GAIA, 20},
		},
	},
	{
		starTypeName:           "Red",
		frequencyForGeneration: 6,
		planetTypesRoll: [][2]int{
			{PLANET_TYPE_NOT_HABITABLE, -1},
			{PLANET_TYPE_RADIATED, 1},
			{PLANET_TYPE_TOXIC, 2},
			{PLANET_TYPE_INFERNO, 3},
			{PLANET_TYPE_DEAD, 4},
			{PLANET_TYPE_TUNDRA, 5},
			{PLANET_TYPE_BARREN, 6},
			{PLANET_TYPE_MINIMAL, 7},
			{PLANET_TYPE_DESERT, 8},
			{PLANET_TYPE_STEPPE, 9},
			{PLANET_TYPE_ARID, 10},
			{PLANET_TYPE_OCEAN, 11},
			{PLANET_TYPE_JUNGLE, 12},
			{PLANET_TYPE_TERRAN, 16},
			{PLANET_TYPE_GAIA, 20},
		},
	},
	{
		starTypeName:           "Green",
		frequencyForGeneration: 5,
		planetTypesRoll: [][2]int{
			{PLANET_TYPE_NOT_HABITABLE, -1},
			{PLANET_TYPE_RADIATED, 1},
			{PLANET_TYPE_TOXIC, 2},
			{PLANET_TYPE_INFERNO, 3},
			{PLANET_TYPE_DEAD, 4},
			{PLANET_TYPE_TUNDRA, 5},
			{PLANET_TYPE_BARREN, 6},
			{PLANET_TYPE_MINIMAL, 7},
			{PLANET_TYPE_DESERT, 8},
			{PLANET_TYPE_STEPPE, 9},
			{PLANET_TYPE_ARID, 10},
			{PLANET_TYPE_OCEAN, 11},
			{PLANET_TYPE_JUNGLE, 12},
			{PLANET_TYPE_TERRAN, 16},
			{PLANET_TYPE_GAIA, 20},
		},
	},
	{
		starTypeName:           "Neutron",
		frequencyForGeneration: 1,
		planetTypesRoll: [][2]int{
			{PLANET_TYPE_NOT_HABITABLE, -1},
			{PLANET_TYPE_RADIATED, 1},
			{PLANET_TYPE_TOXIC, 2},
			{PLANET_TYPE_INFERNO, 3},
			{PLANET_TYPE_DEAD, 4},
			{PLANET_TYPE_TUNDRA, 5},
			{PLANET_TYPE_BARREN, 6},
			{PLANET_TYPE_MINIMAL, 7},
			{PLANET_TYPE_DESERT, 8},
			{PLANET_TYPE_STEPPE, 9},
			{PLANET_TYPE_ARID, 10},
			{PLANET_TYPE_OCEAN, 11},
			{PLANET_TYPE_JUNGLE, 12},
			{PLANET_TYPE_TERRAN, 16},
			{PLANET_TYPE_GAIA, 20},
		},
	},
}

var starNamesList = [...]string{
	// original MOO1 names
	"Ajax", "Alcor", "Anraq", "Antares", "Aquilae", "Argus", "Arietis", "Artemis", "Aurora", "Berel", "Beta Ceti",
	"Bootis", "Capella", "Celtsi", "Centauri", "Collassa", "Crius", "Crypto", "Cygni", "Darrian", "Denubius", "Dolz",
	"Draconis", "Drakka", "Dunatis", "Endoria", "Escalon", "Esper", "Exis", "Firma", "Galos", "Gienah", "Gion", "Gorra",
	"Guradas", "Helos", "Herculis", "Hyades", "Hyboria", "Imra", "Incedius", "Iranha", "Jinga", "Kailis", "Kakata",
	"Keeta", "Klystron", "Kronos", "Kulthos", "Laan", "Lyae", "Maalor", "Maretta", "Misha", "Mobas", "Moro", "Morrig",
	"Mu Delphi", "Neptunus", "Nitzer", "Nordia", "Nyarl", "Obaca", "Omicron", "Paladia", "Paranar", "Phantos", "Phyco",
	"Pollus", "Primodius", "Proteus", "Proxima", "Quayal", "Rana", "Rayden", "Regulus", "Reticuli", "Rha", "Rhilus",
	"Rigel", "Romulas", "Rotan", "Ryoun", "Seidon", "Selia", "Simius", "Spica", "Stalaz", "Talas", "Tao", "Tau Cygni",
	"Tauri", "Thrax", "Toranor", "Trax", "Tyr", "Ukko", "Uxmai", "Vega", "Volantis", "Vox", "Vulcan", "Whynil", "Willow",
	"Xendalla", "Xengara", "Xudax", "Yarrow", "Zhardan", "Zoctan",
	// my names
	"Tau Ceti", "Canis", "Cepheus", "Perseus", "Auriga",
}
