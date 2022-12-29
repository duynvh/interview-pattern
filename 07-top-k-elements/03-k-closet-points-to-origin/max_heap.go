package main

import "container/heap"

type MaxHeap []Point

func newMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].DistFromOrigin() > h[j].DistFromOrigin()
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Top() interface{} {
	return h[0]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
