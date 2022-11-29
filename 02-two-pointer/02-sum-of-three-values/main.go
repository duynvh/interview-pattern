// Solution summary
// 1. Sort the array in asc order
// 2. Loop through the entire array and set up two pointers (low and high) on every iteration
// 3. The low pointer is set to the current loop index + 1, and high is set to the last index of the array
// 4. Calculate the sum of array elements pointed to by the current loop index, and the low and high pointers
// 5. If the sum is equal to the target return true
// 6. If the sum is greater than the target, move the high pointer backward.
// 7. If the sum is less than the target, move the low pointer forward.
// 8. Repeat until the loop has processed the entire array
// 9. If after processing the entire array, we don't find a tripe that matches our requirement, we return false

package main

import (
	"fmt"
	"sort"
	"strings"
)

func findSumOfThree(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))

	low, high, triple := 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		low = i + 1
		high = len(nums) - 1

		for low < high {
			triple = nums[i] + nums[low] + nums[high]
			if triple == target {
				return true
			} else if triple < target {
				low += 1
			} else {
				high -= 1
			}
		}
	}

	return false
}

// Driver code
func main() {
	numsLists := [][]int{
		{3, 7, 1, 2, 8, 4, 5},
		{-1, 2, 1, -4, 5, -3},
		{2, 3, 4, 1, 7, 9},
		{1, -1, 0},
		{2, 4, 2, 7, 6, 3, 1},
	}
	testLists := [][]int{
		{10, 20, 21},
		{-8, 0, 7},
		{8, 10, 20},
		{1, -1, 0},
		{8, 11, 15},
	}
	for i, nList := range numsLists {
		fmt.Printf("%d. Input array: %s\n", i+1, strings.Replace(fmt.Sprint(nList), " ", ", ", -1))
		for _, tList := range testLists[i] {
			if findSumOfThree(nList, tList) {
				fmt.Printf("   Sum for %d exists\n", tList)
			} else {
				fmt.Printf("   Sum for %d does not exist\n", tList)
			}
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
