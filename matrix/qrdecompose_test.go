package matrix

import (
	"fmt"
	"strings"
	"testing"
)

func TestQRDecompose(t *testing.T) {
	A := Matrix[float64]{
		{12, -51},
		{6, 167},
	}

	Q, R, err := QRDecompose(A)
	if err != nil {
		t.Fatal(err)
	}

	// Check A ≈ Q * R
	AR, err := Multiply(Q, R)
	if err != nil {
		t.Fatal(err)
	}

	for i := range A {
		for j := range A[i] {
			if !approxEqual(A[i][j], AR[i][j]) {
				t.Errorf("A != Q*R at [%d][%d]: expected %.6f, got %.6f", i, j, A[i][j], AR[i][j])
			}
		}
	}

	// Check QᵀQ ≈ I
	QT := Transpose(Q)
	QTQ, err := Multiply(QT, Q)
	if err != nil {
		t.Fatal(err)
	}
	for i := range QTQ {
		for j := range QTQ[i] {
			expected := 0.0
			if i == j {
				expected = 1.0
			}
			if !approxEqual(QTQ[i][j], expected) {
				t.Errorf("QᵀQ not identity at [%d][%d]: expected %.6f, got %.6f", i, j, expected, QTQ[i][j])
			}
		}
	}

	// Check R is upper triangular
	for i := range R {
		for j := range R[i] {
			if i > j && !approxEqual(R[i][j], 0) {
				t.Errorf("R not upper triangular at [%d][%d]: expected 0, got %.6f", i, j, R[i][j])
			}
		}
	}
}

func TestQRDecompose_ErrorCases(t *testing.T) {
	// Test case 1: Empty matrix
	t.Run("empty matrix", func(t *testing.T) {
		A := Matrix[float64]{}
		_, _, err := QRDecompose(A)
		if err == nil {
			t.Error("expected error for empty matrix, got nil")
		}
		if err.Error() != "empty matrix" {
			t.Errorf("expected 'empty matrix' error, got: %v", err)
		}
	})

	// Test case 2: Non-rectangular matrix
	t.Run("improper matrix", func(t *testing.T) {
		A := Matrix[float64]{
			{1, 2, 3},
			{4, 5}, // One element less than the first row
		}
		_, _, err := QRDecompose(A)
		if err == nil {
			t.Error("expected error for improper matrix, got nil")
			return
		}

		expectedErrorPrefix := "inconsistent row length at row"
		if !strings.Contains(err.Error(), expectedErrorPrefix) {
			t.Errorf("expected error containing '%s', got: %v", expectedErrorPrefix, err)
		}
	})

	t.Run("non-square matrix", func(t *testing.T) {
		A := Matrix[float64]{
			{1, 2, 3},
			{1, 3, 0},
		}

		_, _, err := QRDecompose(A)
		if err == nil {
			t.Error("expected error for non-square matrix, got nil")
		}
	})

	// Test case 3: Linearly dependent columns
	t.Run("linearly dependent columns", func(t *testing.T) {
		A := Matrix[float64]{
			{1, 2},
			{2, 4}, // Second column is multiple of first column
		}
		_, _, err := QRDecompose(A)
		if err == nil {
			t.Error("expected error for linearly dependent columns, got nil")
		}
		if err.Error() != "linearly dependent columns (zero norm)" {
			t.Errorf("expected 'linearly dependent columns (zero norm)' error, got: %v", err)
		}
	})
}

func BenchmarkQRDecompose(b *testing.B) {
	sizes := []int{10, 50, 100}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%dx%d", size, size), func(b *testing.B) {
			A := randomFloatMatrix(size, size)
			b.ResetTimer()
			for b.Loop() {
				_, _, _ = QRDecompose(A)
			}
		})
	}
}
