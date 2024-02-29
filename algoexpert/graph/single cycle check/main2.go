package main

func HasSingleCycle(array []int) bool {
	seen := map[int]bool{}

	curr := 0

	for jumps := 0; jumps < len(array); jumps += 1 {
		curr = (curr + array[curr]) % len(array)
		if curr < 0 {
			curr += len(array)
		}

		seen[curr] = true
	}

	return len(seen) == len(array)
}
