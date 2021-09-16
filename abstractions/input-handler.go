package abstractions

//go:generate mockgen -destination=./mocks/mock-input-handler.go -package=mocks top-down-tdd/abstractions InputHandler

import (
	userinput "top-down-tdd/user-input"
)

type InputHandler interface {
	GetUserInput() userinput.UserInput
}
