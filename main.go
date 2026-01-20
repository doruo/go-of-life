package main

import (
	"github.com/doruo/gameoflife/gol"
)

func main() {
	screenWidth, screenHeigth := 1920, 1080
	size, lag := 40, 5000

	game := gol.NewGame(screenWidth, screenHeigth, size, lag)
	game.Run()
}
