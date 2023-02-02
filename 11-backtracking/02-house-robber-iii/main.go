package main

import (
	"fmt"
	"strings"
)

func rob(root *BinaryTreeNode) int {
	return max(depthFS(root))
}

func depthFS(root *BinaryTreeNode) []int {
	// Empty tree case
	if root == nil {
		return []int{0, 0}
	}

	// Calculating the amount we can rob from the left children of the node
	leftChildren := depthFS(root.left)

	// Calculating the amount we can rob from the right children of the node
	rightChildren := depthFS(root.right)

	// Adding the maximum we can get from both sides
	notNode := max(leftChildren) + max(rightChildren)

	// Calculating value with node
	node := root.data + leftChildren[1] + rightChildren[1]

	// Returning both results, with node and without node
	return []int{node, notNode}
}

func max(x []int) int {
	if x[0] < x[1] {
		return x[1]
	}

	return x[0]
}

// Driver code
func main() {
	inputList := [][]int{
		{10, 9, 20, 15, 7},
		// {7, 9, 10, 15, 20},
		// {8, 2, 17, 1, 4, 19, 5},
		// {7, 3, 4, 1, 3},
		// {9, 5, 7, 1, 3},
	}

	for _, tree := range inputList {
		inputTree := new(BinaryTree)
		inputTree.BinaryTreeInitWithSlice(tree)
		result := rob(inputTree.root)
		fmt.Printf("\tMaximum amount we can rob without getting caught: %d\n", result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
