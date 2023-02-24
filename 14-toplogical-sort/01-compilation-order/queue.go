package main

import "fmt"

type Node struct {
	value interface{}
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

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return nil
	}

	return q.head.value
}

func (q *Queue) Enqueue(value interface{}) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return nil
	}

	value := q.head.value
	q.head = q.head.next
	q.size--
	return value
}

// Print prints the queue on to the console
func (q *Queue) Print() string {
	if q.Size() == 0 {
		return "[]"
	}
	temp := q.head
	out := "["
	for i := 0; i < q.Size(); i++ {
		out = out + "'" + string(temp.value.(rune)) + "'" + ", "
		temp = temp.next
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}
