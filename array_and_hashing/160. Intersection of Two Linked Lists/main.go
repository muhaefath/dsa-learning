package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	ptrA, ptrB := headA, headB

	for ptrA != ptrB {
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}

		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}

	return ptrA
}

// 1. initiate
// 1.1 length headA
// 1.2 length headB
// 1.3 diff length headA and headB
// 2. loop headA unit headA.Next is nil
// 2.1 length headA ++
// 3. loop headB unit headB.Next is nil
// 3.1 length headB ++
// 4 loop the longest
// 4.1 if the longest < diff, then dont process anything
// 4.2 if the longest == dif, then check if the headA and headB are equal
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	lengthHeadA := 0
	lengthHeadB := 0
	tempA := headA
	tempB := headB
	diff := 0

	for tempA != nil {
		lengthHeadA++
		tempA = tempA.Next
	}

	for tempB != nil {
		lengthHeadB++
		tempB = tempB.Next
	}

	diff = lengthHeadA - lengthHeadB
	isHeadA := true
	if diff < 0 {
		diff = diff * -1
		isHeadA = false
	}

	counter := 0

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}

		if isHeadA {
			if diff <= counter {
				headB = headB.Next
			}

			headA = headA.Next
		} else {
			if diff <= counter {
				headA = headA.Next
			}

			headB = headB.Next
		}

		counter++
	}

	return nil
}
