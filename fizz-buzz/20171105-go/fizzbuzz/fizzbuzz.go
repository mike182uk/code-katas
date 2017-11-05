package fizzbuzz

import (
	"strconv"
	"strings"
)

// Translate translates a given integer to the correct fizzbuzz value
func Translate(val int) string {
	if val%15 == 0 {
		return "FizzBuzz"
	}

	if val%5 == 0 {
		return "Buzz"
	}

	if val%3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(val)
}

// Generate generates fizzbuzz until the given limit
func Generate(limit int) string {
	result := ""
	seperator := " "

	for i := 1; i <= limit; i++ {
		result += Translate(i) + seperator
	}

	return strings.Trim(result, seperator)
}
