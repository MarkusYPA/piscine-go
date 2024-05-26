package piscine

func ListReverse(l *List) {
	lastInd := ListSize(l) - 1
	halfLen := ListSize(l) / 2

	for i := 0; i < halfLen; i++ {
		ListAt(l.Head, i).Data, ListAt(l.Head, lastInd-i).Data = ListAt(l.Head, lastInd-i).Data, ListAt(l.Head, i).Data
	}
}
