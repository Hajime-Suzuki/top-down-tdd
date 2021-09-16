package abstractions

//go:generate mockgen -destination=./mocks/mock-user.go -package=mocks top-down-tdd/abstractions User

type User interface {
	ShowName() string
}
