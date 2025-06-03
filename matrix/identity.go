package matrix

// Identity creates an n×n identity matrix.
//
// An identity matrix has 1s along the main diagonal (top-left to bottom-right)
// and 0s everywhere else. It serves as the multiplicative identity in matrix
// multiplication: A·I = I·A = A for any compatible matrix A.
//
// Parameters:
//   - n: The dimension of the square identity matrix to create
//
// Returns:
//   - Matrix[float64]: An n×n identity matrix
//
// The function returns a Matrix[float64] rather than a generic type for consistency
// with other matrix operations that involve floating-point computation.
func Identity(n int) Matrix[float64] {
	m := make(Matrix[float64], n)
	for i := range n {
		m[i] = make([]float64, n)
		m[i][i] = 1
	}
	return m
}
