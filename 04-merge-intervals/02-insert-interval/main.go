package main

import (
	"fmt"
	"strconv"
)

func insertInterval(existingIntervals []Interval, newInterval Interval) []Interval {
	newStart, newEnd := newInterval.start, newInterval.end

	i, n := 0, len(existingIntervals)

	output := make([]Interval, 0)

	for i < n && newStart > existingIntervals[i].start {
		output = append(output, existingIntervals[i])
		i++
	}

	if len(output) == 0 || output[len(output)-1].end < newStart {
		output = append(output, newInterval)
	} else {
		if output[len(output)-1].end < newEnd {
			output[len(output)-1].end = newEnd
		}
	}

	for i < n {
		if output[len(output)-1].end < existingIntervals[i].start {
			output = append(output, existingIntervals[i])
		} else {
			if existingIntervals[i].end > output[len(output)-1].end {
				output[len(output)-1].end = existingIntervals[i].end
			}
		}
		i++
	}

	fmt.Println(printIntervals(output))
	return output
}

// printIntervals is a helper function
func printIntervals(intervalList []Interval) string {
	out := "["
	for i, v := range intervalList {
		out += "[" + strconv.Itoa(v.start) + ", " + strconv.Itoa(v.end) + "]"
		if i != len(intervalList)-1 {
			out += ", "
		}
	}
	out += "]"
	return out
}

func main() {
	existingIntervals := make([]Interval, 0)
	list := [][]int{
		{1, 3},
		{6, 9},
	}
	for _, value := range list {
		interval := new(Interval)
		interval.IntervalInit(value[0], value[1])
		existingIntervals = append(existingIntervals, *interval)
	}
	newInterval := new(Interval)
	newInterval.IntervalInit(2, 5)
	insertInterval(existingIntervals, *newInterval)
}
