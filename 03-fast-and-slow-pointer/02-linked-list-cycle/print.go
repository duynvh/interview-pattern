package main

import "fmt"

func DisplayLinkedListWithForwardArrow(l *Node) {
	temp := l
	i := 0
	for temp != nil && i < 7 {
		if i == 0 {
			fmt.Printf("\t %d", temp.data)
		} else {
			fmt.Print(temp.data)
		}
		temp = temp.next
		if i == 6 {
			fmt.Printf(" (...)\n")
			break
		}
		if temp != nil {
			fmt.Printf("  → ")
		} else {
			fmt.Printf("  → nil\n")
		}
		i += 1
	}
}
