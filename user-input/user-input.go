package userinput

//go:generate mockgen -destination=./mocks/terminal.go -package=mocks top-down-tdd/user-input TerminalUtil

import (
	"fmt"
	"top-down-tdd/abstractions"
)

type UserInputHandler struct {
	terminal  TerminalUtil
	presenter abstractions.Presenter
}

func NewUserInput(t TerminalUtil, p abstractions.Presenter) abstractions.InputHandler {
	return UserInputHandler{t, p}
}

func (u UserInputHandler) GetUserInput(message string) string {
	u.presenter.Dispay(message)
	var input string
	_, err := u.terminal.GetInput(&input)

	if err != nil {
		panic(err)
	}

	return input
}

// this is used for mock
type TerminalUtil interface {
	GetInput(*string) (int, error)
}

type terminalUtil struct {
}

func (t terminalUtil) GetInput(a *string) (n int, err error) {
	return fmt.Scanln(a)
}

func NewTerminalUtil() TerminalUtil {
	return terminalUtil{}
}
