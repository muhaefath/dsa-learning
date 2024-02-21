package main

import (
	"math"
	"strings"
)

// non empty string
// non negative integer
// convert the charater into next char alphabetical

// 1. convert str to array rune
// 2. create string a-z
// 3. get index of the string
func CaesarCipherEncryptor(str string, key int) string {
	alhpahbet := "abcdefghijklmnopqrstuvwxyz"
	alhpahbetCollection := []rune{}
	for _, v := range alhpahbet {
		alhpahbetCollection = append(alhpahbetCollection, v)
	}

	result := ""

	for _, v := range str {
		i := strings.Index(alhpahbet, string(v))
		newIndex := i + key

		if newIndex >= len(alhpahbet) {
			counter := math.Floor((float64(key) / 26))

			if counter == 0 {
				counter = 1
			}
			newIndex = newIndex - (len(alhpahbet) * int(counter))
		}

		result += string(alhpahbetCollection[newIndex])
	}

	return result
}
