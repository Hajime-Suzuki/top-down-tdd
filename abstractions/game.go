package abstractions

//go:generate mockgen -destination=./mocks/mock-game.go -package=mocks top-down-tdd/abstractions Game

type Game interface {
	InitGame()
	SetMark()
	IsOver() bool
	ShowBoard()
	ShowResultMessage()
}
