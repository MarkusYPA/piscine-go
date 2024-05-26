package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		nd := &NodeI{Data: data_ref}
		return nd
	}

	current := l
	following := l.Next

	if following == nil {
		l.Next.Data = data_ref
		return l
	}

	if data_ref <= l.Data {
		newNd := &NodeI{Data: data_ref, Next: l}
		return newNd
	}

	for following != nil {
		if current.Data <= data_ref && following.Data >= data_ref {
			newNd := &NodeI{Data: data_ref, Next: following}
			current.Next = newNd
			current = newNd
			return l
		}
		current = current.Next
		following = current.Next
	}

	if data_ref <= current.Data { // Current is now the last on the list
		newNd := &NodeI{Data: data_ref}
		current.Next = newNd
	}
	return l
}
