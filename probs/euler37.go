package probs

import (
	"strconv"
	"fmt"
)

func Euler37() int {
	s := 11
	n := 999999
	sum := 0
	count := 0
	for i := s; i < n && count < 11; i++ {
		if isAllPrime(i) {
			fmt.Println(i)
			count++
			sum += i
		}
	}

	return sum
}

func isAllPrime(num int) bool {
	if !IsPrime(num) {
		return false
	}

	strI := strconv.Itoa(num)

	for i := 1; i < len(strI); i++ {
		num, _ = strconv.Atoi(strI[i:])
		if !IsPrime(num) {
			return false
		}

		num, _ = strconv.Atoi(strI[:len(strI) - i])
		if !IsPrime(num) {
			return false
		}
	}

	return true
}
