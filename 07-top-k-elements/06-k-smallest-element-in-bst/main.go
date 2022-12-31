package main

func kthSmallest(root *BinaryTreeNode, k int) int {
	topK := make([]int, 0)

	traversal(root, &topK, k)

	return topK[len(topK)-1]
}

func traversal(root *BinaryTreeNode, topK *[]int, k int) {
	if root == nil {
		return
	}

	traversal(root.left, topK, k)

	if len(*topK) < k {
		*topK = append(*topK, root.data)
	} else {
		return
	}

	traversal(root.right, topK, k)
}
