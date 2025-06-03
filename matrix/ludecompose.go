package matrix

import (
	"errors"
	"math"
)

// LUDecompose performs LU decomposition with partial pivoting on a square matrix.
//
// LU decomposition factorizes a matrix A into the product of a lower triangular
// matrix L and an upper triangular matrix U, optionally with row permutations: PA = LU,
// where P is a permutation matrix represented by the returned swap count.
//
// Parameters:
//   - m: Input matrix of type Matrix[T] where T is int or float64
//
// Returns:
//   - Matrix[float64]: Lower triangular matrix L with 1's on the diagonal
//   - Matrix[float64]: Upper triangular matrix U
//   - int: Number of row swaps performed during pivoting (useful for determinant sign)
//   - error: Returns error if the matrix is empty, non-square, or singular
//
// This implementation uses partial pivoting for numerical stability, which
// selects the largest absolute value in each column as the pivot element.
// The decomposition is useful for solving linear systems, calculating determinants,
// and inverting matrices more efficiently than direct methods.
//
// Time complexity: O(n³) where n is the matrix dimension.
func LUDecompose[T int | float64](m Matrix[T]) (Matrix[float64], Matrix[float64], int, error) {
	if err := m.Validate(); err != nil {
		return nil, nil, 0, err
	}

	n := len(m)
	if n == 0 {
		return nil, nil, 0, errors.New("matrix is empty")
	}

	if !m.isSquare() {
		return nil, nil, 0, errors.New("matrix is not square")
	}

	// Convert input to float64 matrix
	a := make(Matrix[float64], n)
	for i := range n {
		a[i] = make([]float64, n)
		for j := range n {
			a[i][j] = float64(m[i][j])
		}
	}

	// Initialize L and U matrices
	l := make(Matrix[float64], n)
	u := make(Matrix[float64], n)
	for i := range n {
		l[i] = make([]float64, n)
		u[i] = make([]float64, n)
	}

	numSwaps := 0 // Track row swaps for determinant sign calculation

	// Perform LU decomposition with partial pivoting
	for i := range n {
		// Find pivot row with largest absolute value in column i
		maxRow := i
		maxVal := math.Abs(a[i][i])
		for k := i + 1; k < n; k++ {
			absVal := math.Abs(a[k][i])
			if absVal > maxVal {
				maxVal = absVal
				maxRow = k
			}
		}

		// Swap rows if necessary
		if maxRow != i {
			a[i], a[maxRow] = a[maxRow], a[i]

			// Swap already computed parts of L
			for k := range i {
				l[i][k], l[maxRow][k] = l[maxRow][k], l[i][k]
			}
			numSwaps++
		}

		// Compute U row i
		for k := i; k < n; k++ {
			sum := 0.0
			for j := range i {
				sum += l[i][j] * u[j][k]
			}
			u[i][k] = a[i][k] - sum
		}

		// Check for singularity
		if math.Abs(u[i][i]) < 1e-12 {
			return nil, nil, 0, errors.New("matrix is singular")
		}

		// Compute L column i
		for k := i; k < n; k++ {
			if i == k {
				l[i][i] = 1.0 // Diagonal of L is always 1
			} else {
				sum := 0.0
				for j := range i {
					sum += l[k][j] * u[j][i]
				}
				l[k][i] = (a[k][i] - sum) / u[i][i]
			}
		}
	}

	return l, u, numSwaps, nil
}
