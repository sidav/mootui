package game

type starStaticTable struct {
	starTypeName           string
	frequencyForGeneration int // opposite of rarity
}

var starsDataTable = [...]*starStaticTable{
	{
		starTypeName: "Yellow",
		frequencyForGeneration: 3,
	},
	{
		starTypeName: "White",
		frequencyForGeneration: 2,
	},
	{
		starTypeName: "Blue",
		frequencyForGeneration: 3,
	},
	{
		starTypeName: "Red",
		frequencyForGeneration: 6,
	},
	{
		starTypeName: "Green",
		frequencyForGeneration: 5,
	},
	{
		starTypeName: "Neutron",
		frequencyForGeneration: 1,
	},
}

var starNamesList = [...]string{
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
}
