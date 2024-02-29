package main

import (
	"strconv"
)

// given 2 dimensional array
// 0 and 1
// 1=> black
// 0=> white
// island => 1 horizontally and vertically adjacent and dont touch the
// border of the image
// *a group horizontally and vertically adjacent 1 in first, last row
// and first, last column are not island
// island can l shaped
// return modiefied version, where all of the island are removed
// remove an island by replacing it with 0

// 1. initiate MatrixMap[string]bool
// 2. loop start i = 1 until length -1
// 3. loop start j = 1 until length -1
// 4. if MatrixMap == true, then continue
// 5. if matrix[i][j] ==  1,
// 5.1 then call ConvertRemoveIsland(I,j,matrix, MatrixMap, previousIndex ) matrix, MatrixMap, isBorder
// 6. create function: ConvertRemoveIsland
// 6.1 call CheckBorder (I, j, matrix) bool
// 6.1.1 if false -> return matrix
// 6.2 if matrix[i+1][j] == 1,
// 6.2.1 then MatrixMap = true
// 6.2.2 call ConvertRemoveIsland(i+1, j, matrix,MatrixMap, ij )
// 6.2.3 if false, return matrix
func RemoveIslands(matrix [][]int) [][]int {
	matrixMap := make(map[string]bool)

	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			keyString := strconv.Itoa(i) + strconv.Itoa(j)

			if _, ok := matrixMap[keyString]; ok {
				continue
			}

			if matrix[i][j] == 1 {
				matrix, matrixMap, _ = ConvertRemoveIsland(matrix, i, j, matrixMap, keyString, false)
			}
		}
	}

	return matrix
}

func ConvertRemoveIsland(matrix [][]int, i, j int, matrixMap map[string]bool, previousKey string, isCrossBorder bool) ([][]int, map[string]bool, bool) {
	if IsCrossBorder(matrix, i, j) {
		isCrossBorder = true
	}

	tempMatrix := matrix

	currentKey := strconv.Itoa(i) + strconv.Itoa(j)
	leftKey := strconv.Itoa(i-1) + strconv.Itoa(j)
	rightKey := strconv.Itoa(i+1) + strconv.Itoa(j)
	topKey := strconv.Itoa(i) + strconv.Itoa(j-1)
	downKey := strconv.Itoa(i) + strconv.Itoa(j+1)

	if i+1 < len(matrix)-1 && previousKey != rightKey && !matrixMap[rightKey] && matrix[i+1][j] == 1 {
		matrixMap[rightKey] = true
		matrix, matrixMap, isCrossBorder = ConvertRemoveIsland(matrix, i+1, j, matrixMap, currentKey, isCrossBorder)
	}

	if i-1 > 0 && previousKey != leftKey && !matrixMap[leftKey] && matrix[i-1][j] == 1 {
		matrixMap[leftKey] = true
		matrix, matrixMap, isCrossBorder = ConvertRemoveIsland(matrix, i-1, j, matrixMap, currentKey, isCrossBorder)
	}

	if j-1 > 0 && previousKey != topKey && !matrixMap[topKey] && matrix[i][j-1] == 1 {
		matrixMap[topKey] = true
		matrix, matrixMap, isCrossBorder = ConvertRemoveIsland(matrix, i, j-1, matrixMap, currentKey, isCrossBorder)
	}

	if j+1 < len(matrix[0])-1 && previousKey != downKey && !matrixMap[downKey] && matrix[i][j+1] == 1 {
		matrixMap[downKey] = true
		matrix, matrixMap, isCrossBorder = ConvertRemoveIsland(matrix, i, j+1, matrixMap, currentKey, isCrossBorder)
	}

	if isCrossBorder {
		return tempMatrix, matrixMap, isCrossBorder
	}

	matrix[i][j] = 0

	return matrix, matrixMap, false
}

func IsCrossBorder(matrix [][]int, i, j int) bool {
	if i-1 <= 0 && matrix[i-1][j] == 1 {
		return true
	}

	if i+1 >= len(matrix)-1 && matrix[i+1][j] == 1 {
		return true
	}

	if j-1 <= 0 && matrix[i][j-1] == 1 {
		return true
	}

	if j+1 >= len(matrix[0])-1 && matrix[i][j+1] == 1 {
		return true
	}

	return false
}
