package main

import "fmt"

func reverse(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	listToDo := head.next
	reverseList := head
	reverseList.next = nil

	for listToDo != nil {
		temp := listToDo
		listToDo = listToDo.next
		temp.next = reverseList
		reverseList = temp
	}

	return reverseList
}

func main() {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	inputLinkedList := new(LinkedList)
	inputLinkedList.CreateLinkedList(inputList)
	fmt.Print("The original linked list: ")
	DisplayLinkedListWithForwardArrow(inputLinkedList.head)
	result := reverse(inputLinkedList.head)
	fmt.Printf("\n\nThe reversed linked list:   ")
	DisplayLinkedListWithForwardArrow(result)
}
