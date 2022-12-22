package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func kSmallestPairs(arrayOne []int, arrayTwo []int, k int) [][]int {
	minHeap := newMinHeap()

	pairs := make([][]int, 0)

	// iterate over the length of array 1
	for i, _ := range arrayOne {
		// computing sum of pairs all elements of arrayOne with first index
		// of arrayTwo and placing it in the min-heap
		heap.Push(minHeap, Set{n1: arrayOne[i] + arrayTwo[0], n2: i, n3: 0})
	}

	counter := 1

	// iterate over elements of min-heap and only go upto k
	for !minHeap.Empty() && counter <= k {
		// placing sum of the top element of min-heap and
		// its corresponding pairs in i and j
		poppedValue := heap.Pop(minHeap).(Set)
		i, j := poppedValue.n2, poppedValue.n3

		// add pairs with the smallest sum in the new array
		pairs = append(pairs, []int{arrayOne[i], arrayTwo[j]})

		// increment the index for 2nd array, as we've compared all possible
		// pairs with the 1st index of arrayTwo.
		nextElement := j + 1

		// if next element is available for arrayTwo then add it to the heap
		if len(arrayTwo) > nextElement {
			heap.Push(minHeap, Set{n1: arrayOne[i] + arrayTwo[nextElement], n2: i, n3: nextElement})
		}
		counter++
	}

	return pairs
}

func main() {
	// Multiple inputs for efficient results
	arrayOne := [][]int{
		{2, 8, 9},
		{1, 2, 300},
		{1, 1, 2},
		{4, 6},
		{4, 7, 9},
		{1, 1, 2},
	}

	// multiple inputs for efficient results
	arrayTwo := [][]int{
		{1, 3, 6},
		{1, 11, 20, 35, 300},
		{1, 2, 3},
		{2, 3},
		{4, 7, 9},
		{1},
	}
	k := []int{9, 30, 1, 2, 5, 4}

	for i, list := range arrayOne {
		fmt.Printf("%d.\t Input Pairs: %s,%s \n\tK = %d\n", i+1, strings.Replace(fmt.Sprint(list), " ", ", ", -1), strings.Replace(fmt.Sprint(arrayTwo[i]), " ", ", ", -1), k[i])
		fmt.Printf("\tPairs with the smallest sum are: %s\n", strings.Replace(fmt.Sprint(kSmallestPairs(list, arrayTwo[i], k[i])), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
