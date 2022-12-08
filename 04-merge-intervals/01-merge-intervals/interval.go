package main

import "strconv"

type Interval struct {
	start  int
	end    int
	closed bool
}

func (i *Interval) IntervalInit(start int, end int) {
	i.start = start
	i.end = end

	i.closed = true
}

func (i *Interval) setClosed(closed bool) {
	i.closed = closed
}

func (i *Interval) str() string {
	out := ""
	if i.closed {
		out = "[" + strconv.Itoa(i.start) + ", " + strconv.Itoa(i.end) + "]"
	} else {
		out = "(" + strconv.Itoa(i.start) + ", " + strconv.Itoa(i.end) + ")"
	}
	return out
}
