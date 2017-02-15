package probs

import (
	"strconv"
	"math"
)

func Euler40_2() byte {
	numOfDigits := make(map[int]int)
	numOfDigits[1] = 9
	numOfDigits[2] = 180
	numOfDigits[3] = 2700
	numOfDigits[4] = 36000
	numOfDigits[5] = 450000
	numOfDigits[6] = 5400000

	var mult byte = 1

	for i := 10; i <= 1000000; i *= 10 {
		j := 1
		temp := i
		for ; temp > numOfDigits[j]; j++ {
			temp -= numOfDigits[j]
		}

		rem := temp % j
		mod := temp / j

		var num float64  = 0
		if mod == 0 {
			num =  math.Pow10(j - 1)
		} else {
			num = math.Pow10(j - 1) + float64(mod - 1)
		}

		if rem > 0 && mod > 0 {
			num++
		}

		if rem == 0 {
			rem = j
		}

		digit := (strconv.Itoa(int(num))[rem - 1]) - 48
		mult = mult * digit
	}

	return mult
}
