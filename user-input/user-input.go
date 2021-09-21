package userinput

import (
	"top-down-tdd/abstractions"
	"top-down-tdd/utils/terminal"
)

type UserInputHandler struct {
	terminal  terminal.TerminalUtil
	presenter abstractions.Presenter
}

func NewUserInput(t terminal.TerminalUtil) abstractions.InputHandler {
	return UserInputHandler{t, t}
}

func (u UserInputHandler) GetUserInput(message string) string {
	u.presenter.Display(message)
	var input string
	u.terminal.GetInput(&input)

	return input
}
