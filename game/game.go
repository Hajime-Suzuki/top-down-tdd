package game

import (
	"top-down-tdd/abstractions"
	"top-down-tdd/board"
	"top-down-tdd/player"
)

type game struct {
	board   abstractions.Board
	players []abstractions.Player

	inputHandler abstractions.InputHandler
}

func NewGame(inputHandler abstractions.InputHandler) abstractions.Game {
	return &game{
		inputHandler: inputHandler,
	}
}

func (g *game) InitGame() {
	g.board = board.NewBoard()

	userName1, err := g.inputHandler.GetUserInput()
	if err != nil {
		panic(err)
	}

	userName2, err := g.inputHandler.GetUserInput()

	if err != nil {
		panic(err)
	}

	g.players = []abstractions.Player{
		player.NewPlayer(userName1),
		player.NewPlayer(userName2),
	}
}

func (g *game) SetMark() {

}

func (g *game) IsOver() bool {
	return true
}

func (g *game) ShowBoard() {

}

func (g *game) ShowResultMessage() {

}
