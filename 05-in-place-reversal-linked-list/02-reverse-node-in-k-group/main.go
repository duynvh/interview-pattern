package main

import (
	"fmt"
)

func reverse(head *Node, k int) (*Node, *Node, *Node) {
	previous, current, next := new(Node), head, new(Node)
	previous, next = nil, nil
	index := 0

	for current != nil && index < k {
		index += 1
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	return previous, current, next
}

func findLength(start *Node) int {
	current := start
	count := 0
	for current != nil {
		current = current.next
		count++
	}

	return count
}

func reverseLinkedList(head *Node, k int) *Node {
	if k <= 1 || head == nil {
		return head
	}

	i, count := 0, 0
	current, previous := head, new(Node)
	previous = nil
	totalLength := findLength(head)

	for true {
		i++
		lastNodeOfPreviousPart := previous
		lastNodeOfCurrentPart := current
		previous, current, _ = reverse(lastNodeOfCurrentPart, k)
		count += k

		if lastNodeOfPreviousPart != nil {
			lastNodeOfPreviousPart.next = previous
		} else {
			head = previous
		}

		lastNodeOfCurrentPart.next = current
		if current == nil || (totalLength-count) < k {
			break
		}

		previous = lastNodeOfCurrentPart
	}

	return head
}

func main() {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	inputLinkedList := new(LinkedList)
	inputLinkedList.CreateLinkedList(inputList)

	fmt.Print("The original linked list: ")
	DisplayLinkedListWithForwardArrow(inputLinkedList.head)
	result := reverseLinkedList(inputLinkedList.head, 3)
	fmt.Printf("\n\nReversed linked list, with k = %d:   ", 3)
	DisplayLinkedListWithForwardArrow(result)
}
