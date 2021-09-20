package game

import (
	"errors"
	"fmt"
	"testing"
	"top-down-tdd/abstractions"
	"top-down-tdd/abstractions/mocks"
	"top-down-tdd/board"
	"top-down-tdd/players"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Game Suite")
}

func newCtrl() *gomock.Controller {
	return gomock.NewController(GinkgoT())
}

var _ = Describe("Game", func() {

	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = newCtrl()
	})

	Context("InitGame", func() {
		It("create board and players", func() {
			//******
			// ask player names
			// create board
			// create players
			//******

			// given
			defer mockCtrl.Finish()

			inputHandler := mocks.NewMockInputHandler(mockCtrl)
			inputHandler.EXPECT().GetUserInput("Player1: What's your name?").Return("name 1")
			inputHandler.EXPECT().GetUserInput("Player2: What's your name?").Return("name 2")

			subject := game{inputHandler: inputHandler}

			// when
			subject.InitGame()

			// then
			b := board.NewBoard()
			p1 := players.NewPlayer("name 1")
			p2 := players.NewPlayer("name 2")
			ps := []abstractions.Player{p1, p2}

			Expect(subject.board).To(Equal(b))
			Expect(subject.players.GetPlayers()).To(Equal(ps))
		})
	})

	Context("SetMark", func() {
		It("set mark if user input is valid", func() {
			//TODO: make players struct
			//******
			// get current player and its mark
			// ask player input (show current board as well)
			// update board with mark and player input
			// update players' turn
			//******

			// given
			defer mockCtrl.Finish()

			currentPlayer := mocks.NewMockPlayer(mockCtrl)
			currentPlayer.EXPECT().GetMark().Return("o")
			currentPlayer.EXPECT().ShowName().Return("John Doe")

			players := mocks.NewMockPlayers(mockCtrl)
			players.EXPECT().GetCurrentPlayer().Return(currentPlayer)

			// message := gomega.ContainSubstring("John Doe")

			inputHandler := mocks.NewMockInputHandler(mockCtrl)
			//TODO: fix message and use it here
			inputHandler.EXPECT().GetUserInput(gomock.Any()).Return("3")

			board := mocks.NewMockBoard(mockCtrl)
			//TODO: find a better way to achieve this
			board.EXPECT().Show().Return(
				`1 2 3
    o 4 o
    x 5 x`)

			updatedBoard := mocks.NewMockBoard(newCtrl())

			board.EXPECT().Update("o", "3").Times(1).Return(updatedBoard, nil)

			nextPlayers := mocks.NewMockPlayers(newCtrl())
			players.EXPECT().Next().Return(nextPlayers)

			subject := game{
				inputHandler: inputHandler,
				players:      players,
				board:        board,
			}

			// when
			subject.SetMark()

			//then
			Expect(subject.board).To(Equal(updatedBoard))
			Expect(subject.players).To(Equal(nextPlayers))
		})

		It("retry getting user input if invalid", func() {
			// given
			defer mockCtrl.Finish()

			currentPlayer := mocks.NewMockPlayer(mockCtrl)
			currentPlayer.EXPECT().GetMark().AnyTimes().Return("o")
			currentPlayer.EXPECT().ShowName().AnyTimes().Return("John Doe")

			players := mocks.NewMockPlayers(mockCtrl)
			players.EXPECT().GetCurrentPlayer().AnyTimes().Return(currentPlayer)

			errorMessage := "invalid input"

			inputHandler := mocks.NewMockInputHandler(mockCtrl)
			inputHandler.EXPECT().GetUserInput(gomock.Any()).Return("3")

			// when there is error in update, it will retry
			inputHandler.EXPECT().GetUserInput(fmt.Sprintf("%s. Try again:", errorMessage)).Return("5")

			board := mocks.NewMockBoard(mockCtrl)
			board.EXPECT().Show().Return(
				`1 2 3
				o 4 o
				x 5 x`,
			)

			// throws error first time
			board.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, errors.New(errorMessage))

			updatedBoard := mocks.NewMockBoard(newCtrl())
			board.EXPECT().Update("o", "5").Return(updatedBoard, nil)

			nextPlayers := mocks.NewMockPlayers(newCtrl())
			players.EXPECT().Next().Return(nextPlayers)

			subject := game{
				inputHandler: inputHandler,
				players:      players,
				board:        board,
			}

			// when
			subject.SetMark()

			//then
			Expect(subject.board).To(Equal(updatedBoard))
		})
	})

	Context("IsOver", func() {
		It("true if game is over", func() {
			// given
			defer mockCtrl.Finish()

			board := mocks.NewMockBoard(mockCtrl)
			board.EXPECT().IsOver().Times(1).Return(true)

			subject := game{board: board}

			// when
			res := subject.IsOver()

			//then
			Expect(res).To(Equal(true))
		})

		It("false if game is not over", func() {
			// given
			defer mockCtrl.Finish()

			board := mocks.NewMockBoard(mockCtrl)
			board.EXPECT().IsOver().Times(1).Return(false)

			subject := game{board: board}

			// when
			res := subject.IsOver()

			//then
			Expect(res).To(Equal(false))
		})
	})

	Context("ShowBoard", func() {
		It("show board", func() {
			// given
			defer mockCtrl.Finish()

			boardStr := `
			- - -
			- o -
			- - -
			`

			board := mocks.NewMockBoard(mockCtrl)
			board.EXPECT().Show().Return(boardStr)

			presenter := mocks.NewMockPresenter(mockCtrl)
			presenter.EXPECT().Dispay(boardStr).Times(1)

			subject := game{board: board, presenter: presenter}

			// when
			subject.ShowBoard()
		})
	})

	Context("ShowResultMessage", func() {
		When("there is winner", func() {
			It("show a winner message", func() {
				// given
				defer mockCtrl.Finish()

				//********************************
				// 1. get winner mark from board
				// 2. get player by mark from players
				// 3. show message like "player 1234 won!"
				//********************************

				playerName := "John Doe"
				message := "John Doe won!"

				board := mocks.NewMockBoard(mockCtrl)
				board.EXPECT().IsOver().Return(true)
				board.EXPECT().GetWinner().Return("o")

				winner := mocks.NewMockPlayer(mockCtrl)
				winner.EXPECT().ShowName().Times(1).Return(playerName)

				players := mocks.NewMockPlayers(mockCtrl)
				players.EXPECT().GetPlayerByMark("o").Return(winner)

				presenter := mocks.NewMockPresenter(mockCtrl)
				presenter.EXPECT().Dispay(message).Times(1)

				subject := game{
					board:     board,
					players:   players,
					presenter: presenter,
				}

				// when
				subject.ShowResultMessage()
			})
		})

		When("there is no winner", func() {
			It("show draw message", func() {
				// given
				defer mockCtrl.Finish()

				//********************************
				// 1. get winner mark from board
				// 2. if there is no winner, show message like "Draw!"
				//********************************

				message := "Draw!"

				board := mocks.NewMockBoard(mockCtrl)
				board.EXPECT().IsOver().Return(true)
				board.EXPECT().GetWinner().Return("")

				presenter := mocks.NewMockPresenter(mockCtrl)
				presenter.EXPECT().Dispay(message).Times(1)

				subject := game{
					board:     board,
					presenter: presenter,
				}

				// when
				subject.ShowResultMessage()
			})

		})

		It("thrown error when game is not over yet", func() {
			// in case ShowResultMessage is called before game is over, error should be thrown

			// given
			defer mockCtrl.Finish()

			board := mocks.NewMockBoard(mockCtrl)
			board.EXPECT().IsOver().Return(false)

			presenter := mocks.NewMockPresenter(mockCtrl)
			presenter.EXPECT().Dispay(gomock.Any()).Times(0)

			subject := game{
				board:     board,
				presenter: presenter,
			}

			// when
			e := subject.ShowResultMessage()

			//then
			Expect(e.Error()).To(Equal("Game is not over yet"))
		})

	})
})
