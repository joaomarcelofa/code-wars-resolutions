package movezeros

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMoveZeros(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MoveZeros Suite")
}

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})).To(Equal([]int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0}))
	})
})
