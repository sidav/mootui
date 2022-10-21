package game

import (
	"fmt"
	"moocli/fibrandom"
)

var rnd fibrandom.FibRandom

type Game struct {
	Galaxy *galaxyStruct
	Turn   int
}

func InitNewGame() *Game {
	rnd.InitDefault()
	gam := Game{}
	gam.Turn = 1
	gam.Galaxy = generateGalaxy(26, 12, 20)
	fmt.Println("Game init finished.")
	return &gam
}
