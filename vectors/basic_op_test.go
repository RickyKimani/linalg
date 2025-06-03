package vectors

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := Vector[int]{1, 2, 3}
	b := Vector[int]{4, 5, 6}
	sum, err := Add(a, b)
	if err != nil {
		t.Fatal(err)
	}
	expected := Vector[float64]{5, 7, 9}
	for i := range sum {
		if !almostEqual(sum[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, sum)
		}
	}
}

func TestSubtract(t *testing.T) {
	a := Vector[int]{4, 5, 6}
	b := Vector[int]{1, 2, 3}
	diff, err := Subtract(a, b)
	if err != nil {
		t.Fatal(err)
	}
	expected := Vector[float64]{3, 3, 3}
	for i := range diff {
		if !almostEqual(diff[i], expected[i], 1e-6) {
			t.Errorf("Expected %v, got %v", expected, diff)
		}
	}
}
