package userinput_test

import (
	"testing"
	"top-down-tdd/abstractions/mocks"
	userinput "top-down-tdd/user-input"
	uMock "top-down-tdd/user-input/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserInput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserInput Suite")
}

func newCtrl() *gomock.Controller {
	return gomock.NewController(GinkgoT())
}

var _ = Describe("UserInput", func() {
	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = newCtrl()
	})

	Context("GetUserInput", func() {
		It("return user input", func() {
			defer mockCtrl.Finish()
			//given
			var s string
			value := "message"
			terminal := uMock.NewMockTerminalUtil(mockCtrl)
			terminal.EXPECT().GetInput(&s).SetArg(0, value)

			presenter := mocks.NewMockPresenter(mockCtrl)
			presenter.EXPECT().Dispay("select position").Return()

			//when
			userInput := userinput.NewUserInput(terminal, presenter)

			//then
			res := userInput.GetUserInput("select position")
			Expect(res).To(Equal("message"))

		})

	})
})
