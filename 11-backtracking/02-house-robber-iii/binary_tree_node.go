package main

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode

	next   *BinaryTreeNode
	parent *BinaryTreeNode
	count  int
}

func NewBinaryTreeNode(data int) *BinaryTreeNode {
	t := new(BinaryTreeNode)
	t.data = data
	t.left = nil
	t.right = nil
	t.next = nil
	t.parent = nil
	t.count = 0
	return t
}
