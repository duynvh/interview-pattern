package main

import (
	"fmt"
	"strconv"
	"strings"
)

// intervalsIntersection is the challenge function
func intervalsIntersection(intervalsA []Interval, intervalsB []Interval) []Interval {
	// Index "i" to iterate over the length of list a and index "j" to iterate over the length of list b
	i, j := 0, 0

	// To store all intersecting intervals
	intersections := make([]Interval, 0)

	// While loop will break whenever either of the lists ends
	for i < len(intervalsA) && j < len(intervalsB) {
		// Let's check if intervalsA[i] intersect intervalsB[j]
		// 1. start - the potential startpoint of the intersection
		// 2. end - the potential endpoint of the intersection
		start := max(intervalsA[i].start, intervalsB[j].start)
		end := min(intervalsA[i].end, intervalsB[j].end)

		// If this is an actual intersection, add it to the list
		if start <= end {
			intersections = append(intersections, Interval{
				start: start,
				end:   end,
			})
		}

		// Move forward in the list whose interval ends earlier
		if intervalsA[i].end < intervalsB[j].end {
			i++
		} else if intervalsA[i].end > intervalsB[j].end {
			j++
		} else {
			i++
			j++
		}
	}

	return intersections
}

// Native approach
func intervalsIntersection2(intervalsA []Interval, intervalsB []Interval) []Interval {
	// Index "i" to iterate over the length of list a and index "j" to iterate over the length of list b
	// To store all intersecting intervals
	intersections := make([]Interval, 0)

	// While loop will break whenever either of the lists ends
	for i := 0; i < len(intervalsA); i++ {
		for j := 0; j < len(intervalsB); j++ {
			startA, startB, endA, endB := intervalsA[i].start, intervalsB[j].start, intervalsA[i].end, intervalsB[j].end

			if endA < startB || endB < startA {
				continue
			}

			start := max(startA, startB)
			end := min(endA, endB)
			intersections = append(intersections, Interval{
				start: start,
				end:   end,
			})
		}
	}
	return intersections
}

// max function returns the maximum of the integers provided
func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

// min function returns the minimum of the integers provided
func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func printInterval(interval Interval) string {
	out := "[" + strconv.Itoa(interval.start) + ", " + strconv.Itoa(interval.end) + "]"
	return out
}

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

// Driver code
func main() {
	intervalsA := [][][]int{
		{
			{1, 2},
		},
		{
			{1, 4},
			{5, 6},
			{9, 15},
		},
		{
			{3, 6},
			{8, 16},
			{17, 25},
		},
		{
			{4, 7},
			{9, 16},
			{17, 28},
			{8, 10},
			{39, 50},
			{55, 66},
			{70, 89},
		},
		{
			{1, 3},
			{5, 6},
			{7, 8},
			{12, 15},
		},
	}
	intervalsB := [][][]int{
		{
			{1, 2},
		},
		{
			{2, 4},
			{5, 7},
			{9, 15},
		},
		{
			{2, 3},
			{10, 15},
			{18, 23},
		},
		{
			{3, 6},
			{7, 8},
			{9, 10},
			{14, 19},
			{23, 33},
			{35, 40},
			{45, 59},
			{60, 64},
			{68, 76},
		},
		{
			{2, 4},
			{7, 10},
		},
	}

	for i, _ := range intervalsA {
		intervalListA := make([]Interval, 0)
		intervalListB := make([]Interval, 0)
		for _, v := range intervalsA[i] {
			interval := new(Interval)
			interval.IntervalInit(v[0], v[1])
			intervalListA = append(intervalListA, *interval)
		}
		for _, v := range intervalsB[i] {
			interval := new(Interval)
			interval.IntervalInit(v[0], v[1])
			intervalListB = append(intervalListB, *interval)
		}
		fmt.Printf("\tIntersecting intervals in 'A' and 'B' are: %s\n", printIntervals(intervalsIntersection(intervalListA, intervalListB)))

		fmt.Printf("\tIntersecting intervals 2 in 'A' and 'B' are: %s\n", printIntervals(intervalsIntersection2(intervalListA, intervalListB)))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
