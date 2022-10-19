package game

type colony struct {
	faction    *faction
	population int
}

func (c *colony) GetFaction() *faction {
	return c.faction
}
