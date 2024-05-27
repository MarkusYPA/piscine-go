package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	levels := BTreeLevelCount(root)
	applyByLevel(root, f, levels)
}

func applyByLevel(node *TreeNode, f func(...interface{}) (int, error), level int) {
	if node == nil {
		return
	}
	if level == BTreeLevelCount(node) {
		f(node.Data)
	}
	level--
	applyByLevel(node.Left, f, level)
	applyByLevel(node.Right, f, level)
}
