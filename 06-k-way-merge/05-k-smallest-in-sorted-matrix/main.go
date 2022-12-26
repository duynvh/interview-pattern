package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func kthSmallestNumber(matrix [][]int, k int) int {
	minElement := newMinHeap()
	for _, list := range matrix {
		if len(list) == 0 {
			continue
		}

		heap.Push(minElement, Set{n1: list[0], n2: 0, n3: list})
	}

	numberCount, smallestNumber := 0, 0
	for minElement.Empty() == false {
		popValue := heap.Pop(minElement).(Set)
		smallestNumber = popValue.n1

		i, topNumList := popValue.n2, popValue.n3
		numberCount++

		if k == numberCount {
			break
		}

		if i+1 < len(topNumList) {
			heap.Push(minElement, Set{n1: topNumList[i+1], n2: i + 1, n3: topNumList})
		}
	}

	return smallestNumber
}

// Driver code
func main() {
	matrix := [][][]int{
		{{2, 6, 8}, {3, 7, 10}, {5, 8, 11}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{}, {}},
		{{1, 1, 3, 8}, {5, 5, 7, 9}, {3, 5, 8, 12}},
		{{2, 2, 6, 6, 8, 9}, {1, 4, 7, 10, 10, 22}, {5, 7, 9, 11, 13, 15}, {44, 55, 66, 77, 88, 99}},
	}
	k := []int{3, 4, 1, 30, 8}

	for i, tempMatrix := range matrix {
		fmt.Printf("%d.\t Input matrix: %s \n\t K = %d\n", i+1, strings.Replace(fmt.Sprint(tempMatrix), " ", ", ", -1), k[i])
		fmt.Printf("\t %dth smallest number from the given matrix is: %d\n", k[i], kthSmallestNumber(tempMatrix, k[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
