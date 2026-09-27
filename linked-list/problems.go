package linkedlist

import (
	"fmt"
)

// sum two ll where number are in reverse and put the result also in reverse on a new ll
func SumTwoLL() {
	l1 := NewSinglyLinkedList()
	l2 := NewSinglyLinkedList()
	result := NewSinglyLinkedList()

	l1.InsertAtHead(9)
	l1.InsertAtHead(6)
	l1.InsertAtHead(5)

	l1.PrintList()

	fmt.Println("--------")

	l2.InsertAtHead(9)
	l2.InsertAtHead(9)
	l2.InsertAtHead(9)
	l2.InsertAtHead(7)

	l2.PrintList()

	headL1, _ := l1.GetHead()
	headL2, _ := l2.GetHead()

	carry := 0

	for headL1 != nil || headL2 != nil {
		d1 := 0
		d2 := 0
		if headL1 != nil {
			d1 = headL1.Data.(int)
		}
		if headL2 != nil {
			d2 = headL2.Data.(int)
		}

		sum := d1 + d2 + carry
		fmt.Println(d1, d2, sum)
		if sum > 9 {
			carry = sum / 10
			sum = sum % 10
		}

		result.InsertAtTail(sum)

		if headL1 != nil {
			headL1 = headL1.Next
		}

		if headL2 != nil {
			headL2 = headL2.Next
		}

		if headL1 == nil && headL2 == nil {
			result.InsertAtTail(carry)
		}
	}

	fmt.Println("-----------===========")
	result.PrintList()
}
