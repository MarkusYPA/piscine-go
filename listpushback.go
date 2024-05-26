package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func insertNLback(tail *NodeL, data interface{}) *NodeL {
	nL := &NodeL{Data: data}

	if tail == nil {
		return nL
	} else {
		tail.Next = nL
		return nL
	}
}

func ListPushBack(l *List, data interface{}) {
	insertNLback(l.Tail, data)
}
