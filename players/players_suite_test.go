package players

import (
	"testing"
	"top-down-tdd/abstractions"
	"top-down-tdd/abstractions/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPlayers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Players Suite")
}

func newCtrl() *gomock.Controller {
	return gomock.NewController(GinkgoT())
}

var _ = Describe("Players", func() {

	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = newCtrl()
	})

	Context("Players", func() {
		Context("GetCurrentPlayer", func() {
			It("return current player", func() {
				// given
				p1 := mocks.NewMockPlayer(mockCtrl)
				p1.EXPECT().ShowName().Return("p1")
				p2 := mocks.NewMockPlayer(mockCtrl)
				ps := []abstractions.Player{
					p1, p2,
				}

				// when
				subject := Players{ps}

				// then
				res := subject.GetCurrentPlayer()
				Expect(res.ShowName()).To(Equal("p1"))
			})
		})
		Context("Next", func() {
			It("make the next player current", func() {
				// given
				p1 := mocks.NewMockPlayer(mockCtrl)
				p2 := mocks.NewMockPlayer(mockCtrl)
				p2.EXPECT().ShowName().Return("p2")

				ps := []abstractions.Player{
					p1, p2,
				}

				// when
				subject := Players{ps}

				// then
				updated := subject.Next()
				current := updated.GetCurrentPlayer()
				Expect(current.ShowName()).To(Equal("p2"))
			})
		})
	})

})
