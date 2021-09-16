package abstractions

import (
	userinput "top-down-tdd/user-input"
)

type InputHandler interface {
	GetUserInput() userinput.UserInput
}
