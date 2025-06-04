package matrix

// Scale multiplies each element of a matrix by a scalar value.
//
// Parameters:
//   - s: Scalar value of type T where T is int or float64
//   - m: Matrix of type Matrix[E] where E is int or float64
//
// Returns:
//   - Matrix[float64]: A new matrix with each element multiplied by the scalar value
//
// This function preserves the original matrix and returns a new matrix.
// The result is always a float64 matrix to accommodate mixed-type operations.
//
// Time complexity: O(m×n) where m is the number of rows and n is the number of columns.
func Scale[T, E int | float64](s T, m Matrix[E]) Matrix[float64] {
	// Validate matrix structure
	if err := m.Validate(); err != nil {
		// Return empty matrix for invalid input
		return Matrix[float64]{}
	}

	rows := len(m)
	if rows == 0 {
		return Matrix[float64]{}
	}

	cols := len(m[0])
	result := make(Matrix[float64], rows)
	scalar := float64(s)

	for i := range rows {
		result[i] = make([]float64, cols)
		for j := range cols {
			result[i][j] = scalar * float64(m[i][j])
		}
	}

	return result
}
