package presenter_test

import (
	"testing"
	"top-down-tdd/abstractions/mocks"
	userinput "top-down-tdd/user-input"

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

			outputWriter := mocks.NewMockOutputWriter(mockCtrl)
			outputWriter.EXPECT().Output("test message").Return()

			//when
			subject := userinput.NewPresenter(outputWriter)

			//then
			subject.Display("test message")
		})

	})
})
