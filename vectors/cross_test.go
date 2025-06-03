package vectors

import (
	"testing"
)

func TestCross(t *testing.T) {
	a := Vector[int]{1, 0, 0}
	b := Vector[int]{0, 1, 0}
	cross, err := Cross(a, b)
	if err != nil {
		t.Fatal(err)
	}
	expected := Vector[float64]{0, 0, 1}
	for i := range cross {
		if !almostEqual(cross[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, cross)
		}
	}
}

func BenchmarkCross(b *testing.B) {
	a := Vector[float64]{1, 2, 3}
	c := Vector[float64]{4, 5, 6}
	for b.Loop() {
		_, _ = Cross(a, c)
	}
}
