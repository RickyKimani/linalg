package matrix

import (
	"errors"
	"fmt"

	"github.com/rickykimani/linalg/vectors"
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

// MultiplyVector performs matrix-vector multiplication (M × v).
//
// This operation multiplies a matrix by a column vector, producing a new vector.
// The result is computed as result[i] = sum(M[i,j] * v[j]) for all j.
// The number of columns in the matrix must equal the length of the vector.
//
// Parameters:
//   - m: Input matrix of type Matrix[T] where T is int or float64
//   - v: Input vector of type Vector[E] where E is int or float64
//
// Returns:
//   - Matrix[float64]: The resulting vector as a column matrix (n×1) with float64 elements
//   - error: An error if the matrix is invalid, empty, or if dimensions are incompatible
//
// The time complexity is O(m×n) where m is the number of rows and n is the number of columns.
//
// Example:
//
//	Matrix [[1,2], [3,4]] × Vector [5,6] = [[17], [39]]
func MultiplyVector[T, E int | float64](m Matrix[T], v vectors.Vector[E]) (Matrix[float64], error) {
	// Validate matrix structure
	if err := m.Validate(); err != nil {
		return nil, fmt.Errorf("matrix: %w", err)
	}

	// Handle empty inputs
	if len(m) == 0 {
		return nil, errors.New("empty matrix")
	}
	if len(v) == 0 {
		return nil, errors.New("empty vector")
	}

	// Check dimension compatibility
	if len(m[0]) != len(v) {
		return nil, fmt.Errorf("incompatible dimensions: matrix has %d columns, vector has %d elements", len(m[0]), len(v))
	}

	// Perform matrix-vector multiplication
	rows := len(m)
	result := make(Matrix[float64], rows)

	for i := range rows {
		result[i] = make([]float64, 1) // Column vector
		for j := range len(v) {
			result[i][0] += float64(m[i][j]) * float64(v[j])
		}
	}

	return result, nil
}

// VectorMultiply performs vector-matrix multiplication (v × M).
//
// This operation multiplies a row vector by a matrix, producing a new vector.
// The result is computed as result[j] = sum(v[i] * M[i,j]) for all i.
// The length of the vector must equal the number of rows in the matrix.
//
// Parameters:
//   - v: Input vector of type Vector[T] where T is int or float64
//   - m: Input matrix of type Matrix[E] where E is int or float64
//
// Returns:
//   - Matrix[float64]: The resulting vector as a row matrix (1×n) with float64 elements
//   - error: An error if the matrix is invalid, empty, or if dimensions are incompatible
//
// The time complexity is O(m×n) where m is the number of rows and n is the number of columns.
//
// Example:
//
//	Vector [1,2] × Matrix [[3,4], [5,6]] = [[13,16]]
func VectorMultiply[T, E int | float64](v vectors.Vector[T], m Matrix[E]) (Matrix[float64], error) {
	// Validate matrix structure
	if err := m.Validate(); err != nil {
		return nil, fmt.Errorf("matrix: %w", err)
	}

	// Handle empty inputs
	if len(v) == 0 {
		return nil, errors.New("empty vector")
	}
	if len(m) == 0 {
		return nil, errors.New("empty matrix")
	}

	// Check dimension compatibility
	if len(v) != len(m) {
		return nil, fmt.Errorf("incompatible dimensions: vector has %d elements, matrix has %d rows", len(v), len(m))
	}

	// Perform vector-matrix multiplication
	cols := len(m[0])
	result := make(Matrix[float64], 1) // Row vector
	result[0] = make([]float64, cols)

	for j := range cols {
		for i := range len(v) {
			result[0][j] += float64(v[i]) * float64(m[i][j])
		}
	}

	return result, nil
}
