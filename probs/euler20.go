package probs

import "math/big"

func Euler20() int {

	a := big.NewInt(1)

	res := big.NewInt(1)

	for a.Int64() <= 100 {
		res.Mul(res, a)
		a.Add(a, big.NewInt(1))
	}

	strNum := res.String()

	sum := 0
	for _, i := range []byte(strNum) {
		sum += int(i - '0')
	}

	return sum
}
