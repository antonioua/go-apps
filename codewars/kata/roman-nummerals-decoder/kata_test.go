package roman_nummerals_decoder_test

import (
	. "github.com/antonioua/codewars/kata/roman-nummerals-decoder"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test roman to decimal converter", func() {
	It("should give decimal number from roman", func() {
		Expect(Decode("XXI")).To(Equal(21))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("I")).To(Equal(1))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("IV")).To(Equal(4))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("MMVIII")).To(Equal(2008))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("MDCLXVI")).To(Equal(1666))
	})
})
