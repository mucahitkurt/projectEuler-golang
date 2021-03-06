package probs

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const (
	fileName = "C:/work/dev/projects/GoWorkspace/src/github.com/mucahitkurt/sandbox/euler/p022_names.txt"
)

func Euler22() int {
	strArr := ReadFileIntoArray(fileName, "\",\"")

	sort.Strings(strArr)

	sum := 0
	for i, str := range strArr {
		sum += calculateValue(str, i+1)
	}

	return sum
}

func calculateValue(str string, i int) int {
	sumOfPos := 0
	for _, c := range str {
		sumOfPos += int((c - 'A') + 1)
	}
	return sumOfPos * i
}

func ReadFileIntoArray(path string, sep string) []string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Fault:%s", err.Error())
		os.Exit(1)
	}

	strFile := string(dat)
	var strArr []string
	strArr = strings.Split(strFile, sep)

	strArr[0] = strArr[0][1:]
	strLast := strArr[len(strArr) - 1]
	strLast = strLast[:len(strLast) - 1]
	strArr[len(strArr) - 1] = strLast

	return strArr
}
