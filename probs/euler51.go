package probs

import (
	"fmt"
	"sort"
	"strconv"
)

const THRESHOLD = 8

type list []int

type Positions struct {
	positions []list
}

func (p *Positions) addIndex(list list) {
	p.positions = append(p.positions, list)
}

func Euler51() int {

	numStart := 3099999
	numEnd := 10000000

	for i := numStart; i < numEnd; i++ {
		if ln := leastNumberForConsecutivePrimeNumbers(i); ln != -1 {
			return ln
		}
	}
	return -1
}

func leastNumberForConsecutivePrimeNumbers(num int) int {
	strNum := strconv.Itoa(num)
	var strPos string
	for i := 1; i < len(strNum); i++ {
		strPos += fmt.Sprintf("%v", i)
	}
	for i := 1; i <= len(strNum); i++ {
		positions := &Positions{}
		positionCombinations(strPos, "", i, positions)
		for _, pComb := range positions.positions {
			numbers := generateNumbers(num, pComb)
			primes := make([]int, 0)
			for k, v := range numbers {
				if k - len(primes) >= 2 {
					break
				}
				if IsPrime(v) {
					primes = append(primes, v)
				}
				if len(primes) == THRESHOLD {
					sort.Ints(primes)
					fmt.Println(primes)
					return primes[0]
				}
			}
		}
	}
	return -1
}

func positionCombinations(str string, pre string, pCount int, positions *Positions) {
	if len(pre) == pCount {
		postArr := make([]int, 0)
		for _, v := range pre {
			i, _ := strconv.Atoi(string(v))
			postArr = append(postArr, i)
		}
		positions.addIndex(postArr)
		return
	}

	if len(str) == 0 {
		return
	}

	positionCombinations(str[1:], pre+str[0:1], pCount, positions)
	positionCombinations(str[1:], pre+str[0:0], pCount, positions)
}

func generateNumbers(num int, positions []int) []int {
	nums := make([]int, 0)
	for i := 0; i <= 9; i++ {
		strNum := strconv.Itoa(num)
		for _, p := range positions {
			strNum = strNum[:p] + strconv.Itoa(i) + strNum[p+1:]
		}
		newNum, _ := strconv.Atoi(strNum)
		nums = append(nums, newNum)
	}

	return nums
}
