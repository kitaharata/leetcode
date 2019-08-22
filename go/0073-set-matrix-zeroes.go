func setZeroes(matrix [][]int) {
	n, m, l := len(matrix[0]), len(matrix), 1
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			l = 0
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if l == 0 {
			matrix[i][0] = 0
		}
	}
}
