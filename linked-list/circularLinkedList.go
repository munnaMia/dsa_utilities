package linkedlist

type circularLinkedList struct {
}

func NewCircularLinkedList() *circularLinkedList

/*
	Insertion ------------------------------------------------------------
*/

// append element on the end on singly linked list
func (sll *circularLinkedList) InsertAtHead(data any)

// push element on the beginning
func (sll *circularLinkedList) InsertAtTail(data any)

// Adds a node at a specific position.
func (sll *circularLinkedList) InsertAt(index int, data any) error

// Inserts new data right after a specific existing value.
func (sll *circularLinkedList) InsertAfter(targetData any, newData any) error

// Inserts new data right after a specific existing value.
func (sll *circularLinkedList) InsertBefore(targetData any, newData any) error

/*
	Deletation ------------------------------------------------------------
*/

// delete first matched element and return the deleted element
func (sll *circularLinkedList) Delete(data any) (bool, any)

// delete head node.
func (sll *circularLinkedList) DeleteHead() (bool, any)

// delete tail node.
func (sll *circularLinkedList) DeleteTail() (bool, any)

// Removes a node based on its numerical position.
func (sll *circularLinkedList) DeleteAt(index int)

// Keeps the first $n$ elements and deletes the rest.
func (sll *circularLinkedList) Truncate(n int)

/*
	Access & Search Methods ------------------------------------------------------------
*/

// show the head node value
func (sll *circularLinkedList) GetHead() (any, error)

// show the tail node value
func (sll *circularLinkedList) GetTail() (any, error)

// get an element of an given index
func (sll *circularLinkedList) GetAt(index int)

// search an element on linked list and return boolean
func (sll *circularLinkedList) Search(data any) bool

// Returns a simple true/false if the value is in the list.
func (sll *circularLinkedList) Contains(data any)

/*
	Transformation Methods ------------------------------------------------------------
*/

// Replaces a specific value with a new one.
func (sll *circularLinkedList) Update(data, replace any) (bool, any)

// reverse the linked list
func (sll *circularLinkedList) Reverse()

// sort the linked list
func (sll *circularLinkedList) Sort()

// Scans the list and removes nodes with repeating values
func (sll *circularLinkedList) RemoveDuplicates()

// covert the linked list into slice
func (sll *circularLinkedList) ToSlice() []any

/*
	Metadata & Utility Methods ------------------------------------------------------------
*/

// tell how many element the linked list have
func (sll *circularLinkedList) Length() int

// check the linked list is empty or not
func (sll *circularLinkedList) IsEmpty() bool

// Print the single linked list
func (sll *circularLinkedList) PrintList()

// clear the whole linked list
func (sll *circularLinkedList) Clear()
