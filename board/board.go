package board

import "top-down-tdd/abstractions"

type board struct{}

func NewBoard() abstractions.Board {
	return board{}
}

func (b board) HasWinner() bool {
	return true
}
