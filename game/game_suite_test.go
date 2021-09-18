package game

import (
	"testing"
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
		//******
		// ask player names
		// create players
		// create board
		//******

		// given
		defer mockCtrl.Finish()
		userInputMock := mocks.NewMockInputHandler(mockCtrl)

		userInputMock.EXPECT().GetUserInput().Return("name 1")
		userInputMock.EXPECT().GetUserInput().Return("name 2")

		g := game{userInputMock}

		// when
		g.InitGame()

		// then
		b := board.NewBoard()
		p1 := player.NewPlayer("name 1")
		p2 := player.NewPlayer("name 2")
		ps := []player.Player{p1, p2}

		Expect(g.board).To(Equal(b))
		Expect(g.players).To(Equal(ps))
	})

})
