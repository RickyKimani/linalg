package matrix

import (
	"fmt"
	"testing"
)

func TestInverse(t *testing.T) {
	t.Run("Regular cases", func(t *testing.T) {
		tests := []struct {
			name     string
			matrix   Matrix[int]
			expected Matrix[float64]
		}{
			{
				name: "Identity matrix",
				matrix: Matrix[int]{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
				expected: Matrix[float64]{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			{
				name: "3x3 matrix",
				matrix: Matrix[int]{
					{3, 0, 2},
					{2, 0, -2},
					{0, 1, 1},
				},
				expected: Matrix[float64]{
					{0.2, 0.2, 0},
					{-0.2, 0.3, 1},
					{0.2, -0.3, 0},
				},
			},
			{
				name: "4x4 matrix",
				matrix: Matrix[int]{
					{1, 2, 3, 0},
					{0, 1, 4, 4},
					{5, 6, 0, 2},
					{1, 1, 1, 1},
				},
				// Add expected inverse or verify by A*A⁻¹=I check
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				inv, err := Inverse(tt.matrix)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				// Verify inverse is correct
				if len(tt.expected) > 0 && !matricesAlmostEqual(inv, tt.expected, 1e-6) {
					t.Errorf("expected %v, got %v", tt.expected, inv)
				}

				// Verify A*A⁻¹ = I
				product, _ := Multiply(tt.matrix, inv)
				identity := Identity(len(tt.matrix))
				if !matricesAlmostEqual(product, identity, 1e-6) {
					t.Errorf("A*A⁻¹ should equal identity matrix")
				}
			})
		}
	})

	t.Run("Error cases", func(t *testing.T) {
		tests := []struct {
			name     string
			matrix   Matrix[int]
			errorMsg string
		}{
			{
				name: "Singular matrix",
				matrix: Matrix[int]{
					{2, 4},
					{1, 2}, // Row 2 is half of row 1
				},
				errorMsg: "matrix is singular",
			},
			{
				name: "Non-square matrix",
				matrix: Matrix[int]{
					{1, 2, 3},
					{4, 5, 6},
				},
				errorMsg: "cannot invert a non-square matrix",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := Inverse(tt.matrix)
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("expected error message %q, got %q", tt.errorMsg, err.Error())
				}
			})
		}
	})
}

func BenchmarkInverse(b *testing.B) {
	sizes := []int{2, 5, 10, 20}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%dx%d", size, size), func(b *testing.B) {
			m := randomIntMatrix(size, size)
			for b.Loop() {
				_, _ = Inverse(m)
			}
		})
	}
}
