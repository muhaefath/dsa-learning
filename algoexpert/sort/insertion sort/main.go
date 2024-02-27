package main

func InsertionSort(array []int) []int {

	for i := range array {
		for j := i - 1; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}

	return nil
}

func InsertionSort2(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	index := 1
	lastComparedIndex := -1
	for {
		if index == 0 {
			index = lastComparedIndex
			lastComparedIndex = -1
		}

		if array[index] < array[index-1] {
			array[index], array[index-1] = array[index-1], array[index]
			lastComparedIndex = index
			index--
			continue
		}

		if lastComparedIndex > -1 {
			index = lastComparedIndex
			lastComparedIndex = -1
			continue
		}

		index++
		if index == len(array) {
			break
		}
	}

	return array
}
