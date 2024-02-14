package main

import (
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	resp := [][]int{}

	sort.Ints(array)

	for i := 0; i < len(array)-2; i++ {
		left, right := i+1, len(array)-1

		for left < right {
			current := array[i] + array[left] + array[right]

			if current == target {
				resp = append(resp, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1
			} else if current < target {
				left += 1
			} else if current > target {
				right -= 1
			}
		}
	}

	return resp
}

// bruto force
// non empty array
// distinct intergers
// find the triplets of sum
// ordering each triplets
// 1. initiate
// 1.1 resp [][]string
// 2. loop array
// 2.1 targetNew = target - value
// 2.2 loop from j = index
// 2.2.1 loop from z = j
// 2.2.1.1 if targetNew - array[j] == array[z]
//
//	then append to resp with ordering number
//
// 3. loop resp
// 3.1 use sort.Ints(resp[i])
// 4. return resp
func ThreeNumberSum2(array []int, target int) [][]int {
	resp := [][]int{}

	sort.Ints(array)

	for i := 0; i < len(array); i++ {
		targetNew := target - array[i]

		mapTotal := make(map[int]bool)
		for j := i + 1; j < len(array); j++ {
			mapTotal[array[j]] = true
		}

		for j := i + 1; j < len(array); j++ {
			if targetNew-array[j] == array[j] {
				continue
			}

			used, ok := mapTotal[targetNew-array[j]]
			if ok && used {
				resp = append(resp, []int{array[i], array[j], targetNew - array[j]})
				mapTotal[targetNew-array[j]] = false
				mapTotal[array[j]] = false
			}
		}
	}

	for _, v := range resp {
		sort.Ints(v)
	}

	return resp
}
