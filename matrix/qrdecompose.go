package matrix

import (
	"errors"
	"math"
)

// QRDecompose performs QR decomposition of a matrix using the Gram-Schmidt process.
//
// QR decomposition factorizes a matrix A into a product Q*R where Q is an orthogonal
// matrix (QᵀQ = I) and R is an upper triangular matrix. This decomposition is useful
// for solving linear systems, least squares problems, and computing eigenvalues.
//
// Parameters:
//   - A: Input matrix of type Matrix[T] where T is int or float64
//
// Returns:
//   - Matrix[float64]: Orthogonal matrix Q where QᵀQ = I
//   - Matrix[float64]: Upper triangular matrix R
//   - error: Returns error if the matrix is empty, non-rectangular, or has linearly dependent columns
//
// The function uses the classical Gram-Schmidt orthogonalization algorithm.
// If the columns of A are linearly dependent (resulting in a zero norm during
// orthogonalization), the function will return an error.
//
// Time complexity: O(n²m) where n is the number of rows and m is the number of columns.
func QRDecompose[T int | float64](m Matrix[T]) (Matrix[float64], Matrix[float64], error) {
	// Validate matrix structure
	if err := m.Validate(); err != nil {
		return nil, nil, err
	}

	n := len(m)
	if n == 0 {
		return nil, nil, errors.New("empty matrix")
	}

	if !m.isSquare() {
		return nil, nil, errors.New("matrix must be square")
	}

	cols := len(m[0])

	// Initialize Q and R matrices
	q := make(Matrix[float64], n)
	r := make(Matrix[float64], cols)
	for i := range n {
		q[i] = make([]float64, cols)
	}
	for i := range cols {
		r[i] = make([]float64, cols)
	}

	// Gram-Schmidt process
	for j := range cols {
		// Copy column j of m into vector v
		v := make([]float64, n)
		for i := range n {
			v[i] = float64(m[i][j])
		}

		// Orthogonalize against previous columns
		for k := range j {
			// r[k][j] = dot(q_k, m_j)
			var dot float64
			for i := range n {
				dot += q[i][k] * float64(m[i][j])
			}
			r[k][j] = dot

			// v = v - dot * q_k
			for i := range n {
				v[i] -= dot * q[i][k]
			}
		}

		// r[j][j] = ||v||
		var norm float64
		for i := range n {
			norm += v[i] * v[i]
		}
		norm = math.Sqrt(norm)

		// Check for linear dependence
		if norm < 1e-10 {
			return nil, nil, errors.New("linearly dependent columns (zero norm)")
		}
		r[j][j] = norm

		// q[:,j] = v / norm
		for i := range n {
			q[i][j] = v[i] / norm
		}
	}

	return q, r, nil
}
