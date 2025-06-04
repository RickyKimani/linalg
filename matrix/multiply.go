package matrix

import (
	"errors"
	"fmt"
)

// Multiply performs matrix multiplication between two matrices.
//
// Matrix multiplication is calculated as C[i,j] = sum(A[i,k] * B[k,j]) for all k.
// The inner dimensions must match: the number of columns in matrix A must equal
// the number of rows in matrix B.
//
// Parameters:
//   - A: First matrix of type Matrix[T] where T is int or float64
//   - B: Second matrix of type Matrix[E] where E is int or float64
//
// Returns:
//   - Matrix[float64]: The resulting matrix, always with float64 elements to accommodate
//     mixed type operations
//   - error: An error if either matrix is empty or if the dimensions are incompatible
//
// The time complexity is O(n³) for square matrices of size n×n, or more generally
// O(rows × cols × common) where rows and cols are the dimensions of the result matrix
// and common is the shared dimension between the input matrices.
func Multiply[T, E int | float64](a Matrix[T], b Matrix[E]) (Matrix[float64], error) {
	// Validate matrix structure
	if err := a.Validate(); err != nil {
		return nil, fmt.Errorf("first matrix: %w", err)
	}
	if err := b.Validate(); err != nil {
		return nil, fmt.Errorf("second matrix: %w", err)
	}

	// Handle empty matrices
	if len(a) == 0 || len(b) == 0 {
		return nil, errors.New("empty matrix")
	}

	// Check dimension compatibility
	if len(a[0]) != len(b) {
		return nil, errors.New("incompatible dimensions")
	}

	// Perform multiplication
	rows := len(a)
	cols := len(b[0])
	inner := len(b)

	result := make(Matrix[float64], rows)
	for i := range rows {
		result[i] = make([]float64, cols)
		for j := range cols {
			for k := range inner {
				result[i][j] += float64(a[i][k]) * float64(b[k][j])
			}
		}
	}

	return result, nil
}
