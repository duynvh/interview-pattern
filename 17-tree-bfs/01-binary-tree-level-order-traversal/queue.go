package main

import "strconv"

type Node struct {
	value *EduBinaryTreeNode
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Peek() *EduBinaryTreeNode {
	if q.IsEmpty() {
		return nil
	}

	return q.head.value
}

func (q *Queue) Enqueue(data *EduBinaryTreeNode) {
	temp := &Node{data, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}

	q.size++
}

// Dequeue remove an element from the queue
func (q *Queue) Dequeue() *EduBinaryTreeNode {
	if q.IsEmpty() {
		return nil
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value
}

// String function converts queue elements to string
func (q *Queue) String() string {
	output := "["
	if q.head != nil {
		temp := q.head
		for temp != nil {
			output += strconv.Itoa(temp.value.data) + ", "
			temp = temp.next
		}
		output = output[0 : len(output)-2]
	}
	output += "]"
	return output
}
