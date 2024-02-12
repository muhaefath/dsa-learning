package main

func main() {

}

// bruote force solution
// 1. initiate resp []int
// 2. loop array
// 2.1 start I
// 3. nested loop array
// 3.1 start j=i
//  4. if array[i] + array[j] == target
//     then append to resp
//  5. return resp
func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}

	return []int{}
}

// optimize solution
// how to make the time complexity to o(n)
// non empty array
// distinct integer
// return in any order
// 1. initiate result int
// 1.1 initiate map[int]bool
// 2. loop start index = 0
// 3. assign value to map
// 4. loop start index = 0
// 5. get value from map[target - array[I]]
// 5.1 if the map value is false
// 6. append to result
// 7. set the map to true
func TwoNumberSum2(array []int, target int) []int {
	mapArray := make(map[int]bool)

	for _, v := range array {
		mapArray[v] = false
	}

	for _, v := range array {
		if target-v == v {
			continue
		}

		_, ok := mapArray[target-v]
		if ok {
			return []int{v, target - v}

		}

	}

	return []int{}
}
