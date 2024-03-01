package main

import (
	"strconv"
)

// given edges, unweighted, directed graph
// return boolean represent whether the grapn containing a cycle
// cycle => chain of at least one vertext in which the first vertex is the same as the last vertex

// 1. initiate
// 2. loop edges start I
// 3. create function: cycleInGraph(endCycle, target,edges)bool
// 4. if edges[target] == nil, then return false
// 5. loop edges[target]
// 5.1 if value == endCycle, then return true
// 5.2 else  cycleInGraph(endCycle, value, edges)
// 5.2.1 if true, then return true
// 5.3 return false
// 6. call cycleInGraph in first loop
// 7. if cycleInGraph = true, then return
// 8. return false in the end
func CycleInGraph(edges [][]int) bool {
	mapCycle := make(map[string]bool) //combination key = "endCycle:target"
	isCycle := false
	for i := 0; i < len(edges); i++ {
		for j := 0; j < len(edges[i]); j++ {
			mapCycle, isCycle = cycleInGraph(i, edges[i][j], edges, mapCycle)
			if isCycle {
				return true
			}
		}
	}

	return false
}

func cycleInGraph(endCycle, target int, edges [][]int, mapCycle map[string]bool) (map[string]bool, bool) {
	for i := 0; i < len(edges[target]); i++ {
		isCycle := false
		key := strconv.Itoa(endCycle) + strconv.Itoa(edges[target][i])

		if _, ok := mapCycle[key]; ok {
			return mapCycle, false
		}

		mapCycle[key] = true

		if edges[target][i] == endCycle {
			return mapCycle, true
		}

		mapCycle, isCycle = cycleInGraph(endCycle, edges[target][i], edges, mapCycle)
		if isCycle {
			return mapCycle, isCycle
		}

	}

	return mapCycle, false
}
