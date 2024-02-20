package main

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// swap left child to right child

// 1. create function invertBinaryTree
// 2. if tree.Left and tree.Right is null, return
// 3. tempTree = tree.Right
// 4. tree.Right = tree.Left
// 5. tree.Left = tempTree
// 3. if tree.Left is not null
// 3.1 then tree.Left.invertBinaryTree
// 4. if tree.Right is not null
// 4.1 then tree.Right.invertBinaryTree
// 5. return
func (tree *BinaryTree) InvertBinaryTree() {
	tree.invertBinaryTree()
}

func (tree *BinaryTree) invertBinaryTree() {
	if tree.Left == nil && tree.Right == nil {
		return
	}

	tempTree := tree.Right
	tree.Right = tree.Left
	tree.Left = tempTree

	if tree.Left != nil {
		tree.Left.invertBinaryTree()
	}

	if tree.Right != nil {
		tree.Right.invertBinaryTree()
	}

	return
}
