package matrix

import (
	"fmt"
	"testing"
)

// Benchmark matrix inversion with different sizes
func BenchmarkAllocInverse(b *testing.B) {
	sizes := []int{5, 10, 20, 50}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%dx%d", size, size), func(b *testing.B) {
			m := randomIntMatrix(size, size)
			b.ResetTimer() // Don't count setup time
			b.ReportAllocs()

			for b.Loop() {
				inv, _ := Inverse(m)
				// Prevent compiler from optimizing away the result
				b.StopTimer()
				validateInverse(m, inv)
				b.StartTimer()
			}
		})
	}
}
