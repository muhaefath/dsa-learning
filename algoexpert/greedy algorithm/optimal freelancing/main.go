package main

import "sort"

// each job has a deadline
// each job has payment
// return maximum profit that can be obtanied in 7 day period
// each job take 1 full day to complete

// 1. initiate total
// 1.1 initate dayLeft = 7
// 2. sorting jobs by payment
// 3.0 dayLeft-jobs[0].deadline
// 3.0 total += jobs[0].payment
// 3. loop jobs start from 1
// 3.1 if jons[index].deadline == jobs[index-1].deadline
// 3.2 then delet from jobs
// 3.3 if dayleft -  jobs[0].deadline > 0
// 3.4 dayLeft-jobs[0].deadline
// 3.5 total += jobs[0].payment
// 3.6 if dayleft ==0, then return total
func OptimalFreelancing(jobs []map[string]int) int {
	total := 0

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i]["payment"] > jobs[j]["payment"]
	})

	mapDay := make([]bool, 7)
	for i := 0; i < len(jobs); i++ {
		maxTime := min(jobs[i]["deadline"], 7)
		for time := maxTime - 1; time >= 0; time-- {
			if mapDay[time] == false {
				mapDay[time] = true
				total += jobs[i]["payment"]
				break
			}
		}
	}

	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
