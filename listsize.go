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

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	current := l.Head

	for current != nil {
		count++
		current = current.Next
	}

	return count
}
