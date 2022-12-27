package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type KthLargest struct {
	k    int
	topK *MinHeap
}

func (this *KthLargest) KthLargestInit(k int, nums []int) {
	this.k = k
	this.topK = newMinHeap()
	for _, value := range nums {
		heap.Push(this.topK, value)
	}

	for this.topK.Len() > k {
		heap.Pop(this.topK)
	}
}

func (this *KthLargest) Add(value int) int {
	heap.Push(this.topK, value)

	for this.topK.Len() > this.k {
		heap.Pop(this.topK)
	}

	return this.ReturnKthLargest()
}

func (this *KthLargest) ReturnKthLargest() int {
	return this.topK.Top()
}

// Print function to print the heap with markers
func printHeapWithMarkers(minHeap MinHeap, pValue int) string {
	out := "["
	for i := 0; i < minHeap.Len(); i++ {
		if pValue == i {
			out += "«"
			out += strconv.Itoa(minHeap[i]) + "», "
		} else {
			out += strconv.Itoa(minHeap[i]) + ", "
		}
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}

// Driver code
func main() {
	nums := []int{3, 6, 9, 10}
	temp := []int{3, 6, 9, 10}
	val := []int{4, 7, 10, 8, 15}
	fmt.Printf("Initial stream: %s\n\n", strings.Replace(fmt.Sprint(nums), " ", ", ", -1))
	fmt.Print("k: ", 3, "\n")
	kLargest := new(KthLargest)
	kLargest.KthLargestInit(3, nums)
	for _, value := range val {
		fmt.Printf("\tAdding a new number %d to the stream\n", value)
		temp = append(temp, value)
		fmt.Printf("\t\tNumber stream: %s\n", strings.Replace(fmt.Sprint(temp), " ", ", ", -1))
		fmt.Printf("\tKth largest element in the stream: %d\n", kLargest.Add(value))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
