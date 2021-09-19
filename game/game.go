package game

import (
	"errors"
	"fmt"
	"top-down-tdd/abstractions"
	"top-down-tdd/board"
	"top-down-tdd/player"
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

	userName1, err := g.inputHandler.GetUserInput("Player1: What's your name?")
	if err != nil {
		panic(err)
	}

	userName2, err := g.inputHandler.GetUserInput("Player2: What's your name?")

	if err != nil {
		panic(err)
	}

	g.players = []abstractions.Player{
		player.NewPlayer(userName1),
		player.NewPlayer(userName2),
	}
}

func (g *game) SetMark() {
	//TODO: HERE
	g.inputHandler.GetUserInput("")
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
