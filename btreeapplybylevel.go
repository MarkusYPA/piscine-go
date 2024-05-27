package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Initialize a queue
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Traverse nodes at the current level
		for i := 0; i < len(queue); i++ {

			// Remove the first node from the queue and put it into a variable
			node := queue[0]
			queue = queue[1:]

			// Apply function to current node
			f(node)

			// Add the left child to the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// Add the right child to the queue
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
}

/*

BTreeApplyByLevel(
04
├── 07
│   ├── 12
│   │   └── 10
│   └── 05
└── 01
    └── 02
        └── 03
)
 prints
       + + +  +
 "04010203071210" instead of
 "0401070205120310"

*/
