package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(number int) string {
	if number == 0 {
		return "Give me a number other than 0!"
	}
	if (number%3 == 0) && (number%5 == 0) {
		return "FizzBuzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}
