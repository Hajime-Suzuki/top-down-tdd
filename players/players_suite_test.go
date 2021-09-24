package players

import (
	"testing"
	"top-down-tdd/abstractions"

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

	Context("Players", func() {
		Context("GetCurrentPlayer", func() {
			It("return current player", func() {
				// given
				ps := NewPlayers()
				ps.RegisterNewPlayer("p1")
				ps.RegisterNewPlayer("p2")

				// when
				subject := ps.GetCurrentPlayer()

				// then
				Expect(subject.ShowName()).To(Equal("p1"))
			})
		})

		Context("Next", func() {
			It("make the next player current", func() {
				// given
				ps := NewPlayers()
				ps.RegisterNewPlayer("p1")
				ps.RegisterNewPlayer("p2")

				// when
				updated := ps.Next()
				subject := updated.GetCurrentPlayer()

				// then
				Expect(subject.ShowName()).To(Equal("p2"))
			})

			It("current player is the same player if Next is called twice", func() {
				// given
				ps := NewPlayers()
				ps.RegisterNewPlayer("p1")
				ps.RegisterNewPlayer("p2")

				// when
				updated := ps.Next()
				updated = updated.Next()
				subject := updated.GetCurrentPlayer()

				// then
				Expect(subject.ShowName()).To(Equal("p1"))
			})
		})

		Context("RegisterNewPlayer", func() {
			It("add new player", func() {
				// given
				ps := NewPlayers()
				ps.RegisterNewPlayer("p1")
				ps.RegisterNewPlayer("p2")

				// when
				subject := ps.GetPlayers()

				// then
				expected := []abstractions.Player{
					NewPlayer("p1", "o"),
					NewPlayer("p2", "x"),
				}

				Expect(subject).To(Equal(expected))
			})
		})

		Context("GetPlayerByMark", func() {
			It("get 'o' player", func() {
				// given
				ps := NewPlayers()
				ps.RegisterNewPlayer("p1")
				ps.RegisterNewPlayer("p2")

				// when
				subject := ps.GetPlayerByMark("o")

				// then
				Expect(subject.ShowName()).To(Equal("p1"))
				Expect(subject.GetMark()).To(Equal("o"))
			})

			It("get 'x' player", func() {
				// given
				ps := NewPlayers()
				ps.RegisterNewPlayer("p1")
				ps.RegisterNewPlayer("p2")

				// when
				subject := ps.GetPlayerByMark("x")

				// then
				Expect(subject.ShowName()).To(Equal("p2"))
				Expect(subject.GetMark()).To(Equal("x"))
			})
		})
	})
})
