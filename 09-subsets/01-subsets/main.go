package main

import (
	"fmt"
	"math"
	"strings"
)

func getBit(num int, bit int) int {
	// shifts the first operand the specified number of bits to the left
	temp := 1 << bit

	// If binary number and current subset count are equal,
	// return 1 else return 0
	temp = temp & num
	if temp == 0 {
		return 0
	}

	return 1
}

func findAllSubsets(v []int) []Set {
	sets := make([]Set, 0)

	if len(v) == 0 {
		sets = append(sets, *NewSet())
		return sets
	} else {
		base, exp := float64(2), float64(len(v))

		subsetsCount := int(math.Pow(base, exp))
		for i := 0; i < subsetsCount; i++ {
			st := *NewSet()
			for j := 0; j < len(v); j++ {
				// If a specific bit is 1, choose that number from the original set
				// and add it to the current subset
				if getBit(i, j) == 1 && !st.Exists(v[j]) {
					st.Add(v[j])
				}
			}

			if i == 0 {
				sets = append(sets, *NewSet())
			} else {
				sets = append(sets, st)
			}
		}
	}

	return sets
}

// Driver code
func main() {
	array := [][]int{
		// {},
		{2, 5, 7},
		// {1, 2},
		// {1, 2, 3, 4},
		// {7, 3, 1, 5},
	}

	for i, arr := range array {
		// subsets := make([]Set, 0)
		fmt.Printf("%d. Set:     %s\n", i+1, ArrayToString(arr))
		subsets := findAllSubsets(arr)
		fmt.Printf("   Subsets: %s\n", DoubleArrayToString(DictToList(subsets)))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
