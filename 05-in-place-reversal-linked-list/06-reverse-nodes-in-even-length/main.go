package main

import (
	"fmt"
	"strings"
)

func reverseEvenLengthGroups(head *Node) {
	// Point the prev to the node immediately before the current group
	prev := head
	/* The first group is always 1 which is odd and need not be reversed,
	   so, we start with the next group which is 2 */
	l := 2

	for prev.next != nil {
		node := prev
		n := 0
		for i := 0; i < l; i++ {
			if node.next == nil {
				break
			}
			n += 1
			node = node.next
		}

		// If length of the current group is even
		if n%2 == 0 {
			reverse := node.next
			curr := prev.next

			for i := 0; i < n; i++ {
				currNext := curr.next
				curr.next = reverse
				reverse = curr
				curr = currNext
			}

			prevNext := prev.next
			prev.next = node
			prev = prevNext
		} else { // If length of the current group is odd
			prev = node
		}
		l++
	}
}

func main() {
	inputList := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{15},
		{16, 17},
	}
	for index, value := range inputList {
		linkedList := new(LinkedList)
		linkedList.CreateLinkedList(value)
		fmt.Printf("%d.\tIf we reverse the even length groups of the linked list\n\t\t", index+1)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n\n\tWe will get:\n\t\t")
		reverseEvenLengthGroups(linkedList.head)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
