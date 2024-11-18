package binarysearch

func IterativeSearchMatrix(matrix [][]int, target int) bool {
	row, column := len(matrix), len(matrix[0])

	leftPtr, rightPtr := 0, row*column-1

	for leftPtr <= rightPtr {
		midPtr := leftPtr + (rightPtr-leftPtr)/2
		r, c := midPtr/column, midPtr%column

		num := matrix[r][c]
		if num < target {
			leftPtr = midPtr + 1
		} else if num > target {
			rightPtr = midPtr - 1
		} else {
			return true
		}
	}
	return false
}

func RecursiveSearchMatrix(matrix [][]int, target int) bool {
	row, column := len(matrix), len(matrix[0])
	leftPtr, rightPtr := 0, row*column-1
	var searchMatrix func(int, int) bool

	searchMatrix = func(leftPtr, rightPtr int) bool {
		if leftPtr > rightPtr {
			return false
		}

		midPtr := leftPtr + (rightPtr-leftPtr)/2
		r, c := midPtr/column, midPtr%column

		num := matrix[r][c]
		if num == target {
			return true
		}

		if num > target {
			return searchMatrix(leftPtr, midPtr-1)
		}

		return searchMatrix(midPtr+1, rightPtr)
	}

	return searchMatrix(leftPtr, rightPtr)
}
