package euler

import "math/big"

func SimpleEuler16() int {
	a := big.NewInt(2)
	b := big.NewInt(2)

	for i := 999; i > 0; i-- {
		a = a.Mul(a, b)
	}

	strNum := a.String()

	sum := 0
	for _, i := range []byte(strNum) {
		sum += int(i - '0')
	}

	return sum
}

func Euler16() int {
	var a1 []int = []int{3, 2, 7, 3, 3, 9, 0, 6, 0, 7, 8, 9, 6, 1, 4, 1, 8, 7, 0, 0, 1, 3, 1, 8, 9, 6, 9, 6, 8, 2, 7, 5, 9, 9, 1, 5, 2, 2, 1, 6, 6, 4, 2, 0, 4, 6, 0, 4, 3, 0, 6, 4, 7, 8, 9, 4, 8, 3, 2, 9, 1, 3, 6, 8, 0, 9, 6, 1, 3, 3, 7, 9, 6, 4, 0, 4, 6, 7, 4, 5, 5, 4, 8, 8, 3, 2, 7, 0, 0, 9, 2, 3, 2, 5, 9, 0, 4, 1, 5, 7, 1, 5, 0, 8, 8, 6, 6, 8, 4, 1, 2, 7, 5, 6, 0, 0, 7, 1, 0, 0, 9, 2, 1, 7, 2, 5, 6, 5, 4, 5, 8, 8, 5, 3, 9, 3, 0, 5, 3, 3, 2, 8, 5, 2, 7, 5, 8, 9, 3, 7, 6}
	var a2 []int = []int{3, 2, 7, 3, 3, 9, 0, 6, 0, 7, 8, 9, 6, 1, 4, 1, 8, 7, 0, 0, 1, 3, 1, 8, 9, 6, 9, 6, 8, 2, 7, 5, 9, 9, 1, 5, 2, 2, 1, 6, 6, 4, 2, 0, 4, 6, 0, 4, 3, 0, 6, 4, 7, 8, 9, 4, 8, 3, 2, 9, 1, 3, 6, 8, 0, 9, 6, 1, 3, 3, 7, 9, 6, 4, 0, 4, 6, 7, 4, 5, 5, 4, 8, 8, 3, 2, 7, 0, 0, 9, 2, 3, 2, 5, 9, 0, 4, 1, 5, 7, 1, 5, 0, 8, 8, 6, 6, 8, 4, 1, 2, 7, 5, 6, 0, 0, 7, 1, 0, 0, 9, 2, 1, 7, 2, 5, 6, 5, 4, 5, 8, 8, 5, 3, 9, 3, 0, 5, 3, 3, 2, 8, 5, 2, 7, 5, 8, 9, 3, 7, 6}

	//var a1 []int = []int{1, 0, 0, 0}
	//var a2 []int = []int{1, 0, 0}

	sumA := timesOfArray(a1, a2)
	sum := 0
	//fmt.Printf("Result:%v\n", sumA)
	for _, i := range sumA {
		sum += i
	}

	return sum
}

type IntArr struct {
	data []int
}

func timesOfArray(a1, a2 []int) []int {
	var res []IntArr
	var level int = 0

	for i := (len(a1) - 1); i >= 0; i-- {
		subSet := IntArr{}
		rem := 0
		for j := (len(a2) - 1); j >= 0; j-- {
			temp := a1[i] * a2[j]
			temp += rem
			subSet.data = append(subSet.data, temp%10)
			rem = temp / 10
		}
		if rem != 0 {
			subSet.data = append(subSet.data, rem)
		}

		subSet.data = reverse(subSet.data)

		tempLevel := level
		for tempLevel > 0 {
			subSet.data = append(subSet.data, 0)
			tempLevel--
		}
		level++
		//fmt.Printf("Appending:%v\n", subSet)
		subSet.data = reverse(subSet.data)
		res = append(res, subSet)
	}

	//fmt.Printf("Result:%v\n", res)

	var resIntArr []int
	maxLen := 0
	i := 0
	rem := 0
	for true {
		sum := 0
		for _, intArr := range res {
			if maxLen < len(intArr.data) {
				maxLen = len(intArr.data)
			}
			if len(intArr.data) > i {
				sum += intArr.data[i]
			}
		}
		sum += rem
		resIntArr = append(resIntArr, sum%10)
		rem = sum / 10

		i++
		if i >= maxLen {
			break
		}
	}

	for rem != 0 {
		resIntArr = append(resIntArr, rem%10)
		rem = rem / 10
	}

	return reverse(resIntArr)
}

func reverse(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}
	//fmt.Printf("Before Reverse:%v\n", numbers)

	newNumbers := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		newNumbers[i] = numbers[(len(numbers) - 1 - i)]
	}

	//fmt.Printf("After Reverse:%v\n", newNumbers)
	return newNumbers
}
