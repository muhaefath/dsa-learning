package main

import "fmt"

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

// 3 input ancestratree
// first input -> top ancestor
// second and third descendant in ancestral
// return youngest common ancestor

// 1. initiate currentOne := descendantOne
// 1.1 initate currentTwo := descendantTwo
// 2. infinite loop
// 3. if  currentOne = currentTwo, then return currentTwo
// 4. else currentTwo = currentTwo.Ancestor
// 5. if currentTwo = nil,
// 5.1 then currentTwo = descendantTwo
// 5.2 then currentOne = currentOne.Ancestor
// 6. if currentOne == nil, then break
func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	currentOne := descendantOne
	currentTwo := descendantTwo

	for {
		fmt.Println("currentOne: ", currentOne.Name)
		fmt.Println("currentTwo: ", currentTwo.Name)

		if currentOne == currentTwo {
			return currentTwo
		}

		currentTwo = currentTwo.Ancestor
		if currentTwo == nil {
			currentTwo = descendantTwo
			currentOne = currentOne.Ancestor
		}

		if currentOne == nil {
			break
		}
	}

	return nil
}
