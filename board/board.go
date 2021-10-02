package board

import (
	"fmt"
	"strconv"
	"top-down-tdd/abstractions"
)

type board struct {
	board  [][]string
	*stats // calculated when IsOver or GetWinner is called
}

type stats struct {
	winner   string
	hasEmpty bool
}

const BOARD_SIZE = 3

func NewBoard() abstractions.Board {
	initialBoard := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	return board{board: initialBoard}
}

func (b board) IsOver() bool {
	s := b.calculateStats()
	return s.winner != "" || !s.hasEmpty
}

func (b board) Show() string {
	output := ""

	for rowNum, cols := range b.board {
		for colNum, mark := range cols {

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
	s := b.calculateStats()
	return s.winner
}

func (b board) Update(playerMark string, position string) (abstractions.Board, error) {
	err := validate(b.board, position)

	if err != nil {
		return board{}, err
	}

	// initialize 2 dimension slice
	newBoard := make([][]string, BOARD_SIZE)
	for i := range newBoard {
		newBoard[i] = make([]string, BOARD_SIZE)
	}

	for i, col := range b.board {
		for j, mark := range col {
			if strconv.Itoa(i*BOARD_SIZE+1+j) == position {
				newBoard[i][j] = playerMark
			} else {
				newBoard[i][j] = mark
			}
		}
	}

	return board{board: newBoard, stats: nil}, nil
}

func (b board) calculateStats() stats {
	if b.stats != nil {
		return *b.stats
	}

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

	_, mark := isComplete(lines)
	hasEmpty := hasEmptySpot(lines)

	b.stats = &stats{
		winner:   mark,
		hasEmpty: hasEmpty,
	}

	return *b.stats
}

func hasEmptySpot(b [][]string) bool {
	hasEmpty := false
	for _, columns := range b {
		for _, mark := range columns {
			if hasEmpty {
				return hasEmpty
			}
			if isEmpty(mark) {
				hasEmpty = true
			}
		}
	}

	return hasEmpty
}

// checks if line has the same 3 marks
func isComplete(lines [][]string) (isCompleted bool, mark string) {
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
				mark = ls[i]
			}
		}
	}
	return isCompleted, mark
}

func isEmpty(s string) bool {
	return s == "-"
}
