package main

type Node struct {
	data int
	next *Node
}

func NewNode(data int, next *Node) *Node {
	return &Node{
		data: data,
		next: next,
	}
}

func InitNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}
