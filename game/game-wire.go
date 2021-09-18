//+build wireinject

package game

import (
	"github.com/google/wire"
	. "top-down-tdd/user-input"
)

func InitializeGame() (Game, error) {
	wire.Build(NewGame, NewUserInput)
	return Game{}, nil
}
