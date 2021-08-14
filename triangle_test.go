package triangle_test

import (
	. "github.com/adelowo/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Triangle", func() {

	It("Should return an error if the sides don't make up a triangle", func() {

		got, err := KindFromSides(0, -1, 10)

		Expect(err).To(HaveOccurred())
		Expect(got).To(Equal(NaT))
	})
})
