package arrays

func PascalsTriangle(numRows int) [][]int {
	result := make([][]int, numRows)

	for r := range numRows {
		result[r] = append(result[r], 1)
		for c := 1; c <= r; c++ {
			if c == r {
				result[r] = append(result[r], 1)
				break
			}

			result[r] = append(result[r], result[r-1][c-1]+result[r-1][c])
		}
	}

	return result
}

func PascalsTriangleMath(numRows int) [][]int {
	// (r,c) cell of pascal triangles is equal to
	// (r!)/(c!*(r-c)!)
	fact := func(n int) int {
		res := 1
		for i := 2; i <= n; i++ {
			res *= i
		}
		return res
	}

	result := make([][]int, numRows)

	for r := range numRows {
		result[r] = append(result[r], 1)
		for c := 1; c <= r; c++ {
			val := fact(r) / (fact(c) * fact(r-c))
			result[r] = append(result[r], val)
		}
	}

	return result
}

func PascalsTriangleMathOptimized(numRows int) [][]int {
	generateRows := func(row int) []int {
		res := make([]int, row)
		res[0] = 1
		ans := 1
		for col := 1; col < row; col++ {
			ans *= (row - col)
			ans /= col
			res[col] = ans
		}
		return res
	}

	results := make([][]int, 0)
	for r := 1; r <= numRows; r++ {
		results = append(results, generateRows(r))
	}
	return results
}
