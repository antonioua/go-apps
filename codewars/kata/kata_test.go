package kata

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example tests", func() {
	It("Three hoops is not enough", func() {
		Expect(HoopCount(3)).To(Equal("Keep at it until you get it"))
	})

	It("Twelve hoops is good", func() {
		Expect(HoopCount(12)).To(Equal("Great, now move on to tricks"))
	})
})
