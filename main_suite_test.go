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

	It("when there is winner, show result message", func() {
		defer mockCtrl.Finish() //! if this is not called, test passed when mock is not called...

		//given
		m := mocks.NewMockGame(mockCtrl)

		m.EXPECT().InitGame().Times(1)
		m.EXPECT().IsOver().Times(1).Return(true)

		// if game is over, no more update
		m.EXPECT().SetMark().Times(0)
		m.EXPECT().ShowBoard().Times(0)

		m.EXPECT().ShowResultMessage().Times(1)

		// when
		Start(m)
	})

	It("when there is no winner yet, update board", func() {
		defer mockCtrl.Finish() //! if this is not called, test passed when mock is not called...

		//given
		m := mocks.NewMockGame(mockCtrl)

		m.EXPECT().InitGame().Times(1)

		gomock.InOrder(
			m.EXPECT().IsOver().Return(false),
			//* in order test to be finished, it must return true
			m.EXPECT().IsOver().Return(true),
		)

		// if game is over, update
		setMark := m.EXPECT().SetMark().Times(1)
		showBoard := m.EXPECT().ShowBoard().Times(1)

		// if game is over, show result. but make sure it's after update operation
		m.EXPECT().
			ShowResultMessage().
			Times(1).
			After(setMark).
			After(showBoard)

		// when
		Start(m)
	})
})
