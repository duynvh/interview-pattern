package main

import "strconv"

type Interval struct {
	start  int
	end    int
	closed bool
}

type Position []Interval

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

func (pos Position) Len() int {
	return len(pos)
}

func (pos Position) Less(i, j int) bool {
	if pos[i].start == pos[j].start {
		return pos[i].end >= pos[j].end
	}

	return pos[i].start <= pos[j].start
}

func (pos Position) Swap(i, j int) {
	pos[i], pos[j] = pos[j], pos[i]
}
