package arrays

func Search2DMatrixBinary(matrix [][]int, target int) bool {
	nRow := len(matrix)
	nCol := len(matrix[0])

	leftPtr, rightPtr := 0, nRow*nCol-1
	for leftPtr <= rightPtr {
		mid := (leftPtr + rightPtr) / 2
		r := mid / nCol
		c := mid % nCol

		if matrix[r][c] < target {
			leftPtr = mid + 1
		} else if matrix[r][c] > target {
			rightPtr = mid - 1
		} else {
			return true
		}
	}
	return false
}
