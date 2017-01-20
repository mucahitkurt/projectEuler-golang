package probs

import (
	"strconv"
	"strings"
)

func Euler32() int {
	numbers := "123456789"
	pandigital := make(map[int]bool)

	for i := 1; i <= len(numbers); i++ {
		for j := i + 1; j <= len(numbers); j++ {
			a, _ := strconv.Atoi(numbers[0:i])
			b, _ := strconv.Atoi(numbers[i:j])
			c := a * b
			if isPandigital(c, numbers, j) {
				pandigital[c] = true
			}
		}
	}

	sum := 0
	for key, _ := range(pandigital) {
		sum += key
	}

	return sum
}

func findAllPandigitals(prefix string, digits map[int]bool, pandigital map[int]bool)  {
	for k, v := range(digits) {
		if !v {
			a := prefix + strconv.Itoa(k)
			digits[k] = true
			allBs := []int {}
			findAllB("",digits, allBs)
			addAllCrossProduct(a, allBs, pandigital, digits)
		}
	}

}

func addAllCrossProduct(a string, bs []int, pandigital map[int]bool, digits map[int]bool) {

}

func findAllB(prefix string, digits map[int]bool, allBs []int)   {
	for k, v := range(digits) {
		if !v {
			b := prefix + strconv.Itoa(k)
			num, _ := strconv.Atoi(b)
			allBs = append(allBs, num)
			digits[k] = true
			findAllB(b, digits, allBs)
			digits[k] = false
		}
	}
}

func isPandigital(product int, numbers string, start int) bool {
	productStr := strconv.Itoa(product)
	if len(productStr) != 9 - start {
		return false
	}

	for _, num := range(numbers[start:]) {
		if strings.Index(productStr, string(num)) < 0 {
			return false
		}
	}


	return true
}
