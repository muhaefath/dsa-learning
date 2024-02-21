package main

import (
	"math"
)

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// -1. add left and right
// -2. left - right
// -3. left / right => is decimal, need to get the close to 0
// -4. left * right

// 1. create function evaluteExpressionTree(tree, total  )int
// 2. if left and right null, then return int
// 3. if left is not null and left.value < 0
// 3.1 then totalLeft = evaluteExpressionTree(tree.left, total)
// 4. if right is not null and right.value < 0
// 4.1 then totalRight = evaluteExpressionTree(tree.right, total)
// 5. create function convert(operator, totalLeft, totalRight) int
// 5.1 if -1, then left + right
// 5.2 if -2, then left - right
// 5.3 if -3, then float64 (left /right) -> convert to closest 0
// 5.4 if -4, then left * right
// 6. return result
func EvaluateExpressionTree(tree *BinaryTree) int {
	return evaluteExpressionTree(tree)
}

func evaluteExpressionTree(tree *BinaryTree) int {
	if tree.Left == nil && tree.Right == nil {
		return tree.Value
	}

	leftTotal, rightTotal := 0, 0
	if tree.Left != nil {
		if tree.Left.Value < 0 {
			leftTotal = evaluteExpressionTree(tree.Left)
		} else {
			leftTotal = tree.Left.Value
		}
	}

	if tree.Right != nil {
		if tree.Right.Value < 0 {
			rightTotal = evaluteExpressionTree(tree.Right)
		} else {
			rightTotal = tree.Right.Value
		}
	}

	result := convert(tree.Value, leftTotal, rightTotal)
	return result

}

func convert(operator, left, right int) int {
	switch operator {
	case -1:
		return left + right
	case -2:
		return left - right
	case -3:
		total := float64(left) / float64(right)
		if total > 0 {
			return int(math.Floor(total))
		}

		return int(math.Ceil(total))
	case -4:
		return left * right
	}

	return -1
}
