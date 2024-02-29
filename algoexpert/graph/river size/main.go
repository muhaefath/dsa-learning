package main

import (
	"strconv"
)

// given 2 dimension array
// potentially unqeual height and width
// 0 => land
// 1 => river (horizontal or vertical but not diagonally)
// river can twist, => doesnt have to straight vertical or horizontal line.
// it can be L-shaped
// return an array of size of all rivers
// no need ordering

// 1. initiate map[string]bool
// 1.1 resp []int
// 2. loop I
// 3. loop j
// 4. if map[i+j] = true, then continue
// 5. if matrix[i][j] == 0, then continue
// 6. else j, counter, map = convert(I,j, resp, 1, map)
// 7. resp = append counter
// 8. create function: convert(matrix, I,j, counter, map)j, counter, map
// 8.1 if matrix[i+1][j] == 1
// 8.1.1 then counter++
// 8.1.2 _, counter, map := convert(i+1, j, counter, map)
// 8.2 if matrix[i][j+1] == 1
// 8.2.1 then counter
// 8.2.2 ma[] = true
// 8.2.3 j, counter, map:= convert(i, j+1, counter, map)
// 9. return resp
func RiverSizes(matrix [][]int) []int {
	mapRiver := make(map[string]bool)
	resp := []int{}
	mapRiver[strconv.Itoa(0)+strconv.Itoa(0)] = true

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			keyString := strconv.Itoa(i) + strconv.Itoa(j)
			if _, ok := mapRiver[keyString]; ok && keyString != "00" {
				continue
			}

			if matrix[i][j] == 1 {
				counter := 1
				counter, mapRiver = ConvertRiver(matrix, i, j, counter, mapRiver, keyString)
				resp = append(resp, counter)
			}
		}
	}

	return resp
}

func ConvertRiver(matrix [][]int, i, j int, counter int, mapRiver map[string]bool, previousKey string) (int, map[string]bool) {
	keyStringDown := strconv.Itoa(i+1) + strconv.Itoa(j)
	keyStringUp := strconv.Itoa(i-1) + strconv.Itoa(j)
	keyStringRight := strconv.Itoa(i) + strconv.Itoa(j+1)
	keyStringLeft := strconv.Itoa(i) + strconv.Itoa(j-1)
	currentKey := strconv.Itoa(i) + strconv.Itoa(j)

	if !mapRiver[keyStringDown] && previousKey != keyStringDown && i+1 < len(matrix) && matrix[i+1][j] == 1 {
		counter++

		keyString := keyStringDown

		mapRiver[keyString] = true
		counter, mapRiver = ConvertRiver(matrix, i+1, j, counter, mapRiver, currentKey)
	}

	if !mapRiver[keyStringUp] && previousKey != keyStringUp && i-1 >= 0 && matrix[i-1][j] == 1 {
		counter++

		keyString := keyStringUp

		mapRiver[keyString] = true
		counter, mapRiver = ConvertRiver(matrix, i-1, j, counter, mapRiver, currentKey)
	}

	if !mapRiver[keyStringRight] && previousKey != keyStringRight && j+1 < len(matrix[i]) && matrix[i][j+1] == 1 {

		counter++

		keyString := keyStringRight

		mapRiver[keyString] = true
		counter, mapRiver = ConvertRiver(matrix, i, j+1, counter, mapRiver, currentKey)
	}

	if !mapRiver[keyStringLeft] && previousKey != keyStringLeft && j-1 >= 0 && matrix[i][j-1] == 1 {
		counter++

		keyString := keyStringLeft

		mapRiver[keyString] = true
		counter, mapRiver = ConvertRiver(matrix, i, j-1, counter, mapRiver, currentKey)
	}

	return counter, mapRiver
}
