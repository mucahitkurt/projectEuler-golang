package probs

import (
	"math"
)

func Euler27() int {

	//n2+an+b where |a|<1000 and |b|≤1000
	var maxPrimeNumber = 0
	var lastA, lastB = 0, 0
	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			var n = 0
			for ; ; n++ {
				if !IsPrime((n * n) + a * n + b) {
					break
				}
			}
			if n > maxPrimeNumber {
				maxPrimeNumber = n
				lastA = a
				lastB = b
			}
		}
	}

	return lastA * lastB
}

func IsPrime(n int) bool {
	if n  < 0 {
		return false
	}
	if n == 1 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}
