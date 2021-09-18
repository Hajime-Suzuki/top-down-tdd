package userinput

import "top-down-tdd/abstractions"

type UserInputHandler struct {
}

func NewUserInput() abstractions.InputHandler {
	return UserInputHandler{}
}

func (u UserInputHandler) GetUserInput() (string, error) {
	return "TODO: IMPLEMENT", nil
}
