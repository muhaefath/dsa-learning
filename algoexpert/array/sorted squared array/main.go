package main

// non empty array of int
// sorted acs
// return new array same length with squared of the interger
// 1. initiate resp []int
// 1.1 initiate smallestIndex, biggestIndex
// 2. create function translate to turn minus interger to integer
// 3. loop until resp == len(array)
// 4. get smallestValue = array[smallestIndex]
// 4.1 get higestValue = array[biggestIndex]
//  5. check if translate(value smallestValue * smallestValue) >
//     translate(value higestValue * higestValue)
//     then append to resp smallestValue * smallestValue
//     smallestIndex++
//
// 5.1 else append to resp higestValue * higestValue
//
//	biggestIndex++
//
// 6 return resp
func SortedSquaredArray(array []int) []int {
	resp := make([]int, len(array))
	smallestIndex := 0
	biggestIndex := len(array) - 1

	for i := len(array) - 1; i >= 0; i-- {
		smallestValue := array[smallestIndex]
		higestValue := array[biggestIndex]

		if translate(smallestValue) > translate(higestValue) {
			resp[i] = smallestValue * smallestValue
			smallestIndex++
		} else {
			resp[i] = higestValue * higestValue
			biggestIndex--
		}
	}

	return resp
}

func translate(req int) int {
	if req < 0 {
		return -req
	}

	return req
}
