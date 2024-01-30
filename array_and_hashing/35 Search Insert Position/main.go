package main

import "fmt"

func main() {
	value := []int{1, 3, 5, 6}
	fmt.Println("total: ", value)
	fmt.Println("total: ", searchInsert(value, 5))
	fmt.Println("total: ", value)

}

// 1. initiate value
// 2. loop nums
// 2.1 if nums[i] < target
// 2.1.1 return i-1
// 2.2 if nums[i] == target
// 2.2.1 return i
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= target {
			if nums[i] == target {
				return i
			}

			return i - 1
		}
	}

	return 0
}
