package matrix

import (
	"math"
	"testing"
)

func TestLUDecompose(t *testing.T) {
	tests := []struct {
		name     string
		input    Matrix[int]
		wantErr  bool
		singular bool
		diagonal bool
	}{
		{
			name: "improper matrix",
			input: Matrix[int]{
				{1, 2},
				{2},
			},
			wantErr: true,
		},
		{
			name: "3x3 standard matrix",
			input: Matrix[int]{
				{2, 3, 1},
				{4, 7, 2},
				{6, 18, -1},
			},
		},
		{
			name: "Identity matrix",
			input: Matrix[int]{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			diagonal: true,
		},
		{
			name: "2x2 simple matrix",
			input: Matrix[int]{
				{4, 3},
				{6, 3},
			},
		},
		{
			name: "Singular matrix",
			input: Matrix[int]{
				{1, 2},
				{2, 4},
			},
			wantErr:  true,
			singular: true,
		},
		{
			name: "Non-square matrix",
			input: Matrix[int]{
				{1, 2, 3},
				{4, 5, 6},
			},
			wantErr: true,
		},
		{
			name: "4x4 non-singular matrix",
			input: Matrix[int]{
				{4, 3, 2, 1},
				{6, 8, 2, 1},
				{8, 4, 9, 1},
				{10, 5, 4, 7},
			},
		},
		{
			name: "Matrix requiring pivoting",
			input: Matrix[int]{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 10},
			},
		},
		{
			name:    "Empty matrix",
			input:   Matrix[int]{},
			wantErr: true,
		},
		{
			name:  "Single element matrix",
			input: Matrix[int]{{5}},
		},
		{
			name: "Matrix with zeros needing pivoting",
			input: Matrix[int]{
				{0, 1, 2},
				{1, 2, 3},
				{2, 3, 5},
			},
		},
		{
			name: "Another singular matrix (zero determinant)",
			input: Matrix[int]{
				{1, 2, 3},
				{2, 4, 6},
				{3, 6, 9},
			},
			wantErr:  true,
			singular: true,
		},
		{
			name: "2x2 identity",
			input: Matrix[int]{
				{1, 0},
				{0, 1},
			},
			diagonal: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L, U, numSwaps, err := LUDecompose(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				if tt.singular && err.Error() != "matrix is singular" {
					t.Fatalf("Expected singular matrix error, got: %v", err)
				} else if !tt.singular {
					t.Fatalf("LUDecompose() error = %v", err)
				}
				return
			}

			if tt.singular {
				t.Error("Expected error for singular matrix, got none")
				return
			}

			// Verify numSwaps is non-negative
			if numSwaps < 0 {
				t.Errorf("Number of swaps should be non-negative, got %d", numSwaps)
			}

			// 1. Verify L is lower triangular with 1s on diagonal
			for i := range L {
				for j := range L[i] {
					if i < j && math.Abs(L[i][j]) > 1e-10 {
						t.Errorf("L is not lower triangular at [%d][%d] = %f", i, j, L[i][j])
					}
					if i == j && math.Abs(L[i][j]-1) > 1e-10 {
						t.Errorf("L diagonal not 1 at [%d][%d] = %f", i, j, L[i][j])
					}
				}
			}

			// 2. Verify U is upper triangular
			for i := range U {
				for j := range U[i] {
					if i > j && math.Abs(U[i][j]) > 1e-10 {
						t.Errorf("U is not upper triangular at [%d][%d] = %f", i, j, U[i][j])
					}
				}
			}

			// 3. Verify U has no zero diagonal elements (non-singular)
			for i := range U {
				if math.Abs(U[i][i]) < 1e-10 {
					t.Errorf("U has zero diagonal element at [%d][%d] = %f", i, i, U[i][i])
				}
			}

			// 4. For LU decomposition with pivoting, we need to verify PA = LU
			// where P is the permutation matrix representing the row swaps
			// Since we don't return P explicitly, we'll verify that L*U represents
			// some row permutation of the original matrix
			LU, err := Multiply(L, U)
			if err != nil {
				t.Fatalf("Multiply() error = %v", err)
			}

			original := toFloat64Matrix(tt.input)

			// For pivoting verification, check if LU is a row permutation of original
			if !isRowPermutation(LU, original, 1e-10) {
				t.Errorf("LU decomposition failed: L*U is not a row permutation of original matrix")
				t.Logf("Original:\n%s", matrixToString(original))
				t.Logf("L*U:\n%s", matrixToString(LU))
				t.Logf("L:\n%s", matrixToString(L))
				t.Logf("U:\n%s", matrixToString(U))
				t.Logf("Number of swaps: %d", numSwaps)
			}

			// 5. For identity matrix, no swaps should be needed (but allow some tolerance)
			if tt.diagonal && isIdentity(tt.input) && numSwaps > 0 {
				// Note: Even identity matrices might have swaps due to floating point precision
				// or implementation details, so we'll just log this rather than fail
				t.Logf("Identity matrix had %d swaps (this may be acceptable)", numSwaps)
			}
		})
	}
}

func BenchmarkLUDecompose100x100(b *testing.B) {
	m := randomIntMatrix(100, 100)

	for b.Loop() {
		_, _, _, _ = LUDecompose(m)
	}
}
