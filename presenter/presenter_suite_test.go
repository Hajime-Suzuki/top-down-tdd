package presenter_test

import (
	"testing"
	"top-down-tdd/presenter"
	terminalMock "top-down-tdd/utils/terminal/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPresenter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Presenter Suite")
}

func newCtrl() *gomock.Controller {
	return gomock.NewController(GinkgoT())
}

var _ = Describe("Presenter", func() {
	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = newCtrl()
	})

	Context("Display", func() {
		It("display message", func() {
			defer mockCtrl.Finish()
			//given

			t := terminalMock.NewMockTerminalUtil(mockCtrl)
			t.EXPECT().Display("test message").Return()

			//when
			subject := presenter.NewPresenter(t)

			//then
			subject.Display("test message")
		})

	})
})
