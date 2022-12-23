package main

import (
	"fmt"
	"strings"
)

func merge2List(head1 *Node, head2 *Node) *Node {
    dummy := &Node{data: -1}
    prev := dummy
    
    for head1 != nil && head2 != nil {
        if head1.data < head2.data {
            prev.next = head1
            head1 = head1.next
        } else {
            prev.next = head2
            head2 = head2.next
        }

        prev = prev.next
    }

    if head1 == nil {
        prev.next = head2
    } else {
        prev.next = head1
    }

    return dummy.next
}

func mergeKLists(lists []*LinkedList) *Node {
    if len(lists) > 0 {
        step := 1
        for step < len(lists) {
            for i := 0; i < len(lists) - step; i += step * 2 {
                lists[i].head = merge2List(lists[i].head, lists[i + step].head)
            }
            step *= 2
        }
        return lists[0].head
    }

    return nil
}

func main() {
	inputLists := [][][]int{
		{
			{21, 23, 42},
			{1, 2, 4},
		},
		{
			{11, 41, 51},
			{21, 23, 42},
		},
		{
			{2},
			{1, 2, 4},
			{25, 56, 66, 72},
		},
		{
			{11, 41, 51},
			{2},
			{2},
			{2},
			{1, 2, 4},
		},
		{
			{10, 30},
			{15, 25},
			{1, 7},
			{3, 9},
			{100, 300},
			{115, 125},
			{10, 70},
			{30, 90},
		},
	}

	for index, value := range inputLists {
		fmt.Printf("%d.\tInput lists:\n", index+1)
		llLists := make([]*LinkedList, 0)
		for _, v := range value {
			a := new(LinkedList)
			a.CreateLinkedList(v)
			llLists = append(llLists, a)
			fmt.Printf("\t")
			DisplayLinkedListWithForwardArrow(a.head)
			fmt.Printf("\n")
		}
		fmt.Printf("\tMerges list:\n")
		fmt.Printf("\t")
		DisplayLinkedListWithForwardArrow(mergeKLists(llLists))
		fmt.Printf("\n")
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
