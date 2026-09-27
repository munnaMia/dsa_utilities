package linkedlist

import (
	"fmt"
)

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// /*
// 	Insertion ------------------------------------------------------------
// */

// append element on the beginning on singly linked list
func (dll *DoublyLinkedList) InsertAtHead(data any) {
	defer dll.incrementCounter()

	// create a new node
	newNode := &Node{
		Prev: nil,
		Data: data,
		Next: dll.head,
	}

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	dll.head.Prev = newNode
	dll.head = newNode
}

// push element on the end
func (dll *DoublyLinkedList) InsertAtTail(data any) {
	defer dll.incrementCounter()

	if dll.IsEmpty() {
		dll.head = &Node{
			Prev: nil,
			Data: data,
			Next: nil,
		}
		dll.tail = dll.head
		return
	}

	dll.tail.Next = &Node{
		Prev: dll.tail,
		Data: data,
		Next: nil,
	}

	dll.tail = dll.tail.Next
}

// Adds a node at a specific position.
func (dll *DoublyLinkedList) InsertAt(index int, data any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}
	if index > dll.length-1 {
		fmt.Println("index out of range")
		return
	}

	if index == 0 {
		newNode := &Node{
			Prev: nil,
			Data: data,
			Next: dll.head,
		}
		dll.head.Prev = newNode
		dll.head = newNode
		dll.incrementCounter()
		return
	}

	counter := 1

	current := dll.head.Next // start from index 1

	for current != nil {
		if counter == index {
			newNode := &Node{
				Prev: current.Prev,
				Data: data,
				Next: current,
			}
			current.Prev.Next = newNode
			current.Prev = newNode
			dll.incrementCounter()
			return
		}

		current = current.Next
		counter++
	}

	for counter <= index {
		dll.tail.Next = &Node{}
		dll.tail = dll.tail.Next
		dll.incrementCounter()

		if counter == index {
			dll.tail.Data = data
			return
		}
		counter++
	}
}

// Inserts new data right after a specific existing value.
func (dll *DoublyLinkedList) InsertAfter(targetData any, data any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	targetNode, _ := dll.Search(targetData)

	if targetNode == nil {
		return
	}

	// insert at tail
	if dll.tail == targetNode {
		dll.InsertAtTail(data)
		dll.incrementCounter()
		return
	}

	newNode := &Node{
		Prev: targetNode,
		Data: data,
		Next: targetNode.Next,
	}

	targetNode.Next.Prev = newNode
	targetNode.Next = newNode

	dll.incrementCounter()
}

// Inserts new data right before a specific existing value.
func (dll *DoublyLinkedList) InsertBefore(targetData any, data any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	targetNode, _ := dll.Search(targetData)

	if targetNode == nil {
		return
	}

	// insert at head
	if dll.head == targetNode {
		dll.InsertAtHead(data)
		dll.incrementCounter()
		return
	}

	newNode := &Node{
		Prev: targetNode.Prev,
		Data: data,
		Next: targetNode,
	}

	targetNode.Prev.Next = newNode
	targetNode.Prev = newNode
	dll.incrementCounter()
}

// /*
// 	Deletation ------------------------------------------------------------
// */

// delete first matched element and return the deleted element
func (dll *DoublyLinkedList) Delete(data any) (bool, any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}
	targetNode, _ := dll.Search(data)

	if dll.length == 1 {
		if targetNode.Data == dll.head.Data {
			oldData := dll.head.Data
			dll.head = nil
			dll.tail = nil
			dll.decrementCounter()
			return true, oldData
		}
	}

	if dll.head == targetNode {
		oldData := dll.head.Data
		dll.head = dll.head.Next
		dll.head.Prev = nil
		dll.decrementCounter()
		return true, oldData
	}

	if dll.tail == targetNode {
		oldData := dll.tail.Data
		dll.tail = dll.tail.Prev
		dll.tail.Next = nil
		dll.decrementCounter()
		return true, oldData
	}

	oldData := targetNode.Data
	preNode := targetNode.Prev
	postNode := targetNode.Next

	preNode.Next = postNode
	postNode.Prev = preNode
	dll.decrementCounter()
	return true, oldData
}

// delete head node.
func (dll *DoublyLinkedList) DeleteHead() (bool, any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}
	oldData := dll.head.Data

	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
		dll.decrementCounter()
		return true, oldData
	}

	dll.head = dll.head.Next
	dll.head.Prev = nil
	dll.decrementCounter()
	return true, oldData
}

// delete tail node.
func (dll *DoublyLinkedList) DeleteTail() (bool, any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}
	oldData := dll.tail.Data

	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
		dll.decrementCounter()
		return true, oldData
	}

	dll.tail = dll.tail.Prev
	dll.tail.Next = nil
	dll.decrementCounter()

	return true, oldData
}

// Removes a node based on its numerical position.
func (dll *DoublyLinkedList) DeleteAt(index int) (bool, any) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}

	if index > dll.length-1 {
		fmt.Println("index out of range")
		return false, nil
	}

	counter := 0

	current := dll.head

	if dll.length == 1 {
		if counter == index {
			oldData := dll.head.Data
			dll.head = nil
			dll.tail = nil
			dll.decrementCounter()
			return true, oldData
		}
	}

	if index == 0 {
		return dll.DeleteHead()
	}

	if index == dll.length-1 {
		return dll.DeleteTail()
	}

	for counter <= index {
		if counter == index {
			oldData := current.Data
			preNode := current.Prev
			postNode := current.Next

			preNode.Next = postNode
			postNode.Prev = preNode

			dll.decrementCounter()
			return true, oldData
		}
		counter++
		current = current.Next
	}

	return false, nil
}

// Keeps the first $n$ elements and deletes the rest.
func (dll *DoublyLinkedList) Truncate(n int) error {
	if n == 0 || n < 0 {
		return nil
	}

	if dll.Length() < n {
		return fmt.Errorf("not enough elements in the linked list")
	}

	currentNode, err := dll.GetAt(n - 1)
	if err != nil {
		return err
	}

	currentNode.Next = nil
	dll.tail = currentNode

	return nil
}

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// show the head node
func (dll *DoublyLinkedList) GetHead() (*Node, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return dll.head, nil
}

// show the head node value
func (dll *DoublyLinkedList) GetHeadData() (any, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return dll.head.Data, nil
}

// show the tail node
func (dll *DoublyLinkedList) GetTail() (*Node, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return dll.tail, nil
}

// show the tail node value
func (dll *DoublyLinkedList) GetTailData() (any, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return dll.tail.Data, nil
}

// get an element of an given index and a bool status that the index exist or not
func (dll *DoublyLinkedList) GetAt(index int) (*Node, error) {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := dll.head
	counter := 0

	for currentNode != nil && counter < dll.length {
		if counter == index {
			return currentNode, nil
		}
		currentNode = currentNode.Next
		counter++
	}

	return nil, fmt.Errorf("linked list is empty")

}

// search an element on linked list and return boolean
func (dll *DoublyLinkedList) Search(data any) (*Node, error) {
	if dll.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := dll.head

	for currentNode != nil {
		if currentNode.Data == data {
			return currentNode, nil
		}
		currentNode = currentNode.Next
	}

	return nil, fmt.Errorf("element not found")
}

// Returns a simple true/false if the value is in the list.
func (dll *DoublyLinkedList) Contains(data any) bool {
	_, err := dll.Search(data)

	if err != nil {
		return false
	}

	return true
}

// /*
// 	Transformation Methods ------------------------------------------------------------
// */

// Replaces a specific value with a new one.
func (dll *DoublyLinkedList) Update(data, replace any) error {
	if dll.IsEmpty() {
		return fmt.Errorf("Linked list is empty.")
	}

	targetNode, err := dll.Search(data)
	if err != nil {
		return err
	}

	targetNode.Data = replace
	return nil
}

// reverse the linked list
func (dll *DoublyLinkedList) Reverse() {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	current := dll.head

	for current != nil {
		temp := current.Prev
		current.Prev = current.Next
		current.Next = temp

		current = current.Prev
	}
	temp := dll.head

	dll.head = dll.tail
	dll.tail = temp
}

// Scans the list and removes nodes with repeating values
func (dll *DoublyLinkedList) RemoveDuplicates() {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	seen := make(map[any]bool)

	current := dll.head
	seen[current.Data] = true

	for current.Next != nil {
		if seen[current.Next.Data] {
			current.Next = current.Next.Next
		} else {
			seen[current.Next.Data] = true
			current = current.Next
		}

	}
}

// covert the linked list into slice
func (dll *DoublyLinkedList) ToSlice() []any {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return []any{}
	}

	sllSlice := []any{}
	currentNode := dll.head
	for currentNode != nil {
		sllSlice = append(sllSlice, currentNode.Data)

		currentNode = currentNode.Next
	}

	return sllSlice
}

// /*
// 	Metadata & Utility Methods ------------------------------------------------------------
// */

// tell how many element the linked list have
func (dll *DoublyLinkedList) Length() int {
	return dll.length
}

// check the linked list is empty or not
func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head == nil
}

// Print the single linked list
func (dll *DoublyLinkedList) PrintList() {
	if dll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	currentNode := dll.head

	for currentNode != nil {
		fmt.Println("Data :", currentNode.Data)
		currentNode = currentNode.Next
	}
}

// clear the whole linked list
func (dll *DoublyLinkedList) Clear() {
	dll.head = nil
	dll.tail = nil
}

// /*
// 	private helper methods --------------------------------------------------------------------
// */

// increment after eash deletation
func (dll *DoublyLinkedList) incrementCounter() {
	dll.length++ // just increate by one
}

// decrement after eash deletation
func (dll *DoublyLinkedList) decrementCounter() {
	dll.length--
}
