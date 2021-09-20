package players

import "top-down-tdd/abstractions"

type Players struct {
}

func NewPlayers() abstractions.Players {
	return Players{}
}

func (ps Players) GetCurrentPlayer() abstractions.Player {
	return Player{}
}

func (ps Players) Next() abstractions.Players {
	return Players{}
}

func (ps Players) RegisterNewPlayer(name string) {
	panic("TODO IMPLEMENT")
}

func (ps Players) GetPlayers() []abstractions.Player {
	panic("TODO IMPLEMENT")
}

type Player struct {
	name string
}

func NewPlayer(name string) Player {
	return Player{name}
}

func (p Player) ShowName() string {
	return p.name
}

func (p Player) GetMark() string {
	return "TODO: IMPLEMENT"
}
