package game

func init() {
	performTechSanityChecks()
}

// should be run when initializing game, checks if data tables are sane
func performTechSanityChecks() {
	for i, _ := range techTable {
		for _, tech := range techTable[i] {
			if tech.givesShipSystemWithName != "" {
				if GetShipSystemByName(tech.givesShipSystemWithName) == nil {
					panic("No such system: " + tech.givesShipSystemWithName + " (given by " + tech.Name + ")")
				}
			}
		}
	}
}
