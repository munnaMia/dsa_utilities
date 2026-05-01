package linkedlist





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
