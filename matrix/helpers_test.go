package matrix

import (
	"fmt"
	"math"
	"math/rand"
)

func matricesAlmostEqual(a, b Matrix[float64], epsilon float64) bool {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range a[i] {
			if math.Abs(a[i][j]-b[i][j]) > epsilon {
				return false
			}
		}
	}
	return true
}

func approxEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func toFloat64Matrix(m Matrix[int]) Matrix[float64] {
	result := make(Matrix[float64], len(m))
	for i := range m {
		result[i] = make([]float64, len(m[i]))
		for j := range m[i] {
			result[i][j] = float64(m[i][j])
		}
	}
	return result
}

func matrixToString(m Matrix[float64]) string {
	var s string
	for _, row := range m {
		for _, val := range row {
			s += fmt.Sprintf("%8.4f ", val)
		}
		s += "\n"
	}
	return s
}

func randomIntMatrix(rows, cols int) Matrix[int] {
	m := make(Matrix[int], rows)
	for i := range m {
		m[i] = make([]int, cols)
		for j := range m[i] {
			m[i][j] = rand.Intn(10)
		}
	}
	return m
}

func randomFloatMatrix(rows, cols int) Matrix[float64] {
	m := make(Matrix[float64], rows)
	for i := range m {
		m[i] = make([]float64, cols)
		for j := range m[i] {
			m[i][j] = rand.Float64() * 10
		}
	}
	return m
}

// Helper function to check if two rows are almost equal
func rowsAlmostEqual(row1, row2 []float64, tolerance float64) bool {
	if len(row1) != len(row2) {
		return false
	}
	for i := range row1 {
		if math.Abs(row1[i]-row2[i]) > tolerance {
			return false
		}
	}
	return true
}

// Helper function to check if a matrix is the identity matrix
func isIdentity(m Matrix[int]) bool {
	if len(m) == 0 {
		return false
	}
	n := len(m)
	for i := range n {
		if len(m[i]) != n {
			return false
		}
		for j := range n {
			expected := 0
			if i == j {
				expected = 1
			}
			if m[i][j] != expected {
				return false
			}
		}
	}
	return true
}

// Helper function to check if matrix2 is a row permutation of matrix1
func isRowPermutation(matrix1, matrix2 Matrix[float64], tolerance float64) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}
	if len(matrix1) == 0 {
		return true
	}
	if len(matrix1[0]) != len(matrix2[0]) {
		return false
	}

	n := len(matrix1)
	used := make([]bool, n)

	for i := range n {
		found := false
		for j := range n {
			if used[j] {
				continue
			}
			if rowsAlmostEqual(matrix1[i], matrix2[j], tolerance) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Helper function to prevent compiler optimizations
func validateInverse(m Matrix[int], inv Matrix[float64]) {
	if len(m) > 0 && len(inv) > 0 {
		// Just access some elements to ensure the result isn't optimized away
		_ = inv[0][0]
	}
}
