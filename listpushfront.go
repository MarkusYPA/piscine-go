package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		// If the list is empty, set both Head and Tail to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, append the new node to the start of the list
		newNode.Next = l.Head
		l.Head = newNode
	}
}
