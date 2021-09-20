package abstractions

//go:generate mockgen -destination=./mocks/mock-players.go -package=mocks top-down-tdd/abstractions Players,Player

type Players interface {
	GetCurrentPlayer() Player
	Next() Players
	RegisterNewPlayer(name string)
}

type Player interface {
	ShowName() string
	GetMark() string
}
