package main

import (
	"top-down-tdd/abstractions"
)

func main() {

}

func Start(game abstractions.Game) {
	game.InitGame()

	for !game.IsOver() {
		game.SetMark()
		game.ShowBoard()
	}

	game.ShowResultMessage()
}
