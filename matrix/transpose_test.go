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
