package probs

import (
	"strconv"
	"strings"
)

type List struct {
	Values map[string]bool
}

func Euler41() int {
	strNum := "7654321"
	cmb := List{Values:make(map[string]bool)}
	CreateAllCombinations(strNum, "", &cmb)

	max := 0

	for i := 7; i > 0; i-- {
		for num, _ := range cmb.Values {
			i, _ := strconv.Atoi(num)
			if IsPrime(i) && i > max{
				max = i
			}
		}
		newValues := make(map[string]bool)
		for num, _ := range cmb.Values {
			index := strings.Index(num, strconv.Itoa(i))
			newVal := num[:index]
			if (index + 1) < len(num) {
				newVal += num[index + 1:]
			}
			if !newValues[newVal] {
				newValues[newVal] = true
			}
		}
		cmb.Values = newValues
	}

	return max
}

func CreateAllCombinations(str string, pre string, values *List) {
	if len(str) == 0 {
		values.Values[pre] = true
		return
	}

	for i := 0; i < len(str); i++ {
		rem := str[:i] + str[i + 1:]
		CreateAllCombinations(rem, pre+string(str[i]), values)
	}
}
