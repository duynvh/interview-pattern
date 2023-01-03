package main

import (
	"fmt"
	"strings"
)

func binarySearchRotated(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	if start > end {
		return -1
	}

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return mid
		}

		// start to mid is sorted
		if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < target && target <= nums[end] { // mid to end is sorted
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

// Driver code
func main() {
	numsList := [][]int{
		{6, 7, 1, 2, 3, 4, 5},
		{6, 7, 1, 2, 3, 4, 5},
		{4, 5, 6, 1, 2, 3},
		{4, 5, 6, 1, 2, 3},
	}
	targetList := []int{3, 6, 3, 6}

	for index, value := range numsList {
		fmt.Printf("%d.\tRotated array: %s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\tTarget %d found at index %d\n", targetList[index], binarySearchRotated(value, targetList[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
