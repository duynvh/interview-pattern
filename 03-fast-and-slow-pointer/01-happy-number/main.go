package main

import (
	"fmt"
	"strings"
)

func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res *= digit
	}
	return res
}

func sumDigits(number int) int {
	sum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		sum += pow(digit, 2)
	}

	return sum
}

// func isHappy(n int) bool { // Naive approach
// 	sum := sumDigits(n)
// 	mapSum := make(map[int]int)
// 	for sum != 1 && sum != n {
// 		sum = sumDigits(sum)
// 		if _, ok := mapSum[sum]; ok {
// 			break
// 		}
// 		mapSum[sum] = 1
// 	}

// 	return sum == 1
// }

func isHappy(n int) bool {
	slowPointer, fastPointer := n, sumDigits(n)

	for fastPointer != 1 && slowPointer != fastPointer {
		slowPointer = sumDigits(slowPointer)

		fastPointer = sumDigits(sumDigits(fastPointer))
	}

	return fastPointer == 1
}

func main() {
	array := []int{1, 5, 19, 25, 7}
	for i, v := range array {
		fmt.Printf("%d.\tInput number: %d\n", i+1, v)
		fmt.Printf("\tIs it a happy number? %t\n", isHappy(v))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
