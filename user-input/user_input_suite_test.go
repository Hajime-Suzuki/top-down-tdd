package userinput_test

import (
	"testing"
	userinput "top-down-tdd/user-input"
	terminal "top-down-tdd/utils/terminal/mocks"

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
			terminal := terminal.NewMockTerminalUtil(mockCtrl)
			terminal.EXPECT().GetInput(&s).SetArg(0, value)
			terminal.EXPECT().Display("select position").Return()

			//when
			subject := userinput.NewUserInput(terminal)

			//then
			res := subject.GetUserInput("select position")
			Expect(res).To(Equal("message"))

		})

	})
})
