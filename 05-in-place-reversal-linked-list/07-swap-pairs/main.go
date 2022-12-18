package main

import (
	"fmt"
	"strings"
)

func swapPairs(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	list := head.next
	var previousNode *Node

	for head != nil && head.next != nil {
		nextNode := head.next
		head.next = nextNode.next
		nextNode.next = head

		if previousNode != nil {
			previousNode.next = nextNode
		}

		previousNode = head
		head = head.next
	}

	return list
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
		DisplayLinkedListWithForwardArrow(swapPairs(linkedList.head))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
