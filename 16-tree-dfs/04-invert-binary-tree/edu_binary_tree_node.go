package main

type EduBinaryTreeNode struct {
    data int
    left *EduBinaryTreeNode
    right *EduBinaryTreeNode

    next *EduBinaryTreeNode
    parent *EduBinaryTreeNode
    count int
}

func NewBinaryTreeNode(data int) *EduBinaryTreeNode {
    t := new(EduBinaryTreeNode)
    t.data = data
	t.left = nil
	t.right = nil
	t.next = nil
	t.parent = nil
	t.count = 0
	return t
}