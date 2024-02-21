package main

import "fmt"

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

// has parent
// node successor = next node to be visited
// no successor if last node

// 1. create findsuccessor(tree, resp []BinaryTree)[]int
// 2. if left and right is null, then return resp and resp = append (tree)
// 3. if left is not null,
// 3.1 then resp = findsuccessor(tree.Left, resp)
// 4. if right is not null,
// 4.1 then resp = findsuccessor(tree.right, resp)
// 5. resp = append (tree.value)
// 6. return resp
// 7. loop resp
// 7.1 if value == node.Value
// 7.1.1 if index == len(resp)-1 then return nil
// 7.2. return value
func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	orderTraversal := findsuccessor(tree, []BinaryTree{})
	for idx, v := range orderTraversal {
		if v.Value == node.Value {
			fmt.Println("idx: ", idx)
			fmt.Println("length: ", len(orderTraversal)-1)

			if idx == len(orderTraversal)-1 {
				return nil
			}

			return &orderTraversal[idx+1]
		}
	}

	return nil
}

func findsuccessor(tree *BinaryTree, resp []BinaryTree) []BinaryTree {
	fmt.Println("=========")
	fmt.Println("value: ", tree.Value)

	if tree.Left == nil && tree.Right == nil {
		resp = append(resp, *tree)
		fmt.Println("resp leaf: ", resp)

		return resp
	}

	if tree.Left != nil {
		resp = findsuccessor(tree.Left, resp)
		fmt.Println("left resp: ", resp)

	}

	if tree.Right != nil {
		resp = findsuccessor(tree.Right, resp)
		fmt.Println("right resp: ", resp)

	}

	resp = append(resp, *tree)
	fmt.Println("resp3: ", resp)

	return resp
}
