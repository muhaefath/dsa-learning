package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// node depth -> distance between node and tree root
// return sum of node depth for all the node

// 1. call nodeDepth (root, total) int
// 2. if left and right is null, then return total
// 3. if left is not null, then
// 3.1 total += real (tree.Left.Value - tree.value)
// 3.2 total = nodeDepth(tree.Left, total)
// 4. if right is not null, then
// 4.1 total += real (tree.Right.Value - tree.value)
// 4.2 total = nodeDepth(tree.Right, total)
// 5. return total
func NodeDepths(root *BinaryTree) int {
	return nodeDepths(root, 0, 0)
}

func nodeDepths(tree *BinaryTree, total, position int) int {
	if tree.Left == nil && tree.Right == nil {
		return total
	}

	if tree.Left != nil {
		leftPosition := position + 1

		total += leftPosition
		total = nodeDepths(tree.Left, total, leftPosition)
	}

	if tree.Right != nil {
		rightPosition := position + 1

		total += rightPosition
		total = nodeDepths(tree.Right, total, rightPosition)
	}

	return total
}
