package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("final: ", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println("final: ", groupAnagrams([]string{""}))
	fmt.Println("final: ", groupAnagrams([]string{"a"}))

}

// 1. initiate a map for similar anagram
// 2. loop the strs
// 3. create a function sorting a string
// 3.1 sort the string
// 3.2 return sorted string
// 4. check in mapping with key sorted string
// 5. if exist will append the value to key
// 6. if not exist will create map with new key
// 7. loop the map and insert to return variable
func groupAnagrams(strs []string) [][]string {
	returnVariable := [][]string{}
	mapAnagram := make(map[string][]string)

	for _, v := range strs {
		value, _ := mapAnagram[sortingAnagram(v)]

		value = append(value, v)
		mapAnagram[sortingAnagram(v)] = value
	}

	for _, v := range mapAnagram {
		returnVariable = append(returnVariable, v)
	}

	return returnVariable
}

func sortingAnagram(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

// func groupAnagrams(strs []string) [][]string {
// 	returnVariable := [][]string{}
// 	mapAnagram := make(map[string][]string)

// 	for _, v := range strs {
// 		value, _ := mapAnagram[sortingAnagram(v)]
// 		value = append(value, v)
// 		mapAnagram[sortingAnagram(v)] = value

// 		fmt.Println("mapAnagram: ", value)
// 	}

// 	for _, v := range mapAnagram {
// 		returnVariable = append(returnVariable, v)
// 		fmt.Println("returnVariable: ", returnVariable)
// 	}

// 	return returnVariable
// }

// func sortingAnagram(s string) string {
// 	chars := []rune(s)
// 	sort.Slice(chars, func(i, j int) bool {
// 		return chars[i] < chars[j]
// 	})

// 	fmt.Println("sortingAnagram: ", string(chars))
// 	return string(chars)
// }
