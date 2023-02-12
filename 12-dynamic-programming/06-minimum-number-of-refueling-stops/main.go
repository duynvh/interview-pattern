package main

import (
	"fmt"
	"math"
	"strings"
)

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)

	maxDistance := make([]int, n+1)
	maxDistance[0] = startFuel

	for index, value := range stations {
		station, fuel := value[0], value[1]

		for j := index; j >= 0; j-- {
			if maxDistance[j] >= station {
				nextDistance := maxDistance[j+1]
				newDistance := maxDistance[j] + fuel
				maxDistance[j+1] = int(math.Max(float64(nextDistance), float64(newDistance)))
			}
		}
	}

	for i := 0; i < n+1; i++ {
		if maxDistance[i] >= target {
			return i
		}
	}

	return -1
}

// Driver code
func main() {
	stations := [][][]int{
		{},
		{
			{10, 60},
			{20, 25},
			{30, 30},
			{60, 40},
		},
		{
			{2, 5},
			{3, 1},
			{6, 3},
			{12, 6},
		},
		{
			{140, 200},
			{160, 130},
			{310, 200},
			{330, 250},
		},
		{
			{310, 160},
			{380, 620},
			{700, 89},
			{850, 190},
			{900, 360},
		},
	}
	targetFuel := []int{3, 120, 15, 570, 1360}
	startFuel := []int{3, 10, 3, 140, 380}

	for index, _ := range stations {
		fmt.Printf("%d.\tStations: %s\n", index+1, strings.Replace(fmt.Sprint(stations[index]), " ", ", ", -1))
		fmt.Printf("\tTarget fuel: %d\n", targetFuel[index])
		fmt.Printf("\tStarting fuel: %d\n", startFuel[index])
		fmt.Printf("\n\tProcessing...\n")
		fmt.Printf("\n\tMinimum number of refueling stops: %d\n", minRefuelStops(targetFuel[index], startFuel[index], stations[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
