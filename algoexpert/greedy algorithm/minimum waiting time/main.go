package main

import "sort"

// non empty array
// positive inetgers
// query can executed in any order
// only one query can be executed
// query waiting time = amount of time must be wait before execution start
// return minimum amount of total waiting time for all queries

// 1. initiate totalWaiting
// 1.1 initiate resp
// 2. sorting queries
// 3. loop queries start index 0 until length
// 4. resp += totalWaiting
// 5. totalWaiting += value
func MinimumWaitingTime(queries []int) int {
	totalWaiting, resp := 0, 0

	sort.Ints(queries)

	for i := 0; i < len(queries); i++ {
		resp += totalWaiting
		totalWaiting += queries[i]
	}

	return resp
}

// 0 + 1 + 3 + 5 + 8
// 0 + 1 + 2 + 3
