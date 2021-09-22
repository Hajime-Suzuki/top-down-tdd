package userinput

import (
	"fmt"
	"top-down-tdd/abstractions"
)

type TerminalUserInputHandler struct {
}

func NewUserInput() abstractions.InputHandler {
	return TerminalUserInputHandler{}
}

func (u TerminalUserInputHandler) GetUserInput() string {
	var input string
	if _, err := fmt.Scanln(input); err != nil {
		panic(err)
	}
	return input
}
