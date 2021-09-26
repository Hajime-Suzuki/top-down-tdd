package board

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func validate(board [][]string, position string) error {
	p, _ := strconv.Atoi(position)
	pf := float64(p)
	if pf > BOARD_SIZE*BOARD_SIZE {
		return errors.New("position is too big")
	}

	row := math.Floor(pf / BOARD_SIZE)
	col := math.Mod(pf, row*BOARD_SIZE)

	fmt.Println(row, col)
	markInPos := board[int(row)][int(col)]

	if !isEmpty(markInPos) {
		return errors.New(fmt.Sprintf("the position %s is not empty", markInPos))
	}

	return nil
}
