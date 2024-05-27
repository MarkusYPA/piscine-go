package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Initialize a queue
	queue := []TreeNode{*root}

	for len(queue) > 0 {
		// Traverse nodes at the current level
		for i := 0; i < len(queue); i++ {

			// Remove the first node from the queue and put it into a variable
			node := queue[0]
			queue = queue[1:]

			// Apply function to current node
			f(node.Data)

			// Add the left child to the queue
			if node.Left != nil {
				queue = append(queue, *node.Left)
			}

			// Add the right child to the queue
			if node.Right != nil {
				queue = append(queue, *node.Right)
			}
			// fmt.Println("Queue  ---  ", queue)
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
 "&{0xc00010c270 0xc00010c2d0 <nil> 04}&{<nil> 0xc00010c3f0 0xc00010c210 01}&{0xc00010c330 0xc00010c390 0xc00010c210 07}&{<nil> 0xc00010c450 0xc00010c270 02}&{<nil> <nil> 0xc00010c2d0 05}&{0xc00010c4b0 <nil> 0xc00010c2d0 12}&{<nil> <nil> 0xc00010c3f0 03}&{<nil> <nil> 0xc00010c390 10}"
 instead of
 "0401070205120310"
exit status 1

*/
