package linkedlist

import (
	"fmt"
)

type singlyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{
		length: 0,
	}
}

// append element on the end on singly linked list
func (sll *singlyLinkedList) Push(data int) {
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
func (sll *singlyLinkedList) AddFirst(data int) {
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

// tell how many element the linked list have
func (sll *singlyLinkedList) Length() int {
	return sll.length
}

func (sll *singlyLinkedList) incrementCounter() {
	sll.length++
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

// reverse the linked list
func (sll *singlyLinkedList) Reverse() {
	tempLinkedList := NewSinglyLinkedList()

	currentNode := sll.head

	if currentNode == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	for currentNode != nil {
		tempLinkedList.AddFirst(currentNode.data)
		currentNode = currentNode.next
	}

	sll.head = tempLinkedList.head
}

// delete first matched element and return the deleted element
func (sll *singlyLinkedList) Delete(data int) (bool, int) {
	previous := sll.head
	current := sll.head.next

	if previous == nil {
		fmt.Println("Linked list is empty.")
		return false, 0
	}

	// if the head match first
	if previous.data == data {
		sll.head = sll.head.next
		return true, sll.head.data
	}

	for current != nil {
		if current.data == data {
			oldData := current.data
			previous.next = current.next // unlink

			return true, oldData
		}
		// move the pointer
		previous = current
		current = current.next
	}

	return false, 0
}

// delete head node.
func (sll *singlyLinkedList) DeleteHead(data int) (bool, int) {
	currentNode := sll.head

	if currentNode == nil {
		fmt.Println("Linked list is empty.")
		return false, 0
	}

	deletedValue := sll.head.data

	if sll.head.next == nil {
		sll.head = nil
		return true, deletedValue
	}

	sll.head = sll.head.next
	return true, deletedValue
}

// check the linked list is empty or not
func (sll *singlyLinkedList) IsEmpty() bool {
	if sll.head == nil {
		return true
	}
	return false
}

// covert the linked list into slice
func (sll *singlyLinkedList) ToSlice() []int {
	llSlice := []int{}

	tempHead := sll.head

	for tempHead != nil {
		llSlice = append(llSlice, tempHead.data)
		tempHead = tempHead.next
	}

	return llSlice
}
