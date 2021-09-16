package abstractions

//go:generate mockgen -destination=./mocks/mock-presenter.go -package=mocks top-down-tdd/abstractions Presenter

type Presenter interface {
	ShowBoard(board Board) string
	ShowMessage(user User) string
}
