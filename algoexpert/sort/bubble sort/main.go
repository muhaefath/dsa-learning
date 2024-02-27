package main

func BubbleSort(array []int) []int {
	index := 1
	isNeedSwap := false
	if len(array) <= 1 {
		return array
	}

	for {
		if array[index] < array[index-1] {
			temp := array[index]
			array[index] = array[index-1]
			array[index-1] = temp
			isNeedSwap = true
			continue
		}

		index++
		if index == len(array) {
			if isNeedSwap {
				index = 1
				isNeedSwap = false
			} else {
				break
			}
		}
	}

	return array
}
