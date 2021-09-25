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

	Context("IsOver", func() {
		When("Game is not over", func() {
			It("return false 1", func() {

				// given
				data := [][]string{
					{"-", "-", "-"},
					{"-", "-", "-"},
					{"-", "-", "-"},
				}
				b := board{data}

				// when
				subject := b.IsOver()

				// then
				Expect(subject).To(Equal(false))
			})
			It("return false 2", func() {

				// given
				data := [][]string{
					{"x", "x", "o"},
					{"x", "-", "-"},
					{"o", "o", "x"},
				}
				b := board{data}

				// when
				subject := b.IsOver()

				// then
				Expect(subject).To(Equal(false))
			})
		})
		When("Game is over", func() {
			When("horizontal", func() {
				It("return true if there are 3 marks horizontally (first row)", func() {
					// given
					data := [][]string{
						{"x", "x", "x"},
						{"-", "o", "-"},
						{"-", "-", "x"},
					}
					b := board{data}

					// when
					subject := b.IsOver()

					// then
					Expect(subject).To(Equal(true))
				})
				It("return true if there are 3 marks horizontally (second row)", func() {
					// given
					data := [][]string{
						{"o", "-", "x"},
						{"o", "o", "o"},
						{"-", "-", "x"},
					}
					b := board{data}

					// when
					subject := b.IsOver()

					// then
					Expect(subject).To(Equal(true))
				})
				It("return true if there are 3 marks horizontally (third row)", func() {
					// given
					data := [][]string{
						{"o", "-", "x"},
						{"-", "o", "o"},
						{"x", "x", "x"},
					}
					b := board{data}

					// when
					subject := b.IsOver()

					// then
					Expect(subject).To(Equal(true))
				})
			})

			When("vertical", func() {

				It("return true if there are 3 marks (first column)", func() {

					// given
					data := [][]string{
						{"o", "x", "x"},
						{"o", "x", "-"},
						{"o", "-", "x"},
					}
					b := board{data}

					// when
					subject := b.IsOver()

					// then
					Expect(subject).To(Equal(true))
				})

				It("return true if there are 3 marks (second column)", func() {
					// given
					data := [][]string{
						{"-", "x", "o"},
						{"-", "x", "o"},
						{"-", "x", "-"},
					}
					b := board{data}

					// when
					subject := b.IsOver()

					// then
					Expect(subject).To(Equal(true))
				})

				It("return true if there are 3 marks horizontally (third column)", func() {
					// given
					data := [][]string{
						{"o", "-", "x"},
						{"-", "-", "x"},
						{"x", "o", "x"},
					}
					b := board{data}

					// when
					subject := b.IsOver()

					// then
					Expect(subject).To(Equal(true))
				})

			})
		})

	})
})
