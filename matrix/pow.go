package matrix

import (
	"errors"
	"fmt"
)

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
				return nil, err // Dead
			}
		}
		n /= 2
		if n > 0 {
			// Square the base
			var err error
			base, err = Multiply(base, base)
			if err != nil {
				return nil, err //Dead
			}
		}
	}

	return result, nil
}
