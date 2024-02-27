package main

import (
	"math"
	"sort"
)

// 1. initiate resp
// 2. loop array
// 3. check if larger than 0, 1, 2
// 3.1 then assign to the index, and change the related index
// 4. return resp
func FindThreeLargestNumbers(array []int) []int {
	resp := make([]int, 3)
	resp[0] = math.MinInt
	resp[1] = math.MinInt
	resp[2] = math.MinInt

	for _, v := range array {
		resp = swapeMax(resp, v)
	}

	resp[0], resp[2] = resp[2], resp[0]
	return resp
}

func swapeMax(resp []int, value int) []int {
	if resp[0] < value {
		resp[0], resp[1], resp[2] = value, resp[0], resp[1]
	} else if resp[1] < value {
		resp[1], resp[2] = value, resp[1]
	} else if resp[2] < value {
		resp[2] = value
	}

	return resp
}

// an array min 3 integer
// without sorting
// return sorted array of three largest integer
// need return duplicate

// sort array
// get 3 first number
// return it
func FindThreeLargestNumbers2(array []int) []int {
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})

	return []int{array[2], array[1], array[0]}
}
