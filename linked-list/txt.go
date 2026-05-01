package linkedlist



// // search an element on linked list and return boolean
// func (sll *singlyLinkedList) Search(data any) bool {
// 	tempNode := sll.head

// 	if tempNode == nil {
// 		fmt.Println("Linked list is empty.")
// 		return false
// 	}

// 	for tempNode != nil {
// 		if tempNode.data == data {
// 			return true
// 		}
// 		tempNode = tempNode.next
// 	}

// 	return false
// }

// // search an element on linked list and update with the replace data if exist and update the first match data.
// // return the old data and success status
// func (sll *singlyLinkedList) Update(data, replace any) (bool, any) {
// 	tempNode := sll.head

// 	if tempNode == nil {
// 		fmt.Println("Linked list is empty.")
// 		return false, 0
// 	}

// 	for tempNode != nil {
// 		if tempNode.data == data {
// 			oldData := tempNode.data
// 			tempNode.data = replace

// 			return true, oldData
// 		}
// 		tempNode = tempNode.next
// 	}

// 	return false, 0
// }

// // reverse the linked list
// func (sll *singlyLinkedList) Reverse() {
// 	tempLinkedList := NewSinglyLinkedList()

// 	currentNode := sll.head

// 	if currentNode == nil {
// 		fmt.Println("Linked list is empty.")
// 		return
// 	}

// 	for currentNode != nil {
// 		tempLinkedList.InsertAtTail(currentNode.data)
// 		currentNode = currentNode.next
// 	}

// 	sll.head = tempLinkedList.head
// }

// // delete first matched element and return the deleted element
// func (sll *singlyLinkedList) Delete(data any) (bool, any) {
// 	previous := sll.head
// 	current := sll.head.next

// 	if previous == nil {
// 		fmt.Println("Linked list is empty.")
// 		return false, 0
// 	}

// 	// if the head match first
// 	if previous.data == data {
// 		sll.head = sll.head.next
// 		return true, sll.head.data
// 	}

// 	for current != nil {
// 		if current.data == data {
// 			oldData := current.data
// 			previous.next = current.next // unlink

// 			return true, oldData
// 		}
// 		// move the pointer
// 		previous = current
// 		current = current.next
// 	}

// 	return false, 0
// }

// // delete head node.
// func (sll *singlyLinkedList) DeleteHead() (bool, any) {
// 	currentNode := sll.head

// 	if currentNode == nil {
// 		fmt.Println("Linked list is empty.")
// 		return false, 0
// 	}

// 	deletedValue := sll.head.data

// 	if sll.head.next == nil {
// 		sll.head = nil
// 		return true, deletedValue
// 	}

// 	sll.head = sll.head.next
// 	return true, deletedValue
// }



// // covert the linked list into slice
// func (sll *singlyLinkedList) ToSlice() []any {
// 	llSlice := []any{}

// 	tempHead := sll.head

// 	for tempHead != nil {
// 		llSlice = append(llSlice, tempHead.data)
// 		tempHead = tempHead.next
// 	}

// 	return llSlice
// }
