package board

import (
	"fmt"
	"strconv"
	"top-down-tdd/abstractions"
)

type board struct {
	board [][]string
}

var BOARD_SIZE = 3

func NewBoard() abstractions.Board {
	return board{}
}

func (b board) IsOver() bool {
	return true
}

func (b board) Show() string {
	output := ""

	for rowNum, cols := range b.board {
		for colNum, col := range cols {

			symbol := col

			if symbol == "-" {
				symbol = strconv.Itoa(BOARD_SIZE*rowNum + colNum + 1)
			}

			output += fmt.Sprintf("%s ", symbol)
		}
		output += "\n"
	}

	return output
}

func (b board) GetWinner() string {
	return "TODO: IMPLEMENT"
}

func (b board) Update(mark string, position string) (abstractions.Board, error) {
	//TODO: IMPLEMENT
	return board{}, nil
}
