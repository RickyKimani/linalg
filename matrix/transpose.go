package matrix

// Transpose returns the transpose of a matrix.
//
// The transpose of a matrix is formed by flipping the matrix over its main diagonal,
// switching the row and column indices of each element: [A^T][i][j] = A[j][i].
//
// Parameters:
//   - m: Input matrix of type Matrix[T] where T is int or float64
//
// Returns:
//   - Matrix[T]: A new matrix that is the transpose of the input matrix
//
// For an m×n matrix, the transpose will be an n×m matrix.
// This function preserves the original matrix and returns a new matrix.
//
// Time complexity: O(m×n) where m is the number of rows and n is the number of columns.
func Transpose[T int | float64](m Matrix[T]) Matrix[T] {
	// Validate matrix structure
	if err := m.Validate(); err != nil {
		// Return empty matrix for invalid input
		return Matrix[T]{}
	}

	rows := len(m)
	if rows == 0 {
		return Matrix[T]{} // Return empty matrix for empty input
	}

	cols := len(m[0])
	result := make(Matrix[T], cols)

	for i := range cols {
		result[i] = make([]T, rows)
		for j := range rows {
			result[i][j] = m[j][i]
		}
	}

	return result
}
