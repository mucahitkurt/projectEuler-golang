package euler

import (
	"fmt"
	"time"
)

func Euler23() int {

	//perfect number: sum of proper divisior equal the number itself
	//deficient: sum of proper division is less than n
	//abundant: sum of proper division is greater than n

	// 12 smallest abundant number
	// 24 smallest number that can be written as sum of two abundant numbers
	// all integers greater than 28123 can be written as sum of two abundant numbers
	// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
	start := time.Now()
	sum := 0

	var abondant [28123]bool

	for i := 1; i < 28123; i++ {
		if isAbondant(i) {
			abondant[i] = true
		}
	}

	for i := 1; i < 28123; i++ {
		if !isSumOfTwoAbondant(i, abondant) {
			sum += i
		}
	}
	fmt.Printf("Time elapsed:%v\n", time.Since(start))

	return sum
}

func isAbondant(num int) bool {
	if num < 12 {
		return false
	}

	sum := FindSumOfProperDivisors2(num)

	if sum > num {
		return true
	}

	return false
}

func isSumOfTwoAbondant(num int, abondant [28123]bool) bool {
	for a := 1; a <= num/2; a++ {
		if abondant[a] && abondant[num-a] {
			return true
		}
	}

	return false
}
