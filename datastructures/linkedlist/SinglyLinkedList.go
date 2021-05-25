package linkedlist

import "fmt"

type SinglyLinkedList struct {
	head *SinglyNode
	size int
}

func (list *SinglyLinkedList) Reverse() {
	
}

func (list *SinglyLinkedList) Search(value int) bool {

	current := list.head

	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}

	return false
}

func (list *SinglyLinkedList) InsertAtHead(value int) {
	newNode := SinglyNode{value: value}
	newNode.next = list.head
	list.head = &newNode
	list.size++
}

func (list *SinglyLinkedList) InsertAtTail(value int) {
	newNode := SinglyNode{value: value}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = &newNode

	list.size++
}

func (list *SinglyLinkedList) InsertAfter(value int, newValue int) {
	newNode := SinglyNode{value: newValue}

	current := list.head

	for current != nil && current.GetValue() != value {
		current = current.next
	}

	if current != nil {
		newNode.SetNext(current.GetNext())
		current.SetNext(newNode)
		return
	}

	list.size++
}

func (list *SinglyLinkedList) Delete(value int) {

	if list.IsEmpty() {
		return
	}

	if list.head.value == value {
		list.deleteAtHead()
		return
	}

	current := list.head
	var previous *SinglyNode

	for current != nil {
		if current.value == value {
			previous.next = current.next
			return
		}
		previous = current
		current = current.next
	}
	list.size--
}

func (list *SinglyLinkedList) deleteAtHead() {
	list.head = list.head.next
	list.size--
}

func (list SinglyLinkedList) Print() {

	current := list.head

	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}

	fmt.Println("nil")

}

func (list *SinglyLinkedList) IsEmpty() bool {
	return list.head == nil 
}
