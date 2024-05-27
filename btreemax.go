package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	maxNode := root

	for current != nil {
		// checking if current is the biggest so far
		if current.Data > maxNode.Data {
			maxNode = current
		}

		left := current.Left
		right := current.Right

		// Making current the bigger (or non-nil) of the two, or nil.
		if left != nil && right != nil {
			if left.Data > right.Data { // Right should be bigger
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
	return maxNode
}
