package probs

func Euler28() int {
	var ur, ul, dr, dl = 1,1,1,1
	var diff = 2
	var sum = 1
	var n = 1001

	for ; ur < n*n; {
		dr += diff
		diff += 2
		dl += diff
		diff += 2
		ul += diff
		diff += 2
		ur += diff
		diff += 2

		sum += dr + dl + ul + ur
	}

	return sum
}