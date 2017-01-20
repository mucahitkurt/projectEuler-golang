package probs

import (
	"strconv"
)

func Euler34() int {
	memo := [10]int{}
	n := 9999999
	sum := 0
	for i := 3; i < n; i++ {
		if curiousNumber(i, memo) {
			sum += i
		}
	}

	return sum
}

func curiousNumber(i int, memo [10]int) bool {
	iStr := strconv.Itoa(i)
	sum := 0
	for _, v := range iStr {
		vNum, _ := strconv.Atoi(string(v))
		if memo[vNum] <= 0 {
			memo[vNum] = factorial(vNum)
		}
		sum += memo[vNum]
	}

	return sum == i
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * factorial(i - 1)
}
