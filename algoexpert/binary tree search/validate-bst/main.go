package main

import (
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// 1. create a function validateBst(*BST)
// 2. if tree.Left.Value < tree.Value and tree.Left != nil, then validateBst(tree.Left)
// 2.1 else return false
// 3. if tree.Right.Value >= tree.Value and tree.Right != nil, then validateBst(tree.Right)
// 3.1 else return false
// 4. return true
func (tree *BST) ValidateBst() bool {
	return tree.validateBst(math.MinInt32, math.MaxInt32)
}

func (tree *BST) validateBst(minValue, maxValue int) bool {
	if tree.Value < minValue || tree.Value >= maxValue {
		return false
	}

	if tree.Left != nil && !tree.Left.validateBst(minValue, tree.Value) {
		return false
	}

	if tree.Right != nil && !tree.Right.validateBst(tree.Value, maxValue) {
		return false
	}

	return true

}
