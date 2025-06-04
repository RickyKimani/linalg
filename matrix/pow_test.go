package matrix

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	mat := Matrix[int]{
		{3, 8},
		{4, 6},
	}

	// Identity matrix
	identity := Matrix[float64]{
		{1, 0},
		{0, 1},
	}

	matSquared := Matrix[float64]{
		{41, 72},
		{36, 68},
	}

	// Non-square matrix
	nonSquare := Matrix[int]{
		{1, 2, 3},
		{4, 5, 6},
	}

	tests := []struct {
		name    string
		matrix  Matrix[int]
		power   int
		want    Matrix[float64]
		wantErr bool
	}{
		{
			name:   "Power 0",
			matrix: mat,
			power:  0,
			want:   identity,
		},
		{
			name:   "Power 1",
			matrix: mat,
			power:  1,
			want: Matrix[float64]{
				{3, 8},
				{4, 6},
			},
		},
		{
			name:   "Power 2",
			matrix: mat,
			power:  2,
			want:   matSquared,
		},
		{
			name:   "Power 3",
			matrix: mat,
			power:  3,
			want: Matrix[float64]{
				{411, 760},
				{380, 696},
			},
		},
		{
			name: "Negative Power",
			matrix: Matrix[int]{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			power: -1,
			// Inverse of mat (calculated separately)
			want: Identity(3),
		},
		{
			name:    "Non-square Matrix",
			matrix:  nonSquare,
			power:   2,
			wantErr: true,
		},
		{
			name: "Singular Matrix",
			matrix: Matrix[int]{
				{1, 2},
				{2, 4}, // Linearly dependent rows
			},
			power:   -1, // Can't invert singular matrix
			wantErr: true,
		},
		{
			name: "Inconsistent matrix",
			matrix: Matrix[int]{
				{1, 2},
				{3},
			},
			power:   2,
			wantErr: true,
		},
		{
			name:    "Empty Matrix",
			matrix:  make(Matrix[int], 0),
			power:   2,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pow(tt.matrix, tt.power)

			// Check error conditions
			if (err != nil) != tt.wantErr {
				t.Errorf("Pow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return // No need to check result if we expected an error
			}

			// Compare dimensions
			if len(got) != len(tt.want) || len(got[0]) != len(tt.want[0]) {
				t.Errorf("Pow() matrix dimensions mismatch: got %dx%d, want %dx%d",
					len(got), len(got[0]), len(tt.want), len(tt.want[0]))
				return
			}

			// Compare elements with small epsilon for floating point comparison
			epsilon := 1e-12
			for i := range got {
				for j := range got[i] {
					if math.Abs(got[i][j]-tt.want[i][j]) > epsilon {
						t.Errorf("Pow()[%d][%d] = %v, want %v", i, j, got[i][j], tt.want[i][j])
					}
				}
			}
		})
	}
}
