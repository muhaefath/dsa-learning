package main

// n-1 + n-2
func GetNthFib(n int) int {
	resp := 0
	prevFirst := 0
	prevSecond := 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			resp = 0
			prevSecond = 0
		} else if i == 2 {
			resp = 1
			prevFirst = 1
		} else if i == 3 {
			resp = 1
			prevFirst = 1
			prevSecond = 1
		} else {
			resp = prevFirst + prevSecond
			prevSecond = prevFirst
			prevFirst = resp
		}
	}

	return resp
}
