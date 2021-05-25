package linkedlist

type SinglyNode struct {
	value int
	next  *SinglyNode
}

func (node SinglyNode) GetValue() int {
	return node.value
}

func (node SinglyNode) GetNext() SinglyNode {
	return *node.next
}

func (node *SinglyNode) SetNext(next SinglyNode) {
	node.next = &next
}
