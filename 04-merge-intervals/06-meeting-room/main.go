package main

import (
	"fmt"
	"sort"
)

func canAttendMeeting(intervals []Interval) bool {
	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1].end > intervals[i].start {
			return false
		}
	}

	return true
}

// Driver code
func main() {
	existingIntervals := make([]Interval, 0)
	list := [][]int{
		{7, 10},
		{2, 4},
		{15, 20},
	}
	for _, value := range list {
		interval := new(Interval)
		interval.IntervalInit(value[0], value[1])
		existingIntervals = append(existingIntervals, *interval)
	}

	fmt.Println(canAttendMeeting(existingIntervals))
}
