package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func canPartitionArray(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	sum = sum / 2
	matrix := make([][]int, sum+1)
	for i, _ := range matrix {
		matrix[i] = make([]int, len(nums)+1)
		for j, _ := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	for i := 0; i < len(nums)+1; i++ {
		matrix[0][i] = 1
	}

	for i := 1; i < sum+1; i++ {
		matrix[i][0] = 0
	}

	for i := 1; i < sum+1; i++ {
		for j := 1; j < len(nums)+1; j++ {
			if nums[j-1] > i {
				matrix[i][j] = matrix[i][j-1]
			} else {
				if matrix[i-nums[j-1]][j-1] == 1 {
					matrix[i][j] = matrix[i-nums[j-1]][j-1]
				} else {
					matrix[i][j] = matrix[i][j-1]
				}
			}
		}
	}

	// Print 2d array in matrix format
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 0, ' ', 0)
	for _, arr := range matrix {
		for _, v := range arr {
			fmt.Fprint(w, "\t", v)
		}
		fmt.Fprint(w, "\n")
	}
	w.Flush()

	if matrix[sum][len(nums)] == 0 {
		return false
	}

	return true
}

// Driver code
func main() {
	array := [][]int{
		// {3, 1, 1, 2, 2, 1},
		// {1, 3, 7, 3},
		// {1, 2, 3},
		// {1, 2, 5},
		{1, 3, 4, 8},
		// {1, 2, 3, 2, 3, 5},
		// {1, 5, 3, 2, 3, 19, 3},
		// {1, 2, 3, 5, 3, 2, 1},
	}

	for i, arr := range array {
		fmt.Printf("%d. Input array: %s\n", i+1, strings.Replace(fmt.Sprint(arr), " ", ", ", -1))
		result := canPartitionArray(arr)
		fmt.Printf("\nCan we partition array into equal sum?: %v\n", result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
