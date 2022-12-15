package main

import (
	"fmt"
	"strings"
)

func reverseBetween(head *Node, left int, right int) *Node {
	ptr := head

	previous := new(Node)
	previous = nil

	revHead := new(Node)
	revHead = nil

	count := 1

	for count < left && ptr != nil {
		previous = ptr
		ptr = ptr.next
		count++
	}

	if ptr != nil {
		nextNode := ptr
		rightNode := new(Node)
		rightNode = nil

		for count <= right && nextNode != nil {
			rightNode = nextNode
			nextNode = nextNode.next
			count++
		}

		if rightNode != nil {
			revHead = reverse(ptr, left, right)
			if previous != nil {
				previous.next = revHead
			}

			if nextNode != nil {
				tmp := revHead
				for tmp.next != nil {
					tmp = tmp.next
				}

				tmp.next = nextNode
			}
		}
	}

	if previous != nil {
		return head
	} else {
		return revHead
	}
}

func reverse(head *Node, left int, right int) *Node {
	newHead := new(Node)
	newHead = nil

	ptr := head
	for right >= left {
		nextNode := ptr.next

		ptr.next = newHead
		newHead = ptr

		ptr = nextNode

		right -= 1
	}

	return newHead
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
	left := []int{
		1,
		3,
		2,
		1,
		1,
	}
	right := []int{
		5,
		6,
		4,
		3,
		2,
	}

	for index, value := range inputList {
		inputLinkedList := new(LinkedList)
		inputLinkedList.CreateLinkedList(value)
		fmt.Printf("%d.\tOriginal linked list is: ", index+1)
		DisplayLinkedListWithForwardArrow(inputLinkedList.head)
		if left[index] <= 0 {
			fmt.Printf("\tThe expected 'left' and 'right' to have value from 1 to length of the linked list only\n")
		} else {
			result := reverseBetween(inputLinkedList.head, left[index], right[index])
			fmt.Printf("\n\tReversed linked list is: ")
			DisplayLinkedListWithForwardArrow(result)
			fmt.Printf("\n")
			fmt.Printf("%s\n", strings.Repeat("-", 100))
		}
	}
}
