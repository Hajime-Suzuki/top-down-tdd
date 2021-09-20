package players

import "top-down-tdd/abstractions"

type Players struct {
	players []abstractions.Player
}

func NewPlayers() abstractions.Players {
	return &Players{}
}

func (ps *Players) GetCurrentPlayer() abstractions.Player {
	return Player{}
}

func (ps *Players) Next() abstractions.Players {
	return &Players{}
}

func (ps *Players) RegisterNewPlayer(name string) {
	ps.players = append(ps.players, NewPlayer(name))
}

// used for testing
func (ps *Players) GetPlayers() []abstractions.Player {
	return ps.players
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
