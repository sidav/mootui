package game

import (
	"fmt"
	"moocli/fibrandom"
)

var rnd fibrandom.FibRandom

type Game struct {
	Galaxy *galaxyStruct
}

func InitNewGame() *Game {
	rnd.InitDefault()
	gam := Game{}
	gam.Galaxy = generateGalaxy(26, 12, 20)
	fmt.Println("Game init finished.")
	return &gam
}
