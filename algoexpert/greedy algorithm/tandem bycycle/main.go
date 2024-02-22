package main

import "sort"

// tandem bicycle by to people a and b
// speed depends on fastest speed of the person
// postive integer
// return maxium possible total speed or minimum possible total speed

// 1. initiate resp
// 2. sort red
// 3. sort blue
// 4. loop red
// 5. resp += max(red[I] , blue[length(red)-1 - I])
// 6. return resp
func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	resp := 0
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)

	for i := 0; i < len(redShirtSpeeds); i++ {
		if fastest {
			resp += MaxSpeed(redShirtSpeeds[i], blueShirtSpeeds[len(redShirtSpeeds)-1-i])
		} else {
			resp += MaxSpeed(redShirtSpeeds[i], blueShirtSpeeds[i])
		}
	}

	return resp
}

func MaxSpeed(a, b int) int {
	if a > b {
		return a
	}

	return b
}
