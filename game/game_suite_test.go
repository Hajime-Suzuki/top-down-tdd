package game

import (
	"testing"
	"top-down-tdd/abstractions"
	"top-down-tdd/abstractions/mocks"
	"top-down-tdd/board"
	"top-down-tdd/player"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Game Suite")
}

var _ = Describe("Game", func() {

	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
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
			p1 := player.NewPlayer("name 1")
			p2 := player.NewPlayer("name 2")
			ps := []abstractions.Player{p1, p2}

			Expect(subject.board).To(Equal(b))
			Expect(subject.players).To(Equal(ps))
		})
	})

	Context("SetMark", func() {
		It("set mark if user input is valid", func() {
			//TODO: make players struct
			//******
			// get current player and its mark
			// ask player input
			// update board with mark and player input
			// update players' turn
			//******

			// given
			defer mockCtrl.Finish()

			currentPlayer := mocks.NewMockPlayer(mockCtrl)
			currentPlayer.EXPECT().GetMark().Return("o")
			currentPlayer.EXPECT().ShowName().MinTimes(1).Return("John Doe")

			nextPlayer := mocks.NewMockPlayer(mockCtrl)
			// used only for assertion. remove later
			nextPlayer.EXPECT().ShowName().AnyTimes().Return("user2")

			players := []abstractions.Player{
				currentPlayer,
				nextPlayer,
			}

			message := `
	John Doe, select position:
	1 2 3
	o 4 o
	x 5 x
	`

			inputHandler := mocks.NewMockInputHandler(mockCtrl)
			inputHandler.EXPECT().GetUserInput(message).Times(1).Return("3")

			board := mocks.NewMockBoard(mockCtrl)
			//TODO: find a better way to achieve this
			board.EXPECT().Show().Return(
				`1 2 3
	o 4 o
	x 5 x`,
			)

			updatedBoard := mocks.NewMockBoard(gomock.NewController(GinkgoT()))

			board.EXPECT().Update("o", "3").Times(1).Return(updatedBoard, nil)

			subject := game{
				inputHandler: inputHandler,
				players:      players,
				board:        board,
			}

			// when
			subject.SetMark()

			//then
			updatedPlayers := []abstractions.Player{
				nextPlayer,
				currentPlayer,
			}

			// fmt.Println(subject.players[0].ShowName())
			// fmt.Println(updatedPlayers[0].ShowName())

			// fmt.Println(subject.players[0] == updatedPlayers[0])
			// fmt.Println(subject.players[0] == updatedPlayers[1])

			//TODO: Expect doesn't work as expected when players are NOT swapped in implementation.
			Expect(subject.players[0]).To(Equal(updatedPlayers[0]))      //! Passes even if players are not swapped
			Expect(subject.players[0] == updatedPlayers[0]).To(BeTrue()) //! Fails as expected when players are not swapped

			Expect(subject.board).To(Equal(updatedBoard))
		})

		It("retry getting user input if invalid", func() {
			Skip("TODO: IMPLEMENT")
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

				player1 := mocks.NewMockPlayer(mockCtrl)
				player1.EXPECT().ShowName().Return(playerName)
				player1.EXPECT().GetMark().Return("o")

				player2 := mocks.NewMockPlayer(mockCtrl)
				player2.EXPECT().GetMark().Return("x")

				players := []abstractions.Player{
					player1,
					player2,
				}

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
