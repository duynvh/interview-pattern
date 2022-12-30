package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func topKFrequent(nums []int, k int) []int {
	// Find the frequency of each number
	numFrequencyMap := make(map[int]int)
	for _, n := range nums {
		numFrequencyMap[n] = numFrequencyMap[n] + 1
	}
	topKElements := newMinHeap()

	// Go through all numbers of the numFrequencyMap and push them in the minHeap, which will have
	// top k frequent numbers. If the heap size is more than k, we remove the smallest (top) number
	for num, frequency := range numFrequencyMap {
		heap.Push(topKElements, Pair{num, frequency})

		if topKElements.Len() > k {
			heap.Pop(topKElements)
		}
	}

	// Create an array of top k numbers
	topNumbers := make([]int, 0)
	for !topKElements.Empty() {
		pair := topKElements.Top()
		heap.Pop(topKElements)
		topNumbers = append(topNumbers, pair.first)
	}

	return topNumbers
}

func main() {
	nums := [][]int{
		{1, 3, 5, 12, 11, 12, 11, 12, 5},
		{1, 3, 5, 14, 18, 14, 5},
		{2, 3, 4, 5, 6, 7, 7},
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{2, 4, 3, 2, 3, 4, 5, 4, 4, 4},
		{1, 1, 1, 1, 1, 1},
		{2, 3},
	}
	k := []int{3, 2, 1, 1, 3, 4, 3}

	for i, num := range nums {
		fmt.Printf("%d.\tInput: (%s, %d)\n", i+1, strings.Replace(fmt.Sprint(num), " ", ", ", -1), k[i])
		fmt.Printf("\tTop %d frequent elements: %s\n", k[i], strings.Replace(fmt.Sprint(topKFrequent(num, k[i])), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
