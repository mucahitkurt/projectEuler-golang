package probs

import (
	"math"
	"fmt"
)

func Euler39() int {
	solutions := make(map[int]int)

	for p := 1000; p >= 5; p-- {
		for a := 1; a < 998; a++ {
			for b := 1; b <= p-a; b++ {
				a2 := math.Pow(float64(a), 2)
				b2 := math.Pow(float64(b), 2)
				c := p - (a+b)
				c2 := math.Pow(float64(c), 2)

				if (a2 + b2) == c2 {
					solutions[p] = solutions[p] + 1
				}
			}
		}
	}

	maxVal := 0
	maxP := 0
	for k, v := range solutions {
		if v > maxVal {
			maxVal = v
			maxP = k
		}
	}

	fmt.Println(solutions)

	return maxP
}
