package board

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func validate(board [][]string, position string) error {
	p, err := strconv.Atoi(position)
	if err != nil {
		return errors.New("Position should be number")
	}

	pos := float64(p)
	if pos > BOARD_SIZE*BOARD_SIZE {
		return errors.New("Position is too big")
	}

	row := math.Floor((pos - 1) / BOARD_SIZE)
	col := math.Mod((pos - 1), math.Max(row, 1)*BOARD_SIZE)

	targetMark := board[int(row)][int(col)]

	if !isEmpty(targetMark) {
		return errors.New(fmt.Sprintf("Position %s is not empty", position))
	}

	return nil
}
