package main

import (
	"fmt"
	"strings"
)

func swapNodes(head *Node, k int) *Node {
	length := 0

	front := new(Node)
	end := new(Node)

	curr := head
	for curr != nil {
		length++

		if end != nil {
			end = end.next
		}

		if length == k {
			front = curr
			end = head
		}

		curr = curr.next
	}

	SwapTwoNodes(front, end)
	return head
}

func SwapTwoNodes(node1 *Node, node2 *Node) {
	temp := node1.data
	node1.data = node2.data
	node2.data = temp
}

// Driver code
func main() {
	inputList := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{6, 9, 3, 10, 7, 4, 6},
		{6, 9, 3, 4},
		{6, 2, 3, 6, 9},
		{6, 2},
	}
	k := []int{2, 3, 2, 3, 1}
	for index, value := range inputList {
		linkedList := new(LinkedList)
		linkedList.CreateLinkedList(value)
		fmt.Printf("%d.\tOriginal linked list is: ", index+1)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n\tk: %d\n", k[index])
		if k[index] <= 0 {
			fmt.Print("\tThe expected 'k' to have value from 1 to length of the linked list only.\n")
		} else {
			swapNodes(linkedList.head, k[index])
			fmt.Print("\tLinked list with swapped values: ")
			DisplayLinkedListWithForwardArrow(linkedList.head)
		}
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
