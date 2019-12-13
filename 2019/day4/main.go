package main

import "fmt"

func main() {

	count := 0
	for i := 382345; i <= 843167; i++ {
		// for i := 382345; i <= 382395; i++ {
		if MeetsRules(i) {
			count++
		}
	}

	fmt.Println(count)
}

// MeetsRules TODO godoc
func MeetsRules(i int) bool {

	digits := make([]int, 6)

	digits[5] = i % 10
	digits[4] = (i / 10) % 10
	digits[3] = (i / 100) % 10
	digits[2] = (i / 1000) % 10
	digits[1] = (i / 10000) % 10
	digits[0] = (i / 100000) % 10

	// Two adjacent digits must be the same
	if (digits[0] == digits[1] && digits[1] != digits[2]) ||
		(digits[1] == digits[2] && digits[0] != digits[1] && digits[2] != digits[3]) ||
		(digits[2] == digits[3] && digits[1] != digits[2] && digits[3] != digits[4]) ||
		(digits[3] == digits[4] && digits[2] != digits[3] && digits[4] != digits[5]) ||
		(digits[4] == digits[5] && digits[3] != digits[4]) {

		// Values must not decrease
		if digits[0] <= digits[1] &&
			digits[1] <= digits[2] &&
			digits[2] <= digits[3] &&
			digits[3] <= digits[4] &&
			digits[4] <= digits[5] {

			return true
		}
	}

	return false
}
