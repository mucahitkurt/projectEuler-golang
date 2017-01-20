package probs

import (
	"math"
	"strconv"
)

func Euler30() int {
	sum := 0
	//9^5 = 59049, 9^5*digitNumber, max number cannot be 7 digits
	for i := 2; i < 999999; i++ {
		if isSumOfFivePowerOfDigitsEqualToNumber(i) {
			sum += i
		}
	}

	return sum
 }

func isSumOfFivePowerOfDigitsEqualToNumber(i int) bool  {
	strNum := strconv.Itoa(i)
	sum := 0

	for i := 0; i < len(strNum); i++ {
		n, _ := strconv.ParseFloat(strNum[i:i+1], 64)
		sum += int(math.Pow(n, 5))
	}

	return sum == i

}