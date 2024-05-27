package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	toReplace := BTreeSearchItem(root, node.Data)
	toReplace.Left = rplc.Left
	toReplace.Right = rplc.Right
	toReplace.Parent = rplc.Parent
	toReplace.Data = rplc.Data

	return root
}
