package main

import (
	"fmt"
	"strings"
)

func findCorruptPair(nums []int) []int {
	swap := func(first int, second int) {
		temp := nums[first]
		nums[first] = nums[second]
		nums[second] = temp
	}

	for i := 0; i < len(nums); {
		correct := nums[i] - 1

		if nums[i] != nums[correct] {
			swap(i, correct)
		} else {
			i += 1
		}
	}

	missing, duplicated := 0, 0
	for j, num := range nums {
		if num != j+1 {
			duplicated = num
			missing = j + 1
		}
	}

	return []int{missing, duplicated}
}

// Driver code
func main() {
	array := [][]int{
		{3, 1, 2, 5, 2},
		// {3, 1, 2, 3, 6, 4},
		// {4, 1, 2, 1, 6, 3},
		// {4, 3, 4, 5, 1},
		// {5, 3, 5, 6, 2, 1},
		// {1, 2, 3, 4, 5},
		// {5, 4, 3, 2, 1},
	}

	for i, arr := range array {
		fmt.Printf("%d.\tGiven array: %s\n", i+1, strings.Replace(fmt.Sprint(arr), " ", ", ", -1))
		result := findCorruptPair(array[i])
		if result[0] == 0 && result[1] == 0 {
			fmt.Printf("\n\tNo corrupt pair found.\n")
		} else {
			fmt.Printf("\n\tCorrupt pair: %d, %d\n", result[0], result[1])
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
