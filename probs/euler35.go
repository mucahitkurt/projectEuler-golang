package probs

import (
	"strconv"
)

func Euler35() int {
	n := 1000000
	count := 0
	for i := 2; i < n; i++ {
		if circularPrime(i) {
			count++
		}
	}
	return count
}
func circularPrime(i int) bool {
	iStr := strconv.Itoa(i)
	for i := 0; i < len(iStr); i++ {
		newNum, _ := strconv.Atoi(iStr[i:] + iStr[0:i])
		if !IsPrime(newNum) {
			return false
		}
	}

	return true
}
