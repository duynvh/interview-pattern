package main

type EduLinkedList struct {
    head *EduLinkedListNode
    tail *EduLinkedListNode
    size int
}

func LinkedListInit() *EduLinkedList {
    ll := &EduLinkedList{
        head: nil,
        tail: nil,
        size: 0,
    }

    return ll
}

func (ll *EduLinkedList) MoveToHead(node *EduLinkedListNode) {
    if node == nil {
        return
    }

    if node.prev != nil {
        node.prev.next = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    }

    if node == ll.head {
        ll.head = ll.head.next
    }

    if node == ll.tail {
        ll.tail = ll.tail.prev
        if ll.tail != nil {
            ll.tail.next = nil
        }
    }

    if ll.head == nil {
        ll.tail = node
        ll.head = node
    } else {
        node.next = ll.head
        ll.head.prev = node
        ll.head = node
    }
}

func (ll *EduLinkedList) InsertAtHead(pair [2]int) {
    newNode := LinkedListNodeInitWithPair(pair)
    if ll.head == nil {
        ll.tail = newNode
        ll.head = newNode
    } else {
        newNode.next = ll.head
        ll.head.prev = newNode
        ll.head = newNode
    }

    ll.size += 1
}

func (ll *EduLinkedList) InsertAtTail(pair [2]int) {
    newNode := LinkedListNodeInitWithPair(pair)
    if ll.tail == nil {
        ll.tail = newNode
        ll.head = newNode
        newNode.next = nil
    } else {
        ll.tail.next = newNode
        newNode.prev = ll.tail
        ll.tail = newNode
        newNode.next = nil
    }

    ll.size += 1
}

func (ll *EduLinkedList) Remove(pair [2]int) {
    i := ll.head
    for i != nil {
        if i.pair == pair {
            ll.RemoveNode(i)
            return
        }

        i = i.next
    }
}

func (ll *EduLinkedList) RemoveNode(node *EduLinkedListNode) {
    if node == nil {
        return
    }

    if node.prev != nil {
        node.prev.next = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    }

    if node == ll.head {
        ll.head = ll.head.next
    }

    if node == ll.tail {
        ll.tail = ll.tail.prev
        if ll.tail != nil {
            ll.tail.next = nil
        }
    }

    ll.size -= 1
    node = nil
}

// RemoveHead will remove the head from the EduLinkedList
func (ll *EduLinkedList) RemoveHead() {
	ll.RemoveNode(ll.head)
}

// RemoveTail will remove the tail from the EduLinkedList
func (ll *EduLinkedList) RemoveTail() {
	ll.RemoveNode(ll.tail)
}

// GetHead will return the head of the EduLinkedListNode
func (ll *EduLinkedList) GetHead() *EduLinkedListNode {
	return ll.head
}

// GetTail will return the tail of the Edulinkedlist
func (ll *EduLinkedList) GetTail() *EduLinkedListNode {
	return ll.tail
}