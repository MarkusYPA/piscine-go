package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil} // "&"" because Head, Tail and Next use pointers

	if l.Head == nil {
		// If the list is empty, set both Head and Tail to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, append the new node to the end of the list
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
