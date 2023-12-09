package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("isPalindrome: ", isPalindrome(121))
	fmt.Println("isPalindrome: ", isPalindrome(1221))
	fmt.Println("isPalindrome: ", isPalindrome(1212))
	fmt.Println("isPalindrome: ", isPalindrome(12321))
}

// 1. convert to string
// 2. split into array
// 3. loop all the array
// 4. set maximum loop length / 2
// 5. check if the first and last index is different, return false
func isPalindrome(x int) bool {
	xString := strconv.Itoa(x)
	xStringArray := []rune(xString)

	lastIndex := len(xStringArray) - 1
	for i := 0; i < len(xStringArray)/2; i++ {
		if xStringArray[i] != xStringArray[lastIndex] {
			return false
		}

		lastIndex--
	}

	return true
}
