package linkedlist

import "fmt"

type CircularLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// /*
// 	Insertion ------------------------------------------------------------
// */

// append element on the end on singly linked list
func (cll *CircularLinkedList) InsertAtHead(data any) {
	defer cll.incrementCounter()

	newNode := &Node{
		Prev: cll.tail,
		Data: data,
		Next: cll.head,
	}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
	}
	cll.tail.Next = newNode
	cll.head.Prev = newNode
	cll.head = newNode
}

// push element on the beginning
func (cll *CircularLinkedList) InsertAtTail(data any) {
	defer cll.incrementCounter()

	newNode := &Node{
		Prev: cll.tail,
		Data: data,
		Next: cll.head,
	}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
	}

	cll.tail.Next = newNode
	cll.head.Prev = newNode
	cll.tail = newNode
}

// Adds a node at a specific position.
func (cll *CircularLinkedList) InsertAt(index int, data any) {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	if index > cll.length-1 {
		fmt.Println("index out of range")
		return
	}

	if index == 0 {
		newNode := &Node{
			Prev: cll.tail,
			Data: data,
			Next: cll.head,
		}
		cll.head.Prev = newNode
		cll.tail.Next = newNode
		cll.head = newNode

		cll.incrementCounter()
		return
	}

	counter := 1

	current := cll.head.Next // start from index 1

	for {
		if counter == index {
			newNode := &Node{
				Prev: current.Prev,
				Data: data,
				Next: current,
			}
			current.Prev.Next = newNode
			current.Prev = newNode
			cll.incrementCounter()
			return
		}

		current = current.Next
		counter++

		if current == cll.head {
			break
		}
	}

	for counter <= index {
		cll.tail.Next = &Node{}
		cll.tail = cll.tail.Next
		cll.incrementCounter()

		if counter == index {
			cll.tail.Data = data
			return
		}
		counter++
	}
}

// Inserts new data right after a specific existing value.
func (cll *CircularLinkedList) InsertAfter(targetData any, data any) {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	targetNode, _ := cll.Search(targetData)

	if targetNode == nil {
		return
	}

	// insert at tail
	if cll.tail == targetNode {
		cll.InsertAtTail(data)
		cll.incrementCounter()
		return
	}

	newNode := &Node{
		Prev: targetNode,
		Data: data,
		Next: targetNode.Next,
	}

	targetNode.Next.Prev = newNode
	targetNode.Next = newNode

	cll.incrementCounter()
}

// Inserts new data right after a specific existing value.
func (cll *CircularLinkedList) InsertBefore(targetData any, data any) {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	targetNode, _ := cll.Search(targetData)

	if targetNode == nil {
		return
	}

	// insert at head
	if cll.head == targetNode {
		cll.InsertAtHead(data)
		cll.incrementCounter()
		return
	}

	newNode := &Node{
		Prev: targetNode.Prev,
		Data: data,
		Next: targetNode,
	}

	targetNode.Prev.Next = newNode
	targetNode.Prev = newNode
	cll.incrementCounter()
}

// /*
// 	Deletation ------------------------------------------------------------
// */

// delete first matched element and return the deleted element
func (cll *CircularLinkedList) Delete(data any) (bool, any) {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}

	targetNode, _ := cll.Search(data)

	if cll.length == 1 {
		if targetNode.Data == cll.head.Data {
			oldData := cll.head.Data
			cll.head = nil
			cll.tail = nil

			cll.decrementCounter()
			return true, oldData
		}
	}

	if cll.head == targetNode {
		oldData := cll.head.Data
		cll.head = cll.head.Next
		cll.tail.Next = cll.head

		cll.decrementCounter()
		return true, oldData
	}

	if cll.tail == targetNode {
		oldData := cll.tail.Data
		cll.tail = cll.tail.Prev
		cll.tail.Next = cll.head
		cll.head.Prev = cll.tail

		cll.decrementCounter()
		return true, oldData
	}

	oldData := targetNode.Data
	preNode := targetNode.Prev
	postNode := targetNode.Next

	preNode.Next = postNode
	postNode.Prev = preNode
	cll.decrementCounter()
	return true, oldData
}

// delete head node.
func (cll *CircularLinkedList) DeleteHead() (bool, any) {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}
	oldData := cll.head.Data

	if cll.length == 1 {
		cll.head = nil
		cll.tail = nil
		cll.decrementCounter()
		return true, oldData
	}

	cll.head = cll.head.Next
	cll.head.Prev = cll.tail
	cll.tail.Next = cll.head
	cll.decrementCounter()
	return true, oldData
}

// delete tail node.
func (cll *CircularLinkedList) DeleteTail() (bool, any) {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, 0
	}
	oldData := cll.tail.Data

	if cll.length == 1 {
		cll.head = nil
		cll.tail = nil
		cll.decrementCounter()
		return true, oldData
	}

	cll.tail = cll.tail.Prev
	cll.tail.Next = cll.head
	cll.head.Prev = cll.tail
	cll.decrementCounter()

	return true, oldData
}

// Removes a node based on its numerical position.
func (cll *CircularLinkedList) DeleteAt(index int) (bool, any) {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return false, nil
	}

	if index > cll.length-1 {
		fmt.Println("index out of range")
		return false, nil
	}

	counter := 0

	current := cll.head

	if cll.length == 1 {
		if counter == index {
			oldData := cll.head.Data
			cll.head = nil
			cll.tail = nil
			cll.decrementCounter()
			return true, oldData
		}
	}

	if index == 0 {
		return cll.DeleteHead()
	}

	if index == cll.length-1 {
		return cll.DeleteTail()
	}

	for counter <= index {
		if counter == index {
			oldData := current.Data
			preNode := current.Prev
			postNode := current.Next

			preNode.Next = postNode
			postNode.Prev = preNode
			cll.decrementCounter()
			return true, oldData
		}
		counter++
		current = current.Next
	}

	return false, nil
}

// Keeps the first $n$ elements and deletes the rest.
func (cll *CircularLinkedList) Truncate(n int) error {
	if n == 0 || n < 0 {
		return nil
	}

	if cll.Length() < n {
		return fmt.Errorf("not enough elements in the linked list")
	}

	currentNode, err := cll.GetAt(n - 1)
	if err != nil {
		return err
	}

	cll.tail = currentNode
	cll.tail.Next = cll.head
	cll.head.Prev = cll.tail

	return nil
}

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// show the head node
func (cll *CircularLinkedList) GetHead() (*Node, error) {
	if cll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return cll.head, nil
}

// show the tail node
func (cll *CircularLinkedList) GetTail() (*Node, error) {
	if cll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return cll.tail, nil
}

// show the head node value
func (cll *CircularLinkedList) GetHeadData() (any, error) {
	if cll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return cll.head.Data, nil
}

// show the tail node value
func (cll *CircularLinkedList) GetTailData() (any, error) {
	if cll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}
	return cll.tail.Data, nil
}

// get an element of an given index and a bool status that the index exist or not
func (cll *CircularLinkedList) GetAt(index int) (*Node, error) {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := cll.head
	counter := 0

	for counter < cll.length {
		if counter == index {
			return currentNode, nil
		}
		currentNode = currentNode.Next
		counter++

		if currentNode == cll.head {
			break
		}
	}

	return nil, fmt.Errorf("linked list is empty")
}

// search an element on linked list and return boolean
func (cll *CircularLinkedList) Search(data any) (*Node, error) {
	if cll.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	currentNode := cll.head

	for {
		if currentNode.Data == data {
			return currentNode, nil
		}
		currentNode = currentNode.Next

		if currentNode == cll.head {
			break
		}
	}

	return nil, fmt.Errorf("element not found")
}

// Returns a simple true/false if the value is in the list.
func (cll *CircularLinkedList) Contains(data any) bool {
	if _, err := cll.Search(data); err == nil {
		return true
	}
	return false
}

// /*
// 	Transformation Methods ------------------------------------------------------------
// */

// Replaces a specific value with a new one.
func (cll *CircularLinkedList) Update(data, replace any) (*Node, error) {
	if cll.IsEmpty() {
		return nil, fmt.Errorf("Linked list is empty.")
	}

	targetNode, err := cll.Search(data)
	if err != nil {
		return nil, err
	}

	oldData := targetNode

	targetNode.Data = replace

	return oldData, nil
}

// reverse the linked list
func (cll *CircularLinkedList) Reverse() {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	tempCll := NewCircularLinkedList()

	current := cll.head

	for {
		tempCll.InsertAtHead(current.Data)
		current = current.Next

		if current == cll.head {
			break
		}
	}

	cll.head = tempCll.head
	cll.tail = tempCll.tail
}

// Scans the list and removes nodes with repeating values
func (cll *CircularLinkedList) RemoveDuplicates() {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	seen := make(map[any]bool)

	current := cll.head
	seen[current.Data] = true

	for {
		if seen[current.Next.Data] {
			if current.Next == cll.head {
				cll.head = cll.head.Next
			}
			current.Next = current.Next.Next
		} else {
			seen[current.Next.Data] = true
			current = current.Next
		}

		if current.Next == cll.head {
			break
		}

	}
}

// covert the linked list into slice
func (cll *CircularLinkedList) ToSlice() []any {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return nil
	}

	current := cll.head
	slc := make([]any, 0)

	for {
		slc = append(slc, current.Data)
		current = current.Next

		if current == cll.head {
			break
		}
	}

	return slc
}

// /*
// 	Metadata & Utility Methods ------------------------------------------------------------
// */

// tell how many element the linked list have
func (cll *CircularLinkedList) Length() int {
	return cll.length
}

// check the linked list is empty or not
func (cll *CircularLinkedList) IsEmpty() bool {
	return cll.head == nil
}

// Print the single linked list
func (cll *CircularLinkedList) PrintList() {
	if cll.IsEmpty() {
		fmt.Println("Linked list is empty.")
		return
	}

	current := cll.head

	for {
		fmt.Println("Data :", current.Data)
		current = current.Next

		if current == cll.head {
			break
		}
	}
}

// clear the whole linked list
func (cll *CircularLinkedList) Clear() {
	cll.head = nil
	cll.tail = nil
}

// /*
// 	private helper methods --------------------------------------------------------------------
// */

// increment after eash deletation
func (cll *CircularLinkedList) incrementCounter() {
	cll.length++ // just increate by one
}

// decrement after eash deletation
func (cll *CircularLinkedList) decrementCounter() {
	cll.length--
}
