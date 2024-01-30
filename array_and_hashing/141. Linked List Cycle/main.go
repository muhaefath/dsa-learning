package main

import "fmt"

func main() {
	fmt.Println("total: ", hasCycle)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. add nextHead
// 2. add validation if nextHead is equal currentHead it means the looping happend
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {

		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

// bruto force
// 1. initiate map headMap
// 2. loop until head.Next == nil
// 2.1 if map[key] == true then return true
// 2.2 set map[key] == true
// 2.3 set next to val
// 3. return false

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}

	headMap := make(map[ListNode]bool)
	currentHead := head

	for currentHead.Next != nil {
		if headMap[*currentHead] {
			return true
		}

		headMap[*currentHead] = true
		currentHead = currentHead.Next
	}

	return false
}
