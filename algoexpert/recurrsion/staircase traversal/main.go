package main

func StaircaseTraversal(height int, maxSteps int) int {
	return numberOfWays(height, maxSteps)
}

func numberOfWays(height, maxStep int) int {
	if height <= 1 {
		return 1
	}

	ways := 0
	for step := 1; step < min(maxStep, height)+1; step++ {
		ways += numberOfWays(height-step, maxStep)
	}

	return ways
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
