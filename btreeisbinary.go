package piscine

func BTreeIsBinary(root *TreeNode) bool {
	thisIsBin := true
	leftIsBin := true
	rightIsBin := true

	if root.Left != nil {
		leftIsBin = BTreeIsBinary(root.Left)
		thisIsBin = root.Data > root.Left.Data
	}
	if root.Right != nil {
		rightIsBin = BTreeIsBinary(root.Right)
		thisIsBin = root.Data <= root.Right.Data
	}

	if !thisIsBin || !leftIsBin || !rightIsBin {
		return false
	}

	return true
}
