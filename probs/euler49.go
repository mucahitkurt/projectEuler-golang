package probs

import (
	"strconv"
	"math"
)

func Euler49() string {
	n := 9999
	for i := 1488; i < n; i++ {
		if !IsPrime(i) {
			continue
		}

		strNum := strconv.Itoa(i)
		cmb := List{Values:make(map[string]bool)}
		CreateAllCombinations(strNum, "", &cmb)

		for k, _ := range cmb.Values {
			secondNum, _ := strconv.Atoi(k)
			if i == secondNum || !IsPrime(secondNum) {
				continue
			}

			dif := int(math.Abs(float64(i - secondNum)))
			thirdNum := 0
			strResultNum := ""
			if i > secondNum {
				thirdNum = i + dif
				strResultNum = strconv.Itoa(secondNum) + strconv.Itoa(i) + strconv.Itoa(thirdNum)
			} else {
				thirdNum = secondNum + dif
				strResultNum = strconv.Itoa(i) + strconv.Itoa(secondNum) + strconv.Itoa(thirdNum)
			}
			if cmb.Values[strconv.Itoa(thirdNum)] && IsPrime(thirdNum) {
				return strResultNum
			}
		}
	}

	return ""
}
