package linkedlist

import "fmt"

type singlyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

/*
	Insertion ------------------------------------------------------------
*/

// append element on the end on singly linked list
func (sll *singlyLinkedList) InsertAtHead(data any) {
	defer sll.incrementCounter()

	newNode := &Node{
		data: data,
		next: sll.head,
	}

	sll.head = newNode
}

// push element on the beginning
func (sll *singlyLinkedList) InsertAtTail(data any) {
	defer sll.incrementCounter()

	if sll.IsEmpty() {
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

// // Adds a node at a specific position.
// func (sll *singlyLinkedList) InsertAt(index int, data any) error

// // Inserts new data right after a specific existing value.
// func (sll *singlyLinkedList) InsertAfter(targetData any, newData any) error

// // Inserts new data right after a specific existing value.
// func (sll *singlyLinkedList) InsertBefore(targetData any, newData any) error

// /*
// 	Deletation ------------------------------------------------------------
// */

// // delete first matched element and return the deleted element
// func (sll *singlyLinkedList) Delete(data any) (bool, any)

// // delete head node.
// func (sll *singlyLinkedList) DeleteHead() (bool, any)

// // delete tail node.
// func (sll *singlyLinkedList) DeleteTail() (bool, any)

// // Removes a node based on its numerical position.
// func (sll *singlyLinkedList) DeleteAt(index int)

// // Keeps the first $n$ elements and deletes the rest.
// func (sll *singlyLinkedList) Truncate(n int)

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// show the head node value
func (sll *singlyLinkedList) GetHead() (any, error) {
	if sll.IsEmpty() {
		return 0, fmt.Errorf("Linked list is empty.")
	}
	return sll.head.data, nil
}

// show the tail node value
func (sll *singlyLinkedList) GetTail() (any, error) {
	if sll.IsEmpty() {
		return 0, fmt.Errorf("Linked list is empty.")
	}
	return sll.tail.data, nil
}

// // get an element of an given index
// func (sll *singlyLinkedList) GetAt(index int)

// // search an element on linked list and return boolean
func (sll *singlyLinkedList) Search(data any) bool {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false
	}

	currentNode := sll.head

	for currentNode != nil {
		if currentNode.data == data {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

// // Returns a simple true/false if the value is in the list.
// func (sll *singlyLinkedList) Contains(data any)

// /*
// 	Transformation Methods ------------------------------------------------------------
// */

// // Replaces a specific value with a new one.
// func (sll *singlyLinkedList) Update(data, replace any) (bool, any)

// // reverse the linked list
// func (sll *singlyLinkedList) Reverse()

// // sort the linked list
// func (sll *singlyLinkedList) Sort()

// // Scans the list and removes nodes with repeating values
// func (sll *singlyLinkedList) RemoveDuplicates()

// // covert the linked list into slice
// func (sll *singlyLinkedList) ToSlice() []any

// /*
// 	Metadata & Utility Methods ------------------------------------------------------------
// */

// // tell how many element the linked list have
// func (sll *singlyLinkedList) Length() int {
// 	return sll.length
// }

// check the linked list is empty or not
func (sll *singlyLinkedList) IsEmpty() bool {
	return sll.head == nil
}

// // Print the single linked list
func (sll *singlyLinkedList) PrintList() {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	currentNode := sll.head

	for currentNode != nil {
		fmt.Println("Data :", currentNode.data)
		currentNode = currentNode.next
	}
}

// // clear the whole linked list
// func (sll *singlyLinkedList) Clear()

/*
	private helper methods --------------------------------------------------------------------
*/

func (sll *singlyLinkedList) incrementCounter() {
	sll.length++ // just increate by one
}
