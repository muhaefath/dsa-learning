package main

func main() {

}

// non empty array
// subsequence = set number in the other array (same order)
// bruto force
// 1. initiate currentIndex
// 2. loop array
// 2.1 check if sequence[currentIndex] = value
// 2.1.1 if true, then currentIndex ++
// 3. return currentIndex == len(sequence)-1
func IsValidSubsequence(array []int, sequence []int) bool {
	if len(array) < len(sequence) {
		return false
	}

	currentIndex := 0

	for _, v := range array {
		if sequence[currentIndex] == v {
			currentIndex++

			if len(sequence) == currentIndex {
				return true
			}
		}
	}

	return false
}
