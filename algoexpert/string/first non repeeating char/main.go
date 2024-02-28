package main

import "math"

// lowercase str
// return index of first non repeating char

// 1. init map[rune]int
// 1.1 smallestIndex := math.Max
// 2. loop str
// 3. if map not found, then map[value] = idx
// 4. if map found, then map[value] = -1
// 5. loop map
// 6. if value < smallestIndex, then smallestIndex = value
// 7. return smallestIndex
func FirstNonRepeatingCharacter(str string) int {
	mapStr := make(map[rune]int)
	smallestIndex := math.MaxInt
	isDuplicateFound := false

	for idx, v := range str {
		_, ok := mapStr[v]
		if ok {
			mapStr[v] = -1
			continue
		}

		mapStr[v] = idx
	}

	for _, v := range mapStr {
		if smallestIndex > v && v != -1 {
			isDuplicateFound = true
			smallestIndex = v
		}
	}

	if !isDuplicateFound {
		smallestIndex = -1
	}

	return smallestIndex
}
