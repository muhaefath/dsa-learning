package main

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		smallest := array[i]
		smallestIndex := i
		for j := i; j < len(array); j++ {
			if array[j] < smallest {
				smallest = array[j]
				smallestIndex = j
			}
		}

		array[i], array[smallestIndex] = array[smallestIndex], array[i]
	}

	return array
}
