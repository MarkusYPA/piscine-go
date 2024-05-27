package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	BTreeApplyInorder(root.Right, f)
	f(root.Data)
}

/*
prints
"01\na\nb\nc\nd\nq\nr\nt\nx\ny\nz\n08\n" instead of
"01\na\nd\nc\nb\nq\nr\nt\ny\nz\nx\n08\n"
*/
