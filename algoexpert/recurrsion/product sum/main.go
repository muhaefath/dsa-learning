package main

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.

// return product sum
// special array = integer or other spesial array

// 1. inititae total
// 2. create function: CheckReturnArray([]interface{}),counter)int
// 2.1 CheckReturnArray(array, counter) int
// 2.2 loop array
// 2.3 if array, then call CheckReturnArray(array, counter+1)
// 2.4 else, resp += value
// 2.5 return resp*counter
// 3. if value = aray, then call CheckReturnArray(array,2)
// 3.1 total += resp
// 3.2 else total += resp
// 4. return resp
func ProductSum(array []interface{}) int {
	return CheckReturnArray(array, 1)
}

func CheckReturnArray(array []interface{}, counter int) int {
	resp := 0

	for _, v := range array {
		caseResult, ok := v.(SpecialArray)

		if ok {
			resp += CheckReturnArray(caseResult, counter+1)
		} else {
			resp += v.(int)
		}
	}

	return resp * counter
}
