package main

import "strconv"

type Interval struct {
	start  int
	end    int
}

func (i *Interval) IntervalInit(start int, end int) {
	i.start = start
	i.end = end

}

func (i *Interval) str() string {
	out := "(" + strconv.Itoa(i.start) + ", " + strconv.Itoa(i.end) + ")"
	return out
}
