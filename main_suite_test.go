package main_test

import (
	"testing"
	. "top-down-tdd"
	"top-down-tdd/abstractions/mocks"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTopDownTdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TopDownTdd Suite")
}

var _ = Describe("Main", func() {
	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	It("when start game, game.startGame method is called", func() {
		defer mockCtrl.Finish() //! if this is not called, test passed when mock is not called...
		m := mocks.NewMockGame(mockCtrl)
		m.EXPECT().StartGame()
		Start(m)
	})
})
