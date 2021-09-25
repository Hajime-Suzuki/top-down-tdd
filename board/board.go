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

func isEmpty(s string) bool {
	return s == "-"
}

func (b board) IsOver() bool {
	isOver := false

	// horizontal
	for row := 0; row < BOARD_SIZE; row++ {
		complete := true
		firstMarkInRow := b.board[row][0]

		for col := 0; col < BOARD_SIZE; col++ {
			currentMark := b.board[row][col]

			// if mark is "-" or symbol is not the same as the first item of the row, do not consider this row anymore
			if firstMarkInRow != currentMark || isEmpty(currentMark) {
				complete = false
				break
			}
		}

		// if all the marks in the row, game is over
		if complete == true {
			isOver = true
			break
		}
	}

	// vertical
	for col := 0; col < BOARD_SIZE; col++ {
		complete := true
		firstMarkInCol := b.board[0][col]

		for row := 0; row < BOARD_SIZE; row++ {
			currentMark := b.board[row][col]

			// if mark is "-" or symbol is not the same as the first item of the row, do not consider this row anymore
			if firstMarkInCol != currentMark || isEmpty(currentMark) {
				complete = false
				break
			}
		}

		// if all the marks in the row, game is over
		if complete == true {
			isOver = true
			break
		}
	}
	return isOver
}

func (b board) Show() string {
	output := ""

	for rowNum, cols := range b.board {
		for colNum, col := range cols {

			mark := col

			if isEmpty(mark) {
				mark = strconv.Itoa(BOARD_SIZE*rowNum + colNum + 1)
			}

			output += fmt.Sprintf("%s ", mark)
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
