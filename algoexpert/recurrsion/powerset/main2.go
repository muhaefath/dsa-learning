package main

import "fmt"

// powerset
// 1. empty array
// 2. the value itself
// 3. all the value
// 4. combine except the value

// 1. initiate resp[][]int
// 2. resp = append empty array
// 3. resp = append all array
// 4. loop array
// 4.1 resp = append vaue
// 5. create function powerSet(newArray []int)[][]int
// 5.0 if len (newArray) == 2, then return resp[][]int{0,1}
// 5.1 loop newArray
// 5.2 initiate tempArray = make([]int, i+1)
// 5.3 copy (tempArray, newArray[:i)
// 5.3 tempArray = append newArray[:i+1]
// 5.4 call powerSet(tempArray)
// 5.5 resp append the resp
// 5.6 return resp
func Powerset(array []int) [][]int {
	return powerSet(array, len(array)-1)
}

func powerSet(array []int, index int) [][]int {
	if index < 0 {
		return [][]int{{}}
	}

	subset := powerSet(array, index-1)
	ele := array[index]
	length := len(subset)

	for i := 0; i < length; i++ {
		currentSubset := subset[i]
		newSubset := append([]int{}, currentSubset...)
		fmt.Println("newSubset1: ", newSubset)

		newSubset = append(newSubset, ele)
		fmt.Println("newSubset2: ", newSubset)

		subset = append(subset, newSubset)
	}

	fmt.Println("subset5: ", subset)

	return subset
}
