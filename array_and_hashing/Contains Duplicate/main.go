package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

// 1. initiate a map
// 2. loop the nums
// 3. check if the map is exist
// 4. return true if exist
func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool)

	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			return true
		}

		numsMap[v] = true
	}
	return false
}
