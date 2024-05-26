package piscine

// Returns the node at index pos after l
func ListAt(l *NodeL, pos int) *NodeL {
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
