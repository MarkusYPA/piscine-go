package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Tail != nil {
		if l2.Head != nil {
			l1.Tail.Next = l2.Head
		}
		return
	}

	if l1 == nil && l2 != nil {
		l1 = &List{l2.Head, l2.Tail}
		return
	}

	if l1 != nil && l2 == nil {
		l2 = &List{l1.Head, l1.Tail}
	}
}
