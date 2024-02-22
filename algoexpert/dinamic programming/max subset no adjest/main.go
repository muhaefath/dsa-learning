package main

// postive integer
// return maximum sum of non adjacent element

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}

	maxSum := make([]int, len(array))

	maxSum[0] = array[0]
	maxSum[1] = Max(array[0], array[1])

	for i := 2; i < len(array); i++ {
		maxSum[i] = Max(maxSum[i-1], maxSum[i-2]+array[i])
	}

	return maxSum[len(array)-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
