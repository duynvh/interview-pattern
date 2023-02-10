package main

import (
	"fmt"
	"strings"
)

func tribonacci(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		} else {
			return 1
		}
	}

	// Initializing first three numbers
	firstNum, secondNum, thirdNum := 0, 1, 1
	for i := 3; i <= n; i++ {
		firstNum, secondNum, thirdNum = secondNum, thirdNum, firstNum+secondNum+thirdNum
	}

	return thirdNum
}

// Driver code
func main() {
	inputs := []int{4, 5, 25, 17, 19}

	for index, value := range inputs {
		tr := tribonacci(value)
		fmt.Printf("%d. The %dth Tribonacci number is:  %d\n", index+1, value, tr)
		fmt.Printf("%s\n\n", strings.Repeat("-", 100))
	}
}
