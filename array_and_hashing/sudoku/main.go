package main

import (
	"fmt"
)

func main() {
	// fmt.Println("final: ", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println("final: ", isValidSudoku([][]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}))
	// fmt.Println("final: ", topKFrequent([]int{1, 4, 3, 2, 2, 3, 3, 2, 3, 1}, 3))

}

func isValidSudoku(board [][]string) bool {

	// 1. initiate resp
	// 2. initaite double array string
	// 2.1 set length first index is 3
	// 3. loop i
	// 3.1 loop j
	// 3.1.1 assign horizontal index to first index
	// 3.1.1.2 map[0][i] = board[i][j]
	// 3.1.2 assign vertical index to second index
	// 3.1.2.2 map[0][j] = board[i][j]
	// 3.1.3 assign box index to third index
	// 3.1.2.2 map[0][( (i%3) +3 ) + (j%3)] = board[i][j]

	return true
}

// pake mod 3 untuk cek titik 3x3 nya
//    0 1 2  3 4 5  6 7 8
//
// 0  0 1 2  3 4 5  6 7 8
// 1  1 2 3  4 5 6  7 8 9
// 2  2 3 4  5 6 7  8 9 10

// 3  3 4 5  6 7 8  9  10 11
// 4  4 5 6  7 8 9  10 11 12
// 5  5 6 7  8 9 10 11 12 13

// 6  6 7 8  9  10 11 12 13 14
// 7  7 8 9  10 11 12 13 14 15
// 8  8 9 10 11 12 13 14 15 16
