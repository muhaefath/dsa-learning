package main

import (
	"sort"
)

// a list of unique string
// return semordnilap pairs
// semordnilap pair => same combination of every char
// ordering not matter

// 1. init resp[][]string
// 1.1 create string alphabet
// 1.2 init mapChar[rune]int
// 1.4 init currentIndex
// 1.5 init index
// 2. loop alphabet and assign mapChar
// 4. infinite loop
// 5. if words[currentIndex] = index
// 5.1 then append to resp
// 6. index ++
// 7. if index == len(words)
// 7.1 then currentIndex++
// 7.2 then index = currentIndex+1
// 8. if currentIndex == len(words)
// 8.1 then break
// 9. return resp
func Semordnilap(words []string) [][]string {
	resp := [][]string{}
	mapWords := make(map[string]string)
	currentIndex := 0
	index := 1

	if len(words) == 0 {
		return resp
	}

	for _, value := range words {
		temp := []rune{}
		for _, v := range value {
			temp = append(temp, v)
		}

		sort.Slice(temp, func(i, j int) bool {
			return temp[i] > temp[j]
		})

		for _, v := range temp {
			mapWords[value] += string(v)
		}
	}

	for {
		if mapWords[words[currentIndex]] == mapWords[words[index]] {
			resp = append(resp, []string{words[currentIndex], words[index]})
		}

		index++
		if index == len(words) {
			currentIndex++
			index = currentIndex + 1
		}

		if currentIndex == len(words)-1 {
			break
		}
	}

	return resp
}
