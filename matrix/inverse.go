package matrix

import (
	"errors"
	"math"
)

// Inverse calculates the inverse of a matrix using Gauss-Jordan elimination.
//
// Parameters:
//   - m: A square matrix of type Matrix[T] where T is int or float64
//
// Returns:
//   - Matrix[float64]: The inverse of the input matrix
//   - error: Returns an error if the matrix is non-square or singular (not invertible)
//
// The function uses the Gauss-Jordan elimination method with an augmented matrix [A|I]
// to transform A into the identity matrix, simultaneously transforming I into A⁻¹.
// For numerical stability, the function performs row swapping when a pivot element is zero.
//
// Note: The inverse only exists for square matrices with non-zero determinant (non-singular).
func Inverse[T int | float64](m Matrix[T]) (Matrix[float64], error) {
	// Validate input
	if err := m.Validate(); err != nil {
		return nil, err
	}
	if !m.isSquare() {
		return nil, errors.New("cannot invert a non-square matrix")
	}
	n := len(m)

	// Create augmented matrix [A | I]
	A := make(Matrix[float64], n)
	for i := range n {
		A[i] = make([]float64, 2*n)
		for j := range n {
			A[i][j] = float64(m[i][j])
		}
		A[i][n+i] = 1 // Identity matrix in right half
	}

	// Gauss-Jordan elimination
	const epsilon = 1e-10 // Small value for numerical stability

	for i := range n {
		// Find row with maximum pivot (partial pivoting)
		maxRow := i
		maxVal := math.Abs(A[i][i])

		for k := i + 1; k < n; k++ {
			absVal := math.Abs(A[k][i])
			if absVal > maxVal {
				maxRow = k
				maxVal = absVal
			}
		}

		// Check if matrix is singular
		if maxVal < epsilon {
			return nil, errors.New("matrix is singular")
		}

		// Swap rows if needed
		if maxRow != i {
			A[i], A[maxRow] = A[maxRow], A[i]
		}

		// Normalize the pivot row
		pivot := A[i][i]
		for j := range 2 * n {
			A[i][j] /= pivot
		}

		// Eliminate column i from all other rows
		for k := range n {
			if k != i {
				factor := A[k][i]
				for j := range 2 * n {
					A[k][j] -= factor * A[i][j]
				}
			}
		}
	}

	// Extract inverse matrix
	inverse := make(Matrix[float64], n)
	for i := range n {
		inverse[i] = make([]float64, n)
		for j := range n {
			inverse[i][j] = A[i][n+j]
		}
	}

	return inverse, nil
}
