package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	current := l
	subsequent := l.Next

	for subsequent != nil {
		// This can only happen for the first item on the list
		if data_ref < current.Data {
			newNode := &NodeI{Data: data_ref, Next: current}
			return newNode
		}

		if data_ref < subsequent.Data {
			newNode := &NodeI{Data: data_ref, Next: subsequent}
			current.Next = newNode
			return l
		}

		current = subsequent
		subsequent = current.Next
	}

	if data_ref >= current.Data {
		newNode := &NodeI{Data: data_ref}
		current.Next = newNode
		return l
	}

	return l
}
