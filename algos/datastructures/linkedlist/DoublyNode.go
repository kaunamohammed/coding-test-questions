package linkedlist

type DoublyNode struct {
	value int
	previous *DoublyNode
	next *DoublyNode
}

func (node DoublyNode) GetValue()int {
	return node.value
}

func (node DoublyNode) GetPrevious() DoublyNode {
	return *node.previous
}

func (node DoublyNode) GetNext() DoublyNode {
	return *node.next
}

func (node *DoublyNode) SetPrevious(previous DoublyNode) {
	node.previous = &previous
}

func (node *DoublyNode) SetNext(next DoublyNode) {
	node.next = &next
}