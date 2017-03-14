package probs

import (
	"math"
)

func Euler47() int {
	n := 1000000
	numOfPrimeFactorsForEachNumber := 4
	memo := make(map[int][]int)

	for i := 647; i < n; i++ {
		primeFactorsSetForConsecutiveNumbers := make(map[int]bool)
		if isConsecutiveNumberThatHasDifferentPrimeFactorsFound(i, numOfPrimeFactorsForEachNumber, primeFactorsSetForConsecutiveNumbers, memo) {
			return i
		}
	}

	return 0
}

func isConsecutiveNumberThatHasDifferentPrimeFactorsFound(i int, numOfPrimeFactors int, primeFactorsSet map[int]bool, memo map[int][]int) bool {
	for j := 0; j < numOfPrimeFactors; j++ {
		if !isConsecutiveRuleValid(i+j, primeFactorsSet, numOfPrimeFactors, (j+1)*numOfPrimeFactors, memo) {
			return false
		}
	}
	return true
}

func isConsecutiveRuleValid(num int, primeFactorsSet map[int]bool, numOfPrimeFactors int, totalNumOfPrimeFactors int, memo map[int][]int) bool {
	var primeFactors []int
	if memo[num] != nil {
		primeFactors = memo[num]
	} else {
		primeFactors = PrimeFactors(num)
	}

	if len(primeFactors) != numOfPrimeFactors {
		return false
	}

	for i := 0; i < numOfPrimeFactors; i++ {
		primeFactorsSet[primeFactors[i]] = true
	}

	if len(primeFactorsSet) != totalNumOfPrimeFactors {
		return false
	}

	return true
}

func PrimeFactors(num int) []int {
	if IsPrime(num) {
		pf := [1]int{num}
		return pf[:]
	}
	pf := []int{}
	for i := 2; i < int(math.Sqrt(float64(num))); i ++ {
		if num%i == 0 {
			power, num := base(num, i)
			pf = append(pf, int(math.Pow(float64(i), float64(power))))
			pf = append(pf, PrimeFactors(num)...)
			break
		}
	}

	return pf[:]
}

func base(num, primeFactor int) (power int, base int) {
	power = 0
	for num != 0 && num%primeFactor == 0 {
		num = num / primeFactor
		power++
	}

	return power, num
}
