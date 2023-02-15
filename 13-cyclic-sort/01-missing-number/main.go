package main

import (
	"fmt"
	"strings"
)

func findMissingNumber(nums []int) int {
	lenNums := len(nums)
	index := 0

	for index < lenNums {
		value := nums[index]

		if value < lenNums && value != nums[value] {
			nums[index], nums[value] = nums[value], nums[index]
			continue
		}

		index++
	}

	for i, x := range nums {
		if i != x {
			return i
		}
	}

	return lenNums
}

// Driver code
func main() {
	inputNumbers := [][]int{
		{4, 0, 3, 1},
		// {8, 3, 5, 2, 4, 6, 0, 1},
		// {1, 2, 3, 4, 6, 7, 8, 9, 10, 11},
		// {0},
		// {1, 2, 3, 0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23}
	}

	for i, x := range inputNumbers {
		fmt.Printf("%d.\tnums: %s\n", i+1, strings.Replace(fmt.Sprint(x), " ", ", ", -1))
		fmt.Printf("\n\tMissing number: %d\n", findMissingNumber(x))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
