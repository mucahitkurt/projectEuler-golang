package euler

import "math/big"

func Euler25() string {
	f1 := big.NewInt(1)
	f2 := big.NewInt(1)
	count := big.NewInt(2)

	for {
		temp := f2
		f2 = f1.Add(f1, f2)
		f1 = temp
		count = count.Add(count, big.NewInt(1))
		if len(f2.String()) == 1000 {
			break
		}
	}

	return count.String()
}
