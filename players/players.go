package players

import (
	"top-down-tdd/abstractions"
)

type Players struct {
	players []abstractions.Player
	marks   []string
	markIdx int
}

func NewPlayers() abstractions.Players {
	return &Players{
		players: []abstractions.Player{},
		marks:   []string{"o", "x"},
		markIdx: 0,
	}
}

func (ps *Players) GetCurrentPlayer() abstractions.Player {
	return ps.players[0]
}

func (ps *Players) Next() abstractions.Players {
	updated := append(ps.players[1:], ps.players[0])
	return &Players{
		players: updated,
		marks:   ps.marks,
		markIdx: ps.markIdx,
	}
}

func (ps *Players) RegisterNewPlayer(name string) {
	// TODO: fail when there are more than 2 players
	playerMark := ps.marks[ps.markIdx]
	ps.players = append(ps.players, NewPlayer(name, playerMark))
	ps.markIdx += 1
}

func (ps *Players) GetPlayerByMark(mark string) abstractions.Player {
	var player abstractions.Player

	for _, p := range ps.players {
		if p.GetMark() == mark {
			player = p
		}
	}

	return player
}

// used for testing
func (ps *Players) GetPlayers() []abstractions.Player {
	return ps.players
}
