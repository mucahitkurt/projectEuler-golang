package probs

import (
	"github.com/adonovan/gopl.io/ch11/word1"
	"strconv"
)

func Euler36() int {
	n := 1000000
	sum := 0
	for i := 1; i < n; i++ {
		tenBaseNum := strconv.Itoa(i)
		twoBaseNum := strconv.FormatInt(int64(i), 2)
		if word.IsPalindrome(tenBaseNum) && word.IsPalindrome(twoBaseNum) {
			sum += i
		}
	}
	return sum
}
