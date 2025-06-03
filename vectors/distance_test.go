package vectors

import (
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	a := Vector[int]{1, 2, 3}
	b := Vector[int]{4, 6, 3}
	dist, err := EuclideanDistance(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if !almostEqual(dist, 5.0, 1e-6) {
		t.Errorf("Expected distance 5.0, got %v", dist)
	}
}
func TestManhattanDistance(t *testing.T) {
	a := Vector[int]{1, 2, 3}
	b := Vector[int]{4, 6, 3}

	dist, err := ManhattanDistance(a, b)
	if err != nil || !almostEqual(dist, 7.0, 1e-6) { // |4-1| + |6-2| + |3-3| = 3+4+0 = 7
		t.Errorf("Expected Manhattan distance 7.0, got %v", dist)
	}
}

func TestChebyshevDistance(t *testing.T) {
	a := Vector[int]{1, 2, 3}
	b := Vector[int]{4, 6, 3}

	dist, err := ChebyshevDistance(a, b)
	if err != nil || !almostEqual(dist, 4.0, 1e-6) { // max(|4-1|, |6-2|, |3-3|) = max(3, 4, 0) = 4
		t.Errorf("Expected Chebyshev distance 4.0, got %v", dist)
	}
}
