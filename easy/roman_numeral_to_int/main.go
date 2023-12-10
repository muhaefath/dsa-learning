package main

import (
	"fmt"
)

func main() {
	fmt.Println("isPalindrome: ", romanToInt("I"))
	fmt.Println("isPalindrome: ", romanToInt("II"))
	fmt.Println("isPalindrome: ", romanToInt("III"))
	fmt.Println("isPalindrome: ", romanToInt("IV"))
	fmt.Println("isPalindrome: ", romanToInt("V"))
	fmt.Println("isPalindrome: ", romanToInt("VI"))
	fmt.Println("isPalindrome: ", romanToInt("VII"))
	fmt.Println("isPalindrome: ", romanToInt("VIII"))
	fmt.Println("isPalindrome: ", romanToInt("IX"))
	fmt.Println("isPalindrome: ", romanToInt("X"))
	fmt.Println("isPalindrome: ", romanToInt("MCMXCIV"))

}

// 1. convert string to array of char
// 2. define library for Roman as
// 2. define position for Roman as a map
// 4. loop
// 5. convert Roman to int
// 6. check what is the next index based on value and position and assign the isAdded
// 7. if isAdded = true, then add the number and also check the next
// 8. if isAdded = false, then subs the number
func romanToInt(s string) int {
	xStringArray := []rune(s)
	mapRomanLibrary := make(map[string]int)
	// isAdded := true
	total := 0
	lastTemp := 0

	mapRomanLibrary["I"] = 1
	mapRomanLibrary["V"] = 5
	mapRomanLibrary["X"] = 10
	mapRomanLibrary["L"] = 50
	mapRomanLibrary["C"] = 100
	mapRomanLibrary["D"] = 500
	mapRomanLibrary["M"] = 1000

	for index := len(xStringArray) - 1; index >= 0; index-- {
		valueToAdd := mapRomanLibrary[string(xStringArray[index])]

		if valueToAdd >= lastTemp {
			total += valueToAdd
		} else {
			total -= valueToAdd
		}

		if total < 0 {
			total *= -1
		}

		lastTemp = valueToAdd
	}

	return total
}
