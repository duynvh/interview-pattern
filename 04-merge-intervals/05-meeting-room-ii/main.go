package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals []Interval) int {
	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))

	for i := 0; i < len(intervals); i++ {
		starts[i] = intervals[i].start
		ends[i] = intervals[i].end
	}

	sort.Slice(starts, func(i, j int) bool {
		return starts[i] < starts[j]
	})

	sort.Slice(ends, func(i, j int) bool {
		return ends[i] < ends[j]
	})
	rooms, endsItr := 0, 0

	for i := 0; i < len(starts); i++ {
		if starts[i] < ends[endsItr] {
			rooms++
		} else {
			endsItr++
		}
	}

	return rooms
}

// Driver code
func main() {
	existingIntervals := make([]Interval, 0)
	list := [][]int{
		{2, 8},
		{3, 4},
		{3, 9},
		{5, 11},
		{8, 20},
		{11, 15},
	}
	for _, value := range list {
		interval := new(Interval)
		interval.IntervalInit(value[0], value[1])
		existingIntervals = append(existingIntervals, *interval)
	}

	fmt.Println(minMeetingRooms(existingIntervals))
}
