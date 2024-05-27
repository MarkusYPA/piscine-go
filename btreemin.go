package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	minNode := root

	for current != nil {
		// checking if current is the smallest so far
		if current.Data < minNode.Data {
			minNode = current
		}

		left := current.Left
		right := current.Right

		// Making current the bigger (or non-nil) of the two, or nil.
		if left != nil && right != nil {
			if left.Data < right.Data { // Left should be smaller
				current = left
			} else {
				current = right
			}
		} else if left != nil {
			current = left
		} else if right != nil {
			current = right
		} else {
			current = nil
		}
	}
	return minNode
}
