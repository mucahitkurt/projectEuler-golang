package probs

import (
	"strconv"
	"fmt"
)

func Euler33() float64 {
	var curiousFractions float64 = 1.0

	for i := 10.0; i < 99; i++ {
		for j := i+1; j <= 99; j++ {
			iStr := strconv.Itoa(int(i))
			jStr := strconv.Itoa(int(j))

			if iStr[1:2] != jStr[0:1] {
				continue
			}

			halfI, _ := strconv.Atoi(iStr[0:1])
			halfJ, _ := strconv.Atoi(jStr[1:2])

			if halfJ == 0 || halfI == 0 {
				continue
			}

			halfFI := float64(halfI)
			halfFJ := float64(halfJ)

			if i/j == halfFI/halfFJ {
				fmt.Printf("I:%f J:%f\n", i,j)
				curiousFractions *= i / j
			}
		}
	}

	return curiousFractions
}
