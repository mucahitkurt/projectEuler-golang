package probs

import "time"

func Euler19() int {
	start := time.Date(1901, time.January, 01, 0, 0, 0, 0, time.UTC)

	finish := time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC)

	countOfSunday := 0

	for start.Before(finish) {
		start = start.AddDate(0, 1, 0)
		if start.Weekday() == time.Sunday {
			countOfSunday++
		}
	}

	return countOfSunday
}
