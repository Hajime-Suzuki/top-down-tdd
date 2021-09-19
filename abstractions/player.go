package abstractions

//go:generate mockgen -destination=./mocks/mock-player.go -package=mocks top-down-tdd/abstractions Player

type Player interface {
	ShowName() string
	GetMark() string
}
