package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Tail != nil {
		if l2.Head != nil {
			l1.Tail.Next = l2.Head
		}
	}

	if l1 == nil && l2 != nil {
		l1 = l2
	}

	if l1 != nil && l2 == nil {
		l2 = l1
	}
}
