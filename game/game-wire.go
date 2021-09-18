//+build wireinject

package game

import (
	"github.com/google/wire"
	"top-down-tdd/abstractions"
	. "top-down-tdd/user-input"
)

func InitializeGame() (abstractions.Game, error) {
	wire.Build(NewGame, NewUserInput)
	return game{}, nil
}
