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
	lines := [][]string{}

	// add horizontal lines
	for row := 0; row < BOARD_SIZE; row++ {
		lines = append(lines, b.board[row])
	}

	// add vertical lines
	for col := 0; col < BOARD_SIZE; col++ {
		verticals := []string{}
		for row := 0; row < BOARD_SIZE; row++ {
			verticals = append(verticals, b.board[row][col])
		}
		lines = append(lines, verticals)
	}

	// add diagonal lines
	diagonal := [][]string{{}, {}}
	for i := 0; i < BOARD_SIZE; i++ {
		diagonal[0] = append(diagonal[0], b.board[i][i])
		diagonal[1] = append(diagonal[1], b.board[i][BOARD_SIZE-1-i])
	}
	lines = append(lines, diagonal...)

	return isComplete(lines)
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

func isEmpty(s string) bool {
	return s == "-"
}

// checks if line has the same 3 marks
func isComplete(lines [][]string) bool {
	isCompleted := false

	for _, ls := range lines {
		if isCompleted {
			break
		}

		for i := 1; i < BOARD_SIZE; i++ {
			current := ls[i]
			prev := ls[i-1]

			if current != prev || isEmpty(current) {
				break
			}

			if i == BOARD_SIZE-1 {
				isCompleted = true
			}
		}
	}
	return isCompleted
}
