package main

import (
	"fmt"
	"math"
	"strings"
)

func minSubArrayLen(target int, nums []int) int {
	windowSize := math.MaxInt64

	totalSum, start := 0, 0

	for end, _ := range nums {
		totalSum += nums[end]

		for totalSum >= target {
			currLen := end - start + 1
			windowSize = min(windowSize, currLen)

			totalSum -= nums[start]
			start += 1
		}
	}

	if windowSize != math.MaxInt64 {
		return windowSize
	} else {
		return 0
	}
}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func main() {
	targets := []int{7, 4, 11, 10, 5, 15}
	numsList := [][]int{{2, 3, 1, 2, 4, 3}, {1, 4, 4}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 2, 3, 4}, {1, 2, 1, 3}, {5, 4, 9, 8, 11, 3, 7, 12, 15, 44}}
	for i, target := range targets {
		result := minSubArrayLen(target, numsList[i])
		fmt.Printf("%d.\tminSubArrayLen(%d, %s): %d\n", i+1, target, strings.Replace(fmt.Sprint(numsList[i]), " ", ", ", -1), result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
