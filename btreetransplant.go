package piscine

// In the tree root, replaces the subtree node with rplc
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	toReplace := BTreeSearchItem(root, node.Data)
	toReplace.Left = rplc.Left
	toReplace.Right = rplc.Right
	toReplace.Data = rplc.Data
	// Don't change the parent

	return root
}
