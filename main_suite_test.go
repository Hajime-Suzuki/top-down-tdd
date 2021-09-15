package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTopDownTdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TopDownTdd Suite")
}

var _ = Describe("Main", func() {
	It("runs game", func() {
		Expect(0).To(Equal(0))
	})
})
