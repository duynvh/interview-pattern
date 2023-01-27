package main

import (
	"fmt"
	"strings"
)

func sum(array []int) int {
	s := 0
	for _, num := range array {
		s += num
	}

	return s
}

func gasStationJourney(gas []int, cost []int) int {
	if sum(cost) > sum(gas) {
		return -1
	}

	totalGas, startingIndex := 0, 0
	for i := range gas {
		totalGas = totalGas + (gas[i] - cost[i])

		if totalGas < 0 {
			totalGas = 0
			startingIndex = i + 1
		}
	}

	return startingIndex
}

// Driver code
func main() {
	// gas := [][]int{{1, 2, 3, 4, 5}, {2, 3, 4}, {1, 1, 1, 1, 1}, {
	// 	1, 1, 1, 1, 10}, {1, 1, 1, 1, 1}, {1, 2, 3, 4, 5}}
	// cost := [][]int{{3, 4, 5, 1, 2}, {3, 4, 3}, {1, 2, 3, 4, 5}, {
	// 	2, 2, 1, 3, 1}, {1, 0, 1, 2, 3}, {1, 2, 3, 4, 5}}
	gas := [][]int{{3, 1, 1}}
	cost := [][]int{{1, 2, 2}}
	for i := range gas {
		fmt.Printf("%d.\tGas = %s\n", i+1, strings.Replace(fmt.Sprint(gas[i]), " ", ", ", -1))
		fmt.Printf("\tCost = %s\n", strings.Replace(fmt.Sprint(cost[i]), " ", ", ", -1))
		fmt.Printf("\n\tThe index of the gas station we can start our journey from is %d (If it's -1, then \n\tthat means no solution exists)\n", gasStationJourney(gas[i], cost[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
