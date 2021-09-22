package presenter

import (
	"fmt"
	"top-down-tdd/abstractions"
)

type TerminalPresenter struct {
}

func NewPresenter() abstractions.Presenter {
	return TerminalPresenter{}
}

func (p TerminalPresenter) Display(message string) {
	if _, err := fmt.Println(message); err != nil {
		panic(err)
	}
}
