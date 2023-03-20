package main

type EduBinaryTree struct {
	root *EduBinaryTreeNode
}

func (t *EduBinaryTree) BinaryTreeInit() {
	t.root = nil
}

/*
BinaryTreeInitWithData initializes the data members

	of BinaryTree using the given data
*/
func (t *EduBinaryTree) BinaryTreeInitWithData(data int) {
	t.root = NewBinaryTreeNode(data)
}

/* BinaryTreeInitWithSlice initializes the data members
   of EduBinaryTree using the given data slice */
func (t *EduBinaryTree) BinaryTreeInitWithSlice(data []int) {
	t.root = nil
	for _, v := range data {
		t.Insert(v)
	}
}

/* Insert inserts an integer into the binary search tree */
func (t *EduBinaryTree) Insert(nodeData int) {
	newNode := NewBinaryTreeNode(nodeData)
	if t.root == nil {
		t.root = newNode
	} else {
		var parent *EduBinaryTreeNode = nil
		tempNode := t.root

		for tempNode != nil {
			parent = tempNode
			if nodeData <= tempNode.data {
				tempNode = tempNode.left
			} else {
				tempNode = tempNode.right
			}
		}

		if nodeData <= parent.data {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}
}

/* FindInBSTRec is a helper function used by FindInBST to
   find the given data in the binary search tree */
func (t *EduBinaryTree) FindInBSTRec(node *EduBinaryTreeNode, nodeData int) *EduBinaryTreeNode {
	if node == nil {
		return nil
	}

	if node.data == nodeData {
		return node
	} else if node.data > nodeData {
		return t.FindInBSTRec(node.left, nodeData)
	} else {
		return t.FindInBSTRec(node.right, nodeData)
	}
}

/* FindInBST is used to find the given data in the binary search tree */
func (t *EduBinaryTree) FindInBST(nodeData int) *EduBinaryTreeNode {
	return t.FindInBSTRec(t.root, nodeData)
}

// GetTreeDeepCopyRec is a helper function used by GetTreeDeepCopy
func (t *EduBinaryTree) GetTreeDeepCopyRec(node *EduBinaryTreeNode) *EduBinaryTreeNode {
	if node != nil {
		newNode := &EduBinaryTreeNode{data: node.data}
		newNode.left = t.GetTreeDeepCopyRec(node.left)
		newNode.right = t.GetTreeDeepCopyRec(node.right)
		return newNode
	} else {
		return nil
	}
}

// GetTreeDeepCopy is used to make a deep copy of the BinaryTree
func (t *EduBinaryTree) GetTreeDeepCopy() *EduBinaryTreeNode {

	if t.root == nil {
		return nil
	} else {
		treeCopy := EduBinaryTree{}
		treeCopy.root = t.GetTreeDeepCopyRec(t.root)
		return treeCopy.root
	}
}
