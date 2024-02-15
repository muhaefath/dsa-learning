package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("log: ", SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17}))
}

// non empty array of integer
// absolut difference  = (array1 - array2) *-1
// 1. initiate resp[]int
// 1.1 init smallestDiff
// 2. sort array1 and array2
// 3. init pointer1 and pointer2
// 4. loop pointer1 != len(array1) && pointer2 != len(array2)
// 4.1 diff =  real(array1[pointer1] - array2[pointer2])
// 4.2 if array1[pointer1] < array2[pointer2], then pointer1++
// 4.3 else if array1[pointer1] > array2[pointer2], then pointer2++
// 4.4 else return array1[pointer1], array2[pointer2]
// 4.5 if smallestDiff > diff, then append resp and smallestDiff = diff
// 5. return resp
func SmallestDifference(array1, array2 []int) []int {
	resp := make([]int, 2)
	smallestDiff := 0

	sort.Ints(array1)
	sort.Ints(array2)

	pointer1 := 0
	pointer2 := 0

	for pointer1 < len(array1) && pointer2 < len(array2) {
		diff := real(array1[pointer1] - array2[pointer2])

		if pointer1 == 0 && pointer2 == 0 {
			smallestDiff = diff
			resp[0] = array1[pointer1]
			resp[1] = array2[pointer2]
		}

		if smallestDiff > diff {
			smallestDiff = diff
			resp[0] = array1[pointer1]
			resp[1] = array2[pointer2]
		}

		if array1[pointer1] < array2[pointer2] {
			pointer1++
		} else if array1[pointer1] > array2[pointer2] {
			pointer2++
		} else {
			return []int{array1[pointer1], array2[pointer2]}
		}

	}

	return resp
}

// non empty array of integer
// absolut difference  = (array1 - array2) *-1
// 1. initiate resp[]int
// 1.1 smallestDiff
// 2. loop array1
// 2.0 loop array2
// 2.1 diff =  real(real(array1[i]) - real(array2))
// 2.2 if index == 0, then smallestDiff = diff. resp[0] = array1[i]
// 2.3 if smallestDiff > diff, then smallestDiff = diff.  resp[0] = array1[i]
// 3. create function real
// 4. return resp
func SmallestDifference2(array1, array2 []int) []int {
	resp := make([]int, 2)
	smallestDiff := 0

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			diff := real(array1[i] - array2[j])
			if i == 0 && j == 0 {
				smallestDiff = diff
				resp[0] = array1[i]
				resp[1] = array2[j]
			}

			if smallestDiff > diff {
				smallestDiff = diff
				resp[0] = array1[i]
				resp[1] = array2[j]
			}
		}
	}

	return resp
}

func real(req int) int {
	if req < 0 {
		return -req
	}

	return req
}
