package main

import (
	"fmt"
	"math"
)

func MinHeightBST(array []int) *BST {
	return Construct(array, 0, len(array)-1)
}

func Construct(array []int, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return nil
	}

	midIdx := (startIdx + endIdx) / 2
	bst := BST{
		Value: array[midIdx],
	}

	bst.Left = Construct(array, startIdx, endIdx-1)
	bst.Right = Construct(array, startIdx-1, endIdx)

	return &bst
}

// 2. parentIndex := ceil(array / 2)
// 3. initiate parentBST
// 4. loop array
// 4.1 if index = parentIndex, then continue
// 4.2 parentBST.Insert(value)
// 5. return parentBST
func MinHeightBST2(array []int) *BST {
	parentIndex := math.Floor(float64(len(array)) / 2)
	insertedBST := 1

	parentBST := BST{
		Value: array[int64(parentIndex)],
	}

	tempParentIndex := parentIndex
	parentIndex = math.Floor(float64(parentIndex) / 2)

	leftIndex := tempParentIndex - parentIndex
	rightIndex := tempParentIndex + parentIndex

	for insertedBST < len(array) {
		// fmt.Println("tempParentIndex: ", tempParentIndex)

		// parentIndex = math.Floor(float64(parentIndex) / 2)
		// fmt.Println("parentIndex: ", parentIndex)

		tempLeftIndex := leftIndex
		tempRightIndex := rightIndex

		leftIndex = math.Floor(float64(leftIndex) / 2)
		leftIndex = tempLeftIndex - leftIndex

		rightIndex = math.Floor(float64(rightIndex) / 2)
		rightIndex = tempRightIndex + rightIndex

		fmt.Println("leftIndex: ", leftIndex)

		fmt.Println("rightIndex: ", rightIndex)

		parentBST.Insert(array[int64(leftIndex)])
		fmt.Println("left: ", array[int64(leftIndex)])

		insertedBST++

		parentBST.Insert(array[int64(rightIndex)])
		fmt.Println("right: ", array[int64(rightIndex)])

		insertedBST++

		fmt.Println("insertedBST: ", insertedBST)
		fmt.Println("===========: ")

	}

	return &parentBST
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
