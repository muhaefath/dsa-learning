package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("kasur", "rusak"))
	fmt.Println(isAnagram("ab", "a"))
	fmt.Println(isAnagram("aa", "a"))
	fmt.Println(isAnagram("aacc", "ccac"))

}

// 1. split the string into char
// 2. validate if length not same then return false
// 3. initiate the map for schar
// 4. initiate the map for tchar
// 5. loop the first string and assign the value to map
// 6. loop the second sting and check the validity
// 5. if the string not found the return false
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapStringsChar := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		mapStringsChar[rune(s[i])]++
		mapStringsChar[rune(t[i])]--
	}

	for _, v := range mapStringsChar {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagramMedium(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	totalSum := 0
	totalMin := 0
	mapStringsChar := make(map[rune]int)

	for _, v := range s {
		mapStringsChar[v] = mapStringsChar[v] + 1
		totalSum++
	}

	for _, v := range t {
		value, exist := mapStringsChar[v]
		if !exist || value == 0 {
			return false
		}

		mapStringsChar[v] = value - 1
		totalMin++
	}

	return totalSum == totalMin
}

func isAnagramBrutoForce(s string, t string) bool {
	sChar := []rune(s)
	tChar := []rune(t)

	if len(sChar) != len(tChar) {
		return false
	}

	mapStringsChar := make(map[rune]int)
	mapStringtChar := make(map[rune]int)

	for _, v := range sChar {
		mapStringsChar[v] = mapStringsChar[v] + 1
	}

	for _, v := range tChar {
		mapStringtChar[v] = mapStringtChar[v] + 1
	}

	for key, value := range mapStringtChar {
		if mapStringsChar[key] != value {
			return false
		}
	}

	return true
}
