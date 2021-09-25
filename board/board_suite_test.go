package board

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBoard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Board Suite")
}

var _ = Describe("Board", func() {

	Context("Show", func() {
		It("returns representation of board", func() {
			// given
			data := [][]string{
				{"x", "-", "-"},
				{"-", "o", "-"},
				{"-", "-", "x"},
			}
			b := board{data}

			// when
			subject := b.Show()

			// then
			Expect(subject).To(Equal(`x 2 3 
4 o 6 
7 8 x 
`))
		})
	})
})
