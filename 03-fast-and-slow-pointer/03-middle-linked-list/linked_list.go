package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertNodeAtHead(node *Node) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *LinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := InitNode(lst[i])
		l.InsertNodeAtHead(newNode)
	}
}

func (l *LinkedList) Display() {
	temp := l.head
	fmt.Print("[")

	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}
