package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
	reverseArray() is a helper function used by main

to reverse an integer array
*/
func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func levelOrderTraversal(root *EduBinaryTreeNode) string {
	result := ""

	if root == nil {
		return "nil"
	} else {
		result = ""
		queues := make([]*Queue, 0)
		queues = append(queues, &Queue{nil, nil, 0})
		queues = append(queues, &Queue{nil, nil, 0})
		currentQueue, nextQueue := queues[0], queues[1]
		currentQueue.Enqueue(root)
		levelNumber := 0
		n := 0

		for currentQueue.Size() > 0 {
			n++
			temp := currentQueue.Dequeue()
			result += strconv.Itoa(temp.data)

			if temp.left != nil {
				nextQueue.Enqueue(temp.left)
			}

			if temp.right != nil {
				nextQueue.Enqueue(temp.right)
			}

			if currentQueue.Size() == 0 {
				levelNumber++
				if nextQueue.Size() > 0 {
					result += " : "
				}

				currentQueue = queues[levelNumber%2]
				nextQueue = queues[(levelNumber+1)%2]
			} else {
				result += ", "
			}
		}
	}

	return result
}

func levelOrderTraversal2(root *EduBinaryTreeNode) string {
	result := ""
	if root == nil {
		result = "nil"
	} else {
		currentQueue := new(Queue)
		dummyNode := NewBinaryTreeNode(0)
		currentQueue.Enqueue(root)
		currentQueue.Enqueue(dummyNode)

		for !currentQueue.IsEmpty() {
			temp := currentQueue.Dequeue()
			result += strconv.Itoa(temp.data)

			if temp.left != nil {
				currentQueue.Enqueue(temp.left)
			}

			if temp.right != nil {
				currentQueue.Enqueue(temp.right)
			}

			if currentQueue.Peek() == dummyNode {
				currentQueue.Dequeue()

				if !currentQueue.IsEmpty() {
					result += " : "
					currentQueue.Enqueue(dummyNode)
				}
			} else {
				result += ", "
			}
		}
	}

	return result
}

// Driver code
func main() {
	// Creating a binary tree
	input1 := []int{100, 50, 200, 25, 75, 350}
	tree1 := new(EduBinaryTree)
	tree1.BinaryTreeInitWithSlice(input1)

	// Creating a right degenerate binary tree
	input2 := input1
	sort.Ints(input2)
	tree2 := new(EduBinaryTree)
	tree2.BinaryTreeInitWithSlice(input2)

	// Creating a left degenerate binary tree
	input2 = reverseArray(input2)
	tree3 := new(EduBinaryTree)
	tree3.BinaryTreeInitWithSlice(input2)

	// Creating a single node binary tree
	tree4 := new(EduBinaryTree)
	tree4.BinaryTreeInitWithData(100)

	roots := []*EduBinaryTreeNode{tree1.root, tree2.root, tree3.root, tree4.root, nil}

	for i := range roots {
		fmt.Printf("\n\tLevel order traversal: %s\n", levelOrderTraversal(roots[i]))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
