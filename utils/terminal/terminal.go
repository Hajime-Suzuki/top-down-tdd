package terminal

//go:generate mockgen -destination=./mocks/terminal.go -package=mocks top-down-tdd/utils/terminal TerminalUtil

import "fmt"

type TerminalUtil interface {
	GetInput(*string)
	Dispay(string)
}

type terminalUtil struct {
}

func (t terminalUtil) GetInput(a *string) {
	if _, err := fmt.Scanln(a); err != nil {
		panic(err)
	}
}

func (t terminalUtil) Dispay(a string) {
	if _, err := fmt.Println(a); err != nil {
		panic(err)
	}
}

func NewTerminalUtil() TerminalUtil {
	return terminalUtil{}
}
