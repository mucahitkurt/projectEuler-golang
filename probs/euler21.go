package probs

import (
	"fmt"
	"time"
)

func Euler21() int {
	s := time.Now()
	amicable := make(map[int]int)

	var res int = 0

	for i := 220; i < 10000; i++ {

		if amicable[i] != 0 {
			res += i
		} else {
			sum := FindSumOfProperDivisors2(i)
			if sum == i || sum == 0 {
				continue
			}
			sumA := FindSumOfProperDivisors2(sum)

			if sumA == i {
				res += i
				amicable[sumA] = sum
			}
		}
	}
	fmt.Println(time.Now().Sub(s))
	return res
}

func findSumOfProperDivisors(num int) int {
	sum := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum
}

func FindSumOfProperDivisors2(num int) int {
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if num/i != i {
				sum += num / i
			}
		}

	}
	return sum

}

/*
int sum_of_divisors(int num)
{
    int sum=1;
    for(int i=2; i*i<=num; sum+=num%i?0:(i*i==num?i:i+num/i) ,i++);
    return sum;
}

*/
