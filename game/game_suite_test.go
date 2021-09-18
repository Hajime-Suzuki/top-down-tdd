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

			inputHandler.EXPECT().GetUserInput().Return("name 1", nil)

			inputHandler.EXPECT().GetUserInput().Return("name 2", nil)

			g := game{inputHandler: inputHandler}

			// when
			g.InitGame()

			// then
			b := board.NewBoard()
			p1 := player.NewPlayer("name 1")
			p2 := player.NewPlayer("name 2")
			ps := []abstractions.Player{p1, p2}

			Expect(g.board).To(Equal(b))
			Expect(g.players).To(Equal(ps))
		})
	})

})
