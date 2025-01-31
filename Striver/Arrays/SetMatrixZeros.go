// P-73 Set Matrix Zeros
package arrays

func SimpleMatrixZerosSimple(matrix [][]int) {
	zeroRow := make(map[int]struct{})
	zeroCol := make(map[int]struct{})

	for r := range matrix {
		for c := range matrix[0] {
			if matrix[r][c] == 0 {
				zeroRow[r] = struct{}{}
				zeroCol[c] = struct{}{}
			}
		}
	}

	for r := range matrix {
		for c := range matrix[0] {
			_, ok1 := zeroRow[r]
			_, ok2 := zeroCol[c]
			if ok1 || ok2 {
				matrix[r][c] = 0
			}
		}
	}
}

func SimpleMatrixZerosSpaceOptimized(matrix [][]int) {
	// traverse first row and first column to search for zero
	col0 := 1
	for c := range matrix[0] {
		if matrix[0][c] == 0 {
			col0 = 0
		}
	}

	for r := range matrix {
		if matrix[r][0] == 0 {
			matrix[0][0] = 0
		}
	}

	// traverse all cell from second row and second column
	// setting (i,j) = 0 to (i,0) = 0 and (0,j) = 0
	row := len(matrix)
	col := len(matrix[0])
	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}

	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// now for first row and column
	if matrix[0][0] == 0 {
		for r := 0; r < row; r++ {
			matrix[r][0] = 0
		}
	}

	if col0 == 0 {
		for c := 0; c < col; c++ {
			matrix[0][c] = 0
		}
	}
}
