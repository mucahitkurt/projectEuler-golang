package probs

import (
	"math/big"
)

func Euler29() int {
	set := make(map[string]bool)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			ba := big.NewInt(int64(a))
			bb := big.NewInt(int64(b))
			pow := ba.Exp(ba,bb, big.NewInt(0))
			set[pow.String()] = true
		}
	}
	return len(set)
}
