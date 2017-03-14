package probs

import "math"

func Euler46() int {
	for i := 35; i < 10000; i += 2 {
		if IsPrime(i) {
			continue
		}
		found := false
		for j := 1; j < int(math.Sqrt(float64(i))); j++ {
			k := i - 2*(int(math.Pow(float64(j), 2)))
			if IsPrime(k) {
				found = true
				break
			}
		}

		if !found {
			return i
		}
	}
	return 0
}
