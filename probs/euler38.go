package probs

import (
	"strconv"
	"sort"
	"strings"
)

func Euler38() string  {
	digits := "123456789"

	for i := 1; i < 9999; i++ {
		num := "";
		for j := 1; j < 5; j++ {
			if len(num) < 9 {
				num += strconv.Itoa(i * j)
			}
		}
		arr := strings.Split(num, "")
		sort.Strings(arr)
		numSorted := strings.Join(arr, "")

		if numSorted == digits {
			return num
		}
 	}


	return "0"
}