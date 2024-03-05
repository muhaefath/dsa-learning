package main

// array of unique integer
// return array of all permutation
// no particular order
// if input array is empty, return array empty

// 1. inititae resp[][]int
// 2. create function: getPermutationInt(integer, array)[]int
// 2.1. getPermutationInt(target int, array []int)[][]int
// 2.1.1 if len () == 2, then return [][]int{{0,1}{}1,0}
// 2.2. loop array
// 2.3. if value == target, then continue
// 2.4 tempResp := getPermutationInt(value, array)
// 2.5. insert value to first index in tempResp
// 2.6. resp = append tempResp
// 2.7 return resp
// 3. return resp
func GetPermutations(array []int) [][]int {
	if len(array) == 1 {
		return [][]int{{1}}
	}

	return getPermutationInt(array)
}

func getPermutationInt(array []int) [][]int {
	resp := [][]int{}

	if len(array) == 2 {
		resp = append(resp, []int{array[0], array[1]})
		resp = append(resp, []int{array[1], array[0]})
		return resp
	}

	for i := len(array) - 1; i >= 0; i-- {
		// assign array execpt index
		arrayAssignee := make([]int, i)

		copy(arrayAssignee, array[:i])
		arrayAssignee = append(arrayAssignee, array[i+1:]...)

		// add value to last index
		tempResps := getPermutationInt(arrayAssignee)
		for idx, tempResp := range tempResps {
			tempResp = append(tempResp, array[i])
			tempResps[idx] = tempResp
		}

		resp = append(resp, tempResps...)
	}

	return resp
}
