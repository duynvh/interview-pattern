package main

type EduLinkedListNode struct {
    second int
    first int
    pair [2]int
    next *EduLinkedListNode
    prev *EduLinkedListNode
}

func LinkedListNodeInitWithPair(pair [2]int) *EduLinkedListNode {
    setNode := &EduLinkedListNode{
        second: pair[1],
        first: pair[0],
        pair: pair,
        next: nil,
        prev: nil,
    }

    return setNode
}