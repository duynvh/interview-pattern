package main

import (
	"fmt"
	"strings"
)

func diameterHelper(node *EduBinaryTreeNode, diameterRes int) (int, int) {
	if node == nil {
		return 0, diameterRes
	}

	lh, diameterRes := diameterHelper(node.left, diameterRes)
	rh, diameterRes := diameterHelper(node.right, diameterRes)

	diameterRes = max(diameterRes, (lh + rh))
	return max(lh, rh) + 1, diameterRes
}

func diameterOfBinaryTree(root *EduBinaryTreeNode) int {
	if root == nil {
		return 0
	}

	diameterRes := 0
	_, diameterRes = diameterHelper(root, diameterRes)
	return diameterRes
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

// main is used for driver code
func main() {
	inputTrees := [][]int{
		{2, 1, 4, 3, 5, 6, 7},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{45, 32, 23, 21, 19, 18, 1},
		{5, 3, 4, 1, 2, 6, 7, 8, 9},
		{-1, -5, -8, -3, 1, 5, 3},
	}

	for i, input := range inputTrees {
		tree := new(EduBinaryTree)
		tree.BinaryTreeInitWithSlice(input)
		fmt.Printf("%d.\tOriginal tree:\n", i+1)
		// DisplayTree() is a custom function used for prinitng purposes
		fmt.Printf("\n\tDiameter of tree: %d\n", diameterOfBinaryTree(tree.root))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
