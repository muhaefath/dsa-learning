package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

// 1. create function findSuccessor(tree, node.Value) *BinaryTree
// 2. if node == tree.value, then
// 2.1 if tree.right != nil, then return tree.right
// 2.2 if tree.right == nil, then callCheckParent
// 3. create function callCheckParent(currentTree, parentTree) *BinaryTree
// 3.1 if parentTree.Left.value = currentTree.Value
// 3.1.1 then return parentTree.Left.value
// 3.2 if parentTree.Left.value != currentTree.Value
// 3.2.2 then call callCheckParent(parentTree, parentTree.Parent)
// 4. if left and right is null, then return nil
// 5. if left is not nil, then findSuccessor(left, node)
// 6. if right is not nil, then findSuccessor(right, node)
// 7. return nil
func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	return findSuccessor(tree, node.Value)
}

func findSuccessor(tree *BinaryTree, node int) *BinaryTree {
	if node == tree.Value {
		if tree.Right != nil {
			return checkMostLeft(tree.Right)
		}

		return checkParent(tree, tree.Parent)
	}

	if tree.Left == nil && tree.Right == nil {
		return nil
	}

	if tree.Left != nil {
		resp := findSuccessor(tree.Left, node)
		if resp != nil {
			return resp
		}
	}

	if tree.Right != nil {
		resp := findSuccessor(tree.Right, node)
		if resp != nil {
			return resp
		}
	}

	return nil
}

func checkParent(currentTree *BinaryTree, parentTree *BinaryTree) *BinaryTree {
	if parentTree == nil {
		return parentTree
	}

	if parentTree.Left == currentTree {
		return parentTree
	}

	return checkParent(parentTree, parentTree.Parent)
}

func checkMostLeft(currentTree *BinaryTree) *BinaryTree {
	if currentTree == nil {
		return currentTree
	}

	if currentTree.Left != nil {
		return checkMostLeft(currentTree.Left)
	}

	return currentTree
}
