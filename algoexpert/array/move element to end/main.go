package main

import "fmt"

func main() {
	fmt.Println("log: ", MoveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2))
}

func MoveElementToEnd(array []int, toMove int) []int {
	i := 0
	j := len(array) - 1

	for i < j {
		for i < j && array[j] == toMove {
			j--
		}

		if array[i] == toMove {
			array[i], array[j] = array[j], array[i]
		}

		i++
	}

	return array
}

// 1. loop array
// 1.1 if toMove == v, then append array with v
// 1.2 				   and assign value array from index[1]
// 2. return array
func MoveElementToEnd3(array []int, toMove int) []int {
	idx := 0
	i := 0
	length := len(array)

	for length > 0 {
		if array[i] == toMove {
			array = append(array, array[i])
			array = append(array[:idx], array[idx+1:]...)

			fmt.Println("idx: ", idx)
			fmt.Println("array 2: ", array)
			fmt.Println("========")
		} else {
			idx++
			i++
		}

		length--
	}

	return array
}

// move the same integer with target to the last array
// no need maintain the order
// bruto force solution
// 1. resp []int
// 1.1 resp counter
// 2. loop array
// 2.1 if value == toMove, then counter ++ and continue
// 2.2 else append resp
// 3. loop until counter =0
// 3.1 append resp toMove
// 4. return resp
func MoveElementToEnd2(array []int, toMove int) []int {
	resp := []int{}
	counter := 0

	for _, v := range array {
		if v == toMove {
			counter++
		} else {
			resp = append(resp, v)
		}
	}

	for counter > 0 {
		resp = append(resp, toMove)
		counter--
	}

	return resp
}
