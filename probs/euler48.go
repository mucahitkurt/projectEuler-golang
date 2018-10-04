package probs

import "math/big"

func Euler48() string {
	n := 1000
	sum := big.NewInt(0)
	for i := 1; i <= n; i++ {
		x := big.NewInt(int64(i))
		x = x.Exp(x, x, nil)

		sum = x.Add(x, sum)
	}

	strNum := sum.String()

	return strNum[len(strNum)-10:]
}
