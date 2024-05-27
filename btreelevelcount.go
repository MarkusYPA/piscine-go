package piscine

func BTreeLevelCount(root *TreeNode) int {
	return levelCounter(root, 0)
}

func levelCounter(node *TreeNode, count int) int {
	if node == nil {
		return count
	}
	count++
	left := levelCounter(node.Left, count)
	right := levelCounter(node.Right, count)
	if left > right {
		return left
	}
	return right
}
