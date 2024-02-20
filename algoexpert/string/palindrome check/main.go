package main

// non empty string
// return palindrom or note
// palindrom is similar string for backward or forward
// 1. split into rune
// 1.1 initiate startIndex and finishIndex
// 1.2 startIndex = 0
// 1.3 finishIndex = len(splitStr) -1
// 2. inifinte loop
// 3. if startIndex >= finishIndex, then break
// 4. if splitStr[startIndex] != splitStr[endIndex], then return false
// 5. startIndex++ and finishIndex--
// 6. return true
func IsPalindrome(str string) bool {
	splitStr := make([]rune, len(str))
	for idx, v := range str {
		splitStr[idx] = v
	}

	startIndex, endIndex := 0, len(str)-1

	for {
		if startIndex >= endIndex {
			break
		}

		if splitStr[startIndex] != splitStr[endIndex] {
			return false
		}

		startIndex++
		endIndex--
	}

	return true
}
