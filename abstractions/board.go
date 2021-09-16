package abstractions

//go:generate mockgen -destination=./mocks/mock-board.go -package=mocks top-down-tdd/abstractions Board

type Board interface {
	HasWinner() bool
}
