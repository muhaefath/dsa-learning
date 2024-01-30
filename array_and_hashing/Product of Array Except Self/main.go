package main

import "fmt"

func main() {
	fmt.Println("final: ", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println("final: ", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

// 1. initiate resp []int
// 2. initiate templeft and tempright
// 3. loop nums to left
// 	  start from 1
//    limit len nums
// 3.1 tempLeft *= nums [i-1]
// 3.2 tempRight *= nums [i+1]
// 3.3 resp[i] *= tempLeft
// 3.4 resp[length(nums)-i+1] *= tempRight
// 4. return resp

func productExceptSelf(nums []int) []int {
	resp := make([]int, len(nums))

	tempLeft := 1
	tempRight := 1

	resp[0] = 1
	resp[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		tempLeft *= nums[i-1]
		tempRight *= nums[len(nums)-i-1+1]

		if len(nums)-i-1 > i {
			resp[i] = tempLeft
			resp[len(nums)-i-1] = tempRight
		} else if len(nums)-i-1 == i {
			resp[i] = tempLeft
			resp[len(nums)-i-1] *= tempRight
		} else {
			resp[i] *= tempLeft
			resp[len(nums)-i-1] *= tempRight
		}
	}

	return resp
}

//  1. initiate resp []int
//  2. initiate temp
//  3. loop nums to left
//     limit length of nums -1
//
// 3.1 temp += nums [i]
// 3.2 resp[i] += temp
//  4. loop nums to right
//     limit 1
//
// 4.1 temp += nums [i]
// 4.2 resp[i] += temp
// 5. return resp
func productExceptSelf2(nums []int) []int {

	resp := make([]int, len(nums))

	temp := 1

	for i := 1; i < len(nums); i++ {
		temp *= nums[i-1]
		resp[i] = temp
		fmt.Println("resp: ", i, " = ", resp[i])
	}
	temp = 1

	resp[0] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		temp *= nums[i+1]
		resp[i] *= temp

		fmt.Println("resp: ", i, " = ", resp[i])
	}

	return resp
}

// from chatgpt
func productExceptSelfChatGPT(nums []int) []int {
	length := len(nums)

	// Initialize the result slice
	resp := make([]int, length)

	// Initialize left and right products
	leftProduct := 1
	rightProduct := 1

	// Calculate left products and store them in the result slice
	for i := 0; i < length; i++ {
		resp[i] = leftProduct
		leftProduct *= nums[i]
	}

	// Calculate right products and update the result slice
	for i := length - 1; i >= 0; i-- {
		resp[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return resp
}

// 1 2 3
// 1. [0] += 2 && []
// 2. [1] += 3
// 3. [2] += x

// [2] += 2
// [1] += 1
// [0] += y

// 1 2 3 4
// 0. do nothing
// 1. [1] = 1 && [2]+1 *= 4
// 2. [2] = 2 && [1]+1 *= 12
// 3. [3] = 6 && [0]+1 *= 24
//    [i]         [len - i -1]
