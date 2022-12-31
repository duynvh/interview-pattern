package main

type BinaryTree struct {
    root *BinaryTreeNode
}

func (t *BinaryTree) BinaryTreeInit() {
    t.root = nil
}

func (t *BinaryTree) BinaryTreeInitWithData(data int) {
    t.root = NewBinaryTreeNode(data)
}

func (t *BinaryTree) Insert(nodeData int) {
    newNode := NewBinaryTreeNode(nodeData)

    if t.root == nil {
        t.root = newNode
    } else {
        var parent *BinaryTreeNode = nil
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

func (t *BinaryTree) FindInBSTRec(node *BinaryTreeNode, nodeData int) *BinaryTreeNode {
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

func (t *BinaryTree) FindInBST(nodeData int) *BinaryTreeNode {
    return t.FindInBSTRec(t.root, nodeData)
}

func (t *BinaryTree) GetTreeDeepCopyRec(node *BinaryTreeNode) *BinaryTreeNode {
    if node != nil {
        newNode := &BinaryTreeNode{data: node.data}
        newNode.left = t.GetTreeDeepCopyRec(node.left)
        newNode.right = t.GetTreeDeepCopyRec(node.right)
        return newNode
    } else {
        return nil
    }
}

func (t *BinaryTree) GetTreeDeepCopy() *BinaryTreeNode {
    if t.root == nil {
        return nil
    } else {
        treeCopy := BinaryTree{}
        treeCopy.root = t.GetTreeDeepCopyRec(t.root)
        return treeCopy.root
   }
}