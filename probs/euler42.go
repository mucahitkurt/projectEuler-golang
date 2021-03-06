package probs

import (
	"strings"
	"math"
)

const (
	filepath = "C:/work/dev/projects/GoWorkspace/src/github.com/mucahitkurt/projectEuler-golang/probs/p042_words.txt"
)

func Euler42() int {

	strArr := ReadFileIntoArray(filepath, "\",\"")
	count := 0
	for _, str := range strArr {
		if isTriangleWord(str) {
			count++
		}
	}

	return count
}

func isTriangleWord(str string) bool {
	code := findWordCode(str)
	return IsTriangleNumber(int64(code))
}

func IsTriangleNumber(num int64) bool {
	sqrtCode := math.Sqrt(float64(num * 2))
	sqrtCodeInt := int64(sqrtCode)

	if ((sqrtCodeInt * (sqrtCodeInt + 1)) / 2) == num {
		return true
	}

	return false
}

func findWordCode(str string) int {
	str = strings.ToUpper(str)
	code := 0
	base := 'A'
	for _, chr := range str {
		code += int(chr - base) + 1
	}
	return code
}
