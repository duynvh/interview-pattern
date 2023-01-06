package main

import (
	"fmt"
	"strings"
)

func binarySearchRotated(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1

	if start > end {
		return false
	}

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}

		// if left half is sorted
		if nums[mid] > nums[start] {
			if target < nums[mid] && target >= nums[start] {
				end = mid
			} else {
				start = mid + 1
			}

			// if right half is sorted
		} else if nums[mid] < nums[start] {
			if target > nums[mid] && target < nums[start] {
				start = mid + 1
			} else {
				end = mid
			}
		} else {
			start++
		}
	}

	return false
}

// Driver code
func main() {
	numsList := [][]int{
		{-13, 3, 12, 13, 19, 57},
	}
	targetList := []int{19}

	for index, value := range numsList {
		fmt.Printf("%d.\tRotated array: %s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\tTarget %d found at index %d\n", targetList[index], binarySearchRotated(value, targetList[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
