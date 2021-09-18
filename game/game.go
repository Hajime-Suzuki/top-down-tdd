package game

import "top-down-tdd/abstractions"

type Game struct {
	userInputHandler abstractions.InputHandler
}

func NewGame(inputHandler abstractions.InputHandler) Game {
	return Game{
		userInputHandler: inputHandler,
	}
}
