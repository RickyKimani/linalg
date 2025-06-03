package matrix

import (
	"math"
	"testing"
)

func TestEigenvaluesQR(t *testing.T) {
	tests := []struct {
		name      string
		matrix    Matrix[float64]
		maxIter   int
		tol       float64
		expected  []float64
		wantError bool
		errorMsg  string
	}{
		{
			name: "2x2 symmetric matrix",
			matrix: Matrix[float64]{
				{2, 1},
				{1, 2},
			},
			maxIter:   100,
			tol:       1e-10,
			expected:  []float64{3, 1},
			wantError: false,
		},
		{
			name:      "Empty matrix",
			matrix:    Matrix[float64]{},
			maxIter:   100,
			tol:       1e-10,
			wantError: true,
			errorMsg:  "matrix cannot be empty",
		},
		{
			name: "Non-square matrix",
			matrix: Matrix[float64]{
				{1, 2, 3},
				{4, 5, 6},
			},
			maxIter:   100,
			tol:       1e-10,
			wantError: true,
			errorMsg:  "matrix must be square",
		},
		{
			name: "Zero iterations",
			matrix: Matrix[float64]{
				{2, 1},
				{1, 2},
			},
			maxIter:   0,
			tol:       1e-10,
			expected:  []float64{2, 2}, // Original diagonal elements
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evs, err := EigenvaluesQR(tt.matrix, tt.maxIter, tt.tol)

			if tt.wantError {
				if err == nil {
					t.Errorf("Expected error but got nil")
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("Expected error message %q but got %q", tt.errorMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(evs) != len(tt.expected) {
				t.Errorf("Expected %d eigenvalues but got %d", len(tt.expected), len(evs))
				return
			}

			for i := range evs {
				match := false
				for _, exp := range tt.expected {
					if math.Abs(evs[i]-exp) < 1e-6 {
						match = true
						break
					}
				}
				if !match {
					t.Errorf("Unexpected eigenvalue: %.6f", evs[i])
				}
			}
		})
	}
}
