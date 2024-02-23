package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// at least one node
// return middle
// if even, return second

// 1. create function checkMiddleNode(linkedList, counter)
// 1.2 counter ++
// 1.3 checkMiddleNode(linkedList.Next, counter)
// 2. create function middleNode(linkedList, math.Round(counter / 2) , current)
// 2.1 if current == counter, then return
// 2.2 middleNode(linkedList.next,counter, current++ )
// 3.  return
func MiddleNode(linkedList *LinkedList) *LinkedList {
	counter := checkMiddleNode(linkedList, 0)
	index := int(counter / 2)
	resp := middleNode(linkedList, index, 0)
	return &resp
}

func checkMiddleNode(linkedList *LinkedList, counter int) int {
	if linkedList == nil {
		return counter
	}

	counter++
	counter = checkMiddleNode(linkedList.Next, counter)
	return counter
}

func middleNode(linkedList *LinkedList, counter int, current int) LinkedList {
	if counter == current {
		return *linkedList
	}

	return middleNode(linkedList.Next, counter, current+1)
}
