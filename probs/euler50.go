package probs

import (
	"math"
	"sort"
)

func Euler50() int {
	n := 1000000
	primes := Sieve(n)
	sort.Ints(primes)
	max := 0
	prime := 0
	for i := 0; i < len(primes); i++ {
		sum := primes[i]
		for j := i + 1; j < len(primes); j++ {
			sum += primes[j]
			if sum > n {
				break
			}
			if IsPrime(sum) && (j - i + 1) > max {
				max = j - i + 1
				prime = sum
			}
		}
	}

	return prime
}

func Sieve(num int) []int {
	nums := make(map[int]bool)
	for i := 2; i <= num; i++ {
		nums[i] = true
	}

	for i := 2; i < int(math.Sqrt(float64(num))); i++ {
		if nums[i] {
			for j := i * i; j <= num; j += i {
				nums[j] = false
			}
		}
	}

	primes := []int{}
	for k, _ := range nums {
		if nums[k] {
			primes = append(primes, k)
		}
	}

	return primes
}
