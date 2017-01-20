package probs

import (
	"strconv"
	"strings"
	"sort"
	"fmt"
)

func Euler32() int {
	digits := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5 : false,
		6 : false,
		7 : false,
		8 : false,
		9 : false,
	}
	pandigital := make(map[int]bool)

	findAllPandigitals("", digits, pandigital)

	sum := 0
	fmt.Println(len(pandigital))
	for key, _ := range (pandigital) {
		sum += key
	}

	return sum
}

func findAllPandigitals(prefix string, digits map[int]bool, pandigital map[int]bool) {
	for k, v := range (digits) {
		if !v {
			a := prefix + strconv.Itoa(k)
			digits[k] = true
			findAllB(a, "", digits, pandigital)
			findAllPandigitals(a, digits, pandigital)
			digits[k] = false
		}
	}
}

func findAllB(aStr string, prefix string, digits map[int]bool, pandigital map[int]bool) {
	for k, v := range (digits) {
		if !v {
			bStr := prefix + strconv.Itoa(k)
			b, _ := strconv.Atoi(bStr)
			a, _ := strconv.Atoi(aStr)
			digits[k] = true

			if isPandigital(a*b, digits) {
				pandigital[a * b] = true
			}
			findAllB(aStr, bStr, digits, pandigital)
			digits[k] = false
		}
	}
}

func isPandigital(product int, digits map[int]bool) bool {
	strDigits := ""
	for k, v := range (digits) {
		if !v {
			strDigits += strconv.Itoa(k)
		}
	}

	productStr := strconv.Itoa(product)
	productStr = SortString(productStr)
	strDigits = SortString(strDigits)

	if len(strDigits) == len(productStr) && productStr == strDigits {
		return true
	}

	return false
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
