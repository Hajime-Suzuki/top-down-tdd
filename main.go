package main

import (
	"fmt"
	"os"
	"top-down-tdd/abstractions"
	"top-down-tdd/game"
)

func main() {
	g, err := game.MakeGame()
	if err != nil {
		fmt.Printf("failed to create game: %s\n", err)
		os.Exit(2)
	}
	Start(g)
}

func Start(game abstractions.Game) {
	game.InitGame()

	for !game.IsOver() {
		game.SetMark()
		game.ShowBoard()
	}

	game.ShowResultMessage()
}
