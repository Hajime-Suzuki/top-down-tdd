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

func makeDefaultPlayers(ps []abstractions.Player) Players {
	return Players{ps, []string{"o", "x"}, 0}
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
				subject := makeDefaultPlayers(ps)

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
				subject := makeDefaultPlayers(ps)

				// then
				updated := subject.Next()
				current := updated.GetCurrentPlayer()
				Expect(current.ShowName()).To(Equal("p2"))
			})
			It("current player is the same player if Next is called twice", func() {
				// given
				p1 := mocks.NewMockPlayer(mockCtrl)
				p1.EXPECT().ShowName().Return("p1")
				p2 := mocks.NewMockPlayer(mockCtrl)

				ps := []abstractions.Player{
					p1, p2,
				}

				// when
				subject := makeDefaultPlayers(ps)
				// then
				updated := subject.Next()
				updated = updated.Next()
				current := updated.GetCurrentPlayer()
				Expect(current.ShowName()).To(Equal("p1"))
			})
		})
		Context("RegisterNewPlayer", func() {
			It("add new player", func() {
				playerName1 := "p1"
				playreMark1 := "o"
				playerName2 := "p2"
				playerMark2 := "x"

				// when
				subject := NewPlayers()
				subject.RegisterNewPlayer(playerName1)
				subject.RegisterNewPlayer(playerName2)

				// then
				p1 := subject.GetPlayers()[0]
				p2 := subject.GetPlayers()[1]
				Expect(p1.ShowName()).To(Equal(playerName1))
				Expect(p1.GetMark()).To(Equal(playreMark1))
				Expect(p2.ShowName()).To(Equal(playerName2))
				Expect(p2.GetMark()).To(Equal(playerMark2))
			})
		})

	})

})
