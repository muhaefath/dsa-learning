package main

import "fmt"

// return minimum number of edit operation that need to be performed
// on the first string to obtain second string
// there are 3 operation:
// insertion char
// delete char
// subtitute char

// 1. initiate minimumOperation
// 1.1 lengthChecking := 0
// 1.2 if a > b, then length = 1
// 1.3 if a < b, then length = -1
// 2. loop a
// 3. if a[i] != b[i]
// 4. if lengthChecking == 0, then subtitute
// 4.1 then minimumOperation++
// 5. if lengthChecking == 1, then add
// 5.1 then minimumOperation++
// 6. if lengthChecking == -1 then remove
// 6.1 then minimumOperation++
// 7. if a[i] == b[i], then index++
// 8. if index == len(b), then break
// 9. return minimumOperation
func LevenshteinDistance(a, b string) int {
	minimumOperation := 0
	index := 0

	aSlice := make([]rune, len(a))
	bSlice := make([]rune, len(b))

	for idx, v := range a {
		aSlice[idx] = v
	}

	for idx, v := range b {
		bSlice[idx] = v
	}

	for {
		fmt.Println("=========")

		fmt.Println("index: ", index)
		fmt.Println("aSlice[index]: ", string(aSlice[index]))
		fmt.Println("bSlice[index]: ", string(bSlice[index]))
		fmt.Println("minimumOperation: ", minimumOperation)

		if aSlice[index] != bSlice[index] {
			if len(aSlice) == len(bSlice) {
				aSlice[index] = bSlice[index]
			} else if len(aSlice) < len(bSlice) {
				aSlice = append(aSlice, bSlice[index])

				temp := aSlice[index]
				for i := index + 1; i < len(aSlice)-1; i++ {
					temp = aSlice[i]
					aSlice[i+1] = temp
				}

				aSlice[index] = bSlice[index]
			} else {
				for i := index; i < len(aSlice)-1; i++ {
					aSlice[i] = aSlice[i+1]
				}

				aSlice = aSlice[:len(aSlice)-1]
			}

			minimumOperation++
			continue
		}

		index++
		if index == len(bSlice) {
			break
		}
	}

	return minimumOperation
}
