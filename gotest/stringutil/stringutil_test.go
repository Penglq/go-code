package stringutil_test

import (
	. "go-code/gotest/stringutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stringutil", func() {
	var (
		str  string
		want string
	)

	BeforeEach(func() {
		str = "abcdefg"
		want = "gfedcba"
	})

	Describe("string length", func() {
		Context("With more than 300 pages", func() {
			It("string", func() {
				//Ω(Reverse(str)).Should(Equal(want))
				Expect(Reverse(str)).To(Equal(want))
			})
		})
	})

})
