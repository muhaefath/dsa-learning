package main

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	currentChanged := 0
	for _, coin := range coins {
		if coin > currentChanged+1 {
			return currentChanged + 1
		}

		currentChanged += coin
	}

	return currentChanged + 1
}
