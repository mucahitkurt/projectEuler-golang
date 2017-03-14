package probs

func Euler45() int64 {
	from := 144
	to := 100000

	for i := from; i <= to; i++ {
		hex := getHexagon(i)
		if IsPentagonalNum(hex) && IsTriangleNumber(hex) {
			return hex
		}
	}

	return 0
}

func getHexagon(i int) int64 {
	return int64(i * ((2 * i) - 1))
}