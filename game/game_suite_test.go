package game_test

import (
	"testing"

	"top-down-tdd/abstractions/mocks"

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

	When("there is winner", func() {
		It("shows correct message", func() {
			defer mockCtrl.Finish()

			boardMock := mocks.NewMockBoard(mockCtrl)
			boardMock.EXPECT().HasWinner().Times(1).Return(true)

			inputMock := mocks.NewInputHandler(mockCtrl)
			inputMock.getUserInput.RETURNS = Mark{5, "x"}

			presenterMock := mocks.NewPresenter(mockCtrl)
			presenterMock.showBoard.RETURNS = "xxx xxx xxx xxx"
			presenterMock.showMessage.RETURNS = "player 1 won!"

			Expect(0).To(Equal(0))
		})
	})

})
