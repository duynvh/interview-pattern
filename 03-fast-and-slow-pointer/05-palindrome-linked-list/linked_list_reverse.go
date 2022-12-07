package main

func ReverseList(slowPtr *Node) *Node {
	reverse := new(Node)
	reverse = nil

	for slowPtr != nil {
		next := slowPtr.next
		slowPtr.next = reverse
		reverse = slowPtr
		slowPtr = next
	}

	return reverse
}
