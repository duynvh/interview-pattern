package main

import (
	"fmt"
	"strings"
)

func circularArrayLoop(arr []int) bool {
	currentDirection := false
	for i, v := range arr {
		if abs(arr[i]) == len(arr) {
			continue
		}

		if v >= 0 {
			currentDirection = true
		} else {
			currentDirection = false
		}

		slowPtr := i
		fastPtr := i

		for slowPtr != fastPtr || slowPtr != -1 || fastPtr != -1 {
			slowPtr = nextStep(arr, slowPtr, currentDirection)
			fastPtr = nextStep(arr, fastPtr, currentDirection)

			if fastPtr != -1 {
				fastPtr = nextStep(arr, fastPtr, currentDirection)
			}

			if slowPtr == -1 || fastPtr == -1 || slowPtr == fastPtr {
				break
			}
		}

		if slowPtr == fastPtr && slowPtr != -1 {
			return true
		}
	}

	return false
}

func nextStep(array []int, currentIndex int, currentDirection bool) int {
	nextDirection := false
	if array[currentIndex] >= 0 {
		nextDirection = true
	}

	if nextDirection != currentDirection {
		return -1
	}

	findStep := currentIndex + array[currentIndex]
	findStep = findStep % len(array)

	return findStep
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	inputArray := [][]int{
		{1, 3, -2, -4, 1},
		{2, 1, -1, -2, 2},
		{2, 1, -1, -2},
		{3, -2, -1, 1},
		{1, 2, -3, 3, 4, 7, 1},
	}

	for i, arr := range inputArray {
		fmt.Printf("%d.\tCircular array = %s\n", i+1, strings.Replace(fmt.Sprint(arr), " ", ", ", -1))
		fmt.Print("\n\t\tProcessing: \n")
		fmt.Printf("\n\tFound loop = %v\n", circularArrayLoop(arr))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
