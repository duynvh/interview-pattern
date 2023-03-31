package main

import (
	"fmt"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func fractionToDecimal(numerator int, denominator int) string {
	result, remainders := "", make(map[int]int)

	if numerator == 0 {
		return "0"
	}

	condition1, condition2 := 0, 0
	if numerator < 0 {
		condition1 = 1
	}

	if denominator < 0 {
		condition2 = 1
	}

	if condition1^condition2 == 1 {
		result += "-"

		numerator = abs(numerator)
		denominator = abs(denominator)
	}

	quotinent := numerator / denominator
	remainder := (numerator % denominator) * 10
	result += strconv.Itoa(quotinent)

	if remainder == 0 {
		return result
	} else {
		result += "."

		for remainder != 0 {
			if beginning, ok := remainders[remainder]; ok {
				left := result[0:beginning]
				right := result[beginning:len(result)]
				result = left + "(" + right + ")"
				return result
			}

			remainders[remainder] = len(result)
			quotinent = remainder / denominator
			result += strconv.Itoa(quotinent)
			remainder = (remainder % denominator) * 10
		}
	}

	return result
}

// Driver code
func main() {
	numerator := []int{0, 4, 5, 2, 47, 93, -5, 47, -4}
	denominator := []int{4, 2, 333, 3, 18, 7, 333, -18, -2}

	for i := range numerator {
		fmt.Printf("%d.\tInput: fractionToDecimal(%d, %d)\n", i+1, numerator[i], denominator[i])
		result := fractionToDecimal(numerator[i], denominator[i])
		fmt.Printf("\tOutput: %s\n", result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
