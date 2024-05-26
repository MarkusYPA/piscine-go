package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	iterator := n1
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n2

	ListSort(n1)

	return n1
}
