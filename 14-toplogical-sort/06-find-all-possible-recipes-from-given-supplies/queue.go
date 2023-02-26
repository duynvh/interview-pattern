package main

type Node struct {
    value interface{}
    next *Node
}

type Queue struct {
    head *Node
    tail *Node
    size int
}

func (q Queue) Size() int {
    return q.size
}

func (q Queue) IsEmpty() bool {
    return q.size == 0
}

func (q Queue) Peak() interface{} {
    if q.IsEmpty() {
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
        return nil
    }

    value := q.head.value
    q.head = q.head.next
    q.size--
    return value
}