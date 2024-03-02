package main

func createMatrix(rows, cols int) [][]int {
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
		for j := range result[i] {
			result[i][j] = i * j
		}
	}
	return result
}
