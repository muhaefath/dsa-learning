package main

import "fmt"

func main() {
	value := []int{1, 1, 2}
	fmt.Println("total: ", value)
	fmt.Println("total: ", removeDuplicates(value))
	fmt.Println("total: ", value)

}

func removeDuplicatesChatGPT(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastIndex := 1

	for i := 1; i < len(nums); i++ {
		if nums[lastIndex-1] != nums[i] {
			nums[lastIndex] = nums[i]
			lastIndex++
		}
	}

	return lastIndex
}

//  1. initiate variable
//     lastIndex = int
//     lastValue = nums[0]
//  2. loop nums start index 1
//
// 2.1 if lastValue != nums[i]
// 2.1.1 nums[lastIndex] = nums[i]
// 2.1.2 lastIndex++
// 2.1.3 lastValue = nums[i]
func removeDuplicates(nums []int) int {
	lastValue := nums[0]
	lastIndex := 1
	for i := 1; i < len(nums); i++ {
		if lastValue != nums[i] {
			nums[lastIndex] = nums[i]
			lastIndex++
			lastValue = nums[i]
		}
	}

	return lastIndex
}

//  1. initiate variable
//     lastIndex = int
//     mapNums [int]counter
//  2. loop nums
//
// 2.1 check if exist in map
// 2.2 if not exist
// 2.2.1 will asign nums[i] to map
// 2.2.2 update nums[lastIndex] = nums[i]
// 2.2.3 lastIndex++
// 2.3.2 if exist ->
// 3 return len(mapNums)
func removeDuplicates2(nums []int) int {
	mapNums := make(map[int]bool)
	lastIndex := 0
	for i := 0; i < len(nums); i++ {
		_, found := mapNums[nums[i]]
		if !found {
			mapNums[nums[i]] = true
			nums[lastIndex] = nums[i]
			lastIndex++
		}
	}

	return len(mapNums)
}

//  1. initiate variable
//     k = int
//     mapNums [int]counter
//     respFunc []int
//  2. loop nums
//
// 2.1 check if exist in map
// 2.1.1 if not exist will asign to map
// 2.1.2 if exist -> append to respFunc
// 3. print respFunc
// 4 return len(mapNums)
func removeDuplicates3(nums []int) int {
	mapNums := make(map[int]int)
	respFunc := []int{}

	for i := 0; i < len(nums); i++ {
		counter, found := mapNums[nums[i]]
		if !found {
			mapNums[nums[i]] = 1
			respFunc = append(respFunc, nums[i])
		} else {
			mapNums[nums[i]] = counter + 1
		}
	}

	for i := 0; i < len(respFunc); i++ {
		nums[i] = respFunc[i]
	}

	return len(mapNums)
}
