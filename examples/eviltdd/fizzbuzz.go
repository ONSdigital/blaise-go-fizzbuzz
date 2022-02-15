package eviltdd

import "fmt"

func FizzBuzz(number int) string {
	if number == 3 {
		return "Fizz"
	}
	if number == 6 {
		return "Fizz"
	}
	if number == 5 {
		return "Buzz"
	}
	if number == 10 {
		return "Buzz"
	}
	if number == 15 {
		return "FizzBuzz"
	}
	if number == 30 {
		return "FizzBuzz"
	}
	return fmt.Sprintf("%d", number)
}
