package matrix

import (
	"errors"
	"fmt"
)

// Add combines two matrices by adding their corresponding elements.
//
// Parameters:
//   - a: First matrix of type Matrix[T] where T is int or float64
//   - b: Second matrix of type Matrix[E] where E is int or float64
//
// Returns:
//   - Matrix[float64]: A new matrix where each element is the sum of the
//     corresponding elements in a and b
//   - error: An error if matrices have incompatible dimensions or are empty
//
// Both matrices must have identical dimensions (rows × columns).
// The result is always a float64 matrix to accommodate mixed-type operations.
func Add[T, E int | float64](a Matrix[T], b Matrix[E]) (Matrix[float64], error) {
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
	if len(a) != len(b) {
		return nil, errors.New("incompatible dimensions")
	}

	rows := len(a)
	result := make(Matrix[float64], rows)

	// Perform addition
	for i := range rows {
		cols := len(a[i])
		if len(b[i]) != cols {
			return nil, errors.New("incompatible row lengths")
		}

		result[i] = make([]float64, cols)
		for j := range cols {
			result[i][j] = float64(a[i][j]) + float64(b[i][j])
		}
	}

	return result, nil
}

// Subtract creates a new matrix by subtracting the elements of the second matrix
// from the corresponding elements of the first matrix.
//
// Parameters:
//   - a: First matrix of type Matrix[T] where T is int or float64
//   - b: Second matrix of type Matrix[E] where E is int or float64
//
// Returns:
//   - Matrix[float64]: A new matrix where each element is a[i][j] - b[i][j]
//   - error: An error if matrices have incompatible dimensions or are empty
//
// Both matrices must have identical dimensions (rows × columns).
// The result is always a float64 matrix to accommodate mixed-type operations.
func Subtract[T, E int | float64](a Matrix[T], b Matrix[E]) (Matrix[float64], error) {
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
	if len(a) != len(b) {
		return nil, errors.New("incompatible dimensions")
	}

	rows := len(a)
	result := make(Matrix[float64], rows)

	// Perform subtraction
	for i := range rows {
		cols := len(a[i])
		if len(b[i]) != cols {
			return nil, errors.New("incompatible row lengths")
		}

		result[i] = make([]float64, cols)
		for j := range cols {
			result[i][j] = float64(a[i][j]) - float64(b[i][j])
		}
	}

	return result, nil
}
