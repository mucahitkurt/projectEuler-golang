package probs

func Euler31(n int) int  {
	coins := []int {1,2,5,10,20,50,100,200}

	return findAllWays(coins,0, n)
}

func findAllWays(coins []int, index int, rem int) int  {
	if rem == 0 {
		return 1
	}

	if index >= len(coins) {
		return 0
	}

	if rem < 0 {
		return 0
	}

	ways := 0
	for i := 0; ; i++ {
		newRem := rem - (i*coins[index])
		if newRem < 0 {
			break
		}
		ways += findAllWays(coins, index+1, newRem)
	}

	return ways

}
