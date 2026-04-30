package linkedlist

import "fmt"

type singlyLinkedList struct {
	head  *Node
	tail  *Node
	count int
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{
		count: 0,
	}
}

// append element on the end on singly linked list
func (sll *singlyLinkedList) AppendItem(data int) {
	defer sll.incrementCounter()

	// if the linked list is empty
	if sll.head == nil {
		sll.head = &Node{
			data: data,
		}

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
	defer sll.incrementCounter()

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

// search an element on linked list and return boolean
func (sll *singlyLinkedList) Search(data int) bool {
	tempNode := sll.head

	if tempNode == nil {
		fmt.Println("Linked list is empty.")
		return false
	}

	for tempNode != nil {
		if tempNode.data == data {
			return true
		}
		tempNode = tempNode.next
	}

	return false
}

// search an element on linked list and update with the replace data if exist and update the first match data.
// return the old data and success status
func (sll *singlyLinkedList) Update(data, replace int) (bool, int) {
	tempNode := sll.head

	if tempNode == nil {
		fmt.Println("Linked list is empty.")
		return false, 0
	}

	for tempNode != nil {
		if tempNode.data == data {
			oldData := tempNode.data
			tempNode.data = replace

			return true, oldData
		}
		tempNode = tempNode.next
	}

	return false, 0
}

// tell how many element the linked list have
func (sll *singlyLinkedList) Count() int {
	return sll.count
}

func (sll *singlyLinkedList) incrementCounter() {
	sll.count++
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
