package presenter

import (
	"top-down-tdd/abstractions"
	"top-down-tdd/utils/terminal"
)

// TODO: use terminal as presenter
type TerminalPresenter struct {
	terminal terminal.TerminalUtil
}

func NewPresenter(t terminal.TerminalUtil) abstractions.Presenter {
	return TerminalPresenter{t}
}

func (p TerminalPresenter) Display(msg string) {
	p.terminal.Display(msg)
}
