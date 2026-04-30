package linkedlist

type singlyLinkedList struct {
}

func NewSinglyLinkedList() *singlyLinkedList

/*
	Insertion ------------------------------------------------------------
*/

// append element on the end on singly linked list
func (sll *singlyLinkedList) InsertAtHead(data any)

// push element on the beginning
func (sll *singlyLinkedList) InsertAtTail(data any)

// Adds a node at a specific position.
func (sll *singlyLinkedList) InsertAt(index int, data any) error

// Inserts new data right after a specific existing value.
func (sll *singlyLinkedList) InsertAfter(targetData any, newData any) error

// Inserts new data right after a specific existing value.
func (sll *singlyLinkedList) InsertBefore(targetData any, newData any) error

/*
	Deletation ------------------------------------------------------------
*/

// delete first matched element and return the deleted element
func (sll *singlyLinkedList) Delete(data any) (bool, any)

// delete head node.
func (sll *singlyLinkedList) DeleteHead() (bool, any)

// delete tail node.
func (sll *singlyLinkedList) DeleteTail() (bool, any)

// Removes a node based on its numerical position.
func (sll *singlyLinkedList) DeleteAt(index int)

// Keeps the first $n$ elements and deletes the rest.
func (sll *singlyLinkedList) Truncate(n int)

/*
	Access & Search Methods ------------------------------------------------------------
*/

// show the head node value
func (sll *singlyLinkedList) GetHead() (any, error)

// show the tail node value
func (sll *singlyLinkedList) GetTail() (any, error)

// get an element of an given index
func (sll *singlyLinkedList) GetAt(index int)

// search an element on linked list and return boolean
func (sll *singlyLinkedList) Search(data any) bool

// Returns a simple true/false if the value is in the list.
func (sll *singlyLinkedList) Contains(data any)

/*
	Transformation Methods ------------------------------------------------------------
*/

// Replaces a specific value with a new one.
func (sll *singlyLinkedList) Update(data, replace any) (bool, any)

// reverse the linked list
func (sll *singlyLinkedList) Reverse()

// sort the linked list
func (sll *singlyLinkedList) Sort()

// Scans the list and removes nodes with repeating values
func (sll *singlyLinkedList) RemoveDuplicates()

// covert the linked list into slice
func (sll *singlyLinkedList) ToSlice() []any

/*
	Metadata & Utility Methods ------------------------------------------------------------
*/

// tell how many element the linked list have
func (sll *singlyLinkedList) Length() int

// check the linked list is empty or not
func (sll *singlyLinkedList) IsEmpty() bool

// Print the single linked list
func (sll *singlyLinkedList) PrintList()

// clear the whole linked list
func (sll *singlyLinkedList) Clear()
