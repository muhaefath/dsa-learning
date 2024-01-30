package main

import "fmt"

func main() {
	value := []int{1, 3, 5, 6}
	fmt.Println("total: ", plusOne(value))
	fmt.Println("total: ", plusOne([]int{4, 3, 2, 1}))
	fmt.Println("total: ", plusOne([]int{9}))

}

// 1. initiate value
// 2. if last digit = 9
// 2.1
// 2.1 digits[len(digits)-1] = 0
// 2.2 digits[len(digits)-2] += 1
// 3. else
// 3.1 digits[len(digits)-1] += 1
// 4. return digits
func plusOne(digits []int) []int {
	lendigits := len(digits)
	if digits[lendigits-1] == 9 {
		for i := lendigits - 1; i >= 0; i-- {
			digits[i] = 0
			if i == 0 {
				digits[0] = 1
				digits = append(digits, 0)
				return digits
			}

			if digits[i-1] != 9 {
				digits[i-1] += 1
				return digits
			}
		}
	} else {
		digits[lendigits-1] += 1
	}

	return digits
}

// 1. initiate value
// 2. if last digit = 9
// 2.1
// 2.1 digits[len(digits)-1] = 0
// 2.2 digits[len(digits)-2] += 1
// 3. else
// 3.1 digits[len(digits)-1] += 1
// 4. return digits
func plusOne2(digits []int) []int {

	lendigits := len(digits)
	if digits[lendigits-1] == 9 {
		for i := lendigits - 1; i >= 0; i-- {
			digits[i] = 0
			if i == 0 {
				tempDigit := []int{}
				tempDigit = append(tempDigit, 1)
				for i := 0; i < lendigits; i++ {
					tempDigit = append(tempDigit, digits[i])
				}

				return tempDigit
			}

			if digits[i-1] != 9 {
				digits[i-1] += 1
				break
			}
		}
	} else {
		digits[lendigits-1] += 1
	}

	return digits
}
