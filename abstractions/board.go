package abstractions

//go:generate mockgen -destination=./mocks/mock-board.go -package=mocks top-down-tdd/abstractions Board

type Board interface {
	IsOver() bool
	Show() string
	GetWinner() string // returns mark
	Update(mark string, position string) (Board, error)
}
