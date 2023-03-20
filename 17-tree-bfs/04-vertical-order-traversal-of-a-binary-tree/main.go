package main

import (
	"fmt"
	"strings"
)

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func verticalOrder(root *EduBinaryTreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	nodeList := make(map[int][]int)
	minColumn := 0
	maxIndex := 0

	queue := new(Queue)
	queue.Enqueue(root, 0)

	for !queue.IsEmpty() {
		node, column := queue.Dequeue()

		if node != nil {
			temp := nodeList[column]
			temp = append(temp, node.data)
			nodeList[column] = temp

			minColumn = min(minColumn, column)
			maxIndex = max(maxIndex, column)

			queue.Enqueue(node.left, column-1)
			queue.Enqueue(node.right, column+1)
		}
	}

	result := make([][]int, 0)
	for x := minColumn; x < maxIndex+1; x++ {
		result = append(result, nodeList[x])
	}

	return result
}

// Driver code
func main() {
	input := [][]int{
        {1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 6, 5, 7},
		{100, 50, 200, 25, 75, 300, 10, 350, 15},
		{20, 40, 50, 90, 67, 94},
		{-10, -23, 45, 25, 46},
		{0, -10, -5, 3, 6, 4, 8, 5, 70},
		{100, 70, 60, 40, 30, 20, 10},
	}
	for i := range input {
		tree := new(EduBinaryTree)
		tree.BinaryTreeInitWithSlice(input[i])
		fmt.Printf("%d.\tBinary tree:\n", i+1)
		fmt.Printf("\tVertical order traversal: %s\n", strings.Replace(fmt.Sprint(verticalOrder(tree.root)), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
