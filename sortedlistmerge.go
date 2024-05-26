package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil && n2 != nil {
		ListSort(n2)
		return n2
	}
	if n1 != nil && n2 == nil {
		ListSort(n1)
		return n1
	}

	if n1 == nil && n2 == nil {
		return n1
	}

	iterator := n1
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n2

	ListSort(n1)

	return n1
}
