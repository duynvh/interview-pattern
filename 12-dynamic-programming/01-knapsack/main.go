package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func findMaxKnapsackProfit(capacity int, weights []int, values []int) int {
	valuesLength := len(values)
	if capacity <= 0 || valuesLength == 0 || len(weights) != valuesLength {
		return 0
	}

	profits := make([]int, capacity+1)
	for i, _ := range profits {
		profits[i] = 0
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 0, ' ', tabwriter.Debug)
	for i := 0; i < valuesLength; i++ {
		for c := capacity; c > -1; c-- {
			if weights[i] <= c {
				newProfit := profits[c-weights[i]] + values[i]
				profits[c] = max(profits[c], newProfit)
			}
		}
	}
	w.Flush()

	return profits[capacity]
}

// Driver code
func main() {
	weights := []int{1, 2, 3, 5}
	values := []int{1, 5, 4, 8}
	capacity := 6

	fmt.Printf("We have the following list of values and weights for the capacity %d:\n", capacity)
	fmt.Printf("\n%s\n", strings.Repeat("-", 23))
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 0, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Weights", "\t", "Values")
	fmt.Fprintln(w, strings.Repeat("-", 10), "\t", strings.Repeat("-", 10))
	for i, _ := range values {
		fmt.Fprintln(w, weights[i], "\t", values[i])
	}
	w.Flush()
	result := findMaxKnapsackProfit(capacity, weights, values)
	fmt.Printf("\nThe maximum profit found: %d\n", result)
}
