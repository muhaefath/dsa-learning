package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// *notes
// use recursive
// if leaf node, then return int
// if left still exist, sum left
// if right still exist, sum right

// 1. crete func branchSums(*BinaryTree, total, resp)[]int
// 2. total += currentTree.Value
// 3. if root.Left is not null, resp append branchSums(root.Left, total)
// 4. if root.Right is not null , resp append branchSums(root.Right, total)
// 5. if root.Right and left null, return total
// 5. return resp
func BranchSums(root *BinaryTree) []int {
	// Write your code here.

	resp := branchSums(root, 0, []int{})
	return resp
}

func branchSums(currentTree *BinaryTree, total int, resp []int) []int {
	total += currentTree.Value
	if currentTree.Left != nil {
		resp = branchSums(currentTree.Left, total, resp)
	}

	if currentTree.Right != nil {
		resp = branchSums(currentTree.Right, total, resp)
	}

	if currentTree.Left == nil && currentTree.Right == nil {
		resp = append(resp, total)
		return resp
	}

	return resp
}
