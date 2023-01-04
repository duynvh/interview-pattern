package main

import (
	"fmt"
	"strings"
)

func binarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2

		if array[mid] == target {
			return mid
		}

		if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

// abs function takes absolute of the given integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestElements(numbers []int, k int, num int) []int {
	if len(numbers) == k {
		return numbers
	}

	left := binarySearch(numbers, num) - 1
	right := left + 1
    /* Expand the window towards the side with the closer number
		Be careful to not go out of bounds with the pointers
		|a - num| < |b - num|,
		|a - num| == |b - num| */
	
    for right - left - 1 < k {
        if left == -1 {
            right += 1
            continue
        }
        /* Expand the window towards the side with the closer number
            Be careful to not go out of bounds with the pointers
            |a - num| < |b - num|,
            |a - num| == |b - num| */
        if right == len(numbers) || abs(numbers[left] - num) <= abs(numbers[right] - num) {
            left -= 1
        } else {
            right += 1
        }
    }

	return numbers[left+1 : right]
}

// Driver code
func main() {
	inputList := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5},
		{1, 2, 4, 5, 6},
		{1, 2, 3, 4, 5, 10},
	}
	k := []int{4, 4, 2, 3}
	num := []int{4, 3, 10, -5}

	for index, value := range inputList {
		fmt.Printf("%d.\tThe %d closest elements for the number %d in the array %s are:\n", index+1, k[index], num[index], strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\t %s\n", strings.Replace(fmt.Sprint(findClosestElements(value, k[index], num[index])), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
