package probs

func Euler44() int {
	//pentagonal number Pi = n(3n-1)/2
	n := 5000
	pentagonals := generatePentagonal(n)

	for distance := 1; distance < n; distance++ {
		for i := 1; i+distance <= n; i++ {
			curDif := pentagonals[i + distance] - pentagonals[i]
			curSum := pentagonals[i + distance] + pentagonals[i]
			if IsPentagonal(curDif, pentagonals) && IsPentagonal(curSum, pentagonals) {
				return curDif
			}
		}
	}

	return 0
}

func IsPentagonal(num int, pentagonals map[int]int) bool {
	start := 1
	end := len(pentagonals)

	for start <= end {
		mid := (start + end) / 2
		if pentagonals[mid] == num {
			return true
		} else if pentagonals[mid] > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}

func generatePentagonal(n int) map[int]int {
	pentagonals := make(map[int]int)

	for i := 1; i <= n; i++ {
		pen := (i * ((3 * i) - 1)) / 2
		pentagonals[i] = pen
	}

	return pentagonals
}
