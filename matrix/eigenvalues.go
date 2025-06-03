package matrix

import (
	"errors"
	"math"
)

// EigenvaluesQR computes the eigenvalues of a square matrix using the QR algorithm.
//
// Parameters:
//   - m: Input matrix of type Matrix[T] where T is int or float64
//   - maxIter: Maximum number of iterations for the algorithm
//   - tol: Convergence tolerance - algorithm stops when eigenvalue changes are below this threshold
//
// Returns:
//   - []float64: Array containing the eigenvalues
//   - error: Returns error if the matrix is empty, non-square, or if internal operations fail
//
// The algorithm performs iterative QR decompositions and multiplications:
// M = Q*R, then M = R*Q for each iteration
// The process terminates early if eigenvalues converge within tolerance.
//
// Recommended parameter values:
//   - maxIter = 200: Empirical testing shows this is sufficient for most matrices; higher values
//     rarely provide significant improvement
//   - tol = 1e-14: Provides near machine precision accuracy; smaller values typically don't
//     improve results due to floating-point limitations
//
// Limitations:
//   - Complex eigenvalues are ignored.
func EigenvaluesQR[T int | float64](m Matrix[T], maxIter int, tol float64) ([]float64, error) {
	// Validate input matrix
	if err := m.Validate(); err != nil {
		return nil, err
	}

	n := len(m)
	if n == 0 {
		return nil, errors.New("matrix cannot be empty")
	}
	if !m.isSquare() {
		return nil, errors.New("matrix must be square")
	}

	// Convert input to float64 matrix
	current := gtoFloat64Matrix(m)

	if maxIter <= 0 {
		eigenvalues := make([]float64, n)
		for i := range n {
			eigenvalues[i] = current[i][i]
		}
		return eigenvalues, nil
	}

	prevDiag := make([]float64, n)
	for iter := range maxIter {

		for i := range n {
			prevDiag[i] = current[i][i]
		}

		Q, R, err := QRDecompose(current)
		if err != nil {
			return nil, err
		}

		current, err = Multiply(R, Q)
		if err != nil {
			return nil, err
		}

		if iter > 0 {
			converged := true
			for i := range n {
				if math.Abs(current[i][i]-prevDiag[i]) > tol {
					converged = false
					break
				}
			}
			if converged {
				break
			}
		}
	}

	eigenvalues := make([]float64, n)
	for i := range n {
		eigenvalues[i] = current[i][i]
	}

	return eigenvalues, nil
}
