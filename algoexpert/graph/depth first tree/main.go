package main

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

// implement depth first search
// traverses the tree with depth first search approach
// navigating tree from left to right
// store all names in the array and return it

// 1. array = append name
// 2. loop children
// 2.1 array = value[index].DepthFirstSearch(array)
// 3. return array
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)

	for _, value := range n.Children {
		array = value.DepthFirstSearch(array)
	}

	return array
}
