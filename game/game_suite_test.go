package game_test

// import (
// 	"testing"

// 	"top-down-tdd/abstractions/mocks"
// 	"top-down-tdd/board"
// 	"top-down-tdd/user"
// 	userinput "top-down-tdd/user-input"

// 	"github.com/golang/mock/gomock"
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )
//TODO: fix main first
// func TestGame(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Game Suite")
// }

// var _ = Describe("Game", func() {

// 	var (
// 		mockCtrl *gomock.Controller
// 	)

// 	BeforeEach(func() {
// 		mockCtrl = gomock.NewController(GinkgoT())
// 	})

// 	When("there is winner", func() {
// 		It("shows correct message", func() {
// 			defer mockCtrl.Finish()

// 			user := user.NewUser("user 1234")

// 			boardMock := mocks.NewMockBoard(mockCtrl)
// 			boardMock.EXPECT().HasWinner().Times(1).Return(true)

// 			inputMock := mocks.NewMockInputHandler(mockCtrl)
// 			inputMock.EXPECT().GetUserInput().Times(1).Return(
// 				userinput.NewUserInput(0, 0),
// 			)

// 			presenterMock := mocks.NewMockPresenter(mockCtrl)
// 			expectedBoard := board.NewBoard()
// 			presenterMock.EXPECT().ShowBoard(expectedBoard).Times(1).Return("xxxxxxx")
// 			presenterMock.EXPECT().ShowMessage(user).Times(1).Return("player 1 won!")

// 			Expect(0).To(Equal(0))
// 		})
// 	})

// })
