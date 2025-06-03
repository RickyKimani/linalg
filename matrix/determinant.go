package matrix

import (
	"errors"
	"fmt"
)

// Det calculates the determinant of a square matrix using LU decomposition.
//
// The determinant is a scalar value that can be calculated from the elements
// of a square matrix and encodes certain properties of the linear transformation
// described by the matrix.
//
// Parameters:
//   - m: A square matrix of type Matrix[T] where T is int or float64
//
// Returns:
//   - float64: The determinant of the matrix
//   - error: An error if the matrix is empty or not square
//
// For a singular matrix (one that does not have an inverse), the function returns 0.
// The implementation uses LU decomposition with partial pivoting, which is numerically
// stable and efficient for most matrices.
func Det[T int | float64](m Matrix[T]) (float64, error) {
	// Validate input
	if err := m.Validate(); err != nil {
		return 0, fmt.Errorf("invalid matrix: %w", err)
	}

	if len(m) == 0 {
		return 0, errors.New("matrix is empty")
	}

	if !m.isSquare() {
		return 0, errors.New("matrix is not square")
	}

	// Calculate using LU decomposition
	_, U, numSwaps, err := LUDecompose(m)
	if err != nil {
		// For singular matrices, return 0 determinant
		if err.Error() == "matrix is singular" {
			return 0.0, nil
		}
		return 0, err
	}

	// Multiply diagonal elements
	det := 1.0
	for i := range len(m) {
		det *= U[i][i]
	}

	// Adjust sign based on row swaps
	if numSwaps%2 != 0 {
		det = -det
	}

	return det, nil
}
