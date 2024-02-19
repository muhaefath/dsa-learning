package main

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// 1. initiate current tree
// 2. infinite loop
// 3. if value < currentTree.value
// 3.1 if currentTree.Left != nil, then currentTree = currentTree.Left
// 3.2 else currentTree.Left = new tree
// 3.3 return tree
// 4. if value > currentTree.value
// 4.1 if currentTree.Right != nil, then currentTree = currentTree.Right
// 4.2 else currentTree.Right = new tree
// 4.3 return tree
func (tree *BST) Insert(value int) *BST {
	newTree := &BST{Value: value}

	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.Insert(value)
		} else {
			tree.Left = newTree
		}
	} else {
		if tree.Right != nil {
			tree.Right.Insert(value)
		} else {
			tree.Right = newTree
		}
	}

	return tree
}

// 1. initiate current tree
// 2. infinite loop
// 3. if value less than currentTree.Left, then currentTree = Left
// 4. if value more than currentTree.Right then currentTree = Right
func (tree *BST) Contains(value int) bool {
	if value < tree.Value {
		if tree.Left == nil {
			return false
		} else {
			return tree.Left.Contains(value)
		}
	} else if value > tree.Value {
		if tree.Right == nil {
			return false
		} else {
			return tree.Right.Contains(value)
		}
	}

	return true
}

// remove first instance of given value
// if single node tree, will do nothing
// if value is leaf tree, will set the parent left / right null
// if value has left and right, will change to closet value
// if value has left / right only, will change the value to child

// 1. currentTree = tree
// 1.1 if tree doesnt have any children, return tree
// 2. search with Contains
// 3. if value is leaf tree, will set the parent left / right null
// 4. if  value has left / right only, will change the value to child
// 5. if value has left and right, will change to closet value
func (tree *BST) Remove(value int) *BST {
	tree.remove(value, nil)
	return tree
}

func (tree *BST) remove(value int, parent *BST) {
	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.remove(value, tree)
		}
	} else if value > tree.Value {
		if tree.Right != nil {
			tree.Right.remove(value, tree)
		}
	} else {
		if tree.Left != nil && tree.Right != nil {
			tree.Value = tree.Right.getMinValue()
			tree.Right.remove(tree.Value, tree)
		} else if parent == nil {
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Right = tree.Left.Right
				tree.Left = tree.Left.Left
			} else if tree.Right != nil {
				tree.Value = tree.Right.Value
				tree.Right = tree.Right.Right
				tree.Left = tree.Right.Left
			}
		} else if parent.Left == tree {
			if tree.Left != nil {
				parent.Left = tree.Left
			} else {
				parent.Left = tree.Right
			}
		} else if parent.Right == tree {
			if tree.Left != nil {
				parent.Right = tree.Left
			} else {
				parent.Right = tree.Right
			}
		}
	}

}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}

	return tree.Left.getMinValue()
}
