package main

import "strconv"

type Dequeue struct {
	items []*EduBinaryTreeNode
}

func NewDequeue() *Dequeue {
	return new(Dequeue)
}

func (s *Dequeue) PushFront(item *EduBinaryTreeNode) {
	temp := []*EduBinaryTreeNode{item}
	s.items = append(temp, s.items...)
}

func (s *Dequeue) PushBack(item *EduBinaryTreeNode) {
	s.items = append(s.items, item)
}

func (s *Dequeue) PopFront() *EduBinaryTreeNode {
	defer func() {
		s.items = s.items[1:]
	}()
	return s.items[0]
}

func (s *Dequeue) PopBack() *EduBinaryTreeNode {
	i := len(s.items) - 1
	defer func() {
		s.items = append(s.items[:i], s.items[i+1:]...)
	}()

	return s.items[i]
}

// Front will return the element from the front of the dequeue
func (s *Dequeue) Front() *EduBinaryTreeNode {
	return s.items[0]
}

// Back will return the element from the back of the dequeue
func (s *Dequeue) Back() *EduBinaryTreeNode {
	return s.items[len(s.items)-1]
}

// Empty will check if the dequeue is empty or not
func (s *Dequeue) Empty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

// Len will return the length of the dequeue
func (s *Dequeue) Len() int {
	return len(s.items)
}

// String function converts queue elements to string
func (s *Dequeue) String() string {
	output := "["
	for i := range s.items {
		output += strconv.Itoa(s.items[i].data) + ", "
	}
	return output[:len(output)-2] + "]"
}
