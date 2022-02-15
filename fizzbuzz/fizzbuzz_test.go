package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ONSDigital/blaise-go-fizzbuzz/fizzbuzz"
)

var _ = Describe("#FizzBuzz", func() {
	Context("when the number provided is divisible by 3", func() {
		It("returns 'Fizz'", func() {
			Expect(fizzbuzz.FizzBuzz(3)).To(Equal("Fizz"))
		})
	})

	Context("when the number provided is divisible by 5", func() {
		It("returns 'Buzz'", func() {
			Expect(fizzbuzz.FizzBuzz(5)).To(Equal("Buzz"))
		})
	})

	Context("when the number provided is divisible by both 3 and 5", func() {
		It("returns 'FizzBuzz'", func() {
			Expect(fizzbuzz.FizzBuzz(15)).To(Equal("FizzBuzz"))
		})
	})

	Context("when the number is not divisible by 3 or 5", func() {
		It("returns the number", func() {
			Expect(fizzbuzz.FizzBuzz(2)).To(Equal("2"))
		})
	})
})
