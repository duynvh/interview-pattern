package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func employeeFreeTime(schedule [][]Interval) []Interval {
	ans := make([]Interval, 0)
	intervals := make([]Interval, 0)

	for i, _ := range schedule {
		for j, _ := range schedule[i] {
			intervals = append(intervals, schedule[i][j])
		}
	}

	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	prevEnd := intervals[0].end
	for _, interval := range intervals {
		if interval.start > prevEnd {
			tempInterval := new(Interval)
			tempInterval.IntervalInit(prevEnd, interval.start)
			ans = append(ans, *tempInterval)
		}
		prevEnd = max(prevEnd, interval.end)
	}

	return ans
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

// intervalToString is used to convert Interval to string for
// printing purposes
func intervalToString(IntervalArray []Interval) string {
	str := "["
	for i, v := range IntervalArray {
		str += "[" + strconv.Itoa(v.start) + ", " + strconv.Itoa(v.end) + "]"
		if i+1 < len(IntervalArray) {
			str += ", "
		}
	}
	str += "]"
	return str
}

// Driver code
func main() {
	inputs := [][][]Interval{
		{{Interval{1, 2}, Interval{5, 6}}, {Interval{1, 3}}, {Interval{4, 10}}},
		{{Interval{1, 3}, Interval{6, 7}}, {Interval{2, 4}}, {Interval{2, 5}, Interval{9, 12}}},
		{{Interval{2, 3}, Interval{7, 9}}, {Interval{1, 4}, Interval{6, 7}}},
		{{Interval{3, 5}, Interval{8, 10}}, {Interval{4, 6}, Interval{9, 12}}, {Interval{5, 6}, Interval{8, 10}}},
		{{Interval{1, 3}, Interval{6, 9}, Interval{10, 11}}, {Interval{3, 4}, Interval{7, 12}}, {Interval{1, 3}, Interval{7, 10}}, {Interval{1, 4}}, {Interval{7, 10}, Interval{11, 12}}},
		{{Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}}, {Interval{2, 3}, Interval{4, 5}, Interval{6, 8}}},
		{{Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}, Interval{9, 10}, Interval{11, 12}}, {Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}, Interval{9, 10}, Interval{11, 12}}, {Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}, Interval{9, 10}, Interval{11, 12}}, {Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}, Interval{9, 10}, Interval{11, 12}}},
	}
	for i, schedule := range inputs {
		fmt.Printf("%d.\tEmployee Schedules:\n", i+1)
		for _, s := range schedule {
			fmt.Printf("\t\t%s\n", intervalToString(s))
		}
		fmt.Printf("\tEmployees' free time %s\n", intervalToString(employeeFreeTime(schedule)))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
