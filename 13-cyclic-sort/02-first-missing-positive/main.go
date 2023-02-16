package main

import (
	"fmt"
	"strings"
)

func smallestMissingPositiveInteger(nums []int) int {
	for i, _ := range nums {
		correctSpot := nums[i] - 1
		for 1 <= nums[i] && nums[i] <= len(nums) && nums[i] != nums[correctSpot] {
			nums[i], nums[correctSpot] = nums[correctSpot], nums[i]
			correctSpot = nums[i] - 1
		}
	}

	for i, num := range nums {
		if i+1 != num {
			return i + 1
		}
	}

	return len(nums) + 1
}

// Driver code
func main() {

	array := [][]int{
		// {1, 2, 3, 4},
		{-1, 3, 5, 7, 1},
		// {1, 5, 4, 3, 2},
		// {-1, 0, 2, 1, 4},
		// {1, 4, 3},
	}

	for i, arr := range array {
		fmt.Printf("%d.\tThe first missing positive integer in the array %s is:\n", i+1, strings.Replace(fmt.Sprint(arr), " ", ", ", -1))
		fmt.Printf("\t%d\n", smallestMissingPositiveInteger(arr))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
