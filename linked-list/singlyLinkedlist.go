package linkedlist

import (
	"fmt"
)

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

/*
	Insertion ------------------------------------------------------------
*/

// append element on the beginning on singly linked list
func (sll *SinglyLinkedList) InsertAtHead(data any) {
	defer sll.incrementCounter()

	newNode := &Node{
		data: data,
		next: sll.head,
	}

	sll.head = newNode
	if sll.IsEmpty() {
		sll.tail = newNode
	}
}

// push element on the end
func (sll *SinglyLinkedList) InsertAtTail(data any) {
	defer sll.incrementCounter()

	newNode := &Node{
		data: data,
		next: nil,
	}

	if sll.IsEmpty() {
		sll.head = newNode
		sll.tail = newNode
		return
	}

	sll.tail.next = newNode
	sll.tail = newNode
}

// Adds a node at a specific position or append empty node.
func (sll *SinglyLinkedList) InsertAt(index int, data any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	if index > sll.length-1 {
		fmt.Println("Out of index bound")
		return
	}

	// if the head match first
	if index == 0 {
		sll.InsertAtHead(data)
		return
	}

	previous := sll.head
	current := sll.head.next

	counter := 1 // as first index already check and current point start with the secound element

	for current != nil {

		if counter == index {
			newNode := &Node{
				data: data,
				next: current,
			}
			previous.next = newNode
			sll.incrementCounter()
			return
		}

		previous = current
		current = current.next
		counter++
	}
}

// Inserts new data right after a specific existing value. if nothing match it do nothing just like u.
func (sll *SinglyLinkedList) InsertAfter(targetData any, data any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	targetNode, err := sll.Search(targetData)
	if err != nil {
		fmt.Println("targeted node not found, err:", err)
		return
	}

	if sll.tail == targetNode {
		sll.InsertAtTail(data)
		return
	}

	tempNode := targetNode.next
	targetNode.next = &Node{
		data: data,
		next: tempNode,
	}
	sll.incrementCounter()
}

// Inserts new data right after a specific existing value.
func (sll *SinglyLinkedList) InsertBefore(targetData any, data any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	previous := sll.head
	current := sll.head.next

	// if the head match first
	if previous.data == targetData {
		sll.InsertAtHead(targetData)
		return
	}

	for current != nil {
		if current.data == targetData {
			previous.next = &Node{
				data: data,
				next: current,
			}
			sll.incrementCounter()
			return
		}

		// move the pointer
		previous = current
		current = current.next
	}

	fmt.Println("targeted node not found")
}

// /*
// 	Deletation ------------------------------------------------------------
// */

// delete first matched element and return the deleted element. false mean data not exist or list is empty
func (sll *SinglyLinkedList) Delete(data any) (bool, any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, nil
	}

	previous := sll.head
	current := sll.head.next

	// if the head match first
	if previous.data == data {
		temp := sll.head.data
		sll.head = sll.head.next
		sll.decrementCounter()
		return true, temp
	}

	for current != nil {
		if current.data == data {
			temp := current.data
			previous.next = current.next
			sll.decrementCounter()
			return true, temp
		}
		// move the pointer
		previous = current
		current = current.next
	}

	return false, nil
}

// delete head node.
func (sll *SinglyLinkedList) DeleteHead() (bool, any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, nil
	}

	temp := sll.head.data
	sll.head = sll.head.next
	sll.decrementCounter()
	return true, temp
}

// delete tail node.
func (sll *SinglyLinkedList) DeleteTail() (bool, any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, nil
	}

	temp := sll.tail.data

	// if only one node available
	if sll.Length() == 1 {
		sll.head = nil
		sll.tail = nil
		sll.decrementCounter()
		return true, temp
	}

	current := sll.head

	for current.next != nil {
		if current.next == sll.tail {
			current.next = nil
			sll.tail = current
			sll.decrementCounter()
			return true, temp
		}

		current = current.next
	}

	return false, nil
}

// Removes a node based on its numerical position.
func (sll *SinglyLinkedList) DeleteAt(index int) (bool, any) {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, nil
	}

	if index == 0 {
		return sll.DeleteHead()
	}

	previous := sll.head
	current := sll.head.next
	currentIdx := 1

	for current != nil {
		if currentIdx == index {
			temp := current.data
			previous.next = current.next
			sll.decrementCounter()
			return true, temp
		}

		currentIdx++

		previous = current
		current = current.next
	}

	return false, nil
}

// Keeps the first $n$ elements and deletes the rest.
func (sll *SinglyLinkedList) Truncate(n int) error {
	if n == 0 {
		return nil
	}

	if sll.Length() < n {
		return fmt.Errorf("not enough elements in the linked list")
	}

	currentNode, err := sll.GetAt(n - 1)
	if err != nil {
		return err
	}

	currentNode.next = nil
	sll.tail = currentNode

	return nil
}

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// show the head node value
func (sll *SinglyLinkedList) GetHead() (*Node, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return sll.head, nil
}

// show the head node value
func (sll *SinglyLinkedList) GetHeadData() (any, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return sll.head.data, nil
}

// show the tail node
func (sll *SinglyLinkedList) GetTail() (*Node, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return sll.tail, nil
}

// show the tail node value
func (sll *SinglyLinkedList) GetTailData() (any, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return sll.tail.data, nil
}

// get an node by index
func (sll *SinglyLinkedList) GetAt(index int) (*Node, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := sll.head
	counter := 0

	for currentNode != nil && counter < sll.length {
		if counter == index {
			return currentNode, nil
		}
		currentNode = currentNode.next
		counter++
	}

	return nil, fmt.Errorf("index not available")
}

// search an element on linked list and return boolean
func (sll *SinglyLinkedList) Search(data any) (*Node, error) {
	if sll.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := sll.head

	for currentNode != nil {
		if currentNode.data == data {
			return currentNode, nil
		}
		currentNode = currentNode.next
	}

	return nil, fmt.Errorf("element not found")
}

// Returns a simple true/false if the value is in the list.
func (sll *SinglyLinkedList) Contains(data any) bool {
	if _, err := sll.Search(data); err != nil {
		return false
	}

	return true
}

/*
	Transformation Methods ------------------------------------------------------------
*/

// Replaces a specific value with a new one.
func (sll *SinglyLinkedList) Update(data, replace any) error {
	if sll.IsEmpty() {
		return fmt.Errorf("Linked list is empty.")
	}

	targetNode, err := sll.Search(data)
	if err != nil {
		return err
	}

	targetNode.data = replace

	return nil
}

// reverse the linked list
func (sll *SinglyLinkedList) Reverse() {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	tempLinkedList := NewSinglyLinkedList()

	currentNode := sll.head

	for currentNode != nil {
		tempLinkedList.InsertAtHead(currentNode.data)
		currentNode = currentNode.next
	}

	sll.head = tempLinkedList.head
	sll.tail = tempLinkedList.tail
}

// Scans the list and removes nodes with repeating values
func (sll *SinglyLinkedList) RemoveDuplicates() {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	seen := make(map[any]bool)

	current := sll.head
	seen[current.data] = true

	for current.next != nil {
		if seen[current.next.data] {
			current.next = current.next.next
		} else {
			seen[current.next.data] = true
			current = current.next

		}

	}
}

// covert the linked list into slice
func (sll *SinglyLinkedList) ToSlice() []any {
	if sll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return []any{}
	}

	sllSlice := []any{}
	currentNode := sll.head
	for currentNode != nil {
		sllSlice = append(sllSlice, currentNode.data)

		currentNode = currentNode.next
	}

	return sllSlice
}

// covert the slice into linked list
func (sll *SinglyLinkedList) ToLinkedlist(s []any) *SinglyLinkedList {
	newll := NewSinglyLinkedList()

	for _, v := range s {
		newll.InsertAtTail(v)
	}

	return newll
}

/*
	Metadata & Utility Methods ------------------------------------------------------------
*/

// tell how many element the linked list have
func (sll *SinglyLinkedList) Length() int {
	return sll.length
}

// check the linked list is empty or not
func (sll *SinglyLinkedList) IsEmpty() bool {
	return sll.length == 0
}

// Print the single linked list
func (sll *SinglyLinkedList) PrintList() {
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

// clear the whole linked list
func (sll *SinglyLinkedList) Clear() {
	sll.head = nil
	sll.tail = nil
}

/*
	private helper methods --------------------------------------------------------------------
*/

// increment after eash insertion
func (sll *SinglyLinkedList) incrementCounter() {
	sll.length++
}

// decrement after eash deletation
func (sll *SinglyLinkedList) decrementCounter() {
	sll.length--
}
