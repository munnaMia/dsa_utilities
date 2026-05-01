package linkedlist

type doublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// /*
// 	Insertion ------------------------------------------------------------
// */

// // append element on the end on singly linked list
// func (sll *doublyLinkedList) InsertAtHead(data any)

// // push element on the beginning
// func (sll *doublyLinkedList) InsertAtTail(data any)

// // Adds a node at a specific position.
// func (sll *doublyLinkedList) InsertAt(index int, data any) error

// // Inserts new data right after a specific existing value.
// func (sll *doublyLinkedList) InsertAfter(targetData any, newData any) error

// // Inserts new data right after a specific existing value.
// func (sll *doublyLinkedList) InsertBefore(targetData any, newData any) error

// /*
// 	Deletation ------------------------------------------------------------
// */

// // delete first matched element and return the deleted element
// func (sll *doublyLinkedList) Delete(data any) (bool, any)

// // delete head node.
// func (sll *doublyLinkedList) DeleteHead() (bool, any)

// // delete tail node.
// func (sll *doublyLinkedList) DeleteTail() (bool, any)

// // Removes a node based on its numerical position.
// func (sll *doublyLinkedList) DeleteAt(index int)

// // Keeps the first $n$ elements and deletes the rest.
// func (sll *doublyLinkedList) Truncate(n int)

// /*
// 	Access & Search Methods ------------------------------------------------------------
// */

// // show the head node value
// func (sll *doublyLinkedList) GetHead() (any, error)

// // show the tail node value
// func (sll *doublyLinkedList) GetTail() (any, error)

// // get an element of an given index and a bool status that the index exist or not
// func (sll *doublyLinkedList) GetAt(index int) (bool , any)

// // search an element on linked list and return boolean
// func (sll *doublyLinkedList) Search(data any) bool

// // Returns a simple true/false if the value is in the list.
// func (sll *doublyLinkedList) Contains(data any) bool

// /*
// 	Transformation Methods ------------------------------------------------------------
// */

// // Replaces a specific value with a new one.
// func (sll *doublyLinkedList) Update(data, replace any) (bool, any)

// // reverse the linked list
// func (sll *doublyLinkedList) Reverse()

// // sort the linked list
// func (sll *doublyLinkedList) Sort()

// // Scans the list and removes nodes with repeating values
// func (sll *doublyLinkedList) RemoveDuplicates()

// // covert the linked list into slice
// func (sll *doublyLinkedList) ToSlice() []any

// /*
// 	Metadata & Utility Methods ------------------------------------------------------------
// */

// // tell how many element the linked list have
// func (sll *doublyLinkedList) Length() int

// // check the linked list is empty or not
// func (sll *doublyLinkedList) IsEmpty() bool

// // Print the single linked list
// func (sll *doublyLinkedList) PrintList()

// // clear the whole linked list
// func (sll *doublyLinkedList) Clear()

// /*
// 	private helper methods --------------------------------------------------------------------
// */
