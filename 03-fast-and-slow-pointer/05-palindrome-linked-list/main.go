package main

import (
	"fmt"
	"strings"
)

// palindrome function checks if the given linked list is a palindrome or not
func palindrome(head *Node) bool {
	// Initialize variables with head
	slow := head
	fast := head
	midNode := head
	revertData := new(Node)
	revertData = nil

	// Traverse LinkedList through fast and slow pointers to get the middle node
	for fast != nil && fast.next != nil {
		midNode = slow
		slow = slow.next
		fast = fast.next.next
	}

	// Fast pointer of odd linked list will point to last node of linked list but it will point to NULL for even linked list
	savedOddMidNode := new(Node)
	savedOddMidNode = nil
	if fast != nil {
		savedOddMidNode = slow
		slow = slow.next
	}
	fmt.Printf("\n")
	fmt.Println(slow.data)

	// It will skip the first half
	midNode.next = nil

	// Pass middle node as a head to reverse function to revert the next half of LinkedList
	revertData = ReverseList(slow)

	// Pass second half reverted data to compareTwoHalves function to check the palindrome property
	check := false
	check = compareTwoHalves(head, revertData)

	// Revert second half back to the original list
	revertData = ReverseList(revertData)

	// Connect both halves
	if savedOddMidNode != nil {
		midNode.next = savedOddMidNode
		savedOddMidNode.next = revertData
	} else {
		midNode.next = revertData
	}

	// Return true if there's only one node or both are pointing to nil
	if head == nil || revertData == nil {
		return true
	}
	if check {
		return true
	}
	return false
}

func compareTwoHalves(firstHalf *Node, secondHalf *Node) bool {
	check := false
	for firstHalf != nil && secondHalf != nil {
		if firstHalf.data != secondHalf.data {
			check = false
			break
		} else {
			check = true
		}

		firstHalf = firstHalf.next
		secondHalf = secondHalf.next
	}

	return check
}

// Driver code
func main() {
	inputArray := [][]int{
		{2, 4, 6, 4, 2},
		{0, 3, 5, 5, 0},
		{9, 7, 4, 4, 7, 9},
		{5, 4, 7, 9, 4, 5},
		{5, 9, 8, 3, 8, 9, 5},
	}

	for i, input := range inputArray {
		linkedList := new(LinkedList)
		linkedList.CreateLinkedList(input)
		fmt.Printf("%d\tLinked List: ", i+1)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n")
		fmt.Printf("\tIs is Palindrome = ")
		if palindrome(linkedList.head) {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
