package matrix

import (
	"errors"
	"fmt"
)

// Trace calculates the sum of elements on the main diagonal of a square matrix.
//
// The trace of a matrix is defined as the sum of all diagonal elements from the
// top-left to the bottom-right: tr(A) = ∑ A[i][i] for i = 0 to n-1.
//
// Parameters:
//   - m: Input matrix of type Matrix[T] where T is int or float64
//
// Returns:
//   - T: The trace value, with the same type as the input matrix elements
//   - error: Returns error if the matrix is empty or non-square
//
// The trace is only defined for square matrices. For non-square matrices, this
// function will return an error.
//
// Time complexity: O(n) where n is the dimension of the square matrix.
func Trace[T int | float64](m Matrix[T]) (T, error) {
	// Validate matrix structure
	if err := m.Validate(); err != nil {
		return 0, fmt.Errorf("invalid matrix: %w", err)
	}

	if len(m) == 0 {
		return 0, errors.New("cannot find trace of an empty matrix")
	}

	if !m.isSquare() {
		return 0, errors.New("cannot find trace of a non-square matrix")
	}

	var trace T
	n := len(m)

	for i := range n {
		trace += m[i][i]
	}

	return trace, nil
}
