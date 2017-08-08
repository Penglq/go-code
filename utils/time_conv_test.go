package utils_test

import (
	. "go-code/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TimeConv", func() {
	var (
		timestamp []int64
		want      []string
	)

	BeforeEach(func() {
		By("before test") // 会打印出步骤
		timestamp = append(timestamp, 1501814315)
		want = append(want, `2017-08-04 10:38:35`)
	})

	Context("timestamp to dateString", func() {
		It("timestamp", func() {
			By("Entering 0") // 会打印出步骤
			Expect(GetDateTimeFromTime(timestamp[0])).To(Equal(want[0]))

		})

		Measure("it should do something hard efficiently", func(b Benchmarker) {
			runtime := b.Time("runtime", func() {
				Expect(GetDateTimeFromTime(timestamp[0])).To(Equal(want[0]))
			})

			Expect(runtime.Seconds()).To(BeNumerically("<", 0), "shouldn't take too long.")

			Ω(runtime.Seconds()).Should(BeNumerically("<", 0.2), "shouldn't take too long.")

			// 只是自己预期的值
			b.RecordValue("disk usage (in MB)", 1)
		}, 10)
	})
})
