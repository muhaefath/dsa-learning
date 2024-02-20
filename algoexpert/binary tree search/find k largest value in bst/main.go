package main

import (
	"sort"
)

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// k = positive integer
// return kth largest ineterger contained in BST
// k is less than or equal to node
// duplicate integers, treated as distinct value

// 1. create a findKLargest(tree, k, collection) []int
// 2. call pushCollection(tree.Value, collection)
// 3 if tree.Left == nil && tree.Right == nil
// 3.1 then return collection
// 4. if tree.Left != nil,
// 4.1   then findKLargest(tree.Left, k, collection)
// 5. if tree.Right != nil,
// 5.1   then findKLargest(tree.Right, k, collection)
// 6. create function pushCollection(tree.Value, collection)
// 6.1 append tree.Value to collection
// 6.2 sort.Int(collection)
// 6.3 collection = collection[0:len(collection)-1]
// 7. return collection[k]

func FindKthLargestValueInBst(tree *BST, k int) int {
	collection := findKLargest(tree, k, []int{})
	return collection[k-1]
}

func findKLargest(tree *BST, k int, collection []int) []int {

	collection = append(collection, tree.Value)
	sort.Slice(collection, func(i, j int) bool {
		return collection[j] < collection[i]
	})

	if len(collection) > k {
		collection = collection[0 : len(collection)-1]
	}

	if tree.Left != nil {
		collection = findKLargest(tree.Left, k, collection)
	}

	if tree.Right != nil {
		collection = findKLargest(tree.Right, k, collection)
	}

	return collection
}
