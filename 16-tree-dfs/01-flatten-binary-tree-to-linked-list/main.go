package main

var treeForPrinting *EduBinaryTree

func flattenTree(root *EduBinaryTreeNode) *EduBinaryTreeNode {
	if root == nil {
		return nil
	}

	current := root

	for current != nil {
		if current.left != nil {
			last := current.left

			for last.right != nil {
				last = last.right
			}

			last.right = current.right
			current.right = current.left
			current.left = nil
		}

		current = current.right
	}

	return root
}
