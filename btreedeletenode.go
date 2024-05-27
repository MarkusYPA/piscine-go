package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	// Save the data from deletee's right
	dataToMove := []string{}
	dataToMove = getTreeData(node.Right, dataToMove)

	// Replace deletee with its left
	root = BTreeTransplant(root, node, node.Left)

	for _, st := range dataToMove {
		BTreeInsertData(root, st)
	}

	return root
}

func getTreeData(root *TreeNode, data []string) []string {
	if root == nil {
		return data
	}
	data = append(data, root.Data)
	getTreeData(root.Left, data)
	getTreeData(root.Right, data)
	return data
}
