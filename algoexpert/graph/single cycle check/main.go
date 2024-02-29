package main

import "fmt"

// 1. initate resp:= false
// 1.1 initiate index := 0
// 1.2 initiate mapCycle[int]bool
// 1.3 initiate counter :=1
// 2. infinite loop
// 3. map[array[index]] = true
// 4. index = array[index]
// 4.1 counter++
// 5. if index < 0, then index = len(array) - index
// 6. if index > len(array)-1, then index = index - len(array)
// 7. if counter == len(array), then break
// 8. loop mapCycle, if there is false then return false
// 9. else return true
func HasSingleCycle(array []int) bool {
	index := 0
	mapCycle := make(map[int]bool)
	counter := 1

	for counter <= len(array) {
		fmt.Println("index: ", index)

		mapCycle[index] = true
		index = index + array[index]
		counter++

		if index < 0 {
			fmt.Println("index1: ", index)
			index = len(array) + index

		} else if index > len(array)-1 {
			fmt.Println("index3: ", index)

			index = index - len(array)
		}

	}

	// for key, v := range mapCycle {
	// 	if !v {
	// 		fmt.Println("key: ", key)
	// 		return false
	// 	}
	// }
	fmt.Println("key: ", len(mapCycle))
	return len(mapCycle) == len(array)
}
