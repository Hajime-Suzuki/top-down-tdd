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
	players abstractions.Players

	inputHandler abstractions.InputHandler
	presenter    abstractions.Presenter
}

func NewGame(
	inputHandler abstractions.InputHandler,
	presenter abstractions.Presenter,
) abstractions.Game {
	return &game{
		inputHandler: inputHandler,
		presenter:    presenter,
	}
}

func (g *game) InitGame() {
	g.board = board.NewBoard()
	g.players = players.NewPlayers()

	g.presenter.Display("Player1: What's your name?")
	userName1 := g.inputHandler.GetUserInput()

	g.presenter.Display("Player2: What's your name?")
	userName2 := g.inputHandler.GetUserInput()

	g.players.RegisterNewPlayer(userName1)
	g.players.RegisterNewPlayer(userName2)
}

func (g *game) SetMark() {
	p := g.players.GetCurrentPlayer()

	g.presenter.Display(fmt.Sprintf(`
	%s, select position:
	%s`, p.ShowName(), g.board.Show()))
	input := g.inputHandler.GetUserInput()

	mark := p.GetMark()
	updated, err := g.board.Update(mark, input)

	for err != nil {
		// when:
		// input is not number
		// input is not within available spot
		g.presenter.Display(fmt.Sprintf("%s. Try again:", err.Error()))
		input := g.inputHandler.GetUserInput()

		updated, err = g.board.Update(mark, input)
	}

	g.players = g.players.Next()
	g.board = updated
}

func (g *game) IsOver() bool {
	return g.board.IsOver()
}

func (g *game) ShowBoard() {
	b := g.board.Show()
	g.presenter.Display(b)
}

func (g *game) ShowResultMessage() error {
	if !g.board.IsOver() {
		return errors.New("Game is not over yet")
	}

	winnerMark := g.board.GetWinner()

	if winnerMark == "" {
		g.presenter.Display("Draw!")
		return nil
	}

	player := g.players.GetPlayerByMark(winnerMark)

	msg := fmt.Sprintf("%s won!", player.ShowName())
	g.presenter.Display(msg)
	return nil
}
