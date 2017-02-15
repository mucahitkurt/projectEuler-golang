package probs

import "strconv"

func Euler43() int {
	strPan := "9876543210"
	cmb := List{Values:make(map[string]bool)}
	CreateAllCombinations(strPan, "", &cmb)
	sum := 0
	for val, _ := range cmb.Values {
		if val[0] == '0' {
			continue
		}

		if isSubstringDivisible(val) {
			i, _ := strconv.Atoi(val)
			sum += i
		}

	}

	return sum
}

func isSubstringDivisible(val string) bool {
	divider := []int{2,3,5,7,11,13,17}
	start, end := 1,4

	for i := 0; i < len(divider);i++{
		num, _ := strconv.Atoi(val[start:end])
		if num%divider[i] != 0 {
			return false
		}
		start++
		end++
	}

	return true
}
