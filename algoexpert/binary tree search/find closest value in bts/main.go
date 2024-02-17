package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// 1. initiate closestValue
// 1.1 initiate smallestDiff
// 2. loop until Left & Right == nil
// 2.1 diffCurrentValue := real (value - target)
// 2.2 diffLeft := real (Left.value - target)
// 2.3 diffRight := real (Right. value - target)
// 2.4 if diffLeft > diffRight, then current = Left
// 2.5 if diffRight < diffLeft, then current = Right

// 2.1 diffCurrentValue := real (value - target)
// 2.1.1 if diffCurrentValue < smallestDiff, then smallestDiff = diffCurrentValue and closestValue = currentTree.Value
// 2.2  leftDiff := real (value -1 - target)
// 2.3 rightDiff := real (value +1 - target)
// 2.4 if leftDiff < smallestDiff, then current = Left
// 2.5 if rightDiff < smallestDiff, then current = right
func (tree *BST) FindClosestValue(target int) int {
	closestValue := tree.Value
	smallestDiff := real(tree.Value - target)
	currentTree := tree

	for {
		diffCurent := real(currentTree.Value - target)
		if diffCurent == 0 {
			return currentTree.Value
		}

		if diffCurent < smallestDiff {
			smallestDiff = diffCurent
			closestValue = currentTree.Value
		}

		if currentTree.Left == nil && currentTree.Right == nil {
			break
		}

		diffLeft := diffCurent
		if currentTree.Left != nil {
			diffLeft = real(currentTree.Value - 1 - target)
		}

		diffRight := diffCurent
		if currentTree.Right != nil {
			diffRight = real(currentTree.Value + 1 - target)
		}

		if diffLeft < diffCurent || currentTree.Right == nil {
			currentTree = currentTree.Left
		} else if diffRight < diffCurent || currentTree.Left == nil {
			currentTree = currentTree.Right
		}
	}

	return closestValue
}

func real(req int) int {
	if req < 0 {
		return -req
	}

	return req
}
