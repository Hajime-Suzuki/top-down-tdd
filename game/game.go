package game

import (
	"errors"
	"fmt"
	"top-down-tdd/abstractions"
	"top-down-tdd/board"
	"top-down-tdd/players"
)

type game struct {
	board   abstractions.Board
	players []abstractions.Player

	inputHandler abstractions.InputHandler
	presenter    abstractions.Presenter
}

func NewGame(inputHandler abstractions.InputHandler) abstractions.Game {
	return &game{
		inputHandler: inputHandler,
	}
}

func (g *game) InitGame() {
	g.board = board.NewBoard()

	userName1 := g.inputHandler.GetUserInput("Player1: What's your name?")

	userName2 := g.inputHandler.GetUserInput("Player2: What's your name?")

	g.players = []abstractions.Player{
		players.NewPlayer(userName1),
		players.NewPlayer(userName2),
	}
}

func (g *game) SetMark() {
	p := g.players[0]

	input := g.inputHandler.GetUserInput(fmt.Sprintf(`
	%s, select position:
	%s
	`, p.ShowName(), g.board.Show()))

	mark := p.GetMark()
	updated, err := g.board.Update(mark, input)

	for err != nil {
		// when:
		// input is not number
		// input is not within available spot
		input := g.inputHandler.GetUserInput(fmt.Sprintf("%s. Try again:", err.Error()))

		updated, err = g.board.Update(mark, input)
	}

	g.board = updated
	//TODO: order players programatically in case there are more than 2 players
	g.players = []abstractions.Player{g.players[1], g.players[0]}
}

func (g *game) IsOver() bool {
	return g.board.IsOver()
}

func (g *game) ShowBoard() {
	b := g.board.Show()
	g.presenter.Dispay(b)
}

func (g *game) ShowResultMessage() error {
	if !g.board.IsOver() {
		return errors.New("Game is not over yet")
	}

	winnerMark := g.board.GetWinner()

	if winnerMark == "" {
		g.presenter.Dispay("Draw!")
		return nil
	}

	player := getPlayerByMark(g.players, winnerMark)

	msg := fmt.Sprintf("%s won!", player.ShowName())
	g.presenter.Dispay(msg)
	return nil
}

func getPlayerByMark(players []abstractions.Player, mark string) abstractions.Player {
	var player abstractions.Player

	for _, p := range players {
		if p.GetMark() == mark {
			player = p
		}
	}

	return player
}
