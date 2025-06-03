package vectors

import (
	"testing"
)

func TestDot(t *testing.T) {
	a := Vector[float64]{1, 2, 3}
	b := Vector[float64]{4, -5, 6}
	dot, err := Dot(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if !almostEqual(dot, 12, 1e-6) {
		t.Errorf("Expected 12, got %v", dot)
	}

	_, err = Dot(Vector[int]{0, 1}, a)
	if err == nil {
		t.Errorf("expected dimension error got %v", err)
	}

	_, err = Dot(a, Vector[float64]{})
	if err == nil {
		t.Errorf("expected empty vector error got %v", err)
	}

}

func BenchmarkDot(b *testing.B) {
	a := Vector[float64]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := Vector[float64]{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for b.Loop() {
		_, _ = Dot(a, c)
	}
}
