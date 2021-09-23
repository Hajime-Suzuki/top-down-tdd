package players

import "top-down-tdd/abstractions"

type Players struct {
	players []abstractions.Player
}

func NewPlayers() abstractions.Players {
	return &Players{}
}

func (ps *Players) GetCurrentPlayer() abstractions.Player {
	return ps.players[0]
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

func (ps *Players) GetPlayerByMark(mark string) abstractions.Player {
	return Player{}
}
