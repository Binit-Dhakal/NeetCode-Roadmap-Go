package arrays

func RotateImage(matrix [][]int) {
	// as it is square matrix
	n := len(matrix)

	// change row to column and column to row
	for r := range n {
		for c := range n {
			if c > r {
				matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
			}
		}
	}

	// reverse column Or each row of matrix
	for c := range n {
		if c >= n-c-1 {
			break
		}
		for r := range n {
			matrix[r][c], matrix[r][n-c-1] = matrix[r][n-c-1], matrix[r][c]
		}
	}
}
