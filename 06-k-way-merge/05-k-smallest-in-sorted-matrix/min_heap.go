package main

import "container/heap"

type Set struct {
	n1 int
	n2 int
	n3 []int
}

type MinHeap []Set

func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Empty() bool {
	return len(h) == 0
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].n1 < h[j].n1
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Top() interface{} {
	return h[0]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
