package main

import "sort"

// even number
// half red and half blue
// rearrange into two rows
// each row should contain the same number of student
// red shirt must be in the same row
// blue shirt must be in the same row
// student must be taller than in front of
// postive integer
// return whether or not following the guidelines

// 1. initiate resp
// 2. sorting red and blue
// 3. isRed := red[0] > blue[0]
// 4. loop red
// 5. if isRed true, then
// 5.1 if red[i] < blue[i], then return false
// 6. if idRed false, then
// 6.1 if red[i] > blue[i], then return false
// 7. return true

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)

	isRed := redShirtHeights[0] > blueShirtHeights[0]

	for i := 0; i < len(redShirtHeights); i++ {
		if isRed && redShirtHeights[i] <= blueShirtHeights[i] {
			return false
		} else if !isRed && redShirtHeights[i] >= blueShirtHeights[i] {
			return false
		}
	}

	return true
}
