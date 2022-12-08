package main

import (
	"fmt"
	"sort"
	"strings"
)

func mergeIntervals(v []Interval) []Interval {
	if len(v) == 0 {
		return nil
	}

	sort.Sort(Position(v))
	result := make([]Interval, 0)

	result = append(result, v[0])

	for i := 1; i < len(v); i++ {
		lastAddedInterval := &result[len(result)-1]

		curStart := v[i].start
		curEnd := v[i].end

		prevEnd := lastAddedInterval.end

		if prevEnd >= curStart {
			if curEnd > prevEnd {
				lastAddedInterval.end = curEnd
			}
		} else {
			result = append(result, v[i])
		}
	}

	return result
}

func main() {
	inputList := [][][]int{
		{{1, 5},
			{3, 7},
			{4, 6},
		},
		{
			{1, 5},
			{4, 6},
			{6, 8},
			{11, 15},
		},
		{
			{3, 7},
			{6, 8},
			{10, 12},
			{11, 15},
		},
		{
			{1, 5},
		},
		{
			{1, 9},
			{4, 4},
			{3, 8},
		},
		{
			{1, 2},
			{3, 4},
			{8, 8},
		},
		{
			{1, 5},
			{1, 3},
		},
		{
			{1, 5},
			{6, 9},
		},
		{
			{0, 0},
			{1, 18},
			{1, 3},
		},
		{
			{1, 4},
			{0, 4},
		},
	}

	for index, value := range inputList {
		input := make([]Interval, 0)
		for _, v := range value {
			interval := new(Interval)
			interval.IntervalInit(v[0], v[1])
			input = append(input, *interval)
		}

		fmt.Printf("%d. Intervals to merge: %s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		result := mergeIntervals(input)
		fmt.Printf("\nMerged intervals:\t[")
		for i, v := range result {
			fmt.Printf("[%d, %d]", v.start, v.end)
			if i != len(result)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("]\n")
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
