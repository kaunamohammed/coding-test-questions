package linkedlist

import "fmt"

type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
	size int
}

func (list *DoublyLinkedList) InsertAtHead(value int) {

	newNode := DoublyNode{value: value}
	newNode.next = list.head

	// theres a head
	if !list.IsEmpty() {
		list.head.previous = &newNode
	} else {
		list.tail = &newNode
	}

	list.head = &newNode
	list.size++
}

func (list *DoublyLinkedList) InsertAtTail(value int) {
	newNode := DoublyNode{value: value}

	newNode.previous = list.tail
	list.tail.next = &newNode
	list.tail = &newNode

}

func (list *DoublyLinkedList) InsertAfter(value int, newValue int) {

	if list.IsEmpty() {
		list.InsertAtHead(newValue)
		return
	}

	newNode := DoublyNode{value: newValue}

	current := list.head

	for current != nil {
		if current.value == value {
			current.next.previous = &newNode
			newNode.next = current.next
			current.next = &newNode
			newNode.previous = current
		}
		current = current.next
	}

	list.size++
}

func (list *DoublyLinkedList) Delete(value int) {

	if list.IsEmpty() {
		return
	}

	current := list.head

	if current.value == value {
		list.deleteAtHead()
		return
	}

	for current != nil {
		if current.value == value {
			current.previous.next = current.next
			if current.next != nil {
				current.next.previous = current.previous
			}
			list.size--
			return
		}
		current = current.next
	}

}

func (list *DoublyLinkedList) deleteAtHead() {
	list.head = list.head.next
	list.size--
}

func (list DoublyLinkedList) Print() {

	current := list.head

	for current.next != nil {
		fmt.Print(current.value, " <-> ")
		current = current.next
	}

	fmt.Println(current.value, "-> nil")

}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.head == nil && list.tail == nil
}

func (list DoublyLinkedList) GetLength() int {
	return list.size
}