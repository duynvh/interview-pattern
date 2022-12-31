package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func findKthLargest(arr []int, k int) int {
	kNumbersMinHeap := newMinHeap()

	for i := 0; i < k; i++ {
		heap.Push(kNumbersMinHeap, arr[i])
	}

	for i := k; i < len(arr); i++ {
		if arr[i] > kNumbersMinHeap.Top() {
			heap.Pop(kNumbersMinHeap)
			heap.Push(kNumbersMinHeap, arr[i])
		}
	}

	return kNumbersMinHeap.Top()
}

// Driver code
func main() {
	inputList := [][]int{
		{1, 5, 12, 2, 11, 9, 7, 30, 20},
		{23, 13, 17, 19, 10},
		{3, 2, 5, 6, 7},
		{1, 4, 6, 0, 2},
		{1, 2, 3, 4, 5, 6, 7},
	}
	kList := []int{3, 4, 5, 1, 7}

	for i, list := range inputList {
		fmt.Printf("%d.\tInput array: %s, k: %d\n", i+1, strings.Replace(fmt.Sprint(list), " ", ", ", -1), kList[i])
		fmt.Printf("\tkth largest number: %d\n", findKthLargest(list, kList[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
