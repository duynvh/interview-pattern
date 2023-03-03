package main

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}

	return s.head.value
}

func (s *Stack) Push(value interface{}) {
	s.head = &Node{value, s.head}
	s.size++
}

func (s *Stack) Pop() {
	if s.Empty() {
		return
	}

	s.head = s.head.next
	s.size--
}
