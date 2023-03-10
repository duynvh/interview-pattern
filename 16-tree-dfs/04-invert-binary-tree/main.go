package main

import (
	"fmt"
	"strings"
)

var change int
var masterRoot *EduBinaryTreeNode

func mirrorBinaryTree(root *EduBinaryTreeNode) *EduBinaryTreeNode {
	if root == nil {
		return nil
	}

	if root.left != nil {
		mirrorBinaryTree(root.left)
	}

	if root.right != nil {
		mirrorBinaryTree(root.right)
	}

	root.left, root.right = root.right, root.left
	return root
}

func main() {

	inputTrees := [][]int{
		{100, 50, 200, 25, 75, 125, 350},
		{100, 50, 200, 25, 110, 125, 350},
		{100, 50, 200, 25, 75, 90, 350},
		{25, 50, 75, 100, 125, 350},
		{350, 125, 100, 75, 50, 25},
		{100},
	}

	for i, input := range inputTrees {
		tree := new(EduBinaryTree)
		tree.BinaryTreeInitWithSlice(input)

		fmt.Printf("%d.\tBinary tree:\n", i+1)
		change = 0
		masterRoot = tree.root
		mirrorBinaryTree(tree.root)
		fmt.Printf("\n\tMirrored binary tree:\n")
		fmt.Printf("%s\n\n", strings.Repeat("-", 97))
	}
}
