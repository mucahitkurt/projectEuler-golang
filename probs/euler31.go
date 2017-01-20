package probs

import "strconv"

func Euler31(n int) int  {
	coins := []int {1,2,5,10,20,50,100,200}
	//different alternatives can be used as a cache data structure such as 2D array
	memo := make(map[string]int)
	return findAllWays(coins,0, n, memo)
}

func findAllWays(coins []int, index int, rem int, memo map[string]int) int  {
	key := strconv.Itoa(index)
	key += strconv.Itoa(rem)

	ways, ok := memo[key]
	if ok {
		return ways
	}

	if rem == 0 {
		return 1
	}

	if index >= len(coins) {
		return 0
	}

	if rem < 0 {
		return 0
	}

	ways = 0
	for i := 0; ; i++ {
		newRem := rem - (i*coins[index])
		if newRem < 0 {
			break
		}
		ways += findAllWays(coins, index+1, newRem, memo)
	}

	memo[key] = ways
	return ways

}