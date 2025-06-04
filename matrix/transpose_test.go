package matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTransposeGeneral(t *testing.T) {
	m := Matrix[int]{
		{1, 2, 3},
		{4, 5, 6},
	}
	want := Matrix[int]{
		{1, 4},
		{2, 5},
		{3, 6},
	}
	got := Transpose(m)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transpose = %v; want %v", got, want)
	}
}

func TestTransposeErr(t *testing.T) {
	result := make(Matrix[int], 0)

	tests := []struct {
		name   string
		matrix Matrix[int]
		want   Matrix[int]
	}{
		{
			name: "inconsistent matrix",
			matrix: Matrix[int]{
				{1, 2},
				{0},
			},
			want: result,
		},

		{
			name:   "empty matrix",
			matrix: make(Matrix[int], 0),
			want:   result,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transpose := Transpose(tt.matrix)
			if !reflect.DeepEqual(tt.want, transpose) {
				t.Fatal("Transpose error")
			}
		})
	}
}

// Benchmark matrix transpose with different sizes
func BenchmarkTranspose(b *testing.B) {
	sizes := []int{100, 500, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%dx%d", size, size), func(b *testing.B) {
			m := randomIntMatrix(size, size)
			b.ResetTimer()
			b.ReportAllocs() // Also measure allocations

			for b.Loop() {
				_ = Transpose(m)
			}
		})
	}
}
