package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l
	size := 0
	for current != nil {
		size++
		current = current.Next
	}

	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if NodeIat(l, i).Data > NodeIat(l, j).Data {
				NodeIat(l, i).Data, NodeIat(l, j).Data = NodeIat(l, j).Data, NodeIat(l, i).Data
			}
		}
	}

	return l
}

// Returns the node at index pos after l
func NodeIat(l *NodeI, pos int) *NodeI {
	count := 0
	current := l // Head

	for current != nil && count < pos {
		count++
		current = current.Next
	}

	if current == nil {
		return nil
	}

	return current
}
