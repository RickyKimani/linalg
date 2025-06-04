package matrix

import (
	"math"
	"testing"
)

const tolerance = 1e-9

func TestDeterminants(t *testing.T) {
	tests := []struct {
		name    string
		matrix  Matrix[int]
		want    float64
		wantErr bool
	}{
		{
			name:    "empty matrix",
			matrix:  make(Matrix[int], 0),
			want:    0,
			wantErr: true,
		},
		{
			name: "improper matrix",
			matrix: Matrix[int]{
				{1, 2},
				{3},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "1x1 matrix",
			matrix: Matrix[int]{
				{5},
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "2x2 standard",
			matrix: Matrix[int]{
				{3, 8},
				{4, 6},
			},
			want:    -14,
			wantErr: false,
		},
		{
			name: "3x3 standard",
			matrix: Matrix[int]{
				{3, 2, 4},
				{2, 0, 2},
				{4, 2, 3},
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "Identity matrix 3x3",
			matrix: Matrix[int]{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Singular matrix (det 0)",
			matrix: Matrix[int]{
				{2, 2},
				{2, 2},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Negative values",
			matrix: Matrix[int]{
				{-2, -1},
				{-3, -4},
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Zeros matrix",
			matrix: Matrix[int]{
				{0, 0},
				{0, 0},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Non-square matrix (2x3)",
			matrix: Matrix[int]{
				{1, 2, 3},
				{4, 5, 6},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Upper triangular 3x3",
			matrix: Matrix[int]{
				{2, -1, 3},
				{0, 4, 5},
				{0, 0, 6},
			},
			want:    48, // product of diagonals
			wantErr: false,
		},
		{
			name: "Lower triangular 3x3",
			matrix: Matrix[int]{
				{2, 0, 0},
				{-1, 4, 0},
				{3, 5, 6},
			},
			want:    48,
			wantErr: false,
		},
		{
			name: "4x4 matrix",
			matrix: Matrix[int]{
				{1, 0, 2, -1},
				{3, 0, 0, 5},
				{2, 1, 4, -3},
				{1, 0, 5, 0},
			},
			want:    30,
			wantErr: false,
		},
		{
			name: "5x5 matrix",
			matrix: Matrix[int]{{2, 3, -1, 3, 3},
				{0, -1, 5, 2, -1},
				{0, 0, 3, 9, 2},
				{0, 0, 0, -1, 3},
				{0, 0, 0, 0, 5},
			},
			want: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Det(tt.matrix)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Det() expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("Det() unexpected error: %v", err)
				return
			}
			if math.Abs(got-tt.want) > tolerance {
				t.Errorf("Det() = %v; want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDet2x2(b *testing.B) {
	m := Matrix[int]{
		{4, 7},
		{2, 6},
	}
	for b.Loop() {
		_, _ = Det(m)
	}
}

func BenchmarkDet10x10(b *testing.B) {
	m := randomIntMatrix(4096, 4096)
	for b.Loop() {
		_, _ = Det(m)
	}
}
