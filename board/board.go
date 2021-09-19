package board

import "top-down-tdd/abstractions"

type board struct{}

func NewBoard() abstractions.Board {
	return board{}
}

func (b board) IsOver() bool {
	return true
}

func (b board) Show() string {
	return "TODO: IMPLEMENT"
}

func (b board) GetWinner() string {
	return "TODO: IMPLEMENT"
}
