package board

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func validate(board [][]string, position string) error {
	p, _ := strconv.Atoi(position)
	pos := float64(p)
	if pos > BOARD_SIZE*BOARD_SIZE {
		return errors.New("Position is too big")
	}

	row := math.Floor((pos - 1) / BOARD_SIZE)
	col := math.Mod((pos - 1), row*BOARD_SIZE)

	fmt.Println(row, col)
	targetMark := board[int(row)][int(col)]

	if !isEmpty(targetMark) {
		return errors.New(fmt.Sprintf("Position %s is not empty", position))
	}

	return nil
}
