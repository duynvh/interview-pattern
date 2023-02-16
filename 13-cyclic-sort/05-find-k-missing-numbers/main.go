package main

import (
	"fmt"
	"sort"
	"strings"
)

func firstKMissingNumbers(nums []int, k int) []int {
	result := make([]int, 0)
	sort.Ints(nums)
	i := 0
	for i < len(nums) && nums[i] <= 0 {
		i++
	}

	count, curr := 0, 1
	for count < k && i < len(nums) {
		if nums[i] != curr {
			result = append(result, curr)
			count++
		} else {
			i++
		}

		curr++
	}

	for count < k {
		result = append(result, curr)
		curr++
		count++
	}

	return result
}

// Driver code
func main() {
	array := [][]int{
		{1, 2, 3, 0, 4, 9, 7},
		// {3, 1, 2, 3, 6, 4},
		// {4, 1, 2, 1, 6, 3},
		// {4, 3, 4, 5, 1},
		// {5, 3, 5, 6, 2, 1},
		// {1, 2, 3, 4, 5},
		// {5, 4, 3, 2, 1},
	}

	for i, arr := range array {
		fmt.Printf("%d.\tGiven array: %s\n", i+1, strings.Replace(fmt.Sprint(firstKMissingNumbers(arr, 4)), " ", ", ", -1))
		// result := findCorruptPair(array[i])
		// if result[0] == 0 && result[1] == 0 {
		// 	fmt.Printf("\n\tNo corrupt pair found.\n")
		// } else {
		// 	fmt.Printf("\n\tCorrupt pair: %d, %d\n", result[0], result[1])
		// }
		// fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
