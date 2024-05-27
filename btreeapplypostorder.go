package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)
	BTreeApplyInorder(root.Left, f)
	BTreeApplyInorder(root.Right, f)
}
