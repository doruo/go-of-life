package main

import "github.com/doruo/gameoflife/gol"

func main() {
	size, lag := 45, 800
	game := gol.NewGame(size, lag)
	game.Run()
}
