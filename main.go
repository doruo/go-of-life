package main

import (
	"github.com/doruo/gameoflife/gol"
)

func main() {
	width, heigth := 5, 5
	lag := 1000
	game := gol.NewGame(width, heigth, lag)
	game.Run()
}
