//+build wireinject

package game

import "github.com/google/wire"

func InitializeGame() (Game, error) {
	wire.Build(NewGame)
	return Game{}, nil
}
