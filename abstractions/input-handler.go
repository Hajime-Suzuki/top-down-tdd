package abstractions

//go:generate mockgen -destination=./mocks/mock-input-handler.go -package=mocks top-down-tdd/abstractions InputHandler

type InputHandler interface {
	GetUserInput(message string) (string, error)
}
