package main

import (
	"fmt"
	"strings"
)

func zigzagLevelOrder(root *EduBinaryTreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	results := make([][]int, 0)
	reverse := false
	dq := new(Dequeue)
	dq.PushFront(root)

	for !dq.Empty() {
		size := dq.Len()
		results = append(results, make([]int, 0))

		for i := 0; i < size; i++ {
			if !reverse {
				node := dq.PopFront()
				results[len(results)-1] = append(results[len(results)-1], node.data)
				if node.left != nil {
					dq.PushBack(node.left)
				}
				if node.right != nil {
					dq.PushBack(node.right)
				}
			} else {
				node := dq.PopBack()
				results[len(results)-1] = append(results[len(results)-1], node.data)
				if node.right != nil {
					dq.PushFront(node.right)
				}
				if node.left != nil {
					dq.PushFront(node.left)
				}
			}
		}
		reverse = !reverse
	}

	return results
}

// Driver code
func main() {
	inputTrees := [][]int{
		{100, 50, 200, 25, 75, 350},
		{25, 50, 75, 100, 200, 350},
		{350, 75, 25, 200, 50, 100},
		{100},
		nil,
	}
	for i, input := range inputTrees {
		testCasetree := new(EduBinaryTree)
		testCasetree.BinaryTreeInitWithSlice(input)
		fmt.Printf("%d.\tBinary Tree:\n", i+1)
		fmt.Printf("\n\tThe zigzag level order traversal is: %s\n", strings.Replace(fmt.Sprint(zigzagLevelOrder(testCasetree.root)), " ", ", ", -1))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
