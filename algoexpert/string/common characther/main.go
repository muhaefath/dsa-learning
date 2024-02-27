package main

import (
	"math"
)

// non empty list of non empty string
// return a list of char that are commong to all string
// ignoring multiplicity
// unsorted

// 1. initiate map resp[]int
// 1.1 inititiate lessLengthIndex
// 2. loop strings
// 3. check less length index
// 4. asssign all char in string to resp[value] = true
// 5. loop strings
// 6. if index == lessLengthIndex, then continue
// 7. split to rune
// 8. loop spliRune
// 9. if resp[value] += 1
// 10. loop resp
// 11. if counter == length(strings)
// 11.1 then append to commonChar
func CommonCharacters(strings []string) []string {
	mapChar := make(map[rune]bool)
	lessLengthIndex := 0
	lessLength := math.MaxInt64

	for idx, v := range strings {
		if lessLength > len(v) {
			lessLengthIndex = idx
		}
	}

	for _, v := range strings[lessLengthIndex] {
		mapChar[v] = true
	}

	for i, value := range strings {
		if i == lessLengthIndex {
			continue
		}

		mapStringValue := make(map[rune]bool)
		for _, v := range value {
			mapStringValue[v] = true
		}

		for v, _ := range mapChar {
			if _, ok := mapStringValue[v]; !ok {
				delete(mapChar, v)
			}
		}
	}

	commonChar := []string{}
	for key, _ := range mapChar {
		commonChar = append(commonChar, string(key))
	}

	return commonChar
}
