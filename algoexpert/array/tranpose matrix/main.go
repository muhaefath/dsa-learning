package main

// 2d array of integers matrix
// return tranpose of the matrix
// tranpose matrix = flipped version of the original matrix accross its main diagonal
// top-left -> bottom-right
// widht and height are not necessariliy the same

// 1. initiate length matrix
// 1.2 initiate resp[length height][length width]
// 2. loop matrix
// 3. resp[j][i] = matrix[i][j]
// 4. return resp
func TransposeMatrix(matrix [][]int) [][]int {
	lengthX := len(matrix[0])
	lengthY := len(matrix)

	resp := make([][]int, lengthX)
	for i := range resp {
		resp[i] = make([]int, lengthY)
	}

	for y := 0; y < lengthY; y++ {
		for x := 0; x < lengthX; x++ {
			resp[x][y] = matrix[y][x]
		}
	}

	return resp
}
