package main

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
func Powerset2(array []int) [][]int {
	resp := [][]int{}
	resp = append(resp, []int{})
	allValue := []int{}

	if len(array) == 1 {
		resp = append(resp, []int{1})
		return resp
	}

	for idx, value := range array {
		resp = append(resp, []int{value})
		allValue = append(allValue, value)

		tempArray := make([]int, idx)
		copy(tempArray, array[:idx])
		tempArray = append(tempArray, array[idx+1:]...)

		respPowerSet := powerSet2(tempArray)
		resp = append(resp, respPowerSet...)
	}

	if len(allValue) > 0 {
		resp = append(resp, allValue)
	}

	return resp
}

func powerSet2(array []int) [][]int {
	resp := [][]int{}
	allValue := []int{}

	if len(array) == 2 {
		resp = append(resp, []int{array[0], array[1]})
		return resp
	}

	for idx, value := range array {
		tempArray := make([]int, idx)
		allValue = append(allValue, value)

		copy(tempArray, array[:idx])
		tempArray = append(tempArray, array[idx+1:]...)
		respPowerSet := powerSet2(tempArray)
		resp = append(resp, respPowerSet...)

		// for i, v := range respPowerSet {
		// 	tempResp := make([]int, 0)
		// 	tempResp = append(tempResp, value)
		// 	tempResp = append(tempResp, v...)
		// 	respPowerSet[i] = tempResp
		// }

		// resp = append(resp, respPowerSet...)
	}

	if len(array) > 2 {
		resp = append(resp, allValue)
	}

	return resp
}
