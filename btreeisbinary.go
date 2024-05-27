package piscine

func BTreeIsBinary(root *TreeNode) bool {
	rightIsBin := true
	leftIsBin := true

	if root == nil {
		return true
	}

	if root.Left != nil {
		return root.Data > root.Left.Data
	}
	if root.Right != nil {
		return root.Data <= root.Right.Data
	}

	leftIsBin = BTreeIsBinary(root.Left)
	rightIsBin = BTreeIsBinary(root.Right)

	if !leftIsBin || !rightIsBin {
		return false
	}

	return true
}
