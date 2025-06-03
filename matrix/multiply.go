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

// ScalarMultiply multiplies each element of a matrix by a scalar value.
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
func ScalarMultiply[T, E int | float64](s T, m Matrix[E]) Matrix[float64] {
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

// Pow raises a square matrix to the specified integer power.
//
// Parameters:
//   - m: A square matrix of type Matrix[T] where T is int or float64
//   - n: Integer power to which the matrix is raised
//
// Returns:
//   - Matrix[float64]: The resulting matrix after raising to the power n
//   - error: An error if the matrix is not square or if the operation cannot be performed
//
// Special cases:
//   - n = 0: Returns the identity matrix of the same size
//   - n = 1: Returns a copy of the input matrix
//   - n < 0: Returns the inverse of the matrix raised to the absolute value of n
//
// Time complexity: O(n × m³) where n is the power and m is the matrix dimension.
func Pow[T int | float64](m Matrix[T], n int) (Matrix[float64], error) {
	// Validate matrix structure
	if err := m.Validate(); err != nil {
		return nil, fmt.Errorf("invalid matrix: %w", err)
	}

	if len(m) == 0 {
		return nil, errors.New("empty matrix")
	}

	if !m.isSquare() {
		return nil, errors.New("matrix must be square")
	}

	size := len(m)

	// Handle special cases
	if n == 0 {
		// Return identity matrix
		return Identity(size), nil
	}

	if n < 0 {
		// For negative power, compute inverse first
		inv, err := Inverse(m)
		if err != nil {
			return nil, fmt.Errorf("cannot compute negative power: %w", err)
		}
		return Pow(inv, -n)
	}

	if n == 1 {
		// Return a copy of the matrix
		return gtoFloat64Matrix(m), nil
	}

	// For powers > 1, use binary exponentiation for efficiency
	// Start with identity matrix as result
	result := Identity(size)

	// Convert m to float64 matrix
	base := gtoFloat64Matrix(m)

	// Binary exponentiation: x^n = (x^(n/2))² if n is even, or x·(x^(n/2))² if n is odd
	for n > 0 {
		if n%2 == 1 {
			// Multiply result by base if current bit is 1
			var err error
			result, err = Multiply(result, base)
			if err != nil {
				return nil, err
			}
		}
		n /= 2
		if n > 0 {
			// Square the base
			var err error
			base, err = Multiply(base, base)
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}
