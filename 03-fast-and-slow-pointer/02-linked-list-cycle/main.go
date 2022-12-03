package main

import (
	"fmt"
	"strings"
)

func detectCycle(head *Node) bool {
	slow := head
	fast := head
	i, j := 0, 0
	DisplayLinkedListWithForwardArrow(head)
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		i += 1
		j += 2

		if slow == fast {
			return true
		}
	}

	return false
}

// Driver code
func main() {
	inputArray := [][]int{
		{2, 4, 6, 8, 10, 12},
		// {1, 3, 5, 7, 9, 11},
		// {0, 1, 2, 3, 4, 6},
		// {3, 4, 7, 9, 11, 17},
		// {5, 1, 4, 9, 2, 3},
	}
	for i, input := range inputArray {
		linkedList := new(LinkedList)
		linkedList.CreateLinkedList(input)
		if i%2 == 0 {
			linkedList.head.next.next.next.next.next.next = linkedList.head.next
		}
		fmt.Printf("%d\tInput:\n", i+1)
		fmt.Printf("\n\tDetected Cycle = %v\n", detectCycle(linkedList.head))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
