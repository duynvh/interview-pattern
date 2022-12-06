package main

import (
	"fmt"
	"strings"
)

func getMiddleNode(head *Node) *Node {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next

		fast = fast.next.next
	}

	return slow
}

// Driver code
func main() {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{3, 2, 1},
		{10},
		{1, 2},
	}
	for i, input := range inputs {
		linkedList := new(LinkedList)
		linkedList.CreateLinkedList(input)
		fmt.Printf("%d.\tLinked list: ", i+1)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n\tMiddle of linked list is: %d\n", getMiddleNode(linkedList.head).data)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}

}
