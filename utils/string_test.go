package utils_test

import (
	. "go-code/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String", func() {
	var (
		str  string
		want string
	)

	BeforeEach(func() {
		By("before test") // 会打印出步骤
		str = "abc"
		want = "cbaa"
	})

	It("reverse string", func() {
		By("Entering") // 会打印出步骤
		Expect(ReverseString(str)).To(Equal(want))
	})
})
