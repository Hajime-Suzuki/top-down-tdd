// +build wireinject

package game

import (
	"github.com/google/wire"
	"top-down-tdd/abstractions"
	"top-down-tdd/presenter"
	userinput "top-down-tdd/user-input"
)

func MakeGame() (abstractions.Game, error) {
	wire.Build(
		newGame,
		presenter.NewTerminalPresenter,
		userinput.NewTerminalUserInputHandler,
	)
	return &game{}, nil
}
