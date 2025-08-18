// Package matrix provides linear algebra operations for matrices.
// It supports operations like addition, multiplication, determinants,
// eigenvalues, LU decomposition, and matrix inversion.
//
// The package uses Go's generics to support both integer and floating-point
// matrices while minimizing code duplication. Operations that inherently
// produce floating-point results (like matrix inversion) will return float64
// matrices regardless of input type.
package matrix

import (
	"fmt"
	"slices"
)

// Matrix represents a mathematical matrix as a slice of slices.
//
// The generic type parameter T is constrained to either int or float64,
// allowing for both integer and floating-point matrices.
//
// A Matrix[T] is represented as [][]T where each inner slice represents
// a row of the matrix. This implementation does not enforce that all rows
// have the same length, so care must be taken to ensure matrix consistency.
type Matrix[T int | float64] [][]T

// NewMatrix creates a new matrix from a 2D slice, ensuring all rows have the same length.
// This is the recommended way to create matrices as it validates the structure.
//
// Parameters:
//   - data: A 2D slice containing the matrix elements
//
// Returns:
//   - Matrix[float64]: A new matrix with all values converted to float64
//   - error: An error if the input data has inconsistent row lengths
//
// The function performs a deep copy of the input data and converts all elements
// to float64, so modifications to the original slice will not affect the returned matrix.
func NewMatrix[T int | float64](data [][]T) (Matrix[float64], error) {
	if len(data) == 0 {
		return Matrix[float64]{}, nil // Empty matrix is valid
	}

	rowLength := len(data[0])
	for i, row := range data {
		if len(row) != rowLength {
			return nil, fmt.Errorf("inconsistent row length at row %d: expected %d, got %d",
				i, rowLength, len(row))
		}
	}

	// Create a deep copy and convert to float64
	result := make(Matrix[float64], len(data))
	for i, row := range data {
		result[i] = make([]float64, len(row))
		for j, val := range row {
			result[i][j] = float64(val)
		}
	}

	return result, nil
}

// Validate checks if all rows in the matrix have the same length.
func (m Matrix[T]) Validate() error { //TODO: Use validate in important operations
	if len(m) == 0 {
		return nil // Empty matrix is valid
	}

	rowLength := len(m[0])
	for i, row := range m {
		if len(row) != rowLength {
			return fmt.Errorf("inconsistent row length at row %d: expected %d, got %d",
				i, rowLength, len(row))
		}
	}

	return nil
}

// isSquare returns true if the matrix has the same number of rows and columns.
//
// An empty matrix is not considered square. For a non-empty matrix,
// the function checks if the number of rows equals the number of columns
// in the first row.
func (m Matrix[T]) isSquare() bool {
	if len(m) == 0 {
		return false
	}
	return len(m) == len(m[0])
}

// cloneMatrix creates a deep copy of the input matrix.
//
// All rows and elements are copied, so modifications to the clone
// will not affect the original matrix.
func cloneMatrix[T int | float64](m Matrix[T]) Matrix[T] {
	clone := make(Matrix[T], len(m))
	for i := range m {
		clone[i] = slices.Clone(m[i])
	}
	return clone
}

// abs returns the absolute value of a number.
//
// This is a generic function that works with both int and float64 types.
func abs[T int | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// gtoFloat64Matrix converts a matrix of type T to a matrix of float64.
//
// This function is useful for operations that need to preserve precision
// or that inherently produce floating-point results.
func gtoFloat64Matrix[T int | float64](m Matrix[T]) Matrix[float64] {
	result := make(Matrix[float64], len(m))
	for i := range m {
		result[i] = make([]float64, len(m[i]))
		for j := range m[i] {
			result[i][j] = float64(m[i][j])
		}
	}
	return result
}

// Rows returns the number of rows in the matrix.
//
// This method provides a convenient way to get the row count of a matrix,
// which is equivalent to calling len() on the matrix but offers better
// semantic clarity in mathematical contexts.
//
// Returns:
//   - int: The number of rows in the matrix
//
// Example:
//
//	mat := Matrix[int]{{1, 2}, {3, 4}, {5, 6}}
//	rows := mat.Rows()  // Returns 3
func (m Matrix[T]) Rows() int {
	return len(m)
}

// Cols returns the number of columns in the matrix.
//
// For an empty matrix, this returns 0. For non-empty matrices,
// it returns the length of the first row (assuming the matrix is valid).
//
// Returns:
//   - int: The number of columns in the matrix
//
// Example:
//
//	mat := Matrix[int]{{1, 2, 3}, {4, 5, 6}}
//	cols := mat.Cols()  // Returns 3
func (m Matrix[T]) Cols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}
