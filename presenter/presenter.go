package presenter

import (
	"top-down-tdd/abstractions"
	"top-down-tdd/utils/terminal"
)

type TerminalPresenter struct {
	terminal terminal.TerminalUtil
}

func NewPresenter() abstractions.Presenter {
	return TerminalPresenter{}
}

func (p TerminalPresenter) Display(msg string) {
	p.terminal.Display(msg)
}
