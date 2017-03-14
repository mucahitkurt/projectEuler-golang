package main

import (
	"fmt"
	"time"
	"github.com/mucahitkurt/projectEuler-golang/probs"
)

func main() {
	fmt.Printf("%d", probs.Euler47())

	//fmt.Println(math.MaxInt64)
	//numbers := "0123456789"
	//fmt.Println(numbers[0])
	//
	//for _, a := range(numbers[5:]) {
	//	fmt.Printf("%c", a)
	//}

	//fmt.Println(strconv.FormatInt(1, 2))
	//
	//str := "abcd"
	//fmt.Println(str[0:])
	//fmt.Println(str[1:])
	//fmt.Println(string(str[0]))

}

func releaseDate() {
	start := time.Date(2016, time.October, 12, 0, 0, 0, 0, time.UTC)

	//finish := time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC)

	countOfSunday := 25

	for countOfSunday > 0 {
		start = start.AddDate(0, 0, 14)
		countOfSunday--
		fmt.Println(start)
	}
}
