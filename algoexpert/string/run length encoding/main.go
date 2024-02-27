package main

import (
	"strconv"
)

// non empty string
// return run-length encoding
// run-length encoding is convert similar char to integer
// ex: aaa => 3a (max 9 char)

// 1. initiate resp
// 1.1 currentRune = rune[0]
// 1.2 counterRune := 1
// 2. split str to rune
// 3. loop rune start from index =1
// 4. if rune[i] == current rune and currentRune < 10
// 4.1 then counterRune ++
// 5. if not equal
// 5.1 resp += counterRune + currentRune
// 5.2 counterRune = 1
// 5.3 currentRune = rune[I]
// 6. return resp
func RunLengthEncoding(str string) string {
	resp := ""
	counterRune := 1
	splitStr := []rune{}

	for _, v := range str {
		splitStr = append(splitStr, v)
	}

	currentRune := splitStr[0]
	for i := 1; i < len(splitStr); i++ {
		if string(splitStr[i]) == string(currentRune) && counterRune < 9 {
			counterRune++
			continue
		}

		resp += "" + strconv.Itoa(counterRune) + string(currentRune)
		counterRune = 1
		currentRune = splitStr[i]
	}

	resp += "" + strconv.Itoa(counterRune) + string(currentRune)
	return resp
}
