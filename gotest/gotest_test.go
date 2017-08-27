package gotest_test

import (
	. "go-code/gotest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gotest", func() {
	var (
		a    float64
		c    float64
		want float64
	)
	BeforeEach(func() {
		a = 4
		c = 2
		want = 2
	})

	Context("first", func() {
		It("it", func() {
			Expect(Division(a, c)).To(Equal(want))
		})

		Measure("it should do something hard efficiently", func(b Benchmarker) {
			runtime := b.Time("runtime", func() {
				Expect(Division(a, c)).To(Equal(want))
			})

			Expect(runtime.Seconds()).To(BeNumerically("<", 2), "shouldn't take too long.")

			Ω(runtime.Seconds()).Should(BeNumerically("<", 2), "shouldn't take too long.")

			// 只是自己预期的值
			b.RecordValue("disk usage (in MB)", 1)
		}, 10)
	})
})
