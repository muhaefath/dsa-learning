package main

// Do not edit the class below except
// for the breadthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

// 1. initiate children*Node
// 2. infinite loop
// 3. array, children = n.GetArray(array, children)
// 4. if children = nil, then break
// 5. create function: GetArray(array, children)
// 5.1 loop n.Children
// 5.2 array = append(array, v.Name)
// 5.3 children = append(children, v.Children)
// 6. return array
func (n *Node) BreadthFirstSearch(array []string) []string {
	children := n.Children
	array = append(array, n.Name)
	for {
		array, children = GetArray(array, children)
		if len(children) == 0 {
			break
		}
	}

	return array
}

func GetArray(resp []string, children []*Node) ([]string, []*Node) {
	tempLength := len(children)
	for _, v := range children {
		resp = append(resp, v.Name)
		children = append(children, v.Children...)
	}

	children = children[tempLength:]
	return resp, children
}
