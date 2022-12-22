package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func kSmallestNumber(lists [][]int, k int) int {
	minElement := newMinHeap()
	for _, list := range lists {
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
		if numberCount == k {
			break
		}

		if (i + 1) < len(topNumList) {
			heap.Push(minElement, Set{n1: topNumList[i+1], n2: i + 1, n3: topNumList})
		}
	}

	return smallestNumber
}

func main() {
	lists := [][][]int{
		{{2, 6, 8}, {3, 6, 10}, {5, 8, 11}},
		{{1, 2, 3}, {4, 5}, {6, 7, 8, 15}, {10, 11, 12, 13}, {5, 10}},
		{{}, {}, {}},
		{{1, 1, 3, 8}, {5, 5, 7, 9}, {3, 5, 8, 12}},
		{{5, 8, 9, 17}, {1, 6, 6, 6}, {8, 17, 23, 24}},
	}
	k := []int{5, 50, 7, 4, 9}

	for i, list := range lists {
		fmt.Printf("%d.\t Input lists: %s \n\t K = %d\n", i+1, strings.Replace(fmt.Sprint(list), " ", ", ", -1), k[i])
		fmt.Printf("\t %dth smallest number from the given lists is: %d\n", k[i], kSmallestNumber(list, k[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
