package game

import "top-down-tdd/abstractions"

type game struct {
	userInputHandler abstractions.InputHandler
}

func NewGame(inputHandler abstractions.InputHandler) abstractions.Game {
	return game{
		userInputHandler: inputHandler,
	}
}

func (g game) InitGame() {

}

func (g game) SetMark() {

}

func (g game) IsOver() bool {
	return true
}

func (g game) ShowBoard() {

}

func (g game) ShowResultMessage() {

}
