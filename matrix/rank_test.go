package matrix

import "testing"

func TestRank(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[float64]
		want   int
	}{
		{
			name:   "empty matrix",
			matrix: Matrix[float64]{},
			want:   0,
		},
		{
			name: "zero matrix",
			matrix: Matrix[float64]{
				{0, 0},
				{0, 0},
			},
			want: 0,
		},
		{
			name: "identity matrix 3x3",
			matrix: Matrix[float64]{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: 3,
		},
		{
			name: "rank-deficient matrix",
			matrix: Matrix[float64]{
				{1, 2, 3},
				{2, 4, 6},
				{3, 6, 9},
			},
			want: 1,
		},
		{
			name: "full rank rectangular matrix",
			matrix: Matrix[float64]{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rank(tt.matrix)
			if got != tt.want {
				t.Errorf("Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}
