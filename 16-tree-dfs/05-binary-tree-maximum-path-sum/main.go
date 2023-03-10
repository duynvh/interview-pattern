package main

import (
	"fmt"
	"math"
	"strings"
)

var maxSum int = math.MinInt32

func maxPathSum(root *EduBinaryTreeNode) int {
	maxContrib(root)
	temp := maxSum
	maxSum = math.MinInt32
	return temp
}

func maxContrib(root *EduBinaryTreeNode) int {
	if root == nil {
		return 0
	}

	maxLeft := maxContrib(root.left)
	maxRight := maxContrib(root.right)

	leftSubtree, rightSubtree := 0, 0

	if maxLeft > 0 {
		leftSubtree = maxLeft
	}

	if maxRight > 0 {
		rightSubtree = maxRight
	}

	priceNewpath := root.data + leftSubtree + rightSubtree

	maxSum = max(maxSum, priceNewpath)

	return root.data + max(leftSubtree, rightSubtree)
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func main() {
	inputTrees := [][]int{
		{-8, 2, 17, 1, 4, 19, 5},
		{7, 3, 4, -1, -3},
		{-10, 9, 20, 30, 16, 15, 7},
		{1, 2, 3},
		{0},
	}

	for i, input := range inputTrees {
		testCaseTree := new(EduBinaryTree)
		testCaseTree.BinaryTreeInitWithSlice(input)
		if i > 0 {
			fmt.Print("\n")
		}

		// Displaying level-order traversal of trees being tested
		fmt.Printf("%d.\tBinary tree: \n", i+1)
		// custom function to print tree

		fmt.Printf("\n\t Maximum path sum:\t%d\n", maxPathSum(testCaseTree.root))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
