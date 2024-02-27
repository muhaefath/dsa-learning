package main

// 1. initiate mapResp := [rune]int
// 2. split charcharter
// 2.1 assign to mapResp ++
// 3. split document
// 3.1 get document char from mapResp
// 3.2 if not found or found but counter is 0, then return false
func GenerateDocument(characters string, document string) bool {
	mapResp := make(map[rune]int)

	for _, v := range characters {
		value, ok := mapResp[v]
		if ok {
			mapResp[v] = value + 1
		} else {
			mapResp[v] = 1
		}
	}

	for _, v := range document {
		value, ok := mapResp[v]
		if !ok {
			return false
		}

		if value == 0 {
			return false
		}

		mapResp[v] = value - 1
	}

	return true
}
