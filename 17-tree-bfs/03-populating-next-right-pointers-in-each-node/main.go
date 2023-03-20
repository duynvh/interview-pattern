package main

func connectNextLevel(head *EduBinaryTreeNode) *EduBinaryTreeNode {
    current := head
    nextLevelHead, prev := new(EduBinaryTreeNode), new(EduBinaryTreeNode)
    nextLevelHead, prev = nil, nil

    for current != nil {
        if current.left != nil && current.right != nil {
            if nextLevelHead == nil {
                nextLevelHead = current.left
            }

            current.left.next = current.right

            if prev != nil {
                prev.next = current.left
            }

            prev = current.right
        } else if current.left != nil {
            if nextLevelHead == nil {
				nextLevelHead = current.left
			}
			if prev != nil {
				prev.next = current.left
			}

			prev = current.left
        } else if current.right != nil {
            			// If only the right child node children is not null
			// then only connect it with previous same level nodes
			if nextLevelHead == nil {
				nextLevelHead = current.right
			}
			if prev != nil {
				prev.next = current.right
			}
			prev = current.right
        }

        // Update current pointer
		current = current.next
    }

    // Pointing the last node (right-most node) of the next level
	// to null
	if prev != nil {
		prev.next = nil
	}

	// Return the head node (left-most node) of the next level
	return nextLevelHead
}