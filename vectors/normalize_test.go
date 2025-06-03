package vectors

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	v := Vector[float64]{3, 4, 0}
	norm, err := Normalize(v)
	if err != nil {
		t.Fatal(err)
	}
	if !almostEqual(Magnitude(norm), 1.0, 1e-6) {
		t.Errorf("Expected magnitude 1, got %v", Magnitude(norm))
	}

	_, err = Normalize(Vector[int]{})
	if err == nil {
		t.Errorf("Expected error for zero vector got %v", err)
	}
	_, err = Normalize(Vector[int]{0, 0, 0})
	if err == nil {
		t.Errorf("Expected error for zero vector %v", err)
	}

}

func BenchmarkNormalize(b *testing.B) {
	v := Vector[float64]{3, 4, 5}
	for b.Loop() {
		_, _ = Normalize(v)
	}
}
