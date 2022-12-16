package main

import (
	"fmt"
	"strings"
)

func reorderList(head *Node) {
	if head == nil {
		return
	}

	// Find the middle of linked list
	// in 1->2->3->4->5->6 find 4
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}

	// reverse the second part of the list
	// convert 1->2->3->4->5->6 into 1->2->3 and 6->5->4
	// reverse the second half in-place
	prev := new(Node)
	prev = nil
	curr := slow
	tmp := new(Node)
	tmp = nil

	for curr != nil {
		tmp = curr.next
		curr.next = prev

		prev = curr
		curr = tmp
	}

	// Merge two sorted linked lists
	// merge 1->2->3 and 6->5->4 into 1->6->2->5->3->4
	first, second := head, prev
	for second.next != nil {
		tmp = first.next
		first.next = second
		first = tmp

		tmp = second.next
		second.next = first
		second = tmp
	}
}

// Driver code
func main() {
	inputLists := [][]int{
		{1, 1, 2, 2, 3, -1, 10, 12},
		// {10, 20, -22, 21, -12},
		// {1, 1, 1},
		// {-2, -5, -6, 0, -1, -4},
		// {3, 1, 5, 7, -4, -2, -1, -6},
	}

	for i, inputList := range inputLists {
		obj := new(LinkedList)
		obj.CreateLinkedList(inputList)
		fmt.Printf("%d. Original list: ", i+1)
		DisplayLinkedListWithForwardArrow(obj.head)
		reorderList(obj.head)
		fmt.Print("\n   After folding: ")
		DisplayLinkedListWithForwardArrow(obj.head)
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
