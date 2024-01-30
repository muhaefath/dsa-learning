package main

import "fmt"

func main() {
	value := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println("total: ", value)
	fmt.Println("total: ", removeElement(value, 2))
	fmt.Println("total: ", value)

}

func removeElement(nums []int, val int) int {
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

// 1. initiate value
// 1.1 respFunc
// 2. loop nums
// 2.1 if nums[i] != val
// 2.1.1 respFunc = append(nums[i])
// 3 loop respFunc
// 3.1 nums[i] = respFunc[i]
func removeElement2(nums []int, val int) int {

	respFunc := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			respFunc = append(respFunc, nums[i])
		}
	}

	for i := 0; i < len(respFunc); i++ {
		nums[i] = respFunc[i]
	}

	return len(respFunc)
}
