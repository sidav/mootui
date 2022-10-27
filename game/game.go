package game

import (
	"fmt"
	"moocli/fibrandom"
	"moocli/graphic_primitives"
)

var rnd fibrandom.FibRandom

type Game struct {
	Galaxy *galaxyStruct
	Turn   int
}

func (g *Game) GetPlayerFaction() *faction {
	for _, f := range g.Galaxy.factions {
		if f.isPlayerControlled {
			return f
		}
	}
	return nil
}

func InitNewGame() *Game {
	rnd.InitDefault()
	gam := Game{}
	gam.Turn = 1
	gam.Galaxy = generateGalaxy(32, 20, 36)
	fmt.Println("Game init finished.")
	return &gam
}

func Distance(fx, fy, tx, ty int) int {
	return graphic_primitives.GetLineDist(fx, fy, tx, ty)
}
