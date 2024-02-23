package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// sorted order
// return with no duplicate value

// 1. if len.Next != nil
// 1.1 if value == next.value
// 1.1.1 if next.next != nil
// 1.1.2 then next = next.next
// 1.2.3 if next.next == nil
// 1.2.4 then next = nil
// 1.3 RemoveDuplicatesFromLinkedList(linkedList.next)
// 2. return linkedList
func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	if linkedList == nil {
		return nil
	}

	if linkedList.Next != nil {
		if linkedList.Next.Value == linkedList.Value {
			linkedList.Next = linkedList.Next.Next
			RemoveDuplicatesFromLinkedList(linkedList)
		} else {
			RemoveDuplicatesFromLinkedList(linkedList.Next)
		}
	}

	return linkedList
}
