package linkedlist

import "fmt"

type singlyLinkedList struct {
	head *Node
	tail *Node
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

// append element on the end on singly linked list
func (sll *singlyLinkedList) AppendItem(data int) {

	// if the linked list is empty
	if sll.head == nil {
		// init head
		sll.head = &Node{
			data: data,
		}

		// init tail
		sll.tail = sll.head
		return
	}

	// storing element on the tail
	sll.tail.next = &Node{
		data: data,
	}

	// move the tail to the end
	sll.tail = sll.tail.next
}

// push element on the beginning
func (sll *singlyLinkedList) ShiftItem(data int) {
	newNode := &Node{
		data: data,
		next: sll.head,
	}

	sll.head = newNode
}

// show the head node value
func (sll *singlyLinkedList) ShowHead() (int, error) {
	tempHead := sll.head

	if tempHead == nil {
		return 0, fmt.Errorf("Linked list is empty.")
	}
	return sll.head.data, nil
}

// show the tail node value
func (sll *singlyLinkedList) ShowTail() (int, error) {
	tempTail := sll.tail

	if tempTail == nil {
		return 0, fmt.Errorf("Linked list is empty.")
	}
	return sll.tail.data, nil
}

// Print the single linked list
func (sll *singlyLinkedList) PrintList() {
	tempHead := sll.head

	if tempHead == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	for tempHead != nil {
		fmt.Println("Data :", tempHead.data)
		tempHead = tempHead.next
	}
}
